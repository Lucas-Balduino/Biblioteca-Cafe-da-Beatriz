package main
import (
	"aplicacao-web/routes"
	"net/http"
)

func main(){
	routes.Rotas()
	http.ListenAndServe("https://lucas-balduino.github.io/Biblioteca-Cafe-da-Beatriz/", nil)
}

