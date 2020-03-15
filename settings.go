package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #include <ultralight.h>
import "C"

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
	C.ulSettingsSetFileSystemPath(s.s, C.ulCreateString(C.CString(path)))
}

// SetLoadShadersFromFileSystem decides whether or not to load and compile
// shaders from the file system or load pre-compiled shaders from memory
func (s *Settings) SetLoadShadersFromFileSystem(enable bool) {
	C.ulSettingsSetLoadShadersFromFileSystem(s.s, C.bool(enable))
}
