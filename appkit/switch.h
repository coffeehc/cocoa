#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Switch_Alloc();

void* C_NSSwitch_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSSwitch_InitWithCoder(void* ptr, void* coder);
void* C_NSSwitch_Init(void* ptr);
void* C_NSSwitch_AllocSwitch();
void* C_NSSwitch_NewSwitch();
void* C_NSSwitch_Autorelease(void* ptr);
void* C_NSSwitch_Retain(void* ptr);
int C_NSSwitch_State(void* ptr);
void C_NSSwitch_SetState(void* ptr, int value);
