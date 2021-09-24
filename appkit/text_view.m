#import "text_view.h"
#import <AppKit/NSTextView.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

void* C_TextView_Alloc() {
    return [NSTextView alloc];
}

void* C_NSTextView_InitWithFrame_TextContainer(void* ptr, CGRect frameRect, void* container) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSTextView* result_ = [nSTextView initWithFrame:frameRect textContainer:(NSTextContainer*)container];
    return result_;
}

void* C_NSTextView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSTextView* result_ = [nSTextView initWithFrame:frameRect];
    return result_;
}

void* C_NSTextView_InitWithCoder(void* ptr, void* coder) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSTextView* result_ = [nSTextView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSTextView_TextView_FieldEditor() {
    NSTextView* result_ = [NSTextView fieldEditor];
    return result_;
}

void* C_NSTextView_Init(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSTextView* result_ = [nSTextView init];
    return result_;
}

void* C_NSTextView_AllocTextView() {
    NSTextView* result_ = [NSTextView alloc];
    return result_;
}

void* C_NSTextView_NewTextView() {
    NSTextView* result_ = [NSTextView new];
    return result_;
}

void* C_NSTextView_Autorelease(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSTextView* result_ = [nSTextView autorelease];
    return result_;
}

void* C_NSTextView_Retain(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSTextView* result_ = [nSTextView retain];
    return result_;
}

void C_NSTextView_TextView_RegisterForServices() {
    [NSTextView registerForServices];
}

void C_NSTextView_ReplaceTextContainer(void* ptr, void* newContainer) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView replaceTextContainer:(NSTextContainer*)newContainer];
}

void C_NSTextView_InvalidateTextContainerOrigin(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView invalidateTextContainerOrigin];
}

void C_NSTextView_ChangeDocumentBackgroundColor(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView changeDocumentBackgroundColor:(id)sender];
}

void C_NSTextView_SetNeedsDisplayInRect_AvoidAdditionalLayout(void* ptr, CGRect rect, bool flag) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setNeedsDisplayInRect:rect avoidAdditionalLayout:flag];
}

void C_NSTextView_DrawInsertionPointInRect_Color_TurnedOn(void* ptr, CGRect rect, void* color, bool flag) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView drawInsertionPointInRect:rect color:(NSColor*)color turnedOn:flag];
}

void C_NSTextView_DrawViewBackgroundInRect(void* ptr, CGRect rect) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView drawViewBackgroundInRect:rect];
}

void C_NSTextView_SetConstrainedFrameSize(void* ptr, CGSize desiredSize) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setConstrainedFrameSize:desiredSize];
}

void C_NSTextView_CleanUpAfterDragOperation(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView cleanUpAfterDragOperation];
}

void C_NSTextView_ShowFindIndicatorForRange(void* ptr, NSRange charRange) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView showFindIndicatorForRange:charRange];
}

void C_NSTextView_SetBaseWritingDirection_Range(void* ptr, int writingDirection, NSRange _range) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setBaseWritingDirection:writingDirection range:_range];
}

void C_NSTextView_Outline(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView outline:(id)sender];
}

void C_NSTextView_ToggleAutomaticQuoteSubstitution(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView toggleAutomaticQuoteSubstitution:(id)sender];
}

void C_NSTextView_ToggleAutomaticLinkDetection(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView toggleAutomaticLinkDetection:(id)sender];
}

void C_NSTextView_SetSelectedRanges_Affinity_StillSelecting(void* ptr, Array ranges, unsigned int affinity, bool stillSelectingFlag) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSMutableArray* objcRanges = [NSMutableArray arrayWithCapacity:ranges.len];
    if (ranges.len > 0) {
    	void** rangesData = (void**)ranges.data;
    	for (int i = 0; i < ranges.len; i++) {
    		void* p = rangesData[i];
    		[objcRanges addObject:(NSValue*)p];
    	}
    }
    [nSTextView setSelectedRanges:objcRanges affinity:affinity stillSelecting:stillSelectingFlag];
}

void C_NSTextView_UpdateInsertionPointStateAndRestartTimer(void* ptr, bool restartFlag) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView updateInsertionPointStateAndRestartTimer:restartFlag];
}

unsigned int C_NSTextView_CharacterIndexForInsertionAtPoint(void* ptr, CGPoint point) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSUInteger result_ = [nSTextView characterIndexForInsertionAtPoint:point];
    return result_;
}

void* C_NSTextView_PreferredPasteboardTypeFromArray_RestrictedToTypesFromArray(void* ptr, Array availableTypes, Array allowedTypes) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSMutableArray* objcAvailableTypes = [NSMutableArray arrayWithCapacity:availableTypes.len];
    if (availableTypes.len > 0) {
    	void** availableTypesData = (void**)availableTypes.data;
    	for (int i = 0; i < availableTypes.len; i++) {
    		void* p = availableTypesData[i];
    		[objcAvailableTypes addObject:(NSString*)p];
    	}
    }
    NSMutableArray* objcAllowedTypes = [NSMutableArray arrayWithCapacity:allowedTypes.len];
    if (allowedTypes.len > 0) {
    	void** allowedTypesData = (void**)allowedTypes.data;
    	for (int i = 0; i < allowedTypes.len; i++) {
    		void* p = allowedTypesData[i];
    		[objcAllowedTypes addObject:(NSString*)p];
    	}
    }
    NSPasteboardType result_ = [nSTextView preferredPasteboardTypeFromArray:objcAvailableTypes restrictedToTypesFromArray:objcAllowedTypes];
    return result_;
}

