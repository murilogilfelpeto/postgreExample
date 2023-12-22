package handlers

import (
	"encoding/json"
	"github.com/murilogilfelpeto/postgreExample/models"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error while decoding todo: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := models.Save(todo)
	if err != nil {
		log.Printf("Error while saving todo: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	todo.ID = id

	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		return
	}
}
