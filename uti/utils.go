package uti

import (
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/glow/unsafex"
	"unsafe"
)

func toPointer(o objc.Object) unsafe.Pointer {
	if unsafex.InterfaceIsNil(o) {
		return nil
	}
	return o.Ptr()
}
