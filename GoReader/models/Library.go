package models

import "time"

type Library struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"uniqueIndex:uniq_library; not null"`
	Path    string `gorm:"uniqueIndex:uniq_library; not null"`
	Picture []byte `gorm:"not null"`
	PinStatus bool `gorm:"not null; default:false"`
	PinRank uint `gorm:"not null;"`
	LastAccessed time.Time
	Books   []Book 
}
