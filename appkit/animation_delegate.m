#import "animation_delegate.h"
#import "_cgo_export.h"

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
