#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "save_panel.h"

void* SavePanel_URL(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel URL];
}

const char* SavePanel_Prompt(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [[savePanel prompt] UTF8String];
}

void SavePanel_SetPrompt(void* ptr, const char* prompt) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setPrompt:[NSString stringWithUTF8String:prompt]];
}

const char* SavePanel_Message(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [[savePanel message] UTF8String];
}

void SavePanel_SetMessage(void* ptr, const char* message) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setMessage:[NSString stringWithUTF8String:message]];
}

const char* SavePanel_NameFieldLabel(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [[savePanel nameFieldLabel] UTF8String];
}

void SavePanel_SetNameFieldLabel(void* ptr, const char* nameFieldLabel) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setNameFieldLabel:[NSString stringWithUTF8String:nameFieldLabel]];
}

void* SavePanel_DirectoryURL(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel directoryURL];
}

void SavePanel_SetDirectoryURL(void* ptr, void* directoryURL) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setDirectoryURL:(NSURL*)directoryURL];
}

void* SavePanel_AccessoryView(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel accessoryView];
}

void SavePanel_SetAccessoryView(void* ptr, void* accessoryView) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setAccessoryView:(NSView*)accessoryView];
}

bool SavePanel_ShowsTagField(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel showsTagField];
}

void SavePanel_SetShowsTagField(void* ptr, bool showsTagField) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setShowsTagField:showsTagField];
}

Array SavePanel_TagNames(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	NSArray* ns_array = [savePanel tagNames];
	int count = [ns_array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = (void*)[[ns_array objectAtIndex:i] UTF8String];}
	Array array;
	array.data = data;
	array.len = count;
	return array;
}

void SavePanel_SetTagNames(void* ptr, Array tagNames) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
    NSMutableArray* objc_tagNames = [[NSMutableArray alloc] init];
    void** tagNamesData = (void**)tagNames.data;
    for (int i = 0; i < tagNames.len; i++) {
    	[objc_tagNames addObject:[NSString stringWithUTF8String:(const char*)tagNamesData[i]]];
    }
	[savePanel setTagNames:objc_tagNames];
}

bool SavePanel_CanCreateDirectories(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel canCreateDirectories];
}

void SavePanel_SetCanCreateDirectories(void* ptr, bool canCreateDirectories) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setCanCreateDirectories:canCreateDirectories];
}

bool SavePanel_CanSelectHiddenExtension(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel canSelectHiddenExtension];
}

void SavePanel_SetCanSelectHiddenExtension(void* ptr, bool canSelectHiddenExtension) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setCanSelectHiddenExtension:canSelectHiddenExtension];
}

bool SavePanel_ShowsHiddenFiles(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel showsHiddenFiles];
}

void SavePanel_SetShowsHiddenFiles(void* ptr, bool showsHiddenFiles) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setShowsHiddenFiles:showsHiddenFiles];
}

bool SavePanel_IsExtensionHidden(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel isExtensionHidden];
}

void SavePanel_SetExtensionHidden(void* ptr, bool extensionHidden) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setExtensionHidden:extensionHidden];
}

bool SavePanel_IsExpanded(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel isExpanded];
}

Array SavePanel_AllowedContentTypes(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	NSArray* ns_array = [savePanel allowedContentTypes];
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

void SavePanel_SetAllowedContentTypes(void* ptr, Array allowedContentTypes) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
    NSMutableArray* objc_allowedContentTypes = [[NSMutableArray alloc] init];
    void** allowedContentTypesData = (void**)allowedContentTypes.data;
    for (int i = 0; i < allowedContentTypes.len; i++) {
    	[objc_allowedContentTypes addObject:(UTType*)allowedContentTypesData[i]];
    }
	[savePanel setAllowedContentTypes:objc_allowedContentTypes];
}

bool SavePanel_AllowsOtherFileTypes(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel allowsOtherFileTypes];
}

void SavePanel_SetAllowsOtherFileTypes(void* ptr, bool allowsOtherFileTypes) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setAllowsOtherFileTypes:allowsOtherFileTypes];
}

bool SavePanel_TreatsFilePackagesAsDirectories(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel treatsFilePackagesAsDirectories];
}

void SavePanel_SetTreatsFilePackagesAsDirectories(void* ptr, bool treatsFilePackagesAsDirectories) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel setTreatsFilePackagesAsDirectories:treatsFilePackagesAsDirectories];
}

void* SavePanel_NewSavePanel() {
	return [NSSavePanel savePanel];
}

unsigned long SavePanel_RunModal(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	return [savePanel runModal];
}

void SavePanel_ValidateVisibleColumns(void* ptr) {
	NSSavePanel* savePanel = (NSSavePanel*)ptr;
	[savePanel validateVisibleColumns];
}

void SavePanel_BeginSheetModalForWindow(void* ptr, void* window, long handlerID) {
    NSSavePanel* savePanel = (NSSavePanel*)ptr;
    [savePanel beginSheetModalForWindow:(NSWindow*)window completionHandler:^(NSInteger result) {
        SavePanel_CallCompletionHandler(handlerID, result);
    }];
}

void SavePanel_BeginWithCompletionHandler(void* ptr, long handlerID) {
    NSSavePanel* savePanel = (NSSavePanel*)ptr;
    [savePanel beginWithCompletionHandler:^(NSInteger result) {
        SavePanel_CallCompletionHandler(handlerID, result);
    }];
}