bool C_NSTextView_ReadSelectionFromPasteboard(void* ptr, void* pboard) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView readSelectionFromPasteboard:(NSPasteboard*)pboard];
    return result_;
}

bool C_NSTextView_ReadSelectionFromPasteboard_Type(void* ptr, void* pboard, void* _type) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView readSelectionFromPasteboard:(NSPasteboard*)pboard type:(NSString*)_type];
    return result_;
}

bool C_NSTextView_WriteSelectionToPasteboard_Type(void* ptr, void* pboard, void* _type) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView writeSelectionToPasteboard:(NSPasteboard*)pboard type:(NSString*)_type];
    return result_;
}

bool C_NSTextView_WriteSelectionToPasteboard_Types(void* ptr, void* pboard, Array types) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSMutableArray* objcTypes = [NSMutableArray arrayWithCapacity:types.len];
    if (types.len > 0) {
    	void** typesData = (void**)types.data;
    	for (int i = 0; i < types.len; i++) {
    		void* p = typesData[i];
    		[objcTypes addObject:(NSString*)p];
    	}
    }
    BOOL result_ = [nSTextView writeSelectionToPasteboard:(NSPasteboard*)pboard types:objcTypes];
    return result_;
}

void C_NSTextView_AlignJustified(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView alignJustified:(id)sender];
}

void C_NSTextView_ChangeAttributes(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView changeAttributes:(id)sender];
}

void C_NSTextView_ChangeColor(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView changeColor:(id)sender];
}

void C_NSTextView_SetAlignment_Range(void* ptr, int alignment, NSRange _range) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAlignment:alignment range:_range];
}

void C_NSTextView_UseStandardKerning(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView useStandardKerning:(id)sender];
}

void C_NSTextView_LowerBaseline(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView lowerBaseline:(id)sender];
}

void C_NSTextView_RaiseBaseline(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView raiseBaseline:(id)sender];
}

void C_NSTextView_TurnOffKerning(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView turnOffKerning:(id)sender];
}

void C_NSTextView_LoosenKerning(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView loosenKerning:(id)sender];
}

void C_NSTextView_TightenKerning(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView tightenKerning:(id)sender];
}

void C_NSTextView_UseStandardLigatures(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView useStandardLigatures:(id)sender];
}

void C_NSTextView_TurnOffLigatures(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView turnOffLigatures:(id)sender];
}

void C_NSTextView_UseAllLigatures(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView useAllLigatures:(id)sender];
}

void C_NSTextView_ClickedOnLink_AtIndex(void* ptr, void* link, unsigned int charIndex) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView clickedOnLink:(id)link atIndex:charIndex];
}

void C_NSTextView_PasteAsPlainText(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView pasteAsPlainText:(id)sender];
}

void C_NSTextView_PasteAsRichText(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView pasteAsRichText:(id)sender];
}

void C_NSTextView_BreakUndoCoalescing(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView breakUndoCoalescing];
}

void C_NSTextView_UpdateFontPanel(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView updateFontPanel];
}

void C_NSTextView_UpdateRuler(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView updateRuler];
}

void C_NSTextView_UpdateDragTypeRegistration(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView updateDragTypeRegistration];
}

NSRange C_NSTextView_SelectionRangeForProposedRange_Granularity(void* ptr, NSRange proposedCharRange, unsigned int granularity) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSRange result_ = [nSTextView selectionRangeForProposedRange:proposedCharRange granularity:granularity];
    return result_;
}

bool C_NSTextView_ShouldChangeTextInRange_ReplacementString(void* ptr, NSRange affectedCharRange, void* replacementString) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView shouldChangeTextInRange:affectedCharRange replacementString:(NSString*)replacementString];
    return result_;
}

bool C_NSTextView_ShouldChangeTextInRanges_ReplacementStrings(void* ptr, Array affectedRanges, Array replacementStrings) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSMutableArray* objcAffectedRanges = [NSMutableArray arrayWithCapacity:affectedRanges.len];
    if (affectedRanges.len > 0) {
    	void** affectedRangesData = (void**)affectedRanges.data;
    	for (int i = 0; i < affectedRanges.len; i++) {
    		void* p = affectedRangesData[i];
    		[objcAffectedRanges addObject:(NSValue*)p];
    	}
    }
    NSMutableArray* objcReplacementStrings = [NSMutableArray arrayWithCapacity:replacementStrings.len];
    if (replacementStrings.len > 0) {
    	void** replacementStringsData = (void**)replacementStrings.data;
    	for (int i = 0; i < replacementStrings.len; i++) {
    		void* p = replacementStringsData[i];
    		[objcReplacementStrings addObject:(NSString*)p];
    	}
    }
    BOOL result_ = [nSTextView shouldChangeTextInRanges:objcAffectedRanges replacementStrings:objcReplacementStrings];
    return result_;
}

void C_NSTextView_DidChangeText(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView didChangeText];
}

NSRange C_NSTextView_SmartDeleteRangeForProposedRange(void* ptr, NSRange proposedCharRange) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSRange result_ = [nSTextView smartDeleteRangeForProposedRange:proposedCharRange];
    return result_;
}

void* C_NSTextView_SmartInsertAfterStringForString_ReplacingRange(void* ptr, void* pasteString, NSRange charRangeToReplace) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSString* result_ = [nSTextView smartInsertAfterStringForString:(NSString*)pasteString replacingRange:charRangeToReplace];
    return result_;
}

