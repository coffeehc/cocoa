#import <Appkit/Appkit.h>
#import "panel.h"

void* C_Panel_Alloc() {
    return [NSPanel alloc];
}

void* C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag) {
    NSPanel* nSPanel = (NSPanel*)ptr;
    NSPanel* result_ = [nSPanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag];
    return result_;
}

void* C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen) {
    NSPanel* nSPanel = (NSPanel*)ptr;
    NSPanel* result_ = [nSPanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag screen:(NSScreen*)screen];
    return result_;
}

void* C_NSPanel_Init(void* ptr) {
    NSPanel* nSPanel = (NSPanel*)ptr;
    NSPanel* result_ = [nSPanel init];
    return result_;
}

bool C_NSPanel_BecomesKeyOnlyIfNeeded(void* ptr) {
    NSPanel* nSPanel = (NSPanel*)ptr;
    BOOL result_ = [nSPanel becomesKeyOnlyIfNeeded];
    return result_;
}

void C_NSPanel_SetBecomesKeyOnlyIfNeeded(void* ptr, bool value) {
    NSPanel* nSPanel = (NSPanel*)ptr;
    [nSPanel setBecomesKeyOnlyIfNeeded:value];
}
