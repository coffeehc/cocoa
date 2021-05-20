#import <Appkit/Appkit.h>
#import "nib.h"

void* C_Nib_Alloc() {
    return [NSNib alloc];
}

void* C_NSNib_InitWithNibNamed_Bundle(void* ptr, void* nibName, void* bundle) {
    NSNib* nSNib = (NSNib*)ptr;
    NSNib* result_ = [nSNib initWithNibNamed:(NSString*)nibName bundle:(NSBundle*)bundle];
    return result_;
}

void* C_NSNib_InitWithNibData_Bundle(void* ptr, Array nibData, void* bundle) {
    NSNib* nSNib = (NSNib*)ptr;
    NSNib* result_ = [nSNib initWithNibData:[[NSData alloc] initWithBytes:(Byte *)nibData.data length:nibData.len] bundle:(NSBundle*)bundle];
    return result_;
}

void* C_NSNib_Init(void* ptr) {
    NSNib* nSNib = (NSNib*)ptr;
    NSNib* result_ = [nSNib init];
    return result_;
}
