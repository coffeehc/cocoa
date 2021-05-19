package appkit

// #include "typesetter.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Typesetter interface {
	objc.Object
	BaselineOffsetInLayoutManager_GlyphIndex(layoutMgr LayoutManager, glyphIndex uint) coregraphics.Float
	SubstituteFontForFont(originalFont Font) Font
	TextTabForGlyphLocation_WritingDirection_MaxLocation(glyphLocation coregraphics.Float, direction WritingDirection, maxLocation coregraphics.Float) TextTab
	SetParagraphGlyphRange_SeparatorGlyphRange(paragraphRange foundation.Range, paragraphSeparatorRange foundation.Range)
	LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) coregraphics.Float
	ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) coregraphics.Float
	ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) coregraphics.Float
	BeginParagraph()
	EndParagraph()
	BeginLineWithGlyphAtIndex(glyphIndex uint)
	EndLineWithGlyphRange(lineGlyphRange foundation.Range)
	LayoutCharactersInRange_ForLayoutManager_MaximumNumberOfLineFragments(characterRange foundation.Range, layoutManager LayoutManager, maxNumLines uint) foundation.Range
	BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect
	HyphenationFactorForGlyphAtIndex(glyphIndex uint) float32
	ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool
	ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool
	SetHardInvalidation_ForGlyphRange(flag bool, glyphRange foundation.Range)
	SetAttachmentSize_ForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range)
	SetDrawsOutsideLineFragment_ForGlyphRange(flag bool, glyphRange foundation.Range)
	SetLineFragmentRect_ForGlyphRange_UsedRect_BaselineOffset(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect, baselineOffset coregraphics.Float)
	SetNotShownAttribute_ForGlyphRange(flag bool, glyphRange foundation.Range)
	LayoutManager() LayoutManager
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
	HyphenationFactor() float32
	SetHyphenationFactor(value float32)
	CurrentTextContainer() TextContainer
	TextContainers() []TextContainer
	LineFragmentPadding() coregraphics.Float
	SetLineFragmentPadding(value coregraphics.Float)
	BidiProcessingEnabled() bool
	SetBidiProcessingEnabled(value bool)
	CurrentParagraphStyle() ParagraphStyle
	AttributedString() foundation.AttributedString
	SetAttributedString(value foundation.AttributedString)
	ParagraphGlyphRange() foundation.Range
	ParagraphSeparatorGlyphRange() foundation.Range
	ParagraphCharacterRange() foundation.Range
	ParagraphSeparatorCharacterRange() foundation.Range
}

type NSTypesetter struct {
	objc.NSObject
}

func MakeTypesetter(ptr unsafe.Pointer) *NSTypesetter {
	if ptr == nil {
		return nil
	}
	return &NSTypesetter{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTypesetter() *NSTypesetter {
	return MakeTypesetter(C.C_Typesetter_Alloc())
}

func (n *NSTypesetter) Init() Typesetter {
	result_ := C.C_NSTypesetter_Init(n.Ptr())
	return MakeTypesetter(result_)
}

func Typesetter_SharedSystemTypesetterForBehavior(behavior TypesetterBehavior) objc.Object {
	result_ := C.C_NSTypesetter_Typesetter_SharedSystemTypesetterForBehavior(C.int(int(behavior)))
	return objc.MakeObject(result_)
}

func (n *NSTypesetter) BaselineOffsetInLayoutManager_GlyphIndex(layoutMgr LayoutManager, glyphIndex uint) coregraphics.Float {
	result_ := C.C_NSTypesetter_BaselineOffsetInLayoutManager_GlyphIndex(n.Ptr(), objc.ExtractPtr(layoutMgr), C.uint(glyphIndex))
	return coregraphics.Float(float64(result_))
}

func (n *NSTypesetter) SubstituteFontForFont(originalFont Font) Font {
	result_ := C.C_NSTypesetter_SubstituteFontForFont(n.Ptr(), objc.ExtractPtr(originalFont))
	return MakeFont(result_)
}

func (n *NSTypesetter) TextTabForGlyphLocation_WritingDirection_MaxLocation(glyphLocation coregraphics.Float, direction WritingDirection, maxLocation coregraphics.Float) TextTab {
	result_ := C.C_NSTypesetter_TextTabForGlyphLocation_WritingDirection_MaxLocation(n.Ptr(), C.double(float64(glyphLocation)), C.int(int(direction)), C.double(float64(maxLocation)))
	return MakeTextTab(result_)
}

func (n *NSTypesetter) SetParagraphGlyphRange_SeparatorGlyphRange(paragraphRange foundation.Range, paragraphSeparatorRange foundation.Range) {
	C.C_NSTypesetter_SetParagraphGlyphRange_SeparatorGlyphRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(paragraphRange)), *(*C.NSRange)(foundation.ToNSRangePointer(paragraphSeparatorRange)))
}

