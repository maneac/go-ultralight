package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib/win/x64' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x64
// #cgo windows,i386 LDFLAGS: -L'./SDK/lib/win/x86' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x86
// #cgo linux LDFLAGS: -L'./SDK/bin/linux' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/linux
// #cgo darwin LDFLAGS: -L'./SDK/bin/mac' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/mac
// #include "ultralight.c"
import "C"
import (
	"strings"
)

/******************************************************************************
 * Strings
 *****************************************************************************/

func ulStrToStr(str C.ULString) string {
	cstring := C.strconv(str)
	return C.GoString(cstring)
}

/******************************************************************************
 * JavaScript
 *****************************************************************************/

var funcMap map[C.JSObjectRef]viewFunc

type viewFunc struct {
	v *View
	f JSBindFunc
}

//export objFunctionCallback
func objFunctionCallback(ctx C.JSContextRef, function C.JSObjectRef, _ C.JSObjectRef,
	argCt C.size_t, parameters *C.JSValueRef, _ *C.JSValueRef) C.JSValueRef {
	args := strings.Split(C.GoString(C.printParams(ctx, parameters, argCt)), "\t\t\t")
	if len(args) == 1 && args[0] == "" {
		args = []string{}
	}
	vf := funcMap[function]
	vf.f(vf.v, args)
	return nil
}
