#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Appearance_Alloc();

void* C_NSAppearance_InitWithAppearanceNamed_Bundle(void* ptr, void* name, void* bundle);
void* C_NSAppearance_InitWithCoder(void* ptr, void* coder);
void* C_NSAppearance_AllocAppearance();
void* C_NSAppearance_Init(void* ptr);
void* C_NSAppearance_NewAppearance();
void* C_NSAppearance_Autorelease(void* ptr);
void* C_NSAppearance_Retain(void* ptr);
void* C_NSAppearance_AppearanceNamed(void* name);
void* C_NSAppearance_BestMatchFromAppearancesWithNames(void* ptr, Array appearances);
void* C_NSAppearance_Name(void* ptr);
void* C_NSAppearance_CurrentDrawingAppearance();
bool C_NSAppearance_AllowsVibrancy(void* ptr);
