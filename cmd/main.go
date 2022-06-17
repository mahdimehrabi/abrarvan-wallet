package main

import (
	"bufio"
	"bytes"
	"challange/app/infrastracture"
	"challange/app/models"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
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
	creditStr, _ := reader.ReadString('\n')
	creditStr = strings.Replace(creditStr, "\n", "", -1)
	credit, err := strconv.ParseFloat(creditStr, 64)
	if err != nil {
		logger.Error("credit must be number or float")
	}

	fmt.Print("How many users can use this code: ")
	consumerCountStr, _ := reader.ReadString('\n')
	consumerCountStr = strings.Replace(consumerCountStr, "\n", "", -1)
	consumerCount, err := strconv.Atoi(consumerCountStr)
	if err != nil {
		logger.Error("consumerCount must be number")
	}

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

	codeModel := models.Code{
		Code:          code,
		Credit:        credit,
		ConsumerCount: consumerCount,
	}
	var buff bytes.Buffer

	err = codeModel.ToJSON(&buff)
	if err != nil {
		logger.Error("Can't convert to json")
		return
	}
	//store key in redis
	err = redis.Set("code_"+code, buff.String(), 48*time.Hour)
	if err != nil {
		logger.Error("Warning failed to store in redis:" + err.Error())
	}
	fmt.Println("Charge/discount code Created Successfully!")
}
