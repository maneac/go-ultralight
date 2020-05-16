package ultralight_test

import (
	"fmt"

	"github.com/maneac/go-ultralight"
)

func ExampleView_BindJSCallback() {
	// Create the app instance
	config := ultralight.CreateConfig()
	renderer := ultralight.CreateRenderer(config)
	view := ultralight.CreateView(renderer, 300, 300, false)

	// Prints each name passed to the 'hello' JavaScript function to the console
	view.BindJSCallback("hello", func(v *ultralight.View, parameters []string) *ultralight.JSValue {
		for i := range parameters {
			if parameters[i] == "" {
				fmt.Println("Hello mystery person!")
				continue
			}
			fmt.Printf("Hello %s!\n", parameters[i])
		}
		return nil
	})

	isReady := false

	// Execute the JavaScript function when the page has finished loading
	view.SetDOMReadyCallback(func() {
		view.EvaluateScript("hello('John', 'Anita', '', 'Derek')")
		isReady = true
	})

	// Give the View something to load to trigger the above function
	view.LoadHTML("<html></html>")

	// Update the View until the DOM is ready
	for !isReady {
		renderer.Update()
	}

	// Will output:
	// Hello John!
	// Hello Anita!
	// Hello mystery person!
	// Hello Derek!
}
