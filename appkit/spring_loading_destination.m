#import "spring_loading_destination.h"
#import "_cgo_export.h"

@implementation NSSpringLoadingDestinationAdaptor


- (void)springLoadingActivated:(BOOL)activated draggingInfo:(id)draggingInfo {
    springLoadingDestination_SpringLoadingActivated_DraggingInfo([self goID], activated, draggingInfo);
}

- (void)springLoadingHighlightChanged:(id)draggingInfo {
    springLoadingDestination_SpringLoadingHighlightChanged([self goID], draggingInfo);
}

- (NSSpringLoadingOptions)springLoadingEntered:(id)draggingInfo {
    unsigned int result_ = springLoadingDestination_SpringLoadingEntered([self goID], draggingInfo);
    return result_;
}

- (NSSpringLoadingOptions)springLoadingUpdated:(id)draggingInfo {
    unsigned int result_ = springLoadingDestination_SpringLoadingUpdated([self goID], draggingInfo);
    return result_;
}

- (void)springLoadingExited:(id)draggingInfo {
    springLoadingDestination_SpringLoadingExited([self goID], draggingInfo);
}

- (void)draggingEnded:(id)draggingInfo {
    springLoadingDestination_DraggingEnded([self goID], draggingInfo);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return SpringLoadingDestination_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapSpringLoadingDestination(uintptr_t goID) {
    NSSpringLoadingDestinationAdaptor* adaptor = [[NSSpringLoadingDestinationAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
