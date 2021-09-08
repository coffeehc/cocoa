#import <Appkit/Appkit.h>
#import "popover.h"

void* C_Popover_Alloc() {
    return [NSPopover alloc];
}

void* C_NSPopover_Init(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    NSPopover* result_ = [nSPopover init];
    return result_;
}

void* C_NSPopover_InitWithCoder(void* ptr, void* coder) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    NSPopover* result_ = [nSPopover initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSPopover_AllocPopover() {
    NSPopover* result_ = [NSPopover alloc];
    return result_;
}

void* C_NSPopover_NewPopover() {
    NSPopover* result_ = [NSPopover new];
    return result_;
}

void* C_NSPopover_Autorelease(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    NSPopover* result_ = [nSPopover autorelease];
    return result_;
}

void* C_NSPopover_Retain(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    NSPopover* result_ = [nSPopover retain];
    return result_;
}

void C_NSPopover_ShowRelativeToRect_OfView_PreferredEdge(void* ptr, CGRect positioningRect, void* positioningView, unsigned int preferredEdge) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    [nSPopover showRelativeToRect:positioningRect ofView:(NSView*)positioningView preferredEdge:preferredEdge];
}

void C_NSPopover_PerformClose(void* ptr, void* sender) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    [nSPopover performClose:(id)sender];
}

void C_NSPopover_Close(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    [nSPopover close];
}

int C_NSPopover_Behavior(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    NSPopoverBehavior result_ = [nSPopover behavior];
    return result_;
}

void C_NSPopover_SetBehavior(void* ptr, int value) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    [nSPopover setBehavior:value];
}

CGRect C_NSPopover_PositioningRect(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    NSRect result_ = [nSPopover positioningRect];
    return result_;
}

void C_NSPopover_SetPositioningRect(void* ptr, CGRect value) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    [nSPopover setPositioningRect:value];
}

void* C_NSPopover_Appearance(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    NSAppearance* result_ = [nSPopover appearance];
    return result_;
}

void C_NSPopover_SetAppearance(void* ptr, void* value) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    [nSPopover setAppearance:(NSAppearance*)value];
}

void* C_NSPopover_EffectiveAppearance(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    NSAppearance* result_ = [nSPopover effectiveAppearance];
    return result_;
}

bool C_NSPopover_Animates(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    BOOL result_ = [nSPopover animates];
    return result_;
}

void C_NSPopover_SetAnimates(void* ptr, bool value) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    [nSPopover setAnimates:value];
}

CGSize C_NSPopover_ContentSize(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    NSSize result_ = [nSPopover contentSize];
    return result_;
}

void C_NSPopover_SetContentSize(void* ptr, CGSize value) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    [nSPopover setContentSize:value];
}

bool C_NSPopover_IsShown(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    BOOL result_ = [nSPopover isShown];
    return result_;
}

bool C_NSPopover_IsDetached(void* ptr) {
    NSPopover* nSPopover = (NSPopover*)ptr;
    BOOL result_ = [nSPopover isDetached];
    return result_;
}
