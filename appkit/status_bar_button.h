#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_StatusBarButton_Alloc();

void* C_NSStatusBarButton_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSStatusBarButton_InitWithCoder(void* ptr, void* coder);
void* C_NSStatusBarButton_Init(void* ptr);
bool C_NSStatusBarButton_AppearsDisabled(void* ptr);
void C_NSStatusBarButton_SetAppearsDisabled(void* ptr, bool value);
