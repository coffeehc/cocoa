#import "layout_manager.h"
#import <AppKit/NSLayoutManager.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_LayoutManager_Alloc() {
    return [NSLayoutManager alloc];
}

void* C_NSLayoutManager_Init(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSLayoutManager* result_ = [nSLayoutManager init];
    return result_;
}

void* C_NSLayoutManager_InitWithCoder(void* ptr, void* coder) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSLayoutManager* result_ = [nSLayoutManager initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSLayoutManager_AllocLayoutManager() {
    NSLayoutManager* result_ = [NSLayoutManager alloc];
    return result_;
}

void* C_NSLayoutManager_NewLayoutManager() {
    NSLayoutManager* result_ = [NSLayoutManager new];
    return result_;
}

void* C_NSLayoutManager_Autorelease(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSLayoutManager* result_ = [nSLayoutManager autorelease];
    return result_;
}

void* C_NSLayoutManager_Retain(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSLayoutManager* result_ = [nSLayoutManager retain];
    return result_;
}

void C_NSLayoutManager_ReplaceTextStorage(void* ptr, void* newTextStorage) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager replaceTextStorage:(NSTextStorage*)newTextStorage];
}

void C_NSLayoutManager_AddTextContainer(void* ptr, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager addTextContainer:(NSTextContainer*)container];
}

void C_NSLayoutManager_InsertTextContainer_AtIndex(void* ptr, void* container, unsigned int index) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager insertTextContainer:(NSTextContainer*)container atIndex:index];
}

void C_NSLayoutManager_RemoveTextContainerAtIndex(void* ptr, unsigned int index) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager removeTextContainerAtIndex:index];
}

void C_NSLayoutManager_SetTextContainer_ForGlyphRange(void* ptr, void* container, NSRange glyphRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setTextContainer:(NSTextContainer*)container forGlyphRange:glyphRange];
}

void C_NSLayoutManager_TextContainerChangedGeometry(void* ptr, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager textContainerChangedGeometry:(NSTextContainer*)container];
}

void C_NSLayoutManager_TextContainerChangedTextView(void* ptr, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager textContainerChangedTextView:(NSTextContainer*)container];
}

CGRect C_NSLayoutManager_UsedRectForTextContainer(void* ptr, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSRect result_ = [nSLayoutManager usedRectForTextContainer:(NSTextContainer*)container];
    return result_;
}

void C_NSLayoutManager_InvalidateDisplayForCharacterRange(void* ptr, NSRange charRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager invalidateDisplayForCharacterRange:charRange];
}

void C_NSLayoutManager_InvalidateDisplayForGlyphRange(void* ptr, NSRange glyphRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager invalidateDisplayForGlyphRange:glyphRange];
}

void C_NSLayoutManager_ProcessEditingForTextStorage_Edited_Range_ChangeInLength_InvalidatedRange(void* ptr, void* textStorage, unsigned int editMask, NSRange newCharRange, int delta, NSRange invalidatedCharRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager processEditingForTextStorage:(NSTextStorage*)textStorage edited:editMask range:newCharRange changeInLength:delta invalidatedRange:invalidatedCharRange];
}

void C_NSLayoutManager_EnsureGlyphsForCharacterRange(void* ptr, NSRange charRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager ensureGlyphsForCharacterRange:charRange];
}

void C_NSLayoutManager_EnsureGlyphsForGlyphRange(void* ptr, NSRange glyphRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager ensureGlyphsForGlyphRange:glyphRange];
}

void C_NSLayoutManager_EnsureLayoutForBoundingRect_InTextContainer(void* ptr, CGRect bounds, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager ensureLayoutForBoundingRect:bounds inTextContainer:(NSTextContainer*)container];
}

void C_NSLayoutManager_EnsureLayoutForCharacterRange(void* ptr, NSRange charRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager ensureLayoutForCharacterRange:charRange];
}

void C_NSLayoutManager_EnsureLayoutForGlyphRange(void* ptr, NSRange glyphRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager ensureLayoutForGlyphRange:glyphRange];
}

void C_NSLayoutManager_EnsureLayoutForTextContainer(void* ptr, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager ensureLayoutForTextContainer:(NSTextContainer*)container];
}

unsigned int C_NSLayoutManager_CharacterIndexForGlyphAtIndex(void* ptr, unsigned int glyphIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSUInteger result_ = [nSLayoutManager characterIndexForGlyphAtIndex:glyphIndex];
    return result_;
}

