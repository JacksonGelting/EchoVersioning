package main

import (
	"github.com/JacksonGelting/EchoVersioning.git/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/health-check", handlers.HealthCheckHandler)
	e.GET("/posts", handlers.PostIndexHandler)
	e.GET("/post/:id", handlers.PostSingleHandler) //not working

	e.Logger.Fatal(e.Start(":1323"))
}
