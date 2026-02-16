package main

import (
	"GoReader/models"
	"errors"
	"io/fs"
	"log"
	"path/filepath"
	"slices"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm/clause"
)

type CreatelibraryResponse struct{
	LibraryId uint
	Response string
}

func(a *App) OpenFolderAndCreateALibrary() CreatelibraryResponse {
	dialogueOptions := runtime.OpenDialogOptions{
		Title: "Select folder",
	}

	root, err := runtime.OpenDirectoryDialog(a.ctx, dialogueOptions)

	if root == "" {
		return CreatelibraryResponse{
			LibraryId: 0,
			Response: "No Directory Selected",
		}
	}
	if err != nil{
		return CreatelibraryResponse{
			LibraryId: 0,
			Response: err.Error(),
		}
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
		return CreatelibraryResponse{
			LibraryId: 0,
			Response: "Failed to generate Library",
		}
	}

	if library.ID != 0{
		err = a.DB.Where("name = ? AND path = ?", folderName, root).First(&library).Error
	}

	if err != nil {
		return CreatelibraryResponse{
			LibraryId: 0,
			Response: "Failed to fetch Library",
		}
	}

	err = a.TraverseThroughDirectoryAndAddToDb(root, library.ID)
	return CreatelibraryResponse{
		LibraryId: library.ID,
		Response: "Sucessfully created new library",
	}
} 

func(a* App) TraverseThroughDirectoryAndAddToDb(path string, LibraryId uint) error{
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
			LibraryID: LibraryId,
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


func(a* App) GetLibrary(id uint) (models.Library, error) {
	if id < 1 {
		return models.Library{}, errors.New("Id cannot be less that 1")
	}
	var library = models.Library{
		ID: id,
	}
	err := a.DB.Preload("books").First(&library).Error

	if err != nil {
		return models.Library{}, err
	}
	
	return library, nil 
}