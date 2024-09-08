package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) { /*É a resposta para requisição em Home*/
	fmt.Fprint(w, "Home Page")

}

func HandleRequest() { /* Lida com todas as requisições em "/". Neste caso, "/home"*/
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	HandleRequest() /*Ele executa*/
}
