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

void* C_NSFormatter_AttributedStringForObjectValue_WithDefaultAttributes(void* ptr, void* obj, Dictionary attrs) {
    NSFormatter* nSFormatter = (NSFormatter*)ptr;
    NSMutableDictionary* objcAttrs = [[NSMutableDictionary alloc] initWithCapacity: attrs.len];
    if (attrs.len > 0) {
    	void** attrsKeyData = (void**)attrs.key_data;
    	void** attrsValueData = (void**)attrs.value_data;
    	for (int i = 0; i < attrs.len; i++) {
    		void* kp = attrsKeyData[i];
    		void* vp = attrsValueData[i];
    		[objcAttrs setObject:(NSAttributedStringKey)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    NSAttributedString* result_ = [nSFormatter attributedStringForObjectValue:(id)obj withDefaultAttributes:objcAttrs];
    return result_;
}

void* C_NSFormatter_EditingStringForObjectValue(void* ptr, void* obj) {
    NSFormatter* nSFormatter = (NSFormatter*)ptr;
    NSString* result_ = [nSFormatter editingStringForObjectValue:(id)obj];
    return result_;
}
