#import <Appkit/Appkit.h>
#import "tint_configuration.h"

void* C_TintConfiguration_Alloc() {
    return [NSTintConfiguration alloc];
}

void* C_NSTintConfiguration_Init(void* ptr) {
    NSTintConfiguration* nSTintConfiguration = (NSTintConfiguration*)ptr;
    NSTintConfiguration* result_ = [nSTintConfiguration init];
    return result_;
}

void* C_NSTintConfiguration_TintConfigurationWithFixedColor(void* color) {
    NSTintConfiguration* result_ = [NSTintConfiguration tintConfigurationWithFixedColor:(NSColor*)color];
    return result_;
}

void* C_NSTintConfiguration_TintConfigurationWithPreferredColor(void* color) {
    NSTintConfiguration* result_ = [NSTintConfiguration tintConfigurationWithPreferredColor:(NSColor*)color];
    return result_;
}

bool C_NSTintConfiguration_AdaptsToUserAccentColor(void* ptr) {
    NSTintConfiguration* nSTintConfiguration = (NSTintConfiguration*)ptr;
    BOOL result_ = [nSTintConfiguration adaptsToUserAccentColor];
    return result_;
}

void* C_NSTintConfiguration_DefaultTintConfiguration() {
    NSTintConfiguration* result_ = [NSTintConfiguration defaultTintConfiguration];
    return result_;
}

void* C_NSTintConfiguration_MonochromeTintConfiguration() {
    NSTintConfiguration* result_ = [NSTintConfiguration monochromeTintConfiguration];
    return result_;
}

void* C_NSTintConfiguration_BaseTintColor(void* ptr) {
    NSTintConfiguration* nSTintConfiguration = (NSTintConfiguration*)ptr;
    NSColor* result_ = [nSTintConfiguration baseTintColor];
    return result_;
}

void* C_NSTintConfiguration_EquivalentContentTintColor(void* ptr) {
    NSTintConfiguration* nSTintConfiguration = (NSTintConfiguration*)ptr;
    NSColor* result_ = [nSTintConfiguration equivalentContentTintColor];
    return result_;
}
