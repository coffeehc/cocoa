#import "nib.h"
#import <AppKit/NSNib.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_Nib_Alloc() {
    return [NSNib alloc];
}

void* C_NSNib_InitWithNibNamed_Bundle(void* ptr, void* nibName, void* bundle) {
    NSNib* nSNib = (NSNib*)ptr;
    NSNib* result_ = [nSNib initWithNibNamed:(NSString*)nibName bundle:(NSBundle*)bundle];
    return result_;
}

void* C_NSNib_InitWithNibData_Bundle(void* ptr, void* nibData, void* bundle) {
    NSNib* nSNib = (NSNib*)ptr;
    NSNib* result_ = [nSNib initWithNibData:(NSData*)nibData bundle:(NSBundle*)bundle];
    return result_;
}

void* C_NSNib_AllocNib() {
    NSNib* result_ = [NSNib alloc];
    return result_;
}

void* C_NSNib_Init(void* ptr) {
    NSNib* nSNib = (NSNib*)ptr;
    NSNib* result_ = [nSNib init];
    return result_;
}

void* C_NSNib_NewNib() {
    NSNib* result_ = [NSNib new];
    return result_;
}

void* C_NSNib_Autorelease(void* ptr) {
    NSNib* nSNib = (NSNib*)ptr;
    NSNib* result_ = [nSNib autorelease];
    return result_;
}

void* C_NSNib_Retain(void* ptr) {
    NSNib* nSNib = (NSNib*)ptr;
    NSNib* result_ = [nSNib retain];
    return result_;
}
