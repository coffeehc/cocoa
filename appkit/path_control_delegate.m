#import <Appkit/Appkit.h>
#import "path_control_delegate.h"
#import "_cgo_export.h"

@implementation NSPathControlDelegateAdaptor


- (BOOL)pathControl:(NSPathControl*)pathControl shouldDragPathComponentCell:(NSPathComponentCell*)pathComponentCell withPasteboard:(NSPasteboard*)pasteboard {
    bool result_ = pathControlDelegate_PathControl_ShouldDragPathComponentCell_WithPasteboard([self goID], pathControl, pathComponentCell, pasteboard);
    return result_;
}

- (NSDragOperation)pathControl:(NSPathControl*)pathControl validateDrop:(id)info {
    unsigned int result_ = pathControlDelegate_PathControl_ValidateDrop([self goID], pathControl, info);
    return result_;
}

- (BOOL)pathControl:(NSPathControl*)pathControl acceptDrop:(id)info {
    bool result_ = pathControlDelegate_PathControl_AcceptDrop([self goID], pathControl, info);
    return result_;
}

- (void)pathControl:(NSPathControl*)pathControl willDisplayOpenPanel:(NSOpenPanel*)openPanel {
    pathControlDelegate_PathControl_WillDisplayOpenPanel([self goID], pathControl, openPanel);
}

- (void)pathControl:(NSPathControl*)pathControl willPopUpMenu:(NSMenu*)menu {
    pathControlDelegate_PathControl_WillPopUpMenu([self goID], pathControl, menu);
}

- (BOOL)pathControl:(NSPathControl*)pathControl shouldDragItem:(NSPathControlItem*)pathItem withPasteboard:(NSPasteboard*)pasteboard {
    bool result_ = pathControlDelegate_PathControl_ShouldDragItem_WithPasteboard([self goID], pathControl, pathItem, pasteboard);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return PathControlDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapPathControlDelegate(uintptr_t goID) {
    NSPathControlDelegateAdaptor* adaptor = [[NSPathControlDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
