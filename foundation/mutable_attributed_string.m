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

void* C_NSMutableAttributedString_InitWithString_Attributes(void* ptr, void* str, Dictionary attrs) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableDictionary* objcAttrs = [NSMutableDictionary dictionaryWithCapacity:attrs.len];
    if (attrs.len > 0) {
    	void** attrsKeyData = (void**)attrs.key_data;
    	void** attrsValueData = (void**)attrs.value_data;
    	for (int i = 0; i < attrs.len; i++) {
    		void* kp = attrsKeyData[i];
    		void* vp = attrsValueData[i];
    		[objcAttrs setObject:(NSAttributedStringKey)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    NSMutableAttributedString* result_ = [nSMutableAttributedString initWithString:(NSString*)str attributes:objcAttrs];
    return result_;
}

void* C_NSMutableAttributedString_InitWithAttributedString(void* ptr, void* attrStr) {
    NSMutableAttributedString* nSMutableAttributedString = (NSMutableAttributedString*)ptr;
    NSMutableAttributedString* result_ = [nSMutableAttributedString initWithAttributedString:(NSAttributedString*)attrStr];
    return result_;
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
