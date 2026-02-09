package models

type Franchise struct {
	Id uint `gorm:"primaryKey"`
	FranchiseName string `gorm:"unique;not null"`
	FranchiseImage string `gorm:"unique;not null"`
	Books []Book `gorm:"many2many:franchise_books;"`
}

/*
CREATE IF NOT EXISTS franchise_book{
	book_id not null,
	fracnchise_id not null,
	foreign KEY (book_id) REFERNCES book(id)
	foreign KEY (franchise_id) REFERNCES book(franchise)
	
}
*/
