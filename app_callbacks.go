package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"

import "unsafe"

var appUpdate func()

// SetUpdateCallback executes the specified function whenever the App updates
func (a *App) SetUpdateCallback(callFunc func()) {
	appUpdate = callFunc
	C.setAppUpdateCallback(a.a)
}

//export appUpdateFunction
func appUpdateFunction(_ unsafe.Pointer) {
	if appUpdate != nil {
		appUpdate()
	}
}
