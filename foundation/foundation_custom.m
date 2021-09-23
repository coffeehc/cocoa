#import "foundation_custom.h"
#import <Foundation/NSString.h>
#import <Foundation/NSData.h>

void* String_New(const char* str) {
    return [NSString stringWithUTF8String:str];
}

const char* String_Value(void* ptr) {
    return [((NSString*)ptr) UTF8String];
}

void* Data_New(void* data, int len) {
    if (len <= 0) {
        return [NSData data];
    }
    return [NSData dataWithBytes:(Byte *)data length:len];
}

Data Data_ToBytes(void* ptr) {
    NSData* nsData = (NSData*)ptr;
    Data array = {
        .data = [nsData bytes],
        .len = nsData.length
    };
    return array;
}

void* Selector_SelectorFromString(const char* name) {
    return NSSelectorFromString([NSString stringWithUTF8String:name]);
}

const char* Selector_StringFromSelector(void *selector) {
    return [NSStringFromSelector((SEL)selector) UTF8String];
}