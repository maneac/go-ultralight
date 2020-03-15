package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #cgo linux LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #cgo darwin LDFLAGS: -L'./SDK/bin' -lUltralight -lUltralightCore -lWebCore -lAppCore -Wl,-rpath,.
// #include <ultralight.h>
import "C"

// Renderer wraps the underlying struct
type Renderer struct {
	r C.ULRenderer
}

// CreateRenderer creates a Renderer instance (create only one per application lifetime)
func CreateRenderer(conf *Config) *Renderer {
	return &Renderer{C.ulCreateRenderer(conf.c)}
}

// Destroy deletes the Renderer instance
func (rend *Renderer) Destroy() {
	C.ulDestroyRenderer(rend.r)
}

// Update dispatches internal Javascript and network callbacks and updates timers
func (rend *Renderer) Update() {
	C.ulUpdate(rend.r)
}

// Render renders all active views to their respective bitmaps
func (rend *Renderer) Render() {
	C.ulRender(rend.r)
}
