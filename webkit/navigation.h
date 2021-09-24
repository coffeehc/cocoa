#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Navigation_Alloc();

void* C_WKNavigation_AllocNavigation();
void* C_WKNavigation_Init(void* ptr);
void* C_WKNavigation_NewNavigation();
void* C_WKNavigation_Autorelease(void* ptr);
void* C_WKNavigation_Retain(void* ptr);
int C_WKNavigation_EffectiveContentMode(void* ptr);
