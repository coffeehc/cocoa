#import <Appkit/Appkit.h>
#import "font.h"

void* C_Font_Alloc() {
    return [NSFont alloc];
}

void* C_NSFont_Init(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSFont* result = [nSFont init];
    return result;
}

void* C_NSFont_FontWithName_Size(void* fontName, double fontSize) {
    NSFont* result = [NSFont fontWithName:(NSString*)fontName size:fontSize];
    return result;
}

void* C_NSFont_FontWithDescriptor_Size(void* fontDescriptor, double fontSize) {
    NSFont* result = [NSFont fontWithDescriptor:(NSFontDescriptor*)fontDescriptor size:fontSize];
    return result;
}

void* C_NSFont_FontWithDescriptor_TextTransform(void* fontDescriptor, void* textTransform) {
    NSFont* result = [NSFont fontWithDescriptor:(NSFontDescriptor*)fontDescriptor textTransform:(NSAffineTransform*)textTransform];
    return result;
}

void* C_NSFont_Font_UserFontOfSize(double fontSize) {
    NSFont* result = [NSFont userFontOfSize:fontSize];
    return result;
}

void* C_NSFont_Font_UserFixedPitchFontOfSize(double fontSize) {
    NSFont* result = [NSFont userFixedPitchFontOfSize:fontSize];
    return result;
}

void* C_NSFont_Font_SystemFontOfSize(double fontSize) {
    NSFont* result = [NSFont systemFontOfSize:fontSize];
    return result;
}

void* C_NSFont_Font_SystemFontOfSize_Weight(double fontSize, double weight) {
    NSFont* result = [NSFont systemFontOfSize:fontSize weight:weight];
    return result;
}

void* C_NSFont_Font_BoldSystemFontOfSize(double fontSize) {
    NSFont* result = [NSFont boldSystemFontOfSize:fontSize];
    return result;
}

void* C_NSFont_Font_MonospacedSystemFontOfSize_Weight(double fontSize, double weight) {
    NSFont* result = [NSFont monospacedSystemFontOfSize:fontSize weight:weight];
    return result;
}

void* C_NSFont_Font_MonospacedDigitSystemFontOfSize_Weight(double fontSize, double weight) {
    NSFont* result = [NSFont monospacedDigitSystemFontOfSize:fontSize weight:weight];
    return result;
}

void* C_NSFont_Font_LabelFontOfSize(double fontSize) {
    NSFont* result = [NSFont labelFontOfSize:fontSize];
    return result;
}

void* C_NSFont_Font_MessageFontOfSize(double fontSize) {
    NSFont* result = [NSFont messageFontOfSize:fontSize];
    return result;
}

void* C_NSFont_Font_MenuBarFontOfSize(double fontSize) {
    NSFont* result = [NSFont menuBarFontOfSize:fontSize];
    return result;
}

void* C_NSFont_Font_MenuFontOfSize(double fontSize) {
    NSFont* result = [NSFont menuFontOfSize:fontSize];
    return result;
}

void* C_NSFont_Font_ControlContentFontOfSize(double fontSize) {
    NSFont* result = [NSFont controlContentFontOfSize:fontSize];
    return result;
}

void* C_NSFont_Font_TitleBarFontOfSize(double fontSize) {
    NSFont* result = [NSFont titleBarFontOfSize:fontSize];
    return result;
}

void* C_NSFont_Font_PaletteFontOfSize(double fontSize) {
    NSFont* result = [NSFont paletteFontOfSize:fontSize];
    return result;
}

void* C_NSFont_Font_ToolTipsFontOfSize(double fontSize) {
    NSFont* result = [NSFont toolTipsFontOfSize:fontSize];
    return result;
}

double C_NSFont_Font_SystemFontSizeForControlSize(unsigned int controlSize) {
    CGFloat result = [NSFont systemFontSizeForControlSize:controlSize];
    return result;
}

void C_NSFont_Set(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    [nSFont set];
}

void C_NSFont_SetInContext(void* ptr, void* graphicsContext) {
    NSFont* nSFont = (NSFont*)ptr;
    [nSFont setInContext:(NSGraphicsContext*)graphicsContext];
}

void C_NSFont_SetUserFont(void* font) {
    [NSFont setUserFont:(NSFont*)font];
}

void C_NSFont_SetUserFixedPitchFont(void* font) {
    [NSFont setUserFixedPitchFont:(NSFont*)font];
}

void* C_NSFont_FontWithSize(void* ptr, double fontSize) {
    NSFont* nSFont = (NSFont*)ptr;
    NSFont* result = [nSFont fontWithSize:fontSize];
    return result;
}

double C_NSFont_Font_SystemFontSize() {
    CGFloat result = [NSFont systemFontSize];
    return result;
}

double C_NSFont_Font_SmallSystemFontSize() {
    CGFloat result = [NSFont smallSystemFontSize];
    return result;
}

double C_NSFont_Font_LabelFontSize() {
    CGFloat result = [NSFont labelFontSize];
    return result;
}

double C_NSFont_PointSize(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result = [nSFont pointSize];
    return result;
}

void* C_NSFont_CoveredCharacterSet(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSCharacterSet* result = [nSFont coveredCharacterSet];
    return result;
}

void* C_NSFont_FontDescriptor(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSFontDescriptor* result = [nSFont fontDescriptor];
    return result;
}

bool C_NSFont_IsFixedPitch(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    BOOL result = [nSFont isFixedPitch];
    return result;
}

unsigned int C_NSFont_MostCompatibleStringEncoding(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSStringEncoding result = [nSFont mostCompatibleStringEncoding];
    return result;
}

unsigned int C_NSFont_NumberOfGlyphs(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSUInteger result = [nSFont numberOfGlyphs];
    return result;
}

void* C_NSFont_DisplayName(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSString* result = [nSFont displayName];
    return result;
}

void* C_NSFont_FamilyName(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSString* result = [nSFont familyName];
    return result;
}

void* C_NSFont_FontName(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSString* result = [nSFont fontName];
    return result;
}

bool C_NSFont_IsVertical(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    BOOL result = [nSFont isVertical];
    return result;
}

void* C_NSFont_VerticalFont(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSFont* result = [nSFont verticalFont];
    return result;
}

double C_NSFont_Ascender(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result = [nSFont ascender];
    return result;
}

double C_NSFont_Descender(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result = [nSFont descender];
    return result;
}

double C_NSFont_CapHeight(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result = [nSFont capHeight];
    return result;
}

double C_NSFont_Leading(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result = [nSFont leading];
    return result;
}

double C_NSFont_XHeight(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result = [nSFont xHeight];
    return result;
}

double C_NSFont_ItalicAngle(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result = [nSFont italicAngle];
    return result;
}

double C_NSFont_UnderlinePosition(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result = [nSFont underlinePosition];
    return result;
}

double C_NSFont_UnderlineThickness(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result = [nSFont underlineThickness];
    return result;
}

CGRect C_NSFont_BoundingRectForFont(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSRect result = [nSFont boundingRectForFont];
    return result;
}

CGSize C_NSFont_MaximumAdvancement(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSSize result = [nSFont maximumAdvancement];
    return result;
}

void* C_NSFont_TextTransform(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSAffineTransform* result = [nSFont textTransform];
    return result;
}
