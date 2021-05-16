#import <Appkit/Appkit.h>
#import "open_panel.h"

void* C_OpenPanel_Alloc() {
    return [NSOpenPanel alloc];
}

void* C_NSOpenPanel_InitWithContentRect_StyleMask_Backing_Defer(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    NSOpenPanel* result = [nSOpenPanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag];
    return result;
}

void* C_NSOpenPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(void* ptr, CGRect contentRect, unsigned int style, unsigned int backingStoreType, bool flag, void* screen) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    NSOpenPanel* result = [nSOpenPanel initWithContentRect:contentRect styleMask:style backing:backingStoreType defer:flag screen:(NSScreen*)screen];
    return result;
}

void* C_NSOpenPanel_Init(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    NSOpenPanel* result = [nSOpenPanel init];
    return result;
}

void* C_NSOpenPanel_OpenPanel_() {
    NSOpenPanel* result = [NSOpenPanel openPanel];
    return result;
}

bool C_NSOpenPanel_CanChooseFiles(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result = [nSOpenPanel canChooseFiles];
    return result;
}

void C_NSOpenPanel_SetCanChooseFiles(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setCanChooseFiles:value];
}

bool C_NSOpenPanel_CanChooseDirectories(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result = [nSOpenPanel canChooseDirectories];
    return result;
}

void C_NSOpenPanel_SetCanChooseDirectories(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setCanChooseDirectories:value];
}

bool C_NSOpenPanel_ResolvesAliases(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result = [nSOpenPanel resolvesAliases];
    return result;
}

void C_NSOpenPanel_SetResolvesAliases(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setResolvesAliases:value];
}

bool C_NSOpenPanel_AllowsMultipleSelection(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result = [nSOpenPanel allowsMultipleSelection];
    return result;
}

void C_NSOpenPanel_SetAllowsMultipleSelection(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setAllowsMultipleSelection:value];
}

bool C_NSOpenPanel_IsAccessoryViewDisclosed(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result = [nSOpenPanel isAccessoryViewDisclosed];
    return result;
}

void C_NSOpenPanel_SetAccessoryViewDisclosed(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setAccessoryViewDisclosed:value];
}

Array C_NSOpenPanel_URLs(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    NSArray* result = [nSOpenPanel URLs];
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

bool C_NSOpenPanel_CanDownloadUbiquitousContents(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result = [nSOpenPanel canDownloadUbiquitousContents];
    return result;
}

void C_NSOpenPanel_SetCanDownloadUbiquitousContents(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setCanDownloadUbiquitousContents:value];
}

bool C_NSOpenPanel_CanResolveUbiquitousConflicts(void* ptr) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    BOOL result = [nSOpenPanel canResolveUbiquitousConflicts];
    return result;
}

void C_NSOpenPanel_SetCanResolveUbiquitousConflicts(void* ptr, bool value) {
    NSOpenPanel* nSOpenPanel = (NSOpenPanel*)ptr;
    [nSOpenPanel setCanResolveUbiquitousConflicts:value];
}
