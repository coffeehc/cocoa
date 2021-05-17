#import <WebKit/WebKit.h>
#import "content_world.h"

void* C_ContentWorld_Alloc() {
    return [WKContentWorld alloc];
}

void* C_WKContentWorld_ContentWorld_WorldWithName(void* name) {
    WKContentWorld* result_ = [WKContentWorld worldWithName:(NSString*)name];
    return result_;
}

void* C_WKContentWorld_ContentWorld_DefaultClientWorld() {
    WKContentWorld* result_ = [WKContentWorld defaultClientWorld];
    return result_;
}

void* C_WKContentWorld_ContentWorld_PageWorld() {
    WKContentWorld* result_ = [WKContentWorld pageWorld];
    return result_;
}

void* C_WKContentWorld_Name(void* ptr) {
    WKContentWorld* wKContentWorld = (WKContentWorld*)ptr;
    NSString* result_ = [wKContentWorld name];
    return result_;
}
