package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gomez1983/go-rest-api/database"
	"github.com/gomez1983/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) { /*É a resposta para requisição em Home*/
	fmt.Fprint(w, "Home Page")

}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id { /*Converte String em inteiro*/
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
