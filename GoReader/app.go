package main

import (
	"GoReader/models"
	"context"
	"os"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

// App struct
type App struct {
	DB  *gorm.DB
	ctx context.Context
	AllowedFileTypes []string
}

// NewApp creates a new App application struct
func NewApp(conn *gorm.DB, AllowedFileTypes []string) *App {
	return &App{DB: conn, AllowedFileTypes: AllowedFileTypes}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetBookPath() (string, error) { // txt only for now
	dialogueOptions := runtime.OpenDialogOptions{
		Title: "Open Book",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Books",
				Pattern:     "*.txt",
			},
		},
		CanCreateDirectories: true,
	}

	file, err := runtime.OpenFileDialog(a.ctx, dialogueOptions)
	if file == "" {
		return "", nil
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	book, err := AddRecentBookOnFileOpen(file)
	if err != nil {
		return "", err
	}

	a.DB.Save(&book)

	return string(data), nil
}

func (a *App) GetBookFromPath(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	book, err := AddRecentBookOnFileOpen(path)
	if err != nil {
		return "", err
	}

	a.DB.Model(&models.Book{}).Where("id = ?", book.Id).Update("last_accessed", time.Now())

	return string(data), nil
}
