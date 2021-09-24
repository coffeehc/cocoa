#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_ProgressIndicator_Alloc();

void* C_NSProgressIndicator_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSProgressIndicator_InitWithCoder(void* ptr, void* coder);
void* C_NSProgressIndicator_Init(void* ptr);
void* C_NSProgressIndicator_AllocProgressIndicator();
void* C_NSProgressIndicator_NewProgressIndicator();
void* C_NSProgressIndicator_Autorelease(void* ptr);
void* C_NSProgressIndicator_Retain(void* ptr);
void C_NSProgressIndicator_StartAnimation(void* ptr, void* sender);
void C_NSProgressIndicator_StopAnimation(void* ptr, void* sender);
void C_NSProgressIndicator_IncrementBy(void* ptr, double delta);
void C_NSProgressIndicator_SizeToFit(void* ptr);
bool C_NSProgressIndicator_UsesThreadedAnimation(void* ptr);
void C_NSProgressIndicator_SetUsesThreadedAnimation(void* ptr, bool value);
double C_NSProgressIndicator_DoubleValue(void* ptr);
void C_NSProgressIndicator_SetDoubleValue(void* ptr, double value);
double C_NSProgressIndicator_MinValue(void* ptr);
void C_NSProgressIndicator_SetMinValue(void* ptr, double value);
double C_NSProgressIndicator_MaxValue(void* ptr);
void C_NSProgressIndicator_SetMaxValue(void* ptr, double value);
unsigned int C_NSProgressIndicator_ControlSize(void* ptr);
void C_NSProgressIndicator_SetControlSize(void* ptr, unsigned int value);
unsigned int C_NSProgressIndicator_ControlTint(void* ptr);
void C_NSProgressIndicator_SetControlTint(void* ptr, unsigned int value);
bool C_NSProgressIndicator_IsBezeled(void* ptr);
void C_NSProgressIndicator_SetBezeled(void* ptr, bool value);
bool C_NSProgressIndicator_IsIndeterminate(void* ptr);
void C_NSProgressIndicator_SetIndeterminate(void* ptr, bool value);
unsigned int C_NSProgressIndicator_Style(void* ptr);
void C_NSProgressIndicator_SetStyle(void* ptr, unsigned int value);
bool C_NSProgressIndicator_IsDisplayedWhenStopped(void* ptr);
void C_NSProgressIndicator_SetDisplayedWhenStopped(void* ptr, bool value);