void* C_NSTextView_SmartInsertBeforeStringForString_ReplacingRange(void* ptr, void* pasteString, NSRange charRangeToReplace) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSString* result_ = [nSTextView smartInsertBeforeStringForString:(NSString*)pasteString replacingRange:charRangeToReplace];
    return result_;
}

void C_NSTextView_ToggleSmartInsertDelete(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView toggleSmartInsertDelete:(id)sender];
}

void C_NSTextView_ToggleContinuousSpellChecking(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView toggleContinuousSpellChecking:(id)sender];
}

void C_NSTextView_ToggleGrammarChecking(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView toggleGrammarChecking:(id)sender];
}

void C_NSTextView_SetSpellingState_Range(void* ptr, int value, NSRange charRange) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setSpellingState:value range:charRange];
}

void C_NSTextView_OrderFrontSharingServicePicker(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView orderFrontSharingServicePicker:(id)sender];
}

unsigned int C_NSTextView_DragOperationForDraggingInfo_Type(void* ptr, void* dragInfo, void* _type) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSDragOperation result_ = [nSTextView dragOperationForDraggingInfo:(id)dragInfo type:(NSString*)_type];
    return result_;
}

bool C_NSTextView_DragSelectionWithEvent_Offset_SlideBack(void* ptr, void* event, CGSize mouseOffset, bool slideBack) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView dragSelectionWithEvent:(NSEvent*)event offset:mouseOffset slideBack:slideBack];
    return result_;
}

void C_NSTextView_StartSpeaking(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView startSpeaking:(id)sender];
}

void C_NSTextView_StopSpeaking(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView stopSpeaking:(id)sender];
}

void C_NSTextView_PerformFindPanelAction(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView performFindPanelAction:(id)sender];
}

void C_NSTextView_OrderFrontLinkPanel(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView orderFrontLinkPanel:(id)sender];
}

void C_NSTextView_OrderFrontListPanel(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView orderFrontListPanel:(id)sender];
}

void C_NSTextView_OrderFrontSpacingPanel(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView orderFrontSpacingPanel:(id)sender];
}

void C_NSTextView_OrderFrontTablePanel(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView orderFrontTablePanel:(id)sender];
}

void C_NSTextView_OrderFrontSubstitutionsPanel(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView orderFrontSubstitutionsPanel:(id)sender];
}

void C_NSTextView_Complete(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView complete:(id)sender];
}

void C_NSTextView_InsertCompletion_ForPartialWordRange_Movement_IsFinal(void* ptr, void* word, NSRange charRange, int movement, bool flag) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView insertCompletion:(NSString*)word forPartialWordRange:charRange movement:movement isFinal:flag];
}

void C_NSTextView_CheckTextInDocument(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView checkTextInDocument:(id)sender];
}

void C_NSTextView_CheckTextInSelection(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView checkTextInSelection:(id)sender];
}

void C_NSTextView_ToggleAutomaticDashSubstitution(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView toggleAutomaticDashSubstitution:(id)sender];
}

void C_NSTextView_ToggleAutomaticDataDetection(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView toggleAutomaticDataDetection:(id)sender];
}

void C_NSTextView_ToggleAutomaticSpellingCorrection(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView toggleAutomaticSpellingCorrection:(id)sender];
}

void C_NSTextView_ToggleAutomaticTextReplacement(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView toggleAutomaticTextReplacement:(id)sender];
}

void C_NSTextView_UpdateQuickLookPreviewPanel(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView updateQuickLookPreviewPanel];
}

void C_NSTextView_ToggleQuickLookPreviewPanel(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView toggleQuickLookPreviewPanel:(id)sender];
}

void C_NSTextView_ChangeLayoutOrientation(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView changeLayoutOrientation:(id)sender];
}

void C_NSTextView_SetLayoutOrientation(void* ptr, int orientation) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setLayoutOrientation:orientation];
}

bool C_NSTextView_PerformValidatedReplacementInRange_WithAttributedString(void* ptr, NSRange _range, void* attributedString) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView performValidatedReplacementInRange:_range withAttributedString:(NSAttributedString*)attributedString];
    return result_;
}

void C_NSTextView_ToggleAutomaticTextCompletion(void* ptr, void* sender) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView toggleAutomaticTextCompletion:(id)sender];
}

void C_NSTextView_UpdateCandidates(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView updateCandidates];
}

void C_NSTextView_UpdateTextTouchBarItems(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView updateTextTouchBarItems];
}

void C_NSTextView_UpdateTouchBarItemIdentifiers(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView updateTouchBarItemIdentifiers];
}

void* C_NSTextView_ScrollableDocumentContentTextView() {
    NSScrollView* result_ = [NSTextView scrollableDocumentContentTextView];
    return result_;
}

void* C_NSTextView_ScrollablePlainDocumentContentTextView() {
    NSScrollView* result_ = [NSTextView scrollablePlainDocumentContentTextView];
    return result_;
}

void* C_NSTextView_ScrollableTextView() {
    NSScrollView* result_ = [NSTextView scrollableTextView];
    return result_;
}

void* C_NSTextView_TextContainer(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSTextContainer* result_ = [nSTextView textContainer];
    return result_;
}

void C_NSTextView_SetTextContainer(void* ptr, void* value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setTextContainer:(NSTextContainer*)value];
}

CGSize C_NSTextView_TextContainerInset(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSSize result_ = [nSTextView textContainerInset];
    return result_;
}

