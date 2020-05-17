package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"
import "unsafe"

// Bitmap wraps the underlying struct
type Bitmap struct {
	b C.ULBitmap
}

// CreateEmptyBitmap will create an empty bitmap
func CreateEmptyBitmap() *Bitmap {
	return &Bitmap{b: C.ulCreateEmptyBitmap()}
}

// CreateBitmap will produce a bitmap with the specified parameters
func CreateBitmap(width, height uint, format BitmapFormat) *Bitmap {
	return &Bitmap{b: C.ulCreateBitmap(C.uint(width), C.uint(height), C.ULBitmapFormat(format))}
}

// CreateCopy creates a copy of the bitmap
func (bit *Bitmap) CreateCopy() *Bitmap {
	return &Bitmap{b: C.ulCreateBitmapFromCopy(bit.b)}
}

// Destroy destroys a bitmap
func (bit *Bitmap) Destroy() {
	C.ulDestroyBitmap(bit.b)
}

// Width returns the width of the bitmap
func (bit *Bitmap) Width() uint {
	return uint(C.ulBitmapGetWidth(bit.b))
}

// Height returns the height of the bitmap
func (bit *Bitmap) Height() uint {
	return uint(C.ulBitmapGetHeight(bit.b))
}

// Format returns the format of the bitmap
func (bit *Bitmap) Format() BitmapFormat {
	return BitmapFormat(C.ulBitmapGetFormat(bit.b))
}

// BytesPerPixel returns the bytes per pixel of the bitmap
func (bit *Bitmap) BytesPerPixel() uint {
	return uint(C.ulBitmapGetBpp(bit.b))
}

// RowBytes returns the bytes per row of the bitmap
func (bit *Bitmap) RowBytes() uint {
	return uint(C.ulBitmapGetRowBytes(bit.b))
}

// IsEmpty returns whether the bitmap is empty or not
func (bit *Bitmap) IsEmpty() bool {
	return bool(C.ulBitmapIsEmpty(bit.b))
}

// Erase resets the bitmap pixels to 0
func (bit *Bitmap) Erase() {
	C.ulBitmapErase(bit.b)
}

// WritePNG writes the Bitmap to the specified path as a PNG
func (bit *Bitmap) WritePNG(path string) bool {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return bool(C.ulBitmapWritePNG(bit.b, cPath))
}
