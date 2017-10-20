package main

import (
	"bitbucket.org/reneval/hikari/service"
	_ "github.com/blevesearch/bleve/analysis/lang/es"
)

//TODO: makes this file a clis

func main() {
	hs := service.HikariService{}
	hs.Run()
}
