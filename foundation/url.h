#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_URL_Alloc();

void* C_NSURL_InitWithString(void* ptr, void* URLString);
void* C_NSURL_InitWithString_RelativeToURL(void* ptr, void* URLString, void* baseURL);
void* C_NSURL_InitFileURLWithPath_IsDirectory(void* ptr, void* path, bool isDir);
void* C_NSURL_InitFileURLWithPath_RelativeToURL(void* ptr, void* path, void* baseURL);
void* C_NSURL_InitFileURLWithPath_IsDirectory_RelativeToURL(void* ptr, void* path, bool isDir, void* baseURL);
void* C_NSURL_InitFileURLWithPath(void* ptr, void* path);
void* C_NSURL_InitAbsoluteURLWithDataRepresentation_RelativeToURL(void* ptr, void* data, void* baseURL);
void* C_NSURL_InitWithDataRepresentation_RelativeToURL(void* ptr, void* data, void* baseURL);
void* C_NSURL_InitWithScheme_Host_Path(void* ptr, void* scheme, void* host, void* path);
void* C_NSURL_Init(void* ptr);
void* C_NSURL_URLWithString(void* URLString);
void* C_NSURL_URLWithString_RelativeToURL(void* URLString, void* baseURL);
void* C_NSURL_FileURLWithPath_IsDirectory(void* path, bool isDir);
void* C_NSURL_FileURLWithPath_RelativeToURL(void* path, void* baseURL);
void* C_NSURL_FileURLWithPath_IsDirectory_RelativeToURL(void* path, bool isDir, void* baseURL);
void* C_NSURL_FileURLWithPath(void* path);
void* C_NSURL_FileURLWithPathComponents(Array components);
void* C_NSURL_AbsoluteURLWithDataRepresentation_RelativeToURL(void* data, void* baseURL);
void* C_NSURL_URLWithDataRepresentation_RelativeToURL(void* data, void* baseURL);
bool C_NSURL_IsFileReferenceURL(void* ptr);
void C_NSURL_RemoveAllCachedResourceValues(void* ptr);
void C_NSURL_RemoveCachedResourceValueForKey(void* ptr, void* key);
void C_NSURL_SetTemporaryResourceValue_ForKey(void* ptr, void* value, void* key);
void* C_NSURL_FileReferenceURL(void* ptr);
void* C_NSURL_URLByAppendingPathComponent(void* ptr, void* pathComponent);
void* C_NSURL_URLByAppendingPathComponent_IsDirectory(void* ptr, void* pathComponent, bool isDirectory);
void* C_NSURL_URLByAppendingPathExtension(void* ptr, void* pathExtension);
Dictionary C_NSURL_URL_ResourceValuesForKeys_FromBookmarkData(Array keys, void* bookmarkData);
bool C_NSURL_StartAccessingSecurityScopedResource(void* ptr);
void C_NSURL_StopAccessingSecurityScopedResource(void* ptr);
void* C_NSURL_DataRepresentation(void* ptr);
bool C_NSURL_IsFileURL(void* ptr);
void* C_NSURL_AbsoluteString(void* ptr);
void* C_NSURL_AbsoluteURL(void* ptr);
void* C_NSURL_BaseURL(void* ptr);
void* C_NSURL_Fragment(void* ptr);
void* C_NSURL_Host(void* ptr);
void* C_NSURL_LastPathComponent(void* ptr);
void* C_NSURL_Password(void* ptr);
void* C_NSURL_Path(void* ptr);
Array C_NSURL_PathComponents(void* ptr);
void* C_NSURL_PathExtension(void* ptr);
void* C_NSURL_Port(void* ptr);
void* C_NSURL_Query(void* ptr);
void* C_NSURL_RelativePath(void* ptr);
void* C_NSURL_RelativeString(void* ptr);
void* C_NSURL_ResourceSpecifier(void* ptr);
void* C_NSURL_Scheme(void* ptr);
void* C_NSURL_StandardizedURL(void* ptr);
void* C_NSURL_User(void* ptr);
void* C_NSURL_FilePathURL(void* ptr);
void* C_NSURL_URLByDeletingLastPathComponent(void* ptr);
void* C_NSURL_URLByDeletingPathExtension(void* ptr);
void* C_NSURL_URLByResolvingSymlinksInPath(void* ptr);
void* C_NSURL_URLByStandardizingPath(void* ptr);
bool C_NSURL_HasDirectoryPath(void* ptr);
