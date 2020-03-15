package ultralight

// #cgo CPPFLAGS: -I"./SDK/include"
// #cgo windows,amd64 LDFLAGS: -L'./SDK/lib' -lUltralight -lUltralightCore -lWebCore -lAppCore
// #include <ultralight.h>
import "C"

// Monitor wraps the underlying struct
type Monitor struct {
	m C.ULMonitor
}

// GetScale returns the DPI scale of the Monitor as a percentage
func (mon *Monitor) GetScale() float64 {
	return float64(C.ulMonitorGetScale(mon.m))
}

// GetWidth returns the width of the monitor (in device coordinates)
func (mon *Monitor) GetWidth() uint {
	return uint(C.ulMonitorGetWidth(mon.m))
}

// GetHeight returns the height of the monitor (in device coordinates)
func (mon *Monitor) GetHeight() uint {
	return uint(C.ulMonitorGetHeight(mon.m))
}
