#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_AttributedString_Alloc();

void* C_NSAttributedString_AllocAttributedString();
void* C_NSAttributedString_Init(void* ptr);
void* C_NSAttributedString_NewAttributedString();
void* C_NSAttributedString_Autorelease(void* ptr);
void* C_NSAttributedString_Retain(void* ptr);
void* C_NSAttributedString_String(void* ptr);
unsigned int C_NSAttributedString_Length(void* ptr);

void* C_MutableAttributedString_Alloc();

void* C_NSMutableAttributedString_AllocMutableAttributedString();
void* C_NSMutableAttributedString_Init(void* ptr);
void* C_NSMutableAttributedString_NewMutableAttributedString();
void* C_NSMutableAttributedString_Autorelease(void* ptr);
void* C_NSMutableAttributedString_Retain(void* ptr);
void* C_NSMutableAttributedString_MutableString(void* ptr);
