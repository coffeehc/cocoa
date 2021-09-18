#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_Exception_Alloc();

void* C_NSException_AllocException();
void* C_NSException_Init(void* ptr);
void* C_NSException_NewException();
void* C_NSException_Autorelease(void* ptr);
void* C_NSException_Retain(void* ptr);
void C_NSException_Raise(void* ptr);
void* C_NSException_Name(void* ptr);
void* C_NSException_Reason(void* ptr);
Array C_NSException_CallStackReturnAddresses(void* ptr);
Array C_NSException_CallStackSymbols(void* ptr);
