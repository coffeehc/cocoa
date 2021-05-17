#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_ContentRuleList_Alloc();

void* C_WKContentRuleList_Init(void* ptr);
void* C_WKContentRuleList_Identifier(void* ptr);
