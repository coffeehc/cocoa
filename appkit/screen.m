#import "screen.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSScreen.h>

void* C_Screen_Alloc() {
    return [NSScreen alloc];
}

void* C_NSScreen_AllocScreen() {
    NSScreen* result_ = [NSScreen alloc];
    return result_;
}

void* C_NSScreen_Init(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    NSScreen* result_ = [nSScreen init];
    return result_;
}

void* C_NSScreen_NewScreen() {
    NSScreen* result_ = [NSScreen new];
    return result_;
}

void* C_NSScreen_Autorelease(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    NSScreen* result_ = [nSScreen autorelease];
    return result_;
}

void* C_NSScreen_Retain(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    NSScreen* result_ = [nSScreen retain];
    return result_;
}

bool C_NSScreen_CanRepresentDisplayGamut(void* ptr, int displayGamut) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    BOOL result_ = [nSScreen canRepresentDisplayGamut:displayGamut];
    return result_;
}

CGRect C_NSScreen_ConvertRectFromBacking(void* ptr, CGRect rect) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    NSRect result_ = [nSScreen convertRectFromBacking:rect];
    return result_;
}

CGRect C_NSScreen_ConvertRectToBacking(void* ptr, CGRect rect) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    NSRect result_ = [nSScreen convertRectToBacking:rect];
    return result_;
}

void* C_NSScreen_MainScreen() {
    NSScreen* result_ = [NSScreen mainScreen];
    return result_;
}

void* C_NSScreen_DeepestScreen() {
    NSScreen* result_ = [NSScreen deepestScreen];
    return result_;
}

Array C_NSScreen_Screens() {
    NSArray* result_ = [NSScreen screens];
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

int32_t C_NSScreen_Depth(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    NSWindowDepth result_ = [nSScreen depth];
    return result_;
}

CGRect C_NSScreen_Frame(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    NSRect result_ = [nSScreen frame];
    return result_;
}

Dictionary C_NSScreen_DeviceDescription(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    NSDictionary* result_ = [nSScreen deviceDescription];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSDeviceDescriptionKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

CGRect C_NSScreen_VisibleFrame(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    NSRect result_ = [nSScreen visibleFrame];
    return result_;
}

void* C_NSScreen_ColorSpace(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    NSColorSpace* result_ = [nSScreen colorSpace];
    return result_;
}

bool C_NSScreen_ScreensHaveSeparateSpaces() {
    BOOL result_ = [NSScreen screensHaveSeparateSpaces];
    return result_;
}

double C_NSScreen_BackingScaleFactor(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    CGFloat result_ = [nSScreen backingScaleFactor];
    return result_;
}

double C_NSScreen_MaximumPotentialExtendedDynamicRangeColorComponentValue(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    CGFloat result_ = [nSScreen maximumPotentialExtendedDynamicRangeColorComponentValue];
    return result_;
}

double C_NSScreen_MaximumExtendedDynamicRangeColorComponentValue(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    CGFloat result_ = [nSScreen maximumExtendedDynamicRangeColorComponentValue];
    return result_;
}

double C_NSScreen_MaximumReferenceExtendedDynamicRangeColorComponentValue(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    CGFloat result_ = [nSScreen maximumReferenceExtendedDynamicRangeColorComponentValue];
    return result_;
}

void* C_NSScreen_LocalizedName(void* ptr) {
    NSScreen* nSScreen = (NSScreen*)ptr;
    NSString* result_ = [nSScreen localizedName];
    return result_;
}
