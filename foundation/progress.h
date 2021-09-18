#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_Progress_Alloc();

void* C_NSProgress_InitWithParent_UserInfo(void* ptr, void* parentProgressOrNil, Dictionary userInfoOrNil);
void* C_NSProgress_AllocProgress();
void* C_NSProgress_Init(void* ptr);
void* C_NSProgress_NewProgress();
void* C_NSProgress_Autorelease(void* ptr);
void* C_NSProgress_Retain(void* ptr);
void* C_NSProgress_DiscreteProgressWithTotalUnitCount(long unitCount);
void* C_NSProgress_ProgressWithTotalUnitCount(long unitCount);
void* C_NSProgress_ProgressWithTotalUnitCount_Parent_PendingUnitCount(long unitCount, void* parent, long portionOfParentTotalUnitCount);
void* C_NSProgress_CurrentProgress();
void C_NSProgress_BecomeCurrentWithPendingUnitCount(void* ptr, long unitCount);
void C_NSProgress_AddChild_WithPendingUnitCount(void* ptr, void* child, long inUnitCount);
void C_NSProgress_ResignCurrent(void* ptr);
void C_NSProgress_Cancel(void* ptr);
void C_NSProgress_Pause(void* ptr);
void C_NSProgress_Resume(void* ptr);
void C_NSProgress_SetUserInfoObject_ForKey(void* ptr, void* objectOrNil, void* key);
void C_NSProgress_Publish(void* ptr);
void C_NSProgress_Unpublish(void* ptr);
void C_NSProgress_Progress_RemoveSubscriber(void* subscriber);
long C_NSProgress_TotalUnitCount(void* ptr);
void C_NSProgress_SetTotalUnitCount(void* ptr, long value);
long C_NSProgress_CompletedUnitCount(void* ptr);
void C_NSProgress_SetCompletedUnitCount(void* ptr, long value);
void* C_NSProgress_LocalizedDescription(void* ptr);
void C_NSProgress_SetLocalizedDescription(void* ptr, void* value);
void* C_NSProgress_LocalizedAdditionalDescription(void* ptr);
void C_NSProgress_SetLocalizedAdditionalDescription(void* ptr, void* value);
double C_NSProgress_FractionCompleted(void* ptr);
bool C_NSProgress_IsCancellable(void* ptr);
void C_NSProgress_SetCancellable(void* ptr, bool value);
bool C_NSProgress_IsCancelled(void* ptr);
bool C_NSProgress_IsPausable(void* ptr);
void C_NSProgress_SetPausable(void* ptr, bool value);
bool C_NSProgress_IsPaused(void* ptr);
bool C_NSProgress_IsIndeterminate(void* ptr);
void* C_NSProgress_Kind(void* ptr);
void C_NSProgress_SetKind(void* ptr, void* value);
Dictionary C_NSProgress_UserInfo(void* ptr);
void* C_NSProgress_FileOperationKind(void* ptr);
void C_NSProgress_SetFileOperationKind(void* ptr, void* value);
void* C_NSProgress_FileURL(void* ptr);
void C_NSProgress_SetFileURL(void* ptr, void* value);
bool C_NSProgress_IsFinished(void* ptr);
bool C_NSProgress_IsOld(void* ptr);
void* C_NSProgress_EstimatedTimeRemaining(void* ptr);
void C_NSProgress_SetEstimatedTimeRemaining(void* ptr, void* value);
void* C_NSProgress_FileCompletedCount(void* ptr);
void C_NSProgress_SetFileCompletedCount(void* ptr, void* value);
void* C_NSProgress_FileTotalCount(void* ptr);
void C_NSProgress_SetFileTotalCount(void* ptr, void* value);
void* C_NSProgress_Throughput(void* ptr);
void C_NSProgress_SetThroughput(void* ptr, void* value);
