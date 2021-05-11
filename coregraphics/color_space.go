package coregraphics

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreGraphics
// #include "color_space.h"
import "C"
import "unsafe"

type ColorSpace struct {
}

// TODO: should be *ColorSpace
type ColorSpaceRef ColorSpace

func ToCGColorSpaceRefPointer(cs ColorSpaceRef) unsafe.Pointer {
	panic("not implemented")
}

func FromCGColorSpaceRefPointer(p unsafe.Pointer) ColorSpaceRef {
	panic("not implemented")
}
