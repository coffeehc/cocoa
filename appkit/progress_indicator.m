#import <Appkit/Appkit.h>
#import "progress_indicator.h"

void* C_ProgressIndicator_Alloc() {
    return [NSProgressIndicator alloc];
}

void* C_NSProgressIndicator_InitWithFrame(void* ptr, CGRect frameRect) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSProgressIndicator* result = [nSProgressIndicator initWithFrame:frameRect];
    return result;
}

void* C_NSProgressIndicator_InitWithCoder(void* ptr, void* coder) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSProgressIndicator* result = [nSProgressIndicator initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSProgressIndicator_Init(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSProgressIndicator* result = [nSProgressIndicator init];
    return result;
}

void C_NSProgressIndicator_StartAnimation(void* ptr, void* sender) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator startAnimation:(id)sender];
}

void C_NSProgressIndicator_StopAnimation(void* ptr, void* sender) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator stopAnimation:(id)sender];
}

void C_NSProgressIndicator_IncrementBy(void* ptr, double delta) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator incrementBy:delta];
}

void C_NSProgressIndicator_SizeToFit(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator sizeToFit];
}

bool C_NSProgressIndicator_UsesThreadedAnimation(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    BOOL result = [nSProgressIndicator usesThreadedAnimation];
    return result;
}

void C_NSProgressIndicator_SetUsesThreadedAnimation(void* ptr, bool value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setUsesThreadedAnimation:value];
}

double C_NSProgressIndicator_DoubleValue(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    double result = [nSProgressIndicator doubleValue];
    return result;
}

void C_NSProgressIndicator_SetDoubleValue(void* ptr, double value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setDoubleValue:value];
}

double C_NSProgressIndicator_MinValue(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    double result = [nSProgressIndicator minValue];
    return result;
}

void C_NSProgressIndicator_SetMinValue(void* ptr, double value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setMinValue:value];
}

double C_NSProgressIndicator_MaxValue(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    double result = [nSProgressIndicator maxValue];
    return result;
}

void C_NSProgressIndicator_SetMaxValue(void* ptr, double value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setMaxValue:value];
}

unsigned int C_NSProgressIndicator_ControlSize(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSControlSize result = [nSProgressIndicator controlSize];
    return result;
}

void C_NSProgressIndicator_SetControlSize(void* ptr, unsigned int value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setControlSize:value];
}

unsigned int C_NSProgressIndicator_ControlTint(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSControlTint result = [nSProgressIndicator controlTint];
    return result;
}

void C_NSProgressIndicator_SetControlTint(void* ptr, unsigned int value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setControlTint:value];
}

bool C_NSProgressIndicator_IsBezeled(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    BOOL result = [nSProgressIndicator isBezeled];
    return result;
}

void C_NSProgressIndicator_SetBezeled(void* ptr, bool value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setBezeled:value];
}

bool C_NSProgressIndicator_IsIndeterminate(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    BOOL result = [nSProgressIndicator isIndeterminate];
    return result;
}

void C_NSProgressIndicator_SetIndeterminate(void* ptr, bool value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setIndeterminate:value];
}

unsigned int C_NSProgressIndicator_Style(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSProgressIndicatorStyle result = [nSProgressIndicator style];
    return result;
}

void C_NSProgressIndicator_SetStyle(void* ptr, unsigned int value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setStyle:value];
}

bool C_NSProgressIndicator_IsDisplayedWhenStopped(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    BOOL result = [nSProgressIndicator isDisplayedWhenStopped];
    return result;
}

void C_NSProgressIndicator_SetDisplayedWhenStopped(void* ptr, bool value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setDisplayedWhenStopped:value];
}
