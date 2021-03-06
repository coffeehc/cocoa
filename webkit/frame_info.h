#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_FrameInfo_Alloc();

void* C_WKFrameInfo_AllocFrameInfo();
void* C_WKFrameInfo_Init(void* ptr);
void* C_WKFrameInfo_NewFrameInfo();
void* C_WKFrameInfo_Autorelease(void* ptr);
void* C_WKFrameInfo_Retain(void* ptr);
bool C_WKFrameInfo_IsMainFrame(void* ptr);
void* C_WKFrameInfo_Request(void* ptr);
void* C_WKFrameInfo_SecurityOrigin(void* ptr);
void* C_WKFrameInfo_WebView(void* ptr);
