#import <Appkit/Appkit.h>
#import "stepper.h"

void* C_Stepper_Alloc() {
    return [NSStepper alloc];
}

void* C_NSStepper_InitWithFrame(void* ptr, CGRect frameRect) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    NSStepper* result = [nSStepper initWithFrame:frameRect];
    return result;
}

void* C_NSStepper_InitWithCoder(void* ptr, void* coder) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    NSStepper* result = [nSStepper initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSStepper_Init(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    NSStepper* result = [nSStepper init];
    return result;
}

double C_NSStepper_MaxValue(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    double result = [nSStepper maxValue];
    return result;
}

void C_NSStepper_SetMaxValue(void* ptr, double value) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    [nSStepper setMaxValue:value];
}

double C_NSStepper_MinValue(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    double result = [nSStepper minValue];
    return result;
}

void C_NSStepper_SetMinValue(void* ptr, double value) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    [nSStepper setMinValue:value];
}

double C_NSStepper_Increment(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    double result = [nSStepper increment];
    return result;
}

void C_NSStepper_SetIncrement(void* ptr, double value) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    [nSStepper setIncrement:value];
}

bool C_NSStepper_Autorepeat(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    BOOL result = [nSStepper autorepeat];
    return result;
}

void C_NSStepper_SetAutorepeat(void* ptr, bool value) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    [nSStepper setAutorepeat:value];
}

bool C_NSStepper_ValueWraps(void* ptr) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    BOOL result = [nSStepper valueWraps];
    return result;
}

void C_NSStepper_SetValueWraps(void* ptr, bool value) {
    NSStepper* nSStepper = (NSStepper*)ptr;
    [nSStepper setValueWraps:value];
}
