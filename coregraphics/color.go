package coregraphics

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreGraphics
// #include "color_space.h"
import "C"
import "unsafe"

type Color struct {
}

type ColorRef Color

func ToCGColorRefPointer(cs ColorRef) unsafe.Pointer {
	panic("not implemented")
}

func FromCGColorRefPointer(p unsafe.Pointer) ColorRef {
	panic("not implemented")
}
