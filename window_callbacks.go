package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #include <ultralight.h>
import "C"

import "unsafe"

var winClose func()
var winResize func(uint, uint)

// SetCloseCallback executes the specified function when the Window closes
func (win *Window) SetCloseCallback(callFunc func()) {
	winClose = callFunc
	C.setWinCloseCallback(win.w)
}

//export winCloseFunction
func winCloseFunction(_ unsafe.Pointer) {
	if winClose != nil {
		winClose()
	}
}

// SetResizeCallback executes the specified function when the Window is
// resized
func (win *Window) SetResizeCallback(callFunc func(width uint, height uint)) {
	winResize = callFunc
	C.setWinResizeCallback(win.w)
}

//export winResizeFunction
func winResizeFunction(_ unsafe.Pointer, width, height C.uint) {
	if winResize != nil {
		winResize(uint(width), uint(height))
	}
}
