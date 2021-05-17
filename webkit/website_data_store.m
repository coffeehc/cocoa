#import <WebKit/WebKit.h>
#import "website_data_store.h"

void* C_WebsiteDataStore_Alloc() {
    return [WKWebsiteDataStore alloc];
}

void* C_WKWebsiteDataStore_WebsiteDataStore_DefaultDataStore() {
    WKWebsiteDataStore* result_ = [WKWebsiteDataStore defaultDataStore];
    return result_;
}

void* C_WKWebsiteDataStore_WebsiteDataStore_NonPersistentDataStore() {
    WKWebsiteDataStore* result_ = [WKWebsiteDataStore nonPersistentDataStore];
    return result_;
}

bool C_WKWebsiteDataStore_IsPersistent(void* ptr) {
    WKWebsiteDataStore* wKWebsiteDataStore = (WKWebsiteDataStore*)ptr;
    BOOL result_ = [wKWebsiteDataStore isPersistent];
    return result_;
}

void* C_WKWebsiteDataStore_HttpCookieStore(void* ptr) {
    WKWebsiteDataStore* wKWebsiteDataStore = (WKWebsiteDataStore*)ptr;
    WKHTTPCookieStore* result_ = [wKWebsiteDataStore httpCookieStore];
    return result_;
}