unsigned int C_NSLayoutManager_GlyphIndexForCharacterAtIndex(void* ptr, unsigned int charIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSUInteger result_ = [nSLayoutManager glyphIndexForCharacterAtIndex:charIndex];
    return result_;
}

bool C_NSLayoutManager_IsValidGlyphIndex(void* ptr, unsigned int glyphIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager isValidGlyphIndex:glyphIndex];
    return result_;
}

int C_NSLayoutManager_PropertyForGlyphAtIndex(void* ptr, unsigned int glyphIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSGlyphProperty result_ = [nSLayoutManager propertyForGlyphAtIndex:glyphIndex];
    return result_;
}

void C_NSLayoutManager_SetAttachmentSize_ForGlyphRange(void* ptr, CGSize attachmentSize, NSRange glyphRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setAttachmentSize:attachmentSize forGlyphRange:glyphRange];
}

void C_NSLayoutManager_SetDrawsOutsideLineFragment_ForGlyphAtIndex(void* ptr, bool flag, unsigned int glyphIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setDrawsOutsideLineFragment:flag forGlyphAtIndex:glyphIndex];
}

void C_NSLayoutManager_SetExtraLineFragmentRect_UsedRect_TextContainer(void* ptr, CGRect fragmentRect, CGRect usedRect, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setExtraLineFragmentRect:fragmentRect usedRect:usedRect textContainer:(NSTextContainer*)container];
}

void C_NSLayoutManager_SetLineFragmentRect_ForGlyphRange_UsedRect(void* ptr, CGRect fragmentRect, NSRange glyphRange, CGRect usedRect) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setLineFragmentRect:fragmentRect forGlyphRange:glyphRange usedRect:usedRect];
}

void C_NSLayoutManager_SetLocation_ForStartOfGlyphRange(void* ptr, CGPoint location, NSRange glyphRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setLocation:location forStartOfGlyphRange:glyphRange];
}

void C_NSLayoutManager_SetNotShownAttribute_ForGlyphAtIndex(void* ptr, bool flag, unsigned int glyphIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setNotShownAttribute:flag forGlyphAtIndex:glyphIndex];
}

CGSize C_NSLayoutManager_AttachmentSizeForGlyphAtIndex(void* ptr, unsigned int glyphIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSSize result_ = [nSLayoutManager attachmentSizeForGlyphAtIndex:glyphIndex];
    return result_;
}

bool C_NSLayoutManager_DrawsOutsideLineFragmentForGlyphAtIndex(void* ptr, unsigned int glyphIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager drawsOutsideLineFragmentForGlyphAtIndex:glyphIndex];
    return result_;
}

unsigned int C_NSLayoutManager_FirstUnlaidCharacterIndex(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSUInteger result_ = [nSLayoutManager firstUnlaidCharacterIndex];
    return result_;
}

unsigned int C_NSLayoutManager_FirstUnlaidGlyphIndex(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSUInteger result_ = [nSLayoutManager firstUnlaidGlyphIndex];
    return result_;
}

CGPoint C_NSLayoutManager_LocationForGlyphAtIndex(void* ptr, unsigned int glyphIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSPoint result_ = [nSLayoutManager locationForGlyphAtIndex:glyphIndex];
    return result_;
}

bool C_NSLayoutManager_NotShownAttributeForGlyphAtIndex(void* ptr, unsigned int glyphIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager notShownAttributeForGlyphAtIndex:glyphIndex];
    return result_;
}

NSRange C_NSLayoutManager_TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(void* ptr, unsigned int glyphIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSRange result_ = [nSLayoutManager truncatedGlyphRangeInLineFragmentForGlyphAtIndex:glyphIndex];
    return result_;
}

CGRect C_NSLayoutManager_BoundingRectForGlyphRange_InTextContainer(void* ptr, NSRange glyphRange, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSRect result_ = [nSLayoutManager boundingRectForGlyphRange:glyphRange inTextContainer:(NSTextContainer*)container];
    return result_;
}

double C_NSLayoutManager_FractionOfDistanceThroughGlyphForPoint_InTextContainer(void* ptr, CGPoint point, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    CGFloat result_ = [nSLayoutManager fractionOfDistanceThroughGlyphForPoint:point inTextContainer:(NSTextContainer*)container];
    return result_;
}

unsigned int C_NSLayoutManager_GlyphIndexForPoint_InTextContainer(void* ptr, CGPoint point, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSUInteger result_ = [nSLayoutManager glyphIndexForPoint:point inTextContainer:(NSTextContainer*)container];
    return result_;
}

