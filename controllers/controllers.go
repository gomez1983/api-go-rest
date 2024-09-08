package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) { /*É a resposta para requisição em Home*/
	fmt.Fprint(w, "Home Page")

}
