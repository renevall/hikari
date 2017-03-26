package domain

//Title struc is the model for a law Title
type Title struct {
	ID       int64     `json:"id" db:"title_id"`
	Name     string    `json:"name"`
	Chapters []Chapter `json:"chapters"`
	LawID    int64     `json:"lawID" db:"law_id"`
	BookID   int64     `json:"bookID" db:"book_id"`
	Reviewed bool      `json:"reviewed"`
}
