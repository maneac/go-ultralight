package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"

// Window wraps the underlying struct
type Window struct {
	w C.ULWindow
}

// CreateWindow creates a new Window with the specified size (in device coordinates)
func CreateWindow(m *Monitor, width, height uint, isTransparent bool, windowType WindowFlag) *Window {
	return &Window{C.ulCreateWindow(m.m, C.uint(width), C.uint(height), C.bool(isTransparent), C.uint(windowType))}
}

// Destroy deletes the Window instance
func (win *Window) Destroy() {
	C.ulDestroyWindow(win.w)
}

// GetWidth returns the width of the Window instance (in device coordinates)
func (win *Window) GetWidth() uint {
	return uint(C.ulWindowGetWidth(win.w))
}

// GetHeight returns the height of the Window instance (in device coordinates)
func (win *Window) GetHeight() uint {
	return uint(C.ulWindowGetHeight(win.w))
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
