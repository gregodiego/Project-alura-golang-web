package main

import (
	"net/http"
	routes "localhost.com/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}