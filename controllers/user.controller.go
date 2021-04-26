package controllers

import (
	"net/http"
	"strconv"

	"github.com/hilbertpurba/PBP/tugas-crud-echo/models"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	result, err := models.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func InsertNewUser(c echo.Context) error {
	name := c.FormValue("name")
	age := c.FormValue("age")
	address := c.FormValue("address")

	atoiAge, err := strconv.Atoi(age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.InsertNewUser(name, atoiAge, address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	age := c.FormValue("age")
	address := c.FormValue("address")

	atoiId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	atoiAge, err := strconv.Atoi(age)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUser(atoiId, name, atoiAge, address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func DeleteUser(c echo.Context) error {
	id := c.FormValue("id")

	atoiId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteUser(atoiId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}
