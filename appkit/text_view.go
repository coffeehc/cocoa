package appkit

// #include "text_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type TextView interface {
	Text
	ReplaceTextContainer(newContainer TextContainer)
	InvalidateTextContainerOrigin()
	ChangeDocumentBackgroundColor(sender objc.Object)
	SetNeedsDisplayInRect_AvoidAdditionalLayout(rect foundation.Rect, flag bool)
	DrawInsertionPointInRect_Color_TurnedOn(rect foundation.Rect, color Color, flag bool)
	DrawViewBackgroundInRect(rect foundation.Rect)
	SetConstrainedFrameSize(desiredSize foundation.Size)
	CleanUpAfterDragOperation()
	ShowFindIndicatorForRange(charRange foundation.Range)
	SetBaseWritingDirection_Range(writingDirection WritingDirection, _range foundation.Range)
	Outline(sender objc.Object)
	ToggleAutomaticQuoteSubstitution(sender objc.Object)
	ToggleAutomaticLinkDetection(sender objc.Object)
	SetSelectedRanges_Affinity_StillSelecting(ranges []foundation.Value, affinity SelectionAffinity, stillSelectingFlag bool)
	UpdateInsertionPointStateAndRestartTimer(restartFlag bool)
	CharacterIndexForInsertionAtPoint(point foundation.Point) uint
	PreferredPasteboardTypeFromArray_RestrictedToTypesFromArray(availableTypes []PasteboardType, allowedTypes []PasteboardType) PasteboardType
	ReadSelectionFromPasteboard(pboard Pasteboard) bool
	ReadSelectionFromPasteboard_Type(pboard Pasteboard, _type PasteboardType) bool
	WriteSelectionToPasteboard_Type(pboard Pasteboard, _type PasteboardType) bool
	WriteSelectionToPasteboard_Types(pboard Pasteboard, types []PasteboardType) bool
	AlignJustified(sender objc.Object)
	ChangeAttributes(sender objc.Object)
	ChangeColor(sender objc.Object)
	SetAlignment_Range(alignment TextAlignment, _range foundation.Range)
	UseStandardKerning(sender objc.Object)
	LowerBaseline(sender objc.Object)
	RaiseBaseline(sender objc.Object)
	TurnOffKerning(sender objc.Object)
	LoosenKerning(sender objc.Object)
	TightenKerning(sender objc.Object)
	UseStandardLigatures(sender objc.Object)
	TurnOffLigatures(sender objc.Object)
	UseAllLigatures(sender objc.Object)
	ClickedOnLink_AtIndex(link objc.Object, charIndex uint)
	PasteAsPlainText(sender objc.Object)
	PasteAsRichText(sender objc.Object)
	BreakUndoCoalescing()
	UpdateFontPanel()
	UpdateRuler()
	UpdateDragTypeRegistration()
	SelectionRangeForProposedRange_Granularity(proposedCharRange foundation.Range, granularity SelectionGranularity) foundation.Range
	ShouldChangeTextInRange_ReplacementString(affectedCharRange foundation.Range, replacementString string) bool
	ShouldChangeTextInRanges_ReplacementStrings(affectedRanges []foundation.Value, replacementStrings []string) bool
	DidChangeText()
	SmartDeleteRangeForProposedRange(proposedCharRange foundation.Range) foundation.Range
	SmartInsertAfterStringForString_ReplacingRange(pasteString string, charRangeToReplace foundation.Range) string
	SmartInsertBeforeStringForString_ReplacingRange(pasteString string, charRangeToReplace foundation.Range) string
	ToggleSmartInsertDelete(sender objc.Object)
	ToggleContinuousSpellChecking(sender objc.Object)
	ToggleGrammarChecking(sender objc.Object)
	SetSpellingState_Range(value int, charRange foundation.Range)
	OrderFrontSharingServicePicker(sender objc.Object)
	DragOperationForDraggingInfo_Type(dragInfo objc.Object, _type PasteboardType) DragOperation
	DragSelectionWithEvent_Offset_SlideBack(event Event, mouseOffset foundation.Size, slideBack bool) bool
	StartSpeaking(sender objc.Object)
	StopSpeaking(sender objc.Object)
	PerformFindPanelAction(sender objc.Object)
	OrderFrontLinkPanel(sender objc.Object)
	OrderFrontListPanel(sender objc.Object)
	OrderFrontSpacingPanel(sender objc.Object)
	OrderFrontTablePanel(sender objc.Object)
	OrderFrontSubstitutionsPanel(sender objc.Object)
	Complete(sender objc.Object)
	InsertCompletion_ForPartialWordRange_Movement_IsFinal(word string, charRange foundation.Range, movement int, flag bool)
	CheckTextInDocument(sender objc.Object)
	CheckTextInSelection(sender objc.Object)
	ToggleAutomaticDashSubstitution(sender objc.Object)
	ToggleAutomaticDataDetection(sender objc.Object)
	ToggleAutomaticSpellingCorrection(sender objc.Object)
	ToggleAutomaticTextReplacement(sender objc.Object)
	UpdateQuickLookPreviewPanel()
	ToggleQuickLookPreviewPanel(sender objc.Object)
	ChangeLayoutOrientation(sender objc.Object)
	SetLayoutOrientation(orientation TextLayoutOrientation)
	PerformValidatedReplacementInRange_WithAttributedString(_range foundation.Range, attributedString foundation.AttributedString) bool
	ToggleAutomaticTextCompletion(sender objc.Object)
	UpdateCandidates()
	UpdateTextTouchBarItems()
	UpdateTouchBarItemIdentifiers()
	TextContainer() TextContainer
	SetTextContainer(value TextContainer)
	TextContainerInset() foundation.Size
	SetTextContainerInset(value foundation.Size)
	TextContainerOrigin() foundation.Point
	LayoutManager() LayoutManager
	TextStorage() TextStorage
	AllowsDocumentBackgroundColorChange() bool
	SetAllowsDocumentBackgroundColorChange(value bool)
	ShouldDrawInsertionPoint() bool
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	DefaultParagraphStyle() ParagraphStyle
	SetDefaultParagraphStyle(value ParagraphStyle)
	AllowsImageEditing() bool
	SetAllowsImageEditing(value bool)
	IsAutomaticQuoteSubstitutionEnabled() bool
	SetAutomaticQuoteSubstitutionEnabled(value bool)
	IsAutomaticLinkDetectionEnabled() bool
	SetAutomaticLinkDetectionEnabled(value bool)
	DisplaysLinkToolTips() bool
	SetDisplaysLinkToolTips(value bool)
	UsesRuler() bool
	SetUsesRuler(value bool)
	SetRulerVisible(value bool)
	UsesInspectorBar() bool
	SetUsesInspectorBar(value bool)
	SelectedRanges() []foundation.Value
	SetSelectedRanges(value []foundation.Value)
	SelectionAffinity() SelectionAffinity
	SelectionGranularity() SelectionGranularity
	SetSelectionGranularity(value SelectionGranularity)
	InsertionPointColor() Color
	SetInsertionPointColor(value Color)
	SelectedTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetSelectedTextAttributes(value map[foundation.AttributedStringKey]objc.Object)
	MarkedTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetMarkedTextAttributes(value map[foundation.AttributedStringKey]objc.Object)
	LinkTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetLinkTextAttributes(value map[foundation.AttributedStringKey]objc.Object)
	ReadablePasteboardTypes() []PasteboardType
	WritablePasteboardTypes() []PasteboardType
	TypingAttributes() map[foundation.AttributedStringKey]objc.Object
	SetTypingAttributes(value map[foundation.AttributedStringKey]objc.Object)
	IsCoalescingUndo() bool
	AcceptableDragTypes() []PasteboardType
	RangeForUserCharacterAttributeChange() foundation.Range
	RangesForUserCharacterAttributeChange() []foundation.Value
	RangeForUserParagraphAttributeChange() foundation.Range
	RangesForUserParagraphAttributeChange() []foundation.Value
	RangeForUserTextChange() foundation.Range
	RangesForUserTextChange() []foundation.Value
	SmartInsertDeleteEnabled() bool
	SetSmartInsertDeleteEnabled(value bool)
	IsContinuousSpellCheckingEnabled() bool
	SetContinuousSpellCheckingEnabled(value bool)
	SpellCheckerDocumentTag() int
	IsGrammarCheckingEnabled() bool
	SetGrammarCheckingEnabled(value bool)
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	UsesFindPanel() bool
	SetUsesFindPanel(value bool)
	RangeForUserCompletion() foundation.Range
	IsAutomaticDashSubstitutionEnabled() bool
	SetAutomaticDashSubstitutionEnabled(value bool)
	IsAutomaticDataDetectionEnabled() bool
	SetAutomaticDataDetectionEnabled(value bool)
	IsAutomaticSpellingCorrectionEnabled() bool
	SetAutomaticSpellingCorrectionEnabled(value bool)
	IsAutomaticTextReplacementEnabled() bool
	SetAutomaticTextReplacementEnabled(value bool)
	UsesFindBar() bool
	SetUsesFindBar(value bool)
	IsIncrementalSearchingEnabled() bool
	SetIncrementalSearchingEnabled(value bool)
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
	IsAutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)
	UsesAdaptiveColorMappingForDarkAppearance() bool
	SetUsesAdaptiveColorMappingForDarkAppearance(value bool)
	UsesRolloverButtonForSelection() bool
	SetUsesRolloverButtonForSelection(value bool)
}

