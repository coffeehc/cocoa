#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Panel_Alloc();

void* C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag);
void* C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen);
void* C_NSPanel_Init(void* ptr);
bool C_NSPanel_BecomesKeyOnlyIfNeeded(void* ptr);
void C_NSPanel_SetBecomesKeyOnlyIfNeeded(void* ptr, bool value);
