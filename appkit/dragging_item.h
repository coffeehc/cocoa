#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

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
