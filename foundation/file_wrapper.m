#import "file_wrapper.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <Foundation/NSFileWrapper.h>

void* C_FileWrapper_Alloc() {
    return [NSFileWrapper alloc];
}

void* C_NSFileWrapper_InitDirectoryWithFileWrappers(void* ptr, Dictionary childrenByPreferredName) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSMutableDictionary* objcChildrenByPreferredName = [NSMutableDictionary dictionaryWithCapacity:childrenByPreferredName.len];
    if (childrenByPreferredName.len > 0) {
    	void** childrenByPreferredNameKeyData = (void**)childrenByPreferredName.key_data;
    	void** childrenByPreferredNameValueData = (void**)childrenByPreferredName.value_data;
    	for (int i = 0; i < childrenByPreferredName.len; i++) {
    		void* kp = childrenByPreferredNameKeyData[i];
    		void* vp = childrenByPreferredNameValueData[i];
    		[objcChildrenByPreferredName setObject:(NSString*)kp forKey:(NSFileWrapper*)vp];
    	}
    }
    NSFileWrapper* result_ = [nSFileWrapper initDirectoryWithFileWrappers:objcChildrenByPreferredName];
    return result_;
}

void* C_NSFileWrapper_InitRegularFileWithContents(void* ptr, void* contents) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper initRegularFileWithContents:(NSData*)contents];
    return result_;
}

void* C_NSFileWrapper_InitSymbolicLinkWithDestinationURL(void* ptr, void* url) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper initSymbolicLinkWithDestinationURL:(NSURL*)url];
    return result_;
}

void* C_NSFileWrapper_InitWithSerializedRepresentation(void* ptr, void* serializeRepresentation) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper initWithSerializedRepresentation:(NSData*)serializeRepresentation];
    return result_;
}

void* C_NSFileWrapper_InitWithCoder(void* ptr, void* inCoder) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper initWithCoder:(NSCoder*)inCoder];
    return result_;
}

void* C_NSFileWrapper_AllocFileWrapper() {
    NSFileWrapper* result_ = [NSFileWrapper alloc];
    return result_;
}

void* C_NSFileWrapper_Init(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper init];
    return result_;
}

void* C_NSFileWrapper_NewFileWrapper() {
    NSFileWrapper* result_ = [NSFileWrapper new];
    return result_;
}

void* C_NSFileWrapper_Autorelease(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper autorelease];
    return result_;
}

void* C_NSFileWrapper_Retain(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper retain];
    return result_;
}

void* C_NSFileWrapper_AddFileWrapper(void* ptr, void* child) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSString* result_ = [nSFileWrapper addFileWrapper:(NSFileWrapper*)child];
    return result_;
}

void C_NSFileWrapper_RemoveFileWrapper(void* ptr, void* child) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    [nSFileWrapper removeFileWrapper:(NSFileWrapper*)child];
}

void* C_NSFileWrapper_AddRegularFileWithContents_PreferredFilename(void* ptr, void* data, void* fileName) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSString* result_ = [nSFileWrapper addRegularFileWithContents:(NSData*)data preferredFilename:(NSString*)fileName];
    return result_;
}

void* C_NSFileWrapper_KeyForFileWrapper(void* ptr, void* child) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSString* result_ = [nSFileWrapper keyForFileWrapper:(NSFileWrapper*)child];
    return result_;
}

bool C_NSFileWrapper_MatchesContentsOfURL(void* ptr, void* url) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    BOOL result_ = [nSFileWrapper matchesContentsOfURL:(NSURL*)url];
    return result_;
}

bool C_NSFileWrapper_IsRegularFile(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    BOOL result_ = [nSFileWrapper isRegularFile];
    return result_;
}

bool C_NSFileWrapper_IsDirectory(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    BOOL result_ = [nSFileWrapper isDirectory];
    return result_;
}

bool C_NSFileWrapper_IsSymbolicLink(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    BOOL result_ = [nSFileWrapper isSymbolicLink];
    return result_;
}

Dictionary C_NSFileWrapper_FileWrappers(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSDictionary* result_ = [nSFileWrapper fileWrappers];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		NSFileWrapper* vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void* C_NSFileWrapper_SymbolicLinkDestinationURL(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSURL* result_ = [nSFileWrapper symbolicLinkDestinationURL];
    return result_;
}

void* C_NSFileWrapper_SerializedRepresentation(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSData* result_ = [nSFileWrapper serializedRepresentation];
    return result_;
}

void* C_NSFileWrapper_Filename(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSString* result_ = [nSFileWrapper filename];
    return result_;
}

void C_NSFileWrapper_SetFilename(void* ptr, void* value) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    [nSFileWrapper setFilename:(NSString*)value];
}

void* C_NSFileWrapper_PreferredFilename(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSString* result_ = [nSFileWrapper preferredFilename];
    return result_;
}

void C_NSFileWrapper_SetPreferredFilename(void* ptr, void* value) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    [nSFileWrapper setPreferredFilename:(NSString*)value];
}

Dictionary C_NSFileWrapper_FileAttributes(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSDictionary* result_ = [nSFileWrapper fileAttributes];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
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

void C_NSFileWrapper_SetFileAttributes(void* ptr, Dictionary value) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSMutableDictionary* objcValue = [NSMutableDictionary dictionaryWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueKeyData = (void**)value.key_data;
    	void** valueValueData = (void**)value.value_data;
    	for (int i = 0; i < value.len; i++) {
    		void* kp = valueKeyData[i];
    		void* vp = valueValueData[i];
    		[objcValue setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSFileWrapper setFileAttributes:objcValue];
}

void* C_NSFileWrapper_RegularFileContents(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSData* result_ = [nSFileWrapper regularFileContents];
    return result_;
}
