#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_PressureConfiguration_Alloc();

void* C_NSPressureConfiguration_InitWithPressureBehavior(void* ptr, int pressureBehavior);
void* C_NSPressureConfiguration_AllocPressureConfiguration();
void* C_NSPressureConfiguration_Init(void* ptr);
void* C_NSPressureConfiguration_NewPressureConfiguration();
void* C_NSPressureConfiguration_Autorelease(void* ptr);
void* C_NSPressureConfiguration_Retain(void* ptr);
void C_NSPressureConfiguration_Set(void* ptr);
int C_NSPressureConfiguration_PressureBehavior(void* ptr);
