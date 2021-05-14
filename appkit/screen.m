#import <Appkit/Appkit.h>
#import "screen.h"

void* C_Screen_Alloc() {
	return [NSScreen alloc];
}

void* C_NSScreen_Init(void* ptr) {
	NSScreen* nSScreen = (NSScreen*)ptr;
	NSScreen* result = [nSScreen init];
	return result;
}

CGRect C_NSScreen_ConvertRectFromBacking(void* ptr, CGRect rect) {
	NSScreen* nSScreen = (NSScreen*)ptr;
	CGRect result = [nSScreen convertRectFromBacking:rect];
	return result;
}

CGRect C_NSScreen_ConvertRectToBacking(void* ptr, CGRect rect) {
	NSScreen* nSScreen = (NSScreen*)ptr;
	CGRect result = [nSScreen convertRectToBacking:rect];
	return result;
}

CGRect C_NSScreen_Frame(void* ptr) {
	NSScreen* nSScreen = (NSScreen*)ptr;
	CGRect result = [nSScreen frame];
	return result;
}

CGRect C_NSScreen_VisibleFrame(void* ptr) {
	NSScreen* nSScreen = (NSScreen*)ptr;
	CGRect result = [nSScreen visibleFrame];
	return result;
}

void* C_NSScreen_ColorSpace(void* ptr) {
	NSScreen* nSScreen = (NSScreen*)ptr;
	NSColorSpace* result = [nSScreen colorSpace];
	return result;
}

double C_NSScreen_BackingScaleFactor(void* ptr) {
	NSScreen* nSScreen = (NSScreen*)ptr;
	double result = [nSScreen backingScaleFactor];
	return result;
}

double C_NSScreen_MaximumPotentialExtendedDynamicRangeColorComponentValue(void* ptr) {
	NSScreen* nSScreen = (NSScreen*)ptr;
	double result = [nSScreen maximumPotentialExtendedDynamicRangeColorComponentValue];
	return result;
}

double C_NSScreen_MaximumExtendedDynamicRangeColorComponentValue(void* ptr) {
	NSScreen* nSScreen = (NSScreen*)ptr;
	double result = [nSScreen maximumExtendedDynamicRangeColorComponentValue];
	return result;
}

double C_NSScreen_MaximumReferenceExtendedDynamicRangeColorComponentValue(void* ptr) {
	NSScreen* nSScreen = (NSScreen*)ptr;
	double result = [nSScreen maximumReferenceExtendedDynamicRangeColorComponentValue];
	return result;
}

void* C_NSScreen_LocalizedName(void* ptr) {
	NSScreen* nSScreen = (NSScreen*)ptr;
	NSString* result = [nSScreen localizedName];
	return result;
}
