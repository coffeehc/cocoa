#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Formatter_Alloc();

void* C_NSFormatter_Init(void* ptr);
void* C_NSFormatter_StringForObjectValue(void* ptr, void* obj);
void* C_NSFormatter_AttributedStringForObjectValue_WithDefaultAttributes(void* ptr, void* obj, Dictionary attrs);
void* C_NSFormatter_EditingStringForObjectValue(void* ptr, void* obj);
