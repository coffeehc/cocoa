#import <Foundation/Foundation.h>
#import "mutable_attributed_string.h"

void* C_MutableAttributedString_Alloc() {
    return [NSMutableAttributedString alloc];
}

void* C_NSMutableAttributedString_InitWithString(void* ptr, void* str) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableAttributedString* result_ = [nSMutableAttributedString initWithString:(NSString*)str];
    return result_;
}

void* C_NSMutableAttributedString_InitWithAttributedString(void* ptr, void* attrStr) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableAttributedString* result_ = [nSMutableAttributedString initWithAttributedString:(NSAttributedString*)attrStr];
    return result_;
}

void* C_NSMutableAttributedString_Init(void* ptr) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableAttributedString* result_ = [nSMutableAttributedString init];
    return result_;
}

void* C_NSMutableAttributedString_MutableString(void* ptr) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableString* result_ = [nSMutableAttributedString mutableString];
    return result_;
}
