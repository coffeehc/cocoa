#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TrackingArea_Alloc();

void* C_NSTrackingArea_InitWithRect_Options_Owner_UserInfo(void* ptr, CGRect rect, unsigned int options, void* owner, Dictionary userInfo);
void* C_NSTrackingArea_Init(void* ptr);
unsigned int C_NSTrackingArea_Options(void* ptr);
void* C_NSTrackingArea_Owner(void* ptr);
CGRect C_NSTrackingArea_Rect(void* ptr);
Dictionary C_NSTrackingArea_UserInfo(void* ptr);
