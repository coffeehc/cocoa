#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Exception_Alloc();

void* C_NSException_Init(void* ptr);
void C_NSException_Raise(void* ptr);
void* C_NSException_Name(void* ptr);
void* C_NSException_Reason(void* ptr);
Array C_NSException_CallStackReturnAddresses(void* ptr);
Array C_NSException_CallStackSymbols(void* ptr);
