#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_SecurityOrigin_Alloc();

void* C_WKSecurityOrigin_AllocSecurityOrigin();
void* C_WKSecurityOrigin_Autorelease(void* ptr);
void* C_WKSecurityOrigin_Retain(void* ptr);
void* C_WKSecurityOrigin_Host(void* ptr);
int C_WKSecurityOrigin_Port(void* ptr);
void* C_WKSecurityOrigin_Protocol(void* ptr);
