#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_Download_Alloc();

void* C_WKDownload_AllocDownload();
void* C_WKDownload_Init(void* ptr);
void* C_WKDownload_NewDownload();
void* C_WKDownload_Autorelease(void* ptr);
void* C_WKDownload_Retain(void* ptr);
void* C_WKDownload_Delegate(void* ptr);
void C_WKDownload_SetDelegate(void* ptr, void* value);
void* C_WKDownload_OriginalRequest(void* ptr);
void* C_WKDownload_WebView(void* ptr);
