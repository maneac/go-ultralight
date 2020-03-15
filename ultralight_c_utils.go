package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #include <ultralight.h>
import "C"

func ulStrToStr(str C.ULString) string {
	cstring := C.strconv(str)
	return C.GoString(cstring)
}
