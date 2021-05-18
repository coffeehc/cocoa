#import <Appkit/Appkit.h>
#import "dock_tile.h"

void* C_DockTile_Alloc() {
    return [NSDockTile alloc];
}

void* C_NSDockTile_Init(void* ptr) {
    NSDockTile* nSDockTile = (NSDockTile*)ptr;
    NSDockTile* result_ = [nSDockTile init];
    return result_;
}

void C_NSDockTile_Display(void* ptr) {
    NSDockTile* nSDockTile = (NSDockTile*)ptr;
    [nSDockTile display];
}

void* C_NSDockTile_ContentView(void* ptr) {
    NSDockTile* nSDockTile = (NSDockTile*)ptr;
    NSView* result_ = [nSDockTile contentView];
    return result_;
}

void C_NSDockTile_SetContentView(void* ptr, void* value) {
    NSDockTile* nSDockTile = (NSDockTile*)ptr;
    [nSDockTile setContentView:(NSView*)value];
}

CGSize C_NSDockTile_Size(void* ptr) {
    NSDockTile* nSDockTile = (NSDockTile*)ptr;
    NSSize result_ = [nSDockTile size];
    return result_;
}

void* C_NSDockTile_Owner(void* ptr) {
    NSDockTile* nSDockTile = (NSDockTile*)ptr;
    id result_ = [nSDockTile owner];
    return result_;
}

bool C_NSDockTile_ShowsApplicationBadge(void* ptr) {
    NSDockTile* nSDockTile = (NSDockTile*)ptr;
    BOOL result_ = [nSDockTile showsApplicationBadge];
    return result_;
}

void C_NSDockTile_SetShowsApplicationBadge(void* ptr, bool value) {
    NSDockTile* nSDockTile = (NSDockTile*)ptr;
    [nSDockTile setShowsApplicationBadge:value];
}

void* C_NSDockTile_BadgeLabel(void* ptr) {
    NSDockTile* nSDockTile = (NSDockTile*)ptr;
    NSString* result_ = [nSDockTile badgeLabel];
    return result_;
}

void C_NSDockTile_SetBadgeLabel(void* ptr, void* value) {
    NSDockTile* nSDockTile = (NSDockTile*)ptr;
    [nSDockTile setBadgeLabel:(NSString*)value];
}
