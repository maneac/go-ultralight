package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #include <ultralight.h>
import "C"

// Overlay wraps the underlying struct
type Overlay struct {
	o C.ULOverlay
}

// CreateOverlay creates a new Overlay instance with the specified width and
// height (in device coordinates) at the offset (x,y) from the top-left
// corner of the Window instance
func CreateOverlay(win *Window, width, height uint, x, y int) *Overlay {
	return &Overlay{C.ulCreateOverlay(win.w, C.uint(width), C.uint(height), C.int(x), C.int(y))}
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
func (ol *Overlay) Resize(width, height uint) {
	C.ulOverlayResize(ol.o, C.uint(width), C.uint(height))
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
