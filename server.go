package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/labstack/echo/v4"
	"gopkg.in/boj/redistore.v1"
	"html/template"
	"io"
	"math/rand"
	"net/http"
)

const SessionId = "id"

func newRedisStore() *redistore.RediStore {
	store, err := redistore.NewRediStore(10, "tcp", "redis:6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	//defer store.Close()

	return store
}

type M map[string]interface{}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

func main() {
	store := newRedisStore()

	e := echo.New()

	e.Renderer = NewRenderer("templates/*.html", true)

	e.Use(echo.WrapMiddleware(context.ClearHandler))

	e.Any("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.GET("/register", func(c echo.Context) error {
		return c.Render(http.StatusOK, "register.html", nil)
	})

	e.POST("/register", func(c echo.Context) error {
		var user User
		//randomIdUser := 123124123
		user.ID = rand.Intn(99999-10000) + 10000
		user.Username = c.FormValue("username")
		user.FirstName = c.FormValue("FirstName")
		user.LastName = c.FormValue("LastName")
		user.Password = c.FormValue("password")

		register := handleRegister(user)
		if !register {
			return c.String(http.StatusInternalServerError, "Register Failed")
		}

		return c.Redirect(http.StatusTemporaryRedirect, "/")
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", nil)
	})

	e.POST("/login", func(c echo.Context) error {
		var user User
		user.Username = c.FormValue("username")
		user.Password = c.FormValue("password")

		login := handleLogin(user)
		if !login {
			return c.String(http.StatusInternalServerError, "login Failed")
		}

		// SET SESSION
		fmt.Println("Before set session")
		session, _ := store.Get(c.Request(), SessionId)
		session.Values["firstname"] = user.Username
		session.Values["lastname"] = user.LastName
		err := session.Save(c.Request(), c.Response())
		if err != nil {
			fmt.Println("ERROR", err.Error())
			return c.String(http.StatusInternalServerError, "Session Failed")
		}
		fmt.Println("after set session")

		return c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
		//return c.String(http.StatusOK, "/login")
	})

	e.Any("/dashboard", func(c echo.Context) error {
		session, err := store.Get(c.Request(), SessionId)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Dashboard Session", session)

		firstname, ok := session.Values["firstname"].(string)
		if !ok {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		lastname, ok := session.Values["lastname"].(string)
		if !ok {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		data := map[string]string{
			"FirstName": firstname,
			"LastName":  lastname,
		}
		fmt.Println(data)
		return c.Render(http.StatusOK, "dashboard.html", data)
	})

	e.GET("/logout", func(c echo.Context) error {
		session, err := store.Get(c.Request(), SessionId)
		if err != nil {
			fmt.Println(err)
		}
		session.Options.MaxAge = -1
		session.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
