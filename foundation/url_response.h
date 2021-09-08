#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_URLResponse_Alloc();

void* C_NSURLResponse_InitWithURL_MIMEType_ExpectedContentLength_TextEncodingName(void* ptr, void* URL, void* MIMEType, int length, void* name);
void* C_NSURLResponse_AllocURLResponse();
void* C_NSURLResponse_Init(void* ptr);
void* C_NSURLResponse_NewURLResponse();
void* C_NSURLResponse_Autorelease(void* ptr);
void* C_NSURLResponse_Retain(void* ptr);
void* C_NSURLResponse_SuggestedFilename(void* ptr);
void* C_NSURLResponse_MIMEType(void* ptr);
void* C_NSURLResponse_TextEncodingName(void* ptr);
void* C_NSURLResponse_URL(void* ptr);
