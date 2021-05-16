#import <Appkit/Appkit.h>
#import "save_panel.h"

void* C_SavePanel_Alloc() {
    return [NSSavePanel alloc];
}

void* C_NSSavePanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSSavePanel* result = [nSSavePanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag];
    return result;
}

void* C_NSSavePanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSSavePanel* result = [nSSavePanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag screen:(NSScreen*)screen];
    return result;
}

void* C_NSSavePanel_Init(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSSavePanel* result = [nSSavePanel init];
    return result;
}

void* C_NSSavePanel_SavePanel_() {
    NSSavePanel* result = [NSSavePanel savePanel];
    return result;
}

int C_NSSavePanel_RunModal(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSModalResponse result = [nSSavePanel runModal];
    return result;
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
    NSURL* result = [nSSavePanel URL];
    return result;
}

void* C_NSSavePanel_Prompt(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSString* result = [nSSavePanel prompt];
    return result;
}

void C_NSSavePanel_SetPrompt(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setPrompt:(NSString*)value];
}

void* C_NSSavePanel_Message(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSString* result = [nSSavePanel message];
    return result;
}

void C_NSSavePanel_SetMessage(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setMessage:(NSString*)value];
}

void* C_NSSavePanel_NameFieldLabel(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSString* result = [nSSavePanel nameFieldLabel];
    return result;
}

void C_NSSavePanel_SetNameFieldLabel(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setNameFieldLabel:(NSString*)value];
}

void* C_NSSavePanel_NameFieldStringValue(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSString* result = [nSSavePanel nameFieldStringValue];
    return result;
}

void C_NSSavePanel_SetNameFieldStringValue(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setNameFieldStringValue:(NSString*)value];
}

void* C_NSSavePanel_DirectoryURL(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSURL* result = [nSSavePanel directoryURL];
    return result;
}

void C_NSSavePanel_SetDirectoryURL(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setDirectoryURL:(NSURL*)value];
}

void* C_NSSavePanel_AccessoryView(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSView* result = [nSSavePanel accessoryView];
    return result;
}

void C_NSSavePanel_SetAccessoryView(void* ptr, void* value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setAccessoryView:(NSView*)value];
}

bool C_NSSavePanel_ShowsTagField(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result = [nSSavePanel showsTagField];
    return result;
}

void C_NSSavePanel_SetShowsTagField(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setShowsTagField:value];
}

Array C_NSSavePanel_TagNames(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSArray* result = [nSSavePanel tagNames];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void C_NSSavePanel_SetTagNames(void* ptr, Array value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSString*)(NSString*)p];
    }
    [nSSavePanel setTagNames:objcValue];
}

bool C_NSSavePanel_CanCreateDirectories(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result = [nSSavePanel canCreateDirectories];
    return result;
}

void C_NSSavePanel_SetCanCreateDirectories(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setCanCreateDirectories:value];
}

bool C_NSSavePanel_CanSelectHiddenExtension(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result = [nSSavePanel canSelectHiddenExtension];
    return result;
}

void C_NSSavePanel_SetCanSelectHiddenExtension(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setCanSelectHiddenExtension:value];
}

bool C_NSSavePanel_ShowsHiddenFiles(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result = [nSSavePanel showsHiddenFiles];
    return result;
}

void C_NSSavePanel_SetShowsHiddenFiles(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setShowsHiddenFiles:value];
}

bool C_NSSavePanel_IsExtensionHidden(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result = [nSSavePanel isExtensionHidden];
    return result;
}

void C_NSSavePanel_SetExtensionHidden(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setExtensionHidden:value];
}

bool C_NSSavePanel_IsExpanded(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result = [nSSavePanel isExpanded];
    return result;
}

Array C_NSSavePanel_AllowedContentTypes(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSArray* result = [nSSavePanel allowedContentTypes];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void C_NSSavePanel_SetAllowedContentTypes(void* ptr, Array value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(UTType*)(UTType*)p];
    }
    [nSSavePanel setAllowedContentTypes:objcValue];
}

bool C_NSSavePanel_AllowsOtherFileTypes(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result = [nSSavePanel allowsOtherFileTypes];
    return result;
}

void C_NSSavePanel_SetAllowsOtherFileTypes(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setAllowsOtherFileTypes:value];
}

bool C_NSSavePanel_TreatsFilePackagesAsDirectories(void* ptr) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    BOOL result = [nSSavePanel treatsFilePackagesAsDirectories];
    return result;
}

void C_NSSavePanel_SetTreatsFilePackagesAsDirectories(void* ptr, bool value) {
    NSSavePanel* nSSavePanel = (NSSavePanel*)ptr;
    [nSSavePanel setTreatsFilePackagesAsDirectories:value];
}
