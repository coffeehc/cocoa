#import <Appkit/Appkit.h>
#import "text_attachment.h"

void* C_TextAttachment_Alloc() {
    return [NSTextAttachment alloc];
}

void* C_NSTextAttachment_InitWithFileWrapper(void* ptr, void* fileWrapper) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSTextAttachment* result_ = [nSTextAttachment initWithFileWrapper:(NSFileWrapper*)fileWrapper];
    return result_;
}

void* C_NSTextAttachment_InitWithData_OfType(void* ptr, Array contentData, void* uti) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSTextAttachment* result_ = [nSTextAttachment initWithData:[[NSData alloc] initWithBytes:(Byte *)contentData.data length:contentData.len] ofType:(NSString*)uti];
    return result_;
}

void* C_NSTextAttachment_Init(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSTextAttachment* result_ = [nSTextAttachment init];
    return result_;
}

CGRect C_NSTextAttachment_Bounds(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSRect result_ = [nSTextAttachment bounds];
    return result_;
}

void C_NSTextAttachment_SetBounds(void* ptr, CGRect value) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    [nSTextAttachment setBounds:value];
}

Array C_NSTextAttachment_Contents(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSData* result_ = [nSTextAttachment contents];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

void C_NSTextAttachment_SetContents(void* ptr, Array value) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    [nSTextAttachment setContents:[[NSData alloc] initWithBytes:(Byte *)value.data length:value.len]];
}

void* C_NSTextAttachment_FileType(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSString* result_ = [nSTextAttachment fileType];
    return result_;
}

void C_NSTextAttachment_SetFileType(void* ptr, void* value) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    [nSTextAttachment setFileType:(NSString*)value];
}

void* C_NSTextAttachment_Image(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSImage* result_ = [nSTextAttachment image];
    return result_;
}

void C_NSTextAttachment_SetImage(void* ptr, void* value) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    [nSTextAttachment setImage:(NSImage*)value];
}

void* C_NSTextAttachment_FileWrapper(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSFileWrapper* result_ = [nSTextAttachment fileWrapper];
    return result_;
}

void C_NSTextAttachment_SetFileWrapper(void* ptr, void* value) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    [nSTextAttachment setFileWrapper:(NSFileWrapper*)value];
}

void* C_NSTextAttachment_AttachmentCell(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    id result_ = [nSTextAttachment attachmentCell];
    return result_;
}

void C_NSTextAttachment_SetAttachmentCell(void* ptr, void* value) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    [nSTextAttachment setAttachmentCell:(id)value];
}
