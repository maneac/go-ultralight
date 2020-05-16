package main

import (
	"fmt"
	"math"

	"github.com/maneac/go-ultralight"
)

const html = `<html>
	<body>
		<script>
		function getString() {
			var returned = goGetString();
			document.getElementById('stringType').innerHTML = "Type: " + typeof returned;
			document.getElementById('stringVal').innerHTML = returned;
		}
		function getBool() {
			var returned = goGetBool();
			document.getElementById('boolType').innerHTML = "Type: " + typeof returned;
			document.getElementById('boolVal').innerHTML = returned;
		}
		function getNumber() {
			var returned = goGetNumber();
			document.getElementById('numType').innerHTML = "Type: " + typeof returned;
			document.getElementById('numVal').innerHTML = returned;
		}
		function getJSON() {
			var returned = goGetJSON();
			document.getElementById('jsonType').innerHTML = "Type: " + typeof returned;
			document.getElementById('jsonVal').innerHTML = returned.message + ": " + returned.value;
		}
		function marshalJSON() {
			var returned = goMarshal(document.getElementById('message').value, document.getElementById('value').value);
			document.getElementById('customJsonValue').innerHTML = returned.message + ": " + returned.value;
		}
		</script>
		<button onclick="getString()">Get a string from Go</button>
		<div id="stringType"></div>
		<div id="stringVal"></div>
		<br>
		<br>
		<button onclick="getBool()">Or maybe a boolean</button>
		<div id="boolType"></div>
		<div id="boolVal"></div>
		<br>
		<br>
		<button onclick="getNumber()">Or even a number</button>
		<div id="numType"></div>
		<div id="numVal"></div>
		<br>
		<br>
		<button onclick="getJSON()">We can even get parsed JSON</button>
		<div id="jsonType"></div>
		<div id="jsonVal"></div>
		<br>
		<br>
		<form>
			Message: <input type="text" id="message"><br>
			Value: <input type="text" id="value"><br><br>
			<button onclick="marshalJSON()">Pass values both ways</button>
		</form>
		<div id="customJsonValue"></div>
	</body>
</html>`

func main() {
	// initialse app
	settings := ultralight.CreateSettings()
	config := ultralight.CreateConfig()
	app := ultralight.CreateApp(settings, config)

	// initialise window
	window := ultralight.CreateWindow(app.GetMainMonitor(), 500, 800, false, ultralight.WindowTitled)
	window.SetTitle("Bonus Tutorial - Returning Values From Go to JavaScript")
	app.SetWindow(window)

	// create overlay and load assets
	overlay := ultralight.CreateOverlay(window, window.GetWidth(), window.GetHeight(), 0, 0)
	overlay.GetView().LoadHTML(html)

	// bind the JavaScript calls
	overlay.GetView().BindJSCallback("goGetString", func(v *ultralight.View, params []string) *ultralight.JSValue {
		return v.JSBindString("a Go string")
	})

	overlay.GetView().BindJSCallback("goGetBool", func(v *ultralight.View, params []string) *ultralight.JSValue {
		return v.JSBindBool(true)
	})

	overlay.GetView().BindJSCallback("goGetNumber", func(v *ultralight.View, params []string) *ultralight.JSValue {
		return v.JSBindNum(math.Pi)
	})

	count := 0
	overlay.GetView().BindJSCallback("goGetJSON", func(v *ultralight.View, params []string) *ultralight.JSValue {
		count++
		message := fmt.Sprintf(`{"message": "Number of calls to this method", "value": %d}`, count)
		return v.JSBindJSON(message)
	})

	overlay.GetView().BindJSCallback("goMarshal", func(v *ultralight.View, params []string) *ultralight.JSValue {
		if len(params) != 2 {
			return nil
		}
		message := fmt.Sprintf(`{"message": %q, "value": %q}`, params[0], params[1])
		return v.JSBindJSON(message)
	})

	// run the app
	app.Run()
}
