#import <Appkit/Appkit.h>
#import "status_bar_button.h"

void* C_StatusBarButton_Alloc() {
    return [NSStatusBarButton alloc];
}

void* C_NSStatusBarButton_InitWithFrame(void* ptr, CGRect frameRect) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    NSStatusBarButton* result = [nSStatusBarButton initWithFrame:frameRect];
    return result;
}

void* C_NSStatusBarButton_InitWithCoder(void* ptr, void* coder) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    NSStatusBarButton* result = [nSStatusBarButton initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSStatusBarButton_Init(void* ptr) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    NSStatusBarButton* result = [nSStatusBarButton init];
    return result;
}

bool C_NSStatusBarButton_AppearsDisabled(void* ptr) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    BOOL result = [nSStatusBarButton appearsDisabled];
    return result;
}

void C_NSStatusBarButton_SetAppearsDisabled(void* ptr, bool value) {
    NSStatusBarButton* nSStatusBarButton = (NSStatusBarButton*)ptr;
    [nSStatusBarButton setAppearsDisabled:value];
}
