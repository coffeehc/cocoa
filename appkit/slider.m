#import <Appkit/Appkit.h>
#import "slider.h"

void* C_Slider_Alloc() {
    return [NSSlider alloc];
}

void* C_NSSlider_InitWithFrame(void* ptr, CGRect frameRect) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    NSSlider* result_ = [nSSlider initWithFrame:frameRect];
    return result_;
}

void* C_NSSlider_InitWithCoder(void* ptr, void* coder) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    NSSlider* result_ = [nSSlider initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSSlider_Init(void* ptr) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    NSSlider* result_ = [nSSlider init];
    return result_;
}

void* C_NSSlider_SliderWithTarget_Action(void* target, void* action) {
    NSSlider* result_ = [NSSlider sliderWithTarget:(id)target action:(SEL)action];
    return result_;
}

void* C_NSSlider_SliderWithValue_MinValue_MaxValue_Target_Action(double value, double minValue, double maxValue, void* target, void* action) {
    NSSlider* result_ = [NSSlider sliderWithValue:value minValue:minValue maxValue:maxValue target:(id)target action:(SEL)action];
    return result_;
}

double C_NSSlider_ClosestTickMarkValueToValue(void* ptr, double value) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    double result_ = [nSSlider closestTickMarkValueToValue:value];
    return result_;
}

int C_NSSlider_IndexOfTickMarkAtPoint(void* ptr, CGPoint point) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    NSInteger result_ = [nSSlider indexOfTickMarkAtPoint:point];
    return result_;
}

CGRect C_NSSlider_RectOfTickMarkAtIndex(void* ptr, int index) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    NSRect result_ = [nSSlider rectOfTickMarkAtIndex:index];
    return result_;
}

double C_NSSlider_TickMarkValueAtIndex(void* ptr, int index) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    double result_ = [nSSlider tickMarkValueAtIndex:index];
    return result_;
}

unsigned int C_NSSlider_SliderType(void* ptr) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    NSSliderType result_ = [nSSlider sliderType];
    return result_;
}

void C_NSSlider_SetSliderType(void* ptr, unsigned int value) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    [nSSlider setSliderType:value];
}

double C_NSSlider_AltIncrementValue(void* ptr) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    double result_ = [nSSlider altIncrementValue];
    return result_;
}

void C_NSSlider_SetAltIncrementValue(void* ptr, double value) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    [nSSlider setAltIncrementValue:value];
}

double C_NSSlider_KnobThickness(void* ptr) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    CGFloat result_ = [nSSlider knobThickness];
    return result_;
}

bool C_NSSlider_IsVertical(void* ptr) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    BOOL result_ = [nSSlider isVertical];
    return result_;
}

void C_NSSlider_SetVertical(void* ptr, bool value) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    [nSSlider setVertical:value];
}

void* C_NSSlider_TrackFillColor(void* ptr) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    NSColor* result_ = [nSSlider trackFillColor];
    return result_;
}

void C_NSSlider_SetTrackFillColor(void* ptr, void* value) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    [nSSlider setTrackFillColor:(NSColor*)value];
}

double C_NSSlider_MaxValue(void* ptr) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    double result_ = [nSSlider maxValue];
    return result_;
}

void C_NSSlider_SetMaxValue(void* ptr, double value) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    [nSSlider setMaxValue:value];
}

double C_NSSlider_MinValue(void* ptr) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    double result_ = [nSSlider minValue];
    return result_;
}

void C_NSSlider_SetMinValue(void* ptr, double value) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    [nSSlider setMinValue:value];
}

bool C_NSSlider_AllowsTickMarkValuesOnly(void* ptr) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    BOOL result_ = [nSSlider allowsTickMarkValuesOnly];
    return result_;
}

void C_NSSlider_SetAllowsTickMarkValuesOnly(void* ptr, bool value) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    [nSSlider setAllowsTickMarkValuesOnly:value];
}

int C_NSSlider_NumberOfTickMarks(void* ptr) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    NSInteger result_ = [nSSlider numberOfTickMarks];
    return result_;
}

void C_NSSlider_SetNumberOfTickMarks(void* ptr, int value) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    [nSSlider setNumberOfTickMarks:value];
}

unsigned int C_NSSlider_TickMarkPosition(void* ptr) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    NSTickMarkPosition result_ = [nSSlider tickMarkPosition];
    return result_;
}

void C_NSSlider_SetTickMarkPosition(void* ptr, unsigned int value) {
    NSSlider* nSSlider = (NSSlider*)ptr;
    [nSSlider setTickMarkPosition:value];
}
