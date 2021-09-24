#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Shadow_Alloc();

void* C_NSShadow_Init(void* ptr);
void* C_NSShadow_AllocShadow();
void* C_NSShadow_NewShadow();
void* C_NSShadow_Autorelease(void* ptr);
void* C_NSShadow_Retain(void* ptr);
void C_NSShadow_Set(void* ptr);
CGSize C_NSShadow_ShadowOffset(void* ptr);
void C_NSShadow_SetShadowOffset(void* ptr, CGSize value);
double C_NSShadow_ShadowBlurRadius(void* ptr);
void C_NSShadow_SetShadowBlurRadius(void* ptr, double value);
void* C_NSShadow_ShadowColor(void* ptr);
void C_NSShadow_SetShadowColor(void* ptr, void* value);
