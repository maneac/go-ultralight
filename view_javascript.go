package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #include <ultralight.h>
import "C"

import "strings"

// JSContext wraps the underlying struct
type JSContext struct {
	jc C.JSContextRef
}

// JSValue wraps the underlying struct
type JSValue struct {
	jv C.JSValueRef
}

var funcMap map[C.JSObjectRef]viewFunc

type viewFunc struct {
	v *View
	f JSBindFunc
}

func init() {
	funcMap = make(map[C.JSObjectRef]viewFunc)
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
