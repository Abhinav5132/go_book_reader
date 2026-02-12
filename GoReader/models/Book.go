package models

import "time"

type Book struct {
	Id uint `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex:idx_name_path; not null"`
	Path string `gorm:"uniqueIndex:idx_name_path; not null"`
	FileType string `gorm:"not null"`
	LastAccessed time.Time
	Franchises []Franchise `gorm:"many2many:franchise_books;"`

	LibraryID uint
	Library Library
}


/*
Book{
	id int AUTOINCREMENT PRIMARY KEY,
	name TEXT UNIQUE NOT NULL,
	path TEXT UNIQUE NOT NULL,
	type TEXT NOT NULL
	ADD METADATA SUCH AS LAST READ PAGE LATER
}
*/