package main

import (
	"log"
	"os"
	"time"

	"GoReader/models"
	"path/filepath"
)

func (a *App) AddRecentBookOnFileOpen(path string){
	
	fileType := FileTypeFromPath(path)

	file, err := os.Open(path)

	if err != nil {
		return
	}
	defer file.Close()
	fileInfo, err:= file.Stat()

	if err != nil {
		log.Println(err)
		return // actuall error handel this later 
	}

	book := models.Book{
		Name: fileInfo.Name(),
		Path: path,
		FileType: fileType, // TODO change to have a michelanious franchise later
		LastAccessed: time.Now(), // TODO FIX this will break if user changes time zones 
	}

	a.DB.Save(book)
}

func FileTypeFromPath(path string) string {
	ext := filepath.Ext(path)

	if ext == "" {
		return "unknown"
	}
	return ext[1:]
}