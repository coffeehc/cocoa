#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

double Stepper_MaxValue(void* ptr);
void Stepper_SetMaxValue(void* ptr, double maxValue);
double Stepper_MinValue(void* ptr);
void Stepper_SetMinValue(void* ptr, double minValue);
double Stepper_Increment(void* ptr);
void Stepper_SetIncrement(void* ptr, double increment);
bool Stepper_Autorepeat(void* ptr);
void Stepper_SetAutorepeat(void* ptr, bool autorepeat);
bool Stepper_ValueWraps(void* ptr);
void Stepper_SetValueWraps(void* ptr, bool valueWraps);

void* Stepper_NewStepper(NSRect frame);
