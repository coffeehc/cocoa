#import <Appkit/Appkit.h>
#import "animatable_property_container.h"
#import "_cgo_export.h"

@implementation NSAnimatablePropertyContainerAdaptor


- (NSObject*)animator {
    void* result_ = animatablePropertyContainer_Animator([self goID]);
    return (NSObject*)result_;
}

- (id)animationForKey:(NSAnimatablePropertyKey)key {
    void* result_ = animatablePropertyContainer_AnimationForKey([self goID], key);
    return (id)result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return AnimatablePropertyContainer_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapAnimatablePropertyContainer(uintptr_t goID) {
    NSAnimatablePropertyContainerAdaptor* adaptor = [[NSAnimatablePropertyContainerAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
