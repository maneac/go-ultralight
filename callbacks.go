package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib/win/x64' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x64
// #cgo windows,i386 LDFLAGS: -L'./SDK/lib/win/x86' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x86
// #cgo linux LDFLAGS: -L'./SDK/bin/linux' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/linux
// #cgo darwin LDFLAGS: -L'./SDK/bin/mac' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/mac
// #include "ultralight.c"
import "C"
import (
	"strings"
	"unsafe"
)

var appUpdate func()
var viewChangeTitle = make(map[C.ULView]func(string))
var viewChangeURL = make(map[C.ULView]func(string))
var viewChangeTooltip = make(map[C.ULView]func(string))
var viewChangeCursor = make(map[C.ULView]func(Cursor))
var viewAddConsoleMessage = make(map[C.ULView]func(MessageSource, MessageLevel, string, uint, uint, string))
var viewBeginLoading = make(map[C.ULView]func())
var viewFinishLoading = make(map[C.ULView]func())
var viewUpdateHistory = make(map[C.ULView]func())
var viewDOMReady = make(map[C.ULView]func())
var winClose func()
var winResize func(int, int)

/******************************************************************************
 * App
 *****************************************************************************/

// SetUpdateCallback .
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

/******************************************************************************
 * View
 *****************************************************************************/

// SetChangeTitleCallback .
func (view *View) SetChangeTitleCallback(callFunc func(string)) {
	viewChangeTitle[view.v] = callFunc
	C.setViewChangeTitleCallback(view.v)
}

//export viewChangeTitleFunction
func viewChangeTitleFunction(_ unsafe.Pointer, v C.ULView, title C.ULString) {
	if viewChangeTitle != nil {
		viewChangeTitle[v](ulStrToStr(title))
	}
}

// SetChangeURLCallback .
func (view *View) SetChangeURLCallback(callFunc func(string)) {
	viewChangeURL[view.v] = callFunc
	C.setViewChangeURLCallback(view.v)
}

//export viewChangeURLFunction
func viewChangeURLFunction(_ unsafe.Pointer, v C.ULView, url C.ULString) {
	if viewChangeURL != nil {
		viewChangeURL[v](ulStrToStr(url))
	}
}

// SetChangeTooltipCallback .
func (view *View) SetChangeTooltipCallback(callFunc func(string)) {
	viewChangeTooltip[view.v] = callFunc
	C.setViewChangeTooltipCallback(view.v)
}

//export viewChangeTooltipFunction
func viewChangeTooltipFunction(_ unsafe.Pointer, v C.ULView, tooltip C.ULString) {
	if viewChangeTooltip != nil {
		viewChangeTooltip[v](ulStrToStr(tooltip))
	}
}

// SetChangeCursorCallback .
func (view *View) SetChangeCursorCallback(callFunc func(Cursor)) {
	viewChangeCursor[view.v] = callFunc
	C.setViewChangeCursorCallback(view.v)
}

//export viewChangeCursorFunction
func viewChangeCursorFunction(_ unsafe.Pointer, v C.ULView, cursor C.ULCursor) {
	if viewChangeCursor != nil {
		viewChangeCursor[v](Cursor(cursor))
	}
}

// SetAddConsoleMessageCallback .
func (view *View) SetAddConsoleMessageCallback(callFunc func(MessageSource, MessageLevel, string, uint, uint, string)) {
	viewAddConsoleMessage[view.v] = callFunc
	C.setViewAddConsoleMessageCallback(view.v)
}

//export viewAddConsoleMessageFunction
func viewAddConsoleMessageFunction(_ unsafe.Pointer, v C.ULView, source C.ULMessageSource,
	level C.ULMessageLevel, message C.ULString, line C.uint, col C.uint, sourceID C.ULString) {
	if viewAddConsoleMessage != nil {
		viewAddConsoleMessage[v](MessageSource(source), MessageLevel(level), ulStrToStr(message),
			uint(line), uint(col), ulStrToStr(sourceID))
	}
}

// SetBeginLoadingCallback .
func (view *View) SetBeginLoadingCallback(callFunc func()) {
	viewBeginLoading[view.v] = callFunc
	C.setViewBeginLoadingCallback(view.v)
}

//export viewBeginLoadingFunction
func viewBeginLoadingFunction(_ unsafe.Pointer, v C.ULView) {
	if viewBeginLoading != nil {
		viewBeginLoading[v]()
	}
}

// SetFinishLoadingCallback .
func (view *View) SetFinishLoadingCallback(callFunc func()) {
	viewFinishLoading[view.v] = callFunc
	C.setViewFinishLoadingCallback(view.v)
}

//export viewFinishLoadingFunction
func viewFinishLoadingFunction(_ unsafe.Pointer, v C.ULView) {
	if viewFinishLoading != nil {
		viewFinishLoading[v]()
	}
}

// SetUpdateHistoryCallback .
func (view *View) SetUpdateHistoryCallback(callFunc func()) {
	viewUpdateHistory[view.v] = callFunc
	C.setViewUpdateHistoryCallback(view.v)
}

//export viewUpdateHistoryFunction
func viewUpdateHistoryFunction(_ unsafe.Pointer, v C.ULView) {
	if viewUpdateHistory != nil {
		viewUpdateHistory[v]()
	}
}

// SetDOMReadyCallback .
func (view *View) SetDOMReadyCallback(callFunc func()) {
	viewDOMReady[view.v] = callFunc
	C.setViewDOMReadyCallback(view.v)
}

//export viewDOMReadyFunction
func viewDOMReadyFunction(_ unsafe.Pointer, v C.ULView) {
	if viewDOMReady != nil {
		viewDOMReady[v]()
	}
}

/******************************************************************************
 * Window
 *****************************************************************************/

// SetCloseCallback .
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

// SetResizeCallback .
func (win *Window) SetResizeCallback(callFunc func(int, int)) {
	winResize = callFunc
	C.setWinResizeCallback(win.w)
}

//export winResizeFunction
func winResizeFunction(_ unsafe.Pointer, width, height C.int) {
	if winResize != nil {
		winResize(int(width), int(height))
	}
}

//export objFunctionCallback
func objFunctionCallback(ctx C.JSContextRef, function C.JSObjectRef, _ C.JSObjectRef,
	argCt C.size_t, parameters *C.JSValueRef, _ *C.JSValueRef) C.JSValueRef {
	args := strings.Split(C.GoString(C.printParams(ctx, parameters, argCt)), "\t\t\t")
	vf := FuncMap[function]
	vf.f(vf.v, args)
	return nil
}
