package models

import "time"

type Book struct {
	Id uint `gorm:"primaryKey"`
	Name string `gorm:"unique; not null"`
	Path string `gorm:"unique; not null"`
	FileType string `gorm:"not null"`
	LastAccessed time.Time `gorm:"not null"`
	Franchises []Franchise `gorm:"many2many:franchise_books;"`
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