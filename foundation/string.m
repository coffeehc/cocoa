#import <Foundation/NSString.h>
#import "string.h"

void* String_New(const char* str) {
    return [NSString stringWithUTF8String:str];
}

const char* String_Value(void* ptr) {
    return [((NSString*)ptr) UTF8String];
}