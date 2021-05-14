#import <Appkit/Appkit.h>
#import "dock_tile.h"

void* C_DockTile_Alloc() {
	return [NSDockTile alloc];
}

void* C_NSDockTile_Init(void* ptr) {
	NSDockTile* nSDockTile = (NSDockTile*)ptr;
	NSDockTile* result = [nSDockTile init];
	return result;
}

void C_NSDockTile_Display(void* ptr) {
	NSDockTile* nSDockTile = (NSDockTile*)ptr;
	[nSDockTile display];
}

void* C_NSDockTile_ContentView(void* ptr) {
	NSDockTile* nSDockTile = (NSDockTile*)ptr;
	NSView* result = [nSDockTile contentView];
	return result;
}

void C_NSDockTile_SetContentView(void* ptr, void* value) {
	NSDockTile* nSDockTile = (NSDockTile*)ptr;
	[nSDockTile setContentView:(NSView*)value];
}

CGSize C_NSDockTile_Size(void* ptr) {
	NSDockTile* nSDockTile = (NSDockTile*)ptr;
	CGSize result = [nSDockTile size];
	return result;
}

void* C_NSDockTile_Owner(void* ptr) {
	NSDockTile* nSDockTile = (NSDockTile*)ptr;
	id result = [nSDockTile owner];
	return result;
}

bool C_NSDockTile_ShowsApplicationBadge(void* ptr) {
	NSDockTile* nSDockTile = (NSDockTile*)ptr;
	bool result = [nSDockTile showsApplicationBadge];
	return result;
}

void C_NSDockTile_SetShowsApplicationBadge(void* ptr, bool value) {
	NSDockTile* nSDockTile = (NSDockTile*)ptr;
	[nSDockTile setShowsApplicationBadge:value];
}

void* C_NSDockTile_BadgeLabel(void* ptr) {
	NSDockTile* nSDockTile = (NSDockTile*)ptr;
	NSString* result = [nSDockTile badgeLabel];
	return result;
}

void C_NSDockTile_SetBadgeLabel(void* ptr, void* value) {
	NSDockTile* nSDockTile = (NSDockTile*)ptr;
	[nSDockTile setBadgeLabel:(NSString*)value];
}
