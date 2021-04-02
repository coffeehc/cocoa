#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "open_panel.h"

bool OpenPanel_CanChooseFiles(void* ptr) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	return [openPanel canChooseFiles];
}

void OpenPanel_SetCanChooseFiles(void* ptr, bool canChooseFiles) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	[openPanel setCanChooseFiles:canChooseFiles];
}

bool OpenPanel_CanChooseDirectories(void* ptr) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	return [openPanel canChooseDirectories];
}

void OpenPanel_SetCanChooseDirectories(void* ptr, bool canChooseDirectories) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	[openPanel setCanChooseDirectories:canChooseDirectories];
}

bool OpenPanel_ResolvesAliases(void* ptr) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	return [openPanel resolvesAliases];
}

void OpenPanel_SetResolvesAliases(void* ptr, bool resolvesAliases) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	[openPanel setResolvesAliases:resolvesAliases];
}

bool OpenPanel_AllowsMultipleSelection(void* ptr) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	return [openPanel allowsMultipleSelection];
}

void OpenPanel_SetAllowsMultipleSelection(void* ptr, bool allowsMultipleSelection) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	[openPanel setAllowsMultipleSelection:allowsMultipleSelection];
}

bool OpenPanel_IsAccessoryViewDisclosed(void* ptr) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	return [openPanel isAccessoryViewDisclosed];
}

void OpenPanel_SetAccessoryViewDisclosed(void* ptr, bool accessoryViewDisclosed) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	[openPanel setAccessoryViewDisclosed:accessoryViewDisclosed];
}

Array OpenPanel_URLs(void* ptr) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	NSArray* ns_array = [openPanel URLs];
	int count = [ns_array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = [ns_array objectAtIndex:i];
	}
	Array array;
	array.data = data;
	array.len = count;
	return array;
}

bool OpenPanel_CanDownloadUbiquitousContents(void* ptr) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	return [openPanel canDownloadUbiquitousContents];
}

void OpenPanel_SetCanDownloadUbiquitousContents(void* ptr, bool canDownloadUbiquitousContents) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	[openPanel setCanDownloadUbiquitousContents:canDownloadUbiquitousContents];
}

bool OpenPanel_CanResolveUbiquitousConflicts(void* ptr) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	return [openPanel canResolveUbiquitousConflicts];
}

void OpenPanel_SetCanResolveUbiquitousConflicts(void* ptr, bool canResolveUbiquitousConflicts) {
	NSOpenPanel* openPanel = (NSOpenPanel*)ptr;
	[openPanel setCanResolveUbiquitousConflicts:canResolveUbiquitousConflicts];
}

void* OpenPanel_NewOpenPanel() {
	return [NSOpenPanel openPanel];
}
