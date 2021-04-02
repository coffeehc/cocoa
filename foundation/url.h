#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool URL_IsFileURL(void* ptr);
const char* URL_AbsoluteString(void* ptr);
void* URL_AbsoluteURL(void* ptr);
void* URL_BaseURL(void* ptr);
const char* URL_Fragment(void* ptr);
const char* URL_Host(void* ptr);
const char* URL_LastPathComponent(void* ptr);
const char* URL_Password(void* ptr);
const char* URL_Path(void* ptr);
Array URL_PathComponents(void* ptr);
const char* URL_PathExtension(void* ptr);
void* URL_Port(void* ptr);
const char* URL_Query(void* ptr);
const char* URL_RelativePath(void* ptr);
const char* URL_RelativeString(void* ptr);
const char* URL_ResourceSpecifier(void* ptr);
const char* URL_Scheme(void* ptr);
void* URL_StandardizedURL(void* ptr);
const char* URL_User(void* ptr);
void* URL_FilePathURL(void* ptr);

void* URL_URLWithString(const char* URLString);
void* URL_URLWithStringRelativeToURL(const char* URLString, void* baseURL);
void* URL_FileURLWithPath(const char* path);
void* URL_FileURLWithPath2(const char* path, bool isDir);
void* URL_FileURLWithPathRelativeToURL(const char* path, void* baseURL);
void* URL_FileURLWithPathRelativeToURL2(const char* path, bool isDir, void* baseURL);
void* URL_FileURLWithPathComponents(Array components);
bool URL_IsEqual(void* ptr, void* anObject);
bool URL_IsFileReferenceURL(void* ptr);
