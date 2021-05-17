#import <WebKit/WebKit.h>
#import "content_rule_list.h"

void* C_ContentRuleList_Alloc() {
    return [WKContentRuleList alloc];
}

void* C_WKContentRuleList_Init(void* ptr) {
    WKContentRuleList* wKContentRuleList = (WKContentRuleList*)ptr;
    WKContentRuleList* result_ = [wKContentRuleList init];
    return result_;
}

void* C_WKContentRuleList_Identifier(void* ptr) {
    WKContentRuleList* wKContentRuleList = (WKContentRuleList*)ptr;
    NSString* result_ = [wKContentRuleList identifier];
    return result_;
}
