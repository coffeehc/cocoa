#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_UserActivity_Alloc();

void* C_NSUserActivity_InitWithActivityType(void* ptr, void* activityType);
void* C_NSUserActivity_AllocUserActivity();
void* C_NSUserActivity_Autorelease(void* ptr);
void* C_NSUserActivity_Retain(void* ptr);
void C_NSUserActivity_BecomeCurrent(void* ptr);
void C_NSUserActivity_ResignCurrent(void* ptr);
void C_NSUserActivity_Invalidate(void* ptr);
void* C_NSUserActivity_ActivityType(void* ptr);
void* C_NSUserActivity_Title(void* ptr);
void C_NSUserActivity_SetTitle(void* ptr, void* value);
void* C_NSUserActivity_RequiredUserInfoKeys(void* ptr);
void C_NSUserActivity_SetRequiredUserInfoKeys(void* ptr, void* value);
void* C_NSUserActivity_TargetContentIdentifier(void* ptr);
void C_NSUserActivity_SetTargetContentIdentifier(void* ptr, void* value);
bool C_NSUserActivity_NeedsSave(void* ptr);
void C_NSUserActivity_SetNeedsSave(void* ptr, bool value);
void* C_NSUserActivity_Keywords(void* ptr);
void C_NSUserActivity_SetKeywords(void* ptr, void* value);
void* C_NSUserActivity_PersistentIdentifier(void* ptr);
void C_NSUserActivity_SetPersistentIdentifier(void* ptr, void* value);
bool C_NSUserActivity_IsEligibleForHandoff(void* ptr);
void C_NSUserActivity_SetEligibleForHandoff(void* ptr, bool value);
bool C_NSUserActivity_IsEligibleForSearch(void* ptr);
void C_NSUserActivity_SetEligibleForSearch(void* ptr, bool value);
bool C_NSUserActivity_IsEligibleForPublicIndexing(void* ptr);
void C_NSUserActivity_SetEligibleForPublicIndexing(void* ptr, bool value);
void* C_NSUserActivity_ExpirationDate(void* ptr);
void C_NSUserActivity_SetExpirationDate(void* ptr, void* value);
void* C_NSUserActivity_Delegate(void* ptr);
void C_NSUserActivity_SetDelegate(void* ptr, void* value);
bool C_NSUserActivity_SupportsContinuationStreams(void* ptr);
void C_NSUserActivity_SetSupportsContinuationStreams(void* ptr, bool value);
void* C_NSUserActivity_WebpageURL(void* ptr);
void C_NSUserActivity_SetWebpageURL(void* ptr, void* value);
void* C_NSUserActivity_ReferrerURL(void* ptr);
void C_NSUserActivity_SetReferrerURL(void* ptr, void* value);
