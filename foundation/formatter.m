#import <Foundation/Foundation.h>
#import "formatter.h"

void* C_Formatter_Alloc() {
    return [NSFormatter alloc];
}

void* C_NSFormatter_Init(void* ptr) {
    NSFormatter* nSFormatter = (NSFormatter*)ptr;
    NSFormatter* result_ = [nSFormatter init];
    return result_;
}

void* C_NSFormatter_StringForObjectValue(void* ptr, void* obj) {
    NSFormatter* nSFormatter = (NSFormatter*)ptr;
    NSString* result_ = [nSFormatter stringForObjectValue:(id)obj];
    return result_;
}

void* C_NSFormatter_EditingStringForObjectValue(void* ptr, void* obj) {
    NSFormatter* nSFormatter = (NSFormatter*)ptr;
    NSString* result_ = [nSFormatter editingStringForObjectValue:(id)obj];
    return result_;
}
