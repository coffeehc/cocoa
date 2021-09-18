#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_ProcessPool_Alloc();

void* C_WKProcessPool_AllocProcessPool();
void* C_WKProcessPool_Init(void* ptr);
void* C_WKProcessPool_NewProcessPool();
void* C_WKProcessPool_Autorelease(void* ptr);
void* C_WKProcessPool_Retain(void* ptr);
