#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TextView_Alloc();

void* C_NSTextView_InitWithFrame_TextContainer(void* ptr, CGRect frameRect, void* container);
void* C_NSTextView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSTextView_InitWithCoder(void* ptr, void* coder);
void* C_NSTextView_TextView_FieldEditor();
void* C_NSTextView_Init(void* ptr);
void* C_NSTextView_AllocTextView();
void* C_NSTextView_NewTextView();
void* C_NSTextView_Autorelease(void* ptr);
void* C_NSTextView_Retain(void* ptr);
void C_NSTextView_TextView_RegisterForServices();
void C_NSTextView_ReplaceTextContainer(void* ptr, void* newContainer);
void C_NSTextView_InvalidateTextContainerOrigin(void* ptr);
void C_NSTextView_ChangeDocumentBackgroundColor(void* ptr, void* sender);
void C_NSTextView_SetNeedsDisplayInRect_AvoidAdditionalLayout(void* ptr, CGRect rect, bool flag);
void C_NSTextView_DrawInsertionPointInRect_Color_TurnedOn(void* ptr, CGRect rect, void* color, bool flag);
void C_NSTextView_DrawViewBackgroundInRect(void* ptr, CGRect rect);
void C_NSTextView_SetConstrainedFrameSize(void* ptr, CGSize desiredSize);
void C_NSTextView_CleanUpAfterDragOperation(void* ptr);
void C_NSTextView_ShowFindIndicatorForRange(void* ptr, NSRange charRange);
void C_NSTextView_SetBaseWritingDirection_Range(void* ptr, int writingDirection, NSRange _range);
void C_NSTextView_Outline(void* ptr, void* sender);
void C_NSTextView_ToggleAutomaticQuoteSubstitution(void* ptr, void* sender);
void C_NSTextView_ToggleAutomaticLinkDetection(void* ptr, void* sender);
void C_NSTextView_SetSelectedRanges_Affinity_StillSelecting(void* ptr, Array ranges, unsigned int affinity, bool stillSelectingFlag);
void C_NSTextView_UpdateInsertionPointStateAndRestartTimer(void* ptr, bool restartFlag);
unsigned int C_NSTextView_CharacterIndexForInsertionAtPoint(void* ptr, CGPoint point);
void* C_NSTextView_PreferredPasteboardTypeFromArray_RestrictedToTypesFromArray(void* ptr, Array availableTypes, Array allowedTypes);
bool C_NSTextView_ReadSelectionFromPasteboard(void* ptr, void* pboard);
bool C_NSTextView_ReadSelectionFromPasteboard_Type(void* ptr, void* pboard, void* _type);
bool C_NSTextView_WriteSelectionToPasteboard_Type(void* ptr, void* pboard, void* _type);
bool C_NSTextView_WriteSelectionToPasteboard_Types(void* ptr, void* pboard, Array types);
void C_NSTextView_AlignJustified(void* ptr, void* sender);
void C_NSTextView_ChangeAttributes(void* ptr, void* sender);
void C_NSTextView_ChangeColor(void* ptr, void* sender);
void C_NSTextView_SetAlignment_Range(void* ptr, int alignment, NSRange _range);
void C_NSTextView_UseStandardKerning(void* ptr, void* sender);
void C_NSTextView_LowerBaseline(void* ptr, void* sender);
void C_NSTextView_RaiseBaseline(void* ptr, void* sender);
void C_NSTextView_TurnOffKerning(void* ptr, void* sender);
void C_NSTextView_LoosenKerning(void* ptr, void* sender);
void C_NSTextView_TightenKerning(void* ptr, void* sender);
void C_NSTextView_UseStandardLigatures(void* ptr, void* sender);
void C_NSTextView_TurnOffLigatures(void* ptr, void* sender);
void C_NSTextView_UseAllLigatures(void* ptr, void* sender);
void C_NSTextView_ClickedOnLink_AtIndex(void* ptr, void* link, unsigned int charIndex);
void C_NSTextView_PasteAsPlainText(void* ptr, void* sender);
void C_NSTextView_PasteAsRichText(void* ptr, void* sender);
void C_NSTextView_BreakUndoCoalescing(void* ptr);
void C_NSTextView_UpdateFontPanel(void* ptr);
void C_NSTextView_UpdateRuler(void* ptr);
void C_NSTextView_UpdateDragTypeRegistration(void* ptr);
NSRange C_NSTextView_SelectionRangeForProposedRange_Granularity(void* ptr, NSRange proposedCharRange, unsigned int granularity);
bool C_NSTextView_ShouldChangeTextInRange_ReplacementString(void* ptr, NSRange affectedCharRange, void* replacementString);
bool C_NSTextView_ShouldChangeTextInRanges_ReplacementStrings(void* ptr, Array affectedRanges, Array replacementStrings);
void C_NSTextView_DidChangeText(void* ptr);
NSRange C_NSTextView_SmartDeleteRangeForProposedRange(void* ptr, NSRange proposedCharRange);
void* C_NSTextView_SmartInsertAfterStringForString_ReplacingRange(void* ptr, void* pasteString, NSRange charRangeToReplace);
void* C_NSTextView_SmartInsertBeforeStringForString_ReplacingRange(void* ptr, void* pasteString, NSRange charRangeToReplace);
void C_NSTextView_ToggleSmartInsertDelete(void* ptr, void* sender);
void C_NSTextView_ToggleContinuousSpellChecking(void* ptr, void* sender);
void C_NSTextView_ToggleGrammarChecking(void* ptr, void* sender);
void C_NSTextView_SetSpellingState_Range(void* ptr, int value, NSRange charRange);
void C_NSTextView_OrderFrontSharingServicePicker(void* ptr, void* sender);
unsigned int C_NSTextView_DragOperationForDraggingInfo_Type(void* ptr, void* dragInfo, void* _type);
bool C_NSTextView_DragSelectionWithEvent_Offset_SlideBack(void* ptr, void* event, CGSize mouseOffset, bool slideBack);
void C_NSTextView_StartSpeaking(void* ptr, void* sender);
void C_NSTextView_StopSpeaking(void* ptr, void* sender);
void C_NSTextView_PerformFindPanelAction(void* ptr, void* sender);
void C_NSTextView_OrderFrontLinkPanel(void* ptr, void* sender);
void C_NSTextView_OrderFrontListPanel(void* ptr, void* sender);
void C_NSTextView_OrderFrontSpacingPanel(void* ptr, void* sender);
void C_NSTextView_OrderFrontTablePanel(void* ptr, void* sender);
void C_NSTextView_OrderFrontSubstitutionsPanel(void* ptr, void* sender);
void C_NSTextView_Complete(void* ptr, void* sender);
void C_NSTextView_InsertCompletion_ForPartialWordRange_Movement_IsFinal(void* ptr, void* word, NSRange charRange, int movement, bool flag);
void C_NSTextView_CheckTextInDocument(void* ptr, void* sender);
void C_NSTextView_CheckTextInSelection(void* ptr, void* sender);
void C_NSTextView_ToggleAutomaticDashSubstitution(void* ptr, void* sender);
void C_NSTextView_ToggleAutomaticDataDetection(void* ptr, void* sender);
void C_NSTextView_ToggleAutomaticSpellingCorrection(void* ptr, void* sender);
void C_NSTextView_ToggleAutomaticTextReplacement(void* ptr, void* sender);
void C_NSTextView_UpdateQuickLookPreviewPanel(void* ptr);
void C_NSTextView_ToggleQuickLookPreviewPanel(void* ptr, void* sender);
void C_NSTextView_ChangeLayoutOrientation(void* ptr, void* sender);
void C_NSTextView_SetLayoutOrientation(void* ptr, int orientation);
bool C_NSTextView_PerformValidatedReplacementInRange_WithAttributedString(void* ptr, NSRange _range, void* attributedString);
void C_NSTextView_ToggleAutomaticTextCompletion(void* ptr, void* sender);
void C_NSTextView_UpdateCandidates(void* ptr);
void C_NSTextView_UpdateTextTouchBarItems(void* ptr);
void C_NSTextView_UpdateTouchBarItemIdentifiers(void* ptr);
void* C_NSTextView_ScrollableDocumentContentTextView();
void* C_NSTextView_ScrollablePlainDocumentContentTextView();
void* C_NSTextView_ScrollableTextView();
void* C_NSTextView_TextContainer(void* ptr);
void C_NSTextView_SetTextContainer(void* ptr, void* value);
CGSize C_NSTextView_TextContainerInset(void* ptr);
void C_NSTextView_SetTextContainerInset(void* ptr, CGSize value);
CGPoint C_NSTextView_TextContainerOrigin(void* ptr);
void* C_NSTextView_LayoutManager(void* ptr);
void* C_NSTextView_TextStorage(void* ptr);
bool C_NSTextView_AllowsDocumentBackgroundColorChange(void* ptr);
void C_NSTextView_SetAllowsDocumentBackgroundColorChange(void* ptr, bool value);
bool C_NSTextView_ShouldDrawInsertionPoint(void* ptr);
Array C_NSTextView_AllowedInputSourceLocales(void* ptr);
void C_NSTextView_SetAllowedInputSourceLocales(void* ptr, Array value);
bool C_NSTextView_AllowsUndo(void* ptr);
void C_NSTextView_SetAllowsUndo(void* ptr, bool value);
void* C_NSTextView_DefaultParagraphStyle(void* ptr);
void C_NSTextView_SetDefaultParagraphStyle(void* ptr, void* value);
bool C_NSTextView_AllowsImageEditing(void* ptr);
void C_NSTextView_SetAllowsImageEditing(void* ptr, bool value);
bool C_NSTextView_IsAutomaticQuoteSubstitutionEnabled(void* ptr);
void C_NSTextView_SetAutomaticQuoteSubstitutionEnabled(void* ptr, bool value);
bool C_NSTextView_IsAutomaticLinkDetectionEnabled(void* ptr);
void C_NSTextView_SetAutomaticLinkDetectionEnabled(void* ptr, bool value);
bool C_NSTextView_DisplaysLinkToolTips(void* ptr);
void C_NSTextView_SetDisplaysLinkToolTips(void* ptr, bool value);
bool C_NSTextView_UsesRuler(void* ptr);
void C_NSTextView_SetUsesRuler(void* ptr, bool value);
void C_NSTextView_SetRulerVisible(void* ptr, bool value);
bool C_NSTextView_UsesInspectorBar(void* ptr);
void C_NSTextView_SetUsesInspectorBar(void* ptr, bool value);
Array C_NSTextView_SelectedRanges(void* ptr);
void C_NSTextView_SetSelectedRanges(void* ptr, Array value);
unsigned int C_NSTextView_SelectionAffinity(void* ptr);
unsigned int C_NSTextView_SelectionGranularity(void* ptr);
void C_NSTextView_SetSelectionGranularity(void* ptr, unsigned int value);
void* C_NSTextView_InsertionPointColor(void* ptr);
void C_NSTextView_SetInsertionPointColor(void* ptr, void* value);
Dictionary C_NSTextView_SelectedTextAttributes(void* ptr);
void C_NSTextView_SetSelectedTextAttributes(void* ptr, Dictionary value);
Dictionary C_NSTextView_MarkedTextAttributes(void* ptr);
void C_NSTextView_SetMarkedTextAttributes(void* ptr, Dictionary value);
Dictionary C_NSTextView_LinkTextAttributes(void* ptr);
void C_NSTextView_SetLinkTextAttributes(void* ptr, Dictionary value);
Array C_NSTextView_ReadablePasteboardTypes(void* ptr);
Array C_NSTextView_WritablePasteboardTypes(void* ptr);
Dictionary C_NSTextView_TypingAttributes(void* ptr);
void C_NSTextView_SetTypingAttributes(void* ptr, Dictionary value);
bool C_NSTextView_IsCoalescingUndo(void* ptr);
Array C_NSTextView_AcceptableDragTypes(void* ptr);
NSRange C_NSTextView_RangeForUserCharacterAttributeChange(void* ptr);
Array C_NSTextView_RangesForUserCharacterAttributeChange(void* ptr);
NSRange C_NSTextView_RangeForUserParagraphAttributeChange(void* ptr);
Array C_NSTextView_RangesForUserParagraphAttributeChange(void* ptr);
NSRange C_NSTextView_RangeForUserTextChange(void* ptr);
Array C_NSTextView_RangesForUserTextChange(void* ptr);
bool C_NSTextView_SmartInsertDeleteEnabled(void* ptr);
void C_NSTextView_SetSmartInsertDeleteEnabled(void* ptr, bool value);
bool C_NSTextView_IsContinuousSpellCheckingEnabled(void* ptr);
void C_NSTextView_SetContinuousSpellCheckingEnabled(void* ptr, bool value);
int C_NSTextView_SpellCheckerDocumentTag(void* ptr);
bool C_NSTextView_IsGrammarCheckingEnabled(void* ptr);
void C_NSTextView_SetGrammarCheckingEnabled(void* ptr, bool value);
bool C_NSTextView_AcceptsGlyphInfo(void* ptr);
void C_NSTextView_SetAcceptsGlyphInfo(void* ptr, bool value);
bool C_NSTextView_UsesFindPanel(void* ptr);
void C_NSTextView_SetUsesFindPanel(void* ptr, bool value);
NSRange C_NSTextView_RangeForUserCompletion(void* ptr);
bool C_NSTextView_IsAutomaticDashSubstitutionEnabled(void* ptr);
void C_NSTextView_SetAutomaticDashSubstitutionEnabled(void* ptr, bool value);
bool C_NSTextView_IsAutomaticDataDetectionEnabled(void* ptr);
void C_NSTextView_SetAutomaticDataDetectionEnabled(void* ptr, bool value);
bool C_NSTextView_IsAutomaticSpellingCorrectionEnabled(void* ptr);
void C_NSTextView_SetAutomaticSpellingCorrectionEnabled(void* ptr, bool value);
bool C_NSTextView_IsAutomaticTextReplacementEnabled(void* ptr);
void C_NSTextView_SetAutomaticTextReplacementEnabled(void* ptr, bool value);
bool C_NSTextView_UsesFindBar(void* ptr);
void C_NSTextView_SetUsesFindBar(void* ptr, bool value);
bool C_NSTextView_IsIncrementalSearchingEnabled(void* ptr);
void C_NSTextView_SetIncrementalSearchingEnabled(void* ptr, bool value);
bool C_NSTextView_AllowsCharacterPickerTouchBarItem(void* ptr);
void C_NSTextView_SetAllowsCharacterPickerTouchBarItem(void* ptr, bool value);
bool C_NSTextView_IsAutomaticTextCompletionEnabled(void* ptr);
void C_NSTextView_SetAutomaticTextCompletionEnabled(void* ptr, bool value);
bool C_NSTextView_UsesAdaptiveColorMappingForDarkAppearance(void* ptr);
void C_NSTextView_SetUsesAdaptiveColorMappingForDarkAppearance(void* ptr, bool value);
bool C_NSTextView_UsesRolloverButtonForSelection(void* ptr);
void C_NSTextView_SetUsesRolloverButtonForSelection(void* ptr, bool value);
bool C_NSTextView_TextView_StronglyReferencesTextStorage();

void* WrapTextViewDelegate(uintptr_t goID);
