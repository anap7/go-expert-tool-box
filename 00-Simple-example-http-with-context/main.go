package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

/*
	Trabalhando com HTTP usando Contextos

	O que é o pacote de contexto do Go e o que faz?
		Esse pacote de contexto permite que passamos as informações dele
		para diversas chamadas no nosso sistema e com a opção de fazer
		com que esses contextos sejam cancelados

	Exemplo: Fazemos uma requisição http e ela vai demorar um pouco para finalizar e
	enquanto estamos fazendo essa requisição http o meu sistema está realizando uma conta
	e dependendo do resultado da criação dessa conta não vou precisar mais do resultado da
	requisição. Nisso temos a opção de continuar aguardando nossa requisição http para pegar
	resultado ou fazer o sistema perceber que não é necessário mais a resposta da requisição
	e interromper.const

	Ou seja, a ideia do contexto é permitir que os nosso programas interropam uma execução
	de algo que não é mais necessário e acordo com o contexto da sua aplicação. E um dos pontos
	que conseguimos trabalhar com contexto é utilizando o tempo.
*/

func main() {
	//Iniciando um contexto
	ctx := context.Background()
	//Criando um contexto com timeout de 1 segundo
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
	//Cancelando o context - O contexto pode ser cancelado a partir dessa função cancel ou o timeout definido acima
	defer cancel()

	//Atribuido nosso contexto uma request
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}