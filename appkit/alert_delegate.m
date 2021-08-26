#import <Appkit/Appkit.h>
#import "alert_delegate.h"
#import "_cgo_export.h"

@implementation NSAlertDelegateAdaptor


- (BOOL)alertShowHelp:(NSAlert*)alert {
    bool result_ = alertDelegate_AlertShowHelp([self goID], alert);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return AlertDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapAlertDelegate(uintptr_t goID) {
    NSAlertDelegateAdaptor* adaptor = [[NSAlertDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
