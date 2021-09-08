#import <Appkit/Appkit.h>
#import "save_panel.h"

void* C_SavePanel_Alloc() {
    return [NSSavePanel alloc];
}

void* C_NSSavePanel_SavePanel_WindowWithContentViewController(void* contentViewController) {
    NSSavePanel* result_ = [NSSavePanel windowWithContentViewController:(NSViewController*)contentViewController];
    return result_;
}

void* C_NSSavePanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSSavePanel* result_ = [nSSavePanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag];
    return result_;
}

void* C_NSSavePanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSSavePanel* result_ = [nSSavePanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag screen:(NSScreen*)screen];
    return result_;
}

void* C_NSSavePanel_Init(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSSavePanel* result_ = [nSSavePanel init];
    return result_;
}

void* C_NSSavePanel_AllocSavePanel() {
    NSSavePanel* result_ = [NSSavePanel alloc];
    return result_;
}

void* C_NSSavePanel_NewSavePanel() {
    NSSavePanel* result_ = [NSSavePanel new];
    return result_;
}

void* C_NSSavePanel_Autorelease(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSSavePanel* result_ = [nSSavePanel autorelease];
    return result_;
}

void* C_NSSavePanel_Retain(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSSavePanel* result_ = [nSSavePanel retain];
    return result_;
}

void* C_NSSavePanel_SavePanel_() {
    NSSavePanel* result_ = [NSSavePanel savePanel];
    return result_;
}

int C_NSSavePanel_RunModal(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSModalResponse result_ = [nSSavePanel runModal];
    return result_;
}

void C_NSSavePanel_ValidateVisibleColumns(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel validateVisibleColumns];
}

void C_NSSavePanel_Ok(void* ptr, void* sender) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel ok:(id)sender];
}

void C_NSSavePanel_Cancel(void* ptr, void* sender) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel cancel:(id)sender];
}

void* C_NSSavePanel_URL(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSURL* result_ = [nSSavePanel URL];
    return result_;
}

void* C_NSSavePanel_Prompt(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSString* result_ = [nSSavePanel prompt];
    return result_;
}

void C_NSSavePanel_SetPrompt(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setPrompt:(NSString*)value];
}

void* C_NSSavePanel_Message(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSString* result_ = [nSSavePanel message];
    return result_;
}

void C_NSSavePanel_SetMessage(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setMessage:(NSString*)value];
}

void* C_NSSavePanel_NameFieldLabel(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSString* result_ = [nSSavePanel nameFieldLabel];
    return result_;
}

void C_NSSavePanel_SetNameFieldLabel(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setNameFieldLabel:(NSString*)value];
}

void* C_NSSavePanel_NameFieldStringValue(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSString* result_ = [nSSavePanel nameFieldStringValue];
    return result_;
}

void C_NSSavePanel_SetNameFieldStringValue(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setNameFieldStringValue:(NSString*)value];
}

void* C_NSSavePanel_DirectoryURL(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSURL* result_ = [nSSavePanel directoryURL];
    return result_;
}

void C_NSSavePanel_SetDirectoryURL(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setDirectoryURL:(NSURL*)value];
}

void* C_NSSavePanel_AccessoryView(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSView* result_ = [nSSavePanel accessoryView];
    return result_;
}

void C_NSSavePanel_SetAccessoryView(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setAccessoryView:(NSView*)value];
}

bool C_NSSavePanel_ShowsTagField(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result_ = [nSSavePanel showsTagField];
    return result_;
}

void C_NSSavePanel_SetShowsTagField(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setShowsTagField:value];
}

Array C_NSSavePanel_TagNames(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSArray* result_ = [nSSavePanel tagNames];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSSavePanel_SetTagNames(void* ptr, Array value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSString*)(NSString*)p];
    	}
    }
    [nSSavePanel setTagNames:objcValue];
}

bool C_NSSavePanel_CanCreateDirectories(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result_ = [nSSavePanel canCreateDirectories];
    return result_;
}

void C_NSSavePanel_SetCanCreateDirectories(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setCanCreateDirectories:value];
}

bool C_NSSavePanel_CanSelectHiddenExtension(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result_ = [nSSavePanel canSelectHiddenExtension];
    return result_;
}

void C_NSSavePanel_SetCanSelectHiddenExtension(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setCanSelectHiddenExtension:value];
}

bool C_NSSavePanel_ShowsHiddenFiles(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result_ = [nSSavePanel showsHiddenFiles];
    return result_;
}

void C_NSSavePanel_SetShowsHiddenFiles(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setShowsHiddenFiles:value];
}

bool C_NSSavePanel_IsExtensionHidden(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result_ = [nSSavePanel isExtensionHidden];
    return result_;
}

void C_NSSavePanel_SetExtensionHidden(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setExtensionHidden:value];
}

bool C_NSSavePanel_IsExpanded(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result_ = [nSSavePanel isExpanded];
    return result_;
}

Array C_NSSavePanel_AllowedContentTypes(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSArray* result_ = [nSSavePanel allowedContentTypes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSSavePanel_SetAllowedContentTypes(void* ptr, Array value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(UTType*)(UTType*)p];
    	}
    }
    [nSSavePanel setAllowedContentTypes:objcValue];
}

bool C_NSSavePanel_AllowsOtherFileTypes(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result_ = [nSSavePanel allowsOtherFileTypes];
    return result_;
}

void C_NSSavePanel_SetAllowsOtherFileTypes(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setAllowsOtherFileTypes:value];
}

bool C_NSSavePanel_TreatsFilePackagesAsDirectories(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result_ = [nSSavePanel treatsFilePackagesAsDirectories];
    return result_;
}

void C_NSSavePanel_SetTreatsFilePackagesAsDirectories(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setTreatsFilePackagesAsDirectories:value];
}
