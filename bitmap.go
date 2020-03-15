package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"

// Bitmap wraps the underlying struct
type Bitmap struct {
	b C.ULBitmap
}

// WritePNG writes the Bitmap to the specified path as a PNG
func (bit *Bitmap) WritePNG(path string) bool {
	return bool(C.ulBitmapWritePNG(bit.b, C.CString(path)))
}
