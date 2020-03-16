package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"
import "unsafe"

// Settings wraps the underlying struct
type Settings struct {
	s C.ULSettings
}

// CreateSettings creates settings with default values
func CreateSettings() *Settings {
	return &Settings{C.ulCreateSettings()}
}

// Destroy deletes the settings
func (s *Settings) Destroy() {
	C.ulDestroySettings(s.s)
}

// SetFileSystemPath sets the root file path for all data used by the app
func (s *Settings) SetFileSystemPath(path string) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	C.ulSettingsSetFileSystemPath(s.s, C.ulCreateString(cPath))
}

// SetLoadShadersFromFileSystem decides whether or not to load and compile
// shaders from the file system or load pre-compiled shaders from memory
func (s *Settings) SetLoadShadersFromFileSystem(enable bool) {
	C.ulSettingsSetLoadShadersFromFileSystem(s.s, C.bool(enable))
}
