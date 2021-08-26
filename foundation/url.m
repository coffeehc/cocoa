#import <Foundation/Foundation.h>
#import "url.h"

void* C_URL_Alloc() {
    return [NSURL alloc];
}

void* C_NSURL_InitWithString(void* ptr, void* URLString) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL initWithString:(NSString*)URLString];
    return result_;
}

void* C_NSURL_InitWithString_RelativeToURL(void* ptr, void* URLString, void* baseURL) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL initWithString:(NSString*)URLString relativeToURL:(NSURL*)baseURL];
    return result_;
}

void* C_NSURL_InitFileURLWithPath_IsDirectory(void* ptr, void* path, bool isDir) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL initFileURLWithPath:(NSString*)path isDirectory:isDir];
    return result_;
}

void* C_NSURL_InitFileURLWithPath_RelativeToURL(void* ptr, void* path, void* baseURL) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL initFileURLWithPath:(NSString*)path relativeToURL:(NSURL*)baseURL];
    return result_;
}

void* C_NSURL_InitFileURLWithPath_IsDirectory_RelativeToURL(void* ptr, void* path, bool isDir, void* baseURL) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL initFileURLWithPath:(NSString*)path isDirectory:isDir relativeToURL:(NSURL*)baseURL];
    return result_;
}

void* C_NSURL_InitFileURLWithPath(void* ptr, void* path) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL initFileURLWithPath:(NSString*)path];
    return result_;
}

void* C_NSURL_InitAbsoluteURLWithDataRepresentation_RelativeToURL(void* ptr, Array data, void* baseURL) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL initAbsoluteURLWithDataRepresentation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] relativeToURL:(NSURL*)baseURL];
    return result_;
}

void* C_NSURL_InitWithDataRepresentation_RelativeToURL(void* ptr, Array data, void* baseURL) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL initWithDataRepresentation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] relativeToURL:(NSURL*)baseURL];
    return result_;
}

void* C_NSURL_InitWithScheme_Host_Path(void* ptr, void* scheme, void* host, void* path) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL initWithScheme:(NSString*)scheme host:(NSString*)host path:(NSString*)path];
    return result_;
}

void* C_NSURL_Init(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL init];
    return result_;
}

void* C_NSURL_URLWithString(void* URLString) {
    NSURL* result_ = [NSURL URLWithString:(NSString*)URLString];
    return result_;
}

void* C_NSURL_URLWithString_RelativeToURL(void* URLString, void* baseURL) {
    NSURL* result_ = [NSURL URLWithString:(NSString*)URLString relativeToURL:(NSURL*)baseURL];
    return result_;
}

void* C_NSURL_URL_FileURLWithPath_IsDirectory(void* path, bool isDir) {
    NSURL* result_ = [NSURL fileURLWithPath:(NSString*)path isDirectory:isDir];
    return result_;
}

void* C_NSURL_URL_FileURLWithPath_RelativeToURL(void* path, void* baseURL) {
    NSURL* result_ = [NSURL fileURLWithPath:(NSString*)path relativeToURL:(NSURL*)baseURL];
    return result_;
}

void* C_NSURL_URL_FileURLWithPath_IsDirectory_RelativeToURL(void* path, bool isDir, void* baseURL) {
    NSURL* result_ = [NSURL fileURLWithPath:(NSString*)path isDirectory:isDir relativeToURL:(NSURL*)baseURL];
    return result_;
}

void* C_NSURL_URL_FileURLWithPath(void* path) {
    NSURL* result_ = [NSURL fileURLWithPath:(NSString*)path];
    return result_;
}

void* C_NSURL_URL_FileURLWithPathComponents(Array components) {
    NSMutableArray* objcComponents = [[NSMutableArray alloc] init];
    if (components.len > 0) {
    	void** componentsData = (void**)components.data;
    	for (int i = 0; i < components.len; i++) {
    		void* p = componentsData[i];
    		[objcComponents addObject:(NSString*)(NSString*)p];
    	}
    }
    NSURL* result_ = [NSURL fileURLWithPathComponents:objcComponents];
    return result_;
}

void* C_NSURL_URL_AbsoluteURLWithDataRepresentation_RelativeToURL(Array data, void* baseURL) {
    NSURL* result_ = [NSURL absoluteURLWithDataRepresentation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] relativeToURL:(NSURL*)baseURL];
    return result_;
}

