#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Screen_Alloc();

void* C_NSScreen_Init(void* ptr);
bool C_NSScreen_CanRepresentDisplayGamut(void* ptr, int displayGamut);
CGRect C_NSScreen_ConvertRectFromBacking(void* ptr, CGRect rect);
CGRect C_NSScreen_ConvertRectToBacking(void* ptr, CGRect rect);
void* C_NSScreen_MainScreen();
void* C_NSScreen_DeepestScreen();
Array C_NSScreen_Screens();
int32_t C_NSScreen_Depth(void* ptr);
CGRect C_NSScreen_Frame(void* ptr);
Dictionary C_NSScreen_DeviceDescription(void* ptr);
CGRect C_NSScreen_VisibleFrame(void* ptr);
void* C_NSScreen_ColorSpace(void* ptr);
bool C_NSScreen_ScreensHaveSeparateSpaces();
double C_NSScreen_BackingScaleFactor(void* ptr);
double C_NSScreen_MaximumPotentialExtendedDynamicRangeColorComponentValue(void* ptr);
double C_NSScreen_MaximumExtendedDynamicRangeColorComponentValue(void* ptr);
double C_NSScreen_MaximumReferenceExtendedDynamicRangeColorComponentValue(void* ptr);
void* C_NSScreen_LocalizedName(void* ptr);
