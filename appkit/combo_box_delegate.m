#import <Appkit/Appkit.h>
#import "combo_box_delegate.h"
#import "_cgo_export.h"

@implementation NSComboBoxDelegateAdaptor


- (void)comboBoxSelectionDidChange:(NSNotification*)notification {
    comboBoxDelegate_ComboBoxSelectionDidChange([self goID], notification);
}

- (void)comboBoxSelectionIsChanging:(NSNotification*)notification {
    comboBoxDelegate_ComboBoxSelectionIsChanging([self goID], notification);
}

- (void)comboBoxWillDismiss:(NSNotification*)notification {
    comboBoxDelegate_ComboBoxWillDismiss([self goID], notification);
}

- (void)comboBoxWillPopUp:(NSNotification*)notification {
    comboBoxDelegate_ComboBoxWillPopUp([self goID], notification);
}

- (NSArray*)textField:(NSTextField*)textField textView:(NSTextView*)textView candidates:(NSArray*)candidates forSelectedRange:(NSRange)selectedRange {
    Array candidatesArray;
    int candidatescount = [candidates count];
    if (candidatescount > 0) {
    	void** candidatesData = malloc(candidatescount * sizeof(void*));
    	for (int i = 0; i < candidatescount; i++) {
    		 void* p = [candidates objectAtIndex:i];
    		 candidatesData[i] = p;
    	}
    	candidatesArray.data = candidatesData;
    	candidatesArray.len = candidatescount;
    }
    Array result_ = comboBoxDelegate_TextField_TextView_Candidates_ForSelectedRange([self goID], textField, textView, candidatesArray, selectedRange);
    NSMutableArray* objcResult_ = [[NSMutableArray alloc] init];
    if (result_.len > 0) {
    	void** result_Data = (void**)result_.data;
    	for (int i = 0; i < result_.len; i++) {
    		void* p = result_Data[i];
    		[objcResult_ addObject:(NSTextCheckingResult*)(NSTextCheckingResult*)p];
    	}
    }
    return objcResult_;
}

- (NSArray*)textField:(NSTextField*)textField textView:(NSTextView*)textView candidatesForSelectedRange:(NSRange)selectedRange {
    Array result_ = comboBoxDelegate_TextField_TextView_CandidatesForSelectedRange([self goID], textField, textView, selectedRange);
    NSMutableArray* objcResult_ = [[NSMutableArray alloc] init];
    if (result_.len > 0) {
    	void** result_Data = (void**)result_.data;
    	for (int i = 0; i < result_.len; i++) {
    		void* p = result_Data[i];
    		[objcResult_ addObject:(NSObject*)(NSObject*)p];
    	}
    }
    return objcResult_;
}

- (BOOL)textField:(NSTextField*)textField textView:(NSTextView*)textView shouldSelectCandidateAtIndex:(NSUInteger)index {
    bool result_ = comboBoxDelegate_TextField_TextView_ShouldSelectCandidateAtIndex([self goID], textField, textView, index);
    return result_;
}

- (BOOL)control:(NSControl*)control isValidObject:(id)obj {
    bool result_ = comboBoxDelegate_Control_IsValidObject([self goID], control, obj);
    return result_;
}

- (void)control:(NSControl*)control didFailToValidatePartialString:(NSString*)_string errorDescription:(NSString*)error {
    comboBoxDelegate_Control_DidFailToValidatePartialString_ErrorDescription([self goID], control, _string, error);
}

- (BOOL)control:(NSControl*)control didFailToFormatString:(NSString*)_string errorDescription:(NSString*)error {
    bool result_ = comboBoxDelegate_Control_DidFailToFormatString_ErrorDescription([self goID], control, _string, error);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldBeginEditing:(NSText*)fieldEditor {
    bool result_ = comboBoxDelegate_Control_TextShouldBeginEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textShouldEndEditing:(NSText*)fieldEditor {
    bool result_ = comboBoxDelegate_Control_TextShouldEndEditing([self goID], control, fieldEditor);
    return result_;
}

- (BOOL)control:(NSControl*)control textView:(NSTextView*)textView doCommandBySelector:(SEL)commandSelector {
    bool result_ = comboBoxDelegate_Control_TextView_DoCommandBySelector([self goID], control, textView, commandSelector);
    return result_;
}

- (void)controlTextDidBeginEditing:(NSNotification*)obj {
    comboBoxDelegate_ControlTextDidBeginEditing([self goID], obj);
}

- (void)controlTextDidChange:(NSNotification*)obj {
    comboBoxDelegate_ControlTextDidChange([self goID], obj);
}

- (void)controlTextDidEndEditing:(NSNotification*)obj {
    comboBoxDelegate_ControlTextDidEndEditing([self goID], obj);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return ComboBoxDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapComboBoxDelegate(uintptr_t goID) {
    NSComboBoxDelegateAdaptor* adaptor = [[NSComboBoxDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
