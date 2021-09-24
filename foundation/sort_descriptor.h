#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_SortDescriptor_Alloc();

void* C_NSSortDescriptor_SortDescriptorWithKey_Ascending(void* key, bool ascending);
void* C_NSSortDescriptor_InitWithKey_Ascending(void* ptr, void* key, bool ascending);
void* C_NSSortDescriptor_SortDescriptorWithKey_Ascending_Selector(void* key, bool ascending, void* selector);
void* C_NSSortDescriptor_InitWithKey_Ascending_Selector(void* ptr, void* key, bool ascending, void* selector);
void* C_NSSortDescriptor_InitWithCoder(void* ptr, void* coder);
void* C_NSSortDescriptor_AllocSortDescriptor();
void* C_NSSortDescriptor_Init(void* ptr);
void* C_NSSortDescriptor_NewSortDescriptor();
void* C_NSSortDescriptor_Autorelease(void* ptr);
void* C_NSSortDescriptor_Retain(void* ptr);
int C_NSSortDescriptor_CompareObject_ToObject(void* ptr, void* object1, void* object2);
void C_NSSortDescriptor_AllowEvaluation(void* ptr);
bool C_NSSortDescriptor_Ascending(void* ptr);
void* C_NSSortDescriptor_Key(void* ptr);
void* C_NSSortDescriptor_Selector(void* ptr);
void* C_NSSortDescriptor_ReversedSortDescriptor(void* ptr);