func (n *NSTypesetter) LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) coregraphics.Float {
	result_ := C.C_NSTypesetter_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(n.Ptr(), C.uint(glyphIndex), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return coregraphics.Float(float64(result_))
}

func (n *NSTypesetter) ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) coregraphics.Float {
	result_ := C.C_NSTypesetter_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(n.Ptr(), C.uint(glyphIndex), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return coregraphics.Float(float64(result_))
}

func (n *NSTypesetter) ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) coregraphics.Float {
	result_ := C.C_NSTypesetter_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(n.Ptr(), C.uint(glyphIndex), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return coregraphics.Float(float64(result_))
}

func (n *NSTypesetter) BeginParagraph() {
	C.C_NSTypesetter_BeginParagraph(n.Ptr())
}

func (n *NSTypesetter) EndParagraph() {
	C.C_NSTypesetter_EndParagraph(n.Ptr())
}

func (n *NSTypesetter) BeginLineWithGlyphAtIndex(glyphIndex uint) {
	C.C_NSTypesetter_BeginLineWithGlyphAtIndex(n.Ptr(), C.uint(glyphIndex))
}

func (n *NSTypesetter) EndLineWithGlyphRange(lineGlyphRange foundation.Range) {
	C.C_NSTypesetter_EndLineWithGlyphRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(lineGlyphRange)))
}

func (n *NSTypesetter) LayoutCharactersInRange_ForLayoutManager_MaximumNumberOfLineFragments(characterRange foundation.Range, layoutManager LayoutManager, maxNumLines uint) foundation.Range {
	result_ := C.C_NSTypesetter_LayoutCharactersInRange_ForLayoutManager_MaximumNumberOfLineFragments(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(characterRange)), objc.ExtractPtr(layoutManager), C.uint(maxNumLines))
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSTypesetter) BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect {
	result_ := C.C_NSTypesetter_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(n.Ptr(), C.uint(glyphIndex), objc.ExtractPtr(textContainer), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(proposedRect))), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(glyphPosition))), C.uint(charIndex))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSTypesetter) HyphenationFactorForGlyphAtIndex(glyphIndex uint) float32 {
	result_ := C.C_NSTypesetter_HyphenationFactorForGlyphAtIndex(n.Ptr(), C.uint(glyphIndex))
	return float32(result_)
}

func (n *NSTypesetter) ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool {
	result_ := C.C_NSTypesetter_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(n.Ptr(), C.uint(charIndex))
	return bool(result_)
}

func (n *NSTypesetter) ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool {
	result_ := C.C_NSTypesetter_ShouldBreakLineByWordBeforeCharacterAtIndex(n.Ptr(), C.uint(charIndex))
	return bool(result_)
}

func (n *NSTypesetter) SetHardInvalidation_ForGlyphRange(flag bool, glyphRange foundation.Range) {
	C.C_NSTypesetter_SetHardInvalidation_ForGlyphRange(n.Ptr(), C.bool(flag), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func (n *NSTypesetter) SetAttachmentSize_ForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range) {
	C.C_NSTypesetter_SetAttachmentSize_ForGlyphRange(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(attachmentSize))), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func (n *NSTypesetter) SetDrawsOutsideLineFragment_ForGlyphRange(flag bool, glyphRange foundation.Range) {
	C.C_NSTypesetter_SetDrawsOutsideLineFragment_ForGlyphRange(n.Ptr(), C.bool(flag), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func (n *NSTypesetter) SetLineFragmentRect_ForGlyphRange_UsedRect_BaselineOffset(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect, baselineOffset coregraphics.Float) {
	C.C_NSTypesetter_SetLineFragmentRect_ForGlyphRange_UsedRect_BaselineOffset(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(fragmentRect))), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(usedRect))), C.double(float64(baselineOffset)))
}

