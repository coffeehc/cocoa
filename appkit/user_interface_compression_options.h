#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_UserInterfaceCompressionOptions_Alloc();

void* C_NSUserInterfaceCompressionOptions_Init(void* ptr);
void* C_NSUserInterfaceCompressionOptions_InitWithIdentifier(void* ptr, void* identifier);
void* C_NSUserInterfaceCompressionOptions_InitWithCoder(void* ptr, void* coder);
bool C_NSUserInterfaceCompressionOptions_ContainsOptions(void* ptr, void* options);
bool C_NSUserInterfaceCompressionOptions_IntersectsOptions(void* ptr, void* options);
void* C_NSUserInterfaceCompressionOptions_OptionsByAddingOptions(void* ptr, void* options);
void* C_NSUserInterfaceCompressionOptions_OptionsByRemovingOptions(void* ptr, void* options);
void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_HideImagesOption();
void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_HideTextOption();
void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_ReduceMetricsOption();
void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_BreakEqualWidthsOption();
void* C_NSUserInterfaceCompressionOptions_UserInterfaceCompressionOptions_StandardOptions();
bool C_NSUserInterfaceCompressionOptions_IsEmpty(void* ptr);
