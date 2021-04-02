#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* ImageView_Image(void* ptr);
void ImageView_SetImage(void* ptr, void* image);
unsigned long ImageView_ImageFrameStyle(void* ptr);
void ImageView_SetImageFrameStyle(void* ptr, unsigned long imageFrameStyle);
unsigned long ImageView_ImageAlignment(void* ptr);
void ImageView_SetImageAlignment(void* ptr, unsigned long imageAlignment);
unsigned long ImageView_ImageScaling(void* ptr);
void ImageView_SetImageScaling(void* ptr, unsigned long imageScaling);
bool ImageView_Animates(void* ptr);
void ImageView_SetAnimates(void* ptr, bool animates);
bool ImageView_IsEditable(void* ptr);
void ImageView_SetEditable(void* ptr, bool editable);
bool ImageView_AllowsCutCopyPaste(void* ptr);
void ImageView_SetAllowsCutCopyPaste(void* ptr, bool allowsCutCopyPaste);
void* ImageView_ContentTintColor(void* ptr);
void ImageView_SetContentTintColor(void* ptr, void* contentTintColor);
void* ImageView_SymbolConfiguration(void* ptr);
void ImageView_SetSymbolConfiguration(void* ptr, void* symbolConfiguration);

void ImageView_ImageViewWithImage(void* image);
