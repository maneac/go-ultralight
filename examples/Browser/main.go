package main

var globalBrowser *browser

func main() {
	globalBrowser = createBrowser()
	globalBrowser.Run()
}
