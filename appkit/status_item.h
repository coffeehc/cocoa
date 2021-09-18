#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_StatusItem_Alloc();

void* C_NSStatusItem_AllocStatusItem();
void* C_NSStatusItem_Init(void* ptr);
void* C_NSStatusItem_NewStatusItem();
void* C_NSStatusItem_Autorelease(void* ptr);
void* C_NSStatusItem_Retain(void* ptr);
void* C_NSStatusItem_StatusBar(void* ptr);
unsigned int C_NSStatusItem_Behavior(void* ptr);
void C_NSStatusItem_SetBehavior(void* ptr, unsigned int value);
void* C_NSStatusItem_Button(void* ptr);
void* C_NSStatusItem_Menu(void* ptr);
void C_NSStatusItem_SetMenu(void* ptr, void* value);
bool C_NSStatusItem_IsVisible(void* ptr);
void C_NSStatusItem_SetVisible(void* ptr, bool value);
double C_NSStatusItem_Length(void* ptr);
void C_NSStatusItem_SetLength(void* ptr, double value);
void* C_NSStatusItem_AutosaveName(void* ptr);
void C_NSStatusItem_SetAutosaveName(void* ptr, void* value);
