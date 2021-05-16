package coregraphics

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreGraphics
// #include "color_space.h"
import "C"
import "unsafe"

type Event struct {
}

type EventRef Event

func ToCGEventRefPointer(cs EventRef) unsafe.Pointer {
	panic("not implemented")
}

func FromCGEventRefPointer(p unsafe.Pointer) EventRef {
	panic("not implemented")
}
