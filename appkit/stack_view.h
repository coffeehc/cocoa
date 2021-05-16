#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_StackView_Alloc();

void* C_NSStackView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSStackView_InitWithCoder(void* ptr, void* coder);
void* C_NSStackView_Init(void* ptr);
void* C_NSStackView_StackViewWithViews(Array views);
void C_NSStackView_AddView_InGravity(void* ptr, void* view, int gravity);
void C_NSStackView_InsertView_AtIndex_InGravity(void* ptr, void* view, unsigned int index, int gravity);
void C_NSStackView_SetViews_InGravity(void* ptr, Array views, int gravity);
void C_NSStackView_RemoveView(void* ptr, void* view);
void C_NSStackView_AddArrangedSubview(void* ptr, void* view);
void C_NSStackView_InsertArrangedSubview_AtIndex(void* ptr, void* view, int index);
void C_NSStackView_RemoveArrangedSubview(void* ptr, void* view);
Array C_NSStackView_ViewsInGravity(void* ptr, int gravity);
float C_NSStackView_ClippingResistancePriorityForOrientation(void* ptr, int orientation);
float C_NSStackView_HuggingPriorityForOrientation(void* ptr, int orientation);
double C_NSStackView_CustomSpacingAfterView(void* ptr, void* view);
void C_NSStackView_SetCustomSpacing_AfterView(void* ptr, double spacing, void* view);
float C_NSStackView_VisibilityPriorityForView(void* ptr, void* view);
void C_NSStackView_SetVisibilityPriority_ForView(void* ptr, float priority, void* view);
void C_NSStackView_SetClippingResistancePriority_ForOrientation(void* ptr, float clippingResistancePriority, int orientation);
void C_NSStackView_SetHuggingPriority_ForOrientation(void* ptr, float huggingPriority, int orientation);
Array C_NSStackView_ArrangedSubviews(void* ptr);
Array C_NSStackView_Views(void* ptr);
Array C_NSStackView_DetachedViews(void* ptr);
int C_NSStackView_Orientation(void* ptr);
void C_NSStackView_SetOrientation(void* ptr, int value);
int C_NSStackView_Alignment(void* ptr);
void C_NSStackView_SetAlignment(void* ptr, int value);
double C_NSStackView_Spacing(void* ptr);
void C_NSStackView_SetSpacing(void* ptr, double value);
NSEdgeInsets C_NSStackView_EdgeInsets(void* ptr);
void C_NSStackView_SetEdgeInsets(void* ptr, NSEdgeInsets value);
int C_NSStackView_Distribution(void* ptr);
void C_NSStackView_SetDistribution(void* ptr, int value);
bool C_NSStackView_DetachesHiddenViews(void* ptr);
void C_NSStackView_SetDetachesHiddenViews(void* ptr, bool value);
