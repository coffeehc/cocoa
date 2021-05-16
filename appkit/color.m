#import <Appkit/Appkit.h>
#import "color.h"

void* C_Color_Alloc() {
    return [NSColor alloc];
}

void* C_NSColor_Init(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result = [nSColor init];
    return result;
}

void* C_NSColor_InitWithCoder(void* ptr, void* coder) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result = [nSColor initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSColor_ColorWithSystemEffect(void* ptr, int systemEffect) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result = [nSColor colorWithSystemEffect:systemEffect];
    return result;
}

void* C_NSColor_ColorUsingColorSpace(void* ptr, void* space) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result = [nSColor colorUsingColorSpace:(NSColorSpace*)space];
    return result;
}

void* C_NSColor_BlendedColorWithFraction_OfColor(void* ptr, double fraction, void* color) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result = [nSColor blendedColorWithFraction:fraction ofColor:(NSColor*)color];
    return result;
}

void* C_NSColor_ColorWithAlphaComponent(void* ptr, double alpha) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result = [nSColor colorWithAlphaComponent:alpha];
    return result;
}

void* C_NSColor_HighlightWithLevel(void* ptr, double val) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result = [nSColor highlightWithLevel:val];
    return result;
}

void* C_NSColor_ShadowWithLevel(void* ptr, double val) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result = [nSColor shadowWithLevel:val];
    return result;
}

void* C_NSColor_ColorFromPasteboard(void* pasteBoard) {
    NSColor* result = [NSColor colorFromPasteboard:(NSPasteboard*)pasteBoard];
    return result;
}

void C_NSColor_WriteToPasteboard(void* ptr, void* pasteBoard) {
    NSColor* nSColor = (NSColor*)ptr;
    [nSColor writeToPasteboard:(NSPasteboard*)pasteBoard];
}

void* C_NSColor_ColorUsingType(void* ptr, int _type) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColor* result = [nSColor colorUsingType:_type];
    return result;
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
    NSColor* result = [NSColor colorNamed:(NSString*)name];
    return result;
}

void* C_NSColor_ColorNamed_Bundle(void* name, void* bundle) {
    NSColor* result = [NSColor colorNamed:(NSString*)name bundle:(NSBundle*)bundle];
    return result;
}

void* C_NSColor_ColorWithCatalogName_ColorName(void* listName, void* colorName) {
    NSColor* result = [NSColor colorWithCatalogName:(NSString*)listName colorName:(NSString*)colorName];
    return result;
}

