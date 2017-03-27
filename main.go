package main

import (
	"bitbucket.org/reneval/hikari/service"
	_ "github.com/blevesearch/blevex/lang/es"
)

//TODO: makes this file a cli

func main() {
	hs := service.HikariService{}
	hs.Run()
}
