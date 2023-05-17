package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Contextos no Golang:
		O contexto serve para controlar operações de uma aplicação e verificar
		se essa operação é útil ou não para a aplicação

		O contexto também é utilizado para guardar informações para que seja possivel
		resgatar essas informações em outra área da aplicação

		Em headers http ou chamadas de filas muita das vezes existem informações nesse
		contexto, exemplo, um co-relation id que nos ajuda a trabalhar com tracing
		distribuido...

	Exemplo:
		Vamos simular uma reserva de hotel cujo tempo de reserva possui um limite
		e nisso vamos aplicar nosso contexto
*/

func main() {
	//Iniciando o contexto
	ctx := context.Background()

	//Aplicando regras ao contexto com timeout de 3 segundos
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	/*O resultado é retornar "Hotel booked" porque
	a reserva será concluída em 1 segundo e o timeout do 
	nosso contexto são 3 segundos*/
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	//Caso o contexto esteja Done
	case <-ctx.Done(): 
		fmt.Println("Hotel booking cancelled. Timeout reached")
		return
	//Caso passe 5 segundos e não foi cancelado o nosso contexto iremos dar sucesso com a reserva
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked")
	}
}