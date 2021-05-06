package uti

import (
	"unsafe"

	"github.com/hsiafan/cocoa/internal/utils"
	"github.com/hsiafan/cocoa/objc"
)

func toPointer(o objc.Object) unsafe.Pointer {
	if utils.InterfaceIsNil(o) {
		return nil
	}
	return o.Ptr()
}
