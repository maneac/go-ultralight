package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib/win/x64' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x64
// #cgo windows,i386 LDFLAGS: -L'./SDK/lib/win/x86' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x86
// #cgo linux LDFLAGS: -L'./SDK/bin/linux' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/linux
// #cgo darwin LDFLAGS: -L'./SDK/bin/mac' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/mac
// #include "ultralight.c"
import "C"

/******************************************************************************
 * Config
 *****************************************************************************/

// CreateConfig .
func CreateConfig() *Config {
	return &Config{C.ulCreateConfig()}
}

// Destroy .
func (conf *Config) Destroy() {
	C.ulDestroyConfig(conf.c)
}

// EnableImages .
func (conf *Config) EnableImages(enabled bool) {
	C.ulConfigSetEnableImages(conf.c, C.bool(enabled))
}

// EnableJavaScript .
func (conf *Config) EnableJavaScript(enabled bool) {
	C.ulConfigSetEnableJavaScript(conf.c, C.bool(enabled))
}

// EnableBGRA .
func (conf *Config) EnableBGRA(enabled bool) {
	C.ulConfigSetUseBGRAForOffscreenRendering(conf.c, C.bool(enabled))
}

// DeviceScaleHint .
func (conf *Config) DeviceScaleHint(value float64) {
	C.ulConfigSetDeviceScaleHint(conf.c, C.double(value))
}

// FontFamilyStandard .
func (conf *Config) FontFamilyStandard(fontName string) {
	C.ulConfigSetFontFamilyStandard(conf.c, C.ulCreateString(C.CString(fontName)))
}

// FontFamilyFixed .
func (conf *Config) FontFamilyFixed(fontName string) {
	C.ulConfigSetFontFamilyFixed(conf.c, C.ulCreateString(C.CString(fontName)))
}

// FontFamilySerif .
func (conf *Config) FontFamilySerif(fontName string) {
	C.ulConfigSetFontFamilySerif(conf.c, C.ulCreateString(C.CString(fontName)))
}

// FontFamilySansSerif .
func (conf *Config) FontFamilySansSerif(fontName string) {
	C.ulConfigSetFontFamilySansSerif(conf.c, C.ulCreateString(C.CString(fontName)))
}

// UserAgent .
func (conf *Config) UserAgent(agent string) {
	C.ulConfigSetUserAgent(conf.c, C.ulCreateString(C.CString(agent)))
}

// UserStylesheet .
func (conf *Config) UserStylesheet(css string) {
	C.ulConfigSetUserStylesheet(conf.c, C.ulCreateString(C.CString(css)))
}

/******************************************************************************
 * Renderer
 *****************************************************************************/

// CreateRenderer .
func (conf *Config) CreateRenderer() *Renderer {
	return &Renderer{C.ulCreateRenderer(conf.c)}
}

// Destroy .
func (rend *Renderer) Destroy() {
	C.ulDestroyRenderer(rend.r)
}

// Update .
func (rend *Renderer) Update() {
	C.ulUpdate(rend.r)
}

// Render .
func (rend *Renderer) Render() {
	C.ulRender(rend.r)
}

/******************************************************************************
* View
*****************************************************************************/

// CreateView .
func (rend *Renderer) CreateView(width, height uint, isTransparent bool) *View {
	return &View{C.ulCreateView(rend.r, C.uint(width), C.uint(height), C.bool(isTransparent))}
}

// Destroy .
func (view *View) Destroy() {
	C.ulDestroyView(view.v)
}

// GetURL .
func (view *View) GetURL() string {
	return ulStrToStr(C.ulViewGetURL(view.v))
}

// GetTitle .
func (view *View) GetTitle() string {
	return ulStrToStr(C.ulViewGetTitle(view.v))
}

// IsLoading .
func (view *View) IsLoading() bool {
	return bool(C.ulViewIsLoading(view.v))
}

// IsBitmapDirty .
func (view *View) IsBitmapDirty() bool {
	return bool(C.ulViewIsBitmapDirty(view.v))
}

// GetBitmap .
func (view *View) GetBitmap() Bitmap {
	return Bitmap{C.ulViewGetBitmap(view.v)}
}

// LoadHTML .
func (view *View) LoadHTML(html string) {
	C.ulViewLoadHTML(view.v, C.ulCreateString(C.CString(html)))
}

// LoadURL .
func (view *View) LoadURL(url string) {
	C.ulViewLoadURL(view.v, C.ulCreateString(C.CString(url)))
}

// Resize .
func (view *View) Resize(width, height uint) {
	C.ulViewResize(view.v, C.uint(width), C.uint(height))
}

// GetJSContext .
func (view *View) GetJSContext() JSContext {
	return JSContext{C.ulViewGetJSContext(view.v)}
}

// EvaluateScript .
func (view *View) EvaluateScript(script string) string {
	return C.GoString(C.evaluateScript(view.v, C.ulCreateString(C.CString(script))))
}

// CanGoBack .
func (view *View) CanGoBack() bool {
	return bool(C.ulViewCanGoBack(view.v))
}

// CanGoForward .
func (view *View) CanGoForward() bool {
	return bool(C.ulViewCanGoForward(view.v))
}

// GoBack .
func (view *View) GoBack() {
	C.ulViewGoBack(view.v)
}

// GoForward .
func (view *View) GoForward() {
	C.ulViewGoForward(view.v)
}

// GoToHistoryOffset .
func (view *View) GoToHistoryOffset(offset int) {
	C.ulViewGoToHistoryOffset(view.v, C.int(offset))
}

// Reload .
func (view *View) Reload() {
	C.ulViewReload(view.v)
}

// Stop .
func (view *View) Stop() {
	C.ulViewStop(view.v)
}

// SetNeedsPaint .
func (view *View) SetNeedsPaint(needsPaint bool) {
	C.ulViewSetNeedsPaint(view.v, C.bool(needsPaint))
}

// GetNeedsPaint .
func (view *View) GetNeedsPaint() bool {
	return bool(C.ulViewGetNeedsPaint(view.v))
}

// BindJSCallback .
func (view *View) BindJSCallback(name string, function JSBindFunc) {
	FuncMap[C.bindScript(view.v, C.CString(name))] = ViewFunc{view, function}
}

// BITMAP

// WritePNG .
func (bit *Bitmap) WritePNG(path string) bool {
	return bool(C.ulBitmapWritePNG(bit.b, C.CString(path)))
}
