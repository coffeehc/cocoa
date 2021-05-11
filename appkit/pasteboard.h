#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Pasteboard_Alloc();

void* C_NSPasteboard_Init(void* ptr);
void* C_NSPasteboard_PasteboardByFilteringData_OfType(Array data, void* _type);
void* C_NSPasteboard_PasteboardByFilteringFile(void* filename);
void* C_NSPasteboard_PasteboardByFilteringTypesInPasteboard(void* pboard);
void* C_NSPasteboard_PasteboardWithName(void* name);
void* C_NSPasteboard_PasteboardWithUniqueName();
int C_NSPasteboard_ClearContents(void* ptr);
bool C_NSPasteboard_SetData_ForType(void* ptr, Array data, void* dataType);
bool C_NSPasteboard_SetPropertyList_ForType(void* ptr, void* plist, void* dataType);
bool C_NSPasteboard_SetString_ForType(void* ptr, void* _string, void* dataType);
unsigned int C_NSPasteboard_IndexOfPasteboardItem(void* ptr, void* pasteboardItem);
Array C_NSPasteboard_DataForType(void* ptr, void* dataType);
void* C_NSPasteboard_PropertyListForType(void* ptr, void* dataType);
void* C_NSPasteboard_StringForType(void* ptr, void* dataType);
bool C_NSPasteboard_WriteFileContents(void* ptr, void* filename);
bool C_NSPasteboard_WriteFileWrapper(void* ptr, void* wrapper);
void* C_NSPasteboard_ReadFileContentsType_ToFile(void* ptr, void* _type, void* filename);
void* C_NSPasteboard_ReadFileWrapper(void* ptr);
void* C_NSPasteboard_Name(void* ptr);
int C_NSPasteboard_ChangeCount(void* ptr);
