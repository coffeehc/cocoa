#import <Foundation/Foundation.h>
#import "url.h"

void* C_URL_Alloc() {
	return [NSURL alloc];
}

void* C_NSURL_InitWithString(void* ptr, void* URLString) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL initWithString:(NSString*)URLString];
	return result;
}

void* C_NSURL_InitWithString_RelativeToURL(void* ptr, void* URLString, void* baseURL) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL initWithString:(NSString*)URLString relativeToURL:(NSURL*)baseURL];
	return result;
}

void* C_NSURL_InitFileURLWithPath_IsDirectory(void* ptr, void* path, bool isDir) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL initFileURLWithPath:(NSString*)path isDirectory:isDir];
	return result;
}

void* C_NSURL_InitFileURLWithPath_RelativeToURL(void* ptr, void* path, void* baseURL) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL initFileURLWithPath:(NSString*)path relativeToURL:(NSURL*)baseURL];
	return result;
}

void* C_NSURL_InitFileURLWithPath_IsDirectory_RelativeToURL(void* ptr, void* path, bool isDir, void* baseURL) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL initFileURLWithPath:(NSString*)path isDirectory:isDir relativeToURL:(NSURL*)baseURL];
	return result;
}

void* C_NSURL_InitFileURLWithPath(void* ptr, void* path) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL initFileURLWithPath:(NSString*)path];
	return result;
}

void* C_NSURL_InitAbsoluteURLWithDataRepresentation_RelativeToURL(void* ptr, Array data, void* baseURL) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL initAbsoluteURLWithDataRepresentation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] relativeToURL:(NSURL*)baseURL];
	return result;
}

void* C_NSURL_InitWithDataRepresentation_RelativeToURL(void* ptr, Array data, void* baseURL) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL initWithDataRepresentation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] relativeToURL:(NSURL*)baseURL];
	return result;
}

void* C_NSURL_Init(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL init];
	return result;
}

void* C_NSURL_URLWithString(void* URLString) {
	NSURL* result = [NSURL URLWithString:(NSString*)URLString];
	return result;
}

void* C_NSURL_URLWithString_RelativeToURL(void* URLString, void* baseURL) {
	NSURL* result = [NSURL URLWithString:(NSString*)URLString relativeToURL:(NSURL*)baseURL];
	return result;
}

void* C_NSURL_URLFileURLWithPath_IsDirectory(void* path, bool isDir) {
	NSURL* result = [NSURL fileURLWithPath:(NSString*)path isDirectory:isDir];
	return result;
}

void* C_NSURL_URLFileURLWithPath_RelativeToURL(void* path, void* baseURL) {
	NSURL* result = [NSURL fileURLWithPath:(NSString*)path relativeToURL:(NSURL*)baseURL];
	return result;
}

void* C_NSURL_URLFileURLWithPath_IsDirectory_RelativeToURL(void* path, bool isDir, void* baseURL) {
	NSURL* result = [NSURL fileURLWithPath:(NSString*)path isDirectory:isDir relativeToURL:(NSURL*)baseURL];
	return result;
}

void* C_NSURL_URLFileURLWithPath(void* path) {
	NSURL* result = [NSURL fileURLWithPath:(NSString*)path];
	return result;
}

void* C_NSURL_URLAbsoluteURLWithDataRepresentation_RelativeToURL(Array data, void* baseURL) {
	NSURL* result = [NSURL absoluteURLWithDataRepresentation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] relativeToURL:(NSURL*)baseURL];
	return result;
}

void* C_NSURL_URLWithDataRepresentation_RelativeToURL(Array data, void* baseURL) {
	NSURL* result = [NSURL URLWithDataRepresentation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] relativeToURL:(NSURL*)baseURL];
	return result;
}

bool C_NSURL_IsFileReferenceURL(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	bool result = [nSURL isFileReferenceURL];
	return result;
}

void C_NSURL_RemoveAllCachedResourceValues(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	[nSURL removeAllCachedResourceValues];
}

void C_NSURL_RemoveCachedResourceValueForKey(void* ptr, void* key) {
	NSURL* nSURL = (NSURL*)ptr;
	[nSURL removeCachedResourceValueForKey:(NSString*)key];
}

void C_NSURL_SetTemporaryResourceValue_ForKey(void* ptr, void* value, void* key) {
	NSURL* nSURL = (NSURL*)ptr;
	[nSURL setTemporaryResourceValue:(id)value forKey:(NSString*)key];
}

void* C_NSURL_FileReferenceURL(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL fileReferenceURL];
	return result;
}

void* C_NSURL_URLByAppendingPathComponent(void* ptr, void* pathComponent) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL URLByAppendingPathComponent:(NSString*)pathComponent];
	return result;
}

void* C_NSURL_URLByAppendingPathComponent_IsDirectory(void* ptr, void* pathComponent, bool isDirectory) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL URLByAppendingPathComponent:(NSString*)pathComponent isDirectory:isDirectory];
	return result;
}

void* C_NSURL_URLByAppendingPathExtension(void* ptr, void* pathExtension) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL URLByAppendingPathExtension:(NSString*)pathExtension];
	return result;
}

bool C_NSURL_StartAccessingSecurityScopedResource(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	bool result = [nSURL startAccessingSecurityScopedResource];
	return result;
}

void C_NSURL_StopAccessingSecurityScopedResource(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	[nSURL stopAccessingSecurityScopedResource];
}

Array C_NSURL_DataRepresentation(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSData* result = [nSURL dataRepresentation];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

bool C_NSURL_IsFileURL(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	bool result = [nSURL isFileURL];
	return result;
}

void* C_NSURL_AbsoluteString(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL absoluteString];
	return result;
}

void* C_NSURL_AbsoluteURL(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL absoluteURL];
	return result;
}

void* C_NSURL_BaseURL(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL baseURL];
	return result;
}

void* C_NSURL_Fragment(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL fragment];
	return result;
}

void* C_NSURL_Host(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL host];
	return result;
}

void* C_NSURL_LastPathComponent(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL lastPathComponent];
	return result;
}

void* C_NSURL_Password(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL password];
	return result;
}

void* C_NSURL_Path(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL path];
	return result;
}

void* C_NSURL_PathExtension(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL pathExtension];
	return result;
}

void* C_NSURL_Port(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSNumber* result = [nSURL port];
	return result;
}

void* C_NSURL_Query(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL query];
	return result;
}

void* C_NSURL_RelativePath(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL relativePath];
	return result;
}

void* C_NSURL_RelativeString(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL relativeString];
	return result;
}

void* C_NSURL_ResourceSpecifier(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL resourceSpecifier];
	return result;
}

void* C_NSURL_Scheme(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL scheme];
	return result;
}

void* C_NSURL_StandardizedURL(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL standardizedURL];
	return result;
}

void* C_NSURL_User(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSString* result = [nSURL user];
	return result;
}

void* C_NSURL_FilePathURL(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL filePathURL];
	return result;
}

void* C_NSURL_URLByDeletingLastPathComponent(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL URLByDeletingLastPathComponent];
	return result;
}

void* C_NSURL_URLByDeletingPathExtension(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL URLByDeletingPathExtension];
	return result;
}

void* C_NSURL_URLByResolvingSymlinksInPath(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL URLByResolvingSymlinksInPath];
	return result;
}

void* C_NSURL_URLByStandardizingPath(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	NSURL* result = [nSURL URLByStandardizingPath];
	return result;
}

bool C_NSURL_HasDirectoryPath(void* ptr) {
	NSURL* nSURL = (NSURL*)ptr;
	bool result = [nSURL hasDirectoryPath];
	return result;
}
