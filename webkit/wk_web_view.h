#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool WKWebView_IsLoading(void* ptr);
double WKWebView_EstimatedProgress(void* ptr);

void* WKWebView_NewWKWebView(NSRect frameRect);
void WKWebView_StopLoading(void* ptr, void* sender);
void WKWebView_Reload(void* ptr, void* sender);
void WKWebView_ReloadFromOrigin(void* ptr, void* sender);
void WKWebView_GoBack(void* ptr, void* sender);
void WKWebView_GoForward(void* ptr, void* sender);
void WKWebView_LoadRequest(void* ptr, void* request);
void WKWebView_LoadHTMLString(void* ptr, const char* string, void* baseURL);
