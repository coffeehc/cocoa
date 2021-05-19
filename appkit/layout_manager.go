package appkit

// #include "layout_manager.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type LayoutManager interface {
	objc.Object
	ReplaceTextStorage(newTextStorage TextStorage)
	AddTextContainer(container TextContainer)
	InsertTextContainer_AtIndex(container TextContainer, index uint)
	RemoveTextContainerAtIndex(index uint)
	SetTextContainer_ForGlyphRange(container TextContainer, glyphRange foundation.Range)
	TextContainerChangedGeometry(container TextContainer)
	TextContainerChangedTextView(container TextContainer)
	UsedRectForTextContainer(container TextContainer) foundation.Rect
	InvalidateDisplayForCharacterRange(charRange foundation.Range)
	InvalidateDisplayForGlyphRange(glyphRange foundation.Range)
	ProcessEditingForTextStorage_Edited_Range_ChangeInLength_InvalidatedRange(textStorage TextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range)
	EnsureGlyphsForCharacterRange(charRange foundation.Range)
	EnsureGlyphsForGlyphRange(glyphRange foundation.Range)
	EnsureLayoutForBoundingRect_InTextContainer(bounds foundation.Rect, container TextContainer)
	EnsureLayoutForCharacterRange(charRange foundation.Range)
	EnsureLayoutForGlyphRange(glyphRange foundation.Range)
	EnsureLayoutForTextContainer(container TextContainer)
	CharacterIndexForGlyphAtIndex(glyphIndex uint) uint
	GlyphIndexForCharacterAtIndex(charIndex uint) uint
	IsValidGlyphIndex(glyphIndex uint) bool
	PropertyForGlyphAtIndex(glyphIndex uint) GlyphProperty
	SetAttachmentSize_ForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range)
	SetDrawsOutsideLineFragment_ForGlyphAtIndex(flag bool, glyphIndex uint)
	SetExtraLineFragmentRect_UsedRect_TextContainer(fragmentRect foundation.Rect, usedRect foundation.Rect, container TextContainer)
	SetLineFragmentRect_ForGlyphRange_UsedRect(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect)
	SetLocation_ForStartOfGlyphRange(location foundation.Point, glyphRange foundation.Range)
	SetNotShownAttribute_ForGlyphAtIndex(flag bool, glyphIndex uint)
	AttachmentSizeForGlyphAtIndex(glyphIndex uint) foundation.Size
	DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool
	FirstUnlaidCharacterIndex() uint
	FirstUnlaidGlyphIndex() uint
	LocationForGlyphAtIndex(glyphIndex uint) foundation.Point
	NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool
	TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) foundation.Range
	BoundingRectForGlyphRange_InTextContainer(glyphRange foundation.Range, container TextContainer) foundation.Rect
	FractionOfDistanceThroughGlyphForPoint_InTextContainer(point foundation.Point, container TextContainer) coregraphics.Float
	GlyphIndexForPoint_InTextContainer(point foundation.Point, container TextContainer) uint
	GlyphRangeForBoundingRect_InTextContainer(bounds foundation.Rect, container TextContainer) foundation.Range
	GlyphRangeForBoundingRectWithoutAdditionalLayout_InTextContainer(bounds foundation.Rect, container TextContainer) foundation.Range
	GlyphRangeForTextContainer(container TextContainer) foundation.Range
	RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) foundation.Range
	DrawBackgroundForGlyphRange_AtPoint(glyphsToShow foundation.Range, origin foundation.Point)
	DrawGlyphsForGlyphRange_AtPoint(glyphsToShow foundation.Range, origin foundation.Point)
	DrawStrikethroughForGlyphRange_StrikethroughType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, baselineOffset coregraphics.Float, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	DrawUnderlineForGlyphRange_UnderlineType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, baselineOffset coregraphics.Float, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	StrikethroughGlyphRange_StrikethroughType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	UnderlineGlyphRange_UnderlineType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	SetLayoutRect_ForTextBlock_GlyphRange(rect foundation.Rect, block TextBlock, glyphRange foundation.Range)
	LayoutRectForTextBlock_GlyphRange(block TextBlock, glyphRange foundation.Range) foundation.Rect
	SetBoundsRect_ForTextBlock_GlyphRange(rect foundation.Rect, block TextBlock, glyphRange foundation.Range)
	BoundsRectForTextBlock_GlyphRange(block TextBlock, glyphRange foundation.Range) foundation.Rect
	ShowAttachmentCell_InRect_CharacterIndex(cell Cell, rect foundation.Rect, attachmentIndex uint)
	RulerAccessoryViewForTextView_ParagraphStyle_Ruler_Enabled(view TextView, style ParagraphStyle, ruler RulerView, isEnabled bool) View
	RulerMarkersForTextView_ParagraphStyle_Ruler(view TextView, style ParagraphStyle, ruler RulerView) []RulerMarker
	LayoutManagerOwnsFirstResponderInWindow(window Window) bool
	DefaultLineHeightForFont(theFont Font) coregraphics.Float
	DefaultBaselineOffsetForFont(theFont Font) coregraphics.Float
	AddTemporaryAttribute_Value_ForCharacterRange(attrName foundation.AttributedStringKey, value objc.Object, charRange foundation.Range)
	RemoveTemporaryAttribute_ForCharacterRange(attrName foundation.AttributedStringKey, charRange foundation.Range)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	TextStorage() TextStorage
	SetTextStorage(value TextStorage)
	AllowsNonContiguousLayout() bool
	SetAllowsNonContiguousLayout(value bool)
	HasNonContiguousLayout() bool
	ShowsInvisibleCharacters() bool
	SetShowsInvisibleCharacters(value bool)
	ShowsControlCharacters() bool
	SetShowsControlCharacters(value bool)
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	BackgroundLayoutEnabled() bool
	SetBackgroundLayoutEnabled(value bool)
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(value bool)
	UsesDefaultHyphenation() bool
	SetUsesDefaultHyphenation(value bool)
	TextContainers() []TextContainer
	GlyphGenerator() GlyphGenerator
	SetGlyphGenerator(value GlyphGenerator)
	NumberOfGlyphs() uint
	ExtraLineFragmentRect() foundation.Rect
	ExtraLineFragmentTextContainer() TextContainer
	ExtraLineFragmentUsedRect() foundation.Rect
	DefaultAttachmentScaling() ImageScaling
	SetDefaultAttachmentScaling(value ImageScaling)
	FirstTextView() TextView
	TextViewForBeginningOfSelection() TextView
	Typesetter() Typesetter
	SetTypesetter(value Typesetter)
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
}

