#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* StatusItem_StatusBar(void* ptr);
unsigned long StatusItem_Behavior(void* ptr);
void StatusItem_SetBehavior(void* ptr, unsigned long behavior);
void* StatusItem_Button(void* ptr);
void* StatusItem_Menu(void* ptr);
void StatusItem_SetMenu(void* ptr, void* menu);
bool StatusItem_IsVisible(void* ptr);
void StatusItem_SetVisible(void* ptr, bool visible);
double StatusItem_Length(void* ptr);
void StatusItem_SetLength(void* ptr, double length);
const char* StatusItem_AutosaveName(void* ptr);
void StatusItem_SetAutosaveName(void* ptr, const char* autosaveName);

