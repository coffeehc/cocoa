#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_HTTPCookieStore_Alloc();

void* C_WKHTTPCookieStore_AllocHTTPCookieStore();
void* C_WKHTTPCookieStore_Autorelease(void* ptr);
void* C_WKHTTPCookieStore_Retain(void* ptr);
void C_WKHTTPCookieStore_AddObserver(void* ptr, void* observer);
void C_WKHTTPCookieStore_RemoveObserver(void* ptr, void* observer);
