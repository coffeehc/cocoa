#import "color_well.h"
#import <AppKit/NSColorWell.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_ColorWell_Alloc() {
    return [NSColorWell alloc];
}

void* C_NSColorWell_InitWithFrame(void* ptr, CGRect frameRect) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    NSColorWell* result_ = [nSColorWell initWithFrame:frameRect];
    return result_;
}

void* C_NSColorWell_InitWithCoder(void* ptr, void* coder) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    NSColorWell* result_ = [nSColorWell initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSColorWell_Init(void* ptr) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    NSColorWell* result_ = [nSColorWell init];
    return result_;
}

void* C_NSColorWell_AllocColorWell() {
    NSColorWell* result_ = [NSColorWell alloc];
    return result_;
}

void* C_NSColorWell_NewColorWell() {
    NSColorWell* result_ = [NSColorWell new];
    return result_;
}

void* C_NSColorWell_Autorelease(void* ptr) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    NSColorWell* result_ = [nSColorWell autorelease];
    return result_;
}

void* C_NSColorWell_Retain(void* ptr) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    NSColorWell* result_ = [nSColorWell retain];
    return result_;
}

void C_NSColorWell_TakeColorFrom(void* ptr, void* sender) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    [nSColorWell takeColorFrom:(id)sender];
}

void C_NSColorWell_Activate(void* ptr, bool exclusive) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    [nSColorWell activate:exclusive];
}

void C_NSColorWell_Deactivate(void* ptr) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    [nSColorWell deactivate];
}

void C_NSColorWell_DrawWellInside(void* ptr, CGRect insideRect) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    [nSColorWell drawWellInside:insideRect];
}

void* C_NSColorWell_Color(void* ptr) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    NSColor* result_ = [nSColorWell color];
    return result_;
}

void C_NSColorWell_SetColor(void* ptr, void* value) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    [nSColorWell setColor:(NSColor*)value];
}

bool C_NSColorWell_IsActive(void* ptr) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    BOOL result_ = [nSColorWell isActive];
    return result_;
}

bool C_NSColorWell_IsBordered(void* ptr) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    BOOL result_ = [nSColorWell isBordered];
    return result_;
}

void C_NSColorWell_SetBordered(void* ptr, bool value) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    [nSColorWell setBordered:value];
}
