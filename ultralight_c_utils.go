package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"

func ulStrToStr(str C.ULString) string {
	cstring := C.strconv(str)
	return C.GoString(cstring)
}
