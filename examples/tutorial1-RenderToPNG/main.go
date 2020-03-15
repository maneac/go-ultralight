package main

import "github.com/maneac/go-ultralight"

// Golang implementation of Ultralight sample Tutorial 1

type myApp struct {
	renderer *ultralight.Renderer
	view     *ultralight.View
	done     bool
}

func createMyApp() *myApp {
	myApp := &myApp{}
	config := ultralight.CreateConfig()
	myApp.renderer = ultralight.CreateRenderer(config)
	myApp.view = ultralight.CreateView(myApp.renderer, 200, 200, false)
	myApp.view.LoadHTML("<h1>Hello</h1><p>Welcome to Ultralight!</p>")

	// Writes the rendered page to a PNG file
	myApp.view.SetFinishLoadingCallback(func() {
		myApp.renderer.Render()
		myApp.view.GetBitmap().WritePNG("result.png")
		myApp.done = true
	})

	return myApp
}

func (myApp *myApp) run() {
	// Continually updates the renderer until the page
	//     loads, calling the PNG writing function
	for !myApp.done {
		myApp.renderer.Update()
	}
}

func main() {
	app := createMyApp()
	app.run()
}
