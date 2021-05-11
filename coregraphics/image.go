package coregraphics

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreGraphics
// #include "image.h"
import "C"
import "unsafe"

type Image struct {
}

// TODO: should be *Image
type ImageRef Image

func ToCGImageRefPointer(ir ImageRef) unsafe.Pointer {
	panic("not implemented")
}

func FromCGImageRefPointer(p unsafe.Pointer) ImageRef {
	panic("not implemented")
}
