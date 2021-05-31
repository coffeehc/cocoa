#import <Appkit/Appkit.h>
#import "dragging_info.h"
#import "_cgo_export.h"

@implementation NSDraggingInfoAdaptor


- (void)slideDraggedImageTo:(NSPoint)screenPoint {
    draggingInfo_SlideDraggedImageTo([self goID], screenPoint);
}

- (void)resetSpringLoading {
    draggingInfo_ResetSpringLoading([self goID]);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return DraggingInfo_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteDraggingInfo([self goID]);
	[super dealloc];
}
@end

void* WrapDraggingInfo(long goID) {
    NSDraggingInfoAdaptor* adaptor = [[NSDraggingInfoAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
