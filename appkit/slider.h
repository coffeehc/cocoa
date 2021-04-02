#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void Slider_SetDelegate(void *ptr, long goID);
unsigned long Slider_SliderType(void* ptr);
void Slider_SetSliderType(void* ptr, unsigned long sliderType);
double Slider_AltIncrementValue(void* ptr);
void Slider_SetAltIncrementValue(void* ptr, double altIncrementValue);
double Slider_KnobThickness(void* ptr);
bool Slider_IsVertical(void* ptr);
void Slider_SetVertical(void* ptr, bool vertical);
void* Slider_TrackFillColor(void* ptr);
void Slider_SetTrackFillColor(void* ptr, void* trackFillColor);
double Slider_MaxValue(void* ptr);
void Slider_SetMaxValue(void* ptr, double maxValue);
double Slider_MinValue(void* ptr);
void Slider_SetMinValue(void* ptr, double minValue);
bool Slider_AllowsTickMarkValuesOnly(void* ptr);
void Slider_SetAllowsTickMarkValuesOnly(void* ptr, bool allowsTickMarkValuesOnly);
long Slider_NumberOfTickMarks(void* ptr);
void Slider_SetNumberOfTickMarks(void* ptr, long numberOfTickMarks);
unsigned long Slider_TickMarkPosition(void* ptr);
void Slider_SetTickMarkPosition(void* ptr, unsigned long tickMarkPosition);

void* Slider_NewSlider(NSRect frame);
void* Slider_SliderWithTarget(void* target, void* action);
void* Slider_SliderWithValue(double value, double minValue, double maxValue, void* target, void* action);
bool Slider_AcceptsFirstMouse(void* ptr, void* event);
double Slider_ClosestTickMarkValueToValue(void* ptr, double value);
long Slider_IndexOfTickMarkAtPoint(void* ptr, NSPoint point);
NSRect Slider_RectOfTickMarkAtIndex(void* ptr, long index);
double Slider_TickMarkValueAtIndex(void* ptr, long index);
