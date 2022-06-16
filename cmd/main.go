package main

import (
	"bufio"
	"challange/app/infrastracture"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

//send console argument for executing command like "go run ./cmd/ seed"
func main() {
	command := os.Args[1]
	switch command {
	//create random amount of users
	case "create_code":
		CreateCode()
	default:
		log.Fatal("Unkown command!")
	}
}

func CreateCode() {
	logger := infrastracture.NewLogger()
	db := infrastracture.NewPgxDB(logger)
	redis := infrastracture.NewRedis()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Create charge code")
	fmt.Println("---------------------")

	fmt.Print("Please enter your charge/discount code: ")
	code, _ := reader.ReadString('\n')
	code = strings.Replace(code, "\n", "", -1)

	fmt.Print("How much is this code credit(in toman): ")
	credit, _ := reader.ReadString('\n')
	credit = strings.Replace(credit, "\n", "", -1)

	fmt.Print("How many users can use this code: ")
	consumerCount, _ := reader.ReadString('\n')
	consumerCount = strings.Replace(consumerCount, "\n", "", -1)

	parameters := []interface{}{
		code, credit, consumerCount,
	}
	rowsAffect, err := db.Exec(context.TODO(),
		"INSERT INTO codes VALUES($1,$2,$3)", parameters)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	if rowsAffect < 1 {
		logger.Error("Can't create code ,maybe its already exist")
		return
	}
	//store key in redis
	err = redis.Set("code_"+code, "ss", 48*time.Hour)
	if err != nil {
		logger.Error("Warning failed to store in redis:" + err.Error())
	}
	fmt.Println("Charge/discount code Created Successfully!")
}
