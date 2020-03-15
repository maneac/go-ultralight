package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #include <ultralight.h>
import "C"

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
	C.ulConfigSetFontFamilyStandard(conf.c, C.ulCreateString(C.CString(fontName)))
}

// FontFamilyFixed sets the default font-family to use for fixed fonts, e.g. <pre>
// and <code> (Default = "Courier New")
func (conf *Config) FontFamilyFixed(fontName string) {
	C.ulConfigSetFontFamilyFixed(conf.c, C.ulCreateString(C.CString(fontName)))
}

// FontFamilySerif sets the default font-family to use for serif fonts
// (Default = "Times New Roman")
func (conf *Config) FontFamilySerif(fontName string) {
	C.ulConfigSetFontFamilySerif(conf.c, C.ulCreateString(C.CString(fontName)))
}

// FontFamilySansSerif sets the default font-family to use for sans-serif fonts
// (Default = "Arial")
func (conf *Config) FontFamilySansSerif(fontName string) {
	C.ulConfigSetFontFamilySansSerif(conf.c, C.ulCreateString(C.CString(fontName)))
}

// UserAgent sets the user agent string
func (conf *Config) UserAgent(agent string) {
	C.ulConfigSetUserAgent(conf.c, C.ulCreateString(C.CString(agent)))
}

// UserStylesheet sets the user stylesheet (CSS) (Default = Empty)
func (conf *Config) UserStylesheet(css string) {
	C.ulConfigSetUserStylesheet(conf.c, C.ulCreateString(C.CString(css)))
}
