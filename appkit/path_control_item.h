#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_PathControlItem_Alloc();

void* C_NSPathControlItem_AllocPathControlItem();
void* C_NSPathControlItem_Init(void* ptr);
void* C_NSPathControlItem_NewPathControlItem();
void* C_NSPathControlItem_Autorelease(void* ptr);
void* C_NSPathControlItem_Retain(void* ptr);
void* C_NSPathControlItem_URL(void* ptr);
void* C_NSPathControlItem_AttributedTitle(void* ptr);
void C_NSPathControlItem_SetAttributedTitle(void* ptr, void* value);
void* C_NSPathControlItem_Image(void* ptr);
void C_NSPathControlItem_SetImage(void* ptr, void* value);
void* C_NSPathControlItem_Title(void* ptr);
void C_NSPathControlItem_SetTitle(void* ptr, void* value);
