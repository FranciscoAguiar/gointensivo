package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/FranciscoAguiar/gointensivo/internal/order/infra/database"
	"github.com/FranciscoAguiar/gointensivo/internal/order/usecase"
	"github.com/FranciscoAguiar/gointensivo/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"

	//sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPrice{OrderRepository: repository}

	out := make(chan amqp.Delivery) //CHANNEL
	go rabbitmq.Consume(ch, out)    //T2
	for msg := range out {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}
		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println(outputDTO)
		time.Sleep(550 * time.Millisecond)

		//println(string(msg.Body)) //T1
	}

}
