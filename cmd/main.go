package main

import (
	"fmt"
	"log"
	"os"
)

//send console argument for executing command like "go run ./cmd/ seed"
func main() {
	arg := os.Args[len(os.Args)-1]
	switch arg {
	//create random amount of users
	case "seed":
		fmt.Println("seed")
	default:
		log.Fatal("Unkown command!")
	}
}
