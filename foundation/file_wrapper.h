#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_FileWrapper_Alloc();

void* C_NSFileWrapper_InitRegularFileWithContents(void* ptr, Array contents);
void* C_NSFileWrapper_InitSymbolicLinkWithDestinationURL(void* ptr, void* url);
void* C_NSFileWrapper_InitWithSerializedRepresentation(void* ptr, Array serializeRepresentation);
void* C_NSFileWrapper_InitWithCoder(void* ptr, void* inCoder);
void* C_NSFileWrapper_Init(void* ptr);
void* C_NSFileWrapper_AddFileWrapper(void* ptr, void* child);
void C_NSFileWrapper_RemoveFileWrapper(void* ptr, void* child);
void* C_NSFileWrapper_AddRegularFileWithContents_PreferredFilename(void* ptr, Array data, void* fileName);
void* C_NSFileWrapper_KeyForFileWrapper(void* ptr, void* child);
bool C_NSFileWrapper_MatchesContentsOfURL(void* ptr, void* url);
bool C_NSFileWrapper_IsRegularFile(void* ptr);
bool C_NSFileWrapper_IsDirectory(void* ptr);
bool C_NSFileWrapper_IsSymbolicLink(void* ptr);
void* C_NSFileWrapper_SymbolicLinkDestinationURL(void* ptr);
Array C_NSFileWrapper_SerializedRepresentation(void* ptr);
void* C_NSFileWrapper_Filename(void* ptr);
void C_NSFileWrapper_SetFilename(void* ptr, void* value);
void* C_NSFileWrapper_PreferredFilename(void* ptr);
void C_NSFileWrapper_SetPreferredFilename(void* ptr, void* value);
Array C_NSFileWrapper_RegularFileContents(void* ptr);
