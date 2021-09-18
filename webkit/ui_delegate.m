#import "ui_delegate.h"
#import "_cgo_export.h"
#import <WebKit/WKUIDelegate.h>

@interface WKUIDelegateAdaptor : NSObject <WKUIDelegate>
@property (assign) uintptr_t goID;
@end

@implementation WKUIDelegateAdaptor


- (WKWebView*)webView:(WKWebView*)webView createWebViewWithConfiguration:(WKWebViewConfiguration*)configuration forNavigationAction:(WKNavigationAction*)navigationAction windowFeatures:(WKWindowFeatures*)windowFeatures {
    void* result_ = uIDelegate_WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures([self goID], webView, configuration, navigationAction, windowFeatures);
    return (WKWebView*)result_;
}

- (void)webViewDidClose:(WKWebView*)webView {
    uIDelegate_WebViewDidClose([self goID], webView);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return UIDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteWebKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapUIDelegate(uintptr_t goID) {
    WKUIDelegateAdaptor* adaptor = [[WKUIDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
