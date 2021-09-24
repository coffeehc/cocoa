#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_MutableAttributedString_Alloc();

void* C_NSMutableAttributedString_AllocMutableAttributedString();
void* C_NSMutableAttributedString_Init(void* ptr);
void* C_NSMutableAttributedString_NewMutableAttributedString();
void* C_NSMutableAttributedString_Autorelease(void* ptr);
void* C_NSMutableAttributedString_Retain(void* ptr);
void* C_NSMutableAttributedString_MutableString(void* ptr);
