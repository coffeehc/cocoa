#import <Appkit/Appkit.h>
#import "box.h"

void* C_Box_Alloc() {
    return [NSBox alloc];
}

void* C_NSBox_InitWithFrame(void* ptr, CGRect frameRect) {
    NSBox* nSBox = (NSBox*)ptr;
    NSBox* result_ = [nSBox initWithFrame:frameRect];
    return result_;
}

void* C_NSBox_InitWithCoder(void* ptr, void* coder) {
    NSBox* nSBox = (NSBox*)ptr;
    NSBox* result_ = [nSBox initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSBox_Init(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSBox* result_ = [nSBox init];
    return result_;
}

void* C_NSBox_AllocBox() {
    NSBox* result_ = [NSBox alloc];
    return result_;
}

void* C_NSBox_NewBox() {
    NSBox* result_ = [NSBox new];
    return result_;
}

void* C_NSBox_Autorelease(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSBox* result_ = [nSBox autorelease];
    return result_;
}

void* C_NSBox_Retain(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSBox* result_ = [nSBox retain];
    return result_;
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
    NSRect result_ = [nSBox borderRect];
    return result_;
}

unsigned int C_NSBox_BoxType(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSBoxType result_ = [nSBox boxType];
    return result_;
}

void C_NSBox_SetBoxType(void* ptr, unsigned int value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setBoxType:value];
}

bool C_NSBox_IsTransparent(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    BOOL result_ = [nSBox isTransparent];
    return result_;
}

void C_NSBox_SetTransparent(void* ptr, bool value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setTransparent:value];
}

void* C_NSBox_Title(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSString* result_ = [nSBox title];
    return result_;
}

void C_NSBox_SetTitle(void* ptr, void* value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setTitle:(NSString*)value];
}

void* C_NSBox_TitleFont(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSFont* result_ = [nSBox titleFont];
    return result_;
}

void C_NSBox_SetTitleFont(void* ptr, void* value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setTitleFont:(NSFont*)value];
}

unsigned int C_NSBox_TitlePosition(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSTitlePosition result_ = [nSBox titlePosition];
    return result_;
}

void C_NSBox_SetTitlePosition(void* ptr, unsigned int value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setTitlePosition:value];
}

void* C_NSBox_TitleCell(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    id result_ = [nSBox titleCell];
    return result_;
}

CGRect C_NSBox_TitleRect(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSRect result_ = [nSBox titleRect];
    return result_;
}

void* C_NSBox_BorderColor(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSColor* result_ = [nSBox borderColor];
    return result_;
}

void C_NSBox_SetBorderColor(void* ptr, void* value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setBorderColor:(NSColor*)value];
}

double C_NSBox_BorderWidth(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    CGFloat result_ = [nSBox borderWidth];
    return result_;
}

void C_NSBox_SetBorderWidth(void* ptr, double value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setBorderWidth:value];
}

double C_NSBox_CornerRadius(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    CGFloat result_ = [nSBox cornerRadius];
    return result_;
}

void C_NSBox_SetCornerRadius(void* ptr, double value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setCornerRadius:value];
}

void* C_NSBox_FillColor(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSColor* result_ = [nSBox fillColor];
    return result_;
}

void C_NSBox_SetFillColor(void* ptr, void* value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setFillColor:(NSColor*)value];
}

void* C_NSBox_ContentView(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSView* result_ = [nSBox contentView];
    return result_;
}

void C_NSBox_SetContentView(void* ptr, void* value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setContentView:(NSView*)value];
}

CGSize C_NSBox_ContentViewMargins(void* ptr) {
    NSBox* nSBox = (NSBox*)ptr;
    NSSize result_ = [nSBox contentViewMargins];
    return result_;
}

void C_NSBox_SetContentViewMargins(void* ptr, CGSize value) {
    NSBox* nSBox = (NSBox*)ptr;
    [nSBox setContentViewMargins:value];
}
