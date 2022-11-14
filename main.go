package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type M map[string]interface{}

// ActionIndex ECHO WRAP HANDLER
var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action index"))
}

// ActionHome ECHO WRAP HANDLER
var ActionHome = http.HandlerFunc(
	func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("from action home"))
	},
)

// ActionAbout ECHO WRAP HANDLER
var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("from action about"))
		},
	),
)

// User handling request payload
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

type User2 struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Age   int    `json:"age" validate:"gte=0,lte=80"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// CUSTOM ERROR VALIDATOR
	e.HTTPErrorHandler = func(err error, context echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		context.Logger().Error(report)
		context.JSON(report.Code, report)

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
				}

				break
			}
		}

		context.Logger().Error(report)
		context.JSON(report.Code, report)
	}

	// http://localhost:1323/
	e.GET("/", func(c echo.Context) error {
		data := "hello from /index"
		return c.String(http.StatusOK, data)
	})

	// http://localhost:1323/json
	e.GET("/json", func(c echo.Context) error {
		data := M{"Message": "Hello", "Counter": 2}
		return c.JSON(http.StatusOK, data)
	})

	// http://localhost:1323/page1?name=fajar
	e.GET("/page1", func(c echo.Context) error {
		data := c.QueryParam("name")
		return c.String(http.StatusOK, "Hello, "+data)
	})

	// http://localhost:1323/page2/fajar
	e.GET("/page2/:name", func(c echo.Context) error {
		name := c.Param("name")
		data := fmt.Sprintf("Hello %s", name)

		return c.String(http.StatusOK, data)
	})

	// http://127.0.0.1:1323/page3/fajar/need%20some%20help
	e.GET("/page3/:name/*", func(c echo.Context) error {
		name := c.Param("name")
		message := c.Param("*")

		data := fmt.Sprintf("Hello %s, I have messag for you %s", name, message)

		return c.String(http.StatusOK, data)
	})

	// curl -X POST -F name=fajar -F message=angry http://127.0.0.1:1323/page4
	e.POST("/page4", func(c echo.Context) error {
		name := c.FormValue("name")
		message := c.FormValue("message")

		data := fmt.Sprintf(
			"Hello %s, i have message for you: %s",
			name,
			strings.Replace(message, "/", "", 1),
		)

		return c.String(http.StatusOK, data)
	})

	e.GET("/action_index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	e.GET("/home", echo.WrapHandler(ActionHome))
	e.GET("/about", ActionAbout)

	// handle Static files
	e.Static("/static", "assets")

	e.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	// Request payload with validator
	e.Any("/user_with_validator", func(c echo.Context) error {
		u := new(User2)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err := c.Validate(u); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
