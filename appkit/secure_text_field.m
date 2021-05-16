#import <Appkit/Appkit.h>
#import "secure_text_field.h"

void* C_SecureTextField_Alloc() {
    return [NSSecureTextField alloc];
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
