#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_BackForwardListItem_Alloc();

void* C_WKBackForwardListItem_Title(void* ptr);
void* C_WKBackForwardListItem_URL(void* ptr);
void* C_WKBackForwardListItem_InitialURL(void* ptr);
