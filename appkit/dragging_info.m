#import "dragging_info.h"
#import "_cgo_export.h"
#import <AppKit/NSDragging.h>

@interface NSDraggingInfoAdaptor : NSObject <NSDraggingInfo>
@property (assign) uintptr_t goID;
@end

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
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapDraggingInfo(uintptr_t goID) {
    NSDraggingInfoAdaptor* adaptor = [[NSDraggingInfoAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
