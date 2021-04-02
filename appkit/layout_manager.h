#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

bool LayoutManager_AllowsNonContiguousLayout(void* ptr);
void LayoutManager_SetAllowsNonContiguousLayout(void* ptr, bool allowsNonContiguousLayout);
bool LayoutManager_HasNonContiguousLayout(void* ptr);
bool LayoutManager_ShowsInvisibleCharacters(void* ptr);
void LayoutManager_SetShowsInvisibleCharacters(void* ptr, bool showsInvisibleCharacters);
bool LayoutManager_ShowsControlCharacters(void* ptr);
void LayoutManager_SetShowsControlCharacters(void* ptr, bool showsControlCharacters);
bool LayoutManager_UsesFontLeading(void* ptr);
void LayoutManager_SetUsesFontLeading(void* ptr, bool usesFontLeading);
bool LayoutManager_BackgroundLayoutEnabled(void* ptr);
void LayoutManager_SetBackgroundLayoutEnabled(void* ptr, bool backgroundLayoutEnabled);
bool LayoutManager_LimitsLayoutForSuspiciousContents(void* ptr);
void LayoutManager_SetLimitsLayoutForSuspiciousContents(void* ptr, bool limitsLayoutForSuspiciousContents);
bool LayoutManager_UsesDefaultHyphenation(void* ptr);
void LayoutManager_SetUsesDefaultHyphenation(void* ptr, bool usesDefaultHyphenation);

void* LayoutManager_NewLayoutManager();
