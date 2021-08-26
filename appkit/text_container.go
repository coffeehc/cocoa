package appkit

// #include "text_container.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TextContainer interface {
	objc.Object
	ReplaceLayoutManager(newLayoutManager LayoutManager)
	LayoutManager() LayoutManager
	SetLayoutManager(value LayoutManager)
	TextView() TextView
	SetTextView(value TextView)
	Size() foundation.Size
	SetSize(value foundation.Size)
	ExclusionPaths() []BezierPath
	SetExclusionPaths(value []BezierPath)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	WidthTracksTextView() bool
	SetWidthTracksTextView(value bool)
	HeightTracksTextView() bool
	SetHeightTracksTextView(value bool)
	MaximumNumberOfLines() uint
	SetMaximumNumberOfLines(value uint)
	LineFragmentPadding() coregraphics.Float
	SetLineFragmentPadding(value coregraphics.Float)
	IsSimpleRectangularTextContainer() bool
}

type NSTextContainer struct {
	objc.NSObject
}

func MakeTextContainer(ptr unsafe.Pointer) NSTextContainer {
	return NSTextContainer{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocTextContainer() NSTextContainer {
	return MakeTextContainer(C.C_TextContainer_Alloc())
}

func (n NSTextContainer) InitWithSize(size foundation.Size) TextContainer {
	result_ := C.C_NSTextContainer_InitWithSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return MakeTextContainer(result_)
}

func (n NSTextContainer) InitWithCoder(coder foundation.Coder) TextContainer {
	result_ := C.C_NSTextContainer_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTextContainer(result_)
}

func (n NSTextContainer) Init() TextContainer {
	result_ := C.C_NSTextContainer_Init(n.Ptr())
	return MakeTextContainer(result_)
}

func (n NSTextContainer) ReplaceLayoutManager(newLayoutManager LayoutManager) {
	C.C_NSTextContainer_ReplaceLayoutManager(n.Ptr(), objc.ExtractPtr(newLayoutManager))
}

func (n NSTextContainer) LayoutManager() LayoutManager {
	result_ := C.C_NSTextContainer_LayoutManager(n.Ptr())
	return MakeLayoutManager(result_)
}

func (n NSTextContainer) SetLayoutManager(value LayoutManager) {
	C.C_NSTextContainer_SetLayoutManager(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextContainer) TextView() TextView {
	result_ := C.C_NSTextContainer_TextView(n.Ptr())
	return MakeTextView(result_)
}

func (n NSTextContainer) SetTextView(value TextView) {
	C.C_NSTextContainer_SetTextView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextContainer) Size() foundation.Size {
	result_ := C.C_NSTextContainer_Size(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSTextContainer) SetSize(value foundation.Size) {
	C.C_NSTextContainer_SetSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSTextContainer) ExclusionPaths() []BezierPath {
	result_ := C.C_NSTextContainer_ExclusionPaths(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]BezierPath, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeBezierPath(r)
	}
	return goResult_
}

func (n NSTextContainer) SetExclusionPaths(value []BezierPath) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTextContainer_SetExclusionPaths(n.Ptr(), cValue)
}

func (n NSTextContainer) LineBreakMode() LineBreakMode {
	result_ := C.C_NSTextContainer_LineBreakMode(n.Ptr())
	return LineBreakMode(uint(result_))
}

func (n NSTextContainer) SetLineBreakMode(value LineBreakMode) {
	C.C_NSTextContainer_SetLineBreakMode(n.Ptr(), C.uint(uint(value)))
}

func (n NSTextContainer) WidthTracksTextView() bool {
	result_ := C.C_NSTextContainer_WidthTracksTextView(n.Ptr())
	return bool(result_)
}

func (n NSTextContainer) SetWidthTracksTextView(value bool) {
	C.C_NSTextContainer_SetWidthTracksTextView(n.Ptr(), C.bool(value))
}

func (n NSTextContainer) HeightTracksTextView() bool {
	result_ := C.C_NSTextContainer_HeightTracksTextView(n.Ptr())
	return bool(result_)
}

func (n NSTextContainer) SetHeightTracksTextView(value bool) {
	C.C_NSTextContainer_SetHeightTracksTextView(n.Ptr(), C.bool(value))
}

func (n NSTextContainer) MaximumNumberOfLines() uint {
	result_ := C.C_NSTextContainer_MaximumNumberOfLines(n.Ptr())
	return uint(result_)
}

func (n NSTextContainer) SetMaximumNumberOfLines(value uint) {
	C.C_NSTextContainer_SetMaximumNumberOfLines(n.Ptr(), C.uint(value))
}

func (n NSTextContainer) LineFragmentPadding() coregraphics.Float {
	result_ := C.C_NSTextContainer_LineFragmentPadding(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSTextContainer) SetLineFragmentPadding(value coregraphics.Float) {
	C.C_NSTextContainer_SetLineFragmentPadding(n.Ptr(), C.double(float64(value)))
}

func (n NSTextContainer) IsSimpleRectangularTextContainer() bool {
	result_ := C.C_NSTextContainer_IsSimpleRectangularTextContainer(n.Ptr())
	return bool(result_)
}
