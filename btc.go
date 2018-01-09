package main

import "github.com/blockcypher/gobcy"

type BitcoinAPI interface {
	GenAddrKeychain() (gobcy.AddrKeychain, error)
}

type BlockcypherAPI struct{
	gobcy.API
}

func NewBlockcypherAPI(token, coin, chain string) *BlockcypherAPI {
	return &BlockcypherAPI{
		API: gobcy.API{
			Token: token,
			Coin: coin,
			Chain: chain,
		},
	}
}