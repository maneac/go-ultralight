package ultralight_test

import (
	"fmt"
	"github.com/maneac/go-ultralight"
)

func ExampleView_BindJSCallback() {
	// Create the app instance
	config := ultralight.CreateConfig()
	renderer := config.CreateRenderer()
	view := renderer.CreateView(300, 300, false)

	// Prints each name passed to the 'hello' JavaScript function to the console
	view.BindJSCallback("hello", func(v *ultralight.View, parameters []string) {
		// Checks if any parameters were passed to the JavaScript function
		if len(parameters) > 0 {
			for i := range parameters {
				if parameters[i] != "" {
					fmt.Printf("Hello %s!\n", parameters[i])
				} else {
					fmt.Println("Hello mystery person!")
				}
			}
		} else {
			fmt.Println("Hello mystery person!")
		}
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
