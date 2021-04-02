#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "status_bar.h"

bool StatusBar_IsVertical(void* ptr) {
	NSStatusBar* statusBar = (NSStatusBar*)ptr;
	return [statusBar isVertical];
}

double StatusBar_Thickness(void* ptr) {
	NSStatusBar* statusBar = (NSStatusBar*)ptr;
	return [statusBar thickness];
}

void* StatusBar_SystemStatusBar() {
	return [NSStatusBar systemStatusBar];
}

void* StatusBar_StatusItemWithLength(void* ptr, double length) {
	NSStatusBar* statusBar = (NSStatusBar*)ptr;
	return [statusBar statusItemWithLength:length];
}

void StatusBar_RemoveStatusItem(void* ptr, void* item) {
	NSStatusBar* statusBar = (NSStatusBar*)ptr;
	[statusBar removeStatusItem:(NSStatusItem*)item];
}