func (n *NSTypesetter) SetNotShownAttribute_ForGlyphRange(flag bool, glyphRange foundation.Range) {
	C.C_NSTypesetter_SetNotShownAttribute_ForGlyphRange(n.Ptr(), C.bool(flag), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func SharedSystemTypesetter() Typesetter {
	result_ := C.C_NSTypesetter_SharedSystemTypesetter()
	return MakeTypesetter(result_)
}

func Typesetter_DefaultTypesetterBehavior() TypesetterBehavior {
	result_ := C.C_NSTypesetter_Typesetter_DefaultTypesetterBehavior()
	return TypesetterBehavior(int(result_))
}

func (n *NSTypesetter) LayoutManager() LayoutManager {
	result_ := C.C_NSTypesetter_LayoutManager(n.Ptr())
	return MakeLayoutManager(result_)
}

func (n *NSTypesetter) UsesFontLeading() bool {
	result_ := C.C_NSTypesetter_UsesFontLeading(n.Ptr())
	return bool(result_)
}

func (n *NSTypesetter) SetUsesFontLeading(value bool) {
	C.C_NSTypesetter_SetUsesFontLeading(n.Ptr(), C.bool(value))
}

func (n *NSTypesetter) TypesetterBehavior() TypesetterBehavior {
	result_ := C.C_NSTypesetter_TypesetterBehavior(n.Ptr())
	return TypesetterBehavior(int(result_))
}

func (n *NSTypesetter) SetTypesetterBehavior(value TypesetterBehavior) {
	C.C_NSTypesetter_SetTypesetterBehavior(n.Ptr(), C.int(int(value)))
}

func (n *NSTypesetter) HyphenationFactor() float32 {
	result_ := C.C_NSTypesetter_HyphenationFactor(n.Ptr())
	return float32(result_)
}

func (n *NSTypesetter) SetHyphenationFactor(value float32) {
	C.C_NSTypesetter_SetHyphenationFactor(n.Ptr(), C.float(value))
}

func (n *NSTypesetter) CurrentTextContainer() TextContainer {
	result_ := C.C_NSTypesetter_CurrentTextContainer(n.Ptr())
	return MakeTextContainer(result_)
}

func (n *NSTypesetter) TextContainers() []TextContainer {
	result_ := C.C_NSTypesetter_TextContainers(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TextContainer, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextContainer(r)
	}
	return goResult_
}

func (n *NSTypesetter) LineFragmentPadding() coregraphics.Float {
	result_ := C.C_NSTypesetter_LineFragmentPadding(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSTypesetter) SetLineFragmentPadding(value coregraphics.Float) {
	C.C_NSTypesetter_SetLineFragmentPadding(n.Ptr(), C.double(float64(value)))
}

func (n *NSTypesetter) BidiProcessingEnabled() bool {
	result_ := C.C_NSTypesetter_BidiProcessingEnabled(n.Ptr())
	return bool(result_)
}

func (n *NSTypesetter) SetBidiProcessingEnabled(value bool) {
	C.C_NSTypesetter_SetBidiProcessingEnabled(n.Ptr(), C.bool(value))
}

func (n *NSTypesetter) CurrentParagraphStyle() ParagraphStyle {
	result_ := C.C_NSTypesetter_CurrentParagraphStyle(n.Ptr())
	return MakeParagraphStyle(result_)
}

func (n *NSTypesetter) AttributedString() foundation.AttributedString {
	result_ := C.C_NSTypesetter_AttributedString(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n *NSTypesetter) SetAttributedString(value foundation.AttributedString) {
	C.C_NSTypesetter_SetAttributedString(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSTypesetter) ParagraphGlyphRange() foundation.Range {
	result_ := C.C_NSTypesetter_ParagraphGlyphRange(n.Ptr())
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSTypesetter) ParagraphSeparatorGlyphRange() foundation.Range {
	result_ := C.C_NSTypesetter_ParagraphSeparatorGlyphRange(n.Ptr())
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSTypesetter) ParagraphCharacterRange() foundation.Range {
	result_ := C.C_NSTypesetter_ParagraphCharacterRange(n.Ptr())
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSTypesetter) ParagraphSeparatorCharacterRange() foundation.Range {
	result_ := C.C_NSTypesetter_ParagraphSeparatorCharacterRange(n.Ptr())
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}
