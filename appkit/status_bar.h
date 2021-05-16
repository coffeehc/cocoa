#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_StatusBar_Alloc();

void* C_NSStatusBar_Init(void* ptr);
void* C_NSStatusBar_StatusItemWithLength(void* ptr, double length);
void C_NSStatusBar_RemoveStatusItem(void* ptr, void* item);
void* C_NSStatusBar_SystemStatusBar();
bool C_NSStatusBar_IsVertical(void* ptr);
double C_NSStatusBar_Thickness(void* ptr);
