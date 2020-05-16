package main

import (
	"log"
	"os"

	"github.com/maneac/go-ultralight"
)

type browser struct {
	app    *ultralight.App
	window *ultralight.Window
	ui     *ui
}

func createBrowser() *browser {
	// get the current directory to load assets from
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	settings := ultralight.CreateSettings()
	settings.SetFileSystemPath(path)

	b := browser{}
	b.app = ultralight.CreateApp(settings, ultralight.CreateConfig())
	b.window = ultralight.CreateWindow(b.app.GetMainMonitor(), 1024, 768, false,
		ultralight.WindowResizable|ultralight.WindowTitled|ultralight.WindowMaximizable)
	b.window.SetTitle("Ultralight Sample - Browser")
	b.app.SetWindow(b.window)
	b.ui = createUI(b.window)
	return &b
}

func (b *browser) Run() {
	b.app.Run()
}
