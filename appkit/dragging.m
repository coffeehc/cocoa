#import "dragging.h"
#import <AppKit/NSDragging.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

@interface NSDraggingSourceAdaptor : NSObject <NSDraggingSource>
@property (assign) uintptr_t goID;
@end

@implementation NSDraggingSourceAdaptor


- (NSDragOperation)draggingSession:(NSDraggingSession*)session sourceOperationMaskForDraggingContext:(NSDraggingContext)context {
    unsigned int result_ = draggingSource_DraggingSession_SourceOperationMaskForDraggingContext([self goID], session, context);
    return result_;
}

- (void)draggingSession:(NSDraggingSession*)session endedAtPoint:(NSPoint)screenPoint operation:(NSDragOperation)operation {
    draggingSource_DraggingSession_EndedAtPoint_Operation([self goID], session, screenPoint, operation);
}

- (void)draggingSession:(NSDraggingSession*)session movedToPoint:(NSPoint)screenPoint {
    draggingSource_DraggingSession_MovedToPoint([self goID], session, screenPoint);
}

- (void)draggingSession:(NSDraggingSession*)session willBeginAtPoint:(NSPoint)screenPoint {
    draggingSource_DraggingSession_WillBeginAtPoint([self goID], session, screenPoint);
}

- (BOOL)ignoreModifierKeysForDraggingSession:(NSDraggingSession*)session {
    bool result_ = draggingSource_IgnoreModifierKeysForDraggingSession([self goID], session);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return DraggingSource_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapDraggingSource(uintptr_t goID) {
    NSDraggingSourceAdaptor* adaptor = [[NSDraggingSourceAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}

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

@interface NSSpringLoadingDestinationAdaptor : NSObject <NSSpringLoadingDestination>
@property (assign) uintptr_t goID;
@end

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
