#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Unit_Alloc();

void* C_NSUnit_InitWithSymbol(void* ptr, void* symbol);
void* C_NSUnit_AllocUnit();
void* C_NSUnit_Autorelease(void* ptr);
void* C_NSUnit_Retain(void* ptr);
void* C_NSUnit_Symbol(void* ptr);
