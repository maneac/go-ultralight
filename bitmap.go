package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
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
