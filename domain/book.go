package domain

type Book struct {
	ID       int64   `json:"id" db:"book_id"`
	Name     string  `json:"name"`
	Text     string  `json:"text" db:"text"`
	LawID    int64   `json:"lawID" db:"law_id"`
	Titles   []Title `json:"titles"`
	Reviewed bool    `json:"reviewed"`
}

//Type returns the type of bleve document mapping
func (b *Book) Type() string {
	return "book"
}
