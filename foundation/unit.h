#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Unit_Alloc();

void* C_NSUnit_InitWithSymbol(void* ptr, void* symbol);
void* C_NSUnit_AllocUnit();
void* C_NSUnit_Autorelease(void* ptr);
void* C_NSUnit_Retain(void* ptr);
void* C_NSUnit_Symbol(void* ptr);
