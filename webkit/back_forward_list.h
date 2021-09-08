#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_BackForwardList_Alloc();

void* C_WKBackForwardList_AllocBackForwardList();
void* C_WKBackForwardList_Autorelease(void* ptr);
void* C_WKBackForwardList_Retain(void* ptr);
void* C_WKBackForwardList_ItemAtIndex(void* ptr, int index);
void* C_WKBackForwardList_BackItem(void* ptr);
void* C_WKBackForwardList_CurrentItem(void* ptr);
void* C_WKBackForwardList_ForwardItem(void* ptr);
Array C_WKBackForwardList_BackList(void* ptr);
Array C_WKBackForwardList_ForwardList(void* ptr);
