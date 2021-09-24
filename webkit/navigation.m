#import "navigation.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <WebKit/WKNavigation.h>

void* C_Navigation_Alloc() {
    return [WKNavigation alloc];
}

void* C_WKNavigation_AllocNavigation() {
    WKNavigation* result_ = [WKNavigation alloc];
    return result_;
}

void* C_WKNavigation_Init(void* ptr) {
    WKNavigation* wKNavigation = (WKNavigation*)ptr;
    WKNavigation* result_ = [wKNavigation init];
    return result_;
}

void* C_WKNavigation_NewNavigation() {
    WKNavigation* result_ = [WKNavigation new];
    return result_;
}

void* C_WKNavigation_Autorelease(void* ptr) {
    WKNavigation* wKNavigation = (WKNavigation*)ptr;
    WKNavigation* result_ = [wKNavigation autorelease];
    return result_;
}

void* C_WKNavigation_Retain(void* ptr) {
    WKNavigation* wKNavigation = (WKNavigation*)ptr;
    WKNavigation* result_ = [wKNavigation retain];
    return result_;
}

int C_WKNavigation_EffectiveContentMode(void* ptr) {
    WKNavigation* wKNavigation = (WKNavigation*)ptr;
    WKContentMode result_ = [wKNavigation effectiveContentMode];
    return result_;
}
