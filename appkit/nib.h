#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Nib_Alloc();

void* C_NSNib_InitWithNibNamed_Bundle(void* ptr, void* nibName, void* bundle);
void* C_NSNib_InitWithNibData_Bundle(void* ptr, void* nibData, void* bundle);
void* C_NSNib_AllocNib();
void* C_NSNib_Init(void* ptr);
void* C_NSNib_NewNib();
void* C_NSNib_Autorelease(void* ptr);
void* C_NSNib_Retain(void* ptr);
