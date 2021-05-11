#import <Foundation/Foundation.h>
#import "progress.h"

void* C_Progress_Alloc() {
	return [NSProgress alloc];
}

void* C_NSProgress_Init(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	NSProgress* result = [nSProgress init];
	return result;
}

void* C_NSProgress_ProgressDiscreteProgressWithTotalUnitCount(long unitCount) {
	NSProgress* result = [NSProgress discreteProgressWithTotalUnitCount:unitCount];
	return result;
}

void* C_NSProgress_ProgressWithTotalUnitCount(long unitCount) {
	NSProgress* result = [NSProgress progressWithTotalUnitCount:unitCount];
	return result;
}

void* C_NSProgress_ProgressWithTotalUnitCount_Parent_PendingUnitCount(long unitCount, void* parent, long portionOfParentTotalUnitCount) {
	NSProgress* result = [NSProgress progressWithTotalUnitCount:unitCount parent:(NSProgress*)parent pendingUnitCount:portionOfParentTotalUnitCount];
	return result;
}

void* C_NSProgress_ProgressCurrentProgress() {
	NSProgress* result = [NSProgress currentProgress];
	return result;
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

void C_NSProgress_ProgressRemoveSubscriber(void* subscriber) {
	[NSProgress removeSubscriber:(id)subscriber];
}

long C_NSProgress_TotalUnitCount(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	long result = [nSProgress totalUnitCount];
	return result;
}

void C_NSProgress_SetTotalUnitCount(void* ptr, long value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setTotalUnitCount:value];
}

long C_NSProgress_CompletedUnitCount(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	long result = [nSProgress completedUnitCount];
	return result;
}

void C_NSProgress_SetCompletedUnitCount(void* ptr, long value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setCompletedUnitCount:value];
}

void* C_NSProgress_LocalizedDescription(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	NSString* result = [nSProgress localizedDescription];
	return result;
}

void C_NSProgress_SetLocalizedDescription(void* ptr, void* value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setLocalizedDescription:(NSString*)value];
}

void* C_NSProgress_LocalizedAdditionalDescription(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	NSString* result = [nSProgress localizedAdditionalDescription];
	return result;
}

void C_NSProgress_SetLocalizedAdditionalDescription(void* ptr, void* value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setLocalizedAdditionalDescription:(NSString*)value];
}

double C_NSProgress_FractionCompleted(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	double result = [nSProgress fractionCompleted];
	return result;
}

bool C_NSProgress_IsCancellable(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	bool result = [nSProgress isCancellable];
	return result;
}

void C_NSProgress_SetCancellable(void* ptr, bool value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setCancellable:value];
}

bool C_NSProgress_IsCancelled(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	bool result = [nSProgress isCancelled];
	return result;
}

bool C_NSProgress_IsPausable(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	bool result = [nSProgress isPausable];
	return result;
}

void C_NSProgress_SetPausable(void* ptr, bool value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setPausable:value];
}

bool C_NSProgress_IsPaused(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	bool result = [nSProgress isPaused];
	return result;
}

bool C_NSProgress_IsIndeterminate(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	bool result = [nSProgress isIndeterminate];
	return result;
}

void* C_NSProgress_Kind(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	NSString* result = [nSProgress kind];
	return result;
}

void C_NSProgress_SetKind(void* ptr, void* value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setKind:(NSString*)value];
}

void* C_NSProgress_FileOperationKind(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	NSString* result = [nSProgress fileOperationKind];
	return result;
}

void C_NSProgress_SetFileOperationKind(void* ptr, void* value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setFileOperationKind:(NSString*)value];
}

void* C_NSProgress_FileURL(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	NSURL* result = [nSProgress fileURL];
	return result;
}

void C_NSProgress_SetFileURL(void* ptr, void* value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setFileURL:(NSURL*)value];
}

bool C_NSProgress_IsFinished(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	bool result = [nSProgress isFinished];
	return result;
}

bool C_NSProgress_IsOld(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	bool result = [nSProgress isOld];
	return result;
}

void* C_NSProgress_EstimatedTimeRemaining(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	NSNumber* result = [nSProgress estimatedTimeRemaining];
	return result;
}

void C_NSProgress_SetEstimatedTimeRemaining(void* ptr, void* value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setEstimatedTimeRemaining:(NSNumber*)value];
}

void* C_NSProgress_FileCompletedCount(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	NSNumber* result = [nSProgress fileCompletedCount];
	return result;
}

void C_NSProgress_SetFileCompletedCount(void* ptr, void* value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setFileCompletedCount:(NSNumber*)value];
}

void* C_NSProgress_FileTotalCount(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	NSNumber* result = [nSProgress fileTotalCount];
	return result;
}

void C_NSProgress_SetFileTotalCount(void* ptr, void* value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setFileTotalCount:(NSNumber*)value];
}

void* C_NSProgress_Throughput(void* ptr) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	NSNumber* result = [nSProgress throughput];
	return result;
}

void C_NSProgress_SetThroughput(void* ptr, void* value) {
	NSProgress* nSProgress = (NSProgress*)ptr;
	[nSProgress setThroughput:(NSNumber*)value];
}
