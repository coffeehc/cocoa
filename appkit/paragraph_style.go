package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "paragraph_style.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ParagraphStyle interface {
	objc.Object
	Alignment() TextAlignment
	FirstLineHeadIndent() coregraphics.Float
	HeadIndent() coregraphics.Float
	TailIndent() coregraphics.Float
	LineHeightMultiple() coregraphics.Float
	MaximumLineHeight() coregraphics.Float
	MinimumLineHeight() coregraphics.Float
	LineSpacing() coregraphics.Float
	ParagraphSpacing() coregraphics.Float
	ParagraphSpacingBefore() coregraphics.Float
	TabStops() []TextTab
	DefaultTabInterval() coregraphics.Float
	TextBlocks() []TextBlock
	TextLists() []TextList
	LineBreakMode() LineBreakMode
	LineBreakStrategy() LineBreakStrategy
	HyphenationFactor() float32
	TighteningFactorForTruncation() float32
	AllowsDefaultTighteningForTruncation() bool
	HeaderLevel() int
	BaseWritingDirection() WritingDirection
}

type NSParagraphStyle struct {
	objc.NSObject
}

func MakeParagraphStyle(ptr unsafe.Pointer) *NSParagraphStyle {
	if ptr == nil {
		return nil
	}
	return &NSParagraphStyle{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocParagraphStyle() *NSParagraphStyle {
	return MakeParagraphStyle(C.C_ParagraphStyle_Alloc())
}

func (n *NSParagraphStyle) Init() ParagraphStyle {
	result := C.C_NSParagraphStyle_Init(n.Ptr())
	return MakeParagraphStyle(result)
}

func ParagraphStyle_DefaultWritingDirectionForLanguage(languageName string) WritingDirection {
	result := C.C_NSParagraphStyle_ParagraphStyle_DefaultWritingDirectionForLanguage(foundation.NewString(languageName).Ptr())
	return WritingDirection(int(result))
}

func DefaultParagraphStyle() ParagraphStyle {
	result := C.C_NSParagraphStyle_DefaultParagraphStyle()
	return MakeParagraphStyle(result)
}

func (n *NSParagraphStyle) Alignment() TextAlignment {
	result := C.C_NSParagraphStyle_Alignment(n.Ptr())
	return TextAlignment(int(result))
}

func (n *NSParagraphStyle) FirstLineHeadIndent() coregraphics.Float {
	result := C.C_NSParagraphStyle_FirstLineHeadIndent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSParagraphStyle) HeadIndent() coregraphics.Float {
	result := C.C_NSParagraphStyle_HeadIndent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSParagraphStyle) TailIndent() coregraphics.Float {
	result := C.C_NSParagraphStyle_TailIndent(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSParagraphStyle) LineHeightMultiple() coregraphics.Float {
	result := C.C_NSParagraphStyle_LineHeightMultiple(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSParagraphStyle) MaximumLineHeight() coregraphics.Float {
	result := C.C_NSParagraphStyle_MaximumLineHeight(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSParagraphStyle) MinimumLineHeight() coregraphics.Float {
	result := C.C_NSParagraphStyle_MinimumLineHeight(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSParagraphStyle) LineSpacing() coregraphics.Float {
	result := C.C_NSParagraphStyle_LineSpacing(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSParagraphStyle) ParagraphSpacing() coregraphics.Float {
	result := C.C_NSParagraphStyle_ParagraphSpacing(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSParagraphStyle) ParagraphSpacingBefore() coregraphics.Float {
	result := C.C_NSParagraphStyle_ParagraphSpacingBefore(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSParagraphStyle) TabStops() []TextTab {
	result := C.C_NSParagraphStyle_TabStops(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]TextTab, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeTextTab(r)
	}
	return goResult
}

func (n *NSParagraphStyle) DefaultTabInterval() coregraphics.Float {
	result := C.C_NSParagraphStyle_DefaultTabInterval(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSParagraphStyle) TextBlocks() []TextBlock {
	result := C.C_NSParagraphStyle_TextBlocks(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]TextBlock, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeTextBlock(r)
	}
	return goResult
}

func (n *NSParagraphStyle) TextLists() []TextList {
	result := C.C_NSParagraphStyle_TextLists(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]TextList, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeTextList(r)
	}
	return goResult
}

func (n *NSParagraphStyle) LineBreakMode() LineBreakMode {
	result := C.C_NSParagraphStyle_LineBreakMode(n.Ptr())
	return LineBreakMode(int(result))
}

func (n *NSParagraphStyle) LineBreakStrategy() LineBreakStrategy {
	result := C.C_NSParagraphStyle_LineBreakStrategy(n.Ptr())
	return LineBreakStrategy(uint(result))
}

func (n *NSParagraphStyle) HyphenationFactor() float32 {
	result := C.C_NSParagraphStyle_HyphenationFactor(n.Ptr())
	return float32(result)
}

func (n *NSParagraphStyle) TighteningFactorForTruncation() float32 {
	result := C.C_NSParagraphStyle_TighteningFactorForTruncation(n.Ptr())
	return float32(result)
}

func (n *NSParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	result := C.C_NSParagraphStyle_AllowsDefaultTighteningForTruncation(n.Ptr())
	return bool(result)
}

func (n *NSParagraphStyle) HeaderLevel() int {
	result := C.C_NSParagraphStyle_HeaderLevel(n.Ptr())
	return int(result)
}

func (n *NSParagraphStyle) BaseWritingDirection() WritingDirection {
	result := C.C_NSParagraphStyle_BaseWritingDirection(n.Ptr())
	return WritingDirection(int(result))
}
