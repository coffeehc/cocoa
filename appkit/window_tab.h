#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_WindowTab_Alloc();

void* C_NSWindowTab_AllocWindowTab();
void* C_NSWindowTab_Init(void* ptr);
void* C_NSWindowTab_NewWindowTab();
void* C_NSWindowTab_Autorelease(void* ptr);
void* C_NSWindowTab_Retain(void* ptr);
void* C_NSWindowTab_Title(void* ptr);
void C_NSWindowTab_SetTitle(void* ptr, void* value);
void* C_NSWindowTab_AttributedTitle(void* ptr);
void C_NSWindowTab_SetAttributedTitle(void* ptr, void* value);
void* C_NSWindowTab_ToolTip(void* ptr);
void C_NSWindowTab_SetToolTip(void* ptr, void* value);
void* C_NSWindowTab_AccessoryView(void* ptr);
void C_NSWindowTab_SetAccessoryView(void* ptr, void* value);
