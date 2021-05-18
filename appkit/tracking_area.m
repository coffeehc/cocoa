#import <Appkit/Appkit.h>
#import "tracking_area.h"

void* C_TrackingArea_Alloc() {
    return [NSTrackingArea alloc];
}

void* C_NSTrackingArea_Init(void* ptr) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    NSTrackingArea* result_ = [nSTrackingArea init];
    return result_;
}

unsigned int C_NSTrackingArea_Options(void* ptr) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    NSTrackingAreaOptions result_ = [nSTrackingArea options];
    return result_;
}

void* C_NSTrackingArea_Owner(void* ptr) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    id result_ = [nSTrackingArea owner];
    return result_;
}

CGRect C_NSTrackingArea_Rect(void* ptr) {
    NSTrackingArea* nSTrackingArea = (NSTrackingArea*)ptr;
    NSRect result_ = [nSTrackingArea rect];
    return result_;
}
