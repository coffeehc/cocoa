#import <Appkit/Appkit.h>
#import "dragging_source.h"
#import "_cgo_export.h"

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
