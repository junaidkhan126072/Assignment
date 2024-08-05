package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/assignment/utils"
	"github.com/streadway/amqp"
)

// Command structure
type Command struct {
	Action string `json:"action"`
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
}

// sendCommand sends commands to RabbitMQ queue
func sendCommand(command Command) {
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

	body, err := json.Marshal(command)
	if err != nil {
		log.Fatalf("Failed to marshal command: %v", err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	log.Printf(" [x] Sent %s\n", body)
}

func main() {
	action := flag.String("action", "", "Available actions (addItem, deleteItem, getItem, getAllItems)")
	key := flag.String("key", "", "Key for the action")
	value := flag.String("value", "", "Value for the action")
	file := flag.String("file", "", "File to read commands from")
	flag.Parse()

	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer f.Close()

		var commands []Command
		if err := json.NewDecoder(f).Decode(&commands); err != nil {
			log.Fatalf("Failed to decode commands: %v", err)
		}

		for _, command := range commands {
			sendCommand(command)
		}
	} else {
		if *action == "" {
			fmt.Println("Action is required")
			return
		}

		command := Command{
			Action: *action,
			Key:    *key,
			Value:  *value,
		}

		sendCommand(command)
	}
}