void C_NSTextView_SetTextContainerInset(void* ptr, CGSize value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setTextContainerInset:value];
}

CGPoint C_NSTextView_TextContainerOrigin(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSPoint result_ = [nSTextView textContainerOrigin];
    return result_;
}

void* C_NSTextView_LayoutManager(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSLayoutManager* result_ = [nSTextView layoutManager];
    return result_;
}

void* C_NSTextView_TextStorage(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSTextStorage* result_ = [nSTextView textStorage];
    return result_;
}

bool C_NSTextView_AllowsDocumentBackgroundColorChange(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView allowsDocumentBackgroundColorChange];
    return result_;
}

void C_NSTextView_SetAllowsDocumentBackgroundColorChange(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAllowsDocumentBackgroundColorChange:value];
}

bool C_NSTextView_ShouldDrawInsertionPoint(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView shouldDrawInsertionPoint];
    return result_;
}

Array C_NSTextView_AllowedInputSourceLocales(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSArray* result_ = [nSTextView allowedInputSourceLocales];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSTextView_SetAllowedInputSourceLocales(void* ptr, Array value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)p];
    	}
    }
    [nSTextView setAllowedInputSourceLocales:objcValue];
}

bool C_NSTextView_AllowsUndo(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView allowsUndo];
    return result_;
}

void C_NSTextView_SetAllowsUndo(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAllowsUndo:value];
}

void* C_NSTextView_DefaultParagraphStyle(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSParagraphStyle* result_ = [nSTextView defaultParagraphStyle];
    return result_;
}

void C_NSTextView_SetDefaultParagraphStyle(void* ptr, void* value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setDefaultParagraphStyle:(NSParagraphStyle*)value];
}

bool C_NSTextView_AllowsImageEditing(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView allowsImageEditing];
    return result_;
}

void C_NSTextView_SetAllowsImageEditing(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAllowsImageEditing:value];
}

bool C_NSTextView_IsAutomaticQuoteSubstitutionEnabled(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView isAutomaticQuoteSubstitutionEnabled];
    return result_;
}

void C_NSTextView_SetAutomaticQuoteSubstitutionEnabled(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAutomaticQuoteSubstitutionEnabled:value];
}

bool C_NSTextView_IsAutomaticLinkDetectionEnabled(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView isAutomaticLinkDetectionEnabled];
    return result_;
}

void C_NSTextView_SetAutomaticLinkDetectionEnabled(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAutomaticLinkDetectionEnabled:value];
}

bool C_NSTextView_DisplaysLinkToolTips(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView displaysLinkToolTips];
    return result_;
}

void C_NSTextView_SetDisplaysLinkToolTips(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setDisplaysLinkToolTips:value];
}

bool C_NSTextView_UsesRuler(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView usesRuler];
    return result_;
}

void C_NSTextView_SetUsesRuler(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setUsesRuler:value];
}

void C_NSTextView_SetRulerVisible(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setRulerVisible:value];
}

bool C_NSTextView_UsesInspectorBar(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView usesInspectorBar];
    return result_;
}

void C_NSTextView_SetUsesInspectorBar(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setUsesInspectorBar:value];
}

Array C_NSTextView_SelectedRanges(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSArray* result_ = [nSTextView selectedRanges];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSTextView_SetSelectedRanges(void* ptr, Array value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSValue*)p];
    	}
    }
    [nSTextView setSelectedRanges:objcValue];
}

unsigned int C_NSTextView_SelectionAffinity(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSSelectionAffinity result_ = [nSTextView selectionAffinity];
    return result_;
}

unsigned int C_NSTextView_SelectionGranularity(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSSelectionGranularity result_ = [nSTextView selectionGranularity];
    return result_;
}

void C_NSTextView_SetSelectionGranularity(void* ptr, unsigned int value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setSelectionGranularity:value];
}

void* C_NSTextView_InsertionPointColor(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSColor* result_ = [nSTextView insertionPointColor];
    return result_;
}

void C_NSTextView_SetInsertionPointColor(void* ptr, void* value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setInsertionPointColor:(NSColor*)value];
}

Dictionary C_NSTextView_SelectedTextAttributes(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSDictionary* result_ = [nSTextView selectedTextAttributes];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSAttributedStringKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void C_NSTextView_SetSelectedTextAttributes(void* ptr, Dictionary value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSTextView setSelectedTextAttributes:objcValue];
}

Dictionary C_NSTextView_MarkedTextAttributes(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSDictionary* result_ = [nSTextView markedTextAttributes];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSAttributedStringKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void C_NSTextView_SetMarkedTextAttributes(void* ptr, Dictionary value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSTextView setMarkedTextAttributes:objcValue];
}

Dictionary C_NSTextView_LinkTextAttributes(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSDictionary* result_ = [nSTextView linkTextAttributes];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSAttributedStringKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void C_NSTextView_SetLinkTextAttributes(void* ptr, Dictionary value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSTextView setLinkTextAttributes:objcValue];
}

