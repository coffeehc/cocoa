#import <Appkit/Appkit.h>
#import "touch_bar_delegate.h"
#import "_cgo_export.h"

@implementation NSTouchBarDelegateAdaptor


- (NSTouchBarItem*)touchBar:(NSTouchBar*)touchBar makeItemForIdentifier:(NSTouchBarItemIdentifier)identifier {
    void* result_ = touchBarDelegate_TouchBar_MakeItemForIdentifier([self goID], touchBar, identifier);
    return (NSTouchBarItem*)result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TouchBarDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteTouchBarDelegate([self goID]);
	[super dealloc];
}
@end

void* WrapTouchBarDelegate(long goID) {
    NSTouchBarDelegateAdaptor* adaptor = [[NSTouchBarDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
