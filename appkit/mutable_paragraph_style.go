package appkit

// #include "mutable_paragraph_style.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type MutableParagraphStyle interface {
	ParagraphStyle
	SetParagraphStyle(obj ParagraphStyle)
	AddTabStop(anObject TextTab)
	RemoveTabStop(anObject TextTab)
	SetAlignment(value TextAlignment)
	SetFirstLineHeadIndent(value coregraphics.Float)
	SetHeadIndent(value coregraphics.Float)
	SetTailIndent(value coregraphics.Float)
	SetLineHeightMultiple(value coregraphics.Float)
	SetMaximumLineHeight(value coregraphics.Float)
	SetMinimumLineHeight(value coregraphics.Float)
	SetLineSpacing(value coregraphics.Float)
	SetParagraphSpacing(value coregraphics.Float)
	SetParagraphSpacingBefore(value coregraphics.Float)
	SetBaseWritingDirection(value WritingDirection)
	SetTabStops(value []TextTab)
	SetDefaultTabInterval(value coregraphics.Float)
	SetTextBlocks(value []TextBlock)
	SetTextLists(value []TextList)
	SetLineBreakMode(value LineBreakMode)
	SetLineBreakStrategy(value LineBreakStrategy)
	SetHyphenationFactor(value float32)
	SetTighteningFactorForTruncation(value float32)
	SetAllowsDefaultTighteningForTruncation(value bool)
	SetHeaderLevel(value int)
}

type NSMutableParagraphStyle struct {
	NSParagraphStyle
}

func MakeMutableParagraphStyle(ptr unsafe.Pointer) NSMutableParagraphStyle {
	return NSMutableParagraphStyle{
		NSParagraphStyle: MakeParagraphStyle(ptr),
	}
}

func AllocMutableParagraphStyle() NSMutableParagraphStyle {
	return MakeMutableParagraphStyle(C.C_MutableParagraphStyle_Alloc())
}

func (n NSMutableParagraphStyle) Init() MutableParagraphStyle {
	result_ := C.C_NSMutableParagraphStyle_Init(n.Ptr())
	return MakeMutableParagraphStyle(result_)
}

func (n NSMutableParagraphStyle) SetParagraphStyle(obj ParagraphStyle) {
	C.C_NSMutableParagraphStyle_SetParagraphStyle(n.Ptr(), objc.ExtractPtr(obj))
}

func (n NSMutableParagraphStyle) AddTabStop(anObject TextTab) {
	C.C_NSMutableParagraphStyle_AddTabStop(n.Ptr(), objc.ExtractPtr(anObject))
}

func (n NSMutableParagraphStyle) RemoveTabStop(anObject TextTab) {
	C.C_NSMutableParagraphStyle_RemoveTabStop(n.Ptr(), objc.ExtractPtr(anObject))
}

func (n NSMutableParagraphStyle) SetAlignment(value TextAlignment) {
	C.C_NSMutableParagraphStyle_SetAlignment(n.Ptr(), C.int(int(value)))
}

func (n NSMutableParagraphStyle) SetFirstLineHeadIndent(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetFirstLineHeadIndent(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetHeadIndent(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetHeadIndent(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetTailIndent(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetTailIndent(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetLineHeightMultiple(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetLineHeightMultiple(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetMaximumLineHeight(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetMaximumLineHeight(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetMinimumLineHeight(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetMinimumLineHeight(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetLineSpacing(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetLineSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetParagraphSpacing(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetParagraphSpacing(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetParagraphSpacingBefore(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetParagraphSpacingBefore(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetBaseWritingDirection(value WritingDirection) {
	C.C_NSMutableParagraphStyle_SetBaseWritingDirection(n.Ptr(), C.int(int(value)))
}

func (n NSMutableParagraphStyle) SetTabStops(value []TextTab) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSMutableParagraphStyle_SetTabStops(n.Ptr(), cValue)
}

func (n NSMutableParagraphStyle) SetDefaultTabInterval(value coregraphics.Float) {
	C.C_NSMutableParagraphStyle_SetDefaultTabInterval(n.Ptr(), C.double(float64(value)))
}

func (n NSMutableParagraphStyle) SetTextBlocks(value []TextBlock) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSMutableParagraphStyle_SetTextBlocks(n.Ptr(), cValue)
}

func (n NSMutableParagraphStyle) SetTextLists(value []TextList) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSMutableParagraphStyle_SetTextLists(n.Ptr(), cValue)
}

func (n NSMutableParagraphStyle) SetLineBreakMode(value LineBreakMode) {
	C.C_NSMutableParagraphStyle_SetLineBreakMode(n.Ptr(), C.uint(uint(value)))
}

func (n NSMutableParagraphStyle) SetLineBreakStrategy(value LineBreakStrategy) {
	C.C_NSMutableParagraphStyle_SetLineBreakStrategy(n.Ptr(), C.uint(uint(value)))
}

func (n NSMutableParagraphStyle) SetHyphenationFactor(value float32) {
	C.C_NSMutableParagraphStyle_SetHyphenationFactor(n.Ptr(), C.float(value))
}

func (n NSMutableParagraphStyle) SetTighteningFactorForTruncation(value float32) {
	C.C_NSMutableParagraphStyle_SetTighteningFactorForTruncation(n.Ptr(), C.float(value))
}

func (n NSMutableParagraphStyle) SetAllowsDefaultTighteningForTruncation(value bool) {
	C.C_NSMutableParagraphStyle_SetAllowsDefaultTighteningForTruncation(n.Ptr(), C.bool(value))
}

func (n NSMutableParagraphStyle) SetHeaderLevel(value int) {
	C.C_NSMutableParagraphStyle_SetHeaderLevel(n.Ptr(), C.int(value))
}
