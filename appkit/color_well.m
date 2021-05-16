#import <Appkit/Appkit.h>
#import "color_well.h"

void* C_ColorWell_Alloc() {
    return [NSColorWell alloc];
}

void* C_NSColorWell_InitWithFrame(void* ptr, CGRect frameRect) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    NSColorWell* result = [nSColorWell initWithFrame:frameRect];
    return result;
}

void* C_NSColorWell_InitWithCoder(void* ptr, void* coder) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    NSColorWell* result = [nSColorWell initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSColorWell_Init(void* ptr) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    NSColorWell* result = [nSColorWell init];
    return result;
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
    NSColor* result = [nSColorWell color];
    return result;
}

void C_NSColorWell_SetColor(void* ptr, void* value) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    [nSColorWell setColor:(NSColor*)value];
}

bool C_NSColorWell_IsActive(void* ptr) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    BOOL result = [nSColorWell isActive];
    return result;
}

bool C_NSColorWell_IsBordered(void* ptr) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    BOOL result = [nSColorWell isBordered];
    return result;
}

void C_NSColorWell_SetBordered(void* ptr, bool value) {
    NSColorWell* nSColorWell = (NSColorWell*)ptr;
    [nSColorWell setBordered:value];
}
