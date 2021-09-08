#import <Appkit/Appkit.h>
#import "color.h"

void* C_Color_Alloc() {
    return [NSColor alloc];
}

void* C_NSColor_Init(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result_ = [nSColor init];
    return result_;
}

void* C_NSColor_InitWithCoder(void* ptr, void* coder) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result_ = [nSColor initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSColor_AllocColor() {
    NSColor* result_ = [NSColor alloc];
    return result_;
}

void* C_NSColor_NewColor() {
    NSColor* result_ = [NSColor new];
    return result_;
}

void* C_NSColor_Autorelease(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result_ = [nSColor autorelease];
    return result_;
}

void* C_NSColor_Retain(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result_ = [nSColor retain];
    return result_;
}

void* C_NSColor_ColorWithSystemEffect(void* ptr, int systemEffect) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result_ = [nSColor colorWithSystemEffect:systemEffect];
    return result_;
}

void* C_NSColor_ColorUsingColorSpace(void* ptr, void* space) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result_ = [nSColor colorUsingColorSpace:(NSColorSpace*)space];
    return result_;
}

void* C_NSColor_BlendedColorWithFraction_OfColor(void* ptr, double fraction, void* color) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result_ = [nSColor blendedColorWithFraction:fraction ofColor:(NSColor*)color];
    return result_;
}

void* C_NSColor_ColorWithAlphaComponent(void* ptr, double alpha) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result_ = [nSColor colorWithAlphaComponent:alpha];
    return result_;
}

void* C_NSColor_HighlightWithLevel(void* ptr, double val) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result_ = [nSColor highlightWithLevel:val];
    return result_;
}

void* C_NSColor_ShadowWithLevel(void* ptr, double val) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result_ = [nSColor shadowWithLevel:val];
    return result_;
}

void* C_NSColor_ColorFromPasteboard(void* pasteBoard) {
    NSColor* result_ = [NSColor colorFromPasteboard:(NSPasteboard*)pasteBoard];
    return result_;
}

void C_NSColor_WriteToPasteboard(void* ptr, void* pasteBoard) {
    NSColor* nSColor = (NSColor*)ptr;
    [nSColor writeToPasteboard:(NSPasteboard*)pasteBoard];
}

void* C_NSColor_ColorUsingType(void* ptr, int _type) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result_ = [nSColor colorUsingType:_type];
    return result_;
}

void C_NSColor_DrawSwatchInRect(void* ptr, CGRect rect) {
    NSColor* nSColor = (NSColor*)ptr;
    [nSColor drawSwatchInRect:rect];
}

void C_NSColor_Set(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    [nSColor set];
}

void C_NSColor_SetFill(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    [nSColor setFill];
}

void C_NSColor_SetStroke(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    [nSColor setStroke];
}

void* C_NSColor_ColorNamed(void* name) {
    NSColor* result_ = [NSColor colorNamed:(NSString*)name];
    return result_;
}

void* C_NSColor_ColorNamed_Bundle(void* name, void* bundle) {
    NSColor* result_ = [NSColor colorNamed:(NSString*)name bundle:(NSBundle*)bundle];
    return result_;
}

void* C_NSColor_ColorWithCatalogName_ColorName(void* listName, void* colorName) {
    NSColor* result_ = [NSColor colorWithCatalogName:(NSString*)listName colorName:(NSString*)colorName];
    return result_;
}

