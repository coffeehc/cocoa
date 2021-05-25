#import <Appkit/Appkit.h>
#import "view_animation.h"

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

void* C_NSViewAnimation_Init(void* ptr) {
    NSViewAnimation* nSViewAnimation = (NSViewAnimation*)ptr;
    NSViewAnimation* result_ = [nSViewAnimation init];
    return result_;
}
