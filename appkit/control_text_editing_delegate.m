#import <Appkit/Appkit.h>
#import "control_text_editing_delegate.h"
#import "_cgo_export.h"

@implementation NSControlTextEditingDelegateAdaptor


- (BOOL)control:(NSControl*)control isValidObject:(id)obj {
    bool result_ = ControlTextEditingDelegate_Control_IsValidObject([self goID], control, obj);
    return result_;
}

- (void)control:(NSControl*)control didFailToValidatePartialString:(NSString*)_string errorDescription:(NSString*)error {
    ControlTextEditingDelegate_Control_DidFailToValidatePartialString_ErrorDescription([self goID], control, _string, error);
}

- (BOOL)control:(NSControl*)control didFailToFormatString:(NSString*)_string errorDescription:(NSString*)error {
    bool result_ = ControlTextEditingDelegate_Control_DidFailToFormatString_ErrorDescription([self goID], control, _string, error);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldBeginEditing:(NSText*)fieldEditor {
    bool result_ = ControlTextEditingDelegate_Control_TextShouldBeginEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldEndEditing:(NSText*)fieldEditor {
    bool result_ = ControlTextEditingDelegate_Control_TextShouldEndEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textView:(NSTextView*)textView doCommandBySelector:(SEL)commandSelector {
    bool result_ = ControlTextEditingDelegate_Control_TextView_DoCommandBySelector([self goID], control, textView, commandSelector);
    return result_;
}

- (void)controlTextDidBeginEditing:(NSNotification*)obj {
    ControlTextEditingDelegate_ControlTextDidBeginEditing([self goID], obj);
}

- (void)controlTextDidChange:(NSNotification*)obj {
    ControlTextEditingDelegate_ControlTextDidChange([self goID], obj);
}

- (void)controlTextDidEndEditing:(NSNotification*)obj {
    ControlTextEditingDelegate_ControlTextDidEndEditing([self goID], obj);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return ControlTextEditingDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteControlTextEditingDelegate([self goID]);
	[super dealloc];
}
@end

void* WrapControlTextEditingDelegate(long goID) {
    NSControlTextEditingDelegateAdaptor* adaptor = [[NSControlTextEditingDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
