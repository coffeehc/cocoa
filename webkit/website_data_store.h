#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <WebKit/WebKit.h>

void* C_WebsiteDataStore_Alloc();

void* C_WKWebsiteDataStore_WebsiteDataStore_DefaultDataStore();
void* C_WKWebsiteDataStore_WebsiteDataStore_NonPersistentDataStore();
bool C_WKWebsiteDataStore_IsPersistent(void* ptr);
void* C_WKWebsiteDataStore_HttpCookieStore(void* ptr);
