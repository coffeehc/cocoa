#import <Appkit/Appkit.h>
#import "box.h"

void* C_Box_Alloc() {
    return [NSBox alloc];
}

void* C_NSBox_InitWithFrame(void* ptr, CGRect frameRect) {
    NSBox* nSBox = (NSBox*)ptr;
    NSBox* result = [nSBox initWithFrame:frameRect];
    return result;
}

void* C_NSBox_InitWithCoder(void* ptr, void* coder) {
    NSBox* nSBox = (NSBox*)ptr;
    NSBox* result = [nSBox initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSBox_Init(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSBox* result = [nSBox init];
    return result;
}

void C_NSBox_SetFrameFromContentFrame(void* ptr, CGRect contentFrame) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setFrameFromContentFrame:contentFrame];
}

void C_NSBox_SizeToFit(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox sizeToFit];
}

CGRect C_NSBox_BorderRect(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSRect result = [nSBox borderRect];
    return result;
}

unsigned int C_NSBox_BoxType(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSBoxType result = [nSBox boxType];
    return result;
}

void C_NSBox_SetBoxType(void* ptr, unsigned int value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setBoxType:value];
}

bool C_NSBox_IsTransparent(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    BOOL result = [nSBox isTransparent];
    return result;
}

void C_NSBox_SetTransparent(void* ptr, bool value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setTransparent:value];
}

void* C_NSBox_Title(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSString* result = [nSBox title];
    return result;
}

void C_NSBox_SetTitle(void* ptr, void* value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setTitle:(NSString*)value];
}

void* C_NSBox_TitleFont(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSFont* result = [nSBox titleFont];
    return result;
}

void C_NSBox_SetTitleFont(void* ptr, void* value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setTitleFont:(NSFont*)value];
}

unsigned int C_NSBox_TitlePosition(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSTitlePosition result = [nSBox titlePosition];
    return result;
}

void C_NSBox_SetTitlePosition(void* ptr, unsigned int value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setTitlePosition:value];
}

void* C_NSBox_TitleCell(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    id result = [nSBox titleCell];
    return result;
}

CGRect C_NSBox_TitleRect(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSRect result = [nSBox titleRect];
    return result;
}

void* C_NSBox_BorderColor(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSColor* result = [nSBox borderColor];
    return result;
}

void C_NSBox_SetBorderColor(void* ptr, void* value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setBorderColor:(NSColor*)value];
}

double C_NSBox_BorderWidth(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    CGFloat result = [nSBox borderWidth];
    return result;
}

void C_NSBox_SetBorderWidth(void* ptr, double value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setBorderWidth:value];
}

double C_NSBox_CornerRadius(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    CGFloat result = [nSBox cornerRadius];
    return result;
}

void C_NSBox_SetCornerRadius(void* ptr, double value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setCornerRadius:value];
}

void* C_NSBox_FillColor(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSColor* result = [nSBox fillColor];
    return result;
}

void C_NSBox_SetFillColor(void* ptr, void* value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setFillColor:(NSColor*)value];
}

void* C_NSBox_ContentView(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSView* result = [nSBox contentView];
    return result;
}

void C_NSBox_SetContentView(void* ptr, void* value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setContentView:(NSView*)value];
}

CGSize C_NSBox_ContentViewMargins(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSSize result = [nSBox contentViewMargins];
    return result;
}

void C_NSBox_SetContentViewMargins(void* ptr, CGSize value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setContentViewMargins:value];
}
