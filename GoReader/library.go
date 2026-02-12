package main

import (
	"GoReader/models"
	"io/fs"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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

	a.DB.Save(&library)
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

		book := models.Book{
			Name:     d.Name(),
			Path:     path,
			FileType: FileTypeFromPath(path),
			LibraryID: libraryID,
		}

		return a.DB.Save(&book).Error
	})
	return err
}