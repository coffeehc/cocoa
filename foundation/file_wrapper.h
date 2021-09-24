#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_FileWrapper_Alloc();

void* C_NSFileWrapper_InitDirectoryWithFileWrappers(void* ptr, Dictionary childrenByPreferredName);
void* C_NSFileWrapper_InitRegularFileWithContents(void* ptr, void* contents);
void* C_NSFileWrapper_InitSymbolicLinkWithDestinationURL(void* ptr, void* url);
void* C_NSFileWrapper_InitWithSerializedRepresentation(void* ptr, void* serializeRepresentation);
void* C_NSFileWrapper_InitWithCoder(void* ptr, void* inCoder);
void* C_NSFileWrapper_AllocFileWrapper();
void* C_NSFileWrapper_Init(void* ptr);
void* C_NSFileWrapper_NewFileWrapper();
void* C_NSFileWrapper_Autorelease(void* ptr);
void* C_NSFileWrapper_Retain(void* ptr);
void* C_NSFileWrapper_AddFileWrapper(void* ptr, void* child);
void C_NSFileWrapper_RemoveFileWrapper(void* ptr, void* child);
void* C_NSFileWrapper_AddRegularFileWithContents_PreferredFilename(void* ptr, void* data, void* fileName);
void* C_NSFileWrapper_KeyForFileWrapper(void* ptr, void* child);
bool C_NSFileWrapper_MatchesContentsOfURL(void* ptr, void* url);
bool C_NSFileWrapper_IsRegularFile(void* ptr);
bool C_NSFileWrapper_IsDirectory(void* ptr);
bool C_NSFileWrapper_IsSymbolicLink(void* ptr);
Dictionary C_NSFileWrapper_FileWrappers(void* ptr);
void* C_NSFileWrapper_SymbolicLinkDestinationURL(void* ptr);
void* C_NSFileWrapper_SerializedRepresentation(void* ptr);
void* C_NSFileWrapper_Filename(void* ptr);
void C_NSFileWrapper_SetFilename(void* ptr, void* value);
void* C_NSFileWrapper_PreferredFilename(void* ptr);
void C_NSFileWrapper_SetPreferredFilename(void* ptr, void* value);
Dictionary C_NSFileWrapper_FileAttributes(void* ptr);
void C_NSFileWrapper_SetFileAttributes(void* ptr, Dictionary value);
void* C_NSFileWrapper_RegularFileContents(void* ptr);
