#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Unit_Alloc();

void* C_NSUnit_InitWithSymbol(void* ptr, void* symbol);
void* C_NSUnit_Symbol(void* ptr);