void* C_NSColor_ColorWithSRGBRed_Green_Blue_Alpha(double red, double green, double blue, double alpha) {
    NSColor* result_ = [NSColor colorWithSRGBRed:red green:green blue:blue alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithDisplayP3Red_Green_Blue_Alpha(double red, double green, double blue, double alpha) {
    NSColor* result_ = [NSColor colorWithDisplayP3Red:red green:green blue:blue alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithRed_Green_Blue_Alpha(double red, double green, double blue, double alpha) {
    NSColor* result_ = [NSColor colorWithRed:red green:green blue:blue alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithCalibratedRed_Green_Blue_Alpha(double red, double green, double blue, double alpha) {
    NSColor* result_ = [NSColor colorWithCalibratedRed:red green:green blue:blue alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithDeviceRed_Green_Blue_Alpha(double red, double green, double blue, double alpha) {
    NSColor* result_ = [NSColor colorWithDeviceRed:red green:green blue:blue alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithCalibratedHue_Saturation_Brightness_Alpha(double hue, double saturation, double brightness, double alpha) {
    NSColor* result_ = [NSColor colorWithCalibratedHue:hue saturation:saturation brightness:brightness alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithDeviceHue_Saturation_Brightness_Alpha(double hue, double saturation, double brightness, double alpha) {
    NSColor* result_ = [NSColor colorWithDeviceHue:hue saturation:saturation brightness:brightness alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithHue_Saturation_Brightness_Alpha(double hue, double saturation, double brightness, double alpha) {
    NSColor* result_ = [NSColor colorWithHue:hue saturation:saturation brightness:brightness alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithColorSpace_Hue_Saturation_Brightness_Alpha(void* space, double hue, double saturation, double brightness, double alpha) {
    NSColor* result_ = [NSColor colorWithColorSpace:(NSColorSpace*)space hue:hue saturation:saturation brightness:brightness alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithDeviceCyan_Magenta_Yellow_Black_Alpha(double cyan, double magenta, double yellow, double black, double alpha) {
    NSColor* result_ = [NSColor colorWithDeviceCyan:cyan magenta:magenta yellow:yellow black:black alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithWhite_Alpha(double white, double alpha) {
    NSColor* result_ = [NSColor colorWithWhite:white alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithCalibratedWhite_Alpha(double white, double alpha) {
    NSColor* result_ = [NSColor colorWithCalibratedWhite:white alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithDeviceWhite_Alpha(double white, double alpha) {
    NSColor* result_ = [NSColor colorWithDeviceWhite:white alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithGenericGamma22White_Alpha(double white, double alpha) {
    NSColor* result_ = [NSColor colorWithGenericGamma22White:white alpha:alpha];
    return result_;
}

void* C_NSColor_ColorWithPatternImage(void* image) {
    NSColor* result_ = [NSColor colorWithPatternImage:(NSImage*)image];
    return result_;
}

void* C_NSColor_ColorWithCGColor(void* cgColor) {
    NSColor* result_ = [NSColor colorWithCGColor:(CGColorRef)cgColor];
    return result_;
}

bool C_NSColor_Color_IgnoresAlpha() {
    BOOL result_ = [NSColor ignoresAlpha];
    return result_;
}

void C_NSColor_Color_SetIgnoresAlpha(bool value) {
    [NSColor setIgnoresAlpha:value];
}

int C_NSColor_NumberOfComponents(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSInteger result_ = [nSColor numberOfComponents];
    return result_;
}

double C_NSColor_AlphaComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor alphaComponent];
    return result_;
}

double C_NSColor_WhiteComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor whiteComponent];
    return result_;
}

double C_NSColor_RedComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor redComponent];
    return result_;
}

double C_NSColor_GreenComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor greenComponent];
    return result_;
}

double C_NSColor_BlueComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor blueComponent];
    return result_;
}

double C_NSColor_CyanComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor cyanComponent];
    return result_;
}

double C_NSColor_MagentaComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor magentaComponent];
    return result_;
}

double C_NSColor_YellowComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor yellowComponent];
    return result_;
}

double C_NSColor_BlackComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor blackComponent];
    return result_;
}

double C_NSColor_HueComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor hueComponent];
    return result_;
}

double C_NSColor_SaturationComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor saturationComponent];
    return result_;
}

double C_NSColor_BrightnessComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result_ = [nSColor brightnessComponent];
    return result_;
}

void* C_NSColor_CatalogNameComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColorListName result_ = [nSColor catalogNameComponent];
    return result_;
}

void* C_NSColor_LocalizedCatalogNameComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSString* result_ = [nSColor localizedCatalogNameComponent];
    return result_;
}

void* C_NSColor_ColorNameComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColorName result_ = [nSColor colorNameComponent];
    return result_;
}

void* C_NSColor_LocalizedColorNameComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSString* result_ = [nSColor localizedColorNameComponent];
    return result_;
}

int C_NSColor_Type(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColorType result_ = [nSColor type];
    return result_;
}

void* C_NSColor_ColorSpace(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColorSpace* result_ = [nSColor colorSpace];
    return result_;
}

