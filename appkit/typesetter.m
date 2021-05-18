#import <Appkit/Appkit.h>
#import "typesetter.h"

void* C_Typesetter_Alloc() {
    return [NSTypesetter alloc];
}

void* C_NSTypesetter_Init(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSTypesetter* result_ = [nSTypesetter init];
    return result_;
}

void* C_NSTypesetter_Typesetter_SharedSystemTypesetterForBehavior(int behavior) {
    id result_ = [NSTypesetter sharedSystemTypesetterForBehavior:behavior];
    return result_;
}

double C_NSTypesetter_BaselineOffsetInLayoutManager_GlyphIndex(void* ptr, void* layoutMgr, unsigned int glyphIndex) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    CGFloat result_ = [nSTypesetter baselineOffsetInLayoutManager:(NSLayoutManager*)layoutMgr glyphIndex:glyphIndex];
    return result_;
}

void* C_NSTypesetter_SubstituteFontForFont(void* ptr, void* originalFont) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSFont* result_ = [nSTypesetter substituteFontForFont:(NSFont*)originalFont];
    return result_;
}

void* C_NSTypesetter_TextTabForGlyphLocation_WritingDirection_MaxLocation(void* ptr, double glyphLocation, int direction, double maxLocation) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSTextTab* result_ = [nSTypesetter textTabForGlyphLocation:glyphLocation writingDirection:direction maxLocation:maxLocation];
    return result_;
}

void C_NSTypesetter_SetParagraphGlyphRange_SeparatorGlyphRange(void* ptr, NSRange paragraphRange, NSRange paragraphSeparatorRange) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setParagraphGlyphRange:paragraphRange separatorGlyphRange:paragraphSeparatorRange];
}

double C_NSTypesetter_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(void* ptr, unsigned int glyphIndex, CGRect rect) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    CGFloat result_ = [nSTypesetter lineSpacingAfterGlyphAtIndex:glyphIndex withProposedLineFragmentRect:rect];
    return result_;
}

double C_NSTypesetter_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(void* ptr, unsigned int glyphIndex, CGRect rect) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    CGFloat result_ = [nSTypesetter paragraphSpacingAfterGlyphAtIndex:glyphIndex withProposedLineFragmentRect:rect];
    return result_;
}

double C_NSTypesetter_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(void* ptr, unsigned int glyphIndex, CGRect rect) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    CGFloat result_ = [nSTypesetter paragraphSpacingBeforeGlyphAtIndex:glyphIndex withProposedLineFragmentRect:rect];
    return result_;
}

void C_NSTypesetter_BeginParagraph(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter beginParagraph];
}

void C_NSTypesetter_EndParagraph(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter endParagraph];
}

void C_NSTypesetter_BeginLineWithGlyphAtIndex(void* ptr, unsigned int glyphIndex) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter beginLineWithGlyphAtIndex:glyphIndex];
}

void C_NSTypesetter_EndLineWithGlyphRange(void* ptr, NSRange lineGlyphRange) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter endLineWithGlyphRange:lineGlyphRange];
}

NSRange C_NSTypesetter_LayoutCharactersInRange_ForLayoutManager_MaximumNumberOfLineFragments(void* ptr, NSRange characterRange, void* layoutManager, unsigned int maxNumLines) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSRange result_ = [nSTypesetter layoutCharactersInRange:characterRange forLayoutManager:(NSLayoutManager*)layoutManager maximumNumberOfLineFragments:maxNumLines];
    return result_;
}

CGRect C_NSTypesetter_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(void* ptr, unsigned int glyphIndex, void* textContainer, CGRect proposedRect, CGPoint glyphPosition, unsigned int charIndex) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSRect result_ = [nSTypesetter boundingBoxForControlGlyphAtIndex:glyphIndex forTextContainer:(NSTextContainer*)textContainer proposedLineFragment:proposedRect glyphPosition:glyphPosition characterIndex:charIndex];
    return result_;
}

float C_NSTypesetter_HyphenationFactorForGlyphAtIndex(void* ptr, unsigned int glyphIndex) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    float result_ = [nSTypesetter hyphenationFactorForGlyphAtIndex:glyphIndex];
    return result_;
}

bool C_NSTypesetter_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(void* ptr, unsigned int charIndex) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    BOOL result_ = [nSTypesetter shouldBreakLineByHyphenatingBeforeCharacterAtIndex:charIndex];
    return result_;
}

bool C_NSTypesetter_ShouldBreakLineByWordBeforeCharacterAtIndex(void* ptr, unsigned int charIndex) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    BOOL result_ = [nSTypesetter shouldBreakLineByWordBeforeCharacterAtIndex:charIndex];
    return result_;
}

