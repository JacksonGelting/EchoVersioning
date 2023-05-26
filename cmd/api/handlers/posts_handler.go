package handlers

import (
	"github.com/JacksonGelting/EchoVersioning.git/cmd/api/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()

	if err != nil {
		//return c.JSON(http.StatusBadRequest, err)
		err := c.String(http.StatusBadGateway, "unable boobs to process method")
		if err != nil {
			panic(err)
		}
	}

	r := make(map[string]interface{})
	r["status"] = "ok"
	r["data"] = data
	return c.JSON(http.StatusOK, r)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
		//HandleErr(err, c)
	}

	data, err := service.GetById(idx)
	if err != nil {
		HandleErr(err, c)
	}
	r := make(map[string]interface{})
	r["status"] = "ok"
	r["data"] = data
	return c.JSON(http.StatusOK, r)

}

func HandleErr(e error, c echo.Context) {
	err := c.String(http.StatusBadGateway, "unable to process method")
	if err != nil {
		panic(err)
	}
}
