#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_Animation_Alloc();

void* C_NSAnimation_InitWithDuration_AnimationCurve(void* ptr, double duration, unsigned int animationCurve);
void* C_NSAnimation_InitWithCoder(void* ptr, void* coder);
void* C_NSAnimation_AllocAnimation();
void* C_NSAnimation_Init(void* ptr);
void* C_NSAnimation_NewAnimation();
void* C_NSAnimation_Autorelease(void* ptr);
void* C_NSAnimation_Retain(void* ptr);
void C_NSAnimation_StartAnimation(void* ptr);
void C_NSAnimation_StopAnimation(void* ptr);
void C_NSAnimation_AddProgressMark(void* ptr, float progressMark);
void C_NSAnimation_RemoveProgressMark(void* ptr, float progressMark);
void C_NSAnimation_StartWhenAnimation_ReachesProgress(void* ptr, void* animation, float startProgress);
void C_NSAnimation_StopWhenAnimation_ReachesProgress(void* ptr, void* animation, float stopProgress);
void C_NSAnimation_ClearStartAnimation(void* ptr);
void C_NSAnimation_ClearStopAnimation(void* ptr);
unsigned int C_NSAnimation_AnimationBlockingMode(void* ptr);
void C_NSAnimation_SetAnimationBlockingMode(void* ptr, unsigned int value);
Array C_NSAnimation_RunLoopModesForAnimating(void* ptr);
unsigned int C_NSAnimation_AnimationCurve(void* ptr);
void C_NSAnimation_SetAnimationCurve(void* ptr, unsigned int value);
double C_NSAnimation_Duration(void* ptr);
void C_NSAnimation_SetDuration(void* ptr, double value);
float C_NSAnimation_FrameRate(void* ptr);
void C_NSAnimation_SetFrameRate(void* ptr, float value);
void* C_NSAnimation_Delegate(void* ptr);
void C_NSAnimation_SetDelegate(void* ptr, void* value);
bool C_NSAnimation_IsAnimating(void* ptr);
float C_NSAnimation_CurrentProgress(void* ptr);
void C_NSAnimation_SetCurrentProgress(void* ptr, float value);
float C_NSAnimation_CurrentValue(void* ptr);
Array C_NSAnimation_ProgressMarks(void* ptr);
void C_NSAnimation_SetProgressMarks(void* ptr, Array value);
