#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_MutableAttributedString_Alloc();

void* C_NSMutableAttributedString_InitWithString(void* ptr, void* str);
void* C_NSMutableAttributedString_InitWithString_Attributes(void* ptr, void* str, Dictionary attrs);
void* C_NSMutableAttributedString_InitWithAttributedString(void* ptr, void* attrStr);
void* C_NSMutableAttributedString_AllocMutableAttributedString();
void* C_NSMutableAttributedString_Init(void* ptr);
void* C_NSMutableAttributedString_NewMutableAttributedString();
void* C_NSMutableAttributedString_Autorelease(void* ptr);
void* C_NSMutableAttributedString_Retain(void* ptr);
void* C_NSMutableAttributedString_MutableString(void* ptr);
