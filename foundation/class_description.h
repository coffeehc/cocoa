#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_ClassDescription_Alloc();

void* C_NSClassDescription_Init(void* ptr);
void C_NSClassDescription_ClassDescriptionInvalidateClassDescriptionCache();
void* C_NSClassDescription_InverseForRelationshipKey(void* ptr, void* relationshipKey);
