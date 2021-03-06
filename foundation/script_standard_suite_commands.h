#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_CloseCommand_Alloc();

void* C_NSCloseCommand_InitWithCommandDescription(void* ptr, void* commandDef);
void* C_NSCloseCommand_InitWithCoder(void* ptr, void* inCoder);
void* C_NSCloseCommand_AllocCloseCommand();
void* C_NSCloseCommand_Autorelease(void* ptr);
void* C_NSCloseCommand_Retain(void* ptr);
unsigned int C_NSCloseCommand_SaveOptions(void* ptr);
