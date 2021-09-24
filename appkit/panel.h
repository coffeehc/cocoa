#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Panel_Alloc();

void* C_NSPanel_Panel_WindowWithContentViewController(void* contentViewController);
void* C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag);
void* C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen);
void* C_NSPanel_Init(void* ptr);
void* C_NSPanel_AllocPanel();
void* C_NSPanel_NewPanel();
void* C_NSPanel_Autorelease(void* ptr);
void* C_NSPanel_Retain(void* ptr);
void C_NSPanel_SetFloatingPanel(void* ptr, bool value);
bool C_NSPanel_BecomesKeyOnlyIfNeeded(void* ptr);
void C_NSPanel_SetBecomesKeyOnlyIfNeeded(void* ptr, bool value);
void C_NSPanel_SetWorksWhenModal(void* ptr, bool value);
