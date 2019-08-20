package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib/win/x64' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo windows,i386 LDFLAGS: -L'./SDK/lib/win/x86' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo linux LDFLAGS: -L'./SDK/bin/linux' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin/mac' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"

/******************************************************************************
 * Config
 *****************************************************************************/

// CreateConfig creates a Config instance
func CreateConfig() *Config {
	return &Config{C.ulCreateConfig()}
}

// Destroy deletes the Config instance
func (conf *Config) Destroy() {
	C.ulDestroyConfig(conf.c)
}

// EnableImages sets whether images should be enabled (Default = true)
func (conf *Config) EnableImages(enabled bool) {
	C.ulConfigSetEnableImages(conf.c, C.bool(enabled))
}

// EnableJavaScript sets whether JavaScript should be enabled (Default = true)
func (conf *Config) EnableJavaScript(enabled bool) {
	C.ulConfigSetEnableJavaScript(conf.c, C.bool(enabled))
}

// EnableBGRA sets wheter to use BGRA byte order (instead of RGBA) for offscreen
// rendering of View bitmaps (Default = false)
func (conf *Config) EnableBGRA(enabled bool) {
	C.ulConfigSetUseBGRAForOffscreenRendering(conf.c, C.bool(enabled))
}

// DeviceScaleHint sets the amount that the application DPI has been scaled,
// which is used for scaling device coordinates to pixels and oversampling
// raster shapes (Default = 1.0)
func (conf *Config) DeviceScaleHint(value float64) {
	C.ulConfigSetDeviceScaleHint(conf.c, C.double(value))
}

// FontFamilyStandard sets the default font-family to use (Default = "Times New Roman")
func (conf *Config) FontFamilyStandard(fontName string) {
	C.ulConfigSetFontFamilyStandard(conf.c, C.ulCreateString(C.CString(fontName)))
}

// FontFamilyFixed sets the default font-family to use for fixed fonts, e.g. <pre>
// and <code> (Default = "Courier New")
func (conf *Config) FontFamilyFixed(fontName string) {
	C.ulConfigSetFontFamilyFixed(conf.c, C.ulCreateString(C.CString(fontName)))
}

// FontFamilySerif sets the default font-family to use for serif fonts
// (Default = "Times New Roman")
func (conf *Config) FontFamilySerif(fontName string) {
	C.ulConfigSetFontFamilySerif(conf.c, C.ulCreateString(C.CString(fontName)))
}

// FontFamilySansSerif sets the default font-family to use for sans-serif fonts
// (Default = "Arial")
func (conf *Config) FontFamilySansSerif(fontName string) {
	C.ulConfigSetFontFamilySansSerif(conf.c, C.ulCreateString(C.CString(fontName)))
}

// UserAgent sets the user agent string
func (conf *Config) UserAgent(agent string) {
	C.ulConfigSetUserAgent(conf.c, C.ulCreateString(C.CString(agent)))
}

// UserStylesheet sets the user stylesheet (CSS) (Default = Empty)
func (conf *Config) UserStylesheet(css string) {
	C.ulConfigSetUserStylesheet(conf.c, C.ulCreateString(C.CString(css)))
}

/******************************************************************************
 * Renderer
 *****************************************************************************/

// CreateRenderer creates a Renderer instance (create only one per application lifetime)
func (conf *Config) CreateRenderer() *Renderer {
	return &Renderer{C.ulCreateRenderer(conf.c)}
}

// Destroy deletes the Renderer instance
func (rend *Renderer) Destroy() {
	C.ulDestroyRenderer(rend.r)
}

// Update dispatches internal Javascript and network callbacks and updates timers
func (rend *Renderer) Update() {
	C.ulUpdate(rend.r)
}

// Render renders all active views to their respective bitmaps
func (rend *Renderer) Render() {
	C.ulRender(rend.r)
}

/******************************************************************************
* View
*****************************************************************************/

// CreateView creates a View instance with the specified size (in device coordinates)
func (rend *Renderer) CreateView(width, height uint, isTransparent bool) *View {
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
	C.ulViewLoadHTML(view.v, C.ulCreateString(C.CString(html)))
}

// LoadURL loads the specified URL into the main frame
func (view *View) LoadURL(url string) {
	C.ulViewLoadURL(view.v, C.ulCreateString(C.CString(url)))
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
	return C.GoString(C.evaluateScript(view.v, C.ulCreateString(C.CString(script))))
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
	funcMap[C.bindScript(view.v, C.CString(name))] = viewFunc{view, function}
}

/******************************************************************************
* Bitmap
*****************************************************************************/

// WritePNG writes the Bitmap to the specified path as a PNG
func (bit *Bitmap) WritePNG(path string) bool {
	return bool(C.ulBitmapWritePNG(bit.b, C.CString(path)))
}
