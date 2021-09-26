#import "event.h"
#import <AppKit/NSEvent.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_Event_Alloc() {
    return [NSEvent alloc];
}

void* C_NSEvent_AllocEvent() {
    NSEvent* result_ = [NSEvent alloc];
    return result_;
}

void* C_NSEvent_Init(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEvent* result_ = [nSEvent init];
    return result_;
}

void* C_NSEvent_NewEvent() {
    NSEvent* result_ = [NSEvent new];
    return result_;
}

void* C_NSEvent_Autorelease(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEvent* result_ = [nSEvent autorelease];
    return result_;
}

void* C_NSEvent_Retain(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEvent* result_ = [nSEvent retain];
    return result_;
}

void* C_NSEvent_MouseEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_EventNumber_ClickCount_Pressure(unsigned int _type, CGPoint location, unsigned int flags, double time, int wNum, void* unusedPassNil, int eNum, int cNum, float pressure) {
    NSEvent* result_ = [NSEvent mouseEventWithType:_type location:location modifierFlags:flags timestamp:time windowNumber:wNum context:(NSGraphicsContext*)unusedPassNil eventNumber:eNum clickCount:cNum pressure:pressure];
    return result_;
}

void* C_NSEvent_EventWithCGEvent(void* cgEvent) {
    NSEvent* result_ = [NSEvent eventWithCGEvent:(CGEventRef)cgEvent];
    return result_;
}

void C_NSEvent_StartPeriodicEventsAfterDelay_WithPeriod(double delay, double period) {
    [NSEvent startPeriodicEventsAfterDelay:delay withPeriod:period];
}

void C_NSEvent_StopPeriodicEvents() {
    [NSEvent stopPeriodicEvents];
}

void* C_NSEvent_TouchesMatchingPhase_InView(void* ptr, unsigned int phase, void* view) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSSet* result_ = [nSEvent touchesMatchingPhase:phase inView:(NSView*)view];
    return result_;
}

void* C_NSEvent_AllTouches(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSSet* result_ = [nSEvent allTouches];
    return result_;
}

void* C_NSEvent_TouchesForView(void* ptr, void* view) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSSet* result_ = [nSEvent touchesForView:(NSView*)view];
    return result_;
}

Array C_NSEvent_CoalescedTouchesForTouch(void* ptr, void* touch) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSArray* result_ = [nSEvent coalescedTouchesForTouch:(NSTouch*)touch];
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

void C_NSEvent_Event_RemoveMonitor(void* eventMonitor) {
    [NSEvent removeMonitor:(id)eventMonitor];
}

void* C_NSEvent_CharactersByApplyingModifiers(void* ptr, unsigned int modifiers) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSString* result_ = [nSEvent charactersByApplyingModifiers:modifiers];
    return result_;
}

unsigned int C_NSEvent_Type(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEventType result_ = [nSEvent type];
    return result_;
}

unsigned int C_NSEvent_ModifierFlags(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEventModifierFlags result_ = [nSEvent modifierFlags];
    return result_;
}

CGPoint C_NSEvent_LocationInWindow(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSPoint result_ = [nSEvent locationInWindow];
    return result_;
}

double C_NSEvent_Timestamp(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSTimeInterval result_ = [nSEvent timestamp];
    return result_;
}

void* C_NSEvent_Window(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSWindow* result_ = [nSEvent window];
    return result_;
}

int C_NSEvent_WindowNumber(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result_ = [nSEvent windowNumber];
    return result_;
}

void* C_NSEvent_CGEvent(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGEventRef result_ = [nSEvent CGEvent];
    return result_;
}

double C_NSEvent_Event_KeyRepeatDelay() {
    NSTimeInterval result_ = [NSEvent keyRepeatDelay];
    return result_;
}

double C_NSEvent_Event_KeyRepeatInterval() {
    NSTimeInterval result_ = [NSEvent keyRepeatInterval];
    return result_;
}

void* C_NSEvent_Characters(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSString* result_ = [nSEvent characters];
    return result_;
}

void* C_NSEvent_CharactersIgnoringModifiers(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSString* result_ = [nSEvent charactersIgnoringModifiers];
    return result_;
}

bool C_NSEvent_IsARepeat(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    BOOL result_ = [nSEvent isARepeat];
    return result_;
}

unsigned int C_NSEvent_Event_PressedMouseButtons() {
    NSUInteger result_ = [NSEvent pressedMouseButtons];
    return result_;
}

double C_NSEvent_Event_DoubleClickInterval() {
    NSTimeInterval result_ = [NSEvent doubleClickInterval];
    return result_;
}

CGPoint C_NSEvent_Event_MouseLocation() {
    NSPoint result_ = [NSEvent mouseLocation];
    return result_;
}

int C_NSEvent_ButtonNumber(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result_ = [nSEvent buttonNumber];
    return result_;
}

int C_NSEvent_ClickCount(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result_ = [nSEvent clickCount];
    return result_;
}

