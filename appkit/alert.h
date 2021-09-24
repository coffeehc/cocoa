#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Alert_Alloc();

void* C_NSAlert_AllocAlert();
void* C_NSAlert_Init(void* ptr);
void* C_NSAlert_NewAlert();
void* C_NSAlert_Autorelease(void* ptr);
void* C_NSAlert_Retain(void* ptr);
void* C_NSAlert_AlertWithError(void* error);
void C_NSAlert_Layout(void* ptr);
int C_NSAlert_RunModal(void* ptr);
void* C_NSAlert_AddButtonWithTitle(void* ptr, void* title);
unsigned int C_NSAlert_AlertStyle(void* ptr);
void C_NSAlert_SetAlertStyle(void* ptr, unsigned int value);
void* C_NSAlert_AccessoryView(void* ptr);
void C_NSAlert_SetAccessoryView(void* ptr, void* value);
bool C_NSAlert_ShowsHelp(void* ptr);
void C_NSAlert_SetShowsHelp(void* ptr, bool value);
void* C_NSAlert_HelpAnchor(void* ptr);
void C_NSAlert_SetHelpAnchor(void* ptr, void* value);
void* C_NSAlert_Delegate(void* ptr);
void C_NSAlert_SetDelegate(void* ptr, void* value);
void* C_NSAlert_SuppressionButton(void* ptr);
bool C_NSAlert_ShowsSuppressionButton(void* ptr);
void C_NSAlert_SetShowsSuppressionButton(void* ptr, bool value);
void* C_NSAlert_InformativeText(void* ptr);
void C_NSAlert_SetInformativeText(void* ptr, void* value);
void* C_NSAlert_MessageText(void* ptr);
void C_NSAlert_SetMessageText(void* ptr, void* value);
void* C_NSAlert_Icon(void* ptr);
void C_NSAlert_SetIcon(void* ptr, void* value);
Array C_NSAlert_Buttons(void* ptr);
void* C_NSAlert_Window(void* ptr);

void* WrapAlertDelegate(uintptr_t goID);
