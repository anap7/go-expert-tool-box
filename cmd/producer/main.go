package main

import (
	"encoding/json"
	"math/rand"
	"module/internal/order/entity"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Pubish(ch *amqp.Channel, order entity.Order) error {
	body, err := json.Marshal(order) //Transformando os dados em um json
	if err != nil {
		return err
	}
	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

//Função para gerar um pedido
func GenerateOrders() entity.Order {
	return entity.Order{
		ID: uuid.New().String(),
		Price: rand.Float64() * 100,
		Tax: rand.Float64() * 10,
	}
}


func main() {
	//Criando a conexão com o rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//Abrindo um canal
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	for i := 0; i < 100000; i++ {
		Pubish(ch, GenerateOrders())
		time.Sleep(300 * time.Millisecond)
	}
}