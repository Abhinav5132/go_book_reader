package models

type Library struct {
	ID uint `gorm:"primaryKey"`
	Name string `gorm:"unique; not null"`
	Path string `gorm:"unique; not null"`
	Picture string `gorm:"not null"`
	Books []Book `gorm:"foreignKey:LibraryID"`
}