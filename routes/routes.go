package routes

import (
	"log"
	"net/http"

	"github.com/gomez1983/go-rest-api/controllers"
)

func HandleRequest() { /* Lida com todas as requisições em "/". Neste caso, "/home"*/
	http.HandleFunc("/", controllers.Home) /*Indica a função "home" de dentro do pacote controllers.go*/
	log.Fatal(http.ListenAndServe(":8000", nil))
}
