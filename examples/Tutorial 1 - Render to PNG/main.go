package main

import "github.com/maneac/go-ultralight"

// Golang implementation of Ultralight sample Tutorial 1

type MyApp struct {
	renderer *ultralight.Renderer
	view     *ultralight.View
	done     bool
}

func createMyApp() *MyApp {
	myApp := &MyApp{}
	config := ultralight.CreateConfig()
	myApp.renderer = config.CreateRenderer()
	myApp.view = myApp.renderer.CreateView(200, 200, false)
	myApp.view.LoadHTML("<h1>Hello</h1><p>Welcome to Ultralight!</p>")

	// Writes the rendered page to a PNG file
	myApp.view.SetFinishLoadingCallback(func() {
		myApp.renderer.Render()
		myApp.view.GetBitmap().WritePNG("result.png")
		myApp.done = true
	})

	return myApp
}

func (myApp *MyApp) Run() {
	// Continually updates the renderer until the page
	//     loads, calling the PNG writing function
	for !myApp.done {
		myApp.renderer.Update()
	}
}

func main() {
	app := createMyApp()
	app.Run()
}
