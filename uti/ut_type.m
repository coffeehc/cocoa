#import <UniformTypeIdentifiers/UTType.h>
#import "ut_type.h"

const char* UTType_Identifier(void* ptr) {
	UTType* uTType = (UTType*)ptr;
	return [[uTType identifier] UTF8String];
}

const char* UTType_PreferredFilenameExtension(void* ptr) {
	UTType* uTType = (UTType*)ptr;
	return [[uTType preferredFilenameExtension] UTF8String];
}

const char* UTType_PreferredMIMEType(void* ptr) {
	UTType* uTType = (UTType*)ptr;
	return [[uTType preferredMIMEType] UTF8String];
}

bool UTType_IsDeclared(void* ptr) {
	UTType* uTType = (UTType*)ptr;
	return [uTType isDeclared];
}

bool UTType_IsDynamic(void* ptr) {
	UTType* uTType = (UTType*)ptr;
	return [uTType isDynamic];
}

bool UTType_IsPublicType(void* ptr) {
	UTType* uTType = (UTType*)ptr;
	return [uTType isPublicType];
}

void* UTType_ReferenceURL(void* ptr) {
	UTType* uTType = (UTType*)ptr;
	return [uTType referenceURL];
}

Array UTType_typesWithTagConformingToType(const char* tag, const char* tagClass, void* supertype) {
	NSArray* ns_array = [UTType typesWithTag:[NSString stringWithUTF8String:tag] tagClass:[NSString stringWithUTF8String:tagClass] conformingToType:(UTType*)supertype];
	int count = [ns_array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = [ns_array objectAtIndex:i];
	}
	Array array;
	array.data = data;
	array.len = count;
	return array;
}

void* UTType_ExportedTypeWithIdentifier(const char* identifier) {
	return [UTType exportedTypeWithIdentifier:[NSString stringWithUTF8String:identifier]];
}

void* UTType_exportedTypeWithIdentifierConformingToType(const char* identifier, void* parentType) {
	return [UTType exportedTypeWithIdentifier:[NSString stringWithUTF8String:identifier] conformingToType:(UTType*)parentType];
}

void* UTType_ImportedTypeWithIdentifier(const char* identifier) {
	return [UTType importedTypeWithIdentifier:[NSString stringWithUTF8String:identifier]];
}

void* UTType_importedTypeWithIdentifierConformingToType(const char* identifier, void* parentType) {
	return [UTType importedTypeWithIdentifier:[NSString stringWithUTF8String:identifier] conformingToType:(UTType*)parentType];
}

void* UTType_TypeWithIdentifier(const char* identifier) {
	return [UTType typeWithIdentifier:[NSString stringWithUTF8String:identifier]];
}

void* UTType_TypeWithFilenameExtension(const char* filenameExtension) {
	return [UTType typeWithFilenameExtension:[NSString stringWithUTF8String:filenameExtension]];
}

void* UTType_typeWithFilenameExtensionConformingToType(const char* filenameExtension, void* supertype) {
	return [UTType typeWithFilenameExtension:[NSString stringWithUTF8String:filenameExtension] conformingToType:(UTType*)supertype];
}

void* UTType_TypeWithMIMEType(const char* mimeType) {
	return [UTType typeWithMIMEType:[NSString stringWithUTF8String:mimeType]];
}

void* UTType_typeWithMIMETypeConformingToType(const char* mimeType, void* supertype) {
	return [UTType typeWithMIMEType:[NSString stringWithUTF8String:mimeType] conformingToType:(UTType*)supertype];
}

void* UTType_typeWithTagConformingToType(const char* tag, const char* tagClass, void* supertype) {
	return [UTType typeWithTag:[NSString stringWithUTF8String:tag] tagClass:[NSString stringWithUTF8String:tagClass] conformingToType:(UTType*)supertype];
}

bool UTType_ConformsToType(void* ptr, void* type_) {
	UTType* uTType = (UTType*)ptr;
	return [uTType conformsToType:(UTType*)type_];
}

bool UTType_IsSubtypeOfType(void* ptr, void* type_) {
	UTType* uTType = (UTType*)ptr;
	return [uTType isSubtypeOfType:(UTType*)type_];
}

bool UTType_IsSupertypeOfType(void* ptr, void* type_) {
	UTType* uTType = (UTType*)ptr;
	return [uTType isSupertypeOfType:(UTType*)type_];
}