Array C_NSTextView_ReadablePasteboardTypes(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSArray* result_ = [nSTextView readablePasteboardTypes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_NSTextView_WritablePasteboardTypes(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSArray* result_ = [nSTextView writablePasteboardTypes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Dictionary C_NSTextView_TypingAttributes(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSDictionary* result_ = [nSTextView typingAttributes];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSAttributedStringKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void C_NSTextView_SetTypingAttributes(void* ptr, Dictionary value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSTextView setTypingAttributes:objcValue];
}

bool C_NSTextView_IsCoalescingUndo(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView isCoalescingUndo];
    return result_;
}

Array C_NSTextView_AcceptableDragTypes(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSArray* result_ = [nSTextView acceptableDragTypes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

NSRange C_NSTextView_RangeForUserCharacterAttributeChange(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSRange result_ = [nSTextView rangeForUserCharacterAttributeChange];
    return result_;
}

Array C_NSTextView_RangesForUserCharacterAttributeChange(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSArray* result_ = [nSTextView rangesForUserCharacterAttributeChange];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

NSRange C_NSTextView_RangeForUserParagraphAttributeChange(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSRange result_ = [nSTextView rangeForUserParagraphAttributeChange];
    return result_;
}

Array C_NSTextView_RangesForUserParagraphAttributeChange(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSArray* result_ = [nSTextView rangesForUserParagraphAttributeChange];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

NSRange C_NSTextView_RangeForUserTextChange(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSRange result_ = [nSTextView rangeForUserTextChange];
    return result_;
}

Array C_NSTextView_RangesForUserTextChange(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSArray* result_ = [nSTextView rangesForUserTextChange];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

bool C_NSTextView_SmartInsertDeleteEnabled(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView smartInsertDeleteEnabled];
    return result_;
}

void C_NSTextView_SetSmartInsertDeleteEnabled(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setSmartInsertDeleteEnabled:value];
}

bool C_NSTextView_IsContinuousSpellCheckingEnabled(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView isContinuousSpellCheckingEnabled];
    return result_;
}

void C_NSTextView_SetContinuousSpellCheckingEnabled(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setContinuousSpellCheckingEnabled:value];
}

int C_NSTextView_SpellCheckerDocumentTag(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSInteger result_ = [nSTextView spellCheckerDocumentTag];
    return result_;
}

bool C_NSTextView_IsGrammarCheckingEnabled(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView isGrammarCheckingEnabled];
    return result_;
}

void C_NSTextView_SetGrammarCheckingEnabled(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setGrammarCheckingEnabled:value];
}

bool C_NSTextView_AcceptsGlyphInfo(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView acceptsGlyphInfo];
    return result_;
}

void C_NSTextView_SetAcceptsGlyphInfo(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAcceptsGlyphInfo:value];
}

bool C_NSTextView_UsesFindPanel(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView usesFindPanel];
    return result_;
}

void C_NSTextView_SetUsesFindPanel(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setUsesFindPanel:value];
}

NSRange C_NSTextView_RangeForUserCompletion(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    NSRange result_ = [nSTextView rangeForUserCompletion];
    return result_;
}

bool C_NSTextView_IsAutomaticDashSubstitutionEnabled(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView isAutomaticDashSubstitutionEnabled];
    return result_;
}

void C_NSTextView_SetAutomaticDashSubstitutionEnabled(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAutomaticDashSubstitutionEnabled:value];
}

bool C_NSTextView_IsAutomaticDataDetectionEnabled(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView isAutomaticDataDetectionEnabled];
    return result_;
}

void C_NSTextView_SetAutomaticDataDetectionEnabled(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAutomaticDataDetectionEnabled:value];
}

bool C_NSTextView_IsAutomaticSpellingCorrectionEnabled(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView isAutomaticSpellingCorrectionEnabled];
    return result_;
}

void C_NSTextView_SetAutomaticSpellingCorrectionEnabled(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAutomaticSpellingCorrectionEnabled:value];
}

bool C_NSTextView_IsAutomaticTextReplacementEnabled(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView isAutomaticTextReplacementEnabled];
    return result_;
}

void C_NSTextView_SetAutomaticTextReplacementEnabled(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAutomaticTextReplacementEnabled:value];
}

bool C_NSTextView_UsesFindBar(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView usesFindBar];
    return result_;
}

void C_NSTextView_SetUsesFindBar(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setUsesFindBar:value];
}

bool C_NSTextView_IsIncrementalSearchingEnabled(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView isIncrementalSearchingEnabled];
    return result_;
}

void C_NSTextView_SetIncrementalSearchingEnabled(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setIncrementalSearchingEnabled:value];
}

bool C_NSTextView_AllowsCharacterPickerTouchBarItem(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView allowsCharacterPickerTouchBarItem];
    return result_;
}

void C_NSTextView_SetAllowsCharacterPickerTouchBarItem(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAllowsCharacterPickerTouchBarItem:value];
}

bool C_NSTextView_IsAutomaticTextCompletionEnabled(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView isAutomaticTextCompletionEnabled];
    return result_;
}

void C_NSTextView_SetAutomaticTextCompletionEnabled(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setAutomaticTextCompletionEnabled:value];
}

bool C_NSTextView_UsesAdaptiveColorMappingForDarkAppearance(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView usesAdaptiveColorMappingForDarkAppearance];
    return result_;
}

void C_NSTextView_SetUsesAdaptiveColorMappingForDarkAppearance(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setUsesAdaptiveColorMappingForDarkAppearance:value];
}

bool C_NSTextView_UsesRolloverButtonForSelection(void* ptr) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    BOOL result_ = [nSTextView usesRolloverButtonForSelection];
    return result_;
}

void C_NSTextView_SetUsesRolloverButtonForSelection(void* ptr, bool value) {
    NSTextView* nSTextView = (NSTextView*)ptr;
    [nSTextView setUsesRolloverButtonForSelection:value];
}

bool C_NSTextView_TextView_StronglyReferencesTextStorage() {
    BOOL result_ = [NSTextView stronglyReferencesTextStorage];
    return result_;
}

