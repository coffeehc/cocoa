#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_AttributedString_Alloc();

void* C_NSAttributedString_Init(void* ptr);
void* C_NSAttributedString_String(void* ptr);
unsigned int C_NSAttributedString_Length(void* ptr);
Array C_NSAttributedString_AttributedString_TextTypes();
Array C_NSAttributedString_AttributedString_TextUnfilteredTypes();
