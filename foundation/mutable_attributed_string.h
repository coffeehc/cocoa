#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_MutableAttributedString_Alloc();

void* C_NSMutableAttributedString_InitWithString(void* ptr, void* str);
void* C_NSMutableAttributedString_InitWithAttributedString(void* ptr, void* attrStr);
void* C_NSMutableAttributedString_Init(void* ptr);
void* C_NSMutableAttributedString_MutableString(void* ptr);
