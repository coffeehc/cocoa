#import <Appkit/Appkit.h>
#import "image_symbol_configuration.h"

void* C_ImageSymbolConfiguration_Alloc() {
    return [NSImageSymbolConfiguration alloc];
}

void* C_NSImageSymbolConfiguration_AllocImageSymbolConfiguration() {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration alloc];
    return result_;
}

void* C_NSImageSymbolConfiguration_Init(void* ptr) {
    NSImageSymbolConfiguration* nSImageSymbolConfiguration = (NSImageSymbolConfiguration*)ptr;
    NSImageSymbolConfiguration* result_ = [nSImageSymbolConfiguration init];
    return result_;
}

void* C_NSImageSymbolConfiguration_NewImageSymbolConfiguration() {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration new];
    return result_;
}

void* C_NSImageSymbolConfiguration_Autorelease(void* ptr) {
    NSImageSymbolConfiguration* nSImageSymbolConfiguration = (NSImageSymbolConfiguration*)ptr;
    NSImageSymbolConfiguration* result_ = [nSImageSymbolConfiguration autorelease];
    return result_;
}

void* C_NSImageSymbolConfiguration_Retain(void* ptr) {
    NSImageSymbolConfiguration* nSImageSymbolConfiguration = (NSImageSymbolConfiguration*)ptr;
    NSImageSymbolConfiguration* result_ = [nSImageSymbolConfiguration retain];
    return result_;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithPointSize_Weight(double pointSize, double weight) {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration configurationWithPointSize:pointSize weight:weight];
    return result_;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithPointSize_Weight_Scale(double pointSize, double weight, int scale) {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration configurationWithPointSize:pointSize weight:weight scale:scale];
    return result_;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithTextStyle(void* style) {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration configurationWithTextStyle:(NSString*)style];
    return result_;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithTextStyle_Scale(void* style, int scale) {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration configurationWithTextStyle:(NSString*)style scale:scale];
    return result_;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithScale(int scale) {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration configurationWithScale:scale];
    return result_;
}
