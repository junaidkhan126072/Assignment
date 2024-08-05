package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/assignment/utils"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/mock"
)

// MockChannel mocks the RabbitMQ channel
type MockChannel struct {
	mock.Mock
}

// QueueDeclare is a mocked method
func (m *MockChannel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args map[string]interface{}) (string, error) {
	argsList := m.Called(name, durable, autoDelete, exclusive, noWait, args)
	return argsList.String(0), argsList.Error(1)
}

// Publish is a mocked method
func (m *MockChannel) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	argsList := m.Called(exchange, key, mandatory, immediate, msg)
	return argsList.Error(0)
}

// MockConnection mocks the RabbitMQ connection
type MockConnection struct {
	mock.Mock
}

// Channel is a mocked method
func (m *MockConnection) Channel() (MockChannel, error) {
	argsList := m.Called()
	return argsList.Get(0).(MockChannel), argsList.Error(1)
}

// MockRabbitMQ to mock connection and channel
func MockRabbitMQ(t *testing.T) (*MockConnection, *MockChannel) {
	conn := new(MockConnection)
	ch := new(MockChannel)

	conn.On("Channel").Return(ch, nil)
	ch.On("QueueDeclare", utils.QueueName, false, false, false, false, nil).Return(utils.QueueName, nil)
	ch.On("Publish", "", utils.QueueName, false, false, mock.Anything).Return(nil)

	return conn, ch
}

// TestMainFunction tests the main function
func TestMainFunction(t *testing.T) {
	// Backup and restore original log functions

	resetFlags := func() {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}

	t.Run("No Flags", func(t *testing.T) {
		resetFlags()
		os.Args = []string{"cmd"}
		main()
	})

	t.Run("Only Action Flag", func(t *testing.T) {
		resetFlags()
		os.Args = []string{"cmd", "-action=addItem"}
		main()
	})

	t.Run("Action and Key Flags", func(t *testing.T) {
		resetFlags()
		os.Args = []string{"cmd", "-action=getItem", "-key=exampleKey"}
		main()
	})

	t.Run("All Flags", func(t *testing.T) {
		resetFlags()
		os.Args = []string{"cmd", "-action=addItem", "-key=exampleKey", "-value=exampleValue"}
		main()
	})

	t.Run("File Flag", func(t *testing.T) {
		resetFlags()

		commands := []Command{
			{Action: "addItem", Key: "key1", Value: "value1"},
			{Action: "getItem", Key: "key2"},
		}
		file, err := ioutil.TempFile("", "commands*.json")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(file.Name())

		if err := json.NewEncoder(file).Encode(commands); err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
		file.Close()

		os.Args = []string{"cmd", "-file=" + file.Name()}
		main()
	})

	t.Run("Invalid File Path", func(t *testing.T) {
		resetFlags()
		os.Args = []string{"cmd", "-file=invalid_path.json"}
		main()
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		resetFlags()

		file, err := ioutil.TempFile("", "invalid*.json")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(file.Name())

		file.WriteString(`{"action": "addItem"}`) // Invalid JSON array
		file.Close()

		os.Args = []string{"cmd", "-file=" + file.Name()}
		main()
	})
}
