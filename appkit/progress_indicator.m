#import <Appkit/Appkit.h>
#import "progress_indicator.h"

void* C_ProgressIndicator_Alloc() {
    return [NSProgressIndicator alloc];
}

void* C_NSProgressIndicator_InitWithFrame(void* ptr, CGRect frameRect) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSProgressIndicator* result_ = [nSProgressIndicator initWithFrame:frameRect];
    return result_;
}

void* C_NSProgressIndicator_InitWithCoder(void* ptr, void* coder) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSProgressIndicator* result_ = [nSProgressIndicator initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSProgressIndicator_Init(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSProgressIndicator* result_ = [nSProgressIndicator init];
    return result_;
}

void* C_NSProgressIndicator_AllocProgressIndicator() {
    NSProgressIndicator* result_ = [NSProgressIndicator alloc];
    return result_;
}

void* C_NSProgressIndicator_NewProgressIndicator() {
    NSProgressIndicator* result_ = [NSProgressIndicator new];
    return result_;
}

void* C_NSProgressIndicator_Autorelease(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSProgressIndicator* result_ = [nSProgressIndicator autorelease];
    return result_;
}

void* C_NSProgressIndicator_Retain(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSProgressIndicator* result_ = [nSProgressIndicator retain];
    return result_;
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
    BOOL result_ = [nSProgressIndicator usesThreadedAnimation];
    return result_;
}

void C_NSProgressIndicator_SetUsesThreadedAnimation(void* ptr, bool value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setUsesThreadedAnimation:value];
}

double C_NSProgressIndicator_DoubleValue(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    double result_ = [nSProgressIndicator doubleValue];
    return result_;
}

void C_NSProgressIndicator_SetDoubleValue(void* ptr, double value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setDoubleValue:value];
}

double C_NSProgressIndicator_MinValue(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    double result_ = [nSProgressIndicator minValue];
    return result_;
}

void C_NSProgressIndicator_SetMinValue(void* ptr, double value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setMinValue:value];
}

double C_NSProgressIndicator_MaxValue(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    double result_ = [nSProgressIndicator maxValue];
    return result_;
}

void C_NSProgressIndicator_SetMaxValue(void* ptr, double value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setMaxValue:value];
}

unsigned int C_NSProgressIndicator_ControlSize(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSControlSize result_ = [nSProgressIndicator controlSize];
    return result_;
}

void C_NSProgressIndicator_SetControlSize(void* ptr, unsigned int value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setControlSize:value];
}

unsigned int C_NSProgressIndicator_ControlTint(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSControlTint result_ = [nSProgressIndicator controlTint];
    return result_;
}

void C_NSProgressIndicator_SetControlTint(void* ptr, unsigned int value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setControlTint:value];
}

bool C_NSProgressIndicator_IsBezeled(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    BOOL result_ = [nSProgressIndicator isBezeled];
    return result_;
}

void C_NSProgressIndicator_SetBezeled(void* ptr, bool value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setBezeled:value];
}

bool C_NSProgressIndicator_IsIndeterminate(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    BOOL result_ = [nSProgressIndicator isIndeterminate];
    return result_;
}

void C_NSProgressIndicator_SetIndeterminate(void* ptr, bool value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setIndeterminate:value];
}

unsigned int C_NSProgressIndicator_Style(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    NSProgressIndicatorStyle result_ = [nSProgressIndicator style];
    return result_;
}

void C_NSProgressIndicator_SetStyle(void* ptr, unsigned int value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setStyle:value];
}

bool C_NSProgressIndicator_IsDisplayedWhenStopped(void* ptr) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    BOOL result_ = [nSProgressIndicator isDisplayedWhenStopped];
    return result_;
}

void C_NSProgressIndicator_SetDisplayedWhenStopped(void* ptr, bool value) {
    NSProgressIndicator* nSProgressIndicator = (NSProgressIndicator*)ptr;
    [nSProgressIndicator setDisplayedWhenStopped:value];
}
