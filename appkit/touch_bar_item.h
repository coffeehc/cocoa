#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TouchBarItem_Alloc();

void* C_NSTouchBarItem_InitWithIdentifier(void* ptr, void* identifier);
void* C_NSTouchBarItem_InitWithCoder(void* ptr, void* coder);
void* C_NSTouchBarItem_AllocTouchBarItem();
void* C_NSTouchBarItem_Autorelease(void* ptr);
void* C_NSTouchBarItem_Retain(void* ptr);
void* C_NSTouchBarItem_Identifier(void* ptr);
float C_NSTouchBarItem_VisibilityPriority(void* ptr);
void C_NSTouchBarItem_SetVisibilityPriority(void* ptr, float value);
bool C_NSTouchBarItem_IsVisible(void* ptr);
void* C_NSTouchBarItem_CustomizationLabel(void* ptr);
void* C_NSTouchBarItem_ViewController(void* ptr);
void* C_NSTouchBarItem_View(void* ptr);
