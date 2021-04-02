#import <Foundation/Foundation.h>
#import "url_request.h"

unsigned long URLRequest_CachePolicy(void* ptr) {
	NSURLRequest* uRLRequest = (NSURLRequest*)ptr;
	return [uRLRequest cachePolicy];
}

void* URLRequest_NewURLRequest(void* url) {
	NSURLRequest* uRLRequest = [NSURLRequest alloc];
	return [[uRLRequest initWithURL:(NSURL*)url] autorelease];
}

void* URLRequest_RequestWithURL(void* url) {
	return [NSURLRequest requestWithURL:(NSURL*)url];
}
