#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Typesetter_Alloc();

void* C_NSTypesetter_Init(void* ptr);
void* C_NSTypesetter_SharedSystemTypesetterForBehavior(int behavior);
double C_NSTypesetter_BaselineOffsetInLayoutManager_GlyphIndex(void* ptr, void* layoutMgr, unsigned int glyphIndex);
void* C_NSTypesetter_SubstituteFontForFont(void* ptr, void* originalFont);
void* C_NSTypesetter_TextTabForGlyphLocation_WritingDirection_MaxLocation(void* ptr, double glyphLocation, int direction, double maxLocation);
void C_NSTypesetter_SetParagraphGlyphRange_SeparatorGlyphRange(void* ptr, NSRange paragraphRange, NSRange paragraphSeparatorRange);
double C_NSTypesetter_LineSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(void* ptr, unsigned int glyphIndex, CGRect rect);
double C_NSTypesetter_ParagraphSpacingAfterGlyphAtIndex_WithProposedLineFragmentRect(void* ptr, unsigned int glyphIndex, CGRect rect);
double C_NSTypesetter_ParagraphSpacingBeforeGlyphAtIndex_WithProposedLineFragmentRect(void* ptr, unsigned int glyphIndex, CGRect rect);
void C_NSTypesetter_BeginParagraph(void* ptr);
void C_NSTypesetter_EndParagraph(void* ptr);
void C_NSTypesetter_BeginLineWithGlyphAtIndex(void* ptr, unsigned int glyphIndex);
void C_NSTypesetter_EndLineWithGlyphRange(void* ptr, NSRange lineGlyphRange);
NSRange C_NSTypesetter_LayoutCharactersInRange_ForLayoutManager_MaximumNumberOfLineFragments(void* ptr, NSRange characterRange, void* layoutManager, unsigned int maxNumLines);
CGRect C_NSTypesetter_BoundingBoxForControlGlyphAtIndex_ForTextContainer_ProposedLineFragment_GlyphPosition_CharacterIndex(void* ptr, unsigned int glyphIndex, void* textContainer, CGRect proposedRect, CGPoint glyphPosition, unsigned int charIndex);
float C_NSTypesetter_HyphenationFactorForGlyphAtIndex(void* ptr, unsigned int glyphIndex);
bool C_NSTypesetter_ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(void* ptr, unsigned int charIndex);
bool C_NSTypesetter_ShouldBreakLineByWordBeforeCharacterAtIndex(void* ptr, unsigned int charIndex);
void C_NSTypesetter_SetHardInvalidation_ForGlyphRange(void* ptr, bool flag, NSRange glyphRange);
void C_NSTypesetter_SetAttachmentSize_ForGlyphRange(void* ptr, CGSize attachmentSize, NSRange glyphRange);
void C_NSTypesetter_SetDrawsOutsideLineFragment_ForGlyphRange(void* ptr, bool flag, NSRange glyphRange);
void C_NSTypesetter_SetLineFragmentRect_ForGlyphRange_UsedRect_BaselineOffset(void* ptr, CGRect fragmentRect, NSRange glyphRange, CGRect usedRect, double baselineOffset);
void C_NSTypesetter_SetNotShownAttribute_ForGlyphRange(void* ptr, bool flag, NSRange glyphRange);
void* C_NSTypesetter_SharedSystemTypesetter();
int C_NSTypesetter_DefaultTypesetterBehavior();
void* C_NSTypesetter_LayoutManager(void* ptr);
bool C_NSTypesetter_UsesFontLeading(void* ptr);
void C_NSTypesetter_SetUsesFontLeading(void* ptr, bool value);
int C_NSTypesetter_TypesetterBehavior(void* ptr);
void C_NSTypesetter_SetTypesetterBehavior(void* ptr, int value);
float C_NSTypesetter_HyphenationFactor(void* ptr);
void C_NSTypesetter_SetHyphenationFactor(void* ptr, float value);
void* C_NSTypesetter_CurrentTextContainer(void* ptr);
Array C_NSTypesetter_TextContainers(void* ptr);
double C_NSTypesetter_LineFragmentPadding(void* ptr);
void C_NSTypesetter_SetLineFragmentPadding(void* ptr, double value);
bool C_NSTypesetter_BidiProcessingEnabled(void* ptr);
void C_NSTypesetter_SetBidiProcessingEnabled(void* ptr, bool value);
void* C_NSTypesetter_CurrentParagraphStyle(void* ptr);
void* C_NSTypesetter_AttributedString(void* ptr);
void C_NSTypesetter_SetAttributedString(void* ptr, void* value);
NSRange C_NSTypesetter_ParagraphGlyphRange(void* ptr);
NSRange C_NSTypesetter_ParagraphSeparatorGlyphRange(void* ptr);
NSRange C_NSTypesetter_ParagraphCharacterRange(void* ptr);
NSRange C_NSTypesetter_ParagraphSeparatorCharacterRange(void* ptr);
Dictionary C_NSTypesetter_AttributesForExtraLineFragment(void* ptr);
