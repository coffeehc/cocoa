#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_BackForwardListItem_Alloc();

void* C_WKBackForwardListItem_AllocBackForwardListItem();
void* C_WKBackForwardListItem_Autorelease(void* ptr);
void* C_WKBackForwardListItem_Retain(void* ptr);
void* C_WKBackForwardListItem_Title(void* ptr);
void* C_WKBackForwardListItem_URL(void* ptr);
void* C_WKBackForwardListItem_InitialURL(void* ptr);
