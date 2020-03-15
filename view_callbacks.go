package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"

import "unsafe"

var viewChangeTitle map[C.ULView]func(string)
var viewChangeURL map[C.ULView]func(string)
var viewChangeTooltip map[C.ULView]func(string)
var viewChangeCursor map[C.ULView]func(Cursor)
var viewAddConsoleMessage map[C.ULView]func(MessageSource, MessageLevel, string, uint, uint, string)
var viewBeginLoading map[C.ULView]func()
var viewFinishLoading map[C.ULView]func()
var viewUpdateHistory map[C.ULView]func()
var viewDOMReady map[C.ULView]func()

func init() {
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

// SetChangeTitleCallback executes the specified function when the page title
// changes
func (view *View) SetChangeTitleCallback(callFunc func(title string)) {
	viewChangeTitle[view.v] = callFunc
	C.setViewChangeTitleCallback(view.v)
}

//export viewChangeTitleFunction
func viewChangeTitleFunction(_ unsafe.Pointer, v C.ULView, title C.ULString) {
	if viewChangeTitle != nil {
		viewChangeTitle[v](ulStrToStr(title))
	}
}

// SetChangeURLCallback executes the specified function when the page URL
// changes
func (view *View) SetChangeURLCallback(callFunc func(url string)) {
	viewChangeURL[view.v] = callFunc
	C.setViewChangeURLCallback(view.v)
}

//export viewChangeURLFunction
func viewChangeURLFunction(_ unsafe.Pointer, v C.ULView, url C.ULString) {
	if viewChangeURL != nil {
		viewChangeURL[v](ulStrToStr(url))
	}
}

// SetChangeTooltipCallback executes the specified function when the tooltip
// changes (usually due to a mouse hover)
func (view *View) SetChangeTooltipCallback(callFunc func(tooltip string)) {
	viewChangeTooltip[view.v] = callFunc
	C.setViewChangeTooltipCallback(view.v)
}

//export viewChangeTooltipFunction
func viewChangeTooltipFunction(_ unsafe.Pointer, v C.ULView, tooltip C.ULString) {
	if viewChangeTooltip != nil {
		viewChangeTooltip[v](ulStrToStr(tooltip))
	}
}

// SetChangeCursorCallback executes the specified function when the mouse cursor
// changes
func (view *View) SetChangeCursorCallback(callFunc func(cursor Cursor)) {
	viewChangeCursor[view.v] = callFunc
	C.setViewChangeCursorCallback(view.v)
}

//export viewChangeCursorFunction
func viewChangeCursorFunction(_ unsafe.Pointer, v C.ULView, cursor C.ULCursor) {
	if viewChangeCursor != nil {
		viewChangeCursor[v](Cursor(cursor))
	}
}

// SetAddConsoleMessageCallback executes the specified function when a message is
// added to the console
func (view *View) SetAddConsoleMessageCallback(callFunc func(
	source MessageSource, level MessageLevel, message string, lineNumber uint, colNumber uint, sourceID string)) {
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

// SetBeginLoadingCallback executes the specified function when the page
// begins loading a new URL
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

// SetFinishLoadingCallback executes the specified function when the page
// finished loading a URL
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

// SetUpdateHistoryCallback executes the specified function when the
// history (back/forward state) is modified
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

// SetDOMReadyCallback executes the specified function when all
// JavaScript has been parsed and the document is ready
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
