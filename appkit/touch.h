#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Touch_Alloc();

void* C_NSTouch_Init(void* ptr);
CGPoint C_NSTouch_LocationInView(void* ptr, void* view);
CGPoint C_NSTouch_PreviousLocationInView(void* ptr, void* view);
int C_NSTouch_Type(void* ptr);
unsigned int C_NSTouch_Phase(void* ptr);
CGPoint C_NSTouch_NormalizedPosition(void* ptr);
bool C_NSTouch_IsResting(void* ptr);
void* C_NSTouch_Device(void* ptr);
CGSize C_NSTouch_DeviceSize(void* ptr);
