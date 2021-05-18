#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Set_Alloc();

void* C_NSSet_InitWithSet(void* ptr, void* set);
void* C_NSSet_InitWithSet_CopyItems(void* ptr, void* set, bool flag);
void* C_NSSet_Init(void* ptr);
void* C_NSSet_InitWithCoder(void* ptr, void* coder);
void* C_NSSet_Set_();
void* C_NSSet_SetWithSet(void* set);
void* C_NSSet_SetByAddingObjectsFromSet(void* ptr, void* other);
void C_NSSet_MakeObjectsPerformSelector(void* ptr, void* aSelector);
void C_NSSet_MakeObjectsPerformSelector_WithObject(void* ptr, void* aSelector, void* argument);
bool C_NSSet_IsSubsetOfSet(void* ptr, void* otherSet);
bool C_NSSet_IntersectsSet(void* ptr, void* otherSet);
bool C_NSSet_IsEqualToSet(void* ptr, void* otherSet);
void* C_NSSet_ValueForKey(void* ptr, void* key);
void C_NSSet_SetValue_ForKey(void* ptr, void* value, void* key);
void C_NSSet_RemoveObserver_ForKeyPath(void* ptr, void* observer, void* keyPath);
void* C_NSSet_DescriptionWithLocale(void* ptr, void* locale);
unsigned int C_NSSet_Count(void* ptr);
void* C_NSSet_Description(void* ptr);
