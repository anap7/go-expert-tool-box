package main

import (
	"html/template"
	"os"
)

/*
	Introdução
	
	Em algum momento você já construiu uma aplicação dinâmica? Exemplo, precisar
	gerar um HTML dinâmico com os valores do banco de dados ou um e-mail com um 
	código dinâmico 

	O Golang possui um sistema de templates embutidos feitos para isso. Iremos fazer vários
	exemplos de templates, nesse arquivo iremos trabalhar com: HTML Templates
*/

type Course struct {
	Name string
	Workload int
}

func main() {
	course := Course{"Go", 40}

	//Criado um template de HTML
	tmp := template.New("CourseTemplate")

	//O .{{Variavel}} são as variáveis para fazer o pare
	tmp, _ = tmp.Parse("Course: {{.Name}} - Workload: {{.Workload}}")
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}