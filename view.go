package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"
import "unsafe"

// View wraps the underlying struct
type View struct {
	v C.ULView
}

// JSBindFunc defines the structure of JavaScript callback functions, where
// 'params' is an array of the parameters passed to the JavaScript function
type JSBindFunc func(view *View, params []string) *JSValue

// CreateView creates a View instance with the specified size (in device coordinates)
func CreateView(rend *Renderer, width, height uint, isTransparent bool) *View {
	return &View{C.ulCreateView(rend.r, C.uint(width), C.uint(height), C.bool(isTransparent))}
}

// Destroy deletes the View instance
func (view *View) Destroy() {
	C.ulDestroyView(view.v)
}

// GetURL returns the current URL
func (view *View) GetURL() string {
	return ulStrToStr(C.ulViewGetURL(view.v))
}

// GetTitle returns the current Title
func (view *View) GetTitle() string {
	return ulStrToStr(C.ulViewGetTitle(view.v))
}

// IsLoading checks if the main frame is loading
func (view *View) IsLoading() bool {
	return bool(C.ulViewIsLoading(view.v))
}

// IsBitmapDirty checks if the bitmap has changed since the last call to GetBitmap()
func (view *View) IsBitmapDirty() bool {
	return bool(C.ulViewIsBitmapDirty(view.v))
}

// GetBitmap returns the bitmap representation of the View
func (view *View) GetBitmap() *Bitmap {
	return &Bitmap{C.ulViewGetBitmap(view.v)}
}

// LoadHTML loads a raw string of HTML into the main frame
func (view *View) LoadHTML(html string) {
	cHTML := C.CString(html)
	defer C.free(unsafe.Pointer(cHTML))
	C.ulViewLoadHTML(view.v, C.ulCreateString(cHTML))
}

// LoadURL loads the specified URL into the main frame
func (view *View) LoadURL(url string) {
	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))
	C.ulViewLoadURL(view.v, C.ulCreateString(cURL))
}

// Resize changes the View dimensions to the specified width
// and height (in device coordinates)
func (view *View) Resize(width, height uint) {
	C.ulViewResize(view.v, C.uint(width), C.uint(height))
}

// GetJSContext gets the JSContext of the current page
func (view *View) GetJSContext() JSContext {
	return JSContext{C.ulViewGetJSContext(view.v)}
}

// EvaluateScript evaluates a raw string of JavaScript, and
// returns the result
func (view *View) EvaluateScript(script string) string {
	cScript := C.CString(script)
	defer C.free(unsafe.Pointer(cScript))
	return C.GoString(C.evaluateScript(view.v, C.ulCreateString(cScript)))
}

// CanGoBack checks if backwards history navigation is available
func (view *View) CanGoBack() bool {
	return bool(C.ulViewCanGoBack(view.v))
}

// CanGoForward checks if forward history navigation is available
func (view *View) CanGoForward() bool {
	return bool(C.ulViewCanGoForward(view.v))
}

// GoBack navigates backwards through the View history
func (view *View) GoBack() {
	C.ulViewGoBack(view.v)
}

// GoForward navigates forwards through the View history
func (view *View) GoForward() {
	C.ulViewGoForward(view.v)
}

// GoToHistoryOffset navigates to the specified offset in the View history
func (view *View) GoToHistoryOffset(offset int) {
	C.ulViewGoToHistoryOffset(view.v, C.int(offset))
}

// Reload refreshes the current page
func (view *View) Reload() {
	C.ulViewReload(view.v)
}

// Stop terminates all page loads
func (view *View) Stop() {
	C.ulViewStop(view.v)
}

// NeedsPaint sets whether the View should be repainted during the next
// call to Renderer.Render()
func (view *View) NeedsPaint(needsPaint bool) {
	C.ulViewSetNeedsPaint(view.v, C.bool(needsPaint))
}

// GetNeedsPaint returns whether the View should be painted during the
// next call to Renderer.Render()
func (view *View) GetNeedsPaint() bool {
	return bool(C.ulViewGetNeedsPaint(view.v))
}

// BindJSCallback executes the specified function when a JavaScript call
// to the 'name' function is made
func (view *View) BindJSCallback(name string, function JSBindFunc) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	funcMap[C.bindScript(view.v, cName)] = viewFunc{view, function}
}

// JSBindString will wrap a string value to be retuned by a bound Go function
func (view *View) JSBindString(val string) *JSValue {
	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cVal))
	return &JSValue{
		jv: C.makeJSValueString(view.v, cVal),
	}
}

// JSBindBool will wrap a boolean value to be retuned by a bound Go function
func (view *View) JSBindBool(val bool) *JSValue {
	return &JSValue{
		jv: C.makeJSValueBool(view.v, C.bool(val)),
	}
}

// JSBindNum will wrap a numeric value to be retuned by a bound Go function
func (view *View) JSBindNum(val float64) *JSValue {
	return &JSValue{
		jv: C.makeJSValueNum(view.v, C.double(val)),
	}
}

// JSBindJSON will wrap a JSON string value to be retuned by a bound Go function
func (view *View) JSBindJSON(val string) *JSValue {
	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cVal))
	return &JSValue{
		jv: C.makeJSValueJSON(view.v, cVal),
	}
}
