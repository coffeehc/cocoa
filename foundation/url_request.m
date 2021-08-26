#import <Foundation/Foundation.h>
#import "url_request.h"

void* C_URLRequest_Alloc() {
    return [NSURLRequest alloc];
}

void* C_NSURLRequest_InitWithURL(void* ptr, void* URL) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSURLRequest* result_ = [nSURLRequest initWithURL:(NSURL*)URL];
    return result_;
}

void* C_NSURLRequest_InitWithURL_CachePolicy_TimeoutInterval(void* ptr, void* URL, unsigned int cachePolicy, double timeoutInterval) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSURLRequest* result_ = [nSURLRequest initWithURL:(NSURL*)URL cachePolicy:cachePolicy timeoutInterval:timeoutInterval];
    return result_;
}

void* C_NSURLRequest_Init(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSURLRequest* result_ = [nSURLRequest init];
    return result_;
}

void* C_NSURLRequest_URLRequest_RequestWithURL(void* URL) {
    NSURLRequest* result_ = [NSURLRequest requestWithURL:(NSURL*)URL];
    return result_;
}

void* C_NSURLRequest_URLRequest_RequestWithURL_CachePolicy_TimeoutInterval(void* URL, unsigned int cachePolicy, double timeoutInterval) {
    NSURLRequest* result_ = [NSURLRequest requestWithURL:(NSURL*)URL cachePolicy:cachePolicy timeoutInterval:timeoutInterval];
    return result_;
}

void* C_NSURLRequest_ValueForHTTPHeaderField(void* ptr, void* field) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSString* result_ = [nSURLRequest valueForHTTPHeaderField:(NSString*)field];
    return result_;
}

unsigned int C_NSURLRequest_CachePolicy(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSURLRequestCachePolicy result_ = [nSURLRequest cachePolicy];
    return result_;
}

void* C_NSURLRequest_HTTPMethod(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSString* result_ = [nSURLRequest HTTPMethod];
    return result_;
}

void* C_NSURLRequest_URL(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSURL* result_ = [nSURLRequest URL];
    return result_;
}

Array C_NSURLRequest_HTTPBody(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSData* result_ = [nSURLRequest HTTPBody];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

void* C_NSURLRequest_MainDocumentURL(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSURL* result_ = [nSURLRequest mainDocumentURL];
    return result_;
}

Dictionary C_NSURLRequest_AllHTTPHeaderFields(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSDictionary* result_ = [nSURLRequest allHTTPHeaderFields];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSString* kp = [result_Keys objectAtIndex:i];
    		NSString* vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

double C_NSURLRequest_TimeoutInterval(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSTimeInterval result_ = [nSURLRequest timeoutInterval];
    return result_;
}

bool C_NSURLRequest_HTTPShouldHandleCookies(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    BOOL result_ = [nSURLRequest HTTPShouldHandleCookies];
    return result_;
}

bool C_NSURLRequest_HTTPShouldUsePipelining(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    BOOL result_ = [nSURLRequest HTTPShouldUsePipelining];
    return result_;
}

bool C_NSURLRequest_AllowsCellularAccess(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    BOOL result_ = [nSURLRequest allowsCellularAccess];
    return result_;
}

bool C_NSURLRequest_AllowsConstrainedNetworkAccess(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    BOOL result_ = [nSURLRequest allowsConstrainedNetworkAccess];
    return result_;
}

bool C_NSURLRequest_AllowsExpensiveNetworkAccess(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    BOOL result_ = [nSURLRequest allowsExpensiveNetworkAccess];
    return result_;
}

unsigned int C_NSURLRequest_NetworkServiceType(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    NSURLRequestNetworkServiceType result_ = [nSURLRequest networkServiceType];
    return result_;
}

bool C_NSURLRequest_URLRequest_SupportsSecureCoding() {
    BOOL result_ = [NSURLRequest supportsSecureCoding];
    return result_;
}

bool C_NSURLRequest_AssumesHTTP3Capable(void* ptr) {
    NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
    BOOL result_ = [nSURLRequest assumesHTTP3Capable];
    return result_;
}
