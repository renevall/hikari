package service

import (
	"net/http"

	"fmt"

	"bitbucket.org/reneval/hikari/domain"
	"github.com/blevesearch/bleve"
	"github.com/gin-gonic/gin"
)

//LawRouting contains the necessary routes to comunicate via HTTP
type LawRouting struct {
	Indexer
	LawIndexBleve bleve.Index
}

//Indexer defines an interface for the Indexing processes
type Indexer interface {
	Add(domain.Law, bleve.Index) error
	Search(string, bleve.Index) (*bleve.SearchResult, error)
	Delete()
}

type Search struct {
	query string
}

//IndexLaw adds a law to the law index
func (law *LawRouting) IndexLaw(c *gin.Context) {
	var newLaw domain.Law
	c.BindJSON(&newLaw)
	//Start index prodecure
	err := law.Indexer.Add(newLaw, law.LawIndexBleve)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Could not Index Law"})
		return
	}
	c.JSON(200, gin.H{"code": 200})
}

//SearchLaw searches for the query string on the Law Index
func (law *LawRouting) SearchLaw(c *gin.Context) {

	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not the expected params"})
		return
	}
	results, err := law.Indexer.Search(query, law.LawIndexBleve)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error",
			"message": "Search service unavalible"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": results.Hits})

}

//DeleteLaw deletes a record from the Index
func (law *LawRouting) DeleteLaw(c *gin.Context) {

}
