package appkit

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
	DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha coregraphics.Float, respectContextIsFlipped bool, hints map[ImageHintKey]objc.Object) bool
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
	LayoutDirection() ImageLayoutDirection
	SetLayoutDirection(value ImageLayoutDirection)
}

type NSImageRep struct {
	objc.NSObject
}

func MakeImageRep(ptr unsafe.Pointer) NSImageRep {
	return NSImageRep{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocImageRep() NSImageRep {
	return MakeImageRep(C.C_ImageRep_Alloc())
}

func (n NSImageRep) Init() ImageRep {
	result_ := C.C_NSImageRep_Init(n.Ptr())
	return MakeImageRep(result_)
}

func (n NSImageRep) InitWithCoder(coder foundation.Coder) ImageRep {
	result_ := C.C_NSImageRep_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeImageRep(result_)
}

func ImageRepsWithContentsOfFile(filename string) []ImageRep {
	result_ := C.C_NSImageRep_ImageRepsWithContentsOfFile(foundation.NewString(filename).Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]ImageRep, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeImageRep(r)
	}
	return goResult_
}

func ImageRepsWithPasteboard(pasteboard Pasteboard) []ImageRep {
	result_ := C.C_NSImageRep_ImageRepsWithPasteboard(objc.ExtractPtr(pasteboard))
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]ImageRep, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeImageRep(r)
	}
	return goResult_
}

func ImageRepsWithContentsOfURL(url foundation.URL) []ImageRep {
	result_ := C.C_NSImageRep_ImageRepsWithContentsOfURL(objc.ExtractPtr(url))
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]ImageRep, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeImageRep(r)
	}
	return goResult_
}

func ImageRepWithContentsOfFile(filename string) ImageRep {
	result_ := C.C_NSImageRep_ImageRepWithContentsOfFile(foundation.NewString(filename).Ptr())
	return MakeImageRep(result_)
}

func ImageRepWithPasteboard(pasteboard Pasteboard) ImageRep {
	result_ := C.C_NSImageRep_ImageRepWithPasteboard(objc.ExtractPtr(pasteboard))
	return MakeImageRep(result_)
}

func ImageRepWithContentsOfURL(url foundation.URL) ImageRep {
	result_ := C.C_NSImageRep_ImageRepWithContentsOfURL(objc.ExtractPtr(url))
	return MakeImageRep(result_)
}

func ImageRep_CanInitWithData(data []byte) bool {
	result_ := C.C_NSImageRep_ImageRep_CanInitWithData(C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return bool(result_)
}

func ImageRep_CanInitWithPasteboard(pasteboard Pasteboard) bool {
	result_ := C.C_NSImageRep_ImageRep_CanInitWithPasteboard(objc.ExtractPtr(pasteboard))
	return bool(result_)
}

func (n NSImageRep) Draw() bool {
	result_ := C.C_NSImageRep_Draw(n.Ptr())
	return bool(result_)
}

func (n NSImageRep) DrawAtPoint(point foundation.Point) bool {
	result_ := C.C_NSImageRep_DrawAtPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return bool(result_)
}

func (n NSImageRep) DrawInRect(rect foundation.Rect) bool {
	result_ := C.C_NSImageRep_DrawInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return bool(result_)
}

func (n NSImageRep) DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(dstSpacePortionRect foundation.Rect, srcSpacePortionRect foundation.Rect, op CompositingOperation, requestedAlpha coregraphics.Float, respectContextIsFlipped bool, hints map[ImageHintKey]objc.Object) bool {
	var cHints C.Dictionary
	if len(hints) > 0 {
		cHintsKeyData := make([]unsafe.Pointer, len(hints))
		cHintsValueData := make([]unsafe.Pointer, len(hints))
		var idx = 0
		for k, v := range hints {
			cHintsKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cHintsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cHints.key_data = unsafe.Pointer(&cHintsKeyData[0])
		cHints.value_data = unsafe.Pointer(&cHintsValueData[0])
		cHints.len = C.int(len(hints))
	}
	result_ := C.C_NSImageRep_DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(dstSpacePortionRect))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(srcSpacePortionRect))), C.uint(uint(op)), C.double(float64(requestedAlpha)), C.bool(respectContextIsFlipped), cHints)
	return bool(result_)
}

func ImageRep_ImageTypes() []string {
	result_ := C.C_NSImageRep_ImageRep_ImageTypes()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func ImageRep_ImageUnfilteredTypes() []string {
	result_ := C.C_NSImageRep_ImageRep_ImageUnfilteredTypes()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSImageRep) Size() foundation.Size {
	result_ := C.C_NSImageRep_Size(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSImageRep) SetSize(value foundation.Size) {
	C.C_NSImageRep_SetSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSImageRep) BitsPerSample() int {
	result_ := C.C_NSImageRep_BitsPerSample(n.Ptr())
	return int(result_)
}

func (n NSImageRep) SetBitsPerSample(value int) {
	C.C_NSImageRep_SetBitsPerSample(n.Ptr(), C.int(value))
}

func (n NSImageRep) ColorSpaceName() ColorSpaceName {
	result_ := C.C_NSImageRep_ColorSpaceName(n.Ptr())
	return ColorSpaceName(foundation.MakeString(result_).String())
}

func (n NSImageRep) SetColorSpaceName(value ColorSpaceName) {
	C.C_NSImageRep_SetColorSpaceName(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSImageRep) HasAlpha() bool {
	result_ := C.C_NSImageRep_HasAlpha(n.Ptr())
	return bool(result_)
}

func (n NSImageRep) SetAlpha(value bool) {
	C.C_NSImageRep_SetAlpha(n.Ptr(), C.bool(value))
}

func (n NSImageRep) IsOpaque() bool {
	result_ := C.C_NSImageRep_IsOpaque(n.Ptr())
	return bool(result_)
}

func (n NSImageRep) SetOpaque(value bool) {
	C.C_NSImageRep_SetOpaque(n.Ptr(), C.bool(value))
}

func (n NSImageRep) PixelsHigh() int {
	result_ := C.C_NSImageRep_PixelsHigh(n.Ptr())
	return int(result_)
}

func (n NSImageRep) SetPixelsHigh(value int) {
	C.C_NSImageRep_SetPixelsHigh(n.Ptr(), C.int(value))
}

func (n NSImageRep) PixelsWide() int {
	result_ := C.C_NSImageRep_PixelsWide(n.Ptr())
	return int(result_)
}

func (n NSImageRep) SetPixelsWide(value int) {
	C.C_NSImageRep_SetPixelsWide(n.Ptr(), C.int(value))
}

func (n NSImageRep) LayoutDirection() ImageLayoutDirection {
	result_ := C.C_NSImageRep_LayoutDirection(n.Ptr())
	return ImageLayoutDirection(int(result_))
}

func (n NSImageRep) SetLayoutDirection(value ImageLayoutDirection) {
	C.C_NSImageRep_SetLayoutDirection(n.Ptr(), C.int(int(value)))
}
