#import <Foundation/Foundation.h>
#import "mutable_attributed_string.h"

void* C_MutableAttributedString_Alloc() {
    return [NSMutableAttributedString alloc];
}

void* C_NSMutableAttributedString_InitWithString(void* ptr, void* str) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableAttributedString* result = [nSMutableAttributedString initWithString:(NSString*)str];
    return result;
}

void* C_NSMutableAttributedString_InitWithAttributedString(void* ptr, void* attrStr) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableAttributedString* result = [nSMutableAttributedString initWithAttributedString:(NSAttributedString*)attrStr];
    return result;
}

void* C_NSMutableAttributedString_Init(void* ptr) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableAttributedString* result = [nSMutableAttributedString init];
    return result;
}

void* C_NSMutableAttributedString_MutableString(void* ptr) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableString* result = [nSMutableAttributedString mutableString];
    return result;
}
