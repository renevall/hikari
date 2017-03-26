package service

import (
	"github.com/gin-gonic/gin"
)

//LawRouting contains the necessary routes to comunicate via HTTP
type LawRouting struct {
	Indexer
}

//Indexer defines an interface for the Indexing processes
type Indexer interface {
	Add()
	Search()
	Delete()
}

//IndexLaw adds a law to the law index
func (law *LawRouting) IndexLaw(c *gin.Context) {

	//Start index prodecure
	law.Indexer.Add()
}

//SearchLaw searches for the query string on the Law Index
func (law *LawRouting) SearchLaw(c *gin.Context) {

}

//DeleteLaw deletes a record from the Index
func (law *LawRouting) DeleteLaw(c *gin.Context) {

}
