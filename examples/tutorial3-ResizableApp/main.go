package main

import "github.com/maneac/go-ultralight"

// Golang implementation of Ultralight sample Tutorial 3

type myApp struct {
	app     *ultralight.App
	window  *ultralight.Window
	overlay *ultralight.Overlay
}

func createMyApp() *myApp {
	myApp := &myApp{}
	// Create the app instance
	settings := ultralight.CreateSettings()
	config := ultralight.CreateConfig()
	myApp.app = ultralight.CreateApp(settings, config)

	// Create the window for the app
	myApp.window = ultralight.CreateWindow(myApp.app.GetMainMonitor(),
		300, 300, false, ultralight.WindowTitled|ultralight.WindowResizable)

	// Set the title of the window
	myApp.window.SetTitle("Tutorial 3 - Resize Me!")

	// Bind the window to the app - MUST be done before creating
	//     overlays or running the app
	myApp.app.SetWindow(myApp.window)

	// Creates an overlay with the same dimensions as the window
	myApp.overlay = ultralight.CreateOverlay(myApp.window, myApp.window.GetWidth(),
		myApp.window.GetHeight(), 0, 0)

	// Loads the HTML string into the overlay's View
	myApp.overlay.GetView().LoadHTML("<center>Hello world!</center>")

	// Resizes the Overlay to the dimensions of the Window when
	//     resized
	myApp.window.SetResizeCallback(func(width, height uint) {
		myApp.overlay.Resize(width, height)
	})

	return myApp
}

func (myApp *myApp) run() {
	myApp.app.Run()
}

func main() {
	app := createMyApp()
	app.run()
}
