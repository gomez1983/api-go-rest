package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gomez1983/go-rest-api/database"
	"github.com/gomez1983/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) { /*É a resposta para requisição em Home*/
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) { /** Define a função que recebe um ResponseWriter e um Request **/
	var p []models.Personalidade /** Declara uma variável do tipo slice de Personalidade **/
	database.DB.Find(&p)         /** Busca todas as personalidades no banco de dados e armazena no slice **/
	json.NewEncoder(w).Encode(p) /** Codifica o slice de personalidades em JSON e envia na resposta **/
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) { /** Define a função que recebe um ResponseWriter e um Request **/
	vars := mux.Vars(r)                      /** Pega as variáveis da rota, como o "id" **/
	id := vars["id"]                         /** Atribui o valor da variável "id" da rota à variável id **/
	var personalidade models.Personalidade   /** Declara uma variável do tipo Personalidade **/
	database.DB.First(&personalidade, id)    /** Busca a primeira personalidade com o id informado no banco de dados **/
	json.NewEncoder(w).Encode(personalidade) /** Codifica a personalidade em JSON e envia na resposta **/
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) { /** Define a função que recebe um ResponseWriter e um Request **/
	var novaPersonalidade models.Personalidade         /** Declara uma nova variável do tipo Personalidade **/
	json.NewDecoder(r.Body).Decode(&novaPersonalidade) /** Decodifica o corpo da requisição JSON e atribui à novaPersonalidade **/
	database.DB.Create(&novaPersonalidade)             /** Cria uma nova entrada no banco de dados com os dados de novaPersonalidade **/
	json.NewEncoder(w).Encode(novaPersonalidade)       /** Codifica a nova personalidade em JSON e envia como resposta **/
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) { /** Define a função que recebe um ResponseWriter e um Request **/
	vars := mux.Vars(r)                      /** Pega as variáveis da rota, como o "id" **/
	id := vars["id"]                         /** Atribui o valor da variável "id" da rota à variável id **/
	var personalidade models.Personalidade   /** Declara uma variável do tipo Personalidade **/
	database.DB.Delete(&personalidade, id)   /** Deleta a personalidade do banco de dados com o ID fornecido **/
	json.NewEncoder(w).Encode(personalidade) /** Codifica a personalidade (mesmo deletada) em JSON e envia como resposta **/
}
