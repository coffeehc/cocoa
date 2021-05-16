#import <Appkit/Appkit.h>
#import "touch.h"

void* C_Touch_Alloc() {
    return [NSTouch alloc];
}

void* C_NSTouch_Init(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSTouch* result = [nSTouch init];
    return result;
}

CGPoint C_NSTouch_LocationInView(void* ptr, void* view) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSPoint result = [nSTouch locationInView:(NSView*)view];
    return result;
}

CGPoint C_NSTouch_PreviousLocationInView(void* ptr, void* view) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSPoint result = [nSTouch previousLocationInView:(NSView*)view];
    return result;
}

int C_NSTouch_Type(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSTouchType result = [nSTouch type];
    return result;
}

unsigned int C_NSTouch_Phase(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSTouchPhase result = [nSTouch phase];
    return result;
}

CGPoint C_NSTouch_NormalizedPosition(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSPoint result = [nSTouch normalizedPosition];
    return result;
}

bool C_NSTouch_IsResting(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    BOOL result = [nSTouch isResting];
    return result;
}

void* C_NSTouch_Device(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    id result = [nSTouch device];
    return result;
}

CGSize C_NSTouch_DeviceSize(void* ptr) {
    NSTouch* nSTouch = (NSTouch*)ptr;
    NSSize result = [nSTouch deviceSize];
    return result;
}
