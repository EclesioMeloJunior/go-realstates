package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Run irá iniciar as rotas do servidor
func Run(port string) *http.Server {
	r := mux.NewRouter()

	apiV1 := r.PathPrefix("/api/v1").Subrouter()

	apiV1.HandleFunc("/realstates", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	addr := fmt.Sprintf("localhost:%s", port)

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
