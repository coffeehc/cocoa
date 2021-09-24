#import "secure_text_field.h"
#import <AppKit/NSSecureTextField.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_SecureTextField_Alloc() {
    return [NSSecureTextField alloc];
}

void* C_NSSecureTextField_SecureTextField_LabelWithAttributedString(void* attributedStringValue) {
    NSSecureTextField* result_ = [NSSecureTextField labelWithAttributedString:(NSAttributedString*)attributedStringValue];
    return result_;
}

void* C_NSSecureTextField_SecureTextField_LabelWithString(void* stringValue) {
    NSSecureTextField* result_ = [NSSecureTextField labelWithString:(NSString*)stringValue];
    return result_;
}

void* C_NSSecureTextField_SecureTextField_TextFieldWithString(void* stringValue) {
    NSSecureTextField* result_ = [NSSecureTextField textFieldWithString:(NSString*)stringValue];
    return result_;
}

void* C_NSSecureTextField_SecureTextField_WrappingLabelWithString(void* stringValue) {
    NSSecureTextField* result_ = [NSSecureTextField wrappingLabelWithString:(NSString*)stringValue];
    return result_;
}

void* C_NSSecureTextField_InitWithFrame(void* ptr, CGRect frameRect) {
    NSSecureTextField* nSSecureTextField = (NSSecureTextField*)ptr;
    NSSecureTextField* result_ = [nSSecureTextField initWithFrame:frameRect];
    return result_;
}

void* C_NSSecureTextField_InitWithCoder(void* ptr, void* coder) {
    NSSecureTextField* nSSecureTextField = (NSSecureTextField*)ptr;
    NSSecureTextField* result_ = [nSSecureTextField initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSSecureTextField_Init(void* ptr) {
    NSSecureTextField* nSSecureTextField = (NSSecureTextField*)ptr;
    NSSecureTextField* result_ = [nSSecureTextField init];
    return result_;
}

void* C_NSSecureTextField_AllocSecureTextField() {
    NSSecureTextField* result_ = [NSSecureTextField alloc];
    return result_;
}

void* C_NSSecureTextField_NewSecureTextField() {
    NSSecureTextField* result_ = [NSSecureTextField new];
    return result_;
}

void* C_NSSecureTextField_Autorelease(void* ptr) {
    NSSecureTextField* nSSecureTextField = (NSSecureTextField*)ptr;
    NSSecureTextField* result_ = [nSSecureTextField autorelease];
    return result_;
}

void* C_NSSecureTextField_Retain(void* ptr) {
    NSSecureTextField* nSSecureTextField = (NSSecureTextField*)ptr;
    NSSecureTextField* result_ = [nSSecureTextField retain];
    return result_;
}