void C_NSTypesetter_SetHardInvalidation_ForGlyphRange(void* ptr, bool flag, NSRange glyphRange) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setHardInvalidation:flag forGlyphRange:glyphRange];
}

void C_NSTypesetter_SetAttachmentSize_ForGlyphRange(void* ptr, CGSize attachmentSize, NSRange glyphRange) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setAttachmentSize:attachmentSize forGlyphRange:glyphRange];
}

void C_NSTypesetter_SetDrawsOutsideLineFragment_ForGlyphRange(void* ptr, bool flag, NSRange glyphRange) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setDrawsOutsideLineFragment:flag forGlyphRange:glyphRange];
}

void C_NSTypesetter_SetLineFragmentRect_ForGlyphRange_UsedRect_BaselineOffset(void* ptr, CGRect fragmentRect, NSRange glyphRange, CGRect usedRect, double baselineOffset) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setLineFragmentRect:fragmentRect forGlyphRange:glyphRange usedRect:usedRect baselineOffset:baselineOffset];
}

void C_NSTypesetter_SetNotShownAttribute_ForGlyphRange(void* ptr, bool flag, NSRange glyphRange) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setNotShownAttribute:flag forGlyphRange:glyphRange];
}

void* C_NSTypesetter_SharedSystemTypesetter() {
    NSTypesetter* result_ = [NSTypesetter sharedSystemTypesetter];
    return result_;
}

int C_NSTypesetter_Typesetter_DefaultTypesetterBehavior() {
    NSTypesetterBehavior result_ = [NSTypesetter defaultTypesetterBehavior];
    return result_;
}

void* C_NSTypesetter_LayoutManager(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSLayoutManager* result_ = [nSTypesetter layoutManager];
    return result_;
}

bool C_NSTypesetter_UsesFontLeading(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    BOOL result_ = [nSTypesetter usesFontLeading];
    return result_;
}

void C_NSTypesetter_SetUsesFontLeading(void* ptr, bool value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setUsesFontLeading:value];
}

int C_NSTypesetter_TypesetterBehavior(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSTypesetterBehavior result_ = [nSTypesetter typesetterBehavior];
    return result_;
}

void C_NSTypesetter_SetTypesetterBehavior(void* ptr, int value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setTypesetterBehavior:value];
}

float C_NSTypesetter_HyphenationFactor(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    float result_ = [nSTypesetter hyphenationFactor];
    return result_;
}

void C_NSTypesetter_SetHyphenationFactor(void* ptr, float value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setHyphenationFactor:value];
}

void* C_NSTypesetter_CurrentTextContainer(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSTextContainer* result_ = [nSTypesetter currentTextContainer];
    return result_;
}

Array C_NSTypesetter_TextContainers(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSArray* result_ = [nSTypesetter textContainers];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

double C_NSTypesetter_LineFragmentPadding(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    CGFloat result_ = [nSTypesetter lineFragmentPadding];
    return result_;
}

void C_NSTypesetter_SetLineFragmentPadding(void* ptr, double value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setLineFragmentPadding:value];
}

bool C_NSTypesetter_BidiProcessingEnabled(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    BOOL result_ = [nSTypesetter bidiProcessingEnabled];
    return result_;
}

void C_NSTypesetter_SetBidiProcessingEnabled(void* ptr, bool value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setBidiProcessingEnabled:value];
}

void* C_NSTypesetter_CurrentParagraphStyle(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSParagraphStyle* result_ = [nSTypesetter currentParagraphStyle];
    return result_;
}

void* C_NSTypesetter_AttributedString(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSAttributedString* result_ = [nSTypesetter attributedString];
    return result_;
}

void C_NSTypesetter_SetAttributedString(void* ptr, void* value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setAttributedString:(NSAttributedString*)value];
}

NSRange C_NSTypesetter_ParagraphGlyphRange(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSRange result_ = [nSTypesetter paragraphGlyphRange];
    return result_;
}

NSRange C_NSTypesetter_ParagraphSeparatorGlyphRange(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSRange result_ = [nSTypesetter paragraphSeparatorGlyphRange];
    return result_;
}

NSRange C_NSTypesetter_ParagraphCharacterRange(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSRange result_ = [nSTypesetter paragraphCharacterRange];
    return result_;
}

NSRange C_NSTypesetter_ParagraphSeparatorCharacterRange(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSRange result_ = [nSTypesetter paragraphSeparatorCharacterRange];
    return result_;
}
