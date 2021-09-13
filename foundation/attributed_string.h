#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_AttributedString_Alloc();

void* C_NSAttributedString_AllocAttributedString();
void* C_NSAttributedString_Init(void* ptr);
void* C_NSAttributedString_NewAttributedString();
void* C_NSAttributedString_Autorelease(void* ptr);
void* C_NSAttributedString_Retain(void* ptr);
void* C_NSAttributedString_String(void* ptr);
unsigned int C_NSAttributedString_Length(void* ptr);
