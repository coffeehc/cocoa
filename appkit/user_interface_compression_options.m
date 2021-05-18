#import <Appkit/Appkit.h>
#import "user_interface_compression_options.h"

void* C_UserInterfaceCompressionOptions_Alloc() {
    return [NSUserInterfaceCompressionOptions alloc];
}

void* C_NSUserInterfaceCompressionOptions_Init(void* ptr) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    NSUserInterfaceCompressionOptions* result_ = [nSUserInterfaceCompressionOptions init];
    return result_;
}

void* C_NSUserInterfaceCompressionOptions_InitWithCompressionOptions(void* ptr, void* options) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    NSUserInterfaceCompressionOptions* result_ = [nSUserInterfaceCompressionOptions initWithCompressionOptions:(NSSet*)options];
    return result_;
}

void* C_NSUserInterfaceCompressionOptions_InitWithIdentifier(void* ptr, void* identifier) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    NSUserInterfaceCompressionOptions* result_ = [nSUserInterfaceCompressionOptions initWithIdentifier:(NSString*)identifier];
    return result_;
}

void* C_NSUserInterfaceCompressionOptions_InitWithCoder(void* ptr, void* coder) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    NSUserInterfaceCompressionOptions* result_ = [nSUserInterfaceCompressionOptions initWithCoder:(NSCoder*)coder];
    return result_;
}

bool C_NSUserInterfaceCompressionOptions_ContainsOptions(void* ptr, void* options) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    BOOL result_ = [nSUserInterfaceCompressionOptions containsOptions:(NSUserInterfaceCompressionOptions*)options];
    return result_;
}

bool C_NSUserInterfaceCompressionOptions_IntersectsOptions(void* ptr, void* options) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    BOOL result_ = [nSUserInterfaceCompressionOptions intersectsOptions:(NSUserInterfaceCompressionOptions*)options];
    return result_;
}

void* C_NSUserInterfaceCompressionOptions_OptionsByAddingOptions(void* ptr, void* options) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    NSUserInterfaceCompressionOptions* result_ = [nSUserInterfaceCompressionOptions optionsByAddingOptions:(NSUserInterfaceCompressionOptions*)options];
    return result_;
}

void* C_NSUserInterfaceCompressionOptions_OptionsByRemovingOptions(void* ptr, void* options) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    NSUserInterfaceCompressionOptions* result_ = [nSUserInterfaceCompressionOptions optionsByRemovingOptions:(NSUserInterfaceCompressionOptions*)options];
    return result_;
}

void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_HideImagesOption() {
    NSUserInterfaceCompressionOptions* result_ = [NSUserInterfaceCompressionOptions hideImagesOption];
    return result_;
}

void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_HideTextOption() {
    NSUserInterfaceCompressionOptions* result_ = [NSUserInterfaceCompressionOptions hideTextOption];
    return result_;
}

void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_ReduceMetricsOption() {
    NSUserInterfaceCompressionOptions* result_ = [NSUserInterfaceCompressionOptions reduceMetricsOption];
    return result_;
}

void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_BreakEqualWidthsOption() {
    NSUserInterfaceCompressionOptions* result_ = [NSUserInterfaceCompressionOptions breakEqualWidthsOption];
    return result_;
}

void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_StandardOptions() {
    NSUserInterfaceCompressionOptions* result_ = [NSUserInterfaceCompressionOptions standardOptions];
    return result_;
}

bool C_NSUserInterfaceCompressionOptions_IsEmpty(void* ptr) {
    NSUserInterfaceCompressionOptions* nSUserInterfaceCompressionOptions = (NSUserInterfaceCompressionOptions*)ptr;
    BOOL result_ = [nSUserInterfaceCompressionOptions isEmpty];
    return result_;
}