@interface NSTextViewDelegateAdaptor : NSObject <NSTextViewDelegate>
@property (assign) uintptr_t goID;
@end

@implementation NSTextViewDelegateAdaptor


- (NSUndoManager*)undoManagerForTextView:(NSTextView*)view {
    void* result_ = textViewDelegate_UndoManagerForTextView([self goID], view);
    return (NSUndoManager*)result_;
}

- (NSString*)textView:(NSTextView*)textView willDisplayToolTip:(NSString*)tooltip forCharacterAtIndex:(NSUInteger)characterIndex {
    void* result_ = textViewDelegate_TextView_WillDisplayToolTip_ForCharacterAtIndex([self goID], textView, tooltip, characterIndex);
    return (NSString*)result_;
}

- (NSURL*)textView:(NSTextView*)textView URLForContentsOfTextAttachment:(NSTextAttachment*)textAttachment atIndex:(NSUInteger)charIndex {
    void* result_ = textViewDelegate_TextView_URLForContentsOfTextAttachment_AtIndex([self goID], textView, textAttachment, charIndex);
    return (NSURL*)result_;
}

- (NSRange)textView:(NSTextView*)textView willChangeSelectionFromCharacterRange:(NSRange)oldSelectedCharRange toCharacterRange:(NSRange)newSelectedCharRange {
    NSRange result_ = textViewDelegate_TextView_WillChangeSelectionFromCharacterRange_ToCharacterRange([self goID], textView, oldSelectedCharRange, newSelectedCharRange);
    return result_;
}

- (NSArray*)textView:(NSTextView*)textView willChangeSelectionFromCharacterRanges:(NSArray*)oldSelectedCharRanges toCharacterRanges:(NSArray*)newSelectedCharRanges {
    Array oldSelectedCharRangesArray;
    int oldSelectedCharRangescount = [oldSelectedCharRanges count];
    if (oldSelectedCharRangescount > 0) {
    	void** oldSelectedCharRangesData = malloc(oldSelectedCharRangescount * sizeof(void*));
    	for (int i = 0; i < oldSelectedCharRangescount; i++) {
    		 void* p = [oldSelectedCharRanges objectAtIndex:i];
    		 oldSelectedCharRangesData[i] = p;
    	}
    	oldSelectedCharRangesArray.data = oldSelectedCharRangesData;
    	oldSelectedCharRangesArray.len = oldSelectedCharRangescount;
    }
    Array newSelectedCharRangesArray;
    int newSelectedCharRangescount = [newSelectedCharRanges count];
    if (newSelectedCharRangescount > 0) {
    	void** newSelectedCharRangesData = malloc(newSelectedCharRangescount * sizeof(void*));
    	for (int i = 0; i < newSelectedCharRangescount; i++) {
    		 void* p = [newSelectedCharRanges objectAtIndex:i];
    		 newSelectedCharRangesData[i] = p;
    	}
    	newSelectedCharRangesArray.data = newSelectedCharRangesData;
    	newSelectedCharRangesArray.len = newSelectedCharRangescount;
    }
    Array result_ = textViewDelegate_TextView_WillChangeSelectionFromCharacterRanges_ToCharacterRanges([self goID], textView, oldSelectedCharRangesArray, newSelectedCharRangesArray);
    NSMutableArray* objcResult_ = [NSMutableArray arrayWithCapacity:result_.len];
    if (result_.len > 0) {
    	void** result_Data = (void**)result_.data;
    	for (int i = 0; i < result_.len; i++) {
    		void* p = result_Data[i];
    		[objcResult_ addObject:(NSValue*)p];
    	}
    }
    return objcResult_;
}

- (void)textViewDidChangeSelection:(NSNotification*)notification {
    textViewDelegate_TextViewDidChangeSelection([self goID], notification);
}

- (NSArray*)textView:(NSTextView*)view writablePasteboardTypesForCell:(id)cell atIndex:(NSUInteger)charIndex {
    Array result_ = textViewDelegate_TextView_WritablePasteboardTypesForCell_AtIndex([self goID], view, cell, charIndex);
    NSMutableArray* objcResult_ = [NSMutableArray arrayWithCapacity:result_.len];
    if (result_.len > 0) {
    	void** result_Data = (void**)result_.data;
    	for (int i = 0; i < result_.len; i++) {
    		void* p = result_Data[i];
    		[objcResult_ addObject:(NSString*)p];
    	}
    }
    return objcResult_;
}

- (BOOL)textView:(NSTextView*)view writeCell:(id)cell atIndex:(NSUInteger)charIndex toPasteboard:(NSPasteboard*)pboard type:(NSPasteboardType)_type {
    bool result_ = textViewDelegate_TextView_WriteCell_AtIndex_ToPasteboard_Type([self goID], view, cell, charIndex, pboard, _type);
    return result_;
}

- (BOOL)textView:(NSTextView*)textView shouldChangeTextInRange:(NSRange)affectedCharRange replacementString:(NSString*)replacementString {
    bool result_ = textViewDelegate_TextView_ShouldChangeTextInRange_ReplacementString([self goID], textView, affectedCharRange, replacementString);
    return result_;
}

