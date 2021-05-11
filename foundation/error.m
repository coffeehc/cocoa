#import <Foundation/Foundation.h>
#import "error.h"

void* C_Error_Alloc() {
	return [NSError alloc];
}

void* C_NSError_Init(void* ptr) {
	NSError* nSError = (NSError*)ptr;
	NSError* result = [nSError init];
	return result;
}

int C_NSError_Code(void* ptr) {
	NSError* nSError = (NSError*)ptr;
	int result = [nSError code];
	return result;
}

void* C_NSError_Domain(void* ptr) {
	NSError* nSError = (NSError*)ptr;
	NSString* result = [nSError domain];
	return result;
}

void* C_NSError_LocalizedDescription(void* ptr) {
	NSError* nSError = (NSError*)ptr;
	NSString* result = [nSError localizedDescription];
	return result;
}

void* C_NSError_LocalizedRecoverySuggestion(void* ptr) {
	NSError* nSError = (NSError*)ptr;
	NSString* result = [nSError localizedRecoverySuggestion];
	return result;
}

void* C_NSError_LocalizedFailureReason(void* ptr) {
	NSError* nSError = (NSError*)ptr;
	NSString* result = [nSError localizedFailureReason];
	return result;
}

void* C_NSError_RecoveryAttempter(void* ptr) {
	NSError* nSError = (NSError*)ptr;
	id result = [nSError recoveryAttempter];
	return result;
}

void* C_NSError_HelpAnchor(void* ptr) {
	NSError* nSError = (NSError*)ptr;
	NSString* result = [nSError helpAnchor];
	return result;
}
