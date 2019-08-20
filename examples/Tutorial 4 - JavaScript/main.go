package main

import "github.com/maneac/go-ultralight"

// Golang implementation of Ultralight sample Tutorial 4

// MyApp handles the overlay logic for a window
type MyApp struct {
	overlay *ultralight.Overlay
}

func createMyApp(window *ultralight.Window) {
	myApp := &MyApp{}

	// Create an Overlay with the same dimensions as the Window
	myApp.overlay = window.CreateOverlay(window.GetWidth(), window.GetHeight(), 0, 0)

	// Load a string of HTML
	myApp.overlay.GetView().LoadHTML(`<html>
	<head>
	  <style type="text/css">
		body { font-family: Arial; text-align: center; }
	  </style>
	</head>
	<body>
		<button onclick="GetMessage()">Get the Secret Message!</button>
		<div id="message"></div>
	</body>
  </html>
	`)

	// Bind the Go function to the JavaScript function 'GetMessage'
	myApp.overlay.GetView().BindJSCallback("GetMessage", func(v *ultralight.View, params []string) {
		v.EvaluateScript("document.getElementById('message').innerHTML = 'Ultralight rocks!';")
	})
}

func main() {
	// Create the App instance
	config := ultralight.CreateConfig()
	app := config.CreateApp()

	// Create a Window
	window := ultralight.CreateWindow(app.GetMainMonitor(), 300, 300, false, ultralight.WindowTitled)

	// Set the title of the Window
	window.SetTitle("Tutorial 4 - JavaScript")

	// Bind the Window to the App instance
	app.SetWindow(window)

	// Creates a MyApp instance to handle the Overlays and JavaScript
	// NOTE: this structure is unnecessary with these bindings
	createMyApp(window)

	// Runs the app
	app.Run()
}
