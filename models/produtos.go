package models

import db "localhost.com/db"

type Produto struct{
	Id int
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

func EditaProduto(id string) Produto{
	db := db.ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)

	if err != nil{
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil{
			panic(err.Error())
		}
		
		produtoParaAtualizar.Id = id;
		produtoParaAtualizar.Nome = nome;
		produtoParaAtualizar.Descricao = descricao;
		produtoParaAtualizar.Preco = preco;
		produtoParaAtualizar.Quantidade = quantidade;
	}
	defer db.Close()
	return produtoParaAtualizar
}

func BuscaTodosOsProdutos() []Produto{
	db := db.ConectaComBancoDeDados()

	selectProdutos, err := db.Query("Select * From produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int){
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil{
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string){
	db := db.ConectaComBancoDeDados()

	deletarProduto, err := db.Prepare("DELETE FROM PRODUTOS WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)
	defer db.Close()
}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int){
	db := db.ConectaComBancoDeDados()

	produtoAtualiza, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil{
		panic(err.Error())
	}
	produtoAtualiza.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}