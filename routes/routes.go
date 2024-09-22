package routes

import (
	"log"
	"net/http"

	"github.com/gomez1983/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() { /** Lida com todas as requisições HTTP **/
	r := mux.NewRouter()                                                                           /** Cria um novo roteador usando o Gorilla Mux **/
	r.HandleFunc("/", controllers.Home)                                                            /** Define a rota "/" e associa à função Home no controllers **/
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")            /** Define a rota "/api/personalidades" para listar todas as personalidades, apenas método GET **/
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")   /** Define a rota "/api/personalidades/{id}" para retornar uma personalidade por ID, apenas método GET **/
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")      /** Define a rota "/api/personalidades" para criar uma nova personalidade, método POST **/
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete") /** Define a rota "/api/personalidades/{id}" para deletar uma personalidade por ID, método DELETE **/
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")        /** Define a rota "/api/personalidades/{id}" para editar uma personalidade por ID, método PUT **/
	log.Fatal(http.ListenAndServe(":8000", r))                                                     /** Inicia o servidor na porta 8000 e encerra se ocorrer um erro fatal **/

}
