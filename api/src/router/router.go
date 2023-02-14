package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}