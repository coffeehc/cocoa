#import <Appkit/Appkit.h>
#import "text_attachment.h"

void* C_TextAttachment_Alloc() {
    return [NSTextAttachment alloc];
}

void* C_NSTextAttachment_InitWithFileWrapper(void* ptr, void* fileWrapper) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSTextAttachment* result = [nSTextAttachment initWithFileWrapper:(NSFileWrapper*)fileWrapper];
    return result;
}

void* C_NSTextAttachment_InitWithData_OfType(void* ptr, Array contentData, void* uti) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSTextAttachment* result = [nSTextAttachment initWithData:[[NSData alloc] initWithBytes:(Byte *)contentData.data length:contentData.len] ofType:(NSString*)uti];
    return result;
}

void* C_NSTextAttachment_Init(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSTextAttachment* result = [nSTextAttachment init];
    return result;
}

CGRect C_NSTextAttachment_Bounds(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSRect result = [nSTextAttachment bounds];
    return result;
}

void C_NSTextAttachment_SetBounds(void* ptr, CGRect value) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    [nSTextAttachment setBounds:value];
}

Array C_NSTextAttachment_Contents(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSData* result = [nSTextAttachment contents];
    Array resultarray;
    resultarray.data = [result bytes];
    resultarray.len = result.length;
    return resultarray;
}

void C_NSTextAttachment_SetContents(void* ptr, Array value) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    [nSTextAttachment setContents:[[NSData alloc] initWithBytes:(Byte *)value.data length:value.len]];
}

void* C_NSTextAttachment_FileType(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSString* result = [nSTextAttachment fileType];
    return result;
}

void C_NSTextAttachment_SetFileType(void* ptr, void* value) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    [nSTextAttachment setFileType:(NSString*)value];
}

void* C_NSTextAttachment_Image(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSImage* result = [nSTextAttachment image];
    return result;
}

void C_NSTextAttachment_SetImage(void* ptr, void* value) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    [nSTextAttachment setImage:(NSImage*)value];
}

void* C_NSTextAttachment_FileWrapper(void* ptr) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    NSFileWrapper* result = [nSTextAttachment fileWrapper];
    return result;
}

void C_NSTextAttachment_SetFileWrapper(void* ptr, void* value) {
    NSTextAttachment* nSTextAttachment = (NSTextAttachment*)ptr;
    [nSTextAttachment setFileWrapper:(NSFileWrapper*)value];
}
