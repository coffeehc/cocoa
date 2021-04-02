#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "text_view.h"

bool TextView_AllowsUndo(void* ptr) {
	NSTextView* textView = (NSTextView*)ptr;
	return [textView allowsUndo];
}

void TextView_SetAllowsUndo(void* ptr, bool allowsUndo) {
	NSTextView* textView = (NSTextView*)ptr;
	[textView setAllowsUndo:allowsUndo];
}

void* TextView_TextContainer(void* ptr) {
	NSTextView* textView = (NSTextView*)ptr;
	return [textView textContainer];
}

void TextView_SetTextContainer(void* ptr, void* textContainer) {
	NSTextView* textView = (NSTextView*)ptr;
	[textView setTextContainer:(NSTextContainer*)textContainer];
}

void* TextView_NewTextView(NSRect frame) {
	NSTextView* textView = [NSTextView alloc];
	return [[textView initWithFrame:frame] autorelease];
}

void* TextView_ScrollableTextView() {
	return [NSTextView scrollableTextView];
}
