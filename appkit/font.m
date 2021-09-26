#import "font.h"
#import <AppKit/NSFont.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_Font_Alloc() {
    return [NSFont alloc];
}

void* C_NSFont_AllocFont() {
    NSFont* result_ = [NSFont alloc];
    return result_;
}

void* C_NSFont_Init(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSFont* result_ = [nSFont init];
    return result_;
}

void* C_NSFont_NewFont() {
    NSFont* result_ = [NSFont new];
    return result_;
}

void* C_NSFont_Autorelease(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSFont* result_ = [nSFont autorelease];
    return result_;
}

void* C_NSFont_Retain(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSFont* result_ = [nSFont retain];
    return result_;
}

void* C_NSFont_FontWithName_Size(void* fontName, double fontSize) {
    NSFont* result_ = [NSFont fontWithName:(NSString*)fontName size:fontSize];
    return result_;
}

void* C_NSFont_FontWithDescriptor_Size(void* fontDescriptor, double fontSize) {
    NSFont* result_ = [NSFont fontWithDescriptor:(NSFontDescriptor*)fontDescriptor size:fontSize];
    return result_;
}

void* C_NSFont_FontWithDescriptor_TextTransform(void* fontDescriptor, void* textTransform) {
    NSFont* result_ = [NSFont fontWithDescriptor:(NSFontDescriptor*)fontDescriptor textTransform:(NSAffineTransform*)textTransform];
    return result_;
}

void* C_NSFont_UserFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont userFontOfSize:fontSize];
    return result_;
}

void* C_NSFont_UserFixedPitchFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont userFixedPitchFontOfSize:fontSize];
    return result_;
}

void* C_NSFont_PreferredFontForTextStyle_Options(void* style, Dictionary options) {
    NSMutableDictionary* objcOptions = [NSMutableDictionary dictionaryWithCapacity:options.len];
    if (options.len > 0) {
    	void** optionsKeyData = (void**)options.key_data;
    	void** optionsValueData = (void**)options.value_data;
    	for (int i = 0; i < options.len; i++) {
    		void* kp = optionsKeyData[i];
    		void* vp = optionsValueData[i];
    		[objcOptions setObject:(id)vp forKey:(id<NSCopying>)(NSString*)kp];
    	}
    }
    NSFont* result_ = [NSFont preferredFontForTextStyle:(NSString*)style options:objcOptions];
    return result_;
}

void* C_NSFont_SystemFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont systemFontOfSize:fontSize];
    return result_;
}

void* C_NSFont_SystemFontOfSize_Weight(double fontSize, double weight) {
    NSFont* result_ = [NSFont systemFontOfSize:fontSize weight:weight];
    return result_;
}

void* C_NSFont_BoldSystemFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont boldSystemFontOfSize:fontSize];
    return result_;
}

void* C_NSFont_MonospacedSystemFontOfSize_Weight(double fontSize, double weight) {
    NSFont* result_ = [NSFont monospacedSystemFontOfSize:fontSize weight:weight];
    return result_;
}

void* C_NSFont_MonospacedDigitSystemFontOfSize_Weight(double fontSize, double weight) {
    NSFont* result_ = [NSFont monospacedDigitSystemFontOfSize:fontSize weight:weight];
    return result_;
}

void* C_NSFont_LabelFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont labelFontOfSize:fontSize];
    return result_;
}

void* C_NSFont_MessageFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont messageFontOfSize:fontSize];
    return result_;
}

void* C_NSFont_MenuBarFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont menuBarFontOfSize:fontSize];
    return result_;
}

void* C_NSFont_MenuFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont menuFontOfSize:fontSize];
    return result_;
}

void* C_NSFont_ControlContentFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont controlContentFontOfSize:fontSize];
    return result_;
}

void* C_NSFont_TitleBarFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont titleBarFontOfSize:fontSize];
    return result_;
}

void* C_NSFont_PaletteFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont paletteFontOfSize:fontSize];
    return result_;
}

void* C_NSFont_ToolTipsFontOfSize(double fontSize) {
    NSFont* result_ = [NSFont toolTipsFontOfSize:fontSize];
    return result_;
}

double C_NSFont_SystemFontSizeForControlSize(unsigned int controlSize) {
    CGFloat result_ = [NSFont systemFontSizeForControlSize:controlSize];
    return result_;
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
    NSFont* result_ = [nSFont fontWithSize:fontSize];
    return result_;
}

double C_NSFont_SystemFontSize() {
    CGFloat result_ = [NSFont systemFontSize];
    return result_;
}

double C_NSFont_SmallSystemFontSize() {
    CGFloat result_ = [NSFont smallSystemFontSize];
    return result_;
}

double C_NSFont_LabelFontSize() {
    CGFloat result_ = [NSFont labelFontSize];
    return result_;
}

double C_NSFont_PointSize(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result_ = [nSFont pointSize];
    return result_;
}

void* C_NSFont_CoveredCharacterSet(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSCharacterSet* result_ = [nSFont coveredCharacterSet];
    return result_;
}

void* C_NSFont_FontDescriptor(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSFontDescriptor* result_ = [nSFont fontDescriptor];
    return result_;
}

bool C_NSFont_IsFixedPitch(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    BOOL result_ = [nSFont isFixedPitch];
    return result_;
}

unsigned int C_NSFont_MostCompatibleStringEncoding(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSStringEncoding result_ = [nSFont mostCompatibleStringEncoding];
    return result_;
}

unsigned int C_NSFont_NumberOfGlyphs(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSUInteger result_ = [nSFont numberOfGlyphs];
    return result_;
}

void* C_NSFont_DisplayName(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSString* result_ = [nSFont displayName];
    return result_;
}

void* C_NSFont_FamilyName(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSString* result_ = [nSFont familyName];
    return result_;
}

void* C_NSFont_FontName(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSString* result_ = [nSFont fontName];
    return result_;
}

bool C_NSFont_IsVertical(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    BOOL result_ = [nSFont isVertical];
    return result_;
}

void* C_NSFont_VerticalFont(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSFont* result_ = [nSFont verticalFont];
    return result_;
}

double C_NSFont_Ascender(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result_ = [nSFont ascender];
    return result_;
}

double C_NSFont_Descender(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result_ = [nSFont descender];
    return result_;
}

double C_NSFont_CapHeight(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result_ = [nSFont capHeight];
    return result_;
}

double C_NSFont_Leading(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result_ = [nSFont leading];
    return result_;
}

double C_NSFont_XHeight(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result_ = [nSFont xHeight];
    return result_;
}

double C_NSFont_ItalicAngle(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result_ = [nSFont italicAngle];
    return result_;
}

double C_NSFont_UnderlinePosition(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result_ = [nSFont underlinePosition];
    return result_;
}

double C_NSFont_UnderlineThickness(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    CGFloat result_ = [nSFont underlineThickness];
    return result_;
}

CGRect C_NSFont_BoundingRectForFont(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSRect result_ = [nSFont boundingRectForFont];
    return result_;
}

CGSize C_NSFont_MaximumAdvancement(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSSize result_ = [nSFont maximumAdvancement];
    return result_;
}

void* C_NSFont_TextTransform(void* ptr) {
    NSFont* nSFont = (NSFont*)ptr;
    NSAffineTransform* result_ = [nSFont textTransform];
    return result_;
}
