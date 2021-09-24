#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_DraggingItem_Alloc();

void* C_NSDraggingItem_InitWithPasteboardWriter(void* ptr, void* pasteboardWriter);
void* C_NSDraggingItem_AllocDraggingItem();
void* C_NSDraggingItem_Autorelease(void* ptr);
void* C_NSDraggingItem_Retain(void* ptr);
void C_NSDraggingItem_SetDraggingFrame_Contents(void* ptr, CGRect frame, void* contents);
CGRect C_NSDraggingItem_DraggingFrame(void* ptr);
void C_NSDraggingItem_SetDraggingFrame(void* ptr, CGRect value);
Array C_NSDraggingItem_ImageComponents(void* ptr);
void* C_NSDraggingItem_Item(void* ptr);

void* C_DraggingImageComponent_Alloc();

void* C_NSDraggingImageComponent_InitWithKey(void* ptr, void* key);
void* C_NSDraggingImageComponent_AllocDraggingImageComponent();
void* C_NSDraggingImageComponent_Autorelease(void* ptr);
void* C_NSDraggingImageComponent_Retain(void* ptr);
void* C_NSDraggingImageComponent_DraggingImageComponentWithKey(void* key);
void* C_NSDraggingImageComponent_Key(void* ptr);
void C_NSDraggingImageComponent_SetKey(void* ptr, void* value);
void* C_NSDraggingImageComponent_Contents(void* ptr);
void C_NSDraggingImageComponent_SetContents(void* ptr, void* value);
CGRect C_NSDraggingImageComponent_Frame(void* ptr);
void C_NSDraggingImageComponent_SetFrame(void* ptr, CGRect value);