- (BOOL)textView:(NSTextView*)textView shouldChangeTextInRanges:(NSArray*)affectedRanges replacementStrings:(NSArray*)replacementStrings {
    Array affectedRangesArray;
    int affectedRangescount = [affectedRanges count];
    if (affectedRangescount > 0) {
    	void** affectedRangesData = malloc(affectedRangescount * sizeof(void*));
    	for (int i = 0; i < affectedRangescount; i++) {
    		 void* p = [affectedRanges objectAtIndex:i];
    		 affectedRangesData[i] = p;
    	}
    	affectedRangesArray.data = affectedRangesData;
    	affectedRangesArray.len = affectedRangescount;
    }
    Array replacementStringsArray;
    int replacementStringscount = [replacementStrings count];
    if (replacementStringscount > 0) {
    	void** replacementStringsData = malloc(replacementStringscount * sizeof(void*));
    	for (int i = 0; i < replacementStringscount; i++) {
    		 void* p = [replacementStrings objectAtIndex:i];
    		 replacementStringsData[i] = p;
    	}
    	replacementStringsArray.data = replacementStringsData;
    	replacementStringsArray.len = replacementStringscount;
    }
    bool result_ = textViewDelegate_TextView_ShouldChangeTextInRanges_ReplacementStrings([self goID], textView, affectedRangesArray, replacementStringsArray);
    return result_;
}

