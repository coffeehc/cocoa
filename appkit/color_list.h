#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ColorList_Alloc();

void* C_NSColorList_InitWithName(void* ptr, void* name);
void* C_NSColorList_InitWithName_FromFile(void* ptr, void* name, void* path);
void* C_NSColorList_AllocColorList();
void* C_NSColorList_Init(void* ptr);
void* C_NSColorList_NewColorList();
void* C_NSColorList_Autorelease(void* ptr);
void* C_NSColorList_Retain(void* ptr);
void* C_NSColorList_ColorListNamed(void* name);
void* C_NSColorList_ColorWithKey(void* ptr, void* key);
void C_NSColorList_InsertColor_Key_AtIndex(void* ptr, void* color, void* key, unsigned int loc);
void C_NSColorList_RemoveColorWithKey(void* ptr, void* key);
void C_NSColorList_SetColor_ForKey(void* ptr, void* color, void* key);
void C_NSColorList_RemoveFile(void* ptr);
Array C_NSColorList_AvailableColorLists();
void* C_NSColorList_Name(void* ptr);
bool C_NSColorList_IsEditable(void* ptr);
Array C_NSColorList_AllKeys(void* ptr);
