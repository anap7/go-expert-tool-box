package main

import (
	"context"
	"fmt"
)

/*
	Passando dados com contextos

	Vamos ver um exemplo de como passar informações para
	outras áreas da aplicação via contexto
*/
func main() {
	ctx := context.Background()
	//O WithValue espera o seu contexto, a chave (token) e o valor (senha)
	ctx = context.WithValue(ctx, "token", "senha")
	//Chamando a função e passando o contexto
	bookHotel(ctx)
}

/*Por convenção da comunidade: Sempre que criar uma função que recebe um contexto, priorize sempre 
passar o contexto primeiro no primeiro parâmetro */
func bookHotel(ctx context.Context) {
	//Pegando o valor do contexto via chave do contexto
	token := ctx.Value("token")
	fmt.Println(token)
}