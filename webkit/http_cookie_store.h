#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_HTTPCookieStore_Alloc();

void C_WKHTTPCookieStore_AddObserver(void* ptr, void* observer);
void C_WKHTTPCookieStore_RemoveObserver(void* ptr, void* observer);
