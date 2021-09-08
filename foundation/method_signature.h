#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_MethodSignature_Alloc();

void* C_NSMethodSignature_AllocMethodSignature();
void* C_NSMethodSignature_Init(void* ptr);
void* C_NSMethodSignature_NewMethodSignature();
void* C_NSMethodSignature_Autorelease(void* ptr);
void* C_NSMethodSignature_Retain(void* ptr);
bool C_NSMethodSignature_IsOneway(void* ptr);
unsigned int C_NSMethodSignature_NumberOfArguments(void* ptr);
unsigned int C_NSMethodSignature_FrameLength(void* ptr);
unsigned int C_NSMethodSignature_MethodReturnLength(void* ptr);