- (NSDictionary*)textView:(NSTextView*)textView shouldChangeTypingAttributes:(NSDictionary*)oldTypingAttributes toAttributes:(NSDictionary*)newTypingAttributes {
    Dictionary oldTypingAttributesArray;
    NSArray * oldTypingAttributesKeys = [oldTypingAttributes allKeys];
    int oldTypingAttributesCount = [oldTypingAttributesKeys count];
    if (oldTypingAttributesCount > 0) {
    	void** oldTypingAttributesKeyData = malloc(oldTypingAttributesCount * sizeof(void*));
    	void** oldTypingAttributesValueData = malloc(oldTypingAttributesCount * sizeof(void*));
    	for (int i = 0; i < oldTypingAttributesCount; i++) {
    		NSString* kp = [oldTypingAttributesKeys objectAtIndex:i];
    		id vp = oldTypingAttributes[kp];
    		 oldTypingAttributesKeyData[i] = kp;
    		 oldTypingAttributesValueData[i] = vp;
    	}
    	oldTypingAttributesArray.key_data = oldTypingAttributesKeyData;
    	oldTypingAttributesArray.value_data = oldTypingAttributesValueData;
    	oldTypingAttributesArray.len = oldTypingAttributesCount;
    }
    Dictionary newTypingAttributesArray;
    NSArray * newTypingAttributesKeys = [newTypingAttributes allKeys];
    int newTypingAttributesCount = [newTypingAttributesKeys count];
    if (newTypingAttributesCount > 0) {
    	void** newTypingAttributesKeyData = malloc(newTypingAttributesCount * sizeof(void*));
    	void** newTypingAttributesValueData = malloc(newTypingAttributesCount * sizeof(void*));
    	for (int i = 0; i < newTypingAttributesCount; i++) {
    		NSAttributedStringKey kp = [newTypingAttributesKeys objectAtIndex:i];
    		id vp = newTypingAttributes[kp];
    		 newTypingAttributesKeyData[i] = kp;
    		 newTypingAttributesValueData[i] = vp;
    	}
    	newTypingAttributesArray.key_data = newTypingAttributesKeyData;
    	newTypingAttributesArray.value_data = newTypingAttributesValueData;
    	newTypingAttributesArray.len = newTypingAttributesCount;
    }
    Dictionary result_ = textViewDelegate_TextView_ShouldChangeTypingAttributes_ToAttributes([self goID], textView, oldTypingAttributesArray, newTypingAttributesArray);
    NSMutableDictionary* objcResult_ = [NSMutableDictionary dictionaryWithCapacity:result_.len];
    if (result_.len > 0) {
    	void** result_KeyData = (void**)result_.key_data;
    	void** result_ValueData = (void**)result_.value_data;
    	for (int i = 0; i < result_.len; i++) {
    		void* kp = result_KeyData[i];
    		void* vp = result_ValueData[i];
    		[objcResult_ setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    return objcResult_;
}

- (void)textViewDidChangeTypingAttributes:(NSNotification*)notification {
    textViewDelegate_TextViewDidChangeTypingAttributes([self goID], notification);
}

- (void)textView:(NSTextView*)textView clickedOnCell:(id)cell inRect:(NSRect)cellFrame atIndex:(NSUInteger)charIndex {
    textViewDelegate_TextView_ClickedOnCell_InRect_AtIndex([self goID], textView, cell, cellFrame, charIndex);
}

- (void)textView:(NSTextView*)textView doubleClickedOnCell:(id)cell inRect:(NSRect)cellFrame atIndex:(NSUInteger)charIndex {
    textViewDelegate_TextView_DoubleClickedOnCell_InRect_AtIndex([self goID], textView, cell, cellFrame, charIndex);
}

- (BOOL)textView:(NSTextView*)textView clickedOnLink:(id)link atIndex:(NSUInteger)charIndex {
    bool result_ = textViewDelegate_TextView_ClickedOnLink_AtIndex([self goID], textView, link, charIndex);
    return result_;
}

- (NSInteger)textView:(NSTextView*)textView shouldSetSpellingState:(NSInteger)value range:(NSRange)affectedCharRange {
    int result_ = textViewDelegate_TextView_ShouldSetSpellingState_Range([self goID], textView, value, affectedCharRange);
    return result_;
}

- (void)textView:(NSTextView*)view draggedCell:(id)cell inRect:(NSRect)rect event:(NSEvent*)event atIndex:(NSUInteger)charIndex {
    textViewDelegate_TextView_DraggedCell_InRect_Event_AtIndex([self goID], view, cell, rect, event, charIndex);
}

- (NSSharingServicePicker*)textView:(NSTextView*)textView willShowSharingServicePicker:(NSSharingServicePicker*)servicePicker forItems:(NSArray*)items {
    Array itemsArray;
    int itemscount = [items count];
    if (itemscount > 0) {
    	void** itemsData = malloc(itemscount * sizeof(void*));
    	for (int i = 0; i < itemscount; i++) {
    		 void* p = [items objectAtIndex:i];
    		 itemsData[i] = p;
    	}
    	itemsArray.data = itemsData;
    	itemsArray.len = itemscount;
    }
    void* result_ = textViewDelegate_TextView_WillShowSharingServicePicker_ForItems([self goID], textView, servicePicker, itemsArray);
    return (NSSharingServicePicker*)result_;
}

- (BOOL)textView:(NSTextView*)textView doCommandBySelector:(SEL)commandSelector {
    bool result_ = textViewDelegate_TextView_DoCommandBySelector([self goID], textView, commandSelector);
    return result_;
}

- (NSMenu*)textView:(NSTextView*)view menu:(NSMenu*)menu forEvent:(NSEvent*)event atIndex:(NSUInteger)charIndex {
    void* result_ = textViewDelegate_TextView_Menu_ForEvent_AtIndex([self goID], view, menu, event, charIndex);
    return (NSMenu*)result_;
}

- (NSArray*)textView:(NSTextView*)textView candidates:(NSArray*)candidates forSelectedRange:(NSRange)selectedRange {
    Array candidatesArray;
    int candidatescount = [candidates count];
    if (candidatescount > 0) {
    	void** candidatesData = malloc(candidatescount * sizeof(void*));
    	for (int i = 0; i < candidatescount; i++) {
    		 void* p = [candidates objectAtIndex:i];
    		 candidatesData[i] = p;
    	}
    	candidatesArray.data = candidatesData;
    	candidatesArray.len = candidatescount;
    }
    Array result_ = textViewDelegate_TextView_Candidates_ForSelectedRange([self goID], textView, candidatesArray, selectedRange);
    NSMutableArray* objcResult_ = [NSMutableArray arrayWithCapacity:result_.len];
    if (result_.len > 0) {
    	void** result_Data = (void**)result_.data;
    	for (int i = 0; i < result_.len; i++) {
    		void* p = result_Data[i];
    		[objcResult_ addObject:(NSTextCheckingResult*)p];
    	}
    }
    return objcResult_;
}

- (NSArray*)textView:(NSTextView*)textView candidatesForSelectedRange:(NSRange)selectedRange {
    Array result_ = textViewDelegate_TextView_CandidatesForSelectedRange([self goID], textView, selectedRange);
    NSMutableArray* objcResult_ = [NSMutableArray arrayWithCapacity:result_.len];
    if (result_.len > 0) {
    	void** result_Data = (void**)result_.data;
    	for (int i = 0; i < result_.len; i++) {
    		void* p = result_Data[i];
    		[objcResult_ addObject:(NSObject*)p];
    	}
    }
    return objcResult_;
}

- (BOOL)textView:(NSTextView*)textView shouldSelectCandidateAtIndex:(NSUInteger)index {
    bool result_ = textViewDelegate_TextView_ShouldSelectCandidateAtIndex([self goID], textView, index);
    return result_;
}

- (NSArray*)textView:(NSTextView*)textView shouldUpdateTouchBarItemIdentifiers:(NSArray*)identifiers {
    Array identifiersArray;
    int identifierscount = [identifiers count];
    if (identifierscount > 0) {
    	void** identifiersData = malloc(identifierscount * sizeof(void*));
    	for (int i = 0; i < identifierscount; i++) {
    		 void* p = [identifiers objectAtIndex:i];
    		 identifiersData[i] = p;
    	}
    	identifiersArray.data = identifiersData;
    	identifiersArray.len = identifierscount;
    }
    Array result_ = textViewDelegate_TextView_ShouldUpdateTouchBarItemIdentifiers([self goID], textView, identifiersArray);
    NSMutableArray* objcResult_ = [NSMutableArray arrayWithCapacity:result_.len];
    if (result_.len > 0) {
    	void** result_Data = (void**)result_.data;
    	for (int i = 0; i < result_.len; i++) {
    		void* p = result_Data[i];
    		[objcResult_ addObject:(NSString*)p];
    	}
    }
    return objcResult_;
}

- (void)textDidChange:(NSNotification*)notification {
    textViewDelegate_TextDidChange([self goID], notification);
}

- (BOOL)textShouldBeginEditing:(NSText*)textObject {
    bool result_ = textViewDelegate_TextShouldBeginEditing([self goID], textObject);
    return result_;
}

- (void)textDidBeginEditing:(NSNotification*)notification {
    textViewDelegate_TextDidBeginEditing([self goID], notification);
}

- (BOOL)textShouldEndEditing:(NSText*)textObject {
    bool result_ = textViewDelegate_TextShouldEndEditing([self goID], textObject);
    return result_;
}

- (void)textDidEndEditing:(NSNotification*)notification {
    textViewDelegate_TextDidEndEditing([self goID], notification);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TextViewDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapTextViewDelegate(uintptr_t goID) {
    NSTextViewDelegateAdaptor* adaptor = [[NSTextViewDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
