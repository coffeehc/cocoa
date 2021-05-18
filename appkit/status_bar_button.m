#import <Appkit/Appkit.h>
#import "status_bar_button.h"

void* C_StatusBarButton_Alloc() {
    return [NSStatusBarButton alloc];
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

bool C_NSStatusBarButton_AppearsDisabled(void* ptr) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    BOOL result_ = [nSStatusBarButton appearsDisabled];
    return result_;
}

void C_NSStatusBarButton_SetAppearsDisabled(void* ptr, bool value) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    [nSStatusBarButton setAppearsDisabled:value];
}
