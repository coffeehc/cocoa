#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_ExtensionContext_Alloc();

void* C_NSExtensionContext_AllocExtensionContext();
void* C_NSExtensionContext_Init(void* ptr);
void* C_NSExtensionContext_NewExtensionContext();
void* C_NSExtensionContext_Autorelease(void* ptr);
void* C_NSExtensionContext_Retain(void* ptr);
void C_NSExtensionContext_CancelRequestWithError(void* ptr, void* error);
Array C_NSExtensionContext_InputItems(void* ptr);