type NSLayoutManager struct {
	objc.NSObject
}

func MakeLayoutManager(ptr unsafe.Pointer) *NSLayoutManager {
	if ptr == nil {
		return nil
	}
	return &NSLayoutManager{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocLayoutManager() *NSLayoutManager {
	return MakeLayoutManager(C.C_LayoutManager_Alloc())
}

func (n *NSLayoutManager) Init() LayoutManager {
	result_ := C.C_NSLayoutManager_Init(n.Ptr())
	return MakeLayoutManager(result_)
}

func (n *NSLayoutManager) InitWithCoder(coder foundation.Coder) LayoutManager {
	result_ := C.C_NSLayoutManager_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeLayoutManager(result_)
}

func (n *NSLayoutManager) ReplaceTextStorage(newTextStorage TextStorage) {
	C.C_NSLayoutManager_ReplaceTextStorage(n.Ptr(), objc.ExtractPtr(newTextStorage))
}

func (n *NSLayoutManager) AddTextContainer(container TextContainer) {
	C.C_NSLayoutManager_AddTextContainer(n.Ptr(), objc.ExtractPtr(container))
}

func (n *NSLayoutManager) InsertTextContainer_AtIndex(container TextContainer, index uint) {
	C.C_NSLayoutManager_InsertTextContainer_AtIndex(n.Ptr(), objc.ExtractPtr(container), C.uint(index))
}

func (n *NSLayoutManager) RemoveTextContainerAtIndex(index uint) {
	C.C_NSLayoutManager_RemoveTextContainerAtIndex(n.Ptr(), C.uint(index))
}

func (n *NSLayoutManager) SetTextContainer_ForGlyphRange(container TextContainer, glyphRange foundation.Range) {
	C.C_NSLayoutManager_SetTextContainer_ForGlyphRange(n.Ptr(), objc.ExtractPtr(container), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func (n *NSLayoutManager) TextContainerChangedGeometry(container TextContainer) {
	C.C_NSLayoutManager_TextContainerChangedGeometry(n.Ptr(), objc.ExtractPtr(container))
}

func (n *NSLayoutManager) TextContainerChangedTextView(container TextContainer) {
	C.C_NSLayoutManager_TextContainerChangedTextView(n.Ptr(), objc.ExtractPtr(container))
}

func (n *NSLayoutManager) UsedRectForTextContainer(container TextContainer) foundation.Rect {
	result_ := C.C_NSLayoutManager_UsedRectForTextContainer(n.Ptr(), objc.ExtractPtr(container))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSLayoutManager) InvalidateDisplayForCharacterRange(charRange foundation.Range) {
	C.C_NSLayoutManager_InvalidateDisplayForCharacterRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(charRange)))
}

func (n *NSLayoutManager) InvalidateDisplayForGlyphRange(glyphRange foundation.Range) {
	C.C_NSLayoutManager_InvalidateDisplayForGlyphRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func (n *NSLayoutManager) ProcessEditingForTextStorage_Edited_Range_ChangeInLength_InvalidatedRange(textStorage TextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range) {
	C.C_NSLayoutManager_ProcessEditingForTextStorage_Edited_Range_ChangeInLength_InvalidatedRange(n.Ptr(), objc.ExtractPtr(textStorage), C.uint(uint(editMask)), *(*C.NSRange)(foundation.ToNSRangePointer(newCharRange)), C.int(delta), *(*C.NSRange)(foundation.ToNSRangePointer(invalidatedCharRange)))
}

func (n *NSLayoutManager) EnsureGlyphsForCharacterRange(charRange foundation.Range) {
	C.C_NSLayoutManager_EnsureGlyphsForCharacterRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(charRange)))
}

func (n *NSLayoutManager) EnsureGlyphsForGlyphRange(glyphRange foundation.Range) {
	C.C_NSLayoutManager_EnsureGlyphsForGlyphRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func (n *NSLayoutManager) EnsureLayoutForBoundingRect_InTextContainer(bounds foundation.Rect, container TextContainer) {
	C.C_NSLayoutManager_EnsureLayoutForBoundingRect_InTextContainer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(bounds))), objc.ExtractPtr(container))
}

