package main

import (
	"fmt"

	"github.com/gomez1983/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest() /*Ele executa*/
}
