#import <WebKit/WebKit.h>
#import "http_cookie_store.h"

void* C_HTTPCookieStore_Alloc() {
    return [WKHTTPCookieStore alloc];
}

void C_WKHTTPCookieStore_AddObserver(void* ptr, void* observer) {
    WKHTTPCookieStore* wKHTTPCookieStore = (WKHTTPCookieStore*)ptr;
    [wKHTTPCookieStore addObserver:(id)observer];
}

void C_WKHTTPCookieStore_RemoveObserver(void* ptr, void* observer) {
    WKHTTPCookieStore* wKHTTPCookieStore = (WKHTTPCookieStore*)ptr;
    [wKHTTPCookieStore removeObserver:(id)observer];
}
