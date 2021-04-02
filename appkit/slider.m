#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "slider.h"
#import "slider_delegate.h"

@implementation NSSliderHandler

- (void)onAction:(id)sender {
	return Slider_Target_Action([self goID], sender);
}

@end
void Slider_SetDelegate(void *ptr, long goID) {
	NSSlider* slider = (NSSlider*)ptr;
	NSSliderHandler* handler = [[NSSliderHandler alloc] init];
	[handler setGoID:goID];
	[slider setTarget:handler];
	[slider setAction:@selector(onAction:)];
}

unsigned long Slider_SliderType(void* ptr) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider sliderType];
}

void Slider_SetSliderType(void* ptr, unsigned long sliderType) {
	NSSlider* slider = (NSSlider*)ptr;
	[slider setSliderType:sliderType];
}

double Slider_AltIncrementValue(void* ptr) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider altIncrementValue];
}

void Slider_SetAltIncrementValue(void* ptr, double altIncrementValue) {
	NSSlider* slider = (NSSlider*)ptr;
	[slider setAltIncrementValue:altIncrementValue];
}

double Slider_KnobThickness(void* ptr) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider knobThickness];
}

bool Slider_IsVertical(void* ptr) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider isVertical];
}

void Slider_SetVertical(void* ptr, bool vertical) {
	NSSlider* slider = (NSSlider*)ptr;
	[slider setVertical:vertical];
}

void* Slider_TrackFillColor(void* ptr) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider trackFillColor];
}

void Slider_SetTrackFillColor(void* ptr, void* trackFillColor) {
	NSSlider* slider = (NSSlider*)ptr;
	[slider setTrackFillColor:(NSColor*)trackFillColor];
}

double Slider_MaxValue(void* ptr) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider maxValue];
}

void Slider_SetMaxValue(void* ptr, double maxValue) {
	NSSlider* slider = (NSSlider*)ptr;
	[slider setMaxValue:maxValue];
}

double Slider_MinValue(void* ptr) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider minValue];
}

void Slider_SetMinValue(void* ptr, double minValue) {
	NSSlider* slider = (NSSlider*)ptr;
	[slider setMinValue:minValue];
}

bool Slider_AllowsTickMarkValuesOnly(void* ptr) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider allowsTickMarkValuesOnly];
}

void Slider_SetAllowsTickMarkValuesOnly(void* ptr, bool allowsTickMarkValuesOnly) {
	NSSlider* slider = (NSSlider*)ptr;
	[slider setAllowsTickMarkValuesOnly:allowsTickMarkValuesOnly];
}

long Slider_NumberOfTickMarks(void* ptr) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider numberOfTickMarks];
}

void Slider_SetNumberOfTickMarks(void* ptr, long numberOfTickMarks) {
	NSSlider* slider = (NSSlider*)ptr;
	[slider setNumberOfTickMarks:numberOfTickMarks];
}

unsigned long Slider_TickMarkPosition(void* ptr) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider tickMarkPosition];
}

void Slider_SetTickMarkPosition(void* ptr, unsigned long tickMarkPosition) {
	NSSlider* slider = (NSSlider*)ptr;
	[slider setTickMarkPosition:tickMarkPosition];
}

void* Slider_NewSlider(NSRect frame) {
	NSSlider* slider = [NSSlider alloc];
	return [[slider initWithFrame:frame] autorelease];
}

void* Slider_SliderWithTarget(void* target, void* action) {
	return [NSSlider sliderWithTarget:(NSObject*)target action:(SEL)action];
}

void* Slider_SliderWithValue(double value, double minValue, double maxValue, void* target, void* action) {
	return [NSSlider sliderWithValue:value minValue:minValue maxValue:maxValue target:(NSObject*)target action:(SEL)action];
}

bool Slider_AcceptsFirstMouse(void* ptr, void* event) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider acceptsFirstMouse:(NSEvent*)event];
}

double Slider_ClosestTickMarkValueToValue(void* ptr, double value) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider closestTickMarkValueToValue:value];
}

long Slider_IndexOfTickMarkAtPoint(void* ptr, NSPoint point) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider indexOfTickMarkAtPoint:point];
}

NSRect Slider_RectOfTickMarkAtIndex(void* ptr, long index) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider rectOfTickMarkAtIndex:index];
}

double Slider_TickMarkValueAtIndex(void* ptr, long index) {
	NSSlider* slider = (NSSlider*)ptr;
	return [slider tickMarkValueAtIndex:index];
}
