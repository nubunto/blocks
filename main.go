package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func main() {
	mux := httprouter.New()
	bc := NewBlockcypherAPI(os.Getenv("BLOCKCYPHER_API_TOKEN"), "btc", "main")
	mux.POST("/addr/", createAddressHandler(bc))

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
