#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Formatter_Alloc();

void* C_NSFormatter_AllocFormatter();
void* C_NSFormatter_Init(void* ptr);
void* C_NSFormatter_NewFormatter();
void* C_NSFormatter_Autorelease(void* ptr);
void* C_NSFormatter_Retain(void* ptr);
void* C_NSFormatter_StringForObjectValue(void* ptr, void* obj);
void* C_NSFormatter_AttributedStringForObjectValue_WithDefaultAttributes(void* ptr, void* obj, Dictionary attrs);
void* C_NSFormatter_EditingStringForObjectValue(void* ptr, void* obj);
