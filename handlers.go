package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func createAddressHandler(api BitcoinAPI) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		addrKeys, err := api.GenAddrKeychain()
		if err != nil {
			fmt.Println("ERR:", err)
			return
		}
		json.NewEncoder(w).Encode(addrKeys)
	}
}
