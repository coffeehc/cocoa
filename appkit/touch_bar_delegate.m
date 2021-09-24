#import "touch_bar_delegate.h"
#import <AppKit/NSTouchBar.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

@interface NSTouchBarDelegateAdaptor : NSObject <NSTouchBarDelegate>
@property (assign) uintptr_t goID;
@end

@implementation NSTouchBarDelegateAdaptor


- (NSTouchBarItem*)touchBar:(NSTouchBar*)touchBar makeItemForIdentifier:(NSTouchBarItemIdentifier)identifier {
    void* result_ = touchBarDelegate_TouchBar_MakeItemForIdentifier([self goID], touchBar, identifier);
    return (NSTouchBarItem*)result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TouchBarDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapTouchBarDelegate(uintptr_t goID) {
    NSTouchBarDelegateAdaptor* adaptor = [[NSTouchBarDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
