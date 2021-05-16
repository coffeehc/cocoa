#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_SecureTextField_Alloc();

void* C_NSSecureTextField_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSSecureTextField_InitWithCoder(void* ptr, void* coder);
void* C_NSSecureTextField_Init(void* ptr);
