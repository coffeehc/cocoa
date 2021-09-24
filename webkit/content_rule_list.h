#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_ContentRuleList_Alloc();

void* C_WKContentRuleList_AllocContentRuleList();
void* C_WKContentRuleList_Init(void* ptr);
void* C_WKContentRuleList_NewContentRuleList();
void* C_WKContentRuleList_Autorelease(void* ptr);
void* C_WKContentRuleList_Retain(void* ptr);
void* C_WKContentRuleList_Identifier(void* ptr);
