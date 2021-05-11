package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "cursor.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Cursor interface {
	objc.Object
	Pop()
	Push()
	Set()
	Image() Image
	HotSpot() foundation.Point
}

type NSCursor struct {
	objc.NSObject
}

func MakeCursor(ptr unsafe.Pointer) *NSCursor {
	if ptr == nil {
		return nil
	}
	return &NSCursor{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocCursor() *NSCursor {
	return MakeCursor(C.C_Cursor_Alloc())
}

func (n *NSCursor) InitWithImage_HotSpot(newImage Image, point foundation.Point) Cursor {
	result := C.C_NSCursor_InitWithImage_HotSpot(n.Ptr(), objc.ExtractPtr(newImage), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return MakeCursor(result)
}

func (n *NSCursor) InitWithCoder(coder foundation.Coder) Cursor {
	result := C.C_NSCursor_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeCursor(result)
}

func (n *NSCursor) Init() Cursor {
	result := C.C_NSCursor_Init(n.Ptr())
	return MakeCursor(result)
}

func CursorHide() {
	C.C_NSCursor_CursorHide()
}

func CursorUnhide() {
	C.C_NSCursor_CursorUnhide()
}

func CursorSetHiddenUntilMouseMoves(flag bool) {
	C.C_NSCursor_CursorSetHiddenUntilMouseMoves(C.bool(flag))
}

func CursorPop() {
	C.C_NSCursor_CursorPop()
}

func (n *NSCursor) Pop() {
	C.C_NSCursor_Pop(n.Ptr())
}

func (n *NSCursor) Push() {
	C.C_NSCursor_Push(n.Ptr())
}

func (n *NSCursor) Set() {
	C.C_NSCursor_Set(n.Ptr())
}

func (n *NSCursor) Image() Image {
	result := C.C_NSCursor_Image(n.Ptr())
	return MakeImage(result)
}

func (n *NSCursor) HotSpot() foundation.Point {
	result := C.C_NSCursor_HotSpot(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}
