#import <Appkit/Appkit.h>
#import "stack_view.h"

void* C_StackView_Alloc() {
    return [NSStackView alloc];
}

void* C_NSStackView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSStackView* result_ = [nSStackView initWithFrame:frameRect];
    return result_;
}

void* C_NSStackView_InitWithCoder(void* ptr, void* coder) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSStackView* result_ = [nSStackView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSStackView_Init(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSStackView* result_ = [nSStackView init];
    return result_;
}

void* C_NSStackView_StackViewWithViews(Array views) {
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	void* p = viewsData[i];
    	[objcViews addObject:(NSView*)(NSView*)p];
    }
    NSStackView* result_ = [NSStackView stackViewWithViews:objcViews];
    return result_;
}

void C_NSStackView_AddView_InGravity(void* ptr, void* view, int gravity) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView addView:(NSView*)view inGravity:gravity];
}

void C_NSStackView_InsertView_AtIndex_InGravity(void* ptr, void* view, unsigned int index, int gravity) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView insertView:(NSView*)view atIndex:index inGravity:gravity];
}

void C_NSStackView_SetViews_InGravity(void* ptr, Array views, int gravity) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	void* p = viewsData[i];
    	[objcViews addObject:(NSView*)(NSView*)p];
    }
    [nSStackView setViews:objcViews inGravity:gravity];
}

void C_NSStackView_RemoveView(void* ptr, void* view) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView removeView:(NSView*)view];
}

void C_NSStackView_AddArrangedSubview(void* ptr, void* view) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView addArrangedSubview:(NSView*)view];
}

void C_NSStackView_InsertArrangedSubview_AtIndex(void* ptr, void* view, int index) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView insertArrangedSubview:(NSView*)view atIndex:index];
}

void C_NSStackView_RemoveArrangedSubview(void* ptr, void* view) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView removeArrangedSubview:(NSView*)view];
}

Array C_NSStackView_ViewsInGravity(void* ptr, int gravity) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSArray* result_ = [nSStackView viewsInGravity:gravity];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

float C_NSStackView_ClippingResistancePriorityForOrientation(void* ptr, int orientation) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSLayoutPriority result_ = [nSStackView clippingResistancePriorityForOrientation:orientation];
    return result_;
}

float C_NSStackView_HuggingPriorityForOrientation(void* ptr, int orientation) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSLayoutPriority result_ = [nSStackView huggingPriorityForOrientation:orientation];
    return result_;
}

double C_NSStackView_CustomSpacingAfterView(void* ptr, void* view) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    CGFloat result_ = [nSStackView customSpacingAfterView:(NSView*)view];
    return result_;
}

void C_NSStackView_SetCustomSpacing_AfterView(void* ptr, double spacing, void* view) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setCustomSpacing:spacing afterView:(NSView*)view];
}

float C_NSStackView_VisibilityPriorityForView(void* ptr, void* view) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSStackViewVisibilityPriority result_ = [nSStackView visibilityPriorityForView:(NSView*)view];
    return result_;
}

void C_NSStackView_SetVisibilityPriority_ForView(void* ptr, float priority, void* view) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setVisibilityPriority:priority forView:(NSView*)view];
}

void C_NSStackView_SetClippingResistancePriority_ForOrientation(void* ptr, float clippingResistancePriority, int orientation) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setClippingResistancePriority:clippingResistancePriority forOrientation:orientation];
}

void C_NSStackView_SetHuggingPriority_ForOrientation(void* ptr, float huggingPriority, int orientation) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setHuggingPriority:huggingPriority forOrientation:orientation];
}

void* C_NSStackView_Delegate(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    id result_ = [nSStackView delegate];
    return result_;
}

void C_NSStackView_SetDelegate(void* ptr, void* value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setDelegate:(id)value];
}

Array C_NSStackView_ArrangedSubviews(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSArray* result_ = [nSStackView arrangedSubviews];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

Array C_NSStackView_Views(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSArray* result_ = [nSStackView views];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

Array C_NSStackView_DetachedViews(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSArray* result_ = [nSStackView detachedViews];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
    return result_Array;
}

int C_NSStackView_Orientation(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSUserInterfaceLayoutOrientation result_ = [nSStackView orientation];
    return result_;
}

void C_NSStackView_SetOrientation(void* ptr, int value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setOrientation:value];
}

int C_NSStackView_Alignment(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSLayoutAttribute result_ = [nSStackView alignment];
    return result_;
}

void C_NSStackView_SetAlignment(void* ptr, int value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setAlignment:value];
}

double C_NSStackView_Spacing(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    CGFloat result_ = [nSStackView spacing];
    return result_;
}

void C_NSStackView_SetSpacing(void* ptr, double value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setSpacing:value];
}

NSEdgeInsets C_NSStackView_EdgeInsets(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSEdgeInsets result_ = [nSStackView edgeInsets];
    return result_;
}

void C_NSStackView_SetEdgeInsets(void* ptr, NSEdgeInsets value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setEdgeInsets:value];
}

int C_NSStackView_Distribution(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSStackViewDistribution result_ = [nSStackView distribution];
    return result_;
}

void C_NSStackView_SetDistribution(void* ptr, int value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setDistribution:value];
}

bool C_NSStackView_DetachesHiddenViews(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    BOOL result_ = [nSStackView detachesHiddenViews];
    return result_;
}

void C_NSStackView_SetDetachesHiddenViews(void* ptr, bool value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setDetachesHiddenViews:value];
}
