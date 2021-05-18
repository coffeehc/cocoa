#import <Appkit/Appkit.h>
#import "stepper.h"

void* C_Stepper_Alloc() {
    return [NSStepper alloc];
}

void* C_NSStepper_InitWithFrame(void* ptr, CGRect frameRect) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    NSStepper* result_ = [nSStepper initWithFrame:frameRect];
    return result_;
}

void* C_NSStepper_InitWithCoder(void* ptr, void* coder) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    NSStepper* result_ = [nSStepper initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSStepper_Init(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    NSStepper* result_ = [nSStepper init];
    return result_;
}

double C_NSStepper_MaxValue(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    double result_ = [nSStepper maxValue];
    return result_;
}

void C_NSStepper_SetMaxValue(void* ptr, double value) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    [nSStepper setMaxValue:value];
}

double C_NSStepper_MinValue(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    double result_ = [nSStepper minValue];
    return result_;
}

void C_NSStepper_SetMinValue(void* ptr, double value) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    [nSStepper setMinValue:value];
}

double C_NSStepper_Increment(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    double result_ = [nSStepper increment];
    return result_;
}

void C_NSStepper_SetIncrement(void* ptr, double value) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    [nSStepper setIncrement:value];
}

bool C_NSStepper_Autorepeat(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    BOOL result_ = [nSStepper autorepeat];
    return result_;
}

void C_NSStepper_SetAutorepeat(void* ptr, bool value) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    [nSStepper setAutorepeat:value];
}

bool C_NSStepper_ValueWraps(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    BOOL result_ = [nSStepper valueWraps];
    return result_;
}

void C_NSStepper_SetValueWraps(void* ptr, bool value) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    [nSStepper setValueWraps:value];
}
