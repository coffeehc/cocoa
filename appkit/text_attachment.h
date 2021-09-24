#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TextAttachment_Alloc();

void* C_NSTextAttachment_InitWithFileWrapper(void* ptr, void* fileWrapper);
void* C_NSTextAttachment_InitWithData_OfType(void* ptr, void* contentData, void* uti);
void* C_NSTextAttachment_AllocTextAttachment();
void* C_NSTextAttachment_Init(void* ptr);
void* C_NSTextAttachment_NewTextAttachment();
void* C_NSTextAttachment_Autorelease(void* ptr);
void* C_NSTextAttachment_Retain(void* ptr);
CGRect C_NSTextAttachment_Bounds(void* ptr);
void C_NSTextAttachment_SetBounds(void* ptr, CGRect value);
void* C_NSTextAttachment_Contents(void* ptr);
void C_NSTextAttachment_SetContents(void* ptr, void* value);
void* C_NSTextAttachment_FileType(void* ptr);
void C_NSTextAttachment_SetFileType(void* ptr, void* value);
void* C_NSTextAttachment_Image(void* ptr);
void C_NSTextAttachment_SetImage(void* ptr, void* value);
void* C_NSTextAttachment_FileWrapper(void* ptr);
void C_NSTextAttachment_SetFileWrapper(void* ptr, void* value);
void* C_NSTextAttachment_AttachmentCell(void* ptr);
void C_NSTextAttachment_SetAttachmentCell(void* ptr, void* value);
