package indexer

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

func BuildIndexMapping() (mapping.IndexMapping, error) {
	esFieldMapping := bleve.NewTextFieldMapping()
	esFieldMapping.Analyzer = "es"

	eventMapping := bleve.NewDocumentMapping()
	eventMapping.AddFieldMappingsAt("Name", esFieldMapping)
	eventMapping.AddFieldMappingsAt("Content", esFieldMapping)

	mapping := bleve.NewIndexMapping()
	mapping.DefaultMapping = eventMapping
	return mapping, nil
}
