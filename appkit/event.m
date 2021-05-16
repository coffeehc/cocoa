#import <Appkit/Appkit.h>
#import "event.h"

void* C_Event_Alloc() {
    return [NSEvent alloc];
}

void* C_NSEvent_Init(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEvent* result = [nSEvent init];
    return result;
}

void* C_NSEvent_Event_MouseEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_EventNumber_ClickCount_Pressure(unsigned int _type, CGPoint location, unsigned int flags, double time, int wNum, void* unusedPassNil, int eNum, int cNum, float pressure) {
    NSEvent* result = [NSEvent mouseEventWithType:_type location:location modifierFlags:flags timestamp:time windowNumber:wNum context:(NSGraphicsContext*)unusedPassNil eventNumber:eNum clickCount:cNum pressure:pressure];
    return result;
}

void* C_NSEvent_EventWithCGEvent(CGEventRef cgEvent) {
    NSEvent* result = [NSEvent eventWithCGEvent:cgEvent];
    return result;
}

void C_NSEvent_Event_StartPeriodicEventsAfterDelay_WithPeriod(double delay, double period) {
    [NSEvent startPeriodicEventsAfterDelay:delay withPeriod:period];
}

void C_NSEvent_Event_StopPeriodicEvents() {
    [NSEvent stopPeriodicEvents];
}

Array C_NSEvent_CoalescedTouchesForTouch(void* ptr, void* touch) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSArray* result = [nSEvent coalescedTouchesForTouch:(NSTouch*)touch];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void C_NSEvent_Event_RemoveMonitor(void* eventMonitor) {
    [NSEvent removeMonitor:(id)eventMonitor];
}

void* C_NSEvent_CharactersByApplyingModifiers(void* ptr, unsigned int modifiers) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSString* result = [nSEvent charactersByApplyingModifiers:modifiers];
    return result;
}

unsigned int C_NSEvent_Type(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEventType result = [nSEvent type];
    return result;
}

unsigned int C_NSEvent_ModifierFlags(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEventModifierFlags result = [nSEvent modifierFlags];
    return result;
}

CGPoint C_NSEvent_LocationInWindow(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSPoint result = [nSEvent locationInWindow];
    return result;
}

double C_NSEvent_Timestamp(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSTimeInterval result = [nSEvent timestamp];
    return result;
}

void* C_NSEvent_Window(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSWindow* result = [nSEvent window];
    return result;
}

int C_NSEvent_WindowNumber(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result = [nSEvent windowNumber];
    return result;
}

CGEventRef C_NSEvent_CGEvent(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGEventRef result = [nSEvent CGEvent];
    return result;
}

double C_NSEvent_Event_KeyRepeatDelay() {
    NSTimeInterval result = [NSEvent keyRepeatDelay];
    return result;
}

double C_NSEvent_Event_KeyRepeatInterval() {
    NSTimeInterval result = [NSEvent keyRepeatInterval];
    return result;
}

void* C_NSEvent_Characters(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSString* result = [nSEvent characters];
    return result;
}

void* C_NSEvent_CharactersIgnoringModifiers(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSString* result = [nSEvent charactersIgnoringModifiers];
    return result;
}

bool C_NSEvent_IsARepeat(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    BOOL result = [nSEvent isARepeat];
    return result;
}

unsigned int C_NSEvent_Event_PressedMouseButtons() {
    NSUInteger result = [NSEvent pressedMouseButtons];
    return result;
}

double C_NSEvent_Event_DoubleClickInterval() {
    NSTimeInterval result = [NSEvent doubleClickInterval];
    return result;
}

CGPoint C_NSEvent_Event_MouseLocation() {
    NSPoint result = [NSEvent mouseLocation];
    return result;
}

int C_NSEvent_ButtonNumber(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result = [nSEvent buttonNumber];
    return result;
}

int C_NSEvent_ClickCount(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result = [nSEvent clickCount];
    return result;
}

int C_NSEvent_EventNumber(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result = [nSEvent eventNumber];
    return result;
}

int C_NSEvent_TrackingNumber(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result = [nSEvent trackingNumber];
    return result;
}

void* C_NSEvent_TrackingArea(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSTrackingArea* result = [nSEvent trackingArea];
    return result;
}

int C_NSEvent_Data1(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result = [nSEvent data1];
    return result;
}

