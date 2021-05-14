#import <Foundation/Foundation.h>
#import "class_description.h"

void* C_ClassDescription_Alloc() {
	return [NSClassDescription alloc];
}

void* C_NSClassDescription_Init(void* ptr) {
	NSClassDescription* nSClassDescription = (NSClassDescription*)ptr;
	NSClassDescription* result = [nSClassDescription init];
	return result;
}

void C_NSClassDescription_ClassDescriptionInvalidateClassDescriptionCache() {
	[NSClassDescription invalidateClassDescriptionCache];
}

void* C_NSClassDescription_InverseForRelationshipKey(void* ptr, void* relationshipKey) {
	NSClassDescription* nSClassDescription = (NSClassDescription*)ptr;
	NSString* result = [nSClassDescription inverseForRelationshipKey:(NSString*)relationshipKey];
	return result;
}
