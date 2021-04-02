#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "stepper.h"

double Stepper_MaxValue(void* ptr) {
	NSStepper* stepper = (NSStepper*)ptr;
	return [stepper maxValue];
}

void Stepper_SetMaxValue(void* ptr, double maxValue) {
	NSStepper* stepper = (NSStepper*)ptr;
	[stepper setMaxValue:maxValue];
}

double Stepper_MinValue(void* ptr) {
	NSStepper* stepper = (NSStepper*)ptr;
	return [stepper minValue];
}

void Stepper_SetMinValue(void* ptr, double minValue) {
	NSStepper* stepper = (NSStepper*)ptr;
	[stepper setMinValue:minValue];
}

double Stepper_Increment(void* ptr) {
	NSStepper* stepper = (NSStepper*)ptr;
	return [stepper increment];
}

void Stepper_SetIncrement(void* ptr, double increment) {
	NSStepper* stepper = (NSStepper*)ptr;
	[stepper setIncrement:increment];
}

bool Stepper_Autorepeat(void* ptr) {
	NSStepper* stepper = (NSStepper*)ptr;
	return [stepper autorepeat];
}

void Stepper_SetAutorepeat(void* ptr, bool autorepeat) {
	NSStepper* stepper = (NSStepper*)ptr;
	[stepper setAutorepeat:autorepeat];
}

bool Stepper_ValueWraps(void* ptr) {
	NSStepper* stepper = (NSStepper*)ptr;
	return [stepper valueWraps];
}

void Stepper_SetValueWraps(void* ptr, bool valueWraps) {
	NSStepper* stepper = (NSStepper*)ptr;
	[stepper setValueWraps:valueWraps];
}

void* Stepper_NewStepper(NSRect frame) {
	NSStepper* stepper = [NSStepper alloc];
	return [[stepper initWithFrame:frame] autorelease];
}
