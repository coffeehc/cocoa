#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_StatusBar_Alloc();

void* C_NSStatusBar_AllocStatusBar();
void* C_NSStatusBar_Init(void* ptr);
void* C_NSStatusBar_NewStatusBar();
void* C_NSStatusBar_Autorelease(void* ptr);
void* C_NSStatusBar_Retain(void* ptr);
void* C_NSStatusBar_StatusItemWithLength(void* ptr, double length);
void C_NSStatusBar_RemoveStatusItem(void* ptr, void* item);
void* C_NSStatusBar_SystemStatusBar();
bool C_NSStatusBar_IsVertical(void* ptr);
double C_NSStatusBar_Thickness(void* ptr);
