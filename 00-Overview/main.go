package main

import "github.com/anap7/go-expert-tool-box/model"

//"errors"
//"fmt"

func main() {
	product01 := model.NewProduct()
	product01.Name = "Carrinho"

	product02 := model.NewProduct()
	product02.Name = "Boneca"

	products := model.Products{}
	products.Add(product01)
	products.Add(product02)
}