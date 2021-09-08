#import <WebKit/WebKit.h>
#import "window_features.h"

void* C_WindowFeatures_Alloc() {
    return [WKWindowFeatures alloc];
}

void* C_WKWindowFeatures_AllocWindowFeatures() {
    WKWindowFeatures* result_ = [WKWindowFeatures alloc];
    return result_;
}

void* C_WKWindowFeatures_Init(void* ptr) {
    WKWindowFeatures* wKWindowFeatures = (WKWindowFeatures*)ptr;
    WKWindowFeatures* result_ = [wKWindowFeatures init];
    return result_;
}

void* C_WKWindowFeatures_NewWindowFeatures() {
    WKWindowFeatures* result_ = [WKWindowFeatures new];
    return result_;
}

void* C_WKWindowFeatures_Autorelease(void* ptr) {
    WKWindowFeatures* wKWindowFeatures = (WKWindowFeatures*)ptr;
    WKWindowFeatures* result_ = [wKWindowFeatures autorelease];
    return result_;
}

void* C_WKWindowFeatures_Retain(void* ptr) {
    WKWindowFeatures* wKWindowFeatures = (WKWindowFeatures*)ptr;
    WKWindowFeatures* result_ = [wKWindowFeatures retain];
    return result_;
}

void* C_WKWindowFeatures_AllowsResizing(void* ptr) {
    WKWindowFeatures* wKWindowFeatures = (WKWindowFeatures*)ptr;
    NSNumber* result_ = [wKWindowFeatures allowsResizing];
    return result_;
}

void* C_WKWindowFeatures_Height(void* ptr) {
    WKWindowFeatures* wKWindowFeatures = (WKWindowFeatures*)ptr;
    NSNumber* result_ = [wKWindowFeatures height];
    return result_;
}

void* C_WKWindowFeatures_Width(void* ptr) {
    WKWindowFeatures* wKWindowFeatures = (WKWindowFeatures*)ptr;
    NSNumber* result_ = [wKWindowFeatures width];
    return result_;
}

void* C_WKWindowFeatures_X(void* ptr) {
    WKWindowFeatures* wKWindowFeatures = (WKWindowFeatures*)ptr;
    NSNumber* result_ = [wKWindowFeatures x];
    return result_;
}

void* C_WKWindowFeatures_Y(void* ptr) {
    WKWindowFeatures* wKWindowFeatures = (WKWindowFeatures*)ptr;
    NSNumber* result_ = [wKWindowFeatures y];
    return result_;
}

void* C_WKWindowFeatures_MenuBarVisibility(void* ptr) {
    WKWindowFeatures* wKWindowFeatures = (WKWindowFeatures*)ptr;
    NSNumber* result_ = [wKWindowFeatures menuBarVisibility];
    return result_;
}

void* C_WKWindowFeatures_StatusBarVisibility(void* ptr) {
    WKWindowFeatures* wKWindowFeatures = (WKWindowFeatures*)ptr;
    NSNumber* result_ = [wKWindowFeatures statusBarVisibility];
    return result_;
}

void* C_WKWindowFeatures_ToolbarsVisibility(void* ptr) {
    WKWindowFeatures* wKWindowFeatures = (WKWindowFeatures*)ptr;
    NSNumber* result_ = [wKWindowFeatures toolbarsVisibility];
    return result_;
}
