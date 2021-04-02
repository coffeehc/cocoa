#import <WebKit/WebKit.h>
#import "wk_web_view.h"

bool WKWebView_IsLoading(void* ptr) {
	WKWebView* wKWebView = (WKWebView*)ptr;
	return [wKWebView isLoading];
}

double WKWebView_EstimatedProgress(void* ptr) {
	WKWebView* wKWebView = (WKWebView*)ptr;
	return [wKWebView estimatedProgress];
}

void* WKWebView_NewWKWebView(NSRect frameRect) {
	WKWebView* wKWebView = [WKWebView alloc];
	return [[wKWebView initWithFrame:frameRect] autorelease];
}

void WKWebView_StopLoading(void* ptr, void* sender) {
	WKWebView* wKWebView = (WKWebView*)ptr;
	[wKWebView stopLoading:(NSObject*)sender];
}

void WKWebView_Reload(void* ptr, void* sender) {
	WKWebView* wKWebView = (WKWebView*)ptr;
	[wKWebView reload:(NSObject*)sender];
}

void WKWebView_ReloadFromOrigin(void* ptr, void* sender) {
	WKWebView* wKWebView = (WKWebView*)ptr;
	[wKWebView reloadFromOrigin:(NSObject*)sender];
}

void WKWebView_GoBack(void* ptr, void* sender) {
	WKWebView* wKWebView = (WKWebView*)ptr;
	[wKWebView goBack:(NSObject*)sender];
}

void WKWebView_GoForward(void* ptr, void* sender) {
	WKWebView* wKWebView = (WKWebView*)ptr;
	[wKWebView goForward:(NSObject*)sender];
}

void WKWebView_LoadRequest(void* ptr, void* request) {
	WKWebView* wKWebView = (WKWebView*)ptr;
	[wKWebView loadRequest:(NSURLRequest*)request];
}

void WKWebView_LoadHTMLString(void* ptr, const char* string, void* baseURL) {
	WKWebView* wKWebView = (WKWebView*)ptr;
	[wKWebView loadHTMLString:[NSString stringWithUTF8String:string] baseURL:(NSURL*)baseURL];
}
