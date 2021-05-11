package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "image_rep.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ImageRep interface {
	objc.Object
	Draw() bool
	DrawAtPoint(point foundation.Point) bool
	DrawInRect(rect foundation.Rect) bool
	Size() foundation.Size
	SetSize(value foundation.Size)
	BitsPerSample() int
	SetBitsPerSample(value int)
	ColorSpaceName() ColorSpaceName
	SetColorSpaceName(value ColorSpaceName)
	HasAlpha() bool
	SetAlpha(value bool)
	IsOpaque() bool
	SetOpaque(value bool)
	PixelsHigh() int
	SetPixelsHigh(value int)
	PixelsWide() int
	SetPixelsWide(value int)
}

type NSImageRep struct {
	objc.NSObject
}

func MakeImageRep(ptr unsafe.Pointer) *NSImageRep {
	if ptr == nil {
		return nil
	}
	return &NSImageRep{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocImageRep() *NSImageRep {
	return MakeImageRep(C.C_ImageRep_Alloc())
}

func (n *NSImageRep) Init() ImageRep {
	result := C.C_NSImageRep_Init(n.Ptr())
	return MakeImageRep(result)
}

func (n *NSImageRep) InitWithCoder(coder foundation.Coder) ImageRep {
	result := C.C_NSImageRep_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeImageRep(result)
}

func ImageRepWithContentsOfFile(filename string) ImageRep {
	result := C.C_NSImageRep_ImageRepWithContentsOfFile(foundation.NewString(filename).Ptr())
	return MakeImageRep(result)
}

func ImageRepWithPasteboard(pasteboard Pasteboard) ImageRep {
	result := C.C_NSImageRep_ImageRepWithPasteboard(objc.ExtractPtr(pasteboard))
	return MakeImageRep(result)
}

func ImageRepWithContentsOfURL(url foundation.URL) ImageRep {
	result := C.C_NSImageRep_ImageRepWithContentsOfURL(objc.ExtractPtr(url))
	return MakeImageRep(result)
}

func ImageRepCanInitWithData(data []byte) bool {
	result := C.C_NSImageRep_ImageRepCanInitWithData(C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return bool(result)
}

func ImageRepCanInitWithPasteboard(pasteboard Pasteboard) bool {
	result := C.C_NSImageRep_ImageRepCanInitWithPasteboard(objc.ExtractPtr(pasteboard))
	return bool(result)
}

func (n *NSImageRep) Draw() bool {
	result := C.C_NSImageRep_Draw(n.Ptr())
	return bool(result)
}

func (n *NSImageRep) DrawAtPoint(point foundation.Point) bool {
	result := C.C_NSImageRep_DrawAtPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return bool(result)
}

func (n *NSImageRep) DrawInRect(rect foundation.Rect) bool {
	result := C.C_NSImageRep_DrawInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return bool(result)
}

func (n *NSImageRep) Size() foundation.Size {
	result := C.C_NSImageRep_Size(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSImageRep) SetSize(value foundation.Size) {
	C.C_NSImageRep_SetSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSImageRep) BitsPerSample() int {
	result := C.C_NSImageRep_BitsPerSample(n.Ptr())
	return int(result)
}

func (n *NSImageRep) SetBitsPerSample(value int) {
	C.C_NSImageRep_SetBitsPerSample(n.Ptr(), C.int(value))
}

func (n *NSImageRep) ColorSpaceName() ColorSpaceName {
	result := C.C_NSImageRep_ColorSpaceName(n.Ptr())
	return ColorSpaceName(foundation.MakeString(result).String())
}

func (n *NSImageRep) SetColorSpaceName(value ColorSpaceName) {
	C.C_NSImageRep_SetColorSpaceName(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSImageRep) HasAlpha() bool {
	result := C.C_NSImageRep_HasAlpha(n.Ptr())
	return bool(result)
}

func (n *NSImageRep) SetAlpha(value bool) {
	C.C_NSImageRep_SetAlpha(n.Ptr(), C.bool(value))
}

func (n *NSImageRep) IsOpaque() bool {
	result := C.C_NSImageRep_IsOpaque(n.Ptr())
	return bool(result)
}

func (n *NSImageRep) SetOpaque(value bool) {
	C.C_NSImageRep_SetOpaque(n.Ptr(), C.bool(value))
}

func (n *NSImageRep) PixelsHigh() int {
	result := C.C_NSImageRep_PixelsHigh(n.Ptr())
	return int(result)
}

func (n *NSImageRep) SetPixelsHigh(value int) {
	C.C_NSImageRep_SetPixelsHigh(n.Ptr(), C.int(value))
}

func (n *NSImageRep) PixelsWide() int {
	result := C.C_NSImageRep_PixelsWide(n.Ptr())
	return int(result)
}

func (n *NSImageRep) SetPixelsWide(value int) {
	C.C_NSImageRep_SetPixelsWide(n.Ptr(), C.int(value))
}
