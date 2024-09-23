package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler { /** Define um middleware que recebe o próximo handler e retorna um handler **/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { /** Cria uma função que lida com as requisições HTTP **/
		w.Header().Set("Content-type", "application/json") /** Define o cabeçalho da resposta HTTP para indicar que o conteúdo é do tipo JSON **/
		next.ServeHTTP(w, r)                               /** Passa a requisição e a resposta para o próximo handler na cadeia **/
	})
}
