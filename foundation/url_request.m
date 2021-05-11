#import <Foundation/Foundation.h>
#import "url_request.h"

void* C_URLRequest_Alloc() {
	return [NSURLRequest alloc];
}

void* C_NSURLRequest_InitWithURL(void* ptr, void* URL) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	NSURLRequest* result = [nSURLRequest initWithURL:(NSURL*)URL];
	return result;
}

void* C_NSURLRequest_Init(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	NSURLRequest* result = [nSURLRequest init];
	return result;
}

void* C_NSURLRequest_URLRequestRequestWithURL(void* URL) {
	NSURLRequest* result = [NSURLRequest requestWithURL:(NSURL*)URL];
	return result;
}

void* C_NSURLRequest_ValueForHTTPHeaderField(void* ptr, void* field) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	NSString* result = [nSURLRequest valueForHTTPHeaderField:(NSString*)field];
	return result;
}

void* C_NSURLRequest_HTTPMethod(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	NSString* result = [nSURLRequest HTTPMethod];
	return result;
}

void* C_NSURLRequest_URL(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	NSURL* result = [nSURLRequest URL];
	return result;
}

Array C_NSURLRequest_HTTPBody(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	NSData* result = [nSURLRequest HTTPBody];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

void* C_NSURLRequest_MainDocumentURL(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	NSURL* result = [nSURLRequest mainDocumentURL];
	return result;
}

double C_NSURLRequest_TimeoutInterval(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	double result = [nSURLRequest timeoutInterval];
	return result;
}

bool C_NSURLRequest_HTTPShouldHandleCookies(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	bool result = [nSURLRequest HTTPShouldHandleCookies];
	return result;
}

bool C_NSURLRequest_HTTPShouldUsePipelining(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	bool result = [nSURLRequest HTTPShouldUsePipelining];
	return result;
}

bool C_NSURLRequest_AllowsCellularAccess(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	bool result = [nSURLRequest allowsCellularAccess];
	return result;
}

bool C_NSURLRequest_AllowsConstrainedNetworkAccess(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	bool result = [nSURLRequest allowsConstrainedNetworkAccess];
	return result;
}

bool C_NSURLRequest_AllowsExpensiveNetworkAccess(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	bool result = [nSURLRequest allowsExpensiveNetworkAccess];
	return result;
}

bool C_NSURLRequest_AssumesHTTP3Capable(void* ptr) {
	NSURLRequest* nSURLRequest = (NSURLRequest*)ptr;
	bool result = [nSURLRequest assumesHTTP3Capable];
	return result;
}
