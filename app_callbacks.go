package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
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
