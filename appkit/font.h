#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Font_Alloc();

void* C_NSFont_Init(void* ptr);
void* C_NSFont_FontWithName_Size(void* fontName, double fontSize);
void* C_NSFont_FontWithDescriptor_Size(void* fontDescriptor, double fontSize);
void* C_NSFont_FontWithDescriptor_TextTransform(void* fontDescriptor, void* textTransform);
void* C_NSFont_UserFontOfSize(double fontSize);
void* C_NSFont_UserFixedPitchFontOfSize(double fontSize);
void* C_NSFont_PreferredFontForTextStyle_Options(void* style, Dictionary options);
void* C_NSFont_SystemFontOfSize(double fontSize);
void* C_NSFont_SystemFontOfSize_Weight(double fontSize, double weight);
void* C_NSFont_BoldSystemFontOfSize(double fontSize);
void* C_NSFont_MonospacedSystemFontOfSize_Weight(double fontSize, double weight);
void* C_NSFont_MonospacedDigitSystemFontOfSize_Weight(double fontSize, double weight);
void* C_NSFont_LabelFontOfSize(double fontSize);
void* C_NSFont_MessageFontOfSize(double fontSize);
void* C_NSFont_MenuBarFontOfSize(double fontSize);
void* C_NSFont_MenuFontOfSize(double fontSize);
void* C_NSFont_ControlContentFontOfSize(double fontSize);
void* C_NSFont_TitleBarFontOfSize(double fontSize);
void* C_NSFont_PaletteFontOfSize(double fontSize);
void* C_NSFont_ToolTipsFontOfSize(double fontSize);
double C_NSFont_SystemFontSizeForControlSize(unsigned int controlSize);
void C_NSFont_Set(void* ptr);
void C_NSFont_SetInContext(void* ptr, void* graphicsContext);
void C_NSFont_SetUserFont(void* font);
void C_NSFont_SetUserFixedPitchFont(void* font);
void* C_NSFont_FontWithSize(void* ptr, double fontSize);
double C_NSFont_SystemFontSize();
double C_NSFont_SmallSystemFontSize();
double C_NSFont_LabelFontSize();
double C_NSFont_PointSize(void* ptr);
void* C_NSFont_CoveredCharacterSet(void* ptr);
void* C_NSFont_FontDescriptor(void* ptr);
bool C_NSFont_IsFixedPitch(void* ptr);
unsigned int C_NSFont_MostCompatibleStringEncoding(void* ptr);
unsigned int C_NSFont_NumberOfGlyphs(void* ptr);
void* C_NSFont_DisplayName(void* ptr);
void* C_NSFont_FamilyName(void* ptr);
void* C_NSFont_FontName(void* ptr);
bool C_NSFont_IsVertical(void* ptr);
void* C_NSFont_VerticalFont(void* ptr);
double C_NSFont_Ascender(void* ptr);
double C_NSFont_Descender(void* ptr);
double C_NSFont_CapHeight(void* ptr);
double C_NSFont_Leading(void* ptr);
double C_NSFont_XHeight(void* ptr);
double C_NSFont_ItalicAngle(void* ptr);
double C_NSFont_UnderlinePosition(void* ptr);
double C_NSFont_UnderlineThickness(void* ptr);
CGRect C_NSFont_BoundingRectForFont(void* ptr);
CGSize C_NSFont_MaximumAdvancement(void* ptr);
void* C_NSFont_TextTransform(void* ptr);
