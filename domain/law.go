package domain

import "time"

type Law struct {
	ID           int       `json:"id" db:"law_id"`
	Name         string    `json:"name"`
	ApprovalDate time.Time `json:"approvalDate" db:"approval_date"`
	PublishDate  time.Time `json:"publishDate" db:"publish_date"`
	Journal      string    `json:"journal"`
	Intro        string    `json:"intro"`
	Reviewed     bool      `json:"reviewed"`
	Revision     int       `json:"rev"`
	Books        []Book    `json:"books"`
	Titles       []Title   `json:"titles"`
	Chapters     []Chapter `json:"chapters"`
	Articles     []Article `json:"articles"`
	Init         string    `json:"init"`
}
