#import <Foundation/Foundation.h>
#import "user_activity.h"

void* C_UserActivity_Alloc() {
	return [NSUserActivity alloc];
}

void* C_NSUserActivity_InitWithActivityType(void* ptr, void* activityType) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	NSUserActivity* result = [nSUserActivity initWithActivityType:(NSString*)activityType];
	return result;
}

void C_NSUserActivity_BecomeCurrent(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity becomeCurrent];
}

void C_NSUserActivity_ResignCurrent(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity resignCurrent];
}

void C_NSUserActivity_Invalidate(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity invalidate];
}

void* C_NSUserActivity_ActivityType(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	NSString* result = [nSUserActivity activityType];
	return result;
}

void* C_NSUserActivity_Title(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	NSString* result = [nSUserActivity title];
	return result;
}

void C_NSUserActivity_SetTitle(void* ptr, void* value) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity setTitle:(NSString*)value];
}

void* C_NSUserActivity_TargetContentIdentifier(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	NSString* result = [nSUserActivity targetContentIdentifier];
	return result;
}

void C_NSUserActivity_SetTargetContentIdentifier(void* ptr, void* value) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity setTargetContentIdentifier:(NSString*)value];
}

bool C_NSUserActivity_NeedsSave(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	bool result = [nSUserActivity needsSave];
	return result;
}

void C_NSUserActivity_SetNeedsSave(void* ptr, bool value) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity setNeedsSave:value];
}

void* C_NSUserActivity_PersistentIdentifier(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	NSString* result = [nSUserActivity persistentIdentifier];
	return result;
}

void C_NSUserActivity_SetPersistentIdentifier(void* ptr, void* value) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity setPersistentIdentifier:(NSString*)value];
}

bool C_NSUserActivity_IsEligibleForHandoff(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	bool result = [nSUserActivity isEligibleForHandoff];
	return result;
}

void C_NSUserActivity_SetEligibleForHandoff(void* ptr, bool value) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity setEligibleForHandoff:value];
}

bool C_NSUserActivity_IsEligibleForSearch(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	bool result = [nSUserActivity isEligibleForSearch];
	return result;
}

void C_NSUserActivity_SetEligibleForSearch(void* ptr, bool value) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity setEligibleForSearch:value];
}

bool C_NSUserActivity_IsEligibleForPublicIndexing(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	bool result = [nSUserActivity isEligibleForPublicIndexing];
	return result;
}

void C_NSUserActivity_SetEligibleForPublicIndexing(void* ptr, bool value) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity setEligibleForPublicIndexing:value];
}

void* C_NSUserActivity_ExpirationDate(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	NSDate* result = [nSUserActivity expirationDate];
	return result;
}

void C_NSUserActivity_SetExpirationDate(void* ptr, void* value) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity setExpirationDate:(NSDate*)value];
}

bool C_NSUserActivity_SupportsContinuationStreams(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	bool result = [nSUserActivity supportsContinuationStreams];
	return result;
}

void C_NSUserActivity_SetSupportsContinuationStreams(void* ptr, bool value) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity setSupportsContinuationStreams:value];
}

void* C_NSUserActivity_WebpageURL(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	NSURL* result = [nSUserActivity webpageURL];
	return result;
}

void C_NSUserActivity_SetWebpageURL(void* ptr, void* value) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity setWebpageURL:(NSURL*)value];
}

void* C_NSUserActivity_ReferrerURL(void* ptr) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	NSURL* result = [nSUserActivity referrerURL];
	return result;
}

void C_NSUserActivity_SetReferrerURL(void* ptr, void* value) {
	NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
	[nSUserActivity setReferrerURL:(NSURL*)value];
}
