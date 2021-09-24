#import "dragging_destination.h"
#import <AppKit/NSDragging.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

@interface NSDraggingDestinationAdaptor : NSObject <NSDraggingDestination>
@property (assign) uintptr_t goID;
@end

@implementation NSDraggingDestinationAdaptor


- (NSDragOperation)draggingEntered:(id)sender {
    unsigned int result_ = draggingDestination_DraggingEntered([self goID], sender);
    return result_;
}

- (BOOL)wantsPeriodicDraggingUpdates {
    bool result_ = draggingDestination_WantsPeriodicDraggingUpdates([self goID]);
    return result_;
}

- (NSDragOperation)draggingUpdated:(id)sender {
    unsigned int result_ = draggingDestination_DraggingUpdated([self goID], sender);
    return result_;
}

- (void)draggingEnded:(id)sender {
    draggingDestination_DraggingEnded([self goID], sender);
}

- (void)draggingExited:(id)sender {
    draggingDestination_DraggingExited([self goID], sender);
}

- (BOOL)prepareForDragOperation:(id)sender {
    bool result_ = draggingDestination_PrepareForDragOperation([self goID], sender);
    return result_;
}

- (BOOL)performDragOperation:(id)sender {
    bool result_ = draggingDestination_PerformDragOperation([self goID], sender);
    return result_;
}

- (void)concludeDragOperation:(id)sender {
    draggingDestination_ConcludeDragOperation([self goID], sender);
}

- (void)updateDraggingItemsForDrag:(id)sender {
    draggingDestination_UpdateDraggingItemsForDrag([self goID], sender);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return DraggingDestination_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapDraggingDestination(uintptr_t goID) {
    NSDraggingDestinationAdaptor* adaptor = [[NSDraggingDestinationAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