void* C_NSColor_CGColor(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGColorRef result_ = [nSColor CGColor];
    return result_;
}

void* C_NSColor_LabelColor() {
    NSColor* result_ = [NSColor labelColor];
    return result_;
}

void* C_NSColor_SecondaryLabelColor() {
    NSColor* result_ = [NSColor secondaryLabelColor];
    return result_;
}

void* C_NSColor_TertiaryLabelColor() {
    NSColor* result_ = [NSColor tertiaryLabelColor];
    return result_;
}

void* C_NSColor_QuaternaryLabelColor() {
    NSColor* result_ = [NSColor quaternaryLabelColor];
    return result_;
}

void* C_NSColor_TextColor() {
    NSColor* result_ = [NSColor textColor];
    return result_;
}

void* C_NSColor_PlaceholderTextColor() {
    NSColor* result_ = [NSColor placeholderTextColor];
    return result_;
}

void* C_NSColor_SelectedTextColor() {
    NSColor* result_ = [NSColor selectedTextColor];
    return result_;
}

void* C_NSColor_TextBackgroundColor() {
    NSColor* result_ = [NSColor textBackgroundColor];
    return result_;
}

void* C_NSColor_SelectedTextBackgroundColor() {
    NSColor* result_ = [NSColor selectedTextBackgroundColor];
    return result_;
}

void* C_NSColor_KeyboardFocusIndicatorColor() {
    NSColor* result_ = [NSColor keyboardFocusIndicatorColor];
    return result_;
}

void* C_NSColor_UnemphasizedSelectedTextColor() {
    NSColor* result_ = [NSColor unemphasizedSelectedTextColor];
    return result_;
}

void* C_NSColor_UnemphasizedSelectedTextBackgroundColor() {
    NSColor* result_ = [NSColor unemphasizedSelectedTextBackgroundColor];
    return result_;
}

void* C_NSColor_LinkColor() {
    NSColor* result_ = [NSColor linkColor];
    return result_;
}

void* C_NSColor_SeparatorColor() {
    NSColor* result_ = [NSColor separatorColor];
    return result_;
}

void* C_NSColor_SelectedContentBackgroundColor() {
    NSColor* result_ = [NSColor selectedContentBackgroundColor];
    return result_;
}

void* C_NSColor_UnemphasizedSelectedContentBackgroundColor() {
    NSColor* result_ = [NSColor unemphasizedSelectedContentBackgroundColor];
    return result_;
}

void* C_NSColor_SelectedMenuItemTextColor() {
    NSColor* result_ = [NSColor selectedMenuItemTextColor];
    return result_;
}

void* C_NSColor_GridColor() {
    NSColor* result_ = [NSColor gridColor];
    return result_;
}

void* C_NSColor_HeaderTextColor() {
    NSColor* result_ = [NSColor headerTextColor];
    return result_;
}