int C_NSEvent_Data2(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result = [nSEvent data2];
    return result;
}

double C_NSEvent_DeltaX(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result = [nSEvent deltaX];
    return result;
}

double C_NSEvent_DeltaY(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result = [nSEvent deltaY];
    return result;
}

double C_NSEvent_DeltaZ(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result = [nSEvent deltaZ];
    return result;
}

float C_NSEvent_Pressure(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    float result = [nSEvent pressure];
    return result;
}

int C_NSEvent_Stage(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result = [nSEvent stage];
    return result;
}

double C_NSEvent_StageTransition(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result = [nSEvent stageTransition];
    return result;
}

int C_NSEvent_PressureBehavior(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSPressureBehavior result = [nSEvent pressureBehavior];
    return result;
}

unsigned int C_NSEvent_CapabilityMask(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result = [nSEvent capabilityMask];
    return result;
}

unsigned int C_NSEvent_DeviceID(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result = [nSEvent deviceID];
    return result;
}

bool C_NSEvent_IsEnteringProximity(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    BOOL result = [nSEvent isEnteringProximity];
    return result;
}

unsigned int C_NSEvent_PointingDeviceID(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result = [nSEvent pointingDeviceID];
    return result;
}

unsigned int C_NSEvent_PointingDeviceSerialNumber(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result = [nSEvent pointingDeviceSerialNumber];
    return result;
}

unsigned int C_NSEvent_PointingDeviceType(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSPointingDeviceType result = [nSEvent pointingDeviceType];
    return result;
}

unsigned int C_NSEvent_SystemTabletID(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result = [nSEvent systemTabletID];
    return result;
}

unsigned int C_NSEvent_TabletID(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result = [nSEvent tabletID];
    return result;
}

unsigned int C_NSEvent_VendorID(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result = [nSEvent vendorID];
    return result;
}

unsigned int C_NSEvent_VendorPointingDeviceType(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSUInteger result = [nSEvent vendorPointingDeviceType];
    return result;
}

int C_NSEvent_AbsoluteX(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result = [nSEvent absoluteX];
    return result;
}

int C_NSEvent_AbsoluteY(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result = [nSEvent absoluteY];
    return result;
}

int C_NSEvent_AbsoluteZ(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSInteger result = [nSEvent absoluteZ];
    return result;
}

unsigned int C_NSEvent_ButtonMask(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEventButtonMask result = [nSEvent buttonMask];
    return result;
}

float C_NSEvent_Rotation(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    float result = [nSEvent rotation];
    return result;
}

float C_NSEvent_TangentialPressure(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    float result = [nSEvent tangentialPressure];
    return result;
}

CGPoint C_NSEvent_Tilt(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSPoint result = [nSEvent tilt];
    return result;
}

void* C_NSEvent_VendorDefined(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    id result = [nSEvent vendorDefined];
    return result;
}

double C_NSEvent_Magnification(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result = [nSEvent magnification];
    return result;
}

bool C_NSEvent_Event_MouseCoalescingEnabled() {
    BOOL result = [NSEvent mouseCoalescingEnabled];
    return result;
}

void C_NSEvent_Event_SetMouseCoalescingEnabled(bool value) {
    [NSEvent setMouseCoalescingEnabled:value];
}

bool C_NSEvent_HasPreciseScrollingDeltas(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    BOOL result = [nSEvent hasPreciseScrollingDeltas];
    return result;
}

double C_NSEvent_ScrollingDeltaX(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result = [nSEvent scrollingDeltaX];
    return result;
}

double C_NSEvent_ScrollingDeltaY(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    CGFloat result = [nSEvent scrollingDeltaY];
    return result;
}

unsigned int C_NSEvent_MomentumPhase(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEventPhase result = [nSEvent momentumPhase];
    return result;
}

unsigned int C_NSEvent_Phase(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    NSEventPhase result = [nSEvent phase];
    return result;
}

bool C_NSEvent_Event_SwipeTrackingFromScrollEventsEnabled() {
    BOOL result = [NSEvent swipeTrackingFromScrollEventsEnabled];
    return result;
}

bool C_NSEvent_IsDirectionInvertedFromDevice(void* ptr) {
    NSEvent* nSEvent = (NSEvent*)ptr;
    BOOL result = [nSEvent isDirectionInvertedFromDevice];
    return result;
}
