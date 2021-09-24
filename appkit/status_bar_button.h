#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_StatusBarButton_Alloc();

void* C_NSStatusBarButton_StatusBarButton_CheckboxWithTitle_Target_Action(void* title, void* target, void* action);
void* C_NSStatusBarButton_StatusBarButton_ButtonWithImage_Target_Action(void* image, void* target, void* action);
void* C_NSStatusBarButton_StatusBarButton_RadioButtonWithTitle_Target_Action(void* title, void* target, void* action);
void* C_NSStatusBarButton_StatusBarButton_ButtonWithTitle_Image_Target_Action(void* title, void* image, void* target, void* action);
void* C_NSStatusBarButton_StatusBarButton_ButtonWithTitle_Target_Action(void* title, void* target, void* action);
void* C_NSStatusBarButton_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSStatusBarButton_InitWithCoder(void* ptr, void* coder);
void* C_NSStatusBarButton_Init(void* ptr);
void* C_NSStatusBarButton_AllocStatusBarButton();
void* C_NSStatusBarButton_NewStatusBarButton();
void* C_NSStatusBarButton_Autorelease(void* ptr);
void* C_NSStatusBarButton_Retain(void* ptr);
bool C_NSStatusBarButton_AppearsDisabled(void* ptr);
void C_NSStatusBarButton_SetAppearsDisabled(void* ptr, bool value);
