#import <Appkit/Appkit.h>
#import "text_delegate.h"
#import "_cgo_export.h"

@implementation NSTextDelegateAdaptor


- (void)textDidChange:(NSNotification*)notification {
    TextDelegate_TextDidChange([self goID], notification);
}

- (BOOL)textShouldBeginEditing:(NSText*)textObject {
    bool result_ = TextDelegate_TextShouldBeginEditing([self goID], textObject);
    return result_;
}

- (void)textDidBeginEditing:(NSNotification*)notification {
    TextDelegate_TextDidBeginEditing([self goID], notification);
}

- (BOOL)textShouldEndEditing:(NSText*)textObject {
    bool result_ = TextDelegate_TextShouldEndEditing([self goID], textObject);
    return result_;
}

- (void)textDidEndEditing:(NSNotification*)notification {
    TextDelegate_TextDidEndEditing([self goID], notification);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TextDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteTextDelegate([self goID]);
	[super dealloc];
}
@end

void* WrapTextDelegate(long goID) {
    NSTextDelegateAdaptor* adaptor = [[NSTextDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
