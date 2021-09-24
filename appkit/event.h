#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_Event_Alloc();

void* C_NSEvent_AllocEvent();
void* C_NSEvent_Init(void* ptr);
void* C_NSEvent_NewEvent();
void* C_NSEvent_Autorelease(void* ptr);
void* C_NSEvent_Retain(void* ptr);
void* C_NSEvent_MouseEventWithType_Location_ModifierFlags_Timestamp_WindowNumber_Context_EventNumber_ClickCount_Pressure(unsigned int _type, CGPoint location, unsigned int flags, double time, int wNum, void* unusedPassNil, int eNum, int cNum, float pressure);
void* C_NSEvent_EventWithCGEvent(void* cgEvent);
void C_NSEvent_StartPeriodicEventsAfterDelay_WithPeriod(double delay, double period);
void C_NSEvent_StopPeriodicEvents();
void* C_NSEvent_TouchesMatchingPhase_InView(void* ptr, unsigned int phase, void* view);
void* C_NSEvent_AllTouches(void* ptr);
void* C_NSEvent_TouchesForView(void* ptr, void* view);
Array C_NSEvent_CoalescedTouchesForTouch(void* ptr, void* touch);
void C_NSEvent_Event_RemoveMonitor(void* eventMonitor);
void* C_NSEvent_CharactersByApplyingModifiers(void* ptr, unsigned int modifiers);
unsigned int C_NSEvent_Type(void* ptr);
unsigned int C_NSEvent_ModifierFlags(void* ptr);
CGPoint C_NSEvent_LocationInWindow(void* ptr);
double C_NSEvent_Timestamp(void* ptr);
void* C_NSEvent_Window(void* ptr);
int C_NSEvent_WindowNumber(void* ptr);
void* C_NSEvent_CGEvent(void* ptr);
double C_NSEvent_Event_KeyRepeatDelay();
double C_NSEvent_Event_KeyRepeatInterval();
void* C_NSEvent_Characters(void* ptr);
void* C_NSEvent_CharactersIgnoringModifiers(void* ptr);
bool C_NSEvent_IsARepeat(void* ptr);
unsigned int C_NSEvent_Event_PressedMouseButtons();
double C_NSEvent_Event_DoubleClickInterval();
CGPoint C_NSEvent_Event_MouseLocation();
int C_NSEvent_ButtonNumber(void* ptr);
int C_NSEvent_ClickCount(void* ptr);
int C_NSEvent_EventNumber(void* ptr);
int C_NSEvent_TrackingNumber(void* ptr);
void* C_NSEvent_TrackingArea(void* ptr);
int C_NSEvent_Data1(void* ptr);
int C_NSEvent_Data2(void* ptr);
double C_NSEvent_DeltaX(void* ptr);
double C_NSEvent_DeltaY(void* ptr);
double C_NSEvent_DeltaZ(void* ptr);
float C_NSEvent_Pressure(void* ptr);
int C_NSEvent_Stage(void* ptr);
double C_NSEvent_StageTransition(void* ptr);
int C_NSEvent_PressureBehavior(void* ptr);
unsigned int C_NSEvent_CapabilityMask(void* ptr);
unsigned int C_NSEvent_DeviceID(void* ptr);
bool C_NSEvent_IsEnteringProximity(void* ptr);
unsigned int C_NSEvent_PointingDeviceID(void* ptr);
unsigned int C_NSEvent_PointingDeviceSerialNumber(void* ptr);
unsigned int C_NSEvent_PointingDeviceType(void* ptr);
unsigned int C_NSEvent_SystemTabletID(void* ptr);
unsigned int C_NSEvent_TabletID(void* ptr);
unsigned int C_NSEvent_VendorID(void* ptr);
unsigned int C_NSEvent_VendorPointingDeviceType(void* ptr);
int C_NSEvent_AbsoluteX(void* ptr);
int C_NSEvent_AbsoluteY(void* ptr);
int C_NSEvent_AbsoluteZ(void* ptr);
unsigned int C_NSEvent_ButtonMask(void* ptr);
float C_NSEvent_Rotation(void* ptr);
float C_NSEvent_TangentialPressure(void* ptr);
CGPoint C_NSEvent_Tilt(void* ptr);
void* C_NSEvent_VendorDefined(void* ptr);
double C_NSEvent_Magnification(void* ptr);
bool C_NSEvent_Event_MouseCoalescingEnabled();
void C_NSEvent_Event_SetMouseCoalescingEnabled(bool value);
bool C_NSEvent_HasPreciseScrollingDeltas(void* ptr);
double C_NSEvent_ScrollingDeltaX(void* ptr);
double C_NSEvent_ScrollingDeltaY(void* ptr);
unsigned int C_NSEvent_MomentumPhase(void* ptr);
unsigned int C_NSEvent_Phase(void* ptr);
bool C_NSEvent_SwipeTrackingFromScrollEventsEnabled();
bool C_NSEvent_IsDirectionInvertedFromDevice(void* ptr);
