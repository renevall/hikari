package indexer

import (
	"fmt"
	"strconv"

	"bitbucket.org/reneval/hikari/domain"
	"github.com/blevesearch/bleve"
)

//LawIndex defines the operations for the Law Index
type LawIndex struct{}

//Add adds a Law to the index
func (li *LawIndex) Add(law domain.Law, index bleve.Index) error {

	// Adding Relevant Info to Index

	if len(law.Books) > 0 {
		for _, book := range law.Books {
			err := index.Index("book."+strconv.FormatInt(book.ID, 10), book)
			// prepareItem(book.ID, book.Name, book.Name, "book", law.ID, law.Name))
			// bookID, err := fillBooks(&book, lawID, l.DB)
			if err != nil {
				return err
			}
			if len(book.Titles) > 0 {
				for _, title := range book.Titles {
					err := index.Index("title."+strconv.FormatInt(title.ID, 10), title)
					// prepareItem(title.ID, title.Name, title.Name, "title", law.ID, law.Name))
					// titleID, err := fillTitles(&title, lawID, bookID, l.DB)
					if err != nil {
						return err
					}
					if len(title.Chapters) > 0 {
						for _, chapter := range title.Chapters {
							err := index.Index("chapter."+strconv.FormatInt(chapter.ID, 10), chapter)
							// prepareItem(chapter.ID, chapter.Name, chapter.Name, "chapter", law.ID, law.Name))
							// chapterID, err := fillChapter(&chapter, lawID, titleID, l.DB)
							if err != nil {
								return err
							}
							if len(chapter.Articles) > 0 {

								for _, article := range chapter.Articles {
									// _, err := fillArticle(&article, lawID, chapterID, l.DB, tx)
									err := index.Index("article."+strconv.Itoa(article.ID), article)
									// prepareItem(int64(article.ID), article.Name, article.Text, "article", law.ID, law.Name))
									if err != nil {
										return nil
									}
								}
							}
						}
					}
				}
			}
		}
	} else if len(law.Titles) > 0 {
		for _, title := range law.Titles {
			err := index.Index("title."+strconv.FormatInt(title.ID, 10), title)
			// prepareItem(title.ID, title.Name, title.Name, "title", law.ID, law.Name))
			if err != nil {
				return err
			}
			if len(title.Chapters) > 0 {
				for _, chapter := range title.Chapters {
					err := index.Index("chapter."+strconv.FormatInt(chapter.ID, 10), chapter)
					// prepareItem(chapter.ID, chapter.Name, chapter.Name, "chapter", law.ID, law.Name))
					// chapterID, err := fillChapter(&chapter, lawID, titleID, l.DB)
					if err != nil {
						return err
					}
					if len(chapter.Articles) > 0 {

						for _, article := range chapter.Articles {
							fmt.Println("Indexing: ", article.Name)
							err := index.Index("article."+strconv.Itoa(article.ID), article)
							// prepareItem(int64(article.ID), article.Name, article.Name, "article", law.ID, law.Name))
							// _, err := fillArticle(&article, lawID, chapterID, l.DB, tx)
							if err != nil {
								return nil
							}
						}
					}
				}
			}
		}
	} else if len(law.Articles) > 0 {

		for _, article := range law.Articles {
			// _, err := fillArticle(&article, lawID, 0, l.DB, tx)
			err := index.Index("article."+strconv.Itoa(article.ID), article)
			// prepareItem(int64(article.ID), article.Name, article.Name, "article", law.ID, law.Name))
			if err != nil {
				return nil
			}
		}
	}

	return nil
}

//Search executes a query to the index
func (li *LawIndex) Search(queryString string, index bleve.Index) (*bleve.SearchResult, error) {
	fmt.Println("Query string is:", queryString)

	// query := bleve.NewPhraseQuery(strings.Split(queryString, " "), "Content")
	query := bleve.NewMatchQuery(queryString)
	search := bleve.NewSearchRequest(query)
	search.Highlight = bleve.NewHighlightWithStyle("html")
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(searchResults)
	return searchResults, nil
}

//Delete removes a record from the index
func (li *LawIndex) Delete() {
}

func prepareItem(id int64, name string, content string, rType string, lawID int, lawName string) domain.LawIndex {
	return domain.LawIndex{
		ID:           id,
		Name:         name,
		Content:      content,
		ResourceType: rType,
		LawID:        lawID,
		LawName:      lawName,
	}
}
