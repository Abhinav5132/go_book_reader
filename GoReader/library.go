package main

import (
	"GoReader/models"
	"io/fs"
	"log"
	"path/filepath"
	"slices"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm/clause"
)

func(a *App) OpenFolderAndCreateALibrary() string {
	dialogueOptions := runtime.OpenDialogOptions{
		Title: "Select folder",
	}

	root, err := runtime.OpenDirectoryDialog(a.ctx, dialogueOptions)

	if root == "" {
		return "No Directory Selected"
	}
	if err != nil{
		return err.Error()
	}
	folderName := filepath.Base(root)

	library := models.Library{
		Name: folderName,
		Picture: "",
		Path: root,
	}
	err = a.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name :"name"},
			{Name: "path"},
		},
		DoNothing: true,
	}).Create(&library).Error

	if err != nil {
		return "Failed to generate Library"
	}
	
	err = a.TraverseThroughDirectoryAndAddToDb(root, library.ID)
	return "Sucessfully created new library"
} 

func(a* App) TraverseThroughDirectoryAndAddToDb(path string, libraryID uint) error{
	err:= filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		fileType := FileTypeFromPath(path)
		if !slices.Contains(a.AllowedFileTypes, fileType){
			return nil
		}

		book := models.Book{
			Name:     d.Name(),
			Path:     path,
			FileType: FileTypeFromPath(path),
			LibraryID: libraryID,
		}
		err = a.DB.Clauses(clause.OnConflict{
			Columns:[]clause.Column{
				{Name: "name"},
				{Name: "path"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"library_id",
				"file_type",
			}),
		}).Create(&book).Error
		
		if err != nil {
			log.Println("DB error:", err)
		}

		return nil
	})
	return err
}