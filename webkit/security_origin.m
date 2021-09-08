#import <WebKit/WebKit.h>
#import "security_origin.h"

void* C_SecurityOrigin_Alloc() {
    return [WKSecurityOrigin alloc];
}

void* C_WKSecurityOrigin_AllocSecurityOrigin() {
    WKSecurityOrigin* result_ = [WKSecurityOrigin alloc];
    return result_;
}

void* C_WKSecurityOrigin_Autorelease(void* ptr) {
    WKSecurityOrigin* wKSecurityOrigin = (WKSecurityOrigin*)ptr;
    WKSecurityOrigin* result_ = [wKSecurityOrigin autorelease];
    return result_;
}

void* C_WKSecurityOrigin_Retain(void* ptr) {
    WKSecurityOrigin* wKSecurityOrigin = (WKSecurityOrigin*)ptr;
    WKSecurityOrigin* result_ = [wKSecurityOrigin retain];
    return result_;
}

void* C_WKSecurityOrigin_Host(void* ptr) {
    WKSecurityOrigin* wKSecurityOrigin = (WKSecurityOrigin*)ptr;
    NSString* result_ = [wKSecurityOrigin host];
    return result_;
}

int C_WKSecurityOrigin_Port(void* ptr) {
    WKSecurityOrigin* wKSecurityOrigin = (WKSecurityOrigin*)ptr;
    NSInteger result_ = [wKSecurityOrigin port];
    return result_;
}

void* C_WKSecurityOrigin_Protocol(void* ptr) {
    WKSecurityOrigin* wKSecurityOrigin = (WKSecurityOrigin*)ptr;
    NSString* result_ = [wKSecurityOrigin protocol];
    return result_;
}
