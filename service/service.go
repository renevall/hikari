package service

import (
	"fmt"
	"log"
	"time"

	"bitbucket.org/reneval/hikari/indexer"
	"github.com/blevesearch/bleve"
	"github.com/gin-gonic/gin"
	cors "gopkg.in/gin-contrib/cors.v1"
)

type HikariService struct {
}

//Run sets up our router and inject dependencies
func (s *HikariService) Run() {

	r := gin.Default()
	// r.Use(cors.Middleware(cors.Config{
	// 	Origins:         "*",
	// 	Methods:         "GET, PUT, POST, DELETE",
	// 	RequestHeaders:  "Origin, Authorization, Content-Type",
	// 	ExposedHeaders:  "",
	// 	MaxAge:          50 * time.Second,
	// 	Credentials:     true,
	// 	ValidateHeaders: false,
	// }))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	lawIndexer := indexer.LawIndex{}
	lawIndex, err := bleve.Open("testlaws.bleve")
	if err == bleve.ErrorIndexPathDoesNotExist {
		log.Printf("Creating new index...")
		indexMapping, err := indexer.BuildIndexMapping()
		if err != nil {
			log.Fatal(err)
		}
		lawIndex, err = bleve.New("testlaws.bleve", indexMapping)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(lawIndex.Name())
	}

	lawRouting := LawRouting{&lawIndexer, lawIndex}

	// to start out, we'll build the ability to add, get, and delete a todo
	r.POST("/law", lawRouting.IndexLaw)
	r.GET("/law/search", lawRouting.SearchLaw)
	r.DELETE("/todo/:id", lawRouting.IndexLaw)

	// we'll pass in configuration later
	r.Run(":8585")
}