void* C_NSColor_ColorWithSRGBRed_Green_Blue_Alpha(double red, double green, double blue, double alpha) {
    NSColor* result = [NSColor colorWithSRGBRed:red green:green blue:blue alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithDisplayP3Red_Green_Blue_Alpha(double red, double green, double blue, double alpha) {
    NSColor* result = [NSColor colorWithDisplayP3Red:red green:green blue:blue alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithRed_Green_Blue_Alpha(double red, double green, double blue, double alpha) {
    NSColor* result = [NSColor colorWithRed:red green:green blue:blue alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithCalibratedRed_Green_Blue_Alpha(double red, double green, double blue, double alpha) {
    NSColor* result = [NSColor colorWithCalibratedRed:red green:green blue:blue alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithDeviceRed_Green_Blue_Alpha(double red, double green, double blue, double alpha) {
    NSColor* result = [NSColor colorWithDeviceRed:red green:green blue:blue alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithCalibratedHue_Saturation_Brightness_Alpha(double hue, double saturation, double brightness, double alpha) {
    NSColor* result = [NSColor colorWithCalibratedHue:hue saturation:saturation brightness:brightness alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithDeviceHue_Saturation_Brightness_Alpha(double hue, double saturation, double brightness, double alpha) {
    NSColor* result = [NSColor colorWithDeviceHue:hue saturation:saturation brightness:brightness alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithHue_Saturation_Brightness_Alpha(double hue, double saturation, double brightness, double alpha) {
    NSColor* result = [NSColor colorWithHue:hue saturation:saturation brightness:brightness alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithColorSpace_Hue_Saturation_Brightness_Alpha(void* space, double hue, double saturation, double brightness, double alpha) {
    NSColor* result = [NSColor colorWithColorSpace:(NSColorSpace*)space hue:hue saturation:saturation brightness:brightness alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithDeviceCyan_Magenta_Yellow_Black_Alpha(double cyan, double magenta, double yellow, double black, double alpha) {
    NSColor* result = [NSColor colorWithDeviceCyan:cyan magenta:magenta yellow:yellow black:black alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithWhite_Alpha(double white, double alpha) {
    NSColor* result = [NSColor colorWithWhite:white alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithCalibratedWhite_Alpha(double white, double alpha) {
    NSColor* result = [NSColor colorWithCalibratedWhite:white alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithDeviceWhite_Alpha(double white, double alpha) {
    NSColor* result = [NSColor colorWithDeviceWhite:white alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithGenericGamma22White_Alpha(double white, double alpha) {
    NSColor* result = [NSColor colorWithGenericGamma22White:white alpha:alpha];
    return result;
}

void* C_NSColor_ColorWithPatternImage(void* image) {
    NSColor* result = [NSColor colorWithPatternImage:(NSImage*)image];
    return result;
}

void* C_NSColor_ColorWithCGColor(CGColorRef cgColor) {
    NSColor* result = [NSColor colorWithCGColor:cgColor];
    return result;
}

bool C_NSColor_Color_IgnoresAlpha() {
    BOOL result = [NSColor ignoresAlpha];
    return result;
}

void C_NSColor_Color_SetIgnoresAlpha(bool value) {
    [NSColor setIgnoresAlpha:value];
}

int C_NSColor_NumberOfComponents(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSInteger result = [nSColor numberOfComponents];
    return result;
}

double C_NSColor_AlphaComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor alphaComponent];
    return result;
}

double C_NSColor_WhiteComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor whiteComponent];
    return result;
}

double C_NSColor_RedComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor redComponent];
    return result;
}

double C_NSColor_GreenComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor greenComponent];
    return result;
}

double C_NSColor_BlueComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor blueComponent];
    return result;
}

double C_NSColor_CyanComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor cyanComponent];
    return result;
}

double C_NSColor_MagentaComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor magentaComponent];
    return result;
}

double C_NSColor_YellowComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor yellowComponent];
    return result;
}

double C_NSColor_BlackComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor blackComponent];
    return result;
}

double C_NSColor_HueComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor hueComponent];
    return result;
}

double C_NSColor_SaturationComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor saturationComponent];
    return result;
}

double C_NSColor_BrightnessComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGFloat result = [nSColor brightnessComponent];
    return result;
}

void* C_NSColor_CatalogNameComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColorListName result = [nSColor catalogNameComponent];
    return result;
}

void* C_NSColor_LocalizedCatalogNameComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSString* result = [nSColor localizedCatalogNameComponent];
    return result;
}

void* C_NSColor_ColorNameComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColorName result = [nSColor colorNameComponent];
    return result;
}

void* C_NSColor_LocalizedColorNameComponent(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSString* result = [nSColor localizedColorNameComponent];
    return result;
}

int C_NSColor_Type(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColorType result = [nSColor type];
    return result;
}

void* C_NSColor_ColorSpace(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSColorSpace* result = [nSColor colorSpace];
    return result;
}

CGColorRef C_NSColor_CGColor(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    CGColorRef result = [nSColor CGColor];
    return result;
}

void* C_NSColor_LabelColor() {
    NSColor* result = [NSColor labelColor];
    return result;
}

void* C_NSColor_SecondaryLabelColor() {
    NSColor* result = [NSColor secondaryLabelColor];
    return result;
}

