#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "layout_manager.h"

bool LayoutManager_AllowsNonContiguousLayout(void* ptr) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	return [layoutManager allowsNonContiguousLayout];
}

void LayoutManager_SetAllowsNonContiguousLayout(void* ptr, bool allowsNonContiguousLayout) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	[layoutManager setAllowsNonContiguousLayout:allowsNonContiguousLayout];
}

bool LayoutManager_HasNonContiguousLayout(void* ptr) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	return [layoutManager hasNonContiguousLayout];
}

bool LayoutManager_ShowsInvisibleCharacters(void* ptr) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	return [layoutManager showsInvisibleCharacters];
}

void LayoutManager_SetShowsInvisibleCharacters(void* ptr, bool showsInvisibleCharacters) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	[layoutManager setShowsInvisibleCharacters:showsInvisibleCharacters];
}

bool LayoutManager_ShowsControlCharacters(void* ptr) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	return [layoutManager showsControlCharacters];
}

void LayoutManager_SetShowsControlCharacters(void* ptr, bool showsControlCharacters) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	[layoutManager setShowsControlCharacters:showsControlCharacters];
}

bool LayoutManager_UsesFontLeading(void* ptr) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	return [layoutManager usesFontLeading];
}

void LayoutManager_SetUsesFontLeading(void* ptr, bool usesFontLeading) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	[layoutManager setUsesFontLeading:usesFontLeading];
}

bool LayoutManager_BackgroundLayoutEnabled(void* ptr) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	return [layoutManager backgroundLayoutEnabled];
}

void LayoutManager_SetBackgroundLayoutEnabled(void* ptr, bool backgroundLayoutEnabled) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	[layoutManager setBackgroundLayoutEnabled:backgroundLayoutEnabled];
}

bool LayoutManager_LimitsLayoutForSuspiciousContents(void* ptr) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	return [layoutManager limitsLayoutForSuspiciousContents];
}

void LayoutManager_SetLimitsLayoutForSuspiciousContents(void* ptr, bool limitsLayoutForSuspiciousContents) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	[layoutManager setLimitsLayoutForSuspiciousContents:limitsLayoutForSuspiciousContents];
}

bool LayoutManager_UsesDefaultHyphenation(void* ptr) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	return [layoutManager usesDefaultHyphenation];
}

void LayoutManager_SetUsesDefaultHyphenation(void* ptr, bool usesDefaultHyphenation) {
	NSLayoutManager* layoutManager = (NSLayoutManager*)ptr;
	[layoutManager setUsesDefaultHyphenation:usesDefaultHyphenation];
}

void* LayoutManager_NewLayoutManager() {
	NSLayoutManager* layoutManager = [NSLayoutManager alloc];
	return [[layoutManager init] autorelease];
}
