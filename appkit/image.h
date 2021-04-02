#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

NSSize Image_Size(void* ptr);
void Image_SetSize(void* ptr, NSSize size);
bool Image_IsTemplate(void* ptr);
void Image_SetTemplate(void* ptr, bool template);
const char* Image_AccessibilityDescription(void* ptr);
void Image_SetAccessibilityDescription(void* ptr, const char* accessibilityDescription);
bool Image_MatchesOnlyOnBestFittingAxis(void* ptr);
void Image_SetMatchesOnlyOnBestFittingAxis(void* ptr, bool matchesOnlyOnBestFittingAxis);
Array Image_ImageTypes();
Array Image_ImageUnfilteredTypes();

void* Image_ImageNamed(const char* name);
const char* Image_Name(void* ptr);
void Image_SetName(void* ptr, const char* name);
void* Image_ImageWithSystemSymbolName(const char* symbolName, const char* description);
void Image_ImageWithSymbolConfiguration(void* ptr, void* configuration);
void* Image_InitByReferencingFile(void* ptr, const char* fileName);
void* Image_InitByReferencingURL(void* ptr, void* url);
void* Image_InitWithContentsOfFile(void* ptr, const char* fileName);
void* Image_InitWithContentsOfURL(void* ptr, void* fileName);
void* Image_InitWithData(void* ptr, Array data);
void* Image_InitWithDataIgnoringOrientation(void* ptr, Array data);
void* Image_InitWithSize(void* ptr, NSSize szie);
