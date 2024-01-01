package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"module/internal/order/infra/database"
	"module/internal/order/usecase"
	"module/pkg/rabbitmq"
	"time"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	//Instanciando e inserirndo o novo pedido
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	//Chamando a função que criamos
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	//Fechando a conexão com o rabbitmq
	defer ch.Close()
	/*Criando um canal de comunição para consumir as mensagens*/
	out := make(chan amqp.Delivery) //channel
	//Esse canal aguarda uma mensagem a ser recebida, caso não receba ela trava a aplicação e o resto executado
	forever := make(chan bool)
	//Passando a conexão do rabbitmq e o canal para que seja possivel passar as mensagens
	go rabbitmq.Consume(ch, out) //T2

	var qtdWorkers = 5
	for i := 0; i <= qtdWorkers; i++ {
		/*Quando o consumer for acionado e o worker for chamado para ler
		a mensagem e salvar no banco, subiremos 5 workers e 5 serviços recebendo
		mensagens do rabbitmq, enquanto um vai processando o outro vai pegando as mensagens.
		E com isso teremos um balanceador de carga, ao invés de um processo recebendo uma
		requisição*/
		go worker(out, &uc, i)
	}
	<-forever
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerID int) {
	for msg := range deliveryMessage {
		//Cada mensagem que chegar vai ser implementada no DTO
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}

		//println(string(msg.Body)) //T1
		//println(&inputDTO)

		inputDTO.Tax = 10.0
		//Salva a mensagem no banco
		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		//Removendo a mensagem da fila
		msg.Ack(false)
		//Imprimindo o resultado da inserção do banco
		fmt.Printf("Worker %d has processed order %s\n", workerID, outputDTO.ID)
		time.Sleep(500 * time.Millisecond)
	}
} 