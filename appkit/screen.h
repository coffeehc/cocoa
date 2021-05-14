#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Screen_Alloc();

void* C_NSScreen_Init(void* ptr);
CGRect C_NSScreen_ConvertRectFromBacking(void* ptr, CGRect rect);
CGRect C_NSScreen_ConvertRectToBacking(void* ptr, CGRect rect);
CGRect C_NSScreen_Frame(void* ptr);
CGRect C_NSScreen_VisibleFrame(void* ptr);
void* C_NSScreen_ColorSpace(void* ptr);
double C_NSScreen_BackingScaleFactor(void* ptr);
double C_NSScreen_MaximumPotentialExtendedDynamicRangeColorComponentValue(void* ptr);
double C_NSScreen_MaximumExtendedDynamicRangeColorComponentValue(void* ptr);
double C_NSScreen_MaximumReferenceExtendedDynamicRangeColorComponentValue(void* ptr);
void* C_NSScreen_LocalizedName(void* ptr);
