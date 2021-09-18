#import "navigation_delegate.h"
#import "_cgo_export.h"
#import <WebKit/WKNavigationDelegate.h>

@interface WKNavigationDelegateAdaptor : NSObject <WKNavigationDelegate>
@property (assign) uintptr_t goID;
@end

@implementation WKNavigationDelegateAdaptor


- (void)webView:(WKWebView*)webView didStartProvisionalNavigation:(WKNavigation*)navigation {
    navigationDelegate_WebView_DidStartProvisionalNavigation([self goID], webView, navigation);
}

- (void)webView:(WKWebView*)webView didReceiveServerRedirectForProvisionalNavigation:(WKNavigation*)navigation {
    navigationDelegate_WebView_DidReceiveServerRedirectForProvisionalNavigation([self goID], webView, navigation);
}

- (void)webView:(WKWebView*)webView didCommitNavigation:(WKNavigation*)navigation {
    navigationDelegate_WebView_DidCommitNavigation([self goID], webView, navigation);
}

- (void)webView:(WKWebView*)webView didFinishNavigation:(WKNavigation*)navigation {
    navigationDelegate_WebView_DidFinishNavigation([self goID], webView, navigation);
}

- (void)webView:(WKWebView*)webView didFailNavigation:(WKNavigation*)navigation withError:(NSError*)error {
    navigationDelegate_WebView_DidFailNavigation_WithError([self goID], webView, navigation, error);
}

- (void)webView:(WKWebView*)webView didFailProvisionalNavigation:(WKNavigation*)navigation withError:(NSError*)error {
    navigationDelegate_WebView_DidFailProvisionalNavigation_WithError([self goID], webView, navigation, error);
}

- (void)webViewWebContentProcessDidTerminate:(WKWebView*)webView {
    navigationDelegate_WebViewWebContentProcessDidTerminate([self goID], webView);
}

- (void)webView:(WKWebView*)webView navigationAction:(WKNavigationAction*)navigationAction didBecomeDownload:(WKDownload*)download {
    navigationDelegate_WebView_NavigationAction_DidBecomeDownload([self goID], webView, navigationAction, download);
}

- (void)webView:(WKWebView*)webView navigationResponse:(WKNavigationResponse*)navigationResponse didBecomeDownload:(WKDownload*)download {
    navigationDelegate_WebView_NavigationResponse_DidBecomeDownload([self goID], webView, navigationResponse, download);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return NavigationDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteWebKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapNavigationDelegate(uintptr_t goID) {
    WKNavigationDelegateAdaptor* adaptor = [[WKNavigationDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
