#import "progress.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSProgress.h>

void* C_Progress_Alloc() {
    return [NSProgress alloc];
}

void* C_NSProgress_InitWithParent_UserInfo(void* ptr, void* parentProgressOrNil, Dictionary userInfoOrNil) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSMutableDictionary* objcUserInfoOrNil = [NSMutableDictionary dictionaryWithCapacity:userInfoOrNil.len];
    if (userInfoOrNil.len > 0) {
    	void** userInfoOrNilKeyData = (void**)userInfoOrNil.key_data;
    	void** userInfoOrNilValueData = (void**)userInfoOrNil.value_data;
    	for (int i = 0; i < userInfoOrNil.len; i++) {
    		void* kp = userInfoOrNilKeyData[i];
    		void* vp = userInfoOrNilValueData[i];
    		[objcUserInfoOrNil setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    NSProgress* result_ = [nSProgress initWithParent:(NSProgress*)parentProgressOrNil userInfo:objcUserInfoOrNil];
    return result_;
}

void* C_NSProgress_AllocProgress() {
    NSProgress* result_ = [NSProgress alloc];
    return result_;
}

void* C_NSProgress_Init(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSProgress* result_ = [nSProgress init];
    return result_;
}

void* C_NSProgress_NewProgress() {
    NSProgress* result_ = [NSProgress new];
    return result_;
}

void* C_NSProgress_Autorelease(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSProgress* result_ = [nSProgress autorelease];
    return result_;
}

void* C_NSProgress_Retain(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSProgress* result_ = [nSProgress retain];
    return result_;
}

void* C_NSProgress_DiscreteProgressWithTotalUnitCount(long unitCount) {
    NSProgress* result_ = [NSProgress discreteProgressWithTotalUnitCount:unitCount];
    return result_;
}

void* C_NSProgress_ProgressWithTotalUnitCount(long unitCount) {
    NSProgress* result_ = [NSProgress progressWithTotalUnitCount:unitCount];
    return result_;
}

void* C_NSProgress_ProgressWithTotalUnitCount_Parent_PendingUnitCount(long unitCount, void* parent, long portionOfParentTotalUnitCount) {
    NSProgress* result_ = [NSProgress progressWithTotalUnitCount:unitCount parent:(NSProgress*)parent pendingUnitCount:portionOfParentTotalUnitCount];
    return result_;
}

void* C_NSProgress_CurrentProgress() {
    NSProgress* result_ = [NSProgress currentProgress];
    return result_;
}

void C_NSProgress_BecomeCurrentWithPendingUnitCount(void* ptr, long unitCount) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress becomeCurrentWithPendingUnitCount:unitCount];
}

void C_NSProgress_AddChild_WithPendingUnitCount(void* ptr, void* child, long inUnitCount) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress addChild:(NSProgress*)child withPendingUnitCount:inUnitCount];
}

void C_NSProgress_ResignCurrent(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress resignCurrent];
}

void C_NSProgress_Cancel(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress cancel];
}

void C_NSProgress_Pause(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress pause];
}

void C_NSProgress_Resume(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress resume];
}

void C_NSProgress_SetUserInfoObject_ForKey(void* ptr, void* objectOrNil, void* key) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setUserInfoObject:(id)objectOrNil forKey:(NSString*)key];
}

void C_NSProgress_Publish(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress publish];
}

void C_NSProgress_Unpublish(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress unpublish];
}

void C_NSProgress_Progress_RemoveSubscriber(void* subscriber) {
    [NSProgress removeSubscriber:(id)subscriber];
}

long C_NSProgress_TotalUnitCount(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSInteger result_ = [nSProgress totalUnitCount];
    return result_;
}

void C_NSProgress_SetTotalUnitCount(void* ptr, long value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setTotalUnitCount:value];
}

long C_NSProgress_CompletedUnitCount(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSInteger result_ = [nSProgress completedUnitCount];
    return result_;
}

void C_NSProgress_SetCompletedUnitCount(void* ptr, long value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setCompletedUnitCount:value];
}

void* C_NSProgress_LocalizedDescription(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSString* result_ = [nSProgress localizedDescription];
    return result_;
}

void C_NSProgress_SetLocalizedDescription(void* ptr, void* value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setLocalizedDescription:(NSString*)value];
}

void* C_NSProgress_LocalizedAdditionalDescription(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSString* result_ = [nSProgress localizedAdditionalDescription];
    return result_;
}

void C_NSProgress_SetLocalizedAdditionalDescription(void* ptr, void* value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setLocalizedAdditionalDescription:(NSString*)value];
}

double C_NSProgress_FractionCompleted(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    double result_ = [nSProgress fractionCompleted];
    return result_;
}

bool C_NSProgress_IsCancellable(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    BOOL result_ = [nSProgress isCancellable];
    return result_;
}

void C_NSProgress_SetCancellable(void* ptr, bool value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setCancellable:value];
}

bool C_NSProgress_IsCancelled(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    BOOL result_ = [nSProgress isCancelled];
    return result_;
}

bool C_NSProgress_IsPausable(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    BOOL result_ = [nSProgress isPausable];
    return result_;
}

void C_NSProgress_SetPausable(void* ptr, bool value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setPausable:value];
}

bool C_NSProgress_IsPaused(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    BOOL result_ = [nSProgress isPaused];
    return result_;
}

bool C_NSProgress_IsIndeterminate(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    BOOL result_ = [nSProgress isIndeterminate];
    return result_;
}

void* C_NSProgress_Kind(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSProgressKind result_ = [nSProgress kind];
    return result_;
}

void C_NSProgress_SetKind(void* ptr, void* value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setKind:(NSString*)value];
}

Dictionary C_NSProgress_UserInfo(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSDictionary* result_ = [nSProgress userInfo];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSProgressUserInfoKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void* C_NSProgress_FileOperationKind(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSProgressFileOperationKind result_ = [nSProgress fileOperationKind];
    return result_;
}

void C_NSProgress_SetFileOperationKind(void* ptr, void* value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setFileOperationKind:(NSString*)value];
}

void* C_NSProgress_FileURL(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSURL* result_ = [nSProgress fileURL];
    return result_;
}

void C_NSProgress_SetFileURL(void* ptr, void* value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setFileURL:(NSURL*)value];
}

bool C_NSProgress_IsFinished(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    BOOL result_ = [nSProgress isFinished];
    return result_;
}

bool C_NSProgress_IsOld(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    BOOL result_ = [nSProgress isOld];
    return result_;
}

void* C_NSProgress_EstimatedTimeRemaining(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSNumber* result_ = [nSProgress estimatedTimeRemaining];
    return result_;
}

void C_NSProgress_SetEstimatedTimeRemaining(void* ptr, void* value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setEstimatedTimeRemaining:(NSNumber*)value];
}

void* C_NSProgress_FileCompletedCount(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSNumber* result_ = [nSProgress fileCompletedCount];
    return result_;
}

void C_NSProgress_SetFileCompletedCount(void* ptr, void* value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setFileCompletedCount:(NSNumber*)value];
}

void* C_NSProgress_FileTotalCount(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSNumber* result_ = [nSProgress fileTotalCount];
    return result_;
}

void C_NSProgress_SetFileTotalCount(void* ptr, void* value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setFileTotalCount:(NSNumber*)value];
}

void* C_NSProgress_Throughput(void* ptr) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    NSNumber* result_ = [nSProgress throughput];
    return result_;
}

void C_NSProgress_SetThroughput(void* ptr, void* value) {
    NSProgress* nSProgress = (NSProgress*)ptr;
    [nSProgress setThroughput:(NSNumber*)value];
}
