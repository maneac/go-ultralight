package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"
import "unsafe"

// KeyEvent wraps the underlying struct
type KeyEvent struct {
	ke C.ULKeyEvent
}

// MouseEvent wraps the underlying struct
type MouseEvent struct {
	me C.ULMouseEvent
}

// ScrollEvent wraps the underlying struct
type ScrollEvent struct {
	se C.ULScrollEvent
}

// CreateKeyEvent creates a new key event
func CreateKeyEvent(eventType KeyEventType,
	modifiers uint, virtKeyCode, nativeKeyCode int,
	text, rawText string, isKeypad, isAutoRepeat, isSystemKey bool) *KeyEvent {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	cRawText := C.CString(rawText)
	defer C.free(unsafe.Pointer(cRawText))
	return &KeyEvent{
		ke: C.ulCreateKeyEvent(C.ULKeyEventType(eventType), C.uint(modifiers),
			C.int(virtKeyCode), C.int(nativeKeyCode),
			C.ulCreateString(cText), C.ulCreateString(cRawText),
			C.bool(isKeypad), C.bool(isAutoRepeat), C.bool(isSystemKey)),
	}
}

// Destroy destroys the key event
func (ke *KeyEvent) Destroy() {
	C.ulDestroyKeyEvent(ke.ke)
}

// CreateMouseEvent creates a new mouse event
func CreateMouseEvent(eventType MouseEventType, x, y int, button MouseButton) *MouseEvent {
	return &MouseEvent{me: C.ulCreateMouseEvent(C.ULMouseEventType(eventType), C.int(x), C.int(y), C.ULMouseButton(button))}
}

// Destroy destroys the mouse event
func (me *MouseEvent) Destroy() {
	C.ulDestroyMouseEvent(me.me)
}

// CreateScrollEvent creates a new scroll event
func CreateScrollEvent(eventType ScrollEventType, deltaX, deltaY int) *ScrollEvent {
	return &ScrollEvent{se: C.ulCreateScrollEvent(C.ULScrollEventType(eventType), C.int(deltaX), C.int(deltaY))}
}

// Destroy destroys the scroll event
func (se *ScrollEvent) Destroy() {
	C.ulDestroyScrollEvent(se.se)
}
