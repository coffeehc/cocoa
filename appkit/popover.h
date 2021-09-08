#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_Popover_Alloc();

void* C_NSPopover_Init(void* ptr);
void* C_NSPopover_InitWithCoder(void* ptr, void* coder);
void* C_NSPopover_AllocPopover();
void* C_NSPopover_NewPopover();
void* C_NSPopover_Autorelease(void* ptr);
void* C_NSPopover_Retain(void* ptr);
void C_NSPopover_ShowRelativeToRect_OfView_PreferredEdge(void* ptr, CGRect positioningRect, void* positioningView, unsigned int preferredEdge);
void C_NSPopover_PerformClose(void* ptr, void* sender);
void C_NSPopover_Close(void* ptr);
int C_NSPopover_Behavior(void* ptr);
void C_NSPopover_SetBehavior(void* ptr, int value);
CGRect C_NSPopover_PositioningRect(void* ptr);
void C_NSPopover_SetPositioningRect(void* ptr, CGRect value);
void* C_NSPopover_Appearance(void* ptr);
void C_NSPopover_SetAppearance(void* ptr, void* value);
void* C_NSPopover_EffectiveAppearance(void* ptr);
bool C_NSPopover_Animates(void* ptr);
void C_NSPopover_SetAnimates(void* ptr, bool value);
CGSize C_NSPopover_ContentSize(void* ptr);
void C_NSPopover_SetContentSize(void* ptr, CGSize value);
bool C_NSPopover_IsShown(void* ptr);
bool C_NSPopover_IsDetached(void* ptr);
