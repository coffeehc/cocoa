#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Switch_Alloc();

void* C_NSSwitch_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSSwitch_InitWithCoder(void* ptr, void* coder);
void* C_NSSwitch_Init(void* ptr);
int C_NSSwitch_State(void* ptr);
void C_NSSwitch_SetState(void* ptr, int value);
