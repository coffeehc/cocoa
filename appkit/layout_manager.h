#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_LayoutManager_Alloc();

void* C_NSLayoutManager_Init(void* ptr);
void* C_NSLayoutManager_InitWithCoder(void* ptr, void* coder);
void C_NSLayoutManager_ReplaceTextStorage(void* ptr, void* newTextStorage);
void C_NSLayoutManager_AddTextContainer(void* ptr, void* container);
void C_NSLayoutManager_InsertTextContainer_AtIndex(void* ptr, void* container, unsigned int index);
void C_NSLayoutManager_RemoveTextContainerAtIndex(void* ptr, unsigned int index);
void C_NSLayoutManager_SetTextContainer_ForGlyphRange(void* ptr, void* container, NSRange glyphRange);
void C_NSLayoutManager_TextContainerChangedGeometry(void* ptr, void* container);
void C_NSLayoutManager_TextContainerChangedTextView(void* ptr, void* container);
CGRect C_NSLayoutManager_UsedRectForTextContainer(void* ptr, void* container);
void C_NSLayoutManager_InvalidateDisplayForCharacterRange(void* ptr, NSRange charRange);
void C_NSLayoutManager_InvalidateDisplayForGlyphRange(void* ptr, NSRange glyphRange);
void C_NSLayoutManager_ProcessEditingForTextStorage_Edited_Range_ChangeInLength_InvalidatedRange(void* ptr, void* textStorage, unsigned int editMask, NSRange newCharRange, int delta, NSRange invalidatedCharRange);
void C_NSLayoutManager_EnsureGlyphsForCharacterRange(void* ptr, NSRange charRange);
void C_NSLayoutManager_EnsureGlyphsForGlyphRange(void* ptr, NSRange glyphRange);
void C_NSLayoutManager_EnsureLayoutForBoundingRect_InTextContainer(void* ptr, CGRect bounds, void* container);
void C_NSLayoutManager_EnsureLayoutForCharacterRange(void* ptr, NSRange charRange);
void C_NSLayoutManager_EnsureLayoutForGlyphRange(void* ptr, NSRange glyphRange);
void C_NSLayoutManager_EnsureLayoutForTextContainer(void* ptr, void* container);
unsigned int C_NSLayoutManager_CharacterIndexForGlyphAtIndex(void* ptr, unsigned int glyphIndex);
unsigned int C_NSLayoutManager_GlyphIndexForCharacterAtIndex(void* ptr, unsigned int charIndex);
bool C_NSLayoutManager_IsValidGlyphIndex(void* ptr, unsigned int glyphIndex);
int C_NSLayoutManager_PropertyForGlyphAtIndex(void* ptr, unsigned int glyphIndex);
void C_NSLayoutManager_SetAttachmentSize_ForGlyphRange(void* ptr, CGSize attachmentSize, NSRange glyphRange);
void C_NSLayoutManager_SetDrawsOutsideLineFragment_ForGlyphAtIndex(void* ptr, bool flag, unsigned int glyphIndex);
void C_NSLayoutManager_SetExtraLineFragmentRect_UsedRect_TextContainer(void* ptr, CGRect fragmentRect, CGRect usedRect, void* container);
void C_NSLayoutManager_SetLineFragmentRect_ForGlyphRange_UsedRect(void* ptr, CGRect fragmentRect, NSRange glyphRange, CGRect usedRect);
void C_NSLayoutManager_SetLocation_ForStartOfGlyphRange(void* ptr, CGPoint location, NSRange glyphRange);
void C_NSLayoutManager_SetNotShownAttribute_ForGlyphAtIndex(void* ptr, bool flag, unsigned int glyphIndex);
CGSize C_NSLayoutManager_AttachmentSizeForGlyphAtIndex(void* ptr, unsigned int glyphIndex);
bool C_NSLayoutManager_DrawsOutsideLineFragmentForGlyphAtIndex(void* ptr, unsigned int glyphIndex);
unsigned int C_NSLayoutManager_FirstUnlaidCharacterIndex(void* ptr);
unsigned int C_NSLayoutManager_FirstUnlaidGlyphIndex(void* ptr);
CGPoint C_NSLayoutManager_LocationForGlyphAtIndex(void* ptr, unsigned int glyphIndex);
bool C_NSLayoutManager_NotShownAttributeForGlyphAtIndex(void* ptr, unsigned int glyphIndex);
NSRange C_NSLayoutManager_TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(void* ptr, unsigned int glyphIndex);
CGRect C_NSLayoutManager_BoundingRectForGlyphRange_InTextContainer(void* ptr, NSRange glyphRange, void* container);
double C_NSLayoutManager_FractionOfDistanceThroughGlyphForPoint_InTextContainer(void* ptr, CGPoint point, void* container);
unsigned int C_NSLayoutManager_GlyphIndexForPoint_InTextContainer(void* ptr, CGPoint point, void* container);
NSRange C_NSLayoutManager_GlyphRangeForBoundingRect_InTextContainer(void* ptr, CGRect bounds, void* container);
NSRange C_NSLayoutManager_GlyphRangeForBoundingRectWithoutAdditionalLayout_InTextContainer(void* ptr, CGRect bounds, void* container);
NSRange C_NSLayoutManager_GlyphRangeForTextContainer(void* ptr, void* container);
NSRange C_NSLayoutManager_RangeOfNominallySpacedGlyphsContainingIndex(void* ptr, unsigned int glyphIndex);
void C_NSLayoutManager_DrawBackgroundForGlyphRange_AtPoint(void* ptr, NSRange glyphsToShow, CGPoint origin);
void C_NSLayoutManager_DrawGlyphsForGlyphRange_AtPoint(void* ptr, NSRange glyphsToShow, CGPoint origin);
void C_NSLayoutManager_DrawStrikethroughForGlyphRange_StrikethroughType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(void* ptr, NSRange glyphRange, int strikethroughVal, double baselineOffset, CGRect lineRect, NSRange lineGlyphRange, CGPoint containerOrigin);
void C_NSLayoutManager_DrawUnderlineForGlyphRange_UnderlineType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(void* ptr, NSRange glyphRange, int underlineVal, double baselineOffset, CGRect lineRect, NSRange lineGlyphRange, CGPoint containerOrigin);
void C_NSLayoutManager_StrikethroughGlyphRange_StrikethroughType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(void* ptr, NSRange glyphRange, int strikethroughVal, CGRect lineRect, NSRange lineGlyphRange, CGPoint containerOrigin);
void C_NSLayoutManager_UnderlineGlyphRange_UnderlineType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(void* ptr, NSRange glyphRange, int underlineVal, CGRect lineRect, NSRange lineGlyphRange, CGPoint containerOrigin);
void C_NSLayoutManager_SetLayoutRect_ForTextBlock_GlyphRange(void* ptr, CGRect rect, void* block, NSRange glyphRange);
CGRect C_NSLayoutManager_LayoutRectForTextBlock_GlyphRange(void* ptr, void* block, NSRange glyphRange);
void C_NSLayoutManager_SetBoundsRect_ForTextBlock_GlyphRange(void* ptr, CGRect rect, void* block, NSRange glyphRange);
CGRect C_NSLayoutManager_BoundsRectForTextBlock_GlyphRange(void* ptr, void* block, NSRange glyphRange);
void C_NSLayoutManager_ShowAttachmentCell_InRect_CharacterIndex(void* ptr, void* cell, CGRect rect, unsigned int attachmentIndex);
void* C_NSLayoutManager_RulerAccessoryViewForTextView_ParagraphStyle_Ruler_Enabled(void* ptr, void* view, void* style, void* ruler, bool isEnabled);
Array C_NSLayoutManager_RulerMarkersForTextView_ParagraphStyle_Ruler(void* ptr, void* view, void* style, void* ruler);
bool C_NSLayoutManager_LayoutManagerOwnsFirstResponderInWindow(void* ptr, void* window);
double C_NSLayoutManager_DefaultLineHeightForFont(void* ptr, void* theFont);
double C_NSLayoutManager_DefaultBaselineOffsetForFont(void* ptr, void* theFont);
void C_NSLayoutManager_AddTemporaryAttribute_Value_ForCharacterRange(void* ptr, void* attrName, void* value, NSRange charRange);
void C_NSLayoutManager_RemoveTemporaryAttribute_ForCharacterRange(void* ptr, void* attrName, NSRange charRange);
void* C_NSLayoutManager_Delegate(void* ptr);
void C_NSLayoutManager_SetDelegate(void* ptr, void* value);
void* C_NSLayoutManager_TextStorage(void* ptr);
void C_NSLayoutManager_SetTextStorage(void* ptr, void* value);
bool C_NSLayoutManager_AllowsNonContiguousLayout(void* ptr);
void C_NSLayoutManager_SetAllowsNonContiguousLayout(void* ptr, bool value);
bool C_NSLayoutManager_HasNonContiguousLayout(void* ptr);
bool C_NSLayoutManager_ShowsInvisibleCharacters(void* ptr);
void C_NSLayoutManager_SetShowsInvisibleCharacters(void* ptr, bool value);
bool C_NSLayoutManager_ShowsControlCharacters(void* ptr);
void C_NSLayoutManager_SetShowsControlCharacters(void* ptr, bool value);
bool C_NSLayoutManager_UsesFontLeading(void* ptr);
void C_NSLayoutManager_SetUsesFontLeading(void* ptr, bool value);
bool C_NSLayoutManager_BackgroundLayoutEnabled(void* ptr);
void C_NSLayoutManager_SetBackgroundLayoutEnabled(void* ptr, bool value);
bool C_NSLayoutManager_LimitsLayoutForSuspiciousContents(void* ptr);
void C_NSLayoutManager_SetLimitsLayoutForSuspiciousContents(void* ptr, bool value);
bool C_NSLayoutManager_UsesDefaultHyphenation(void* ptr);
void C_NSLayoutManager_SetUsesDefaultHyphenation(void* ptr, bool value);
Array C_NSLayoutManager_TextContainers(void* ptr);
void* C_NSLayoutManager_GlyphGenerator(void* ptr);
void C_NSLayoutManager_SetGlyphGenerator(void* ptr, void* value);
unsigned int C_NSLayoutManager_NumberOfGlyphs(void* ptr);
CGRect C_NSLayoutManager_ExtraLineFragmentRect(void* ptr);
void* C_NSLayoutManager_ExtraLineFragmentTextContainer(void* ptr);
CGRect C_NSLayoutManager_ExtraLineFragmentUsedRect(void* ptr);
unsigned int C_NSLayoutManager_DefaultAttachmentScaling(void* ptr);
void C_NSLayoutManager_SetDefaultAttachmentScaling(void* ptr, unsigned int value);
void* C_NSLayoutManager_FirstTextView(void* ptr);
void* C_NSLayoutManager_TextViewForBeginningOfSelection(void* ptr);
void* C_NSLayoutManager_Typesetter(void* ptr);
void C_NSLayoutManager_SetTypesetter(void* ptr, void* value);
int C_NSLayoutManager_TypesetterBehavior(void* ptr);
void C_NSLayoutManager_SetTypesetterBehavior(void* ptr, int value);
