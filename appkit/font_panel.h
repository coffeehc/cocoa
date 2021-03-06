#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_FontPanel_Alloc();

void* C_NSFontPanel_FontPanel_WindowWithContentViewController(void* contentViewController);
void* C_NSFontPanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag);
void* C_NSFontPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen);
void* C_NSFontPanel_Init(void* ptr);
void* C_NSFontPanel_AllocFontPanel();
void* C_NSFontPanel_NewFontPanel();
void* C_NSFontPanel_Autorelease(void* ptr);
void* C_NSFontPanel_Retain(void* ptr);
void C_NSFontPanel_ReloadDefaultFontFamilies(void* ptr);
void C_NSFontPanel_SetPanelFont_IsMultiple(void* ptr, void* fontObj, bool flag);
void* C_NSFontPanel_PanelConvertFont(void* ptr, void* fontObj);
void* C_NSFontPanel_SharedFontPanel();
bool C_NSFontPanel_SharedFontPanelExists();
bool C_NSFontPanel_IsEnabled(void* ptr);
void C_NSFontPanel_SetEnabled(void* ptr, bool value);
void* C_NSFontPanel_AccessoryView(void* ptr);
void C_NSFontPanel_SetAccessoryView(void* ptr, void* value);

void* WrapFontChanging(uintptr_t goID);
