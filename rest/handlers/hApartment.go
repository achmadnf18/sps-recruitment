package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"sps-recruitment/rest/models"
	"strconv"
)

type MSX map[string]interface{}

func GetApartment() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetApartment())
	}
}

func CreateApartment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var apartment models.Appartment

		c.Bind(&apartment)

		id := models.InsertApartment(apartment.Name)

		return c.JSON(http.StatusCreated, MSX{
			"created": id,
		})


	}
}

func UpdateApartment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var apartment models.Appartment

		c.Bind(&apartment)

		ra, err := models.UpdateApartment(apartment.ID, apartment.Name)

		if err != nil{
			return err
		}

		return c.JSON(http.StatusCreated, MSX{
			"rows_affected": ra,
		})
	}
}

func DeleteApartment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		ra, err := models.DeleteApartment(int64(id))

		if err != nil{
			return err
		}

		return c.JSON(http.StatusCreated, MSX{
			"rows_affected": ra,
		})
	}
}