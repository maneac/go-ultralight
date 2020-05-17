package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"
import "unsafe"

// Config wraps the C config stuct
type Config struct {
	c C.ULConfig
}

// CreateConfig creates a Config instance
func CreateConfig() *Config {
	return &Config{C.ulCreateConfig()}
}

// Destroy deletes the Config instance
func (conf *Config) Destroy() {
	C.ulDestroyConfig(conf.c)
}

// EnableImages sets whether images should be enabled (Default = true)
func (conf *Config) EnableImages(enabled bool) {
	C.ulConfigSetEnableImages(conf.c, C.bool(enabled))
}

// EnableJavaScript sets whether JavaScript should be enabled (Default = true)
func (conf *Config) EnableJavaScript(enabled bool) {
	C.ulConfigSetEnableJavaScript(conf.c, C.bool(enabled))
}

// EnableBGRA sets wheter to use BGRA byte order (instead of RGBA) for offscreen
// rendering of View bitmaps (Default = false)
func (conf *Config) EnableBGRA(enabled bool) {
	C.ulConfigSetUseBGRAForOffscreenRendering(conf.c, C.bool(enabled))
}

// DeviceScaleHint sets the amount that the application DPI has been scaled,
// which is used for scaling device coordinates to pixels and oversampling
// raster shapes (Default = 1.0)
func (conf *Config) DeviceScaleHint(value float64) {
	C.ulConfigSetDeviceScaleHint(conf.c, C.double(value))
}

// FontFamilyStandard sets the default font-family to use (Default = "Times New Roman")
func (conf *Config) FontFamilyStandard(fontName string) {
	cFontName := C.CString(fontName)
	defer C.free(unsafe.Pointer(cFontName))
	C.ulConfigSetFontFamilyStandard(conf.c, C.ulCreateString(cFontName))
}

// FontFamilyFixed sets the default font-family to use for fixed fonts, e.g. <pre>
// and <code> (Default = "Courier New")
func (conf *Config) FontFamilyFixed(fontName string) {
	cFontName := C.CString(fontName)
	defer C.free(unsafe.Pointer(cFontName))
	C.ulConfigSetFontFamilyFixed(conf.c, C.ulCreateString(cFontName))
}

// FontFamilySerif sets the default font-family to use for serif fonts
// (Default = "Times New Roman")
func (conf *Config) FontFamilySerif(fontName string) {
	cFontName := C.CString(fontName)
	defer C.free(unsafe.Pointer(cFontName))
	C.ulConfigSetFontFamilySerif(conf.c, C.ulCreateString(cFontName))
}

// FontFamilySansSerif sets the default font-family to use for sans-serif fonts
// (Default = "Arial")
func (conf *Config) FontFamilySansSerif(fontName string) {
	cFontName := C.CString(fontName)
	defer C.free(unsafe.Pointer(cFontName))
	C.ulConfigSetFontFamilySansSerif(conf.c, C.ulCreateString(cFontName))
}

// UserAgent sets the user agent string
func (conf *Config) UserAgent(agent string) {
	cAgent := C.CString(agent)
	defer C.free(unsafe.Pointer(cAgent))
	C.ulConfigSetUserAgent(conf.c, C.ulCreateString(cAgent))
}

// UserStylesheet sets the user stylesheet (CSS) (Default = Empty)
func (conf *Config) UserStylesheet(css string) {
	cCSS := C.CString(css)
	defer C.free(unsafe.Pointer(cCSS))
	C.ulConfigSetUserStylesheet(conf.c, C.ulCreateString(cCSS))
}

// ForceRepaint sets whether the views should be continuously repainted or not. Mainly
// used for diagnosis
func (conf *Config) ForceRepaint(enabled bool) {
	C.ulConfigSetForceRepaint(conf.c, C.bool(enabled))
}

// AnimationTimerDelay sets the amount of time to wait before repainting when a
// CSS animation is active (Default = 1/60)
func (conf *Config) AnimationTimerDelay(delay float64) {
	C.ulConfigSetAnimationTimerDelay(conf.c, C.double(delay))
}

// MemoryCacheSize sets the size of the cache for assets in bytes (Default = 64MB)
func (conf *Config) MemoryCacheSize(size uint) {
	C.ulConfigSetMemoryCacheSize(conf.c, C.uint(size))
}

// PageCacheSize sets the number of pages to keep cached (Default = 0)
func (conf *Config) PageCacheSize(size uint) {
	C.ulConfigSetPageCacheSize(conf.c, C.uint(size))
}
