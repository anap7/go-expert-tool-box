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
	exemplos de templates, nesse arquivo iremos trabalhar com o template.must, um comando
	onde podemos criar um template em um único comando
*/

type Course struct {
	Name string
	Workload int
}

func main() {
	course := Course{"PHP", 50}
	tmp := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Workload: {{.Workload}}"))

	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}