#import "http_cookie_store.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <WebKit/WKHTTPCookieStore.h>

void* C_HTTPCookieStore_Alloc() {
    return [WKHTTPCookieStore alloc];
}

void* C_WKHTTPCookieStore_AllocHTTPCookieStore() {
    WKHTTPCookieStore* result_ = [WKHTTPCookieStore alloc];
    return result_;
}

void* C_WKHTTPCookieStore_Autorelease(void* ptr) {
    WKHTTPCookieStore* wKHTTPCookieStore = (WKHTTPCookieStore*)ptr;
    WKHTTPCookieStore* result_ = [wKHTTPCookieStore autorelease];
    return result_;
}

void* C_WKHTTPCookieStore_Retain(void* ptr) {
    WKHTTPCookieStore* wKHTTPCookieStore = (WKHTTPCookieStore*)ptr;
    WKHTTPCookieStore* result_ = [wKHTTPCookieStore retain];
    return result_;
}

void C_WKHTTPCookieStore_AddObserver(void* ptr, void* observer) {
    WKHTTPCookieStore* wKHTTPCookieStore = (WKHTTPCookieStore*)ptr;
    [wKHTTPCookieStore addObserver:(id)observer];
}

void C_WKHTTPCookieStore_RemoveObserver(void* ptr, void* observer) {
    WKHTTPCookieStore* wKHTTPCookieStore = (WKHTTPCookieStore*)ptr;
    [wKHTTPCookieStore removeObserver:(id)observer];
}
