#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_MethodSignature_Alloc();

void* C_NSMethodSignature_Init(void* ptr);
bool C_NSMethodSignature_IsOneway(void* ptr);
unsigned int C_NSMethodSignature_NumberOfArguments(void* ptr);
unsigned int C_NSMethodSignature_FrameLength(void* ptr);
unsigned int C_NSMethodSignature_MethodReturnLength(void* ptr);
