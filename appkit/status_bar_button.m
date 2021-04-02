#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "status_bar_button.h"

bool StatusBarButton_AppearsDisabled(void* ptr) {
	NSStatusBarButton* statusBarButton = (NSStatusBarButton*)ptr;
	return [statusBarButton appearsDisabled];
}

void StatusBarButton_SetAppearsDisabled(void* ptr, bool appearsDisabled) {
	NSStatusBarButton* statusBarButton = (NSStatusBarButton*)ptr;
	[statusBarButton setAppearsDisabled:appearsDisabled];
}
