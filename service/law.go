package service

import (
	"net/http"

	"bitbucket.org/reneval/hikari/domain"
	"github.com/gin-gonic/gin"
)

//LawRouting contains the necessary routes to comunicate via HTTP
type LawRouting struct {
	Indexer
}

//Indexer defines an interface for the Indexing processes
type Indexer interface {
	Add(domain.Law) error
	Search(string)
	Delete()
}

//IndexLaw adds a law to the law index
func (law *LawRouting) IndexLaw(c *gin.Context) {
	var newLaw domain.Law
	c.BindJSON(&newLaw)
	//Start index prodecure
	err := law.Indexer.Add(newLaw)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "fail", "message": "Could not Index Law"})
	}
	c.JSON(200, gin.H{"code": 200})
}

//SearchLaw searches for the query string on the Law Index
func (law *LawRouting) SearchLaw(c *gin.Context) {
	query := c.Param("query")
	law.Indexer.Search(query)

}

//DeleteLaw deletes a record from the Index
func (law *LawRouting) DeleteLaw(c *gin.Context) {

}
