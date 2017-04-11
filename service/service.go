package service

import (
	"log"

	"bitbucket.org/reneval/hikari/indexer"
	"github.com/blevesearch/bleve"
	"gopkg.in/gin-gonic/gin.v1"
)

type HikariService struct {
}

//Run sets up our router and inject dependencies
func (s *HikariService) Run() {

	r := gin.Default()
	lawIndexer := indexer.LawIndex{}
	lawIndex, err := bleve.Open("testlaws.bleve")
	if err == bleve.ErrorIndexPathDoesNotExist {
		log.Printf("Creating new index...")
		indexMapping, err := indexer.BuildIndexMapping()
		if err != nil {
			log.Fatal(err)
		}
		lawIndex, err = bleve.New("testlaws.bleve", indexMapping)
	}

	lawRouting := LawRouting{&lawIndexer, lawIndex}

	// to start out, we'll build the ability to add, get, and delete a todo
	r.POST("/law", lawRouting.IndexLaw)
	r.GET("/law/search", lawRouting.SearchLaw)
	r.DELETE("/todo/:id", lawRouting.IndexLaw)

	// we'll pass in configuration later
	r.Run(":8585")
}
