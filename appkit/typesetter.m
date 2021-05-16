#import <Appkit/Appkit.h>
#import "typesetter.h"

void* C_Typesetter_Alloc() {
    return [NSTypesetter alloc];
}

void* C_NSTypesetter_Init(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSTypesetter* result = [nSTypesetter init];
    return result;
}

void* C_NSTypesetter_Typesetter_SharedSystemTypesetterForBehavior(int behavior) {
    id result = [NSTypesetter sharedSystemTypesetterForBehavior:behavior];
    return result;
}

double C_NSTypesetter_BaselineOffsetInLayoutManager_GlyphIndex(void* ptr, void* layoutMgr, unsigned int glyphIndex) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    CGFloat result = [nSTypesetter baselineOffsetInLayoutManager:(NSLayoutManager*)layoutMgr glyphIndex:glyphIndex];
    return result;
}

void* C_NSTypesetter_SubstituteFontForFont(void* ptr, void* originalFont) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSFont* result = [nSTypesetter substituteFontForFont:(NSFont*)originalFont];
    return result;
}

void* C_NSTypesetter_TextTabForGlyphLocation_WritingDirection_MaxLocation(void* ptr, double glyphLocation, int direction, double maxLocation) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSTextTab* result = [nSTypesetter textTabForGlyphLocation:glyphLocation writingDirection:direction maxLocation:maxLocation];
    return result;
}

void C_NSTypesetter_SetParagraphGlyphRange_SeparatorGlyphRange(void* ptr, NSRange paragraphRange, NSRange paragraphSeparatorRange) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setParagraphGlyphRange:paragraphRange separatorGlyphRange:paragraphSeparatorRange];
}

double C_NSTypesetter_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(void* ptr, unsigned int glyphIndex, CGRect rect) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    CGFloat result = [nSTypesetter lineSpacingAfterGlyphAtIndex:glyphIndex withProposedLineFragmentRect:rect];
    return result;
}

double C_NSTypesetter_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(void* ptr, unsigned int glyphIndex, CGRect rect) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    CGFloat result = [nSTypesetter paragraphSpacingAfterGlyphAtIndex:glyphIndex withProposedLineFragmentRect:rect];
    return result;
}

double C_NSTypesetter_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(void* ptr, unsigned int glyphIndex, CGRect rect) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    CGFloat result = [nSTypesetter paragraphSpacingBeforeGlyphAtIndex:glyphIndex withProposedLineFragmentRect:rect];
    return result;
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
    NSRange result = [nSTypesetter layoutCharactersInRange:characterRange forLayoutManager:(NSLayoutManager*)layoutManager maximumNumberOfLineFragments:maxNumLines];
    return result;
}

CGRect C_NSTypesetter_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(void* ptr, unsigned int glyphIndex, void* textContainer, CGRect proposedRect, CGPoint glyphPosition, unsigned int charIndex) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSRect result = [nSTypesetter boundingBoxForControlGlyphAtIndex:glyphIndex forTextContainer:(NSTextContainer*)textContainer proposedLineFragment:proposedRect glyphPosition:glyphPosition characterIndex:charIndex];
    return result;
}

float C_NSTypesetter_HyphenationFactorForGlyphAtIndex(void* ptr, unsigned int glyphIndex) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    float result = [nSTypesetter hyphenationFactorForGlyphAtIndex:glyphIndex];
    return result;
}

bool C_NSTypesetter_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(void* ptr, unsigned int charIndex) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    BOOL result = [nSTypesetter shouldBreakLineByHyphenatingBeforeCharacterAtIndex:charIndex];
    return result;
}

bool C_NSTypesetter_ShouldBreakLineByWordBeforeCharacterAtIndex(void* ptr, unsigned int charIndex) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    BOOL result = [nSTypesetter shouldBreakLineByWordBeforeCharacterAtIndex:charIndex];
    return result;
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
    NSTypesetter* result = [NSTypesetter sharedSystemTypesetter];
    return result;
}

int C_NSTypesetter_Typesetter_DefaultTypesetterBehavior() {
    NSTypesetterBehavior result = [NSTypesetter defaultTypesetterBehavior];
    return result;
}

void* C_NSTypesetter_LayoutManager(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSLayoutManager* result = [nSTypesetter layoutManager];
    return result;
}

bool C_NSTypesetter_UsesFontLeading(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    BOOL result = [nSTypesetter usesFontLeading];
    return result;
}

void C_NSTypesetter_SetUsesFontLeading(void* ptr, bool value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setUsesFontLeading:value];
}

int C_NSTypesetter_TypesetterBehavior(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSTypesetterBehavior result = [nSTypesetter typesetterBehavior];
    return result;
}

void C_NSTypesetter_SetTypesetterBehavior(void* ptr, int value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setTypesetterBehavior:value];
}

float C_NSTypesetter_HyphenationFactor(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    float result = [nSTypesetter hyphenationFactor];
    return result;
}

void C_NSTypesetter_SetHyphenationFactor(void* ptr, float value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setHyphenationFactor:value];
}

void* C_NSTypesetter_CurrentTextContainer(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSTextContainer* result = [nSTypesetter currentTextContainer];
    return result;
}

Array C_NSTypesetter_TextContainers(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSArray* result = [nSTypesetter textContainers];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

double C_NSTypesetter_LineFragmentPadding(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    CGFloat result = [nSTypesetter lineFragmentPadding];
    return result;
}

void C_NSTypesetter_SetLineFragmentPadding(void* ptr, double value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setLineFragmentPadding:value];
}

bool C_NSTypesetter_BidiProcessingEnabled(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    BOOL result = [nSTypesetter bidiProcessingEnabled];
    return result;
}

void C_NSTypesetter_SetBidiProcessingEnabled(void* ptr, bool value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setBidiProcessingEnabled:value];
}

void* C_NSTypesetter_CurrentParagraphStyle(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSParagraphStyle* result = [nSTypesetter currentParagraphStyle];
    return result;
}

void* C_NSTypesetter_AttributedString(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSAttributedString* result = [nSTypesetter attributedString];
    return result;
}

void C_NSTypesetter_SetAttributedString(void* ptr, void* value) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    [nSTypesetter setAttributedString:(NSAttributedString*)value];
}

NSRange C_NSTypesetter_ParagraphGlyphRange(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSRange result = [nSTypesetter paragraphGlyphRange];
    return result;
}

NSRange C_NSTypesetter_ParagraphSeparatorGlyphRange(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSRange result = [nSTypesetter paragraphSeparatorGlyphRange];
    return result;
}

NSRange C_NSTypesetter_ParagraphCharacterRange(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSRange result = [nSTypesetter paragraphCharacterRange];
    return result;
}

NSRange C_NSTypesetter_ParagraphSeparatorCharacterRange(void* ptr) {
    NSTypesetter* nSTypesetter = (NSTypesetter*)ptr;
    NSRange result = [nSTypesetter paragraphSeparatorCharacterRange];
    return result;
}