func (n *NSLayoutManager) EnsureLayoutForCharacterRange(charRange foundation.Range) {
	C.C_NSLayoutManager_EnsureLayoutForCharacterRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(charRange)))
}

func (n *NSLayoutManager) EnsureLayoutForGlyphRange(glyphRange foundation.Range) {
	C.C_NSLayoutManager_EnsureLayoutForGlyphRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func (n *NSLayoutManager) EnsureLayoutForTextContainer(container TextContainer) {
	C.C_NSLayoutManager_EnsureLayoutForTextContainer(n.Ptr(), objc.ExtractPtr(container))
}

func (n *NSLayoutManager) CharacterIndexForGlyphAtIndex(glyphIndex uint) uint {
	result_ := C.C_NSLayoutManager_CharacterIndexForGlyphAtIndex(n.Ptr(), C.uint(glyphIndex))
	return uint(result_)
}

func (n *NSLayoutManager) GlyphIndexForCharacterAtIndex(charIndex uint) uint {
	result_ := C.C_NSLayoutManager_GlyphIndexForCharacterAtIndex(n.Ptr(), C.uint(charIndex))
	return uint(result_)
}

func (n *NSLayoutManager) IsValidGlyphIndex(glyphIndex uint) bool {
	result_ := C.C_NSLayoutManager_IsValidGlyphIndex(n.Ptr(), C.uint(glyphIndex))
	return bool(result_)
}

func (n *NSLayoutManager) PropertyForGlyphAtIndex(glyphIndex uint) GlyphProperty {
	result_ := C.C_NSLayoutManager_PropertyForGlyphAtIndex(n.Ptr(), C.uint(glyphIndex))
	return GlyphProperty(int(result_))
}

func (n *NSLayoutManager) SetAttachmentSize_ForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range) {
	C.C_NSLayoutManager_SetAttachmentSize_ForGlyphRange(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(attachmentSize))), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func (n *NSLayoutManager) SetDrawsOutsideLineFragment_ForGlyphAtIndex(flag bool, glyphIndex uint) {
	C.C_NSLayoutManager_SetDrawsOutsideLineFragment_ForGlyphAtIndex(n.Ptr(), C.bool(flag), C.uint(glyphIndex))
}

