#import <Foundation/Foundation.h>
#import "url.h"

bool URL_IsFileURL(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [uRL isFileURL];
}

const char* URL_AbsoluteString(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL absoluteString] UTF8String];
}

void* URL_AbsoluteURL(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [uRL absoluteURL];
}

void* URL_BaseURL(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [uRL baseURL];
}

const char* URL_Fragment(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL fragment] UTF8String];
}

const char* URL_Host(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL host] UTF8String];
}

const char* URL_LastPathComponent(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL lastPathComponent] UTF8String];
}

const char* URL_Password(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL password] UTF8String];
}

const char* URL_Path(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL path] UTF8String];
}

Array URL_PathComponents(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	NSArray* ns_array = [uRL pathComponents];
	int count = [ns_array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = (void*)[[ns_array objectAtIndex:i] UTF8String];}
	Array array;
	array.data = data;
	array.len = count;
	return array;
}

const char* URL_PathExtension(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL pathExtension] UTF8String];
}

void* URL_Port(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [uRL port];
}

const char* URL_Query(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL query] UTF8String];
}

const char* URL_RelativePath(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL relativePath] UTF8String];
}

const char* URL_RelativeString(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL relativeString] UTF8String];
}

const char* URL_ResourceSpecifier(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL resourceSpecifier] UTF8String];
}

const char* URL_Scheme(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL scheme] UTF8String];
}

void* URL_StandardizedURL(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [uRL standardizedURL];
}

const char* URL_User(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [[uRL user] UTF8String];
}

void* URL_FilePathURL(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [uRL filePathURL];
}

void* URL_URLWithString(const char* URLString) {
	return [NSURL URLWithString:[NSString stringWithUTF8String:URLString]];
}

void* URL_URLWithStringRelativeToURL(const char* URLString, void* baseURL) {
	return [NSURL URLWithString:[NSString stringWithUTF8String:URLString] relativeToURL:(NSURL*)baseURL];
}

void* URL_FileURLWithPath(const char* path) {
	return [NSURL fileURLWithPath:[NSString stringWithUTF8String:path]];
}

void* URL_FileURLWithPath2(const char* path, bool isDir) {
	return [NSURL fileURLWithPath:[NSString stringWithUTF8String:path] isDirectory:isDir];
}

void* URL_FileURLWithPathRelativeToURL(const char* path, void* baseURL) {
	return [NSURL fileURLWithPath:[NSString stringWithUTF8String:path] relativeToURL:(NSURL*)baseURL];
}

void* URL_FileURLWithPathRelativeToURL2(const char* path, bool isDir, void* baseURL) {
	return [NSURL fileURLWithPath:[NSString stringWithUTF8String:path] isDirectory:isDir relativeToURL:(NSURL*)baseURL];
}

void* URL_FileURLWithPathComponents(Array components) {
    NSMutableArray* objc_components = [[NSMutableArray alloc] init];
    void** componentsData = (void**)components.data;
    for (int i = 0; i < components.len; i++) {
    	[objc_components addObject:[NSString stringWithUTF8String:(const char*)componentsData[i]]];
    }
	return [NSURL fileURLWithPathComponents:objc_components];
}

bool URL_IsEqual(void* ptr, void* anObject) {
	NSURL* uRL = (NSURL*)ptr;
	return [uRL isEqual:(NSObject*)anObject];
}

bool URL_IsFileReferenceURL(void* ptr) {
	NSURL* uRL = (NSURL*)ptr;
	return [uRL isFileReferenceURL];
}
