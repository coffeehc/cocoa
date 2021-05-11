#import <Appkit/Appkit.h>
#import "tracking_area.h"

void* C_TrackingArea_Alloc() {
	return [NSTrackingArea alloc];
}

void* C_NSTrackingArea_Init(void* ptr) {
	NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
	NSTrackingArea* result = [nSTrackingArea init];
	return result;
}

void* C_NSTrackingArea_Owner(void* ptr) {
	NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
	id result = [nSTrackingArea owner];
	return result;
}

CGRect C_NSTrackingArea_Rect(void* ptr) {
	NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
	CGRect result = [nSTrackingArea rect];
	return result;
}
