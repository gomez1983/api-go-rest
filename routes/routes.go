package routes

import (
	"log"
	"net/http"

	"github.com/gomez1983/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() { /* Lida com todas as requisições em "/". Neste caso, "/home"*/
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home) /*Indica a função "home" de dentro do pacote controllers.go*/
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
