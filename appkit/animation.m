#import "animation.h"
#import <AppKit/NSAnimation.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

void* C_Animation_Alloc() {
    return [NSAnimation alloc];
}

void* C_NSAnimation_InitWithDuration_AnimationCurve(void* ptr, double duration, unsigned int animationCurve) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSAnimation* result_ = [nSAnimation initWithDuration:duration animationCurve:animationCurve];
    return result_;
}

void* C_NSAnimation_InitWithCoder(void* ptr, void* coder) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSAnimation* result_ = [nSAnimation initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSAnimation_AllocAnimation() {
    NSAnimation* result_ = [NSAnimation alloc];
    return result_;
}

void* C_NSAnimation_Init(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSAnimation* result_ = [nSAnimation init];
    return result_;
}

void* C_NSAnimation_NewAnimation() {
    NSAnimation* result_ = [NSAnimation new];
    return result_;
}

void* C_NSAnimation_Autorelease(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSAnimation* result_ = [nSAnimation autorelease];
    return result_;
}

void* C_NSAnimation_Retain(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSAnimation* result_ = [nSAnimation retain];
    return result_;
}

void C_NSAnimation_StartAnimation(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation startAnimation];
}

void C_NSAnimation_StopAnimation(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation stopAnimation];
}

void C_NSAnimation_AddProgressMark(void* ptr, float progressMark) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation addProgressMark:progressMark];
}

void C_NSAnimation_RemoveProgressMark(void* ptr, float progressMark) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation removeProgressMark:progressMark];
}

void C_NSAnimation_StartWhenAnimation_ReachesProgress(void* ptr, void* animation, float startProgress) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation startWhenAnimation:(NSAnimation*)animation reachesProgress:startProgress];
}

void C_NSAnimation_StopWhenAnimation_ReachesProgress(void* ptr, void* animation, float stopProgress) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation stopWhenAnimation:(NSAnimation*)animation reachesProgress:stopProgress];
}

void C_NSAnimation_ClearStartAnimation(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation clearStartAnimation];
}

void C_NSAnimation_ClearStopAnimation(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation clearStopAnimation];
}

unsigned int C_NSAnimation_AnimationBlockingMode(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSAnimationBlockingMode result_ = [nSAnimation animationBlockingMode];
    return result_;
}

void C_NSAnimation_SetAnimationBlockingMode(void* ptr, unsigned int value) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation setAnimationBlockingMode:value];
}

Array C_NSAnimation_RunLoopModesForAnimating(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSArray* result_ = [nSAnimation runLoopModesForAnimating];
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

unsigned int C_NSAnimation_AnimationCurve(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSAnimationCurve result_ = [nSAnimation animationCurve];
    return result_;
}

void C_NSAnimation_SetAnimationCurve(void* ptr, unsigned int value) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation setAnimationCurve:value];
}

double C_NSAnimation_Duration(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSTimeInterval result_ = [nSAnimation duration];
    return result_;
}

void C_NSAnimation_SetDuration(void* ptr, double value) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation setDuration:value];
}

float C_NSAnimation_FrameRate(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    float result_ = [nSAnimation frameRate];
    return result_;
}

void C_NSAnimation_SetFrameRate(void* ptr, float value) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation setFrameRate:value];
}

void* C_NSAnimation_Delegate(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    id result_ = [nSAnimation delegate];
    return result_;
}

void C_NSAnimation_SetDelegate(void* ptr, void* value) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation setDelegate:(id)value];
}

bool C_NSAnimation_IsAnimating(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    BOOL result_ = [nSAnimation isAnimating];
    return result_;
}

float C_NSAnimation_CurrentProgress(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSAnimationProgress result_ = [nSAnimation currentProgress];
    return result_;
}

void C_NSAnimation_SetCurrentProgress(void* ptr, float value) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    [nSAnimation setCurrentProgress:value];
}

float C_NSAnimation_CurrentValue(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    float result_ = [nSAnimation currentValue];
    return result_;
}

Array C_NSAnimation_ProgressMarks(void* ptr) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSArray* result_ = [nSAnimation progressMarks];
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

void C_NSAnimation_SetProgressMarks(void* ptr, Array value) {
    NSAnimation* nSAnimation = (NSAnimation*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSNumber*)p];
    	}
    }
    [nSAnimation setProgressMarks:objcValue];
}

void* C_ViewAnimation_Alloc() {
    return [NSViewAnimation alloc];
}

void* C_NSViewAnimation_InitWithDuration_AnimationCurve(void* ptr, double duration, unsigned int animationCurve) {
    NSViewAnimation* nSViewAnimation = (NSViewAnimation*)ptr;
    NSViewAnimation* result_ = [nSViewAnimation initWithDuration:duration animationCurve:animationCurve];
    return result_;
}

void* C_NSViewAnimation_InitWithCoder(void* ptr, void* coder) {
    NSViewAnimation* nSViewAnimation = (NSViewAnimation*)ptr;
    NSViewAnimation* result_ = [nSViewAnimation initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSViewAnimation_AllocViewAnimation() {
    NSViewAnimation* result_ = [NSViewAnimation alloc];
    return result_;
}

void* C_NSViewAnimation_Init(void* ptr) {
    NSViewAnimation* nSViewAnimation = (NSViewAnimation*)ptr;
    NSViewAnimation* result_ = [nSViewAnimation init];
    return result_;
}

void* C_NSViewAnimation_NewViewAnimation() {
    NSViewAnimation* result_ = [NSViewAnimation new];
    return result_;
}

void* C_NSViewAnimation_Autorelease(void* ptr) {
    NSViewAnimation* nSViewAnimation = (NSViewAnimation*)ptr;
    NSViewAnimation* result_ = [nSViewAnimation autorelease];
    return result_;
}

void* C_NSViewAnimation_Retain(void* ptr) {
    NSViewAnimation* nSViewAnimation = (NSViewAnimation*)ptr;
    NSViewAnimation* result_ = [nSViewAnimation retain];
    return result_;
}

@interface NSAnimationDelegateAdaptor : NSObject <NSAnimationDelegate>
@property (assign) uintptr_t goID;
@end

@implementation NSAnimationDelegateAdaptor


- (void)animationDidEnd:(NSAnimation*)animation {
    animationDelegate_AnimationDidEnd([self goID], animation);
}

- (void)animationDidStop:(NSAnimation*)animation {
    animationDelegate_AnimationDidStop([self goID], animation);
}

- (BOOL)animationShouldStart:(NSAnimation*)animation {
    bool result_ = animationDelegate_AnimationShouldStart([self goID], animation);
    return result_;
}

- (float)animation:(NSAnimation*)animation valueForProgress:(NSAnimationProgress)progress {
    float result_ = animationDelegate_Animation_ValueForProgress([self goID], animation, progress);
    return result_;
}

- (void)animation:(NSAnimation*)animation didReachProgressMark:(NSAnimationProgress)progress {
    animationDelegate_Animation_DidReachProgressMark([self goID], animation, progress);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return AnimationDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapAnimationDelegate(uintptr_t goID) {
    NSAnimationDelegateAdaptor* adaptor = [[NSAnimationDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
