#import <Appkit/Appkit.h>
#import "image_symbol_configuration.h"

void* C_ImageSymbolConfiguration_Alloc() {
    return [NSImageSymbolConfiguration alloc];
}

void* C_NSImageSymbolConfiguration_Init(void* ptr) {
    NSImageSymbolConfiguration* nSImageSymbolConfiguration = (NSImageSymbolConfiguration*)ptr;
    NSImageSymbolConfiguration* result = [nSImageSymbolConfiguration init];
    return result;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithPointSize_Weight(double pointSize, double weight) {
    NSImageSymbolConfiguration* result = [NSImageSymbolConfiguration configurationWithPointSize:pointSize weight:weight];
    return result;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithPointSize_Weight_Scale(double pointSize, double weight, int scale) {
    NSImageSymbolConfiguration* result = [NSImageSymbolConfiguration configurationWithPointSize:pointSize weight:weight scale:scale];
    return result;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithTextStyle(void* style) {
    NSImageSymbolConfiguration* result = [NSImageSymbolConfiguration configurationWithTextStyle:(NSString*)style];
    return result;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithTextStyle_Scale(void* style, int scale) {
    NSImageSymbolConfiguration* result = [NSImageSymbolConfiguration configurationWithTextStyle:(NSString*)style scale:scale];
    return result;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithScale(int scale) {
    NSImageSymbolConfiguration* result = [NSImageSymbolConfiguration configurationWithScale:scale];
    return result;
}
