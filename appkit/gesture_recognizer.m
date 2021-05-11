#import <Appkit/Appkit.h>
#import "gesture_recognizer.h"

void* C_GestureRecognizer_Alloc() {
	return [NSGestureRecognizer alloc];
}

void* C_NSGestureRecognizer_InitWithTarget_Action(void* ptr, void* target, void* action) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	NSGestureRecognizer* result = [nSGestureRecognizer initWithTarget:(id)target action:(SEL)action];
	return result;
}

void* C_NSGestureRecognizer_InitWithCoder(void* ptr, void* coder) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	NSGestureRecognizer* result = [nSGestureRecognizer initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSGestureRecognizer_Init(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	NSGestureRecognizer* result = [nSGestureRecognizer init];
	return result;
}

CGPoint C_NSGestureRecognizer_LocationInView(void* ptr, void* view) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	CGPoint result = [nSGestureRecognizer locationInView:(NSView*)view];
	return result;
}

void C_NSGestureRecognizer_Reset(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer reset];
}

void C_NSGestureRecognizer_MouseDown(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer mouseDown:(NSEvent*)event];
}

void C_NSGestureRecognizer_MouseDragged(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer mouseDragged:(NSEvent*)event];
}

void C_NSGestureRecognizer_MouseUp(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer mouseUp:(NSEvent*)event];
}

void C_NSGestureRecognizer_OtherMouseDown(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer otherMouseDown:(NSEvent*)event];
}

void C_NSGestureRecognizer_OtherMouseDragged(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer otherMouseDragged:(NSEvent*)event];
}

void C_NSGestureRecognizer_OtherMouseUp(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer otherMouseUp:(NSEvent*)event];
}

void C_NSGestureRecognizer_RightMouseDown(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer rightMouseDown:(NSEvent*)event];
}

void C_NSGestureRecognizer_RightMouseDragged(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer rightMouseDragged:(NSEvent*)event];
}

void C_NSGestureRecognizer_RightMouseUp(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer rightMouseUp:(NSEvent*)event];
}

void C_NSGestureRecognizer_MagnifyWithEvent(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer magnifyWithEvent:(NSEvent*)event];
}

void C_NSGestureRecognizer_RotateWithEvent(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer rotateWithEvent:(NSEvent*)event];
}

bool C_NSGestureRecognizer_CanBePreventedByGestureRecognizer(void* ptr, void* preventingGestureRecognizer) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	bool result = [nSGestureRecognizer canBePreventedByGestureRecognizer:(NSGestureRecognizer*)preventingGestureRecognizer];
	return result;
}

bool C_NSGestureRecognizer_CanPreventGestureRecognizer(void* ptr, void* preventedGestureRecognizer) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	bool result = [nSGestureRecognizer canPreventGestureRecognizer:(NSGestureRecognizer*)preventedGestureRecognizer];
	return result;
}

bool C_NSGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(void* ptr, void* otherGestureRecognizer) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	bool result = [nSGestureRecognizer shouldBeRequiredToFailByGestureRecognizer:(NSGestureRecognizer*)otherGestureRecognizer];
	return result;
}

bool C_NSGestureRecognizer_ShouldRequireFailureOfGestureRecognizer(void* ptr, void* otherGestureRecognizer) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	bool result = [nSGestureRecognizer shouldRequireFailureOfGestureRecognizer:(NSGestureRecognizer*)otherGestureRecognizer];
	return result;
}

void C_NSGestureRecognizer_KeyDown(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer keyDown:(NSEvent*)event];
}

void C_NSGestureRecognizer_KeyUp(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer keyUp:(NSEvent*)event];
}

void C_NSGestureRecognizer_TabletPoint(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer tabletPoint:(NSEvent*)event];
}

void C_NSGestureRecognizer_FlagsChanged(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer flagsChanged:(NSEvent*)event];
}

void C_NSGestureRecognizer_PressureChangeWithEvent(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer pressureChangeWithEvent:(NSEvent*)event];
}

