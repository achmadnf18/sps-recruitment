package main

import (
	"github.com/labstack/echo"
	"sps-recruitment/rest/handlers"
)

func main() {
	e := echo.New()

	e.GET("/apartment", handlers.GetApartment())
	e.POST("/apartment", handlers.CreateApartment())
	e.PUT("/apartment", handlers.UpdateApartment())
	e.DELETE("/apartment/:id", handlers.DeleteApartment())

	e.Logger.Fatal(e.Start(":8000"))
}