package service

import (
	"bitbucket.org/reneval/hikari/indexer"
	"github.com/gin-gonic/gin"
)

type HikariService struct {
}

//Run sets up our router and inject dependencies
func (s *HikariService) Run() {

	r := gin.Default()
	lawIndexer := indexer.LawIndex{}

	lawRouting := LawRouting{&lawIndexer}

	// to start out, we'll build the ability to add, get, and delete a todo
	r.POST("/law", lawRouting.IndexLaw)
	r.GET("/law/:query", lawRouting.SearchLaw)
	r.DELETE("/todo/:id", lawRouting.IndexLaw)

	// we'll pass in configuration later
	r.Run(":8888")
}