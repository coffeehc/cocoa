#import "page_layout.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSPageLayout.h>

void* C_PageLayout_Alloc() {
    return [NSPageLayout alloc];
}

void* C_NSPageLayout_AllocPageLayout() {
    NSPageLayout* result_ = [NSPageLayout alloc];
    return result_;
}

void* C_NSPageLayout_Init(void* ptr) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSPageLayout* result_ = [nSPageLayout init];
    return result_;
}

void* C_NSPageLayout_NewPageLayout() {
    NSPageLayout* result_ = [NSPageLayout new];
    return result_;
}

void* C_NSPageLayout_Autorelease(void* ptr) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSPageLayout* result_ = [nSPageLayout autorelease];
    return result_;
}

void* C_NSPageLayout_Retain(void* ptr) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSPageLayout* result_ = [nSPageLayout retain];
    return result_;
}

void* C_NSPageLayout_PageLayout_() {
    NSPageLayout* result_ = [NSPageLayout pageLayout];
    return result_;
}

int C_NSPageLayout_RunModal(void* ptr) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSInteger result_ = [nSPageLayout runModal];
    return result_;
}

int C_NSPageLayout_RunModalWithPrintInfo(void* ptr, void* printInfo) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSInteger result_ = [nSPageLayout runModalWithPrintInfo:(NSPrintInfo*)printInfo];
    return result_;
}

void C_NSPageLayout_AddAccessoryController(void* ptr, void* accessoryController) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    [nSPageLayout addAccessoryController:(NSViewController*)accessoryController];
}

void C_NSPageLayout_RemoveAccessoryController(void* ptr, void* accessoryController) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    [nSPageLayout removeAccessoryController:(NSViewController*)accessoryController];
}

Array C_NSPageLayout_AccessoryControllers(void* ptr) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSArray* result_ = [nSPageLayout accessoryControllers];
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

void* C_NSPageLayout_PrintInfo(void* ptr) {
    NSPageLayout* nSPageLayout = (NSPageLayout*)ptr;
    NSPrintInfo* result_ = [nSPageLayout printInfo];
    return result_;
}
