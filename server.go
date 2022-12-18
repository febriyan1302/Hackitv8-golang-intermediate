package main

import (
	"fmt"
	"github.com/crewjam/saml/samlsp"
	"log"
	"net/http"
)

func main() {
	sp, err := newSamlMiddleware()
	if err != nil {
		log.Fatal(err.Error())
	}

	http.Handle("/saml/", sp)

	http.Handle("/index", sp.RequireAccount(http.HandlerFunc(helloHandler)))

	http.Handle("/hello", sp.RequireAccount(http.HandlerFunc(landingHandler)))

	portString := fmt.Sprintf(":%s", webserverPort)
	fmt.Println("Server started at :", portString)
	err = http.ListenAndServe(portString, nil)
	if err != nil {
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello !"))
	if err != nil {
		return
	}
}

func landingHandler(w http.ResponseWriter, r *http.Request) {
	name := samlsp.AttributeFromContext(r.Context(), "displayName")
	_, err := w.Write([]byte(fmt.Sprintf("Welcome, %s", name)))
	if err != nil {
		return
	}
}
