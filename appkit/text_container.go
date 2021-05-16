package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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

func MakeTextContainer(ptr unsafe.Pointer) *NSTextContainer {
	if ptr == nil {
		return nil
	}
	return &NSTextContainer{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTextContainer() *NSTextContainer {
	return MakeTextContainer(C.C_TextContainer_Alloc())
}

func (n *NSTextContainer) InitWithSize(size foundation.Size) TextContainer {
	result := C.C_NSTextContainer_InitWithSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(size))))
	return MakeTextContainer(result)
}

func (n *NSTextContainer) InitWithCoder(coder foundation.Coder) TextContainer {
	result := C.C_NSTextContainer_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTextContainer(result)
}

func (n *NSTextContainer) Init() TextContainer {
	result := C.C_NSTextContainer_Init(n.Ptr())
	return MakeTextContainer(result)
}

func (n *NSTextContainer) ReplaceLayoutManager(newLayoutManager LayoutManager) {
	C.C_NSTextContainer_ReplaceLayoutManager(n.Ptr(), objc.ExtractPtr(newLayoutManager))
}

func (n *NSTextContainer) LayoutManager() LayoutManager {
	result := C.C_NSTextContainer_LayoutManager(n.Ptr())
	return MakeLayoutManager(result)
}

func (n *NSTextContainer) SetLayoutManager(value LayoutManager) {
	C.C_NSTextContainer_SetLayoutManager(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTextContainer) TextView() TextView {
	result := C.C_NSTextContainer_TextView(n.Ptr())
	return MakeTextView(result)
}

func (n *NSTextContainer) SetTextView(value TextView) {
	C.C_NSTextContainer_SetTextView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTextContainer) Size() foundation.Size {
	result := C.C_NSTextContainer_Size(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSTextContainer) SetSize(value foundation.Size) {
	C.C_NSTextContainer_SetSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSTextContainer) ExclusionPaths() []BezierPath {
	result := C.C_NSTextContainer_ExclusionPaths(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]BezierPath, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeBezierPath(r)
	}
	return goResult
}

func (n *NSTextContainer) SetExclusionPaths(value []BezierPath) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSTextContainer_SetExclusionPaths(n.Ptr(), cValue)
}

func (n *NSTextContainer) LineBreakMode() LineBreakMode {
	result := C.C_NSTextContainer_LineBreakMode(n.Ptr())
	return LineBreakMode(uint(result))
}

func (n *NSTextContainer) SetLineBreakMode(value LineBreakMode) {
	C.C_NSTextContainer_SetLineBreakMode(n.Ptr(), C.uint(uint(value)))
}

func (n *NSTextContainer) WidthTracksTextView() bool {
	result := C.C_NSTextContainer_WidthTracksTextView(n.Ptr())
	return bool(result)
}

func (n *NSTextContainer) SetWidthTracksTextView(value bool) {
	C.C_NSTextContainer_SetWidthTracksTextView(n.Ptr(), C.bool(value))
}

func (n *NSTextContainer) HeightTracksTextView() bool {
	result := C.C_NSTextContainer_HeightTracksTextView(n.Ptr())
	return bool(result)
}

func (n *NSTextContainer) SetHeightTracksTextView(value bool) {
	C.C_NSTextContainer_SetHeightTracksTextView(n.Ptr(), C.bool(value))
}

func (n *NSTextContainer) MaximumNumberOfLines() uint {
	result := C.C_NSTextContainer_MaximumNumberOfLines(n.Ptr())
	return uint(result)
}

func (n *NSTextContainer) SetMaximumNumberOfLines(value uint) {
	C.C_NSTextContainer_SetMaximumNumberOfLines(n.Ptr(), C.uint(value))
}

func (n *NSTextContainer) LineFragmentPadding() coregraphics.Float {
	result := C.C_NSTextContainer_LineFragmentPadding(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSTextContainer) SetLineFragmentPadding(value coregraphics.Float) {
	C.C_NSTextContainer_SetLineFragmentPadding(n.Ptr(), C.double(float64(value)))
}

func (n *NSTextContainer) IsSimpleRectangularTextContainer() bool {
	result := C.C_NSTextContainer_IsSimpleRectangularTextContainer(n.Ptr())
	return bool(result)
}