NSRange C_NSLayoutManager_GlyphRangeForBoundingRect_InTextContainer(void* ptr, CGRect bounds, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSRange result_ = [nSLayoutManager glyphRangeForBoundingRect:bounds inTextContainer:(NSTextContainer*)container];
    return result_;
}

NSRange C_NSLayoutManager_GlyphRangeForBoundingRectWithoutAdditionalLayout_InTextContainer(void* ptr, CGRect bounds, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSRange result_ = [nSLayoutManager glyphRangeForBoundingRectWithoutAdditionalLayout:bounds inTextContainer:(NSTextContainer*)container];
    return result_;
}

NSRange C_NSLayoutManager_GlyphRangeForTextContainer(void* ptr, void* container) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSRange result_ = [nSLayoutManager glyphRangeForTextContainer:(NSTextContainer*)container];
    return result_;
}

NSRange C_NSLayoutManager_RangeOfNominallySpacedGlyphsContainingIndex(void* ptr, unsigned int glyphIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSRange result_ = [nSLayoutManager rangeOfNominallySpacedGlyphsContainingIndex:glyphIndex];
    return result_;
}

void C_NSLayoutManager_DrawBackgroundForGlyphRange_AtPoint(void* ptr, NSRange glyphsToShow, CGPoint origin) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager drawBackgroundForGlyphRange:glyphsToShow atPoint:origin];
}

void C_NSLayoutManager_DrawGlyphsForGlyphRange_AtPoint(void* ptr, NSRange glyphsToShow, CGPoint origin) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager drawGlyphsForGlyphRange:glyphsToShow atPoint:origin];
}

void C_NSLayoutManager_DrawStrikethroughForGlyphRange_StrikethroughType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(void* ptr, NSRange glyphRange, int strikethroughVal, double baselineOffset, CGRect lineRect, NSRange lineGlyphRange, CGPoint containerOrigin) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager drawStrikethroughForGlyphRange:glyphRange strikethroughType:strikethroughVal baselineOffset:baselineOffset lineFragmentRect:lineRect lineFragmentGlyphRange:lineGlyphRange containerOrigin:containerOrigin];
}

void C_NSLayoutManager_DrawUnderlineForGlyphRange_UnderlineType_BaselineOffset_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(void* ptr, NSRange glyphRange, int underlineVal, double baselineOffset, CGRect lineRect, NSRange lineGlyphRange, CGPoint containerOrigin) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager drawUnderlineForGlyphRange:glyphRange underlineType:underlineVal baselineOffset:baselineOffset lineFragmentRect:lineRect lineFragmentGlyphRange:lineGlyphRange containerOrigin:containerOrigin];
}

void C_NSLayoutManager_StrikethroughGlyphRange_StrikethroughType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(void* ptr, NSRange glyphRange, int strikethroughVal, CGRect lineRect, NSRange lineGlyphRange, CGPoint containerOrigin) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager strikethroughGlyphRange:glyphRange strikethroughType:strikethroughVal lineFragmentRect:lineRect lineFragmentGlyphRange:lineGlyphRange containerOrigin:containerOrigin];
}

void C_NSLayoutManager_UnderlineGlyphRange_UnderlineType_LineFragmentRect_LineFragmentGlyphRange_ContainerOrigin(void* ptr, NSRange glyphRange, int underlineVal, CGRect lineRect, NSRange lineGlyphRange, CGPoint containerOrigin) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager underlineGlyphRange:glyphRange underlineType:underlineVal lineFragmentRect:lineRect lineFragmentGlyphRange:lineGlyphRange containerOrigin:containerOrigin];
}

void C_NSLayoutManager_SetLayoutRect_ForTextBlock_GlyphRange(void* ptr, CGRect rect, void* block, NSRange glyphRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setLayoutRect:rect forTextBlock:(NSTextBlock*)block glyphRange:glyphRange];
}

CGRect C_NSLayoutManager_LayoutRectForTextBlock_GlyphRange(void* ptr, void* block, NSRange glyphRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSRect result_ = [nSLayoutManager layoutRectForTextBlock:(NSTextBlock*)block glyphRange:glyphRange];
    return result_;
}

void C_NSLayoutManager_SetBoundsRect_ForTextBlock_GlyphRange(void* ptr, CGRect rect, void* block, NSRange glyphRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setBoundsRect:rect forTextBlock:(NSTextBlock*)block glyphRange:glyphRange];
}

