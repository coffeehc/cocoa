#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_WebView_Alloc();

void* C_WKWebView_InitWithFrame_Configuration(void* ptr, CGRect frame, void* configuration);
void* C_WKWebView_InitWithCoder(void* ptr, void* coder);
void* C_WKWebView_Init(void* ptr);
bool C_WKWebView_WebView_HandlesURLScheme(void* urlScheme);
void* C_WKWebView_LoadRequest(void* ptr, void* request);
void* C_WKWebView_LoadHTMLString_BaseURL(void* ptr, void* _string, void* baseURL);
void* C_WKWebView_LoadFileURL_AllowingReadAccessToURL(void* ptr, void* URL, void* readAccessURL);
void* C_WKWebView_LoadData_MIMEType_CharacterEncodingName_BaseURL(void* ptr, Array data, void* MIMEType, void* characterEncodingName, void* baseURL);
void* C_WKWebView_Reload(void* ptr);
void* C_WKWebView_ReloadFromOrigin(void* ptr);
void C_WKWebView_StopLoading(void* ptr);
void C_WKWebView_SetMagnification_CenteredAtPoint(void* ptr, double magnification, CGPoint point);
void C_WKWebView_GoBack(void* ptr, void* sender);
void C_WKWebView_GoForward(void* ptr, void* sender);
void* C_WKWebView_GoToBackForwardListItem(void* ptr, void* item);
void* C_WKWebView_PrintOperationWithPrintInfo(void* ptr, void* printInfo);
void C_WKWebView_CloseAllMediaPresentations(void* ptr);
void* C_WKWebView_Configuration(void* ptr);
void* C_WKWebView_UIDelegate(void* ptr);
void C_WKWebView_SetUIDelegate(void* ptr, void* value);
void* C_WKWebView_NavigationDelegate(void* ptr);
void C_WKWebView_SetNavigationDelegate(void* ptr, void* value);
bool C_WKWebView_IsLoading(void* ptr);
double C_WKWebView_EstimatedProgress(void* ptr);
void* C_WKWebView_Title(void* ptr);
void* C_WKWebView_URL(void* ptr);
void* C_WKWebView_MediaType(void* ptr);
void C_WKWebView_SetMediaType(void* ptr, void* value);
void* C_WKWebView_CustomUserAgent(void* ptr);
void C_WKWebView_SetCustomUserAgent(void* ptr, void* value);
bool C_WKWebView_HasOnlySecureContent(void* ptr);
double C_WKWebView_PageZoom(void* ptr);
void C_WKWebView_SetPageZoom(void* ptr, double value);
bool C_WKWebView_AllowsMagnification(void* ptr);
void C_WKWebView_SetAllowsMagnification(void* ptr, bool value);
double C_WKWebView_Magnification(void* ptr);
void C_WKWebView_SetMagnification(void* ptr, double value);
bool C_WKWebView_AllowsBackForwardNavigationGestures(void* ptr);
void C_WKWebView_SetAllowsBackForwardNavigationGestures(void* ptr, bool value);
void* C_WKWebView_BackForwardList(void* ptr);
bool C_WKWebView_CanGoBack(void* ptr);
bool C_WKWebView_CanGoForward(void* ptr);
bool C_WKWebView_AllowsLinkPreview(void* ptr);
void C_WKWebView_SetAllowsLinkPreview(void* ptr, bool value);
