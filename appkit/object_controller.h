#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ObjectController_Alloc();

void* C_NSObjectController_InitWithContent(void* ptr, void* content);
void* C_NSObjectController_InitWithCoder(void* ptr, void* coder);
void* C_NSObjectController_Init(void* ptr);
void* C_NSObjectController_AllocObjectController();
void* C_NSObjectController_NewObjectController();
void* C_NSObjectController_Autorelease(void* ptr);
void* C_NSObjectController_Retain(void* ptr);
void C_NSObjectController_PrepareContent(void* ptr);
void* C_NSObjectController_NewObject(void* ptr);
void C_NSObjectController_AddObject(void* ptr, void* object);
void C_NSObjectController_RemoveObject(void* ptr, void* object);
void C_NSObjectController_Add(void* ptr, void* sender);
void C_NSObjectController_Remove(void* ptr, void* sender);
void C_NSObjectController_Fetch(void* ptr, void* sender);
bool C_NSObjectController_ValidateUserInterfaceItem(void* ptr, void* item);
void* C_NSObjectController_Content(void* ptr);
void C_NSObjectController_SetContent(void* ptr, void* value);
bool C_NSObjectController_AutomaticallyPreparesContent(void* ptr);
void C_NSObjectController_SetAutomaticallyPreparesContent(void* ptr, bool value);
bool C_NSObjectController_CanAdd(void* ptr);
bool C_NSObjectController_CanRemove(void* ptr);
bool C_NSObjectController_IsEditable(void* ptr);
void C_NSObjectController_SetEditable(void* ptr, bool value);
void* C_NSObjectController_EntityName(void* ptr);
void C_NSObjectController_SetEntityName(void* ptr, void* value);
bool C_NSObjectController_UsesLazyFetching(void* ptr);
void C_NSObjectController_SetUsesLazyFetching(void* ptr, bool value);
void* C_NSObjectController_FetchPredicate(void* ptr);
void C_NSObjectController_SetFetchPredicate(void* ptr, void* value);
Array C_NSObjectController_SelectedObjects(void* ptr);
void* C_NSObjectController_Selection(void* ptr);
