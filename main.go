// main.go
package main

import (
	"log"

	"github.com/assignment/server"
)

func main() {
	server := server.NewServer()
	server.Run()
	log.Println("Server is running...")
}