CGRect C_NSLayoutManager_BoundsRectForTextBlock_GlyphRange(void* ptr, void* block, NSRange glyphRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSRect result_ = [nSLayoutManager boundsRectForTextBlock:(NSTextBlock*)block glyphRange:glyphRange];
    return result_;
}

void C_NSLayoutManager_ShowAttachmentCell_InRect_CharacterIndex(void* ptr, void* cell, CGRect rect, unsigned int attachmentIndex) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager showAttachmentCell:(NSCell*)cell inRect:rect characterIndex:attachmentIndex];
}

void* C_NSLayoutManager_RulerAccessoryViewForTextView_ParagraphStyle_Ruler_Enabled(void* ptr, void* view, void* style, void* ruler, bool isEnabled) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSView* result_ = [nSLayoutManager rulerAccessoryViewForTextView:(NSTextView*)view paragraphStyle:(NSParagraphStyle*)style ruler:(NSRulerView*)ruler enabled:isEnabled];
    return result_;
}

Array C_NSLayoutManager_RulerMarkersForTextView_ParagraphStyle_Ruler(void* ptr, void* view, void* style, void* ruler) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSArray* result_ = [nSLayoutManager rulerMarkersForTextView:(NSTextView*)view paragraphStyle:(NSParagraphStyle*)style ruler:(NSRulerView*)ruler];
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

bool C_NSLayoutManager_LayoutManagerOwnsFirstResponderInWindow(void* ptr, void* window) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager layoutManagerOwnsFirstResponderInWindow:(NSWindow*)window];
    return result_;
}

double C_NSLayoutManager_DefaultLineHeightForFont(void* ptr, void* theFont) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    CGFloat result_ = [nSLayoutManager defaultLineHeightForFont:(NSFont*)theFont];
    return result_;
}

double C_NSLayoutManager_DefaultBaselineOffsetForFont(void* ptr, void* theFont) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    CGFloat result_ = [nSLayoutManager defaultBaselineOffsetForFont:(NSFont*)theFont];
    return result_;
}

void C_NSLayoutManager_AddTemporaryAttributes_ForCharacterRange(void* ptr, Dictionary attrs, NSRange charRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSMutableDictionary* objcAttrs = [NSMutableDictionary dictionaryWithCapacity:attrs.len];
    if (attrs.len > 0) {
    	void** attrsKeyData = (void**)attrs.key_data;
    	void** attrsValueData = (void**)attrs.value_data;
    	for (int i = 0; i < attrs.len; i++) {
    		void* kp = attrsKeyData[i];
    		void* vp = attrsValueData[i];
    		[objcAttrs setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSLayoutManager addTemporaryAttributes:objcAttrs forCharacterRange:charRange];
}

void C_NSLayoutManager_AddTemporaryAttribute_Value_ForCharacterRange(void* ptr, void* attrName, void* value, NSRange charRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager addTemporaryAttribute:(NSString*)attrName value:(id)value forCharacterRange:charRange];
}

void C_NSLayoutManager_SetTemporaryAttributes_ForCharacterRange(void* ptr, Dictionary attrs, NSRange charRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSMutableDictionary* objcAttrs = [NSMutableDictionary dictionaryWithCapacity:attrs.len];
    if (attrs.len > 0) {
    	void** attrsKeyData = (void**)attrs.key_data;
    	void** attrsValueData = (void**)attrs.value_data;
    	for (int i = 0; i < attrs.len; i++) {
    		void* kp = attrsKeyData[i];
    		void* vp = attrsValueData[i];
    		[objcAttrs setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSLayoutManager setTemporaryAttributes:objcAttrs forCharacterRange:charRange];
}

void C_NSLayoutManager_RemoveTemporaryAttribute_ForCharacterRange(void* ptr, void* attrName, NSRange charRange) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager removeTemporaryAttribute:(NSString*)attrName forCharacterRange:charRange];
}

void* C_NSLayoutManager_Delegate(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    id result_ = [nSLayoutManager delegate];
    return result_;
}

void C_NSLayoutManager_SetDelegate(void* ptr, void* value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setDelegate:(id)value];
}

void* C_NSLayoutManager_TextStorage(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSTextStorage* result_ = [nSLayoutManager textStorage];
    return result_;
}

void C_NSLayoutManager_SetTextStorage(void* ptr, void* value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setTextStorage:(NSTextStorage*)value];
}

