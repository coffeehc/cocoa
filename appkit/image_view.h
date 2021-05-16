#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ImageView_Alloc();

void* C_NSImageView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSImageView_InitWithCoder(void* ptr, void* coder);
void* C_NSImageView_Init(void* ptr);
void* C_NSImageView_ImageViewWithImage(void* image);
void* C_NSImageView_Image(void* ptr);
void C_NSImageView_SetImage(void* ptr, void* value);
unsigned int C_NSImageView_ImageFrameStyle(void* ptr);
void C_NSImageView_SetImageFrameStyle(void* ptr, unsigned int value);
unsigned int C_NSImageView_ImageAlignment(void* ptr);
void C_NSImageView_SetImageAlignment(void* ptr, unsigned int value);
unsigned int C_NSImageView_ImageScaling(void* ptr);
void C_NSImageView_SetImageScaling(void* ptr, unsigned int value);
bool C_NSImageView_Animates(void* ptr);
void C_NSImageView_SetAnimates(void* ptr, bool value);
bool C_NSImageView_IsEditable(void* ptr);
void C_NSImageView_SetEditable(void* ptr, bool value);
bool C_NSImageView_AllowsCutCopyPaste(void* ptr);
void C_NSImageView_SetAllowsCutCopyPaste(void* ptr, bool value);
void* C_NSImageView_ContentTintColor(void* ptr);
void C_NSImageView_SetContentTintColor(void* ptr, void* value);
void* C_NSImageView_SymbolConfiguration(void* ptr);
void C_NSImageView_SetSymbolConfiguration(void* ptr, void* value);
