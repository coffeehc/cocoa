#import <Appkit/Appkit.h>
#import "stack_view.h"

void* C_StackView_Alloc() {
    return [NSStackView alloc];
}

void* C_NSStackView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSStackView* result = [nSStackView initWithFrame:frameRect];
    return result;
}

void* C_NSStackView_InitWithCoder(void* ptr, void* coder) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSStackView* result = [nSStackView initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSStackView_Init(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSStackView* result = [nSStackView init];
    return result;
}

void* C_NSStackView_StackViewWithViews(Array views) {
    NSMutableArray* objcViews = [[NSMutableArray alloc] init];
    void** viewsData = (void**)views.data;
    for (int i = 0; i < views.len; i++) {
    	void* p = viewsData[i];
    	[objcViews addObject:(NSView*)(NSView*)p];
    }
    NSStackView* result = [NSStackView stackViewWithViews:objcViews];
    return result;
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
    NSArray* result = [nSStackView viewsInGravity:gravity];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

float C_NSStackView_ClippingResistancePriorityForOrientation(void* ptr, int orientation) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSLayoutPriority result = [nSStackView clippingResistancePriorityForOrientation:orientation];
    return result;
}

float C_NSStackView_HuggingPriorityForOrientation(void* ptr, int orientation) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSLayoutPriority result = [nSStackView huggingPriorityForOrientation:orientation];
    return result;
}

double C_NSStackView_CustomSpacingAfterView(void* ptr, void* view) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    CGFloat result = [nSStackView customSpacingAfterView:(NSView*)view];
    return result;
}

void C_NSStackView_SetCustomSpacing_AfterView(void* ptr, double spacing, void* view) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setCustomSpacing:spacing afterView:(NSView*)view];
}

float C_NSStackView_VisibilityPriorityForView(void* ptr, void* view) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSStackViewVisibilityPriority result = [nSStackView visibilityPriorityForView:(NSView*)view];
    return result;
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

Array C_NSStackView_ArrangedSubviews(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSArray* result = [nSStackView arrangedSubviews];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

Array C_NSStackView_Views(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSArray* result = [nSStackView views];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

Array C_NSStackView_DetachedViews(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSArray* result = [nSStackView detachedViews];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

int C_NSStackView_Orientation(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSUserInterfaceLayoutOrientation result = [nSStackView orientation];
    return result;
}

void C_NSStackView_SetOrientation(void* ptr, int value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setOrientation:value];
}

int C_NSStackView_Alignment(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSLayoutAttribute result = [nSStackView alignment];
    return result;
}

void C_NSStackView_SetAlignment(void* ptr, int value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setAlignment:value];
}

double C_NSStackView_Spacing(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    CGFloat result = [nSStackView spacing];
    return result;
}

void C_NSStackView_SetSpacing(void* ptr, double value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setSpacing:value];
}

NSEdgeInsets C_NSStackView_EdgeInsets(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSEdgeInsets result = [nSStackView edgeInsets];
    return result;
}

void C_NSStackView_SetEdgeInsets(void* ptr, NSEdgeInsets value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setEdgeInsets:value];
}

int C_NSStackView_Distribution(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    NSStackViewDistribution result = [nSStackView distribution];
    return result;
}

void C_NSStackView_SetDistribution(void* ptr, int value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setDistribution:value];
}

bool C_NSStackView_DetachesHiddenViews(void* ptr) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    BOOL result = [nSStackView detachesHiddenViews];
    return result;
}

void C_NSStackView_SetDetachesHiddenViews(void* ptr, bool value) {
    NSStackView* nSStackView = (NSStackView*)ptr;
    [nSStackView setDetachesHiddenViews:value];
}
