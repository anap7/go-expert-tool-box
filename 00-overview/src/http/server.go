package http

import (
	"net/http"

	"github.com/anap7/go-expert-tool-box/src/model"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/product", w.getAll)
	e.POST("/product", w.createProduct)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, w.Products)
}

func (w WebServer) createProduct(c echo.Context) error {
	product := model.NewProduct()
	//O Bind funciona como o unmarshall para ler o body e associar a struct
	if err := c.Bind(product); err != nil {
		return err
	}
	w.Products.Add(*product)
	return c.JSON(http.StatusCreated, product)
}
