#import "status_bar_button.h"
#import <AppKit/NSStatusBarButton.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_StatusBarButton_Alloc() {
    return [NSStatusBarButton alloc];
}

void* C_NSStatusBarButton_StatusBarButton_CheckboxWithTitle_Target_Action(void* title, void* target, void* action) {
    NSStatusBarButton* result_ = [NSStatusBarButton checkboxWithTitle:(NSString*)title target:(id)target action:(SEL)action];
    return result_;
}

void* C_NSStatusBarButton_StatusBarButton_ButtonWithImage_Target_Action(void* image, void* target, void* action) {
    NSStatusBarButton* result_ = [NSStatusBarButton buttonWithImage:(NSImage*)image target:(id)target action:(SEL)action];
    return result_;
}

void* C_NSStatusBarButton_StatusBarButton_RadioButtonWithTitle_Target_Action(void* title, void* target, void* action) {
    NSStatusBarButton* result_ = [NSStatusBarButton radioButtonWithTitle:(NSString*)title target:(id)target action:(SEL)action];
    return result_;
}

void* C_NSStatusBarButton_StatusBarButton_ButtonWithTitle_Image_Target_Action(void* title, void* image, void* target, void* action) {
    NSStatusBarButton* result_ = [NSStatusBarButton buttonWithTitle:(NSString*)title image:(NSImage*)image target:(id)target action:(SEL)action];
    return result_;
}

void* C_NSStatusBarButton_StatusBarButton_ButtonWithTitle_Target_Action(void* title, void* target, void* action) {
    NSStatusBarButton* result_ = [NSStatusBarButton buttonWithTitle:(NSString*)title target:(id)target action:(SEL)action];
    return result_;
}

void* C_NSStatusBarButton_InitWithFrame(void* ptr, CGRect frameRect) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    NSStatusBarButton* result_ = [nSStatusBarButton initWithFrame:frameRect];
    return result_;
}

void* C_NSStatusBarButton_InitWithCoder(void* ptr, void* coder) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    NSStatusBarButton* result_ = [nSStatusBarButton initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSStatusBarButton_Init(void* ptr) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    NSStatusBarButton* result_ = [nSStatusBarButton init];
    return result_;
}

void* C_NSStatusBarButton_AllocStatusBarButton() {
    NSStatusBarButton* result_ = [NSStatusBarButton alloc];
    return result_;
}

void* C_NSStatusBarButton_NewStatusBarButton() {
    NSStatusBarButton* result_ = [NSStatusBarButton new];
    return result_;
}

void* C_NSStatusBarButton_Autorelease(void* ptr) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    NSStatusBarButton* result_ = [nSStatusBarButton autorelease];
    return result_;
}

void* C_NSStatusBarButton_Retain(void* ptr) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    NSStatusBarButton* result_ = [nSStatusBarButton retain];
    return result_;
}

bool C_NSStatusBarButton_AppearsDisabled(void* ptr) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    BOOL result_ = [nSStatusBarButton appearsDisabled];
    return result_;
}

void C_NSStatusBarButton_SetAppearsDisabled(void* ptr, bool value) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    [nSStatusBarButton setAppearsDisabled:value];
}
