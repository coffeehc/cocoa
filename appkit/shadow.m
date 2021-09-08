#import <Appkit/Appkit.h>
#import "shadow.h"

void* C_Shadow_Alloc() {
    return [NSShadow alloc];
}

void* C_NSShadow_Init(void* ptr) {
    NSShadow* nSShadow = (NSShadow*)ptr;
    NSShadow* result_ = [nSShadow init];
    return result_;
}

void* C_NSShadow_AllocShadow() {
    NSShadow* result_ = [NSShadow alloc];
    return result_;
}

void* C_NSShadow_NewShadow() {
    NSShadow* result_ = [NSShadow new];
    return result_;
}

void* C_NSShadow_Autorelease(void* ptr) {
    NSShadow* nSShadow = (NSShadow*)ptr;
    NSShadow* result_ = [nSShadow autorelease];
    return result_;
}

void* C_NSShadow_Retain(void* ptr) {
    NSShadow* nSShadow = (NSShadow*)ptr;
    NSShadow* result_ = [nSShadow retain];
    return result_;
}

void C_NSShadow_Set(void* ptr) {
    NSShadow* nSShadow = (NSShadow*)ptr;
    [nSShadow set];
}

CGSize C_NSShadow_ShadowOffset(void* ptr) {
    NSShadow* nSShadow = (NSShadow*)ptr;
    NSSize result_ = [nSShadow shadowOffset];
    return result_;
}

void C_NSShadow_SetShadowOffset(void* ptr, CGSize value) {
    NSShadow* nSShadow = (NSShadow*)ptr;
    [nSShadow setShadowOffset:value];
}

double C_NSShadow_ShadowBlurRadius(void* ptr) {
    NSShadow* nSShadow = (NSShadow*)ptr;
    CGFloat result_ = [nSShadow shadowBlurRadius];
    return result_;
}

void C_NSShadow_SetShadowBlurRadius(void* ptr, double value) {
    NSShadow* nSShadow = (NSShadow*)ptr;
    [nSShadow setShadowBlurRadius:value];
}

void* C_NSShadow_ShadowColor(void* ptr) {
    NSShadow* nSShadow = (NSShadow*)ptr;
    NSColor* result_ = [nSShadow shadowColor];
    return result_;
}

void C_NSShadow_SetShadowColor(void* ptr, void* value) {
    NSShadow* nSShadow = (NSShadow*)ptr;
    [nSShadow setShadowColor:(NSColor*)value];
}
