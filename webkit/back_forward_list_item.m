#import <WebKit/WebKit.h>
#import "back_forward_list_item.h"

void* C_BackForwardListItem_Alloc() {
    return [WKBackForwardListItem alloc];
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
