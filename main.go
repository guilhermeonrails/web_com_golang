package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alura/routes"
)

func main() {
	routes.CarregaRotas()

	porta := ":8000"
	fmt.Println("Aguardando requisições na porta" + porta)
	log.Fatal(http.ListenAndServe(porta, nil))
}
