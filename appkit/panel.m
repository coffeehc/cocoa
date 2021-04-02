#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "panel.h"

bool Panel_IsFloatingPanel(void* ptr) {
	NSPanel* panel = (NSPanel*)ptr;
	return [panel isFloatingPanel];
}

void Panel_SetFloatingPanel(void* ptr, bool floatingPanel) {
	NSPanel* panel = (NSPanel*)ptr;
	[panel setFloatingPanel:floatingPanel];
}

bool Panel_BecomesKeyOnlyIfNeeded(void* ptr) {
	NSPanel* panel = (NSPanel*)ptr;
	return [panel becomesKeyOnlyIfNeeded];
}

void Panel_SetBecomesKeyOnlyIfNeeded(void* ptr, bool becomesKeyOnlyIfNeeded) {
	NSPanel* panel = (NSPanel*)ptr;
	[panel setBecomesKeyOnlyIfNeeded:becomesKeyOnlyIfNeeded];
}

bool Panel_WorksWhenModal(void* ptr) {
	NSPanel* panel = (NSPanel*)ptr;
	return [panel worksWhenModal];
}

void Panel_SetWorksWhenModal(void* ptr, bool worksWhenModal) {
	NSPanel* panel = (NSPanel*)ptr;
	[panel setWorksWhenModal:worksWhenModal];
}