func (n *NSLayoutManager) SetExtraLineFragmentRect_UsedRect_TextContainer(fragmentRect foundation.Rect, usedRect foundation.Rect, container TextContainer) {
	C.C_NSLayoutManager_SetExtraLineFragmentRect_UsedRect_TextContainer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(fragmentRect))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(usedRect))), objc.ExtractPtr(container))
}

func (n *NSLayoutManager) SetLineFragmentRect_ForGlyphRange_UsedRect(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect) {
	C.C_NSLayoutManager_SetLineFragmentRect_ForGlyphRange_UsedRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(fragmentRect))), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(usedRect))))
}

func (n *NSLayoutManager) SetLocation_ForStartOfGlyphRange(location foundation.Point, glyphRange foundation.Range) {
	C.C_NSLayoutManager_SetLocation_ForStartOfGlyphRange(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(location))), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func (n *NSLayoutManager) SetNotShownAttribute_ForGlyphAtIndex(flag bool, glyphIndex uint) {
	C.C_NSLayoutManager_SetNotShownAttribute_ForGlyphAtIndex(n.Ptr(), C.bool(flag), C.uint(glyphIndex))
}

func (n *NSLayoutManager) AttachmentSizeForGlyphAtIndex(glyphIndex uint) foundation.Size {
	result_ := C.C_NSLayoutManager_AttachmentSizeForGlyphAtIndex(n.Ptr(), C.uint(glyphIndex))
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSLayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool {
	result_ := C.C_NSLayoutManager_DrawsOutsideLineFragmentForGlyphAtIndex(n.Ptr(), C.uint(glyphIndex))
	return bool(result_)
}

func (n *NSLayoutManager) FirstUnlaidCharacterIndex() uint {
	result_ := C.C_NSLayoutManager_FirstUnlaidCharacterIndex(n.Ptr())
	return uint(result_)
}

func (n *NSLayoutManager) FirstUnlaidGlyphIndex() uint {
	result_ := C.C_NSLayoutManager_FirstUnlaidGlyphIndex(n.Ptr())
	return uint(result_)
}

func (n *NSLayoutManager) LocationForGlyphAtIndex(glyphIndex uint) foundation.Point {
	result_ := C.C_NSLayoutManager_LocationForGlyphAtIndex(n.Ptr(), C.uint(glyphIndex))
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n *NSLayoutManager) NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool {
	result_ := C.C_NSLayoutManager_NotShownAttributeForGlyphAtIndex(n.Ptr(), C.uint(glyphIndex))
	return bool(result_)
}

func (n *NSLayoutManager) TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) foundation.Range {
	result_ := C.C_NSLayoutManager_TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(n.Ptr(), C.uint(glyphIndex))
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSLayoutManager) BoundingRectForGlyphRange_InTextContainer(glyphRange foundation.Range, container TextContainer) foundation.Rect {
	result_ := C.C_NSLayoutManager_BoundingRectForGlyphRange_InTextContainer(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)), objc.ExtractPtr(container))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSLayoutManager) FractionOfDistanceThroughGlyphForPoint_InTextContainer(point foundation.Point, container TextContainer) coregraphics.Float {
	result_ := C.C_NSLayoutManager_FractionOfDistanceThroughGlyphForPoint_InTextContainer(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), objc.ExtractPtr(container))
	return coregraphics.Float(float64(result_))
}

func (n *NSLayoutManager) GlyphIndexForPoint_InTextContainer(point foundation.Point, container TextContainer) uint {
	result_ := C.C_NSLayoutManager_GlyphIndexForPoint_InTextContainer(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))), objc.ExtractPtr(container))
	return uint(result_)
}

func (n *NSLayoutManager) GlyphRangeForBoundingRect_InTextContainer(bounds foundation.Rect, container TextContainer) foundation.Range {
	result_ := C.C_NSLayoutManager_GlyphRangeForBoundingRect_InTextContainer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(bounds))), objc.ExtractPtr(container))
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSLayoutManager) GlyphRangeForBoundingRectWithoutAdditionalLayout_InTextContainer(bounds foundation.Rect, container TextContainer) foundation.Range {
	result_ := C.C_NSLayoutManager_GlyphRangeForBoundingRectWithoutAdditionalLayout_InTextContainer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(bounds))), objc.ExtractPtr(container))
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSLayoutManager) GlyphRangeForTextContainer(container TextContainer) foundation.Range {
	result_ := C.C_NSLayoutManager_GlyphRangeForTextContainer(n.Ptr(), objc.ExtractPtr(container))
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSLayoutManager) RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) foundation.Range {
	result_ := C.C_NSLayoutManager_RangeOfNominallySpacedGlyphsContainingIndex(n.Ptr(), C.uint(glyphIndex))
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n *NSLayoutManager) DrawBackgroundForGlyphRange_AtPoint(glyphsToShow foundation.Range, origin foundation.Point) {
	C.C_NSLayoutManager_DrawBackgroundForGlyphRange_AtPoint(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(glyphsToShow)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(origin))))
}

