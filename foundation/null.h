#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Null_Alloc();

void* C_NSNull_AllocNull();
void* C_NSNull_Init(void* ptr);
void* C_NSNull_NewNull();
void* C_NSNull_Autorelease(void* ptr);
void* C_NSNull_Retain(void* ptr);
void* C_NSNull_Null_();
