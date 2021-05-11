#import <Foundation/Foundation.h>
#import "file_wrapper.h"

void* C_FileWrapper_Alloc() {
	return [NSFileWrapper alloc];
}

void* C_NSFileWrapper_InitRegularFileWithContents(void* ptr, Array contents) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSFileWrapper* result = [nSFileWrapper initRegularFileWithContents:[[NSData alloc] initWithBytes:(Byte *)contents.data length:contents.len]];
	return result;
}

void* C_NSFileWrapper_InitSymbolicLinkWithDestinationURL(void* ptr, void* url) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSFileWrapper* result = [nSFileWrapper initSymbolicLinkWithDestinationURL:(NSURL*)url];
	return result;
}

void* C_NSFileWrapper_InitWithSerializedRepresentation(void* ptr, Array serializeRepresentation) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSFileWrapper* result = [nSFileWrapper initWithSerializedRepresentation:[[NSData alloc] initWithBytes:(Byte *)serializeRepresentation.data length:serializeRepresentation.len]];
	return result;
}

void* C_NSFileWrapper_InitWithCoder(void* ptr, void* inCoder) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSFileWrapper* result = [nSFileWrapper initWithCoder:(NSCoder*)inCoder];
	return result;
}

void* C_NSFileWrapper_Init(void* ptr) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSFileWrapper* result = [nSFileWrapper init];
	return result;
}

void* C_NSFileWrapper_AddFileWrapper(void* ptr, void* child) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSString* result = [nSFileWrapper addFileWrapper:(NSFileWrapper*)child];
	return result;
}

void C_NSFileWrapper_RemoveFileWrapper(void* ptr, void* child) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	[nSFileWrapper removeFileWrapper:(NSFileWrapper*)child];
}

void* C_NSFileWrapper_AddRegularFileWithContents_PreferredFilename(void* ptr, Array data, void* fileName) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSString* result = [nSFileWrapper addRegularFileWithContents:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] preferredFilename:(NSString*)fileName];
	return result;
}

void* C_NSFileWrapper_KeyForFileWrapper(void* ptr, void* child) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSString* result = [nSFileWrapper keyForFileWrapper:(NSFileWrapper*)child];
	return result;
}

bool C_NSFileWrapper_MatchesContentsOfURL(void* ptr, void* url) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	bool result = [nSFileWrapper matchesContentsOfURL:(NSURL*)url];
	return result;
}

bool C_NSFileWrapper_IsRegularFile(void* ptr) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	bool result = [nSFileWrapper isRegularFile];
	return result;
}

bool C_NSFileWrapper_IsDirectory(void* ptr) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	bool result = [nSFileWrapper isDirectory];
	return result;
}

bool C_NSFileWrapper_IsSymbolicLink(void* ptr) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	bool result = [nSFileWrapper isSymbolicLink];
	return result;
}

void* C_NSFileWrapper_SymbolicLinkDestinationURL(void* ptr) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSURL* result = [nSFileWrapper symbolicLinkDestinationURL];
	return result;
}

Array C_NSFileWrapper_SerializedRepresentation(void* ptr) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSData* result = [nSFileWrapper serializedRepresentation];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

void* C_NSFileWrapper_Filename(void* ptr) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSString* result = [nSFileWrapper filename];
	return result;
}

void C_NSFileWrapper_SetFilename(void* ptr, void* value) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	[nSFileWrapper setFilename:(NSString*)value];
}

void* C_NSFileWrapper_PreferredFilename(void* ptr) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSString* result = [nSFileWrapper preferredFilename];
	return result;
}

void C_NSFileWrapper_SetPreferredFilename(void* ptr, void* value) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	[nSFileWrapper setPreferredFilename:(NSString*)value];
}

Array C_NSFileWrapper_RegularFileContents(void* ptr) {
	NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
	NSData* result = [nSFileWrapper regularFileContents];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}
