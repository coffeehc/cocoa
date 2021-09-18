#import "text_delegate.h"
#import "_cgo_export.h"
#import <AppKit/NSText.h>

@interface NSTextDelegateAdaptor : NSObject <NSTextDelegate>
@property (assign) uintptr_t goID;
@end

@implementation NSTextDelegateAdaptor


- (void)textDidChange:(NSNotification*)notification {
    textDelegate_TextDidChange([self goID], notification);
}

- (BOOL)textShouldBeginEditing:(NSText*)textObject {
    bool result_ = textDelegate_TextShouldBeginEditing([self goID], textObject);
    return result_;
}

- (void)textDidBeginEditing:(NSNotification*)notification {
    textDelegate_TextDidBeginEditing([self goID], notification);
}

- (BOOL)textShouldEndEditing:(NSText*)textObject {
    bool result_ = textDelegate_TextShouldEndEditing([self goID], textObject);
    return result_;
}

- (void)textDidEndEditing:(NSNotification*)notification {
    textDelegate_TextDidEndEditing([self goID], notification);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TextDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapTextDelegate(uintptr_t goID) {
    NSTextDelegateAdaptor* adaptor = [[NSTextDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
