#import "content_world.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <WebKit/WKContentWorld.h>

void* C_ContentWorld_Alloc() {
    return [WKContentWorld alloc];
}

void* C_WKContentWorld_AllocContentWorld() {
    WKContentWorld* result_ = [WKContentWorld alloc];
    return result_;
}

void* C_WKContentWorld_Autorelease(void* ptr) {
    WKContentWorld* wKContentWorld = (WKContentWorld*)ptr;
    WKContentWorld* result_ = [wKContentWorld autorelease];
    return result_;
}

void* C_WKContentWorld_Retain(void* ptr) {
    WKContentWorld* wKContentWorld = (WKContentWorld*)ptr;
    WKContentWorld* result_ = [wKContentWorld retain];
    return result_;
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