type NSTextView struct {
	NSText
}

func MakeTextView(ptr unsafe.Pointer) NSTextView {
	return NSTextView{
		NSText: MakeText(ptr),
	}
}

func (n NSTextView) InitWithFrame_TextContainer(frameRect foundation.Rect, container TextContainer) NSTextView {
	result_ := C.C_NSTextView_InitWithFrame_TextContainer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))), objc.ExtractPtr(container))
	return MakeTextView(result_)
}

func (n NSTextView) InitWithFrame(frameRect foundation.Rect) NSTextView {
	result_ := C.C_NSTextView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeTextView(result_)
}

func (n NSTextView) InitWithCoder(coder foundation.Coder) NSTextView {
	result_ := C.C_NSTextView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTextView(result_)
}

func TextView_FieldEditor() NSTextView {
	result_ := C.C_NSTextView_TextView_FieldEditor()
	return MakeTextView(result_)
}

func (n NSTextView) Init() NSTextView {
	result_ := C.C_NSTextView_Init(n.Ptr())
	return MakeTextView(result_)
}

func AllocTextView() NSTextView {
	result_ := C.C_NSTextView_AllocTextView()
	return MakeTextView(result_)
}

func NewTextView() NSTextView {
	result_ := C.C_NSTextView_NewTextView()
	return MakeTextView(result_)
}

func (n NSTextView) Autorelease() NSTextView {
	result_ := C.C_NSTextView_Autorelease(n.Ptr())
	return MakeTextView(result_)
}

func (n NSTextView) Retain() NSTextView {
	result_ := C.C_NSTextView_Retain(n.Ptr())
	return MakeTextView(result_)
}

func TextView_RegisterForServices() {
	C.C_NSTextView_TextView_RegisterForServices()
}

func (n NSTextView) ReplaceTextContainer(newContainer TextContainer) {
	C.C_NSTextView_ReplaceTextContainer(n.Ptr(), objc.ExtractPtr(newContainer))
}

func (n NSTextView) InvalidateTextContainerOrigin() {
	C.C_NSTextView_InvalidateTextContainerOrigin(n.Ptr())
}

func (n NSTextView) ChangeDocumentBackgroundColor(sender objc.Object) {
	C.C_NSTextView_ChangeDocumentBackgroundColor(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) SetNeedsDisplayInRect_AvoidAdditionalLayout(rect foundation.Rect, flag bool) {
	C.C_NSTextView_SetNeedsDisplayInRect_AvoidAdditionalLayout(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), C.bool(flag))
}

func (n NSTextView) DrawInsertionPointInRect_Color_TurnedOn(rect foundation.Rect, color Color, flag bool) {
	C.C_NSTextView_DrawInsertionPointInRect_Color_TurnedOn(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(color), C.bool(flag))
}

func (n NSTextView) DrawViewBackgroundInRect(rect foundation.Rect) {
	C.C_NSTextView_DrawViewBackgroundInRect(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
}

func (n NSTextView) SetConstrainedFrameSize(desiredSize foundation.Size) {
	C.C_NSTextView_SetConstrainedFrameSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(desiredSize))))
}

func (n NSTextView) CleanUpAfterDragOperation() {
	C.C_NSTextView_CleanUpAfterDragOperation(n.Ptr())
}

func (n NSTextView) ShowFindIndicatorForRange(charRange foundation.Range) {
	C.C_NSTextView_ShowFindIndicatorForRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(charRange)))
}

