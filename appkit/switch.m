#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "switch.h"

long Switch_State(void* ptr) {
	NSSwitch* switch_ = (NSSwitch*)ptr;
	return [switch_ state];
}

void Switch_SetState(void* ptr, long state) {
	NSSwitch* switch_ = (NSSwitch*)ptr;
	[switch_ setState:state];
}

void* Switch_NewSwitch(NSRect frame) {
	NSSwitch* switch_ = [NSSwitch alloc];
	return [[switch_ initWithFrame:frame] autorelease];
}
