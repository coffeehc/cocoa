#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ViewAnimation_Alloc();

void* C_NSViewAnimation_InitWithDuration_AnimationCurve(void* ptr, double duration, unsigned int animationCurve);
void* C_NSViewAnimation_InitWithCoder(void* ptr, void* coder);
void* C_NSViewAnimation_Init(void* ptr);
