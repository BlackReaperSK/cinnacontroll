// server.go

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("server/webserver/aminoapps.com/c/my-hero-academia-brasil-050204/page/item/quem-e-cinnamoroll")))
	http.HandleFunc("/rpc", rpcHandler)

	port := "8080"
	fmt.Printf("Servidor web rodando na porta %s...\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %s\n", err.Error())
		os.Exit(1)
	}
}

func rpcHandler(w http.ResponseWriter, r *http.Request) {
	obtainValue := r.URL.Query().Get("obtain")
	idValue := r.URL.Query().Get("id")
	usernameValue := r.URL.Query().Get("username")

	if obtainValue != "" {
		fmt.Print("############### NOVO COMANDO #################\n")
		fmt.Println("Requisição 'obtain':", obtainValue)
	}
	if idValue != "" {
		fmt.Println("Requisição 'id':", idValue)
	}
	if usernameValue != "" {
		fmt.Println("Requisição 'username':", usernameValue)
	}

	fmt.Println() // Linha em branco para separar as requisições
}
