package routes

import (
	"net/http"
)

func Rotas() {
	http.HandleFunc("/", index)
	http.HandleFunc("/inserir", inserir)
	http.HandleFunc("/deletar", deletar)
	http.HandleFunc("/editar", editar)
}
