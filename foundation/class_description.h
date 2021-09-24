#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_ClassDescription_Alloc();

void* C_NSClassDescription_AllocClassDescription();
void* C_NSClassDescription_Init(void* ptr);
void* C_NSClassDescription_NewClassDescription();
void* C_NSClassDescription_Autorelease(void* ptr);
void* C_NSClassDescription_Retain(void* ptr);
void C_NSClassDescription_InvalidateClassDescriptionCache();
void* C_NSClassDescription_InverseForRelationshipKey(void* ptr, void* relationshipKey);
Array C_NSClassDescription_AttributeKeys(void* ptr);
Array C_NSClassDescription_ToManyRelationshipKeys(void* ptr);
Array C_NSClassDescription_ToOneRelationshipKeys(void* ptr);
