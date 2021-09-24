#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_WebsiteDataStore_Alloc();

void* C_WKWebsiteDataStore_AllocWebsiteDataStore();
void* C_WKWebsiteDataStore_Autorelease(void* ptr);
void* C_WKWebsiteDataStore_Retain(void* ptr);
void* C_WKWebsiteDataStore_WebsiteDataStore_DefaultDataStore();
void* C_WKWebsiteDataStore_WebsiteDataStore_NonPersistentDataStore();
void* C_WKWebsiteDataStore_WebsiteDataStore_AllWebsiteDataTypes();
bool C_WKWebsiteDataStore_IsPersistent(void* ptr);
void* C_WKWebsiteDataStore_HttpCookieStore(void* ptr);