void* C_NSURL_URLWithDataRepresentation_RelativeToURL(Array data, void* baseURL) {
    NSURL* result_ = [NSURL URLWithDataRepresentation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] relativeToURL:(NSURL*)baseURL];
    return result_;
}

bool C_NSURL_IsFileReferenceURL(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    BOOL result_ = [nSURL isFileReferenceURL];
    return result_;
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
    NSURL* result_ = [nSURL fileReferenceURL];
    return result_;
}

void* C_NSURL_URLByAppendingPathComponent(void* ptr, void* pathComponent) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL URLByAppendingPathComponent:(NSString*)pathComponent];
    return result_;
}

void* C_NSURL_URLByAppendingPathComponent_IsDirectory(void* ptr, void* pathComponent, bool isDirectory) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL URLByAppendingPathComponent:(NSString*)pathComponent isDirectory:isDirectory];
    return result_;
}

void* C_NSURL_URLByAppendingPathExtension(void* ptr, void* pathExtension) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL URLByAppendingPathExtension:(NSString*)pathExtension];
    return result_;
}

Dictionary C_NSURL_URL_ResourceValuesForKeys_FromBookmarkData(Array keys, Array bookmarkData) {
    NSMutableArray* objcKeys = [[NSMutableArray alloc] init];
    if (keys.len > 0) {
    	void** keysData = (void**)keys.data;
    	for (int i = 0; i < keys.len; i++) {
    		void* p = keysData[i];
    		[objcKeys addObject:(NSURLResourceKey)(NSString*)p];
    	}
    }
    NSDictionary* result_ = [NSURL resourceValuesForKeys:objcKeys fromBookmarkData:[[NSData alloc] initWithBytes:(Byte *)bookmarkData.data length:bookmarkData.len]];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSURLResourceKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

bool C_NSURL_StartAccessingSecurityScopedResource(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    BOOL result_ = [nSURL startAccessingSecurityScopedResource];
    return result_;
}

void C_NSURL_StopAccessingSecurityScopedResource(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    [nSURL stopAccessingSecurityScopedResource];
}

Array C_NSURL_DataRepresentation(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSData* result_ = [nSURL dataRepresentation];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

bool C_NSURL_IsFileURL(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    BOOL result_ = [nSURL isFileURL];
    return result_;
}

void* C_NSURL_AbsoluteString(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL absoluteString];
    return result_;
}

void* C_NSURL_AbsoluteURL(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL absoluteURL];
    return result_;
}

void* C_NSURL_BaseURL(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL baseURL];
    return result_;
}

void* C_NSURL_Fragment(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL fragment];
    return result_;
}

void* C_NSURL_Host(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL host];
    return result_;
}

void* C_NSURL_LastPathComponent(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL lastPathComponent];
    return result_;
}

void* C_NSURL_Password(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL password];
    return result_;
}

void* C_NSURL_Path(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL path];
    return result_;
}

Array C_NSURL_PathComponents(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSArray* result_ = [nSURL pathComponents];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void* C_NSURL_PathExtension(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL pathExtension];
    return result_;
}

void* C_NSURL_Port(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSNumber* result_ = [nSURL port];
    return result_;
}

void* C_NSURL_Query(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL query];
    return result_;
}

void* C_NSURL_RelativePath(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL relativePath];
    return result_;
}

void* C_NSURL_RelativeString(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL relativeString];
    return result_;
}

void* C_NSURL_ResourceSpecifier(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL resourceSpecifier];
    return result_;
}

void* C_NSURL_Scheme(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL scheme];
    return result_;
}

void* C_NSURL_StandardizedURL(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL standardizedURL];
    return result_;
}

void* C_NSURL_User(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSString* result_ = [nSURL user];
    return result_;
}

void* C_NSURL_FilePathURL(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL filePathURL];
    return result_;
}

void* C_NSURL_URLByDeletingLastPathComponent(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL URLByDeletingLastPathComponent];
    return result_;
}

void* C_NSURL_URLByDeletingPathExtension(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL URLByDeletingPathExtension];
    return result_;
}

void* C_NSURL_URLByResolvingSymlinksInPath(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL URLByResolvingSymlinksInPath];
    return result_;
}

void* C_NSURL_URLByStandardizingPath(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    NSURL* result_ = [nSURL URLByStandardizingPath];
    return result_;
}

bool C_NSURL_HasDirectoryPath(void* ptr) {
    NSURL* nSURL = (NSURL*)ptr;
    BOOL result_ = [nSURL hasDirectoryPath];
    return result_;
}
