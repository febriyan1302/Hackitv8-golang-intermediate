package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	tm := NewTodoManager()

	e := echo.New()

	e.GET("/todo", func(c echo.Context) error {
		todos := tm.GetAll()
		return c.JSON(http.StatusOK, todos)
	})

	e.POST("/todo", func(c echo.Context) error {
		requestBody := CreateTodoRequest{}
		err := c.Bind(&requestBody)
		if err != nil {
			return err
		}

		todo := tm.Create(requestBody)

		return c.JSON(http.StatusCreated, todo)
	})

	e.DELETE("/todo/:id", func(c echo.Context) error {
		id := c.Param("id")

		err := tm.Remove(id)
		if err != nil {
			c.Error(err)
			return err
		}

		return c.String(http.StatusOK, "OK")
	})

	e.PUT("/todo", func(c echo.Context) error {
		requestBody := EditTodoRequest{}
		err := c.Bind(&requestBody)
		if err != nil {
			return err
		}

		errEdit := tm.Edit(requestBody.ID, requestBody.Name)
		if errEdit != nil {
			c.Error(errEdit)
			return errEdit
		}

		return c.String(http.StatusOK, "OK")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