void* C_NSColor_TertiaryLabelColor() {
    NSColor* result = [NSColor tertiaryLabelColor];
    return result;
}

void* C_NSColor_QuaternaryLabelColor() {
    NSColor* result = [NSColor quaternaryLabelColor];
    return result;
}

void* C_NSColor_TextColor() {
    NSColor* result = [NSColor textColor];
    return result;
}

void* C_NSColor_PlaceholderTextColor() {
    NSColor* result = [NSColor placeholderTextColor];
    return result;
}

void* C_NSColor_SelectedTextColor() {
    NSColor* result = [NSColor selectedTextColor];
    return result;
}

void* C_NSColor_TextBackgroundColor() {
    NSColor* result = [NSColor textBackgroundColor];
    return result;
}

void* C_NSColor_SelectedTextBackgroundColor() {
    NSColor* result = [NSColor selectedTextBackgroundColor];
    return result;
}

void* C_NSColor_KeyboardFocusIndicatorColor() {
    NSColor* result = [NSColor keyboardFocusIndicatorColor];
    return result;
}

void* C_NSColor_UnemphasizedSelectedTextColor() {
    NSColor* result = [NSColor unemphasizedSelectedTextColor];
    return result;
}

void* C_NSColor_UnemphasizedSelectedTextBackgroundColor() {
    NSColor* result = [NSColor unemphasizedSelectedTextBackgroundColor];
    return result;
}

void* C_NSColor_LinkColor() {
    NSColor* result = [NSColor linkColor];
    return result;
}

void* C_NSColor_SeparatorColor() {
    NSColor* result = [NSColor separatorColor];
    return result;
}

void* C_NSColor_SelectedContentBackgroundColor() {
    NSColor* result = [NSColor selectedContentBackgroundColor];
    return result;
}

void* C_NSColor_UnemphasizedSelectedContentBackgroundColor() {
    NSColor* result = [NSColor unemphasizedSelectedContentBackgroundColor];
    return result;
}

void* C_NSColor_SelectedMenuItemTextColor() {
    NSColor* result = [NSColor selectedMenuItemTextColor];
    return result;
}

void* C_NSColor_GridColor() {
    NSColor* result = [NSColor gridColor];
    return result;
}

void* C_NSColor_HeaderTextColor() {
    NSColor* result = [NSColor headerTextColor];
    return result;
}

