package main

import "github.com/maneac/go-ultralight"

// Golang implementation of Ultralight sample Tutorial 2

func main() {

	// Create the app instance
	config := ultralight.CreateConfig()
	app := config.CreateApp()

	// Create the window for the app
	window := ultralight.CreateWindow(app.GetMainMonitor(), 300, 300, false, ultralight.WindowTitled)

	// Set the title of the window
	window.SetTitle("Tutorial 2 - Basic App")

	// Bind the window to the app - MUST be done before creating
	//     overlays or running the app
	app.SetWindow(window)

	// Creates an overlay with the same dimensions as the window
	overlay := window.CreateOverlay(window.GetWidth(), window.GetHeight(), 0, 0)

	// Loads the HTML string into the overlay's View
	overlay.GetView().LoadHTML("<center>Hello world!</center>")

	// Runs the app
	app.Run()
}
