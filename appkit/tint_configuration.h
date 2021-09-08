#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TintConfiguration_Alloc();

void* C_NSTintConfiguration_TintConfigurationWithFixedColor(void* color);
void* C_NSTintConfiguration_TintConfigurationWithPreferredColor(void* color);
void* C_NSTintConfiguration_AllocTintConfiguration();
void* C_NSTintConfiguration_Init(void* ptr);
void* C_NSTintConfiguration_NewTintConfiguration();
void* C_NSTintConfiguration_Autorelease(void* ptr);
void* C_NSTintConfiguration_Retain(void* ptr);
bool C_NSTintConfiguration_AdaptsToUserAccentColor(void* ptr);
void* C_NSTintConfiguration_DefaultTintConfiguration();
void* C_NSTintConfiguration_MonochromeTintConfiguration();
void* C_NSTintConfiguration_BaseTintColor(void* ptr);
void* C_NSTintConfiguration_EquivalentContentTintColor(void* ptr);
