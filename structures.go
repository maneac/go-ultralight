package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib/win/x64' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x64
// #cgo windows,i386 LDFLAGS: -L'./SDK/lib/win/x86' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x86
// #cgo linux LDFLAGS: -L'./SDK/bin/linux' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/linux
// #cgo darwin LDFLAGS: -L'./SDK/bin/mac' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/mac
// #include "ultralight.c"
import "C"

// JSBindFunc .
type JSBindFunc func(view *View, params []string)

// FuncMap .
var FuncMap = make(map[C.JSObjectRef]ViewFunc)

// ViewFunc .
type ViewFunc struct {
	v *View
	f JSBindFunc
}

// Config .
type Config struct {
	c C.ULConfig
}

// Renderer .
type Renderer struct {
	r C.ULRenderer
}

// View .
type View struct {
	v C.ULView
}

// Bitmap .
type Bitmap struct {
	b C.ULBitmap
}

// JSContext .
type JSContext struct {
	jc C.JSContextRef
}

// JSValue .
type JSValue struct {
	jv C.JSValueRef
}

// App .
type App struct {
	a C.ULApp
}

// Window .
type Window struct {
	w C.ULWindow
}

// Monitor .
type Monitor struct {
	m C.ULMonitor
}

// Overlay .
type Overlay struct {
	o C.ULOverlay
}