bool C_NSLayoutManager_AllowsNonContiguousLayout(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager allowsNonContiguousLayout];
    return result_;
}

void C_NSLayoutManager_SetAllowsNonContiguousLayout(void* ptr, bool value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setAllowsNonContiguousLayout:value];
}

bool C_NSLayoutManager_HasNonContiguousLayout(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager hasNonContiguousLayout];
    return result_;
}

bool C_NSLayoutManager_ShowsInvisibleCharacters(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager showsInvisibleCharacters];
    return result_;
}

void C_NSLayoutManager_SetShowsInvisibleCharacters(void* ptr, bool value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setShowsInvisibleCharacters:value];
}

bool C_NSLayoutManager_ShowsControlCharacters(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager showsControlCharacters];
    return result_;
}

void C_NSLayoutManager_SetShowsControlCharacters(void* ptr, bool value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setShowsControlCharacters:value];
}

bool C_NSLayoutManager_UsesFontLeading(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager usesFontLeading];
    return result_;
}

void C_NSLayoutManager_SetUsesFontLeading(void* ptr, bool value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setUsesFontLeading:value];
}

bool C_NSLayoutManager_BackgroundLayoutEnabled(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager backgroundLayoutEnabled];
    return result_;
}

void C_NSLayoutManager_SetBackgroundLayoutEnabled(void* ptr, bool value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setBackgroundLayoutEnabled:value];
}

bool C_NSLayoutManager_LimitsLayoutForSuspiciousContents(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager limitsLayoutForSuspiciousContents];
    return result_;
}

void C_NSLayoutManager_SetLimitsLayoutForSuspiciousContents(void* ptr, bool value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setLimitsLayoutForSuspiciousContents:value];
}

bool C_NSLayoutManager_UsesDefaultHyphenation(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    BOOL result_ = [nSLayoutManager usesDefaultHyphenation];
    return result_;
}

void C_NSLayoutManager_SetUsesDefaultHyphenation(void* ptr, bool value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setUsesDefaultHyphenation:value];
}

Array C_NSLayoutManager_TextContainers(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSArray* result_ = [nSLayoutManager textContainers];
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

void* C_NSLayoutManager_GlyphGenerator(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSGlyphGenerator* result_ = [nSLayoutManager glyphGenerator];
    return result_;
}

void C_NSLayoutManager_SetGlyphGenerator(void* ptr, void* value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setGlyphGenerator:(NSGlyphGenerator*)value];
}

unsigned int C_NSLayoutManager_NumberOfGlyphs(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSUInteger result_ = [nSLayoutManager numberOfGlyphs];
    return result_;
}

CGRect C_NSLayoutManager_ExtraLineFragmentRect(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSRect result_ = [nSLayoutManager extraLineFragmentRect];
    return result_;
}

void* C_NSLayoutManager_ExtraLineFragmentTextContainer(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSTextContainer* result_ = [nSLayoutManager extraLineFragmentTextContainer];
    return result_;
}

CGRect C_NSLayoutManager_ExtraLineFragmentUsedRect(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSRect result_ = [nSLayoutManager extraLineFragmentUsedRect];
    return result_;
}

unsigned int C_NSLayoutManager_DefaultAttachmentScaling(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSImageScaling result_ = [nSLayoutManager defaultAttachmentScaling];
    return result_;
}

void C_NSLayoutManager_SetDefaultAttachmentScaling(void* ptr, unsigned int value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setDefaultAttachmentScaling:value];
}

void* C_NSLayoutManager_FirstTextView(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSTextView* result_ = [nSLayoutManager firstTextView];
    return result_;
}

void* C_NSLayoutManager_TextViewForBeginningOfSelection(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSTextView* result_ = [nSLayoutManager textViewForBeginningOfSelection];
    return result_;
}

void* C_NSLayoutManager_Typesetter(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSTypesetter* result_ = [nSLayoutManager typesetter];
    return result_;
}

void C_NSLayoutManager_SetTypesetter(void* ptr, void* value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setTypesetter:(NSTypesetter*)value];
}

int C_NSLayoutManager_TypesetterBehavior(void* ptr) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    NSTypesetterBehavior result_ = [nSLayoutManager typesetterBehavior];
    return result_;
}

void C_NSLayoutManager_SetTypesetterBehavior(void* ptr, int value) {
    NSLayoutManager* nSLayoutManager = (NSLayoutManager*)ptr;
    [nSLayoutManager setTypesetterBehavior:value];
}
