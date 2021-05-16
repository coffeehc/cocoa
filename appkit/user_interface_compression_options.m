#import <Appkit/Appkit.h>
#import "user_interface_compression_options.h"

void* C_UserInterfaceCompressionOptions_Alloc() {
    return [NSUserInterfaceCompressionOptions alloc];
}

void* C_NSUserInterfaceCompressionOptions_Init(void* ptr) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    NSUserInterfaceCompressionOptions* result = [nSUserInterfaceCompressionOptions init];
    return result;
}

void* C_NSUserInterfaceCompressionOptions_InitWithIdentifier(void* ptr, void* identifier) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    NSUserInterfaceCompressionOptions* result = [nSUserInterfaceCompressionOptions initWithIdentifier:(NSString*)identifier];
    return result;
}

void* C_NSUserInterfaceCompressionOptions_InitWithCoder(void* ptr, void* coder) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    NSUserInterfaceCompressionOptions* result = [nSUserInterfaceCompressionOptions initWithCoder:(NSCoder*)coder];
    return result;
}

bool C_NSUserInterfaceCompressionOptions_ContainsOptions(void* ptr, void* options) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    BOOL result = [nSUserInterfaceCompressionOptions containsOptions:(NSUserInterfaceCompressionOptions*)options];
    return result;
}

bool C_NSUserInterfaceCompressionOptions_IntersectsOptions(void* ptr, void* options) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    BOOL result = [nSUserInterfaceCompressionOptions intersectsOptions:(NSUserInterfaceCompressionOptions*)options];
    return result;
}

void* C_NSUserInterfaceCompressionOptions_OptionsByAddingOptions(void* ptr, void* options) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    NSUserInterfaceCompressionOptions* result = [nSUserInterfaceCompressionOptions optionsByAddingOptions:(NSUserInterfaceCompressionOptions*)options];
    return result;
}

void* C_NSUserInterfaceCompressionOptions_OptionsByRemovingOptions(void* ptr, void* options) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    NSUserInterfaceCompressionOptions* result = [nSUserInterfaceCompressionOptions optionsByRemovingOptions:(NSUserInterfaceCompressionOptions*)options];
    return result;
}

void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_HideImagesOption() {
    NSUserInterfaceCompressionOptions* result = [NSUserInterfaceCompressionOptions hideImagesOption];
    return result;
}

void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_HideTextOption() {
    NSUserInterfaceCompressionOptions* result = [NSUserInterfaceCompressionOptions hideTextOption];
    return result;
}

void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_ReduceMetricsOption() {
    NSUserInterfaceCompressionOptions* result = [NSUserInterfaceCompressionOptions reduceMetricsOption];
    return result;
}

void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_BreakEqualWidthsOption() {
    NSUserInterfaceCompressionOptions* result = [NSUserInterfaceCompressionOptions breakEqualWidthsOption];
    return result;
}

void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_StandardOptions() {
    NSUserInterfaceCompressionOptions* result = [NSUserInterfaceCompressionOptions standardOptions];
    return result;
}

bool C_NSUserInterfaceCompressionOptions_IsEmpty(void* ptr) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    BOOL result = [nSUserInterfaceCompressionOptions isEmpty];
    return result;
}
