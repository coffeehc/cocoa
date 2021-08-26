#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_Download_Alloc();

void* C_WKDownload_Init(void* ptr);
void* C_WKDownload_Delegate(void* ptr);
void C_WKDownload_SetDelegate(void* ptr, void* value);
void* C_WKDownload_OriginalRequest(void* ptr);
void* C_WKDownload_WebView(void* ptr);