Array C_NSColor_AlternatingContentBackgroundColors() {
    NSArray* result_ = [NSColor alternatingContentBackgroundColors];
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

void* C_NSColor_ControlAccentColor() {
    NSColor* result_ = [NSColor controlAccentColor];
    return result_;
}

void* C_NSColor_ControlColor() {
    NSColor* result_ = [NSColor controlColor];
    return result_;
}

void* C_NSColor_ControlBackgroundColor() {
    NSColor* result_ = [NSColor controlBackgroundColor];
    return result_;
}

void* C_NSColor_ControlTextColor() {
    NSColor* result_ = [NSColor controlTextColor];
    return result_;
}

void* C_NSColor_DisabledControlTextColor() {
    NSColor* result_ = [NSColor disabledControlTextColor];
    return result_;
}

unsigned int C_NSColor_Color_CurrentControlTint() {
    NSControlTint result_ = [NSColor currentControlTint];
    return result_;
}

void* C_NSColor_SelectedControlColor() {
    NSColor* result_ = [NSColor selectedControlColor];
    return result_;
}

void* C_NSColor_SelectedControlTextColor() {
    NSColor* result_ = [NSColor selectedControlTextColor];
    return result_;
}

void* C_NSColor_AlternateSelectedControlTextColor() {
    NSColor* result_ = [NSColor alternateSelectedControlTextColor];
    return result_;
}

void* C_NSColor_ScrubberTexturedBackgroundColor() {
    NSColor* result_ = [NSColor scrubberTexturedBackgroundColor];
    return result_;
}

void* C_NSColor_WindowBackgroundColor() {
    NSColor* result_ = [NSColor windowBackgroundColor];
    return result_;
}

void* C_NSColor_WindowFrameTextColor() {
    NSColor* result_ = [NSColor windowFrameTextColor];
    return result_;
}

void* C_NSColor_UnderPageBackgroundColor() {
    NSColor* result_ = [NSColor underPageBackgroundColor];
    return result_;
}

void* C_NSColor_FindHighlightColor() {
    NSColor* result_ = [NSColor findHighlightColor];
    return result_;
}

void* C_NSColor_HighlightColor() {
    NSColor* result_ = [NSColor highlightColor];
    return result_;
}

void* C_NSColor_ShadowColor() {
    NSColor* result_ = [NSColor shadowColor];
    return result_;
}

void* C_NSColor_SystemBlueColor() {
    NSColor* result_ = [NSColor systemBlueColor];
    return result_;
}

void* C_NSColor_SystemBrownColor() {
    NSColor* result_ = [NSColor systemBrownColor];
    return result_;
}

void* C_NSColor_SystemGrayColor() {
    NSColor* result_ = [NSColor systemGrayColor];
    return result_;
}

void* C_NSColor_SystemGreenColor() {
    NSColor* result_ = [NSColor systemGreenColor];
    return result_;
}

void* C_NSColor_SystemIndigoColor() {
    NSColor* result_ = [NSColor systemIndigoColor];
    return result_;
}

void* C_NSColor_SystemOrangeColor() {
    NSColor* result_ = [NSColor systemOrangeColor];
    return result_;
}

void* C_NSColor_SystemPinkColor() {
    NSColor* result_ = [NSColor systemPinkColor];
    return result_;
}

void* C_NSColor_SystemPurpleColor() {
    NSColor* result_ = [NSColor systemPurpleColor];
    return result_;
}

void* C_NSColor_SystemRedColor() {
    NSColor* result_ = [NSColor systemRedColor];
    return result_;
}

void* C_NSColor_SystemTealColor() {
    NSColor* result_ = [NSColor systemTealColor];
    return result_;
}

void* C_NSColor_SystemYellowColor() {
    NSColor* result_ = [NSColor systemYellowColor];
    return result_;
}

void* C_NSColor_ClearColor() {
    NSColor* result_ = [NSColor clearColor];
    return result_;
}

void* C_NSColor_BlackColor() {
    NSColor* result_ = [NSColor blackColor];
    return result_;
}

void* C_NSColor_BlueColor() {
    NSColor* result_ = [NSColor blueColor];
    return result_;
}

void* C_NSColor_BrownColor() {
    NSColor* result_ = [NSColor brownColor];
    return result_;
}

void* C_NSColor_CyanColor() {
    NSColor* result_ = [NSColor cyanColor];
    return result_;
}

void* C_NSColor_DarkGrayColor() {
    NSColor* result_ = [NSColor darkGrayColor];
    return result_;
}

void* C_NSColor_GrayColor() {
    NSColor* result_ = [NSColor grayColor];
    return result_;
}

void* C_NSColor_GreenColor() {
    NSColor* result_ = [NSColor greenColor];
    return result_;
}

void* C_NSColor_LightGrayColor() {
    NSColor* result_ = [NSColor lightGrayColor];
    return result_;
}

void* C_NSColor_MagentaColor() {
    NSColor* result_ = [NSColor magentaColor];
    return result_;
}

void* C_NSColor_OrangeColor() {
    NSColor* result_ = [NSColor orangeColor];
    return result_;
}

void* C_NSColor_PurpleColor() {
    NSColor* result_ = [NSColor purpleColor];
    return result_;
}

void* C_NSColor_RedColor() {
    NSColor* result_ = [NSColor redColor];
    return result_;
}

void* C_NSColor_WhiteColor() {
    NSColor* result_ = [NSColor whiteColor];
    return result_;
}

void* C_NSColor_YellowColor() {
    NSColor* result_ = [NSColor yellowColor];
    return result_;
}

void* C_NSColor_PatternImage(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSImage* result_ = [nSColor patternImage];
    return result_;
}
