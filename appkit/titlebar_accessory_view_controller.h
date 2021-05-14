#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TitlebarAccessoryViewController_Alloc();

void* C_NSTitlebarAccessoryViewController_InitWithNibName_Bundle(void* ptr, void* nibNameOrNil, void* nibBundleOrNil);
void* C_NSTitlebarAccessoryViewController_InitWithCoder(void* ptr, void* coder);
void* C_NSTitlebarAccessoryViewController_Init(void* ptr);
double C_NSTitlebarAccessoryViewController_FullScreenMinHeight(void* ptr);
void C_NSTitlebarAccessoryViewController_SetFullScreenMinHeight(void* ptr, double value);
int C_NSTitlebarAccessoryViewController_LayoutAttribute(void* ptr);
void C_NSTitlebarAccessoryViewController_SetLayoutAttribute(void* ptr, int value);
bool C_NSTitlebarAccessoryViewController_AutomaticallyAdjustsSize(void* ptr);
void C_NSTitlebarAccessoryViewController_SetAutomaticallyAdjustsSize(void* ptr, bool value);
bool C_NSTitlebarAccessoryViewController_IsHidden(void* ptr);
void C_NSTitlebarAccessoryViewController_SetHidden(void* ptr, bool value);
