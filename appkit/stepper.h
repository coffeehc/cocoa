#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Stepper_Alloc();

void* C_NSStepper_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSStepper_InitWithCoder(void* ptr, void* coder);
void* C_NSStepper_Init(void* ptr);
void* C_NSStepper_AllocStepper();
void* C_NSStepper_NewStepper();
void* C_NSStepper_Autorelease(void* ptr);
void* C_NSStepper_Retain(void* ptr);
double C_NSStepper_MaxValue(void* ptr);
void C_NSStepper_SetMaxValue(void* ptr, double value);
double C_NSStepper_MinValue(void* ptr);
void C_NSStepper_SetMinValue(void* ptr, double value);
double C_NSStepper_Increment(void* ptr);
void C_NSStepper_SetIncrement(void* ptr, double value);
bool C_NSStepper_Autorepeat(void* ptr);
void C_NSStepper_SetAutorepeat(void* ptr, bool value);
bool C_NSStepper_ValueWraps(void* ptr);
void C_NSStepper_SetValueWraps(void* ptr, bool value);
