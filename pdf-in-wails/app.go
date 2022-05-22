package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

func (a *App) OpenAndGetPDFData() ([]byte, error) {
	filters := []runtime.FileFilter{
		{DisplayName: "PDF Documents (*.pdf)", Pattern: "*.pdf"},
	}
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Filters:         filters,
		ShowHiddenFiles: false,
	})
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(filePath)
}
