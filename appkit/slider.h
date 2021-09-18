#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_Slider_Alloc();

void* C_NSSlider_SliderWithTarget_Action(void* target, void* action);
void* C_NSSlider_SliderWithValue_MinValue_MaxValue_Target_Action(double value, double minValue, double maxValue, void* target, void* action);
void* C_NSSlider_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSSlider_InitWithCoder(void* ptr, void* coder);
void* C_NSSlider_Init(void* ptr);
void* C_NSSlider_AllocSlider();
void* C_NSSlider_NewSlider();
void* C_NSSlider_Autorelease(void* ptr);
void* C_NSSlider_Retain(void* ptr);
double C_NSSlider_ClosestTickMarkValueToValue(void* ptr, double value);
int C_NSSlider_IndexOfTickMarkAtPoint(void* ptr, CGPoint point);
CGRect C_NSSlider_RectOfTickMarkAtIndex(void* ptr, int index);
double C_NSSlider_TickMarkValueAtIndex(void* ptr, int index);
unsigned int C_NSSlider_SliderType(void* ptr);
void C_NSSlider_SetSliderType(void* ptr, unsigned int value);
double C_NSSlider_AltIncrementValue(void* ptr);
void C_NSSlider_SetAltIncrementValue(void* ptr, double value);
double C_NSSlider_KnobThickness(void* ptr);
bool C_NSSlider_IsVertical(void* ptr);
void C_NSSlider_SetVertical(void* ptr, bool value);
void* C_NSSlider_TrackFillColor(void* ptr);
void C_NSSlider_SetTrackFillColor(void* ptr, void* value);
double C_NSSlider_MaxValue(void* ptr);
void C_NSSlider_SetMaxValue(void* ptr, double value);
double C_NSSlider_MinValue(void* ptr);
void C_NSSlider_SetMinValue(void* ptr, double value);
bool C_NSSlider_AllowsTickMarkValuesOnly(void* ptr);
void C_NSSlider_SetAllowsTickMarkValuesOnly(void* ptr, bool value);
int C_NSSlider_NumberOfTickMarks(void* ptr);
void C_NSSlider_SetNumberOfTickMarks(void* ptr, int value);
unsigned int C_NSSlider_TickMarkPosition(void* ptr);
void C_NSSlider_SetTickMarkPosition(void* ptr, unsigned int value);
