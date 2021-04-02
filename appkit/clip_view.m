#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "clip_view.h"

NSEdgeInsets ClipView_ContentInsets(void* ptr) {
	NSClipView* clipView = (NSClipView*)ptr;
	return [clipView contentInsets];
}

void ClipView_SetContentInsets(void* ptr, NSEdgeInsets contentInsets) {
	NSClipView* clipView = (NSClipView*)ptr;
	[clipView setContentInsets:contentInsets];
}

void* ClipView_DocumentView(void* ptr) {
	NSClipView* clipView = (NSClipView*)ptr;
	return [clipView documentView];
}

void ClipView_SetDocumentView(void* ptr, void* documentView) {
	NSClipView* clipView = (NSClipView*)ptr;
	[clipView setDocumentView:(NSView*)documentView];
}

bool ClipView_AutomaticallyAdjustsContentInsets(void* ptr) {
	NSClipView* clipView = (NSClipView*)ptr;
	return [clipView automaticallyAdjustsContentInsets];
}

void ClipView_SetAutomaticallyAdjustsContentInsets(void* ptr, bool automaticallyAdjustsContentInsets) {
	NSClipView* clipView = (NSClipView*)ptr;
	[clipView setAutomaticallyAdjustsContentInsets:automaticallyAdjustsContentInsets];
}

NSRect ClipView_DocumentRect(void* ptr) {
	NSClipView* clipView = (NSClipView*)ptr;
	return [clipView documentRect];
}

NSRect ClipView_DocumentVisibleRect(void* ptr) {
	NSClipView* clipView = (NSClipView*)ptr;
	return [clipView documentVisibleRect];
}

bool ClipView_DrawsBackground(void* ptr) {
	NSClipView* clipView = (NSClipView*)ptr;
	return [clipView drawsBackground];
}

void ClipView_SetDrawsBackground(void* ptr, bool drawsBackground) {
	NSClipView* clipView = (NSClipView*)ptr;
	[clipView setDrawsBackground:drawsBackground];
}

void* ClipView_BackgroundColor(void* ptr) {
	NSClipView* clipView = (NSClipView*)ptr;
	return [clipView backgroundColor];
}

void ClipView_SetBackgroundColor(void* ptr, void* backgroundColor) {
	NSClipView* clipView = (NSClipView*)ptr;
	[clipView setBackgroundColor:(NSColor*)backgroundColor];
}

void ClipView_ScrollToPoint(void* ptr, NSPoint newOrigin) {
	NSClipView* clipView = (NSClipView*)ptr;
	[clipView scrollToPoint:newOrigin];
}

void ClipView_Autoscroll(void* ptr, void* event) {
	NSClipView* clipView = (NSClipView*)ptr;
	[clipView autoscroll:(NSEvent*)event];
}

NSRect ClipView_ConstrainBoundsRect(void* ptr, NSRect proposedBounds) {
	NSClipView* clipView = (NSClipView*)ptr;
	return [clipView constrainBoundsRect:proposedBounds];
}