int C_NSEvent_EventNumber(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result_ = [nSEvent eventNumber];
    return result_;
}

int C_NSEvent_TrackingNumber(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result_ = [nSEvent trackingNumber];
    return result_;
}

void* C_NSEvent_TrackingArea(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSTrackingArea* result_ = [nSEvent trackingArea];
    return result_;
}

int C_NSEvent_Data1(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result_ = [nSEvent data1];
    return result_;
}

int C_NSEvent_Data2(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result_ = [nSEvent data2];
    return result_;
}

double C_NSEvent_DeltaX(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result_ = [nSEvent deltaX];
    return result_;
}

double C_NSEvent_DeltaY(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result_ = [nSEvent deltaY];
    return result_;
}

double C_NSEvent_DeltaZ(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result_ = [nSEvent deltaZ];
    return result_;
}

float C_NSEvent_Pressure(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    float result_ = [nSEvent pressure];
    return result_;
}

int C_NSEvent_Stage(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result_ = [nSEvent stage];
    return result_;
}

double C_NSEvent_StageTransition(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result_ = [nSEvent stageTransition];
    return result_;
}

int C_NSEvent_PressureBehavior(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSPressureBehavior result_ = [nSEvent pressureBehavior];
    return result_;
}

unsigned int C_NSEvent_CapabilityMask(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result_ = [nSEvent capabilityMask];
    return result_;
}

unsigned int C_NSEvent_DeviceID(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result_ = [nSEvent deviceID];
    return result_;
}

bool C_NSEvent_IsEnteringProximity(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    BOOL result_ = [nSEvent isEnteringProximity];
    return result_;
}

unsigned int C_NSEvent_PointingDeviceID(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result_ = [nSEvent pointingDeviceID];
    return result_;
}

unsigned int C_NSEvent_PointingDeviceSerialNumber(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result_ = [nSEvent pointingDeviceSerialNumber];
    return result_;
}

unsigned int C_NSEvent_PointingDeviceType(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSPointingDeviceType result_ = [nSEvent pointingDeviceType];
    return result_;
}

unsigned int C_NSEvent_SystemTabletID(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result_ = [nSEvent systemTabletID];
    return result_;
}

unsigned int C_NSEvent_TabletID(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result_ = [nSEvent tabletID];
    return result_;
}

unsigned int C_NSEvent_VendorID(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result_ = [nSEvent vendorID];
    return result_;
}

unsigned int C_NSEvent_VendorPointingDeviceType(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result_ = [nSEvent vendorPointingDeviceType];
    return result_;
}

int C_NSEvent_AbsoluteX(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result_ = [nSEvent absoluteX];
    return result_;
}

int C_NSEvent_AbsoluteY(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result_ = [nSEvent absoluteY];
    return result_;
}

int C_NSEvent_AbsoluteZ(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result_ = [nSEvent absoluteZ];
    return result_;
}

unsigned int C_NSEvent_ButtonMask(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEventButtonMask result_ = [nSEvent buttonMask];
    return result_;
}

float C_NSEvent_Rotation(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    float result_ = [nSEvent rotation];
    return result_;
}

float C_NSEvent_TangentialPressure(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    float result_ = [nSEvent tangentialPressure];
    return result_;
}

CGPoint C_NSEvent_Tilt(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSPoint result_ = [nSEvent tilt];
    return result_;
}

void* C_NSEvent_VendorDefined(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    id result_ = [nSEvent vendorDefined];
    return result_;
}

double C_NSEvent_Magnification(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result_ = [nSEvent magnification];
    return result_;
}

bool C_NSEvent_Event_IsMouseCoalescingEnabled() {
    BOOL result_ = [NSEvent isMouseCoalescingEnabled];
    return result_;
}

void C_NSEvent_Event_SetMouseCoalescingEnabled(bool value) {
    [NSEvent setMouseCoalescingEnabled:value];
}

bool C_NSEvent_HasPreciseScrollingDeltas(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    BOOL result_ = [nSEvent hasPreciseScrollingDeltas];
    return result_;
}

double C_NSEvent_ScrollingDeltaX(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result_ = [nSEvent scrollingDeltaX];
    return result_;
}

double C_NSEvent_ScrollingDeltaY(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result_ = [nSEvent scrollingDeltaY];
    return result_;
}

unsigned int C_NSEvent_MomentumPhase(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEventPhase result_ = [nSEvent momentumPhase];
    return result_;
}

unsigned int C_NSEvent_Phase(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEventPhase result_ = [nSEvent phase];
    return result_;
}

bool C_NSEvent_IsSwipeTrackingFromScrollEventsEnabled() {
    BOOL result_ = [NSEvent isSwipeTrackingFromScrollEventsEnabled];
    return result_;
}

bool C_NSEvent_IsDirectionInvertedFromDevice(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    BOOL result_ = [nSEvent isDirectionInvertedFromDevice];
    return result_;
}
