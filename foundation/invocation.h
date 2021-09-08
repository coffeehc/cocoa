#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Invocation_Alloc();

void* C_NSInvocation_AllocInvocation();
void* C_NSInvocation_Init(void* ptr);
void* C_NSInvocation_NewInvocation();
void* C_NSInvocation_Autorelease(void* ptr);
void* C_NSInvocation_Retain(void* ptr);
void* C_NSInvocation_InvocationWithMethodSignature(void* sig);
void C_NSInvocation_RetainArguments(void* ptr);
void C_NSInvocation_Invoke(void* ptr);
void C_NSInvocation_InvokeWithTarget(void* ptr, void* target);
void* C_NSInvocation_Selector(void* ptr);
void C_NSInvocation_SetSelector(void* ptr, void* value);
void* C_NSInvocation_Target(void* ptr);
void C_NSInvocation_SetTarget(void* ptr, void* value);
bool C_NSInvocation_ArgumentsRetained(void* ptr);
void* C_NSInvocation_MethodSignature(void* ptr);
