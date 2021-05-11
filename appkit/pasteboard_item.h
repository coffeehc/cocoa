#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_PasteboardItem_Alloc();

void* C_NSPasteboardItem_Init(void* ptr);
bool C_NSPasteboardItem_SetData_ForType(void* ptr, Array data, void* _type);
bool C_NSPasteboardItem_SetString_ForType(void* ptr, void* _string, void* _type);
bool C_NSPasteboardItem_SetPropertyList_ForType(void* ptr, void* propertyList, void* _type);
Array C_NSPasteboardItem_DataForType(void* ptr, void* _type);
void* C_NSPasteboardItem_StringForType(void* ptr, void* _type);
void* C_NSPasteboardItem_PropertyListForType(void* ptr, void* _type);