func (n *NSLayoutManager) DrawGlyphsForGlyphRange_AtPoint(glyphsToShow foundation.Range, origin foundation.Point) {
	C.C_NSLayoutManager_DrawGlyphsForGlyphRange_AtPoint(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(glyphsToShow)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(origin))))
}

func (n *NSLayoutManager) DrawStrikethroughForGlyphRange_StrikethroughType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, baselineOffset coregraphics.Float, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	C.C_NSLayoutManager_DrawStrikethroughForGlyphRange_StrikethroughType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)), C.int(int(strikethroughVal)), C.double(float64(baselineOffset)), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(lineRect))), *(*C.NSRange)(foundation.ToNSRangePointer(lineGlyphRange)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(containerOrigin))))
}

func (n *NSLayoutManager) DrawUnderlineForGlyphRange_UnderlineType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, baselineOffset coregraphics.Float, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	C.C_NSLayoutManager_DrawUnderlineForGlyphRange_UnderlineType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)), C.int(int(underlineVal)), C.double(float64(baselineOffset)), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(lineRect))), *(*C.NSRange)(foundation.ToNSRangePointer(lineGlyphRange)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(containerOrigin))))
}

func (n *NSLayoutManager) StrikethroughGlyphRange_StrikethroughType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	C.C_NSLayoutManager_StrikethroughGlyphRange_StrikethroughType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)), C.int(int(strikethroughVal)), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(lineRect))), *(*C.NSRange)(foundation.ToNSRangePointer(lineGlyphRange)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(containerOrigin))))
}

func (n *NSLayoutManager) UnderlineGlyphRange_UnderlineType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	C.C_NSLayoutManager_UnderlineGlyphRange_UnderlineType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)), C.int(int(underlineVal)), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(lineRect))), *(*C.NSRange)(foundation.ToNSRangePointer(lineGlyphRange)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(containerOrigin))))
}

