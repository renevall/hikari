package indexer

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

//BuildIndexMapping sets the mapping for the index
func BuildIndexMapping() (mapping.IndexMapping, error) {
	esFieldMapping := bleve.NewTextFieldMapping()
	esFieldMapping.Analyzer = "es"

	//default
	documentMapping := bleve.NewDocumentMapping()
	documentMapping.AddFieldMappingsAt("Name", esFieldMapping)
	documentMapping.AddFieldMappingsAt("Content", esFieldMapping)

	//law
	lawMapping := bleve.NewDocumentMapping()
	lawMapping.AddFieldMappingsAt("Name", esFieldMapping)

	//article
	articleMapping := bleve.NewDocumentMapping()
	articleMapping.AddFieldMappingsAt("Name", esFieldMapping)
	articleMapping.AddFieldMappingsAt("Content", esFieldMapping)

	//book
	bookMapping := bleve.NewDocumentMapping()
	bookMapping.AddFieldMappingsAt("Name", esFieldMapping)

	//title
	titleMapping := bleve.NewDocumentMapping()
	titleMapping.AddFieldMappingsAt("Name", esFieldMapping)

	//chapter
	chapterMapping := bleve.NewDocumentMapping()
	chapterMapping.AddFieldMappingsAt("Name", esFieldMapping)

	mapping := bleve.NewIndexMapping()
	mapping.DefaultMapping = documentMapping
	mapping.AddDocumentMapping("law", lawMapping)
	mapping.AddDocumentMapping("article", articleMapping)
	mapping.AddDocumentMapping("book", bookMapping)
	mapping.AddDocumentMapping("title", titleMapping)
	mapping.AddDocumentMapping("chapter", chapterMapping)

	return mapping, nil
}
