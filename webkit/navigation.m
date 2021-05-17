#import <WebKit/WebKit.h>
#import "navigation.h"

void* C_Navigation_Alloc() {
    return [WKNavigation alloc];
}

void* C_WKNavigation_Init(void* ptr) {
    WKNavigation* wKNavigation = (WKNavigation*)ptr;
    WKNavigation* result_ = [wKNavigation init];
    return result_;
}

int C_WKNavigation_EffectiveContentMode(void* ptr) {
    WKNavigation* wKNavigation = (WKNavigation*)ptr;
    WKContentMode result_ = [wKNavigation effectiveContentMode];
    return result_;
}
