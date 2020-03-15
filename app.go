package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"

// App wraps the underlying App struct
type App struct {
	a C.ULApp
}

// CreateApp creates an App instance (create only one per application lifetime)
func CreateApp(settings *Settings, conf *Config) *App {
	return &App{C.ulCreateApp(settings.s, conf.c)}
}

// Destroy deletes the App instance
func (a *App) Destroy() {
	C.ulDestroyApp(a.a)
}

// SetWindow sets the main window. This must be called before
// App.Run()
func (a *App) SetWindow(win *Window) {
	C.ulAppSetWindow(a.a, win.w)
}

// GetWindow returns the main Window
func (a *App) GetWindow() *Window {
	return &Window{C.ulAppGetWindow(a.a)}
}

// IsRunning returns whether the App is running
func (a *App) IsRunning() bool {
	return bool(C.ulAppIsRunning(a.a))
}

// GetMainMonitor returns the main Monitor instance
func (a *App) GetMainMonitor() *Monitor {
	return &Monitor{C.ulAppGetMainMonitor(a.a)}
}

// GetRenderer returns the underlying Renderer instance
func (a *App) GetRenderer() *Renderer {
	return &Renderer{C.ulAppGetRenderer(a.a)}
}

// Run executes the main loop. Ensure App.SetWindow() has been called
// prior to calling
func (a *App) Run() {
	C.ulAppRun(a.a)
}

// Quit exits the application
func (a *App) Quit() {
	C.ulAppQuit(a.a)
}
