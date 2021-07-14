#import <Appkit/Appkit.h>
#import "path_cell_delegate.h"
#import "_cgo_export.h"

@implementation NSPathCellDelegateAdaptor


- (void)pathCell:(NSPathCell*)pathCell willDisplayOpenPanel:(NSOpenPanel*)openPanel {
    pathCellDelegate_PathCell_WillDisplayOpenPanel([self goID], pathCell, openPanel);
}

- (void)pathCell:(NSPathCell*)pathCell willPopUpMenu:(NSMenu*)menu {
    pathCellDelegate_PathCell_WillPopUpMenu([self goID], pathCell, menu);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return PathCellDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deletePathCellDelegate([self goID]);
	[super dealloc];
}
@end

void* WrapPathCellDelegate(uintptr_t goID) {
    NSPathCellDelegateAdaptor* adaptor = [[NSPathCellDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
