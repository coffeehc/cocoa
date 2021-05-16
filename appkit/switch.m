#import <Appkit/Appkit.h>
#import "switch.h"

void* C_Switch_Alloc() {
    return [NSSwitch alloc];
}

void* C_NSSwitch_InitWithFrame(void* ptr, CGRect frameRect) {
    NSSwitch* nSSwitch = (NSSwitch*)ptr;
    NSSwitch* result = [nSSwitch initWithFrame:frameRect];
    return result;
}

void* C_NSSwitch_InitWithCoder(void* ptr, void* coder) {
    NSSwitch* nSSwitch = (NSSwitch*)ptr;
    NSSwitch* result = [nSSwitch initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSSwitch_Init(void* ptr) {
    NSSwitch* nSSwitch = (NSSwitch*)ptr;
    NSSwitch* result = [nSSwitch init];
    return result;
}

int C_NSSwitch_State(void* ptr) {
    NSSwitch* nSSwitch = (NSSwitch*)ptr;
    NSControlStateValue result = [nSSwitch state];
    return result;
}

void C_NSSwitch_SetState(void* ptr, int value) {
    NSSwitch* nSSwitch = (NSSwitch*)ptr;
    [nSSwitch setState:value];
}
