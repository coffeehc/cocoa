#import <Appkit/Appkit.h>
#import "text_list.h"

void* C_TextList_Alloc() {
    return [NSTextList alloc];
}

void* C_NSTextList_InitWithMarkerFormat_Options(void* ptr, void* format, unsigned int mask) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSTextList* result = [nSTextList initWithMarkerFormat:(NSString*)format options:mask];
    return result;
}

void* C_NSTextList_Init(void* ptr) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSTextList* result = [nSTextList init];
    return result;
}

void* C_NSTextList_MarkerForItemNumber(void* ptr, int itemNum) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSString* result = [nSTextList markerForItemNumber:itemNum];
    return result;
}

void* C_NSTextList_MarkerFormat(void* ptr) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSTextListMarkerFormat result = [nSTextList markerFormat];
    return result;
}

unsigned int C_NSTextList_ListOptions(void* ptr) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSTextListOptions result = [nSTextList listOptions];
    return result;
}

int C_NSTextList_StartingItemNumber(void* ptr) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    NSInteger result = [nSTextList startingItemNumber];
    return result;
}

void C_NSTextList_SetStartingItemNumber(void* ptr, int value) {
    NSTextList* nSTextList = (NSTextList*)ptr;
    [nSTextList setStartingItemNumber:value];
}
