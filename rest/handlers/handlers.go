package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se obtiene todos los usuarios")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	response := models.GetDefaultResponse()
	user, err := models.GetUser(userID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.NotFound(err.Error())
	} else {
		response.Data = user
	}

	output, _ := json.Marshal(&response)

	fmt.Fprintf(w, string(output))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se crea un usuario")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se actualiza un usuario")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Se elimina un usuario")
}
