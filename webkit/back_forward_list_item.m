#import "back_forward_list_item.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <WebKit/WKBackForwardListItem.h>

void* C_BackForwardListItem_Alloc() {
    return [WKBackForwardListItem alloc];
}

void* C_WKBackForwardListItem_AllocBackForwardListItem() {
    WKBackForwardListItem* result_ = [WKBackForwardListItem alloc];
    return result_;
}

void* C_WKBackForwardListItem_Autorelease(void* ptr) {
    WKBackForwardListItem* wKBackForwardListItem = (WKBackForwardListItem*)ptr;
    WKBackForwardListItem* result_ = [wKBackForwardListItem autorelease];
    return result_;
}

void* C_WKBackForwardListItem_Retain(void* ptr) {
    WKBackForwardListItem* wKBackForwardListItem = (WKBackForwardListItem*)ptr;
    WKBackForwardListItem* result_ = [wKBackForwardListItem retain];
    return result_;
}

void* C_WKBackForwardListItem_Title(void* ptr) {
    WKBackForwardListItem* wKBackForwardListItem = (WKBackForwardListItem*)ptr;
    NSString* result_ = [wKBackForwardListItem title];
    return result_;
}

void* C_WKBackForwardListItem_URL(void* ptr) {
    WKBackForwardListItem* wKBackForwardListItem = (WKBackForwardListItem*)ptr;
    NSURL* result_ = [wKBackForwardListItem URL];
    return result_;
}

void* C_WKBackForwardListItem_InitialURL(void* ptr) {
    WKBackForwardListItem* wKBackForwardListItem = (WKBackForwardListItem*)ptr;
    NSURL* result_ = [wKBackForwardListItem initialURL];
    return result_;
}
