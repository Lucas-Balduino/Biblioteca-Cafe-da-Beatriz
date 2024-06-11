package models

import (
	"biblioteca/database"
)

type Livro struct {
	Id         int
	Nome       string
	Autor      string
	Quantidade int
	Preco      float64
}

func SearchLivro() ([]Livro, error) {
	allLivros, err := database.DB.Query("SELECT * FROM livros ORDER BY id ASC")
	if err != nil {
		return nil, err
	}

	livros := []Livro{}

	for allLivros.Next() {
		var l Livro
		err = allLivros.Scan(&l.Id, &l.Nome, &l.Autor, &l.Quantidade, &l.Preco)
		if err != nil { // checa por erro no scan antes de dar append
			return nil, err
		}
		livros = append(livros, l)
	}

	defer allLivros.Close()

	return livros, nil
}

func CreateLivro(nome, autor string, quantidade int, preco float64) error {
	stmt, err := database.DB.Prepare("INSERT INTO livros (nome, descricao, preco, quantidade) VALUES(?,?,?,?)") // stmt vulgo statement eh a query preparada
	if err != nil {
		return err
	}

	_, err = stmt.Exec(nome, autor, quantidade, preco)
	if err != nil {
		return err
	}

	defer stmt.Close()

	return nil
}

func DeleteLivro(id string) error {
	stmt, err := database.DB.Prepare("DELETE FROM livros WHERE id = $1") // stmt vulgo statement eh a query preparada
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
