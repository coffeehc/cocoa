#import "content_rule_list.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <WebKit/WKContentRuleList.h>

void* C_ContentRuleList_Alloc() {
    return [WKContentRuleList alloc];
}

void* C_WKContentRuleList_AllocContentRuleList() {
    WKContentRuleList* result_ = [WKContentRuleList alloc];
    return result_;
}

void* C_WKContentRuleList_Init(void* ptr) {
    WKContentRuleList* wKContentRuleList = (WKContentRuleList*)ptr;
    WKContentRuleList* result_ = [wKContentRuleList init];
    return result_;
}

void* C_WKContentRuleList_NewContentRuleList() {
    WKContentRuleList* result_ = [WKContentRuleList new];
    return result_;
}

void* C_WKContentRuleList_Autorelease(void* ptr) {
    WKContentRuleList* wKContentRuleList = (WKContentRuleList*)ptr;
    WKContentRuleList* result_ = [wKContentRuleList autorelease];
    return result_;
}

void* C_WKContentRuleList_Retain(void* ptr) {
    WKContentRuleList* wKContentRuleList = (WKContentRuleList*)ptr;
    WKContentRuleList* result_ = [wKContentRuleList retain];
    return result_;
}

void* C_WKContentRuleList_Identifier(void* ptr) {
    WKContentRuleList* wKContentRuleList = (WKContentRuleList*)ptr;
    NSString* result_ = [wKContentRuleList identifier];
    return result_;
}
