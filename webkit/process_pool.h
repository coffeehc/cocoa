#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_ProcessPool_Alloc();

void* C_WKProcessPool_AllocProcessPool();
void* C_WKProcessPool_Init(void* ptr);
void* C_WKProcessPool_NewProcessPool();
void* C_WKProcessPool_Autorelease(void* ptr);
void* C_WKProcessPool_Retain(void* ptr);
