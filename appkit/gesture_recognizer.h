#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_GestureRecognizer_Alloc();

void* C_NSGestureRecognizer_InitWithTarget_Action(void* ptr, void* target, void* action);
void* C_NSGestureRecognizer_InitWithCoder(void* ptr, void* coder);
void* C_NSGestureRecognizer_Init(void* ptr);
CGPoint C_NSGestureRecognizer_LocationInView(void* ptr, void* view);
void C_NSGestureRecognizer_Reset(void* ptr);
void C_NSGestureRecognizer_MouseDown(void* ptr, void* event);
void C_NSGestureRecognizer_MouseDragged(void* ptr, void* event);
void C_NSGestureRecognizer_MouseUp(void* ptr, void* event);
void C_NSGestureRecognizer_OtherMouseDown(void* ptr, void* event);
void C_NSGestureRecognizer_OtherMouseDragged(void* ptr, void* event);
void C_NSGestureRecognizer_OtherMouseUp(void* ptr, void* event);
void C_NSGestureRecognizer_RightMouseDown(void* ptr, void* event);
void C_NSGestureRecognizer_RightMouseDragged(void* ptr, void* event);
void C_NSGestureRecognizer_RightMouseUp(void* ptr, void* event);
void C_NSGestureRecognizer_MagnifyWithEvent(void* ptr, void* event);
void C_NSGestureRecognizer_RotateWithEvent(void* ptr, void* event);
bool C_NSGestureRecognizer_CanBePreventedByGestureRecognizer(void* ptr, void* preventingGestureRecognizer);
bool C_NSGestureRecognizer_CanPreventGestureRecognizer(void* ptr, void* preventedGestureRecognizer);
bool C_NSGestureRecognizer_ShouldBeRequiredToFailByGestureRecognizer(void* ptr, void* otherGestureRecognizer);
bool C_NSGestureRecognizer_ShouldRequireFailureOfGestureRecognizer(void* ptr, void* otherGestureRecognizer);
void C_NSGestureRecognizer_KeyDown(void* ptr, void* event);
void C_NSGestureRecognizer_KeyUp(void* ptr, void* event);
void C_NSGestureRecognizer_TabletPoint(void* ptr, void* event);
void C_NSGestureRecognizer_FlagsChanged(void* ptr, void* event);
void C_NSGestureRecognizer_PressureChangeWithEvent(void* ptr, void* event);
void C_NSGestureRecognizer_TouchesBeganWithEvent(void* ptr, void* event);
void C_NSGestureRecognizer_TouchesCancelledWithEvent(void* ptr, void* event);
void C_NSGestureRecognizer_TouchesEndedWithEvent(void* ptr, void* event);
void C_NSGestureRecognizer_TouchesMovedWithEvent(void* ptr, void* event);
void* C_NSGestureRecognizer_Action(void* ptr);
void C_NSGestureRecognizer_SetAction(void* ptr, void* value);
void* C_NSGestureRecognizer_Target(void* ptr);
void C_NSGestureRecognizer_SetTarget(void* ptr, void* value);
void* C_NSGestureRecognizer_View(void* ptr);
bool C_NSGestureRecognizer_IsEnabled(void* ptr);
void C_NSGestureRecognizer_SetEnabled(void* ptr, bool value);
bool C_NSGestureRecognizer_DelaysPrimaryMouseButtonEvents(void* ptr);
void C_NSGestureRecognizer_SetDelaysPrimaryMouseButtonEvents(void* ptr, bool value);
bool C_NSGestureRecognizer_DelaysSecondaryMouseButtonEvents(void* ptr);
void C_NSGestureRecognizer_SetDelaysSecondaryMouseButtonEvents(void* ptr, bool value);
bool C_NSGestureRecognizer_DelaysOtherMouseButtonEvents(void* ptr);
void C_NSGestureRecognizer_SetDelaysOtherMouseButtonEvents(void* ptr, bool value);
bool C_NSGestureRecognizer_DelaysKeyEvents(void* ptr);
void C_NSGestureRecognizer_SetDelaysKeyEvents(void* ptr, bool value);
bool C_NSGestureRecognizer_DelaysMagnificationEvents(void* ptr);
void C_NSGestureRecognizer_SetDelaysMagnificationEvents(void* ptr, bool value);
bool C_NSGestureRecognizer_DelaysRotationEvents(void* ptr);
void C_NSGestureRecognizer_SetDelaysRotationEvents(void* ptr, bool value);
void* C_NSGestureRecognizer_PressureConfiguration(void* ptr);
void C_NSGestureRecognizer_SetPressureConfiguration(void* ptr, void* value);
