package appkit

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
	Push()
	Set()
	Image() Image
	HotSpot() foundation.Point
}

type NSCursor struct {
	objc.NSObject
}

func MakeCursor(ptr unsafe.Pointer) NSCursor {
	return NSCursor{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSCursor) InitWithImage_HotSpot(newImage Image, point foundation.Point) NSCursor {
	result_ := C.C_NSCursor_InitWithImage_HotSpot(n.Ptr(), objc.ExtractPtr(newImage), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return MakeCursor(result_)
}

func (n NSCursor) InitWithCoder(coder foundation.Coder) NSCursor {
	result_ := C.C_NSCursor_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeCursor(result_)
}

func AllocCursor() NSCursor {
	result_ := C.C_NSCursor_AllocCursor()
	return MakeCursor(result_)
}

func (n NSCursor) Init() NSCursor {
	result_ := C.C_NSCursor_Init(n.Ptr())
	return MakeCursor(result_)
}

func NewCursor() NSCursor {
	result_ := C.C_NSCursor_NewCursor()
	return MakeCursor(result_)
}

func (n NSCursor) Autorelease() NSCursor {
	result_ := C.C_NSCursor_Autorelease(n.Ptr())
	return MakeCursor(result_)
}

func (n NSCursor) Retain() NSCursor {
	result_ := C.C_NSCursor_Retain(n.Ptr())
	return MakeCursor(result_)
}

func Cursor_Hide() {
	C.C_NSCursor_Cursor_Hide()
}

func Cursor_Unhide() {
	C.C_NSCursor_Cursor_Unhide()
}

func Cursor_SetHiddenUntilMouseMoves(flag bool) {
	C.C_NSCursor_Cursor_SetHiddenUntilMouseMoves(C.bool(flag))
}

func Cursor_Pop() {
	C.C_NSCursor_Cursor_Pop()
}

func (n NSCursor) Push() {
	C.C_NSCursor_Push(n.Ptr())
}

func (n NSCursor) Set() {
	C.C_NSCursor_Set(n.Ptr())
}

func (n NSCursor) Image() Image {
	result_ := C.C_NSCursor_Image(n.Ptr())
	return MakeImage(result_)
}

func (n NSCursor) HotSpot() foundation.Point {
	result_ := C.C_NSCursor_HotSpot(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func CurrentCursor() Cursor {
	result_ := C.C_NSCursor_CurrentCursor()
	return MakeCursor(result_)
}

func CurrentSystemCursor() Cursor {
	result_ := C.C_NSCursor_CurrentSystemCursor()
	return MakeCursor(result_)
}

func ArrowCursor() Cursor {
	result_ := C.C_NSCursor_ArrowCursor()
	return MakeCursor(result_)
}

func ContextualMenuCursor() Cursor {
	result_ := C.C_NSCursor_ContextualMenuCursor()
	return MakeCursor(result_)
}

func ClosedHandCursor() Cursor {
	result_ := C.C_NSCursor_ClosedHandCursor()
	return MakeCursor(result_)
}

func CrosshairCursor() Cursor {
	result_ := C.C_NSCursor_CrosshairCursor()
	return MakeCursor(result_)
}

func DisappearingItemCursor() Cursor {
	result_ := C.C_NSCursor_DisappearingItemCursor()
	return MakeCursor(result_)
}

func DragCopyCursor() Cursor {
	result_ := C.C_NSCursor_DragCopyCursor()
	return MakeCursor(result_)
}

func DragLinkCursor() Cursor {
	result_ := C.C_NSCursor_DragLinkCursor()
	return MakeCursor(result_)
}

func IBeamCursor() Cursor {
	result_ := C.C_NSCursor_IBeamCursor()
	return MakeCursor(result_)
}

func OpenHandCursor() Cursor {
	result_ := C.C_NSCursor_OpenHandCursor()
	return MakeCursor(result_)
}

func OperationNotAllowedCursor() Cursor {
	result_ := C.C_NSCursor_OperationNotAllowedCursor()
	return MakeCursor(result_)
}

func PointingHandCursor() Cursor {
	result_ := C.C_NSCursor_PointingHandCursor()
	return MakeCursor(result_)
}

func ResizeDownCursor() Cursor {
	result_ := C.C_NSCursor_ResizeDownCursor()
	return MakeCursor(result_)
}

func ResizeLeftCursor() Cursor {
	result_ := C.C_NSCursor_ResizeLeftCursor()
	return MakeCursor(result_)
}

func ResizeLeftRightCursor() Cursor {
	result_ := C.C_NSCursor_ResizeLeftRightCursor()
	return MakeCursor(result_)
}

func ResizeRightCursor() Cursor {
	result_ := C.C_NSCursor_ResizeRightCursor()
	return MakeCursor(result_)
}

func ResizeUpCursor() Cursor {
	result_ := C.C_NSCursor_ResizeUpCursor()
	return MakeCursor(result_)
}

func ResizeUpDownCursor() Cursor {
	result_ := C.C_NSCursor_ResizeUpDownCursor()
	return MakeCursor(result_)
}

func IBeamCursorForVerticalLayout() Cursor {
	result_ := C.C_NSCursor_IBeamCursorForVerticalLayout()
	return MakeCursor(result_)
}
