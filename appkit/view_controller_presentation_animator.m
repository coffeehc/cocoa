#import "view_controller_presentation_animator.h"
#import "_cgo_export.h"
#import <AppKit/NSViewController.h>

@interface NSViewControllerPresentationAnimatorAdaptor : NSObject <NSViewControllerPresentationAnimator>
@property (assign) uintptr_t goID;
@end

@implementation NSViewControllerPresentationAnimatorAdaptor


- (void)animatePresentationOfViewController:(NSViewController*)viewController fromViewController:(NSViewController*)fromViewController {
    viewControllerPresentationAnimator_AnimatePresentationOfViewController_FromViewController([self goID], viewController, fromViewController);
}

- (void)animateDismissalOfViewController:(NSViewController*)viewController fromViewController:(NSViewController*)fromViewController {
    viewControllerPresentationAnimator_AnimateDismissalOfViewController_FromViewController([self goID], viewController, fromViewController);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return ViewControllerPresentationAnimator_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapViewControllerPresentationAnimator(uintptr_t goID) {
    NSViewControllerPresentationAnimatorAdaptor* adaptor = [[NSViewControllerPresentationAnimatorAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
