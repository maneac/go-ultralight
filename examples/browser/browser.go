package main

import "github.com/maneac/go-ultralight"

type browser struct {
	app    *ultralight.App
	window *ultralight.Window
	ui     *ui
}

func createBrowser() *browser {
	b := browser{}
	b.app = ultralight.CreateApp(ultralight.CreateSettings(), ultralight.CreateConfig())
	b.window = ultralight.CreateWindow(b.app.GetMainMonitor(), 1024, 768, false, ultralight.WindowResizable|ultralight.WindowTitled|ultralight.WindowMaximizable)
	b.window.SetTitle("Ultralight Sample - Browser")
	b.app.SetWindow(b.window)
	b.ui = createUI(b.window)
	return &b
}

func (b *browser) Run() {
	b.app.Run()
}
