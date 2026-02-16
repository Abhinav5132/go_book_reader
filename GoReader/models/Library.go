package models

type Library struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"uniqueIndex:uniq_library; not null"`
	Path    string `gorm:"uniqueIndex:uniq_library; not null"`
	Picture string `gorm:"not null"`
	Books   []Book 
}
