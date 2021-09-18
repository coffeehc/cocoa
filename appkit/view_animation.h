#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_ViewAnimation_Alloc();

void* C_NSViewAnimation_InitWithDuration_AnimationCurve(void* ptr, double duration, unsigned int animationCurve);
void* C_NSViewAnimation_InitWithCoder(void* ptr, void* coder);
void* C_NSViewAnimation_AllocViewAnimation();
void* C_NSViewAnimation_Init(void* ptr);
void* C_NSViewAnimation_NewViewAnimation();
void* C_NSViewAnimation_Autorelease(void* ptr);
void* C_NSViewAnimation_Retain(void* ptr);
