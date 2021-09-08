#import <WebKit/WebKit.h>
#import "web_view.h"

void* C_WebView_Alloc() {
    return [WKWebView alloc];
}

void* C_WKWebView_InitWithFrame_Configuration(void* ptr, CGRect frame, void* configuration) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKWebView* result_ = [wKWebView initWithFrame:frame configuration:(WKWebViewConfiguration*)configuration];
    return result_;
}

void* C_WKWebView_InitWithCoder(void* ptr, void* coder) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKWebView* result_ = [wKWebView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_WKWebView_Init(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKWebView* result_ = [wKWebView init];
    return result_;
}

void* C_WKWebView_AllocWebView() {
    WKWebView* result_ = [WKWebView alloc];
    return result_;
}

void* C_WKWebView_NewWebView() {
    WKWebView* result_ = [WKWebView new];
    return result_;
}

void* C_WKWebView_Autorelease(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKWebView* result_ = [wKWebView autorelease];
    return result_;
}

void* C_WKWebView_Retain(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKWebView* result_ = [wKWebView retain];
    return result_;
}

bool C_WKWebView_WebView_HandlesURLScheme(void* urlScheme) {
    BOOL result_ = [WKWebView handlesURLScheme:(NSString*)urlScheme];
    return result_;
}

void* C_WKWebView_LoadRequest(void* ptr, void* request) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKNavigation* result_ = [wKWebView loadRequest:(NSURLRequest*)request];
    return result_;
}

void* C_WKWebView_LoadHTMLString_BaseURL(void* ptr, void* _string, void* baseURL) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKNavigation* result_ = [wKWebView loadHTMLString:(NSString*)_string baseURL:(NSURL*)baseURL];
    return result_;
}

void* C_WKWebView_LoadFileURL_AllowingReadAccessToURL(void* ptr, void* URL, void* readAccessURL) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKNavigation* result_ = [wKWebView loadFileURL:(NSURL*)URL allowingReadAccessToURL:(NSURL*)readAccessURL];
    return result_;
}

void* C_WKWebView_LoadData_MIMEType_CharacterEncodingName_BaseURL(void* ptr, void* data, void* MIMEType, void* characterEncodingName, void* baseURL) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKNavigation* result_ = [wKWebView loadData:(NSData*)data MIMEType:(NSString*)MIMEType characterEncodingName:(NSString*)characterEncodingName baseURL:(NSURL*)baseURL];
    return result_;
}

void* C_WKWebView_Reload(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKNavigation* result_ = [wKWebView reload];
    return result_;
}

void* C_WKWebView_ReloadFromOrigin(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKNavigation* result_ = [wKWebView reloadFromOrigin];
    return result_;
}

void C_WKWebView_StopLoading(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView stopLoading];
}

void C_WKWebView_SetMagnification_CenteredAtPoint(void* ptr, double magnification, CGPoint point) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView setMagnification:magnification centeredAtPoint:point];
}

void C_WKWebView_GoBack(void* ptr, void* sender) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView goBack:(id)sender];
}

void C_WKWebView_GoForward(void* ptr, void* sender) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView goForward:(id)sender];
}

void* C_WKWebView_GoToBackForwardListItem(void* ptr, void* item) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKNavigation* result_ = [wKWebView goToBackForwardListItem:(WKBackForwardListItem*)item];
    return result_;
}

void* C_WKWebView_PrintOperationWithPrintInfo(void* ptr, void* printInfo) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    NSPrintOperation* result_ = [wKWebView printOperationWithPrintInfo:(NSPrintInfo*)printInfo];
    return result_;
}

void C_WKWebView_CloseAllMediaPresentations(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView closeAllMediaPresentations];
}

void* C_WKWebView_Configuration(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKWebViewConfiguration* result_ = [wKWebView configuration];
    return result_;
}

void* C_WKWebView_UIDelegate(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    id result_ = [wKWebView UIDelegate];
    return result_;
}

void C_WKWebView_SetUIDelegate(void* ptr, void* value) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView setUIDelegate:(id)value];
}

void* C_WKWebView_NavigationDelegate(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    id result_ = [wKWebView navigationDelegate];
    return result_;
}

void C_WKWebView_SetNavigationDelegate(void* ptr, void* value) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView setNavigationDelegate:(id)value];
}

bool C_WKWebView_IsLoading(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    BOOL result_ = [wKWebView isLoading];
    return result_;
}

double C_WKWebView_EstimatedProgress(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    double result_ = [wKWebView estimatedProgress];
    return result_;
}

void* C_WKWebView_Title(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    NSString* result_ = [wKWebView title];
    return result_;
}

void* C_WKWebView_URL(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    NSURL* result_ = [wKWebView URL];
    return result_;
}

void* C_WKWebView_MediaType(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    NSString* result_ = [wKWebView mediaType];
    return result_;
}

void C_WKWebView_SetMediaType(void* ptr, void* value) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView setMediaType:(NSString*)value];
}

void* C_WKWebView_CustomUserAgent(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    NSString* result_ = [wKWebView customUserAgent];
    return result_;
}

void C_WKWebView_SetCustomUserAgent(void* ptr, void* value) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView setCustomUserAgent:(NSString*)value];
}

bool C_WKWebView_HasOnlySecureContent(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    BOOL result_ = [wKWebView hasOnlySecureContent];
    return result_;
}

double C_WKWebView_PageZoom(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    CGFloat result_ = [wKWebView pageZoom];
    return result_;
}

void C_WKWebView_SetPageZoom(void* ptr, double value) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView setPageZoom:value];
}

bool C_WKWebView_AllowsMagnification(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    BOOL result_ = [wKWebView allowsMagnification];
    return result_;
}

void C_WKWebView_SetAllowsMagnification(void* ptr, bool value) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView setAllowsMagnification:value];
}

double C_WKWebView_Magnification(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    CGFloat result_ = [wKWebView magnification];
    return result_;
}

void C_WKWebView_SetMagnification(void* ptr, double value) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView setMagnification:value];
}

bool C_WKWebView_AllowsBackForwardNavigationGestures(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    BOOL result_ = [wKWebView allowsBackForwardNavigationGestures];
    return result_;
}

void C_WKWebView_SetAllowsBackForwardNavigationGestures(void* ptr, bool value) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView setAllowsBackForwardNavigationGestures:value];
}

void* C_WKWebView_BackForwardList(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    WKBackForwardList* result_ = [wKWebView backForwardList];
    return result_;
}

bool C_WKWebView_CanGoBack(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    BOOL result_ = [wKWebView canGoBack];
    return result_;
}

bool C_WKWebView_CanGoForward(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    BOOL result_ = [wKWebView canGoForward];
    return result_;
}

bool C_WKWebView_AllowsLinkPreview(void* ptr) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    BOOL result_ = [wKWebView allowsLinkPreview];
    return result_;
}

void C_WKWebView_SetAllowsLinkPreview(void* ptr, bool value) {
    WKWebView* wKWebView = (WKWebView*)ptr;
    [wKWebView setAllowsLinkPreview:value];
}
