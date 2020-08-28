package domain

// Writer 著者
type Writer struct {
	ID         int
	FamilyName string
	GivenName  string
}

// Writers 著者
type Writers []Writer
