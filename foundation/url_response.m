#import <Foundation/Foundation.h>
#import "url_response.h"

void* C_URLResponse_Alloc() {
    return [NSURLResponse alloc];
}

void* C_NSURLResponse_InitWithURL_MIMEType_ExpectedContentLength_TextEncodingName(void* ptr, void* URL, void* MIMEType, int length, void* name) {
    NSURLResponse* nSURLResponse = (NSURLResponse*)ptr;
    NSURLResponse* result_ = [nSURLResponse initWithURL:(NSURL*)URL MIMEType:(NSString*)MIMEType expectedContentLength:length textEncodingName:(NSString*)name];
    return result_;
}

void* C_NSURLResponse_Init(void* ptr) {
    NSURLResponse* nSURLResponse = (NSURLResponse*)ptr;
    NSURLResponse* result_ = [nSURLResponse init];
    return result_;
}

void* C_NSURLResponse_SuggestedFilename(void* ptr) {
    NSURLResponse* nSURLResponse = (NSURLResponse*)ptr;
    NSString* result_ = [nSURLResponse suggestedFilename];
    return result_;
}

void* C_NSURLResponse_MIMEType(void* ptr) {
    NSURLResponse* nSURLResponse = (NSURLResponse*)ptr;
    NSString* result_ = [nSURLResponse MIMEType];
    return result_;
}

void* C_NSURLResponse_TextEncodingName(void* ptr) {
    NSURLResponse* nSURLResponse = (NSURLResponse*)ptr;
    NSString* result_ = [nSURLResponse textEncodingName];
    return result_;
}

void* C_NSURLResponse_URL(void* ptr) {
    NSURLResponse* nSURLResponse = (NSURLResponse*)ptr;
    NSURL* result_ = [nSURLResponse URL];
    return result_;
}
