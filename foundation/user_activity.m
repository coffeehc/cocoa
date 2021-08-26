#import <Foundation/Foundation.h>
#import "user_activity.h"

void* C_UserActivity_Alloc() {
    return [NSUserActivity alloc];
}

void* C_NSUserActivity_InitWithActivityType(void* ptr, void* activityType) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    NSUserActivity* result_ = [nSUserActivity initWithActivityType:(NSString*)activityType];
    return result_;
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
    NSString* result_ = [nSUserActivity activityType];
    return result_;
}

void* C_NSUserActivity_Title(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    NSString* result_ = [nSUserActivity title];
    return result_;
}

void C_NSUserActivity_SetTitle(void* ptr, void* value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setTitle:(NSString*)value];
}

void* C_NSUserActivity_RequiredUserInfoKeys(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    NSSet* result_ = [nSUserActivity requiredUserInfoKeys];
    return result_;
}

void C_NSUserActivity_SetRequiredUserInfoKeys(void* ptr, void* value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setRequiredUserInfoKeys:(NSSet*)value];
}

void* C_NSUserActivity_TargetContentIdentifier(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    NSString* result_ = [nSUserActivity targetContentIdentifier];
    return result_;
}

void C_NSUserActivity_SetTargetContentIdentifier(void* ptr, void* value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setTargetContentIdentifier:(NSString*)value];
}

bool C_NSUserActivity_NeedsSave(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    BOOL result_ = [nSUserActivity needsSave];
    return result_;
}

void C_NSUserActivity_SetNeedsSave(void* ptr, bool value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setNeedsSave:value];
}

void* C_NSUserActivity_Keywords(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    NSSet* result_ = [nSUserActivity keywords];
    return result_;
}

void C_NSUserActivity_SetKeywords(void* ptr, void* value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setKeywords:(NSSet*)value];
}

void* C_NSUserActivity_PersistentIdentifier(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    NSUserActivityPersistentIdentifier result_ = [nSUserActivity persistentIdentifier];
    return result_;
}

void C_NSUserActivity_SetPersistentIdentifier(void* ptr, void* value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setPersistentIdentifier:(NSString*)value];
}

bool C_NSUserActivity_IsEligibleForHandoff(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    BOOL result_ = [nSUserActivity isEligibleForHandoff];
    return result_;
}

void C_NSUserActivity_SetEligibleForHandoff(void* ptr, bool value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setEligibleForHandoff:value];
}

bool C_NSUserActivity_IsEligibleForSearch(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    BOOL result_ = [nSUserActivity isEligibleForSearch];
    return result_;
}

void C_NSUserActivity_SetEligibleForSearch(void* ptr, bool value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setEligibleForSearch:value];
}

bool C_NSUserActivity_IsEligibleForPublicIndexing(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    BOOL result_ = [nSUserActivity isEligibleForPublicIndexing];
    return result_;
}

void C_NSUserActivity_SetEligibleForPublicIndexing(void* ptr, bool value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setEligibleForPublicIndexing:value];
}

void* C_NSUserActivity_ExpirationDate(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    NSDate* result_ = [nSUserActivity expirationDate];
    return result_;
}

void C_NSUserActivity_SetExpirationDate(void* ptr, void* value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setExpirationDate:(NSDate*)value];
}

void* C_NSUserActivity_Delegate(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    id result_ = [nSUserActivity delegate];
    return result_;
}

void C_NSUserActivity_SetDelegate(void* ptr, void* value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setDelegate:(id)value];
}

bool C_NSUserActivity_SupportsContinuationStreams(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    BOOL result_ = [nSUserActivity supportsContinuationStreams];
    return result_;
}

void C_NSUserActivity_SetSupportsContinuationStreams(void* ptr, bool value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setSupportsContinuationStreams:value];
}

void* C_NSUserActivity_WebpageURL(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    NSURL* result_ = [nSUserActivity webpageURL];
    return result_;
}

void C_NSUserActivity_SetWebpageURL(void* ptr, void* value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setWebpageURL:(NSURL*)value];
}

void* C_NSUserActivity_ReferrerURL(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    NSURL* result_ = [nSUserActivity referrerURL];
    return result_;
}

void C_NSUserActivity_SetReferrerURL(void* ptr, void* value) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    [nSUserActivity setReferrerURL:(NSURL*)value];
}

Array C_NSUserActivity_ContextIdentifierPath(void* ptr) {
    NSUserActivity* nSUserActivity = (NSUserActivity*)ptr;
    NSArray* result_ = [nSUserActivity contextIdentifierPath];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}
