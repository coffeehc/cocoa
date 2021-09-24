#import "website_data_store.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <WebKit/WKWebsiteDataStore.h>

void* C_WebsiteDataStore_Alloc() {
    return [WKWebsiteDataStore alloc];
}

void* C_WKWebsiteDataStore_AllocWebsiteDataStore() {
    WKWebsiteDataStore* result_ = [WKWebsiteDataStore alloc];
    return result_;
}

void* C_WKWebsiteDataStore_Autorelease(void* ptr) {
    WKWebsiteDataStore* wKWebsiteDataStore = (WKWebsiteDataStore*)ptr;
    WKWebsiteDataStore* result_ = [wKWebsiteDataStore autorelease];
    return result_;
}

void* C_WKWebsiteDataStore_Retain(void* ptr) {
    WKWebsiteDataStore* wKWebsiteDataStore = (WKWebsiteDataStore*)ptr;
    WKWebsiteDataStore* result_ = [wKWebsiteDataStore retain];
    return result_;
}

void* C_WKWebsiteDataStore_WebsiteDataStore_DefaultDataStore() {
    WKWebsiteDataStore* result_ = [WKWebsiteDataStore defaultDataStore];
    return result_;
}

void* C_WKWebsiteDataStore_WebsiteDataStore_NonPersistentDataStore() {
    WKWebsiteDataStore* result_ = [WKWebsiteDataStore nonPersistentDataStore];
    return result_;
}

void* C_WKWebsiteDataStore_WebsiteDataStore_AllWebsiteDataTypes() {
    NSSet* result_ = [WKWebsiteDataStore allWebsiteDataTypes];
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
