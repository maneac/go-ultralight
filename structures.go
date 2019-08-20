package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib/win/x64' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo windows,i386 LDFLAGS: -L'./SDK/lib/win/x86' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo linux LDFLAGS: -L'./SDK/bin/linux' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin/mac' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"

// JSBindFunc defines the structure of JavaScript callback functions, where
// 'params' is an array of the parameters passed to the JavaScript function
type JSBindFunc func(view *View, params []string)

type Config struct {
	c C.ULConfig
}

type Renderer struct {
	r C.ULRenderer
}

type View struct {
	v C.ULView
}

type Bitmap struct {
	b C.ULBitmap
}

type JSContext struct {
	jc C.JSContextRef
}

type JSValue struct {
	jv C.JSValueRef
}

type App struct {
	a C.ULApp
}

type Window struct {
	w C.ULWindow
}

type Monitor struct {
	m C.ULMonitor
}

type Overlay struct {
	o C.ULOverlay
}