func (n *NSLayoutManager) SetLayoutRect_ForTextBlock_GlyphRange(rect foundation.Rect, block TextBlock, glyphRange foundation.Range) {
	C.C_NSLayoutManager_SetLayoutRect_ForTextBlock_GlyphRange(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(block), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func (n *NSLayoutManager) LayoutRectForTextBlock_GlyphRange(block TextBlock, glyphRange foundation.Range) foundation.Rect {
	result_ := C.C_NSLayoutManager_LayoutRectForTextBlock_GlyphRange(n.Ptr(), objc.ExtractPtr(block), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSLayoutManager) SetBoundsRect_ForTextBlock_GlyphRange(rect foundation.Rect, block TextBlock, glyphRange foundation.Range) {
	C.C_NSLayoutManager_SetBoundsRect_ForTextBlock_GlyphRange(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(block), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
}

func (n *NSLayoutManager) BoundsRectForTextBlock_GlyphRange(block TextBlock, glyphRange foundation.Range) foundation.Rect {
	result_ := C.C_NSLayoutManager_BoundsRectForTextBlock_GlyphRange(n.Ptr(), objc.ExtractPtr(block), *(*C.NSRange)(foundation.ToNSRangePointer(glyphRange)))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSLayoutManager) ShowAttachmentCell_InRect_CharacterIndex(cell Cell, rect foundation.Rect, attachmentIndex uint) {
	C.C_NSLayoutManager_ShowAttachmentCell_InRect_CharacterIndex(n.Ptr(), objc.ExtractPtr(cell), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), C.uint(attachmentIndex))
}

func (n *NSLayoutManager) RulerAccessoryViewForTextView_ParagraphStyle_Ruler_Enabled(view TextView, style ParagraphStyle, ruler RulerView, isEnabled bool) View {
	result_ := C.C_NSLayoutManager_RulerAccessoryViewForTextView_ParagraphStyle_Ruler_Enabled(n.Ptr(), objc.ExtractPtr(view), objc.ExtractPtr(style), objc.ExtractPtr(ruler), C.bool(isEnabled))
	return MakeView(result_)
}

func (n *NSLayoutManager) RulerMarkersForTextView_ParagraphStyle_Ruler(view TextView, style ParagraphStyle, ruler RulerView) []RulerMarker {
	result_ := C.C_NSLayoutManager_RulerMarkersForTextView_ParagraphStyle_Ruler(n.Ptr(), objc.ExtractPtr(view), objc.ExtractPtr(style), objc.ExtractPtr(ruler))
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]RulerMarker, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeRulerMarker(r)
	}
	return goResult_
}

func (n *NSLayoutManager) LayoutManagerOwnsFirstResponderInWindow(window Window) bool {
	result_ := C.C_NSLayoutManager_LayoutManagerOwnsFirstResponderInWindow(n.Ptr(), objc.ExtractPtr(window))
	return bool(result_)
}

func (n *NSLayoutManager) DefaultLineHeightForFont(theFont Font) coregraphics.Float {
	result_ := C.C_NSLayoutManager_DefaultLineHeightForFont(n.Ptr(), objc.ExtractPtr(theFont))
	return coregraphics.Float(float64(result_))
}

func (n *NSLayoutManager) DefaultBaselineOffsetForFont(theFont Font) coregraphics.Float {
	result_ := C.C_NSLayoutManager_DefaultBaselineOffsetForFont(n.Ptr(), objc.ExtractPtr(theFont))
	return coregraphics.Float(float64(result_))
}

func (n *NSLayoutManager) AddTemporaryAttribute_Value_ForCharacterRange(attrName foundation.AttributedStringKey, value objc.Object, charRange foundation.Range) {
	C.C_NSLayoutManager_AddTemporaryAttribute_Value_ForCharacterRange(n.Ptr(), foundation.NewString(string(attrName)).Ptr(), objc.ExtractPtr(value), *(*C.NSRange)(foundation.ToNSRangePointer(charRange)))
}

func (n *NSLayoutManager) RemoveTemporaryAttribute_ForCharacterRange(attrName foundation.AttributedStringKey, charRange foundation.Range) {
	C.C_NSLayoutManager_RemoveTemporaryAttribute_ForCharacterRange(n.Ptr(), foundation.NewString(string(attrName)).Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(charRange)))
}

func (n *NSLayoutManager) Delegate() objc.Object {
	result_ := C.C_NSLayoutManager_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSLayoutManager) SetDelegate(value objc.Object) {
	C.C_NSLayoutManager_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSLayoutManager) TextStorage() TextStorage {
	result_ := C.C_NSLayoutManager_TextStorage(n.Ptr())
	return MakeTextStorage(result_)
}

func (n *NSLayoutManager) SetTextStorage(value TextStorage) {
	C.C_NSLayoutManager_SetTextStorage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSLayoutManager) AllowsNonContiguousLayout() bool {
	result_ := C.C_NSLayoutManager_AllowsNonContiguousLayout(n.Ptr())
	return bool(result_)
}

func (n *NSLayoutManager) SetAllowsNonContiguousLayout(value bool) {
	C.C_NSLayoutManager_SetAllowsNonContiguousLayout(n.Ptr(), C.bool(value))
}

func (n *NSLayoutManager) HasNonContiguousLayout() bool {
	result_ := C.C_NSLayoutManager_HasNonContiguousLayout(n.Ptr())
	return bool(result_)
}

func (n *NSLayoutManager) ShowsInvisibleCharacters() bool {
	result_ := C.C_NSLayoutManager_ShowsInvisibleCharacters(n.Ptr())
	return bool(result_)
}

func (n *NSLayoutManager) SetShowsInvisibleCharacters(value bool) {
	C.C_NSLayoutManager_SetShowsInvisibleCharacters(n.Ptr(), C.bool(value))
}

func (n *NSLayoutManager) ShowsControlCharacters() bool {
	result_ := C.C_NSLayoutManager_ShowsControlCharacters(n.Ptr())
	return bool(result_)
}

func (n *NSLayoutManager) SetShowsControlCharacters(value bool) {
	C.C_NSLayoutManager_SetShowsControlCharacters(n.Ptr(), C.bool(value))
}

