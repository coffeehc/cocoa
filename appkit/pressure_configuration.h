#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_PressureConfiguration_Alloc();

void* C_NSPressureConfiguration_InitWithPressureBehavior(void* ptr, int pressureBehavior);
void* C_NSPressureConfiguration_Init(void* ptr);
void C_NSPressureConfiguration_Set(void* ptr);
int C_NSPressureConfiguration_PressureBehavior(void* ptr);
