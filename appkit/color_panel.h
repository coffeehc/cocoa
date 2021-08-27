#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ColorPanel_Alloc();

void* C_NSColorPanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag);
void* C_NSColorPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen);
void* C_NSColorPanel_Init(void* ptr);
void C_NSColorPanel_ColorPanel_SetPickerMode(int mode);
void C_NSColorPanel_ColorPanel_SetPickerMask(unsigned int mask);
void C_NSColorPanel_SetAction(void* ptr, void* selector);
void C_NSColorPanel_SetTarget(void* ptr, void* target);
void C_NSColorPanel_AttachColorList(void* ptr, void* colorList);
void C_NSColorPanel_DetachColorList(void* ptr, void* colorList);
bool C_NSColorPanel_ColorPanel_DragColor_WithEvent_FromView(void* color, void* event, void* sourceView);
void* C_NSColorPanel_SharedColorPanel();
bool C_NSColorPanel_SharedColorPanelExists();
int C_NSColorPanel_Mode(void* ptr);
void C_NSColorPanel_SetMode(void* ptr, int value);
void* C_NSColorPanel_AccessoryView(void* ptr);
void C_NSColorPanel_SetAccessoryView(void* ptr, void* value);
bool C_NSColorPanel_IsContinuous(void* ptr);
void C_NSColorPanel_SetContinuous(void* ptr, bool value);
bool C_NSColorPanel_ShowsAlpha(void* ptr);
void C_NSColorPanel_SetShowsAlpha(void* ptr, bool value);
void* C_NSColorPanel_Color(void* ptr);
void C_NSColorPanel_SetColor(void* ptr, void* value);
double C_NSColorPanel_Alpha(void* ptr);
