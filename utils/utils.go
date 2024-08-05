// utils/utils.go
package utils

import (
	"log"
	"os"
)

// Constants for configuration
const (
	RabbitMQURL = "amqp://guest:guest@localhost:5672/"
	QueueName   = "commands"
)

var (
	OutputFile = "output.txt"
)

// AppendToFile appends content to the output file
func AppendToFile(content string) {
	file, err := os.OpenFile(OutputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}
