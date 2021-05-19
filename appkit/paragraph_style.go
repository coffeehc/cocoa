package appkit

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
	result_ := C.C_NSParagraphStyle_Init(n.Ptr())
	return MakeParagraphStyle(result_)
}

func ParagraphStyle_DefaultWritingDirectionForLanguage(languageName string) WritingDirection {
	result_ := C.C_NSParagraphStyle_ParagraphStyle_DefaultWritingDirectionForLanguage(foundation.NewString(languageName).Ptr())
	return WritingDirection(int(result_))
}

func DefaultParagraphStyle() ParagraphStyle {
	result_ := C.C_NSParagraphStyle_DefaultParagraphStyle()
	return MakeParagraphStyle(result_)
}

func (n *NSParagraphStyle) Alignment() TextAlignment {
	result_ := C.C_NSParagraphStyle_Alignment(n.Ptr())
	return TextAlignment(int(result_))
}

func (n *NSParagraphStyle) FirstLineHeadIndent() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_FirstLineHeadIndent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSParagraphStyle) HeadIndent() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_HeadIndent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSParagraphStyle) TailIndent() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_TailIndent(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSParagraphStyle) LineHeightMultiple() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_LineHeightMultiple(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSParagraphStyle) MaximumLineHeight() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_MaximumLineHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSParagraphStyle) MinimumLineHeight() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_MinimumLineHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSParagraphStyle) LineSpacing() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_LineSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSParagraphStyle) ParagraphSpacing() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_ParagraphSpacing(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSParagraphStyle) ParagraphSpacingBefore() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_ParagraphSpacingBefore(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSParagraphStyle) TabStops() []TextTab {
	result_ := C.C_NSParagraphStyle_TabStops(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TextTab, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextTab(r)
	}
	return goResult_
}

func (n *NSParagraphStyle) DefaultTabInterval() coregraphics.Float {
	result_ := C.C_NSParagraphStyle_DefaultTabInterval(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSParagraphStyle) TextBlocks() []TextBlock {
	result_ := C.C_NSParagraphStyle_TextBlocks(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TextBlock, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextBlock(r)
	}
	return goResult_
}

func (n *NSParagraphStyle) TextLists() []TextList {
	result_ := C.C_NSParagraphStyle_TextLists(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TextList, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextList(r)
	}
	return goResult_
}

func (n *NSParagraphStyle) LineBreakMode() LineBreakMode {
	result_ := C.C_NSParagraphStyle_LineBreakMode(n.Ptr())
	return LineBreakMode(uint(result_))
}

func (n *NSParagraphStyle) LineBreakStrategy() LineBreakStrategy {
	result_ := C.C_NSParagraphStyle_LineBreakStrategy(n.Ptr())
	return LineBreakStrategy(uint(result_))
}

func (n *NSParagraphStyle) HyphenationFactor() float32 {
	result_ := C.C_NSParagraphStyle_HyphenationFactor(n.Ptr())
	return float32(result_)
}

func (n *NSParagraphStyle) TighteningFactorForTruncation() float32 {
	result_ := C.C_NSParagraphStyle_TighteningFactorForTruncation(n.Ptr())
	return float32(result_)
}

func (n *NSParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	result_ := C.C_NSParagraphStyle_AllowsDefaultTighteningForTruncation(n.Ptr())
	return bool(result_)
}

func (n *NSParagraphStyle) HeaderLevel() int {
	result_ := C.C_NSParagraphStyle_HeaderLevel(n.Ptr())
	return int(result_)
}

func (n *NSParagraphStyle) BaseWritingDirection() WritingDirection {
	result_ := C.C_NSParagraphStyle_BaseWritingDirection(n.Ptr())
	return WritingDirection(int(result_))
}
