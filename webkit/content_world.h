#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_ContentWorld_Alloc();

void* C_WKContentWorld_AllocContentWorld();
void* C_WKContentWorld_Autorelease(void* ptr);
void* C_WKContentWorld_Retain(void* ptr);
void* C_WKContentWorld_ContentWorld_WorldWithName(void* name);
void* C_WKContentWorld_ContentWorld_DefaultClientWorld();
void* C_WKContentWorld_ContentWorld_PageWorld();
void* C_WKContentWorld_Name(void* ptr);
