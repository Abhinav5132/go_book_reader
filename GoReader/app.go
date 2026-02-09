package main

import (
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

// App struct
type App struct {
	DB *gorm.DB
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp(conn *gorm.DB ) *App {
	return &App{DB: conn}
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
				Pattern: "*.txt",
			},
		},
		CanCreateDirectories: true,
	};
	file, err := runtime.OpenFileDialog(a.ctx, dialogueOptions)

	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	} 

	return string(data), nil
}