func (n *NSLayoutManager) UsesFontLeading() bool {
	result_ := C.C_NSLayoutManager_UsesFontLeading(n.Ptr())
	return bool(result_)
}

func (n *NSLayoutManager) SetUsesFontLeading(value bool) {
	C.C_NSLayoutManager_SetUsesFontLeading(n.Ptr(), C.bool(value))
}

func (n *NSLayoutManager) BackgroundLayoutEnabled() bool {
	result_ := C.C_NSLayoutManager_BackgroundLayoutEnabled(n.Ptr())
	return bool(result_)
}

func (n *NSLayoutManager) SetBackgroundLayoutEnabled(value bool) {
	C.C_NSLayoutManager_SetBackgroundLayoutEnabled(n.Ptr(), C.bool(value))
}

func (n *NSLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	result_ := C.C_NSLayoutManager_LimitsLayoutForSuspiciousContents(n.Ptr())
	return bool(result_)
}

func (n *NSLayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	C.C_NSLayoutManager_SetLimitsLayoutForSuspiciousContents(n.Ptr(), C.bool(value))
}

func (n *NSLayoutManager) UsesDefaultHyphenation() bool {
	result_ := C.C_NSLayoutManager_UsesDefaultHyphenation(n.Ptr())
	return bool(result_)
}

func (n *NSLayoutManager) SetUsesDefaultHyphenation(value bool) {
	C.C_NSLayoutManager_SetUsesDefaultHyphenation(n.Ptr(), C.bool(value))
}

func (n *NSLayoutManager) TextContainers() []TextContainer {
	result_ := C.C_NSLayoutManager_TextContainers(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]TextContainer, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeTextContainer(r)
	}
	return goResult_
}

func (n *NSLayoutManager) GlyphGenerator() GlyphGenerator {
	result_ := C.C_NSLayoutManager_GlyphGenerator(n.Ptr())
	return MakeGlyphGenerator(result_)
}

func (n *NSLayoutManager) SetGlyphGenerator(value GlyphGenerator) {
	C.C_NSLayoutManager_SetGlyphGenerator(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSLayoutManager) NumberOfGlyphs() uint {
	result_ := C.C_NSLayoutManager_NumberOfGlyphs(n.Ptr())
	return uint(result_)
}

func (n *NSLayoutManager) ExtraLineFragmentRect() foundation.Rect {
	result_ := C.C_NSLayoutManager_ExtraLineFragmentRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSLayoutManager) ExtraLineFragmentTextContainer() TextContainer {
	result_ := C.C_NSLayoutManager_ExtraLineFragmentTextContainer(n.Ptr())
	return MakeTextContainer(result_)
}

func (n *NSLayoutManager) ExtraLineFragmentUsedRect() foundation.Rect {
	result_ := C.C_NSLayoutManager_ExtraLineFragmentUsedRect(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSLayoutManager) DefaultAttachmentScaling() ImageScaling {
	result_ := C.C_NSLayoutManager_DefaultAttachmentScaling(n.Ptr())
	return ImageScaling(uint(result_))
}

func (n *NSLayoutManager) SetDefaultAttachmentScaling(value ImageScaling) {
	C.C_NSLayoutManager_SetDefaultAttachmentScaling(n.Ptr(), C.uint(uint(value)))
}

func (n *NSLayoutManager) FirstTextView() TextView {
	result_ := C.C_NSLayoutManager_FirstTextView(n.Ptr())
	return MakeTextView(result_)
}

func (n *NSLayoutManager) TextViewForBeginningOfSelection() TextView {
	result_ := C.C_NSLayoutManager_TextViewForBeginningOfSelection(n.Ptr())
	return MakeTextView(result_)
}

func (n *NSLayoutManager) Typesetter() Typesetter {
	result_ := C.C_NSLayoutManager_Typesetter(n.Ptr())
	return MakeTypesetter(result_)
}

func (n *NSLayoutManager) SetTypesetter(value Typesetter) {
	C.C_NSLayoutManager_SetTypesetter(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSLayoutManager) TypesetterBehavior() TypesetterBehavior {
	result_ := C.C_NSLayoutManager_TypesetterBehavior(n.Ptr())
	return TypesetterBehavior(int(result_))
}

func (n *NSLayoutManager) SetTypesetterBehavior(value TypesetterBehavior) {
	C.C_NSLayoutManager_SetTypesetterBehavior(n.Ptr(), C.int(int(value)))
}
