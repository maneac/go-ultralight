// Package ultralight provides unofficial Golang bindings for the Ultralight UI C++ library,
// found at https://ultralig.ht
// 
// Please view the GitHub repository (https://github.com/maneac/go-ultralight) for
// full setup instructions
package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib/win/x64' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x64
// #cgo windows,i386 LDFLAGS: -L'./SDK/lib/win/x86' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x86
// #cgo linux LDFLAGS: -L'./SDK/bin/linux' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/linux
// #cgo darwin LDFLAGS: -L'./SDK/bin/mac' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/mac
// #include "ultralight.c"
import "C"

func init() {
	FuncMap = make(map[C.JSObjectRef]ViewFunc)
	viewChangeTitle = make(map[C.ULView]func(string))
	viewChangeURL = make(map[C.ULView]func(string))
	viewChangeTooltip = make(map[C.ULView]func(string))
	viewChangeCursor = make(map[C.ULView]func(Cursor))
	viewAddConsoleMessage = make(map[C.ULView]func(MessageSource, MessageLevel, string, uint, uint, string))
	viewBeginLoading = make(map[C.ULView]func())
	viewFinishLoading = make(map[C.ULView]func())
	viewUpdateHistory = make(map[C.ULView]func())
	viewDOMReady = make(map[C.ULView]func())
}