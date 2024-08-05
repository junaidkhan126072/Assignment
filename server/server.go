// server/server.go
package server

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"

	"sync"

	"github.com/assignment/command"
	"github.com/assignment/orderedmap"
	"github.com/assignment/utils"
)

// Server represents the server structure
type Server struct {
	orderedMap *orderedmap.OrderedMap
	wg         sync.WaitGroup
}

// NewServer creates and returns a new Server instance
func NewServer() *Server {
	return &Server{
		orderedMap: orderedmap.NewOrderedMap(),
	}
}

// handleAddItem adds an item to the ordered map
func (s *Server) handleAddItem(key, value string) {
	s.orderedMap.Add(key, value)
}

// handleDeleteItem deletes an item from the ordered map
func (s *Server) handleDeleteItem(key string) {
	s.orderedMap.Delete(key)
}

// handleGetItem retrieves an item from the ordered map
func (s *Server) handleGetItem(key string) {
	value, exists := s.orderedMap.Get(key)
	if exists {
		utils.AppendToFile(fmt.Sprintf("Key: %s, Value: %s\n", key, value))
	} else {
		utils.AppendToFile(fmt.Sprintf("Key: %s not found\n", key))
	}
}

// handleGetAllItems retrieves all items from the ordered map
func (s *Server) handleGetAllItems() {
	items := s.orderedMap.GetAll()
	for _, item := range items {
		utils.AppendToFile(fmt.Sprintf("Key: %s, Value: %s\n", item.Key, item.Val))
	}
}

// processCommand processes a command based on its action
func (s *Server) processCommand(cmd command.Command) {
	defer s.wg.Done()
	switch cmd.Action {
	case "addItem":
		s.handleAddItem(cmd.Key, cmd.Value)
	case "deleteItem":
		s.handleDeleteItem(cmd.Key)
	case "getItem":
		s.handleGetItem(cmd.Key)
	case "getAllItems":
		s.handleGetAllItems()
	default:
		log.Printf("Unknown command: %s\n", cmd.Action)
	}
}

// Run starts the server and listens for messages from RabbitMQ
func (s *Server) Run() {
	conn, err := amqp.Dial(utils.RabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		utils.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var cmd command.Command
			if err := json.Unmarshal(d.Body, &cmd); err != nil {
				log.Printf("Failed to unmarshal command: %v\n", err)
				continue
			}
			s.wg.Add(1)
			go s.processCommand(cmd)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	s.wg.Wait()
}
