#import <WebKit/WebKit.h>
#import "back_forward_list.h"

void* C_BackForwardList_Alloc() {
    return [WKBackForwardList alloc];
}

void* C_WKBackForwardList_AllocBackForwardList() {
    WKBackForwardList* result_ = [WKBackForwardList alloc];
    return result_;
}

void* C_WKBackForwardList_Autorelease(void* ptr) {
    WKBackForwardList* wKBackForwardList = (WKBackForwardList*)ptr;
    WKBackForwardList* result_ = [wKBackForwardList autorelease];
    return result_;
}

void* C_WKBackForwardList_Retain(void* ptr) {
    WKBackForwardList* wKBackForwardList = (WKBackForwardList*)ptr;
    WKBackForwardList* result_ = [wKBackForwardList retain];
    return result_;
}

void* C_WKBackForwardList_ItemAtIndex(void* ptr, int index) {
    WKBackForwardList* wKBackForwardList = (WKBackForwardList*)ptr;
    WKBackForwardListItem* result_ = [wKBackForwardList itemAtIndex:index];
    return result_;
}

void* C_WKBackForwardList_BackItem(void* ptr) {
    WKBackForwardList* wKBackForwardList = (WKBackForwardList*)ptr;
    WKBackForwardListItem* result_ = [wKBackForwardList backItem];
    return result_;
}

void* C_WKBackForwardList_CurrentItem(void* ptr) {
    WKBackForwardList* wKBackForwardList = (WKBackForwardList*)ptr;
    WKBackForwardListItem* result_ = [wKBackForwardList currentItem];
    return result_;
}

void* C_WKBackForwardList_ForwardItem(void* ptr) {
    WKBackForwardList* wKBackForwardList = (WKBackForwardList*)ptr;
    WKBackForwardListItem* result_ = [wKBackForwardList forwardItem];
    return result_;
}

Array C_WKBackForwardList_BackList(void* ptr) {
    WKBackForwardList* wKBackForwardList = (WKBackForwardList*)ptr;
    NSArray* result_ = [wKBackForwardList backList];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_WKBackForwardList_ForwardList(void* ptr) {
    WKBackForwardList* wKBackForwardList = (WKBackForwardList*)ptr;
    NSArray* result_ = [wKBackForwardList forwardList];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}
