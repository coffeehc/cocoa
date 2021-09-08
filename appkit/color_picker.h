#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ColorPicker_Alloc();

void* C_NSColorPicker_InitWithPickerMask_ColorPanel(void* ptr, unsigned int mask, void* owningColorPanel);
void* C_NSColorPicker_AllocColorPicker();
void* C_NSColorPicker_Init(void* ptr);
void* C_NSColorPicker_NewColorPicker();
void* C_NSColorPicker_Autorelease(void* ptr);
void* C_NSColorPicker_Retain(void* ptr);
void C_NSColorPicker_InsertNewButtonImage_In(void* ptr, void* newButtonImage, void* buttonCell);
void C_NSColorPicker_SetMode(void* ptr, int mode);
void C_NSColorPicker_AttachColorList(void* ptr, void* colorList);
void C_NSColorPicker_DetachColorList(void* ptr, void* colorList);
void C_NSColorPicker_ViewSizeChanged(void* ptr, void* sender);
void* C_NSColorPicker_ColorPanel(void* ptr);
void* C_NSColorPicker_ProvideNewButtonImage(void* ptr);
void* C_NSColorPicker_ButtonToolTip(void* ptr);
CGSize C_NSColorPicker_MinContentSize(void* ptr);
