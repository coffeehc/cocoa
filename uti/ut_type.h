#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

const char* UTType_Identifier(void* ptr);
const char* UTType_PreferredFilenameExtension(void* ptr);
const char* UTType_PreferredMIMEType(void* ptr);
bool UTType_IsDeclared(void* ptr);
bool UTType_IsDynamic(void* ptr);
bool UTType_IsPublicType(void* ptr);
void* UTType_ReferenceURL(void* ptr);

Array UTType_typesWithTagConformingToType(const char* tag, const char* tagClass, void* supertype);
void* UTType_ExportedTypeWithIdentifier(const char* identifier);
void* UTType_exportedTypeWithIdentifierConformingToType(const char* identifier, void* parentType);
void* UTType_ImportedTypeWithIdentifier(const char* identifier);
void* UTType_importedTypeWithIdentifierConformingToType(const char* identifier, void* parentType);
void* UTType_TypeWithIdentifier(const char* identifier);
void* UTType_TypeWithFilenameExtension(const char* filenameExtension);
void* UTType_typeWithFilenameExtensionConformingToType(const char* filenameExtension, void* supertype);
void* UTType_TypeWithMIMEType(const char* mimeType);
void* UTType_typeWithMIMETypeConformingToType(const char* mimeType, void* supertype);
void* UTType_typeWithTagConformingToType(const char* tag, const char* tagClass, void* supertype);
bool UTType_ConformsToType(void* ptr, void* type_);
bool UTType_IsSubtypeOfType(void* ptr, void* type_);
bool UTType_IsSupertypeOfType(void* ptr, void* type_);
