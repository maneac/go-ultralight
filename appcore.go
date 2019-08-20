package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib/win/x64' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo windows,i386 LDFLAGS: -L'./SDK/lib/win/x86' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo linux LDFLAGS: -L'./SDK/bin/linux' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin/mac' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"

/******************************************************************************
 * App
 *****************************************************************************/

// CreateApp creates an App instance (create only one per application lifetime)
func (conf *Config) CreateApp() *App {
	return &App{C.ulCreateApp(conf.c)}
}

// Destroy deletes the App instance
func (a *App) Destroy() {
	C.ulDestroyApp(a.a)
}

// SetWindow sets the main window. This must be called before
// App.Run()
func (a *App) SetWindow(win *Window) {
	C.ulAppSetWindow(a.a, win.w)
}

// GetWindow returns the main Window
func (a *App) GetWindow() *Window {
	return &Window{C.ulAppGetWindow(a.a)}
}

// IsRunning returns whether the App is running
func (a *App) IsRunning() bool {
	return bool(C.ulAppIsRunning(a.a))
}

// GetMainMonitor returns the main Monitor instance
func (a *App) GetMainMonitor() *Monitor {
	return &Monitor{C.ulAppGetMainMonitor(a.a)}
}

// GetRenderer returns the underlying Renderer instance
func (a *App) GetRenderer() *Renderer {
	return &Renderer{C.ulAppGetRenderer(a.a)}
}

// Run executes the main loop. Ensure App.SetWindow() has been called
// prior to calling
func (a *App) Run() {
	C.ulAppRun(a.a)
}

// Quit exits the application
func (a *App) Quit() {
	C.ulAppQuit(a.a)
}

/******************************************************************************
 * Monitor
 *****************************************************************************/

// GetScale returns the DPI scale of the Monitor as a percentage
func (mon *Monitor) GetScale() float64 {
	return float64(C.ulMonitorGetScale(mon.m))
}

// GetWidth returns the width of the monitor (in device coordinates)
func (mon *Monitor) GetWidth() uint {
	return uint(C.ulMonitorGetWidth(mon.m))
}

// GetHeight returns the height of the monitor (in device coordinates)
func (mon *Monitor) GetHeight() uint {
	return uint(C.ulMonitorGetHeight(mon.m))
}

/******************************************************************************
 * Window
 *****************************************************************************/

// CreateWindow creates a new Window with the specified size (in device coordinates)
func CreateWindow(m *Monitor, width, height uint, isTransparent bool, windowType WindowFlag) *Window {
	return &Window{C.ulCreateWindow(m.m, C.uint(width), C.uint(height), C.bool(isTransparent), C.uint(windowType))}
}

// Destroy deletes the Window instance
func (win *Window) Destroy() {
	C.ulDestroyWindow(win.w)
}

// GetWidth returns the width of the Window instance (in device coordinates)
func (win *Window) GetWidth() int {
	return int(C.ulWindowGetWidth(win.w))
}

// GetHeight returns the height of the Window instance (in device coordinates)
func (win *Window) GetHeight() int {
	return int(C.ulWindowGetHeight(win.w))
}

// IsFullscreen returns whether the Window is fullscreen
func (win *Window) IsFullscreen() bool {
	return bool(C.ulWindowIsFullscreen(win.w))
}

// GetScale returns the DPI scale of the Window as a percentage
func (win *Window) GetScale() float64 {
	return float64(C.ulWindowGetScale(win.w))
}

// SetTitle sets the title of the Window instance
func (win *Window) SetTitle(title string) {
	C.ulWindowSetTitle(win.w, C.CString(title))
}

// SetCursor sets the cursor for the Window instance
func (win *Window) SetCursor(cursor Cursor) {
	C.ulWindowSetCursor(win.w, C.ULCursor(cursor))
}

// Close closes the Window
func (win *Window) Close() {
	C.ulWindowClose(win.w)
}

// DeviceToPixel converts device coordinates to pixels using the current
// DPI scale
func (win *Window) DeviceToPixel(value int) int {
	return int(C.ulWindowDeviceToPixel(win.w, C.int(value)))
}

// PixelsToDevice converts pixels to device coordinates using the current
// DPI scale
func (win *Window) PixelsToDevice(value int) int {
	return int(C.ulWindowPixelsToDevice(win.w, C.int(value)))
}

/******************************************************************************
 * Overlay
 *****************************************************************************/

// CreateOverlay creates a new Overlay instance with the specified width and
// height (in device coordinates) at the offset (x,y) from the top-left
// corner of the Window instance
func (win *Window) CreateOverlay(width, height int, x, y int) *Overlay {
	return &Overlay{C.ulCreateOverlay(win.w, C.int(width), C.int(height), C.int(x), C.int(y))}
}

// Destroy deletes the Overlay instance
func (ol *Overlay) Destroy() {
	C.ulDestroyOverlay(ol.o)
}

// GetView returns the underlying View instance
func (ol *Overlay) GetView() *View {
	return &View{C.ulOverlayGetView(ol.o)}
}

// GetWidth returns the width of the Overlay (in device coordinates)
func (ol *Overlay) GetWidth() uint {
	return uint(C.ulOverlayGetWidth(ol.o))
}

// GetHeight returns the height of the Overlay (in device coordinates)
func (ol *Overlay) GetHeight() uint {
	return uint(C.ulOverlayGetHeight(ol.o))
}

// GetX returns the horizontal offset of the Overlay relative
// to it's Window (in device coordinates)
func (ol *Overlay) GetX() int {
	return int(C.ulOverlayGetX(ol.o))
}

// GetY returns the vertical offset of the Overlay relative
// to it's Window (in device coordinates)
func (ol *Overlay) GetY() int {
	return int(C.ulOverlayGetY(ol.o))
}

// MoveTo moves the Overlay to the specified (x,y) position
// (in device coordinates)
func (ol *Overlay) MoveTo(x, y int) {
	C.ulOverlayMoveTo(ol.o, C.int(x), C.int(y))
}

// Resize sets the size of the Overlay and it's
// underlying View to the specified dimensions (in device coordinates)
func (ol *Overlay) Resize(width, height int) {
	C.ulOverlayResize(ol.o, C.int(width), C.int(height))
}

// IsHidden returns whether the Overlay is hidden or drawn
func (ol *Overlay) IsHidden() bool {
	return bool(C.ulOverlayIsHidden(ol.o))
}

// Hide stops the Overlay from being drawn
func (ol *Overlay) Hide() {
	C.ulOverlayHide(ol.o)
}

// Show draws the Overlay instance
func (ol *Overlay) Show() {
	C.ulOverlayShow(ol.o)
}

// HasFocus returns whether the Overlay has keyboard focus
func (ol *Overlay) HasFocus() bool {
	return bool(C.ulOverlayHasFocus(ol.o))
}

// Focus grants the Overlay exclusive keyboard focus
func (ol *Overlay) Focus() {
	C.ulOverlayFocus(ol.o)
}

// Unfocus revokes exclusive keyboard focus from the Overlay
func (ol *Overlay) Unfocus() {
	C.ulOverlayUnfocus(ol.o)
}
