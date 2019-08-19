package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib/win/x64' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x64
// #cgo windows,i386 LDFLAGS: -L'./SDK/lib/win/x86' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/lib/win/x86
// #cgo linux LDFLAGS: -L'./SDK/bin/linux' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/linux
// #cgo darwin LDFLAGS: -L'./SDK/bin/mac' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,./SDK/bin/mac
// #include "ultralight.c"
import "C"
import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func ulStrToStr(str C.ULString) string {
	cstring := C.strconv(str)
	return C.GoString(cstring)
}

func Init() {
	create()
}

func create() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	_, fp, _, _ := runtime.Caller(0)
	pkg, err := filepath.Abs(filepath.Join(filepath.Dir(fp), "SDK", "bin"))
	if err != nil {
		log.Fatal(err)
	}
	directory, _ := os.Open(pkg)
	objects, _ := directory.Readdir(-1)
	for _, obj := range objects {
		if filepath.Ext(obj.Name()) == ".dll" {
			src, _ := os.Open(filepath.Join(pkg, obj.Name()))
			defer src.Close()
			dst, _ := os.Create(filepath.Join(dir, obj.Name()))
			defer dst.Close()
			_, _ = io.Copy(dst, src)
		}
	}
}

func init() {
	create()
}