void C_NSGestureRecognizer_TouchesBeganWithEvent(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer touchesBeganWithEvent:(NSEvent*)event];
}

void C_NSGestureRecognizer_TouchesCancelledWithEvent(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer touchesCancelledWithEvent:(NSEvent*)event];
}

void C_NSGestureRecognizer_TouchesEndedWithEvent(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer touchesEndedWithEvent:(NSEvent*)event];
}

void C_NSGestureRecognizer_TouchesMovedWithEvent(void* ptr, void* event) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer touchesMovedWithEvent:(NSEvent*)event];
}

void* C_NSGestureRecognizer_Action(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	SEL result = [nSGestureRecognizer action];
	return result;
}

void C_NSGestureRecognizer_SetAction(void* ptr, void* value) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer setAction:(SEL)value];
}

void* C_NSGestureRecognizer_Target(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	id result = [nSGestureRecognizer target];
	return result;
}

void C_NSGestureRecognizer_SetTarget(void* ptr, void* value) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer setTarget:(id)value];
}

void* C_NSGestureRecognizer_View(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	NSView* result = [nSGestureRecognizer view];
	return result;
}

bool C_NSGestureRecognizer_IsEnabled(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	bool result = [nSGestureRecognizer isEnabled];
	return result;
}

void C_NSGestureRecognizer_SetEnabled(void* ptr, bool value) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer setEnabled:value];
}

bool C_NSGestureRecognizer_DelaysPrimaryMouseButtonEvents(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	bool result = [nSGestureRecognizer delaysPrimaryMouseButtonEvents];
	return result;
}

void C_NSGestureRecognizer_SetDelaysPrimaryMouseButtonEvents(void* ptr, bool value) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer setDelaysPrimaryMouseButtonEvents:value];
}

bool C_NSGestureRecognizer_DelaysSecondaryMouseButtonEvents(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	bool result = [nSGestureRecognizer delaysSecondaryMouseButtonEvents];
	return result;
}

void C_NSGestureRecognizer_SetDelaysSecondaryMouseButtonEvents(void* ptr, bool value) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer setDelaysSecondaryMouseButtonEvents:value];
}

bool C_NSGestureRecognizer_DelaysOtherMouseButtonEvents(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	bool result = [nSGestureRecognizer delaysOtherMouseButtonEvents];
	return result;
}

void C_NSGestureRecognizer_SetDelaysOtherMouseButtonEvents(void* ptr, bool value) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer setDelaysOtherMouseButtonEvents:value];
}

bool C_NSGestureRecognizer_DelaysKeyEvents(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	bool result = [nSGestureRecognizer delaysKeyEvents];
	return result;
}

void C_NSGestureRecognizer_SetDelaysKeyEvents(void* ptr, bool value) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer setDelaysKeyEvents:value];
}

bool C_NSGestureRecognizer_DelaysMagnificationEvents(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	bool result = [nSGestureRecognizer delaysMagnificationEvents];
	return result;
}

void C_NSGestureRecognizer_SetDelaysMagnificationEvents(void* ptr, bool value) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer setDelaysMagnificationEvents:value];
}

bool C_NSGestureRecognizer_DelaysRotationEvents(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	bool result = [nSGestureRecognizer delaysRotationEvents];
	return result;
}

void C_NSGestureRecognizer_SetDelaysRotationEvents(void* ptr, bool value) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer setDelaysRotationEvents:value];
}

void* C_NSGestureRecognizer_PressureConfiguration(void* ptr) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	NSPressureConfiguration* result = [nSGestureRecognizer pressureConfiguration];
	return result;
}

void C_NSGestureRecognizer_SetPressureConfiguration(void* ptr, void* value) {
	NSGestureRecognizer* nSGestureRecognizer = (NSGestureRecognizer*)ptr;
	[nSGestureRecognizer setPressureConfiguration:(NSPressureConfiguration*)value];
}
