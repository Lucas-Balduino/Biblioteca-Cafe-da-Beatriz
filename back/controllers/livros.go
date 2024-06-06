package controllers

import (
	"biblioteca/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseGlob("ARRUMAR O CAMINHO DO TEMPLATE")) //TODO: ARRUMAR O CAMINHO DO TEMPLATE

func execTemplate(write http.ResponseWriter, tmplName string, data interface{}) {
	err := tmpl.ExecuteTemplate(write, tmplName, data)
	if err != nil {
		http.Error(write, "Houve um erro ao carregar o template", http.StatusInternalServerError)
		log.Println("Erro ao executar o template: ", err)
		return
	}
}

func Index(write http.ResponseWriter, read *http.Request) {
	livros, err := models.SearchLivro()
	if err != nil {
		http.Error(write, "Houve um erro ao encontrar os livros", http.StatusInternalServerError)
		log.Println("Erro ao encontrar os livros: ", err)
		return
	}
	execTemplate(write, "Index", livros)
}

func Create(write http.ResponseWriter, read *http.Request) {
	execTemplate(write, "Create", nil)
}

func Delete(write http.ResponseWriter, read *http.Request) {
	idLivro := read.URL.Query().Get("id")
	models.DeleteLivro(idLivro)
	http.Redirect(write, read, "/", http.StatusMovedPermanently) // 301 status code
}

func Insert(write http.ResponseWriter, read *http.Request) {
	if read.Method == "POST" {
		nome := read.FormValue("nome")
		autor := read.FormValue("autor")
		quantidade := read.FormValue("quantidade")
		preco := read.FormValue("preco")

		precoConv, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro ao converter o preco", err)
		}

		quantidadeConv, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro ao converter a quantidade", err)
		}

		models.CreateLivro(nome, autor, quantidadeConv, precoConv)
	}

	http.Redirect(write, read, "/", http.StatusMovedPermanently) // 301 status code
}
