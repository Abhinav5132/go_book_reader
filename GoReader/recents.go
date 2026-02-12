package main

import (
	"log"
	"os"
	"time"

	"GoReader/models"
	"path/filepath"
)

func AddRecentBookOnFileOpen(path string) (models.Book, error){
	
	fileType := FileTypeFromPath(path)

	file, err := os.Open(path)

	if err != nil {
		return models.Book{}, err
	}
	defer file.Close()
	fileInfo, err:= file.Stat()

	if err != nil {
		log.Println(err)
		return models.Book{}, err // actuall error handel this later 
	}

	book := models.Book{
		Name: fileInfo.Name(),
		Path: path,
		FileType: fileType, // TODO change to have a michelanious franchise later
		LastAccessed: time.Now(), // TODO FIX this will break if user changes time zones 
	}

	return book, nil
}

func (a *App) GetFirstTenRecentBooks() []models.Book{
	var recents []models.Book
	a.DB.Limit(10).Order("last_accessed DESC").Find(&recents)

	return recents
} 

func FileTypeFromPath(path string) string {
	ext := filepath.Ext(path)

	if ext == "" {
		return "unknown"
	}
	return ext[1:]
}