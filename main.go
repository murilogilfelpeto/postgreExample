package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/murilogilfelpeto/postgreExample/configs"
	"github.com/murilogilfelpeto/postgreExample/handlers"
	"net/http"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Post("/todos", handlers.Create)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetPort), router)
}
