package domain

// LawIndex Contains the index to be stored no matter the item subclass
// inside a law struct
type LawIndex struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	DocumentType string `json:"type"`
	LawID        int    `json:"law-id"`
	LawName      string `json:"law-name"`
}

func (l *LawIndex) Type() string {
	return l.DocumentType
}
