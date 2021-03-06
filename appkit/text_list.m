#import "text_list.h"
#import <AppKit/NSTextList.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_TextList_Alloc() {
    return [NSTextList alloc];
}

void* C_NSTextList_InitWithMarkerFormat_Options(void* ptr, void* format, unsigned int mask) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSTextList* result_ = [nSTextList initWithMarkerFormat:(NSString*)format options:mask];
    return result_;
}

void* C_NSTextList_AllocTextList() {
    NSTextList* result_ = [NSTextList alloc];
    return result_;
}

void* C_NSTextList_Init(void* ptr) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSTextList* result_ = [nSTextList init];
    return result_;
}

void* C_NSTextList_NewTextList() {
    NSTextList* result_ = [NSTextList new];
    return result_;
}

void* C_NSTextList_Autorelease(void* ptr) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSTextList* result_ = [nSTextList autorelease];
    return result_;
}

void* C_NSTextList_Retain(void* ptr) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSTextList* result_ = [nSTextList retain];
    return result_;
}

void* C_NSTextList_MarkerForItemNumber(void* ptr, int itemNum) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSString* result_ = [nSTextList markerForItemNumber:itemNum];
    return result_;
}

void* C_NSTextList_MarkerFormat(void* ptr) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSTextListMarkerFormat result_ = [nSTextList markerFormat];
    return result_;
}

unsigned int C_NSTextList_ListOptions(void* ptr) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSTextListOptions result_ = [nSTextList listOptions];
    return result_;
}

int C_NSTextList_StartingItemNumber(void* ptr) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSInteger result_ = [nSTextList startingItemNumber];
    return result_;
}

void C_NSTextList_SetStartingItemNumber(void* ptr, int value) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    [nSTextList setStartingItemNumber:value];
}
