package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
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
	//forever := make(chan bool)      //CHANNEL
	go rabbitmq.Consume(ch, out) //T2
	var qtdWorkers = 3
	for i := 1; i <= qtdWorkers; i++ {
		go worker(out, &uc, i) //criando mais threads
	}
	//<-forever
	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		getTotalUC := usecase.GetTotalUseCase{OrderRepository: repository}
		total, err := getTotalUC.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		json.NewEncoder(w).Encode(total)

	})

	http.ListenAndServe(":8080", nil) // criando servidor http
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPrice, worderID int) {
	for msg := range deliveryMessage {
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
		fmt.Printf("Worder %d has processd order %s \n", worderID, outputDTO.ID)
		time.Sleep(1 * time.Second)
	}
}
