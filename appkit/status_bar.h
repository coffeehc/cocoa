#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool StatusBar_IsVertical(void* ptr);
double StatusBar_Thickness(void* ptr);
void* StatusBar_SystemStatusBar();

void* StatusBar_StatusItemWithLength(void* ptr, double length);
void StatusBar_RemoveStatusItem(void* ptr, void* item);