Array C_NSColor_Color_AlternatingContentBackgroundColors() {
    NSArray* result = [NSColor alternatingContentBackgroundColors];
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

void* C_NSColor_ControlAccentColor() {
    NSColor* result = [NSColor controlAccentColor];
    return result;
}

void* C_NSColor_ControlColor() {
    NSColor* result = [NSColor controlColor];
    return result;
}

void* C_NSColor_ControlBackgroundColor() {
    NSColor* result = [NSColor controlBackgroundColor];
    return result;
}

void* C_NSColor_ControlTextColor() {
    NSColor* result = [NSColor controlTextColor];
    return result;
}

void* C_NSColor_DisabledControlTextColor() {
    NSColor* result = [NSColor disabledControlTextColor];
    return result;
}

unsigned int C_NSColor_Color_CurrentControlTint() {
    NSControlTint result = [NSColor currentControlTint];
    return result;
}

void* C_NSColor_SelectedControlColor() {
    NSColor* result = [NSColor selectedControlColor];
    return result;
}

void* C_NSColor_SelectedControlTextColor() {
    NSColor* result = [NSColor selectedControlTextColor];
    return result;
}

void* C_NSColor_AlternateSelectedControlTextColor() {
    NSColor* result = [NSColor alternateSelectedControlTextColor];
    return result;
}

void* C_NSColor_ScrubberTexturedBackgroundColor() {
    NSColor* result = [NSColor scrubberTexturedBackgroundColor];
    return result;
}

void* C_NSColor_WindowBackgroundColor() {
    NSColor* result = [NSColor windowBackgroundColor];
    return result;
}

void* C_NSColor_WindowFrameTextColor() {
    NSColor* result = [NSColor windowFrameTextColor];
    return result;
}

void* C_NSColor_UnderPageBackgroundColor() {
    NSColor* result = [NSColor underPageBackgroundColor];
    return result;
}

void* C_NSColor_FindHighlightColor() {
    NSColor* result = [NSColor findHighlightColor];
    return result;
}

void* C_NSColor_HighlightColor() {
    NSColor* result = [NSColor highlightColor];
    return result;
}

void* C_NSColor_ShadowColor() {
    NSColor* result = [NSColor shadowColor];
    return result;
}

void* C_NSColor_SystemBlueColor() {
    NSColor* result = [NSColor systemBlueColor];
    return result;
}

void* C_NSColor_SystemBrownColor() {
    NSColor* result = [NSColor systemBrownColor];
    return result;
}

void* C_NSColor_SystemGrayColor() {
    NSColor* result = [NSColor systemGrayColor];
    return result;
}

void* C_NSColor_SystemGreenColor() {
    NSColor* result = [NSColor systemGreenColor];
    return result;
}

void* C_NSColor_SystemIndigoColor() {
    NSColor* result = [NSColor systemIndigoColor];
    return result;
}

void* C_NSColor_SystemOrangeColor() {
    NSColor* result = [NSColor systemOrangeColor];
    return result;
}

void* C_NSColor_SystemPinkColor() {
    NSColor* result = [NSColor systemPinkColor];
    return result;
}

void* C_NSColor_SystemPurpleColor() {
    NSColor* result = [NSColor systemPurpleColor];
    return result;
}

void* C_NSColor_SystemRedColor() {
    NSColor* result = [NSColor systemRedColor];
    return result;
}

void* C_NSColor_SystemTealColor() {
    NSColor* result = [NSColor systemTealColor];
    return result;
}

void* C_NSColor_SystemYellowColor() {
    NSColor* result = [NSColor systemYellowColor];
    return result;
}

void* C_NSColor_ClearColor() {
    NSColor* result = [NSColor clearColor];
    return result;
}

void* C_NSColor_BlackColor() {
    NSColor* result = [NSColor blackColor];
    return result;
}

void* C_NSColor_BlueColor() {
    NSColor* result = [NSColor blueColor];
    return result;
}

void* C_NSColor_BrownColor() {
    NSColor* result = [NSColor brownColor];
    return result;
}

void* C_NSColor_CyanColor() {
    NSColor* result = [NSColor cyanColor];
    return result;
}

void* C_NSColor_DarkGrayColor() {
    NSColor* result = [NSColor darkGrayColor];
    return result;
}

void* C_NSColor_GrayColor() {
    NSColor* result = [NSColor grayColor];
    return result;
}

void* C_NSColor_GreenColor() {
    NSColor* result = [NSColor greenColor];
    return result;
}

void* C_NSColor_LightGrayColor() {
    NSColor* result = [NSColor lightGrayColor];
    return result;
}

void* C_NSColor_MagentaColor() {
    NSColor* result = [NSColor magentaColor];
    return result;
}

void* C_NSColor_OrangeColor() {
    NSColor* result = [NSColor orangeColor];
    return result;
}

void* C_NSColor_PurpleColor() {
    NSColor* result = [NSColor purpleColor];
    return result;
}

void* C_NSColor_RedColor() {
    NSColor* result = [NSColor redColor];
    return result;
}

void* C_NSColor_WhiteColor() {
    NSColor* result = [NSColor whiteColor];
    return result;
}

void* C_NSColor_YellowColor() {
    NSColor* result = [NSColor yellowColor];
    return result;
}

void* C_NSColor_PatternImage(void* ptr) {
    NSColor* nSColor = (NSColor*)ptr;
    NSImage* result = [nSColor patternImage];
    return result;
}
