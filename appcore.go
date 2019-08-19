package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib/win/x64' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x64
// #cgo windows,i386 LDFLAGS: -L'./SDK/lib/win/x86' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x86
// #cgo linux LDFLAGS: -L'./SDK/bin/linux' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/linux
// #cgo darwin LDFLAGS: -L'./SDK/bin/mac' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/mac
// #include "ultralight.c"
import "C"

/******************************************************************************
 * App
 *****************************************************************************/

// CreateApp creates an ULApp instance
func (conf *Config) CreateApp() *App {
	return &App{C.ulCreateApp(conf.c)}
}

// Destroy .
func (a *App) Destroy() {
	C.ulDestroyApp(a.a)
}

// SetWindow .
func (a *App) SetWindow(win *Window) {
	C.ulAppSetWindow(a.a, win.w)
}

// GetWindow .
func (a *App) GetWindow() *Window {
	return &Window{C.ulAppGetWindow(a.a)}
}

// IsRunning .
func (a *App) IsRunning() bool {
	return bool(C.ulAppIsRunning(a.a))
}

// GetMainMonitor .
func (a *App) GetMainMonitor() *Monitor {
	return &Monitor{C.ulAppGetMainMonitor(a.a)}
}

// GetRenderer .
func (a *App) GetRenderer() *Renderer {
	return &Renderer{C.ulAppGetRenderer(a.a)}
}

// Run .
func (a *App) Run() {
	C.ulAppRun(a.a)
}

// Quit .
func (a *App) Quit() {
	C.ulAppQuit(a.a)
}

/******************************************************************************
 * Monitor
 *****************************************************************************/

// GetScale .
func (mon *Monitor) GetScale() float64 {
	return float64(C.ulMonitorGetScale(mon.m))
}

// GetWidth .
func (mon *Monitor) GetWidth() uint {
	return uint(C.ulMonitorGetWidth(mon.m))
}

// GetHeight .
func (mon *Monitor) GetHeight() uint {
	return uint(C.ulMonitorGetHeight(mon.m))
}

/******************************************************************************
 * Window
 *****************************************************************************/

// CreateWindow .
func CreateWindow(m *Monitor, width, height uint, isTransparent bool, windowType WindowFlag) *Window {
	return &Window{C.ulCreateWindow(m.m, C.uint(width), C.uint(height), C.bool(isTransparent), C.uint(windowType))}
}

// Destroy .
func (win *Window) Destroy() {
	C.ulDestroyWindow(win.w)
}

// GetWidth .
func (win *Window) GetWidth() int {
	return int(C.ulWindowGetWidth(win.w))
}

// GetHeight .
func (win *Window) GetHeight() int {
	return int(C.ulWindowGetHeight(win.w))
}

// IsFullscreen .
func (win *Window) IsFullscreen() bool {
	return bool(C.ulWindowIsFullscreen(win.w))
}

// GetScale .
func (win *Window) GetScale() float64 {
	return float64(C.ulWindowGetScale(win.w))
}

// SetTitle .
func (win *Window) SetTitle(title string) {
	C.ulWindowSetTitle(win.w, C.CString(title))
}

// SetCursor .
func (win *Window) SetCursor(cursor Cursor) {
	C.ulWindowSetCursor(win.w, C.ULCursor(cursor))
}

// Close .
func (win *Window) Close() {
	C.ulWindowClose(win.w)
}

// DeviceToPixel .
func (win *Window) DeviceToPixel(value int) int {
	return int(C.ulWindowDeviceToPixel(win.w, C.int(value)))
}

// PixelsToDevice .
func (win *Window) PixelsToDevice(value int) int {
	return int(C.ulWindowPixelsToDevice(win.w, C.int(value)))
}

/******************************************************************************
 * Overlay
 *****************************************************************************/

// CreateOverlay .
func (win *Window) CreateOverlay(width, height int, x, y int) *Overlay {
	return &Overlay{C.ulCreateOverlay(win.w, C.int(width), C.int(height), C.int(x), C.int(y))}
}

// Destroy .
func (ol *Overlay) Destroy() {
	C.ulDestroyOverlay(ol.o)
}

// GetView .
func (ol *Overlay) GetView() *View {
	return &View{C.ulOverlayGetView(ol.o)}
}

// GetWidth .
func (ol *Overlay) GetWidth() uint {
	return uint(C.ulOverlayGetWidth(ol.o))
}

// GetHeight .
func (ol *Overlay) GetHeight() uint {
	return uint(C.ulOverlayGetHeight(ol.o))
}

// GetX .
func (ol *Overlay) GetX() int {
	return int(C.ulOverlayGetX(ol.o))
}

// GetY .
func (ol *Overlay) GetY() int {
	return int(C.ulOverlayGetY(ol.o))
}

// MoveTo .
func (ol *Overlay) MoveTo(x, y int) {
	C.ulOverlayMoveTo(ol.o, C.int(x), C.int(y))
}

// Resize .
func (ol *Overlay) Resize(width, height int) {
	C.ulOverlayResize(ol.o, C.int(width), C.int(height))
}

// IsHidden .
func (ol *Overlay) IsHidden() bool {
	return bool(C.ulOverlayIsHidden(ol.o))
}

// Hide .
func (ol *Overlay) Hide() {
	C.ulOverlayHide(ol.o)
}

// Show .
func (ol *Overlay) Show() {
	C.ulOverlayShow(ol.o)
}

// HasFocus .
func (ol *Overlay) HasFocus() bool {
	return bool(C.ulOverlayHasFocus(ol.o))
}

// Focus .
func (ol *Overlay) Focus() {
	C.ulOverlayFocus(ol.o)
}

// Unfocus .
func (ol *Overlay) Unfocus() {
	C.ulOverlayUnfocus(ol.o)
}
