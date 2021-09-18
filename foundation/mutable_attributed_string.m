#import "mutable_attributed_string.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSAttributedString.h>

void* C_MutableAttributedString_Alloc() {
    return [NSMutableAttributedString alloc];
}

void* C_NSMutableAttributedString_AllocMutableAttributedString() {
    NSMutableAttributedString* result_ = [NSMutableAttributedString alloc];
    return result_;
}

void* C_NSMutableAttributedString_Init(void* ptr) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableAttributedString* result_ = [nSMutableAttributedString init];
    return result_;
}

void* C_NSMutableAttributedString_NewMutableAttributedString() {
    NSMutableAttributedString* result_ = [NSMutableAttributedString new];
    return result_;
}

void* C_NSMutableAttributedString_Autorelease(void* ptr) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableAttributedString* result_ = [nSMutableAttributedString autorelease];
    return result_;
}

void* C_NSMutableAttributedString_Retain(void* ptr) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableAttributedString* result_ = [nSMutableAttributedString retain];
    return result_;
}

void* C_NSMutableAttributedString_MutableString(void* ptr) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableString* result_ = [nSMutableAttributedString mutableString];
    return result_;
}
