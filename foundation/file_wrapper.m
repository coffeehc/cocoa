#import <Foundation/Foundation.h>
#import "file_wrapper.h"

void* C_FileWrapper_Alloc() {
    return [NSFileWrapper alloc];
}

void* C_NSFileWrapper_InitRegularFileWithContents(void* ptr, Array contents) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper initRegularFileWithContents:[[NSData alloc] initWithBytes:(Byte *)contents.data length:contents.len]];
    return result_;
}

void* C_NSFileWrapper_InitSymbolicLinkWithDestinationURL(void* ptr, void* url) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper initSymbolicLinkWithDestinationURL:(NSURL*)url];
    return result_;
}

void* C_NSFileWrapper_InitWithSerializedRepresentation(void* ptr, Array serializeRepresentation) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper initWithSerializedRepresentation:[[NSData alloc] initWithBytes:(Byte *)serializeRepresentation.data length:serializeRepresentation.len]];
    return result_;
}

void* C_NSFileWrapper_InitWithCoder(void* ptr, void* inCoder) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper initWithCoder:(NSCoder*)inCoder];
    return result_;
}

void* C_NSFileWrapper_Init(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSFileWrapper* result_ = [nSFileWrapper init];
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

void* C_NSFileWrapper_AddRegularFileWithContents_PreferredFilename(void* ptr, Array data, void* fileName) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSString* result_ = [nSFileWrapper addRegularFileWithContents:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] preferredFilename:(NSString*)fileName];
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

void* C_NSFileWrapper_SymbolicLinkDestinationURL(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSURL* result_ = [nSFileWrapper symbolicLinkDestinationURL];
    return result_;
}

Array C_NSFileWrapper_SerializedRepresentation(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSData* result_ = [nSFileWrapper serializedRepresentation];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
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

Array C_NSFileWrapper_RegularFileContents(void* ptr) {
    NSFileWrapper* nSFileWrapper = (NSFileWrapper*)ptr;
    NSData* result_ = [nSFileWrapper regularFileContents];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}
