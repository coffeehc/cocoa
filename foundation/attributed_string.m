#import <Foundation/Foundation.h>
#import "attributed_string.h"

void* C_AttributedString_Alloc() {
    return [NSAttributedString alloc];
}

void* C_NSAttributedString_AllocAttributedString() {
    NSAttributedString* result_ = [NSAttributedString alloc];
    return result_;
}

void* C_NSAttributedString_Init(void* ptr) {
    NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
    NSAttributedString* result_ = [nSAttributedString init];
    return result_;
}

void* C_NSAttributedString_NewAttributedString() {
    NSAttributedString* result_ = [NSAttributedString new];
    return result_;
}

void* C_NSAttributedString_Autorelease(void* ptr) {
    NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
    NSAttributedString* result_ = [nSAttributedString autorelease];
    return result_;
}

void* C_NSAttributedString_Retain(void* ptr) {
    NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
    NSAttributedString* result_ = [nSAttributedString retain];
    return result_;
}

void* C_NSAttributedString_String(void* ptr) {
    NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
    NSString* result_ = [nSAttributedString string];
    return result_;
}

unsigned int C_NSAttributedString_Length(void* ptr) {
    NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
    NSUInteger result_ = [nSAttributedString length];
    return result_;
}
