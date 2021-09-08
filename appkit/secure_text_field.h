#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_SecureTextField_Alloc();

void* C_NSSecureTextField_SecureTextField_LabelWithAttributedString(void* attributedStringValue);
void* C_NSSecureTextField_SecureTextField_LabelWithString(void* stringValue);
void* C_NSSecureTextField_SecureTextField_TextFieldWithString(void* stringValue);
void* C_NSSecureTextField_SecureTextField_WrappingLabelWithString(void* stringValue);
void* C_NSSecureTextField_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSSecureTextField_InitWithCoder(void* ptr, void* coder);
void* C_NSSecureTextField_Init(void* ptr);
void* C_NSSecureTextField_AllocSecureTextField();
void* C_NSSecureTextField_NewSecureTextField();
void* C_NSSecureTextField_Autorelease(void* ptr);
void* C_NSSecureTextField_Retain(void* ptr);