func (n NSTextView) SetBaseWritingDirection_Range(writingDirection WritingDirection, _range foundation.Range) {
	C.C_NSTextView_SetBaseWritingDirection_Range(n.Ptr(), C.int(int(writingDirection)), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
}

func (n NSTextView) Outline(sender objc.Object) {
	C.C_NSTextView_Outline(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ToggleAutomaticQuoteSubstitution(sender objc.Object) {
	C.C_NSTextView_ToggleAutomaticQuoteSubstitution(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ToggleAutomaticLinkDetection(sender objc.Object) {
	C.C_NSTextView_ToggleAutomaticLinkDetection(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) SetSelectedRanges_Affinity_StillSelecting(ranges []foundation.Value, affinity SelectionAffinity, stillSelectingFlag bool) {
	var cRanges C.Array
	if len(ranges) > 0 {
		cRangesData := make([]unsafe.Pointer, len(ranges))
		for idx, v := range ranges {
			cRangesData[idx] = objc.ExtractPtr(v)
		}
		cRanges.data = unsafe.Pointer(&cRangesData[0])
		cRanges.len = C.int(len(ranges))
	}
	C.C_NSTextView_SetSelectedRanges_Affinity_StillSelecting(n.Ptr(), cRanges, C.uint(uint(affinity)), C.bool(stillSelectingFlag))
}

func (n NSTextView) UpdateInsertionPointStateAndRestartTimer(restartFlag bool) {
	C.C_NSTextView_UpdateInsertionPointStateAndRestartTimer(n.Ptr(), C.bool(restartFlag))
}

func (n NSTextView) CharacterIndexForInsertionAtPoint(point foundation.Point) uint {
	result_ := C.C_NSTextView_CharacterIndexForInsertionAtPoint(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(point))))
	return uint(result_)
}

func (n NSTextView) PreferredPasteboardTypeFromArray_RestrictedToTypesFromArray(availableTypes []PasteboardType, allowedTypes []PasteboardType) PasteboardType {
	var cAvailableTypes C.Array
	if len(availableTypes) > 0 {
		cAvailableTypesData := make([]unsafe.Pointer, len(availableTypes))
		for idx, v := range availableTypes {
			cAvailableTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cAvailableTypes.data = unsafe.Pointer(&cAvailableTypesData[0])
		cAvailableTypes.len = C.int(len(availableTypes))
	}
	var cAllowedTypes C.Array
	if len(allowedTypes) > 0 {
		cAllowedTypesData := make([]unsafe.Pointer, len(allowedTypes))
		for idx, v := range allowedTypes {
			cAllowedTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cAllowedTypes.data = unsafe.Pointer(&cAllowedTypesData[0])
		cAllowedTypes.len = C.int(len(allowedTypes))
	}
	result_ := C.C_NSTextView_PreferredPasteboardTypeFromArray_RestrictedToTypesFromArray(n.Ptr(), cAvailableTypes, cAllowedTypes)
	return PasteboardType(foundation.MakeString(result_).String())
}

func (n NSTextView) ReadSelectionFromPasteboard(pboard Pasteboard) bool {
	result_ := C.C_NSTextView_ReadSelectionFromPasteboard(n.Ptr(), objc.ExtractPtr(pboard))
	return bool(result_)
}

func (n NSTextView) ReadSelectionFromPasteboard_Type(pboard Pasteboard, _type PasteboardType) bool {
	result_ := C.C_NSTextView_ReadSelectionFromPasteboard_Type(n.Ptr(), objc.ExtractPtr(pboard), foundation.NewString(string(_type)).Ptr())
	return bool(result_)
}

func (n NSTextView) WriteSelectionToPasteboard_Type(pboard Pasteboard, _type PasteboardType) bool {
	result_ := C.C_NSTextView_WriteSelectionToPasteboard_Type(n.Ptr(), objc.ExtractPtr(pboard), foundation.NewString(string(_type)).Ptr())
	return bool(result_)
}

func (n NSTextView) WriteSelectionToPasteboard_Types(pboard Pasteboard, types []PasteboardType) bool {
	var cTypes C.Array
	if len(types) > 0 {
		cTypesData := make([]unsafe.Pointer, len(types))
		for idx, v := range types {
			cTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cTypes.data = unsafe.Pointer(&cTypesData[0])
		cTypes.len = C.int(len(types))
	}
	result_ := C.C_NSTextView_WriteSelectionToPasteboard_Types(n.Ptr(), objc.ExtractPtr(pboard), cTypes)
	return bool(result_)
}

func (n NSTextView) AlignJustified(sender objc.Object) {
	C.C_NSTextView_AlignJustified(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ChangeAttributes(sender objc.Object) {
	C.C_NSTextView_ChangeAttributes(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ChangeColor(sender objc.Object) {
	C.C_NSTextView_ChangeColor(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) SetAlignment_Range(alignment TextAlignment, _range foundation.Range) {
	C.C_NSTextView_SetAlignment_Range(n.Ptr(), C.int(int(alignment)), *(*C.NSRange)(foundation.ToNSRangePointer(_range)))
}

func (n NSTextView) UseStandardKerning(sender objc.Object) {
	C.C_NSTextView_UseStandardKerning(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) LowerBaseline(sender objc.Object) {
	C.C_NSTextView_LowerBaseline(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) RaiseBaseline(sender objc.Object) {
	C.C_NSTextView_RaiseBaseline(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) TurnOffKerning(sender objc.Object) {
	C.C_NSTextView_TurnOffKerning(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) LoosenKerning(sender objc.Object) {
	C.C_NSTextView_LoosenKerning(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) TightenKerning(sender objc.Object) {
	C.C_NSTextView_TightenKerning(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) UseStandardLigatures(sender objc.Object) {
	C.C_NSTextView_UseStandardLigatures(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) TurnOffLigatures(sender objc.Object) {
	C.C_NSTextView_TurnOffLigatures(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) UseAllLigatures(sender objc.Object) {
	C.C_NSTextView_UseAllLigatures(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ClickedOnLink_AtIndex(link objc.Object, charIndex uint) {
	C.C_NSTextView_ClickedOnLink_AtIndex(n.Ptr(), objc.ExtractPtr(link), C.uint(charIndex))
}

func (n NSTextView) PasteAsPlainText(sender objc.Object) {
	C.C_NSTextView_PasteAsPlainText(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) PasteAsRichText(sender objc.Object) {
	C.C_NSTextView_PasteAsRichText(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) BreakUndoCoalescing() {
	C.C_NSTextView_BreakUndoCoalescing(n.Ptr())
}

func (n NSTextView) UpdateFontPanel() {
	C.C_NSTextView_UpdateFontPanel(n.Ptr())
}

func (n NSTextView) UpdateRuler() {
	C.C_NSTextView_UpdateRuler(n.Ptr())
}

func (n NSTextView) UpdateDragTypeRegistration() {
	C.C_NSTextView_UpdateDragTypeRegistration(n.Ptr())
}

func (n NSTextView) SelectionRangeForProposedRange_Granularity(proposedCharRange foundation.Range, granularity SelectionGranularity) foundation.Range {
	result_ := C.C_NSTextView_SelectionRangeForProposedRange_Granularity(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(proposedCharRange)), C.uint(uint(granularity)))
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSTextView) ShouldChangeTextInRange_ReplacementString(affectedCharRange foundation.Range, replacementString string) bool {
	result_ := C.C_NSTextView_ShouldChangeTextInRange_ReplacementString(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(affectedCharRange)), foundation.NewString(replacementString).Ptr())
	return bool(result_)
}

func (n NSTextView) ShouldChangeTextInRanges_ReplacementStrings(affectedRanges []foundation.Value, replacementStrings []string) bool {
	var cAffectedRanges C.Array
	if len(affectedRanges) > 0 {
		cAffectedRangesData := make([]unsafe.Pointer, len(affectedRanges))
		for idx, v := range affectedRanges {
			cAffectedRangesData[idx] = objc.ExtractPtr(v)
		}
		cAffectedRanges.data = unsafe.Pointer(&cAffectedRangesData[0])
		cAffectedRanges.len = C.int(len(affectedRanges))
	}
	var cReplacementStrings C.Array
	if len(replacementStrings) > 0 {
		cReplacementStringsData := make([]unsafe.Pointer, len(replacementStrings))
		for idx, v := range replacementStrings {
			cReplacementStringsData[idx] = foundation.NewString(v).Ptr()
		}
		cReplacementStrings.data = unsafe.Pointer(&cReplacementStringsData[0])
		cReplacementStrings.len = C.int(len(replacementStrings))
	}
	result_ := C.C_NSTextView_ShouldChangeTextInRanges_ReplacementStrings(n.Ptr(), cAffectedRanges, cReplacementStrings)
	return bool(result_)
}

func (n NSTextView) DidChangeText() {
	C.C_NSTextView_DidChangeText(n.Ptr())
}

func (n NSTextView) SmartDeleteRangeForProposedRange(proposedCharRange foundation.Range) foundation.Range {
	result_ := C.C_NSTextView_SmartDeleteRangeForProposedRange(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(proposedCharRange)))
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSTextView) SmartInsertAfterStringForString_ReplacingRange(pasteString string, charRangeToReplace foundation.Range) string {
	result_ := C.C_NSTextView_SmartInsertAfterStringForString_ReplacingRange(n.Ptr(), foundation.NewString(pasteString).Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(charRangeToReplace)))
	return foundation.MakeString(result_).String()
}

func (n NSTextView) SmartInsertBeforeStringForString_ReplacingRange(pasteString string, charRangeToReplace foundation.Range) string {
	result_ := C.C_NSTextView_SmartInsertBeforeStringForString_ReplacingRange(n.Ptr(), foundation.NewString(pasteString).Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(charRangeToReplace)))
	return foundation.MakeString(result_).String()
}

func (n NSTextView) ToggleSmartInsertDelete(sender objc.Object) {
	C.C_NSTextView_ToggleSmartInsertDelete(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ToggleContinuousSpellChecking(sender objc.Object) {
	C.C_NSTextView_ToggleContinuousSpellChecking(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ToggleGrammarChecking(sender objc.Object) {
	C.C_NSTextView_ToggleGrammarChecking(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) SetSpellingState_Range(value int, charRange foundation.Range) {
	C.C_NSTextView_SetSpellingState_Range(n.Ptr(), C.int(value), *(*C.NSRange)(foundation.ToNSRangePointer(charRange)))
}

func (n NSTextView) OrderFrontSharingServicePicker(sender objc.Object) {
	C.C_NSTextView_OrderFrontSharingServicePicker(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) DragOperationForDraggingInfo_Type(dragInfo objc.Object, _type PasteboardType) DragOperation {
	result_ := C.C_NSTextView_DragOperationForDraggingInfo_Type(n.Ptr(), objc.ExtractPtr(dragInfo), foundation.NewString(string(_type)).Ptr())
	return DragOperation(uint(result_))
}

func (n NSTextView) DragSelectionWithEvent_Offset_SlideBack(event Event, mouseOffset foundation.Size, slideBack bool) bool {
	result_ := C.C_NSTextView_DragSelectionWithEvent_Offset_SlideBack(n.Ptr(), objc.ExtractPtr(event), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(mouseOffset))), C.bool(slideBack))
	return bool(result_)
}

func (n NSTextView) StartSpeaking(sender objc.Object) {
	C.C_NSTextView_StartSpeaking(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) StopSpeaking(sender objc.Object) {
	C.C_NSTextView_StopSpeaking(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) PerformFindPanelAction(sender objc.Object) {
	C.C_NSTextView_PerformFindPanelAction(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) OrderFrontLinkPanel(sender objc.Object) {
	C.C_NSTextView_OrderFrontLinkPanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) OrderFrontListPanel(sender objc.Object) {
	C.C_NSTextView_OrderFrontListPanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) OrderFrontSpacingPanel(sender objc.Object) {
	C.C_NSTextView_OrderFrontSpacingPanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) OrderFrontTablePanel(sender objc.Object) {
	C.C_NSTextView_OrderFrontTablePanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) OrderFrontSubstitutionsPanel(sender objc.Object) {
	C.C_NSTextView_OrderFrontSubstitutionsPanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) Complete(sender objc.Object) {
	C.C_NSTextView_Complete(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) InsertCompletion_ForPartialWordRange_Movement_IsFinal(word string, charRange foundation.Range, movement int, flag bool) {
	C.C_NSTextView_InsertCompletion_ForPartialWordRange_Movement_IsFinal(n.Ptr(), foundation.NewString(word).Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(charRange)), C.int(movement), C.bool(flag))
}

func (n NSTextView) CheckTextInDocument(sender objc.Object) {
	C.C_NSTextView_CheckTextInDocument(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) CheckTextInSelection(sender objc.Object) {
	C.C_NSTextView_CheckTextInSelection(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ToggleAutomaticDashSubstitution(sender objc.Object) {
	C.C_NSTextView_ToggleAutomaticDashSubstitution(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ToggleAutomaticDataDetection(sender objc.Object) {
	C.C_NSTextView_ToggleAutomaticDataDetection(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ToggleAutomaticSpellingCorrection(sender objc.Object) {
	C.C_NSTextView_ToggleAutomaticSpellingCorrection(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ToggleAutomaticTextReplacement(sender objc.Object) {
	C.C_NSTextView_ToggleAutomaticTextReplacement(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) UpdateQuickLookPreviewPanel() {
	C.C_NSTextView_UpdateQuickLookPreviewPanel(n.Ptr())
}

func (n NSTextView) ToggleQuickLookPreviewPanel(sender objc.Object) {
	C.C_NSTextView_ToggleQuickLookPreviewPanel(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) ChangeLayoutOrientation(sender objc.Object) {
	C.C_NSTextView_ChangeLayoutOrientation(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) SetLayoutOrientation(orientation TextLayoutOrientation) {
	C.C_NSTextView_SetLayoutOrientation(n.Ptr(), C.int(int(orientation)))
}

func (n NSTextView) PerformValidatedReplacementInRange_WithAttributedString(_range foundation.Range, attributedString foundation.AttributedString) bool {
	result_ := C.C_NSTextView_PerformValidatedReplacementInRange_WithAttributedString(n.Ptr(), *(*C.NSRange)(foundation.ToNSRangePointer(_range)), objc.ExtractPtr(attributedString))
	return bool(result_)
}

func (n NSTextView) ToggleAutomaticTextCompletion(sender objc.Object) {
	C.C_NSTextView_ToggleAutomaticTextCompletion(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSTextView) UpdateCandidates() {
	C.C_NSTextView_UpdateCandidates(n.Ptr())
}

func (n NSTextView) UpdateTextTouchBarItems() {
	C.C_NSTextView_UpdateTextTouchBarItems(n.Ptr())
}

func (n NSTextView) UpdateTouchBarItemIdentifiers() {
	C.C_NSTextView_UpdateTouchBarItemIdentifiers(n.Ptr())
}

func ScrollableDocumentContentTextView() ScrollView {
	result_ := C.C_NSTextView_ScrollableDocumentContentTextView()
	return MakeScrollView(result_)
}

func ScrollablePlainDocumentContentTextView() ScrollView {
	result_ := C.C_NSTextView_ScrollablePlainDocumentContentTextView()
	return MakeScrollView(result_)
}

func ScrollableTextView() ScrollView {
	result_ := C.C_NSTextView_ScrollableTextView()
	return MakeScrollView(result_)
}

func (n NSTextView) TextContainer() TextContainer {
	result_ := C.C_NSTextView_TextContainer(n.Ptr())
	return MakeTextContainer(result_)
}

func (n NSTextView) SetTextContainer(value TextContainer) {
	C.C_NSTextView_SetTextContainer(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextView) TextContainerInset() foundation.Size {
	result_ := C.C_NSTextView_TextContainerInset(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSTextView) SetTextContainerInset(value foundation.Size) {
	C.C_NSTextView_SetTextContainerInset(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSTextView) TextContainerOrigin() foundation.Point {
	result_ := C.C_NSTextView_TextContainerOrigin(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSTextView) LayoutManager() LayoutManager {
	result_ := C.C_NSTextView_LayoutManager(n.Ptr())
	return MakeLayoutManager(result_)
}

func (n NSTextView) TextStorage() TextStorage {
	result_ := C.C_NSTextView_TextStorage(n.Ptr())
	return MakeTextStorage(result_)
}

func (n NSTextView) AllowsDocumentBackgroundColorChange() bool {
	result_ := C.C_NSTextView_AllowsDocumentBackgroundColorChange(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAllowsDocumentBackgroundColorChange(value bool) {
	C.C_NSTextView_SetAllowsDocumentBackgroundColorChange(n.Ptr(), C.bool(value))
}

func (n NSTextView) ShouldDrawInsertionPoint() bool {
	result_ := C.C_NSTextView_ShouldDrawInsertionPoint(n.Ptr())
	return bool(result_)
}

func (n NSTextView) AllowedInputSourceLocales() []string {
	result_ := C.C_NSTextView_AllowedInputSourceLocales(n.Ptr())
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

func (n NSTextView) SetAllowedInputSourceLocales(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = foundation.NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTextView_SetAllowedInputSourceLocales(n.Ptr(), cValue)
}

func (n NSTextView) AllowsUndo() bool {
	result_ := C.C_NSTextView_AllowsUndo(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAllowsUndo(value bool) {
	C.C_NSTextView_SetAllowsUndo(n.Ptr(), C.bool(value))
}

func (n NSTextView) DefaultParagraphStyle() ParagraphStyle {
	result_ := C.C_NSTextView_DefaultParagraphStyle(n.Ptr())
	return MakeParagraphStyle(result_)
}

func (n NSTextView) SetDefaultParagraphStyle(value ParagraphStyle) {
	C.C_NSTextView_SetDefaultParagraphStyle(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextView) AllowsImageEditing() bool {
	result_ := C.C_NSTextView_AllowsImageEditing(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAllowsImageEditing(value bool) {
	C.C_NSTextView_SetAllowsImageEditing(n.Ptr(), C.bool(value))
}

func (n NSTextView) IsAutomaticQuoteSubstitutionEnabled() bool {
	result_ := C.C_NSTextView_IsAutomaticQuoteSubstitutionEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAutomaticQuoteSubstitutionEnabled(value bool) {
	C.C_NSTextView_SetAutomaticQuoteSubstitutionEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextView) IsAutomaticLinkDetectionEnabled() bool {
	result_ := C.C_NSTextView_IsAutomaticLinkDetectionEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAutomaticLinkDetectionEnabled(value bool) {
	C.C_NSTextView_SetAutomaticLinkDetectionEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextView) DisplaysLinkToolTips() bool {
	result_ := C.C_NSTextView_DisplaysLinkToolTips(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetDisplaysLinkToolTips(value bool) {
	C.C_NSTextView_SetDisplaysLinkToolTips(n.Ptr(), C.bool(value))
}

func (n NSTextView) UsesRuler() bool {
	result_ := C.C_NSTextView_UsesRuler(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetUsesRuler(value bool) {
	C.C_NSTextView_SetUsesRuler(n.Ptr(), C.bool(value))
}

func (n NSTextView) SetRulerVisible(value bool) {
	C.C_NSTextView_SetRulerVisible(n.Ptr(), C.bool(value))
}

func (n NSTextView) UsesInspectorBar() bool {
	result_ := C.C_NSTextView_UsesInspectorBar(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetUsesInspectorBar(value bool) {
	C.C_NSTextView_SetUsesInspectorBar(n.Ptr(), C.bool(value))
}

func (n NSTextView) SelectedRanges() []foundation.Value {
	result_ := C.C_NSTextView_SelectedRanges(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]foundation.Value, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeValue(r)
	}
	return goResult_
}

func (n NSTextView) SetSelectedRanges(value []foundation.Value) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTextView_SetSelectedRanges(n.Ptr(), cValue)
}

func (n NSTextView) SelectionAffinity() SelectionAffinity {
	result_ := C.C_NSTextView_SelectionAffinity(n.Ptr())
	return SelectionAffinity(uint(result_))
}

func (n NSTextView) SelectionGranularity() SelectionGranularity {
	result_ := C.C_NSTextView_SelectionGranularity(n.Ptr())
	return SelectionGranularity(uint(result_))
}

func (n NSTextView) SetSelectionGranularity(value SelectionGranularity) {
	C.C_NSTextView_SetSelectionGranularity(n.Ptr(), C.uint(uint(value)))
}

func (n NSTextView) InsertionPointColor() Color {
	result_ := C.C_NSTextView_InsertionPointColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSTextView) SetInsertionPointColor(value Color) {
	C.C_NSTextView_SetInsertionPointColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTextView) SelectedTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	result_ := C.C_NSTextView_SelectedTextAttributes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[foundation.AttributedStringKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[foundation.AttributedStringKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSTextView) SetSelectedTextAttributes(value map[foundation.AttributedStringKey]objc.Object) {
	var cValue C.Dictionary
	if len(value) > 0 {
		cValueKeyData := make([]unsafe.Pointer, len(value))
		cValueValueData := make([]unsafe.Pointer, len(value))
		var idx = 0
		for k, v := range value {
			cValueKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cValueValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cValue.key_data = unsafe.Pointer(&cValueKeyData[0])
		cValue.value_data = unsafe.Pointer(&cValueValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTextView_SetSelectedTextAttributes(n.Ptr(), cValue)
}

func (n NSTextView) MarkedTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	result_ := C.C_NSTextView_MarkedTextAttributes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[foundation.AttributedStringKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[foundation.AttributedStringKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSTextView) SetMarkedTextAttributes(value map[foundation.AttributedStringKey]objc.Object) {
	var cValue C.Dictionary
	if len(value) > 0 {
		cValueKeyData := make([]unsafe.Pointer, len(value))
		cValueValueData := make([]unsafe.Pointer, len(value))
		var idx = 0
		for k, v := range value {
			cValueKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cValueValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cValue.key_data = unsafe.Pointer(&cValueKeyData[0])
		cValue.value_data = unsafe.Pointer(&cValueValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTextView_SetMarkedTextAttributes(n.Ptr(), cValue)
}

func (n NSTextView) LinkTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	result_ := C.C_NSTextView_LinkTextAttributes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[foundation.AttributedStringKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[foundation.AttributedStringKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSTextView) SetLinkTextAttributes(value map[foundation.AttributedStringKey]objc.Object) {
	var cValue C.Dictionary
	if len(value) > 0 {
		cValueKeyData := make([]unsafe.Pointer, len(value))
		cValueValueData := make([]unsafe.Pointer, len(value))
		var idx = 0
		for k, v := range value {
			cValueKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cValueValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cValue.key_data = unsafe.Pointer(&cValueKeyData[0])
		cValue.value_data = unsafe.Pointer(&cValueValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTextView_SetLinkTextAttributes(n.Ptr(), cValue)
}

func (n NSTextView) ReadablePasteboardTypes() []PasteboardType {
	result_ := C.C_NSTextView_ReadablePasteboardTypes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]PasteboardType, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = PasteboardType(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n NSTextView) WritablePasteboardTypes() []PasteboardType {
	result_ := C.C_NSTextView_WritablePasteboardTypes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]PasteboardType, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = PasteboardType(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n NSTextView) TypingAttributes() map[foundation.AttributedStringKey]objc.Object {
	result_ := C.C_NSTextView_TypingAttributes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[foundation.AttributedStringKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[foundation.AttributedStringKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSTextView) SetTypingAttributes(value map[foundation.AttributedStringKey]objc.Object) {
	var cValue C.Dictionary
	if len(value) > 0 {
		cValueKeyData := make([]unsafe.Pointer, len(value))
		cValueValueData := make([]unsafe.Pointer, len(value))
		var idx = 0
		for k, v := range value {
			cValueKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cValueValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cValue.key_data = unsafe.Pointer(&cValueKeyData[0])
		cValue.value_data = unsafe.Pointer(&cValueValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSTextView_SetTypingAttributes(n.Ptr(), cValue)
}

func (n NSTextView) IsCoalescingUndo() bool {
	result_ := C.C_NSTextView_IsCoalescingUndo(n.Ptr())
	return bool(result_)
}

func (n NSTextView) AcceptableDragTypes() []PasteboardType {
	result_ := C.C_NSTextView_AcceptableDragTypes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]PasteboardType, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = PasteboardType(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n NSTextView) RangeForUserCharacterAttributeChange() foundation.Range {
	result_ := C.C_NSTextView_RangeForUserCharacterAttributeChange(n.Ptr())
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSTextView) RangesForUserCharacterAttributeChange() []foundation.Value {
	result_ := C.C_NSTextView_RangesForUserCharacterAttributeChange(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]foundation.Value, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeValue(r)
	}
	return goResult_
}

func (n NSTextView) RangeForUserParagraphAttributeChange() foundation.Range {
	result_ := C.C_NSTextView_RangeForUserParagraphAttributeChange(n.Ptr())
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSTextView) RangesForUserParagraphAttributeChange() []foundation.Value {
	result_ := C.C_NSTextView_RangesForUserParagraphAttributeChange(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]foundation.Value, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeValue(r)
	}
	return goResult_
}

func (n NSTextView) RangeForUserTextChange() foundation.Range {
	result_ := C.C_NSTextView_RangeForUserTextChange(n.Ptr())
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSTextView) RangesForUserTextChange() []foundation.Value {
	result_ := C.C_NSTextView_RangesForUserTextChange(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]foundation.Value, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeValue(r)
	}
	return goResult_
}

func (n NSTextView) SmartInsertDeleteEnabled() bool {
	result_ := C.C_NSTextView_SmartInsertDeleteEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetSmartInsertDeleteEnabled(value bool) {
	C.C_NSTextView_SetSmartInsertDeleteEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextView) IsContinuousSpellCheckingEnabled() bool {
	result_ := C.C_NSTextView_IsContinuousSpellCheckingEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetContinuousSpellCheckingEnabled(value bool) {
	C.C_NSTextView_SetContinuousSpellCheckingEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextView) SpellCheckerDocumentTag() int {
	result_ := C.C_NSTextView_SpellCheckerDocumentTag(n.Ptr())
	return int(result_)
}

func (n NSTextView) IsGrammarCheckingEnabled() bool {
	result_ := C.C_NSTextView_IsGrammarCheckingEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetGrammarCheckingEnabled(value bool) {
	C.C_NSTextView_SetGrammarCheckingEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextView) AcceptsGlyphInfo() bool {
	result_ := C.C_NSTextView_AcceptsGlyphInfo(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAcceptsGlyphInfo(value bool) {
	C.C_NSTextView_SetAcceptsGlyphInfo(n.Ptr(), C.bool(value))
}

func (n NSTextView) UsesFindPanel() bool {
	result_ := C.C_NSTextView_UsesFindPanel(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetUsesFindPanel(value bool) {
	C.C_NSTextView_SetUsesFindPanel(n.Ptr(), C.bool(value))
}

func (n NSTextView) RangeForUserCompletion() foundation.Range {
	result_ := C.C_NSTextView_RangeForUserCompletion(n.Ptr())
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSTextView) IsAutomaticDashSubstitutionEnabled() bool {
	result_ := C.C_NSTextView_IsAutomaticDashSubstitutionEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAutomaticDashSubstitutionEnabled(value bool) {
	C.C_NSTextView_SetAutomaticDashSubstitutionEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextView) IsAutomaticDataDetectionEnabled() bool {
	result_ := C.C_NSTextView_IsAutomaticDataDetectionEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAutomaticDataDetectionEnabled(value bool) {
	C.C_NSTextView_SetAutomaticDataDetectionEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextView) IsAutomaticSpellingCorrectionEnabled() bool {
	result_ := C.C_NSTextView_IsAutomaticSpellingCorrectionEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAutomaticSpellingCorrectionEnabled(value bool) {
	C.C_NSTextView_SetAutomaticSpellingCorrectionEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextView) IsAutomaticTextReplacementEnabled() bool {
	result_ := C.C_NSTextView_IsAutomaticTextReplacementEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAutomaticTextReplacementEnabled(value bool) {
	C.C_NSTextView_SetAutomaticTextReplacementEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextView) UsesFindBar() bool {
	result_ := C.C_NSTextView_UsesFindBar(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetUsesFindBar(value bool) {
	C.C_NSTextView_SetUsesFindBar(n.Ptr(), C.bool(value))
}

func (n NSTextView) IsIncrementalSearchingEnabled() bool {
	result_ := C.C_NSTextView_IsIncrementalSearchingEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetIncrementalSearchingEnabled(value bool) {
	C.C_NSTextView_SetIncrementalSearchingEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextView) AllowsCharacterPickerTouchBarItem() bool {
	result_ := C.C_NSTextView_AllowsCharacterPickerTouchBarItem(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAllowsCharacterPickerTouchBarItem(value bool) {
	C.C_NSTextView_SetAllowsCharacterPickerTouchBarItem(n.Ptr(), C.bool(value))
}

func (n NSTextView) IsAutomaticTextCompletionEnabled() bool {
	result_ := C.C_NSTextView_IsAutomaticTextCompletionEnabled(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetAutomaticTextCompletionEnabled(value bool) {
	C.C_NSTextView_SetAutomaticTextCompletionEnabled(n.Ptr(), C.bool(value))
}

func (n NSTextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	result_ := C.C_NSTextView_UsesAdaptiveColorMappingForDarkAppearance(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetUsesAdaptiveColorMappingForDarkAppearance(value bool) {
	C.C_NSTextView_SetUsesAdaptiveColorMappingForDarkAppearance(n.Ptr(), C.bool(value))
}

func (n NSTextView) UsesRolloverButtonForSelection() bool {
	result_ := C.C_NSTextView_UsesRolloverButtonForSelection(n.Ptr())
	return bool(result_)
}

func (n NSTextView) SetUsesRolloverButtonForSelection(value bool) {
	C.C_NSTextView_SetUsesRolloverButtonForSelection(n.Ptr(), C.bool(value))
}

func TextView_StronglyReferencesTextStorage() bool {
	result_ := C.C_NSTextView_TextView_StronglyReferencesTextStorage()
	return bool(result_)
}

type TextViewDelegate struct {
	UndoManagerForTextView                                            func(view TextView) foundation.UndoManager
	TextView_WillDisplayToolTip_ForCharacterAtIndex                   func(textView TextView, tooltip string, characterIndex uint) string
	TextView_URLForContentsOfTextAttachment_AtIndex                   func(textView TextView, textAttachment TextAttachment, charIndex uint) foundation.URL
	TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange   func(textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range
	TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges func(textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.Value
	TextViewDidChangeSelection                                        func(notification foundation.Notification)
	TextView_WritablePasteboardTypesForCell_AtIndex                   func(view TextView, cell objc.Object, charIndex uint) []PasteboardType
	TextView_WriteCell_AtIndex_ToPasteboard_Type                      func(view TextView, cell objc.Object, charIndex uint, pboard Pasteboard, _type PasteboardType) bool
	TextView_ShouldChangeTextInRange_ReplacementString                func(textView TextView, affectedCharRange foundation.Range, replacementString string) bool
	TextView_ShouldChangeTextInRanges_ReplacementStrings              func(textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool
	TextView_ShouldChangeTypingAttributes_ToAttributes                func(textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.Object
	TextViewDidChangeTypingAttributes                                 func(notification foundation.Notification)
	TextView_ClickedOnCell_InRect_AtIndex                             func(textView TextView, cell objc.Object, cellFrame foundation.Rect, charIndex uint)
	TextView_DoubleClickedOnCell_InRect_AtIndex                       func(textView TextView, cell objc.Object, cellFrame foundation.Rect, charIndex uint)
	TextView_ClickedOnLink_AtIndex                                    func(textView TextView, link objc.Object, charIndex uint) bool
	TextView_ShouldSetSpellingState_Range                             func(textView TextView, value int, affectedCharRange foundation.Range) int
	TextView_DraggedCell_InRect_Event_AtIndex                         func(view TextView, cell objc.Object, rect foundation.Rect, event Event, charIndex uint)
	TextView_WillShowSharingServicePicker_ForItems                    func(textView TextView, servicePicker SharingServicePicker, items []objc.Object) SharingServicePicker
	TextView_DoCommandBySelector                                      func(textView TextView, commandSelector objc.Selector) bool
	TextView_Menu_ForEvent_AtIndex                                    func(view TextView, menu Menu, event Event, charIndex uint) Menu
	TextView_Candidates_ForSelectedRange                              func(textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult
	TextView_CandidatesForSelectedRange                               func(textView TextView, selectedRange foundation.Range) []objc.Object
	TextView_ShouldSelectCandidateAtIndex                             func(textView TextView, index uint) bool
	TextView_ShouldUpdateTouchBarItemIdentifiers                      func(textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier
	TextDidChange                                                     func(notification foundation.Notification)
	TextShouldBeginEditing                                            func(textObject Text) bool
	TextDidBeginEditing                                               func(notification foundation.Notification)
	TextShouldEndEditing                                              func(textObject Text) bool
	TextDidEndEditing                                                 func(notification foundation.Notification)
}

func (delegate *TextViewDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapTextViewDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export textViewDelegate_UndoManagerForTextView
func textViewDelegate_UndoManagerForTextView(hp C.uintptr_t, view unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.UndoManagerForTextView(MakeTextView(view))
	return objc.ExtractPtr(result)
}

//export textViewDelegate_TextView_WillDisplayToolTip_ForCharacterAtIndex
func textViewDelegate_TextView_WillDisplayToolTip_ForCharacterAtIndex(hp C.uintptr_t, textView unsafe.Pointer, tooltip unsafe.Pointer, characterIndex C.uint) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_WillDisplayToolTip_ForCharacterAtIndex(MakeTextView(textView), foundation.MakeString(tooltip).String(), uint(characterIndex))
	return foundation.NewString(result).Ptr()
}

//export textViewDelegate_TextView_URLForContentsOfTextAttachment_AtIndex
func textViewDelegate_TextView_URLForContentsOfTextAttachment_AtIndex(hp C.uintptr_t, textView unsafe.Pointer, textAttachment unsafe.Pointer, charIndex C.uint) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_URLForContentsOfTextAttachment_AtIndex(MakeTextView(textView), MakeTextAttachment(textAttachment), uint(charIndex))
	return objc.ExtractPtr(result)
}

//export textViewDelegate_TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange
func textViewDelegate_TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(hp C.uintptr_t, textView unsafe.Pointer, oldSelectedCharRange C.NSRange, newSelectedCharRange C.NSRange) C.NSRange {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange(MakeTextView(textView), foundation.FromNSRangePointer(unsafe.Pointer(&oldSelectedCharRange)), foundation.FromNSRangePointer(unsafe.Pointer(&newSelectedCharRange)))
	return *(*C.NSRange)(foundation.ToNSRangePointer(result))
}

//export textViewDelegate_TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges
func textViewDelegate_TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(hp C.uintptr_t, textView unsafe.Pointer, oldSelectedCharRanges C.Array, newSelectedCharRanges C.Array) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if oldSelectedCharRanges.len > 0 {
		defer C.free(oldSelectedCharRanges.data)
	}
	oldSelectedCharRangesSlice := unsafe.Slice((*unsafe.Pointer)(oldSelectedCharRanges.data), int(oldSelectedCharRanges.len))
	var goOldSelectedCharRanges = make([]foundation.Value, len(oldSelectedCharRangesSlice))
	for idx, r := range oldSelectedCharRangesSlice {
		goOldSelectedCharRanges[idx] = foundation.MakeValue(r)
	}
	if newSelectedCharRanges.len > 0 {
		defer C.free(newSelectedCharRanges.data)
	}
	newSelectedCharRangesSlice := unsafe.Slice((*unsafe.Pointer)(newSelectedCharRanges.data), int(newSelectedCharRanges.len))
	var goNewSelectedCharRanges = make([]foundation.Value, len(newSelectedCharRangesSlice))
	for idx, r := range newSelectedCharRangesSlice {
		goNewSelectedCharRanges[idx] = foundation.MakeValue(r)
	}
	result := delegate.TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges(MakeTextView(textView), goOldSelectedCharRanges, goNewSelectedCharRanges)
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = objc.ExtractPtr(v)
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textViewDelegate_TextViewDidChangeSelection
func textViewDelegate_TextViewDidChangeSelection(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextViewDidChangeSelection(foundation.MakeNotification(notification))
}

//export textViewDelegate_TextView_WritablePasteboardTypesForCell_AtIndex
func textViewDelegate_TextView_WritablePasteboardTypesForCell_AtIndex(hp C.uintptr_t, view unsafe.Pointer, cell unsafe.Pointer, charIndex C.uint) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_WritablePasteboardTypesForCell_AtIndex(MakeTextView(view), objc.MakeObject(cell), uint(charIndex))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textViewDelegate_TextView_WriteCell_AtIndex_ToPasteboard_Type
func textViewDelegate_TextView_WriteCell_AtIndex_ToPasteboard_Type(hp C.uintptr_t, view unsafe.Pointer, cell unsafe.Pointer, charIndex C.uint, pboard unsafe.Pointer, _type unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_WriteCell_AtIndex_ToPasteboard_Type(MakeTextView(view), objc.MakeObject(cell), uint(charIndex), MakePasteboard(pboard), PasteboardType(foundation.MakeString(_type).String()))
	return C.bool(result)
}

//export textViewDelegate_TextView_ShouldChangeTextInRange_ReplacementString
func textViewDelegate_TextView_ShouldChangeTextInRange_ReplacementString(hp C.uintptr_t, textView unsafe.Pointer, affectedCharRange C.NSRange, replacementString unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_ShouldChangeTextInRange_ReplacementString(MakeTextView(textView), foundation.FromNSRangePointer(unsafe.Pointer(&affectedCharRange)), foundation.MakeString(replacementString).String())
	return C.bool(result)
}

//export textViewDelegate_TextView_ShouldChangeTextInRanges_ReplacementStrings
func textViewDelegate_TextView_ShouldChangeTextInRanges_ReplacementStrings(hp C.uintptr_t, textView unsafe.Pointer, affectedRanges C.Array, replacementStrings C.Array) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if affectedRanges.len > 0 {
		defer C.free(affectedRanges.data)
	}
	affectedRangesSlice := unsafe.Slice((*unsafe.Pointer)(affectedRanges.data), int(affectedRanges.len))
	var goAffectedRanges = make([]foundation.Value, len(affectedRangesSlice))
	for idx, r := range affectedRangesSlice {
		goAffectedRanges[idx] = foundation.MakeValue(r)
	}
	if replacementStrings.len > 0 {
		defer C.free(replacementStrings.data)
	}
	replacementStringsSlice := unsafe.Slice((*unsafe.Pointer)(replacementStrings.data), int(replacementStrings.len))
	var goReplacementStrings = make([]string, len(replacementStringsSlice))
	for idx, r := range replacementStringsSlice {
		goReplacementStrings[idx] = foundation.MakeString(r).String()
	}
	result := delegate.TextView_ShouldChangeTextInRanges_ReplacementStrings(MakeTextView(textView), goAffectedRanges, goReplacementStrings)
	return C.bool(result)
}

//export textViewDelegate_TextView_ShouldChangeTypingAttributes_ToAttributes
func textViewDelegate_TextView_ShouldChangeTypingAttributes_ToAttributes(hp C.uintptr_t, textView unsafe.Pointer, oldTypingAttributes C.Dictionary, newTypingAttributes C.Dictionary) C.Dictionary {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if oldTypingAttributes.len > 0 {
		defer C.free(oldTypingAttributes.key_data)
		defer C.free(oldTypingAttributes.value_data)
	}
	oldTypingAttributesKeySlice := unsafe.Slice((*unsafe.Pointer)(oldTypingAttributes.key_data), int(oldTypingAttributes.len))
	oldTypingAttributesValueSlice := unsafe.Slice((*unsafe.Pointer)(oldTypingAttributes.value_data), int(oldTypingAttributes.len))
	var goOldTypingAttributes = make(map[string]objc.Object)
	for idx, k := range oldTypingAttributesKeySlice {
		v := oldTypingAttributesValueSlice[idx]
		goOldTypingAttributes[foundation.MakeString(k).String()] = objc.MakeObject(v)
	}
	if newTypingAttributes.len > 0 {
		defer C.free(newTypingAttributes.key_data)
		defer C.free(newTypingAttributes.value_data)
	}
	newTypingAttributesKeySlice := unsafe.Slice((*unsafe.Pointer)(newTypingAttributes.key_data), int(newTypingAttributes.len))
	newTypingAttributesValueSlice := unsafe.Slice((*unsafe.Pointer)(newTypingAttributes.value_data), int(newTypingAttributes.len))
	var goNewTypingAttributes = make(map[foundation.AttributedStringKey]objc.Object)
	for idx, k := range newTypingAttributesKeySlice {
		v := newTypingAttributesValueSlice[idx]
		goNewTypingAttributes[foundation.AttributedStringKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	result := delegate.TextView_ShouldChangeTypingAttributes_ToAttributes(MakeTextView(textView), goOldTypingAttributes, goNewTypingAttributes)
	var cResult C.Dictionary
	if len(result) > 0 {
		cResultKeyData := make([]unsafe.Pointer, len(result))
		cResultValueData := make([]unsafe.Pointer, len(result))
		var idx = 0
		for k, v := range result {
			cResultKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cResultValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cResult.key_data = unsafe.Pointer(&cResultKeyData[0])
		cResult.value_data = unsafe.Pointer(&cResultValueData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textViewDelegate_TextViewDidChangeTypingAttributes
func textViewDelegate_TextViewDidChangeTypingAttributes(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextViewDidChangeTypingAttributes(foundation.MakeNotification(notification))
}

//export textViewDelegate_TextView_ClickedOnCell_InRect_AtIndex
func textViewDelegate_TextView_ClickedOnCell_InRect_AtIndex(hp C.uintptr_t, textView unsafe.Pointer, cell unsafe.Pointer, cellFrame C.CGRect, charIndex C.uint) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextView_ClickedOnCell_InRect_AtIndex(MakeTextView(textView), objc.MakeObject(cell), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&cellFrame))), uint(charIndex))
}

//export textViewDelegate_TextView_DoubleClickedOnCell_InRect_AtIndex
func textViewDelegate_TextView_DoubleClickedOnCell_InRect_AtIndex(hp C.uintptr_t, textView unsafe.Pointer, cell unsafe.Pointer, cellFrame C.CGRect, charIndex C.uint) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextView_DoubleClickedOnCell_InRect_AtIndex(MakeTextView(textView), objc.MakeObject(cell), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&cellFrame))), uint(charIndex))
}

//export textViewDelegate_TextView_ClickedOnLink_AtIndex
func textViewDelegate_TextView_ClickedOnLink_AtIndex(hp C.uintptr_t, textView unsafe.Pointer, link unsafe.Pointer, charIndex C.uint) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_ClickedOnLink_AtIndex(MakeTextView(textView), objc.MakeObject(link), uint(charIndex))
	return C.bool(result)
}

//export textViewDelegate_TextView_ShouldSetSpellingState_Range
func textViewDelegate_TextView_ShouldSetSpellingState_Range(hp C.uintptr_t, textView unsafe.Pointer, value C.int, affectedCharRange C.NSRange) C.int {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_ShouldSetSpellingState_Range(MakeTextView(textView), int(value), foundation.FromNSRangePointer(unsafe.Pointer(&affectedCharRange)))
	return C.int(result)
}

//export textViewDelegate_TextView_DraggedCell_InRect_Event_AtIndex
func textViewDelegate_TextView_DraggedCell_InRect_Event_AtIndex(hp C.uintptr_t, view unsafe.Pointer, cell unsafe.Pointer, rect C.CGRect, event unsafe.Pointer, charIndex C.uint) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextView_DraggedCell_InRect_Event_AtIndex(MakeTextView(view), objc.MakeObject(cell), foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&rect))), MakeEvent(event), uint(charIndex))
}

//export textViewDelegate_TextView_WillShowSharingServicePicker_ForItems
func textViewDelegate_TextView_WillShowSharingServicePicker_ForItems(hp C.uintptr_t, textView unsafe.Pointer, servicePicker unsafe.Pointer, items C.Array) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if items.len > 0 {
		defer C.free(items.data)
	}
	itemsSlice := unsafe.Slice((*unsafe.Pointer)(items.data), int(items.len))
	var goItems = make([]objc.Object, len(itemsSlice))
	for idx, r := range itemsSlice {
		goItems[idx] = objc.MakeObject(r)
	}
	result := delegate.TextView_WillShowSharingServicePicker_ForItems(MakeTextView(textView), MakeSharingServicePicker(servicePicker), goItems)
	return objc.ExtractPtr(result)
}

//export textViewDelegate_TextView_DoCommandBySelector
func textViewDelegate_TextView_DoCommandBySelector(hp C.uintptr_t, textView unsafe.Pointer, commandSelector unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_DoCommandBySelector(MakeTextView(textView), objc.Selector(commandSelector))
	return C.bool(result)
}

//export textViewDelegate_TextView_Menu_ForEvent_AtIndex
func textViewDelegate_TextView_Menu_ForEvent_AtIndex(hp C.uintptr_t, view unsafe.Pointer, menu unsafe.Pointer, event unsafe.Pointer, charIndex C.uint) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_Menu_ForEvent_AtIndex(MakeTextView(view), MakeMenu(menu), MakeEvent(event), uint(charIndex))
	return objc.ExtractPtr(result)
}

//export textViewDelegate_TextView_Candidates_ForSelectedRange
func textViewDelegate_TextView_Candidates_ForSelectedRange(hp C.uintptr_t, textView unsafe.Pointer, candidates C.Array, selectedRange C.NSRange) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if candidates.len > 0 {
		defer C.free(candidates.data)
	}
	candidatesSlice := unsafe.Slice((*unsafe.Pointer)(candidates.data), int(candidates.len))
	var goCandidates = make([]foundation.TextCheckingResult, len(candidatesSlice))
	for idx, r := range candidatesSlice {
		goCandidates[idx] = foundation.MakeTextCheckingResult(r)
	}
	result := delegate.TextView_Candidates_ForSelectedRange(MakeTextView(textView), goCandidates, foundation.FromNSRangePointer(unsafe.Pointer(&selectedRange)))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = objc.ExtractPtr(v)
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textViewDelegate_TextView_CandidatesForSelectedRange
func textViewDelegate_TextView_CandidatesForSelectedRange(hp C.uintptr_t, textView unsafe.Pointer, selectedRange C.NSRange) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_CandidatesForSelectedRange(MakeTextView(textView), foundation.FromNSRangePointer(unsafe.Pointer(&selectedRange)))
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = objc.ExtractPtr(v)
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textViewDelegate_TextView_ShouldSelectCandidateAtIndex
func textViewDelegate_TextView_ShouldSelectCandidateAtIndex(hp C.uintptr_t, textView unsafe.Pointer, index C.uint) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextView_ShouldSelectCandidateAtIndex(MakeTextView(textView), uint(index))
	return C.bool(result)
}

//export textViewDelegate_TextView_ShouldUpdateTouchBarItemIdentifiers
func textViewDelegate_TextView_ShouldUpdateTouchBarItemIdentifiers(hp C.uintptr_t, textView unsafe.Pointer, identifiers C.Array) C.Array {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	if identifiers.len > 0 {
		defer C.free(identifiers.data)
	}
	identifiersSlice := unsafe.Slice((*unsafe.Pointer)(identifiers.data), int(identifiers.len))
	var goIdentifiers = make([]TouchBarItemIdentifier, len(identifiersSlice))
	for idx, r := range identifiersSlice {
		goIdentifiers[idx] = TouchBarItemIdentifier(foundation.MakeString(r).String())
	}
	result := delegate.TextView_ShouldUpdateTouchBarItemIdentifiers(MakeTextView(textView), goIdentifiers)
	var cResult C.Array
	if len(result) > 0 {
		cResultData := make([]unsafe.Pointer, len(result))
		for idx, v := range result {
			cResultData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cResult.data = unsafe.Pointer(&cResultData[0])
		cResult.len = C.int(len(result))
	}
	return cResult
}

//export textViewDelegate_TextDidChange
func textViewDelegate_TextDidChange(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextDidChange(foundation.MakeNotification(notification))
}

//export textViewDelegate_TextShouldBeginEditing
func textViewDelegate_TextShouldBeginEditing(hp C.uintptr_t, textObject unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextShouldBeginEditing(MakeText(textObject))
	return C.bool(result)
}

//export textViewDelegate_TextDidBeginEditing
func textViewDelegate_TextDidBeginEditing(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextDidBeginEditing(foundation.MakeNotification(notification))
}

//export textViewDelegate_TextShouldEndEditing
func textViewDelegate_TextShouldEndEditing(hp C.uintptr_t, textObject unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	result := delegate.TextShouldEndEditing(MakeText(textObject))
	return C.bool(result)
}

//export textViewDelegate_TextDidEndEditing
func textViewDelegate_TextDidEndEditing(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	delegate.TextDidEndEditing(foundation.MakeNotification(notification))
}

//export TextViewDelegate_RespondsTo
func TextViewDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*TextViewDelegate)
	switch selName {
	case "undoManagerForTextView:":
		return delegate.UndoManagerForTextView != nil
	case "textView:willDisplayToolTip:forCharacterAtIndex:":
		return delegate.TextView_WillDisplayToolTip_ForCharacterAtIndex != nil
	case "textView:URLForContentsOfTextAttachment:atIndex:":
		return delegate.TextView_URLForContentsOfTextAttachment_AtIndex != nil
	case "textView:willChangeSelectionFromCharacterRange:toCharacterRange:":
		return delegate.TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange != nil
	case "textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:":
		return delegate.TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges != nil
	case "textViewDidChangeSelection:":
		return delegate.TextViewDidChangeSelection != nil
	case "textView:writablePasteboardTypesForCell:atIndex:":
		return delegate.TextView_WritablePasteboardTypesForCell_AtIndex != nil
	case "textView:writeCell:atIndex:toPasteboard:type:":
		return delegate.TextView_WriteCell_AtIndex_ToPasteboard_Type != nil
	case "textView:shouldChangeTextInRange:replacementString:":
		return delegate.TextView_ShouldChangeTextInRange_ReplacementString != nil
	case "textView:shouldChangeTextInRanges:replacementStrings:":
		return delegate.TextView_ShouldChangeTextInRanges_ReplacementStrings != nil
	case "textView:shouldChangeTypingAttributes:toAttributes:":
		return delegate.TextView_ShouldChangeTypingAttributes_ToAttributes != nil
	case "textViewDidChangeTypingAttributes:":
		return delegate.TextViewDidChangeTypingAttributes != nil
	case "textView:clickedOnCell:inRect:atIndex:":
		return delegate.TextView_ClickedOnCell_InRect_AtIndex != nil
	case "textView:doubleClickedOnCell:inRect:atIndex:":
		return delegate.TextView_DoubleClickedOnCell_InRect_AtIndex != nil
	case "textView:clickedOnLink:atIndex:":
		return delegate.TextView_ClickedOnLink_AtIndex != nil
	case "textView:shouldSetSpellingState:range:":
		return delegate.TextView_ShouldSetSpellingState_Range != nil
	case "textView:draggedCell:inRect:event:atIndex:":
		return delegate.TextView_DraggedCell_InRect_Event_AtIndex != nil
	case "textView:willShowSharingServicePicker:forItems:":
		return delegate.TextView_WillShowSharingServicePicker_ForItems != nil
	case "textView:doCommandBySelector:":
		return delegate.TextView_DoCommandBySelector != nil
	case "textView:menu:forEvent:atIndex:":
		return delegate.TextView_Menu_ForEvent_AtIndex != nil
	case "textView:candidates:forSelectedRange:":
		return delegate.TextView_Candidates_ForSelectedRange != nil
	case "textView:candidatesForSelectedRange:":
		return delegate.TextView_CandidatesForSelectedRange != nil
	case "textView:shouldSelectCandidateAtIndex:":
		return delegate.TextView_ShouldSelectCandidateAtIndex != nil
	case "textView:shouldUpdateTouchBarItemIdentifiers:":
		return delegate.TextView_ShouldUpdateTouchBarItemIdentifiers != nil
	case "textDidChange:":
		return delegate.TextDidChange != nil
	case "textShouldBeginEditing:":
		return delegate.TextShouldBeginEditing != nil
	case "textDidBeginEditing:":
		return delegate.TextDidBeginEditing != nil
	case "textShouldEndEditing:":
		return delegate.TextShouldEndEditing != nil
	case "textDidEndEditing:":
		return delegate.TextDidEndEditing != nil
	default:
		return false
	}
}
