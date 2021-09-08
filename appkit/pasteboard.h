#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Pasteboard_Alloc();

void* C_NSPasteboard_AllocPasteboard();
void* C_NSPasteboard_Init(void* ptr);
void* C_NSPasteboard_NewPasteboard();
void* C_NSPasteboard_Autorelease(void* ptr);
void* C_NSPasteboard_Retain(void* ptr);
void* C_NSPasteboard_PasteboardByFilteringData_OfType(void* data, void* _type);
void* C_NSPasteboard_PasteboardByFilteringFile(void* filename);
void* C_NSPasteboard_PasteboardByFilteringTypesInPasteboard(void* pboard);
void* C_NSPasteboard_PasteboardWithName(void* name);
void* C_NSPasteboard_PasteboardWithUniqueName();
int C_NSPasteboard_ClearContents(void* ptr);
bool C_NSPasteboard_SetData_ForType(void* ptr, void* data, void* dataType);
bool C_NSPasteboard_SetPropertyList_ForType(void* ptr, void* plist, void* dataType);
bool C_NSPasteboard_SetString_ForType(void* ptr, void* _string, void* dataType);
unsigned int C_NSPasteboard_IndexOfPasteboardItem(void* ptr, void* pasteboardItem);
void* C_NSPasteboard_DataForType(void* ptr, void* dataType);
void* C_NSPasteboard_PropertyListForType(void* ptr, void* dataType);
void* C_NSPasteboard_StringForType(void* ptr, void* dataType);
void* C_NSPasteboard_AvailableTypeFromArray(void* ptr, Array types);
bool C_NSPasteboard_CanReadItemWithDataConformingToTypes(void* ptr, Array types);
Array C_NSPasteboard_Pasteboard_TypesFilterableTo(void* _type);
int C_NSPasteboard_PrepareForNewContentsWithOptions(void* ptr, unsigned int options);
int C_NSPasteboard_DeclareTypes_Owner(void* ptr, Array newTypes, void* newOwner);
int C_NSPasteboard_AddTypes_Owner(void* ptr, Array newTypes, void* newOwner);
bool C_NSPasteboard_WriteFileContents(void* ptr, void* filename);
bool C_NSPasteboard_WriteFileWrapper(void* ptr, void* wrapper);
void* C_NSPasteboard_ReadFileContentsType_ToFile(void* ptr, void* _type, void* filename);
void* C_NSPasteboard_ReadFileWrapper(void* ptr);
void* C_NSPasteboard_GeneralPasteboard();
Array C_NSPasteboard_PasteboardItems(void* ptr);
Array C_NSPasteboard_Types(void* ptr);
void* C_NSPasteboard_Name(void* ptr);
int C_NSPasteboard_ChangeCount(void* ptr);
