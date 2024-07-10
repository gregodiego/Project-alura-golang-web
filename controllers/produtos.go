package controllers

import (
	models "localhost.com/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "index", produtos)
}

func New(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		
		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil{
			log.Println("Erro na conversão do preço, err")
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil{
			log.Println("Erro na conversão da quantidade", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request){
	idDoProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idDoProduto)

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request){
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "edit", produto)
}


func Update(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil{
			log.Println("Erro na conversão do ID para int", err)
		}
		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil{
			log.Println("Erro na conversão do preço", err)
		}
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err!= nil{
			log.Println("Erro ao converter quantidade para inteiro", err)
		}

		models.AtualizaProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301 )
}
