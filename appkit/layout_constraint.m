#import <Appkit/Appkit.h>
#import "layout_constraint.h"

void* C_LayoutConstraint_Alloc() {
    return [NSLayoutConstraint alloc];
}

void* C_NSLayoutConstraint_LayoutConstraint_ConstraintWithItem_Attribute_RelatedBy_ToItem_Attribute_Multiplier_Constant(void* view1, int attr1, int relation, void* view2, int attr2, double multiplier, double c) {
    NSLayoutConstraint* result_ = [NSLayoutConstraint constraintWithItem:(id)view1 attribute:attr1 relatedBy:relation toItem:(id)view2 attribute:attr2 multiplier:multiplier constant:c];
    return result_;
}

void* C_NSLayoutConstraint_AllocLayoutConstraint() {
    NSLayoutConstraint* result_ = [NSLayoutConstraint alloc];
    return result_;
}

void* C_NSLayoutConstraint_Init(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutConstraint init];
    return result_;
}

void* C_NSLayoutConstraint_NewLayoutConstraint() {
    NSLayoutConstraint* result_ = [NSLayoutConstraint new];
    return result_;
}

void* C_NSLayoutConstraint_Autorelease(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutConstraint autorelease];
    return result_;
}

void* C_NSLayoutConstraint_Retain(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutConstraint retain];
    return result_;
}

Array C_NSLayoutConstraint_LayoutConstraint_ConstraintsWithVisualFormat_Options_Metrics_Views(void* format, unsigned int opts, Dictionary metrics, Dictionary views) {
    NSMutableDictionary* objcMetrics = [NSMutableDictionary dictionaryWithCapacity:metrics.len];
    if (metrics.len > 0) {
    	void** metricsKeyData = (void**)metrics.key_data;
    	void** metricsValueData = (void**)metrics.value_data;
    	for (int i = 0; i < metrics.len; i++) {
    		void* kp = metricsKeyData[i];
    		void* vp = metricsValueData[i];
    		[objcMetrics setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    NSMutableDictionary* objcViews = [NSMutableDictionary dictionaryWithCapacity:views.len];
    if (views.len > 0) {
    	void** viewsKeyData = (void**)views.key_data;
    	void** viewsValueData = (void**)views.value_data;
    	for (int i = 0; i < views.len; i++) {
    		void* kp = viewsKeyData[i];
    		void* vp = viewsValueData[i];
    		[objcViews setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    NSArray* result_ = [NSLayoutConstraint constraintsWithVisualFormat:(NSString*)format options:opts metrics:objcMetrics views:objcViews];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSLayoutConstraint_LayoutConstraint_ActivateConstraints(Array constraints) {
    NSMutableArray* objcConstraints = [NSMutableArray arrayWithCapacity:constraints.len];
    if (constraints.len > 0) {
    	void** constraintsData = (void**)constraints.data;
    	for (int i = 0; i < constraints.len; i++) {
    		void* p = constraintsData[i];
    		[objcConstraints addObject:(NSLayoutConstraint*)p];
    	}
    }
    [NSLayoutConstraint activateConstraints:objcConstraints];
}

void C_NSLayoutConstraint_LayoutConstraint_DeactivateConstraints(Array constraints) {
    NSMutableArray* objcConstraints = [NSMutableArray arrayWithCapacity:constraints.len];
    if (constraints.len > 0) {
    	void** constraintsData = (void**)constraints.data;
    	for (int i = 0; i < constraints.len; i++) {
    		void* p = constraintsData[i];
    		[objcConstraints addObject:(NSLayoutConstraint*)p];
    	}
    }
    [NSLayoutConstraint deactivateConstraints:objcConstraints];
}

bool C_NSLayoutConstraint_IsActive(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    BOOL result_ = [nSLayoutConstraint isActive];
    return result_;
}

void C_NSLayoutConstraint_SetActive(void* ptr, bool value) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    [nSLayoutConstraint setActive:value];
}

void* C_NSLayoutConstraint_FirstItem(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    id result_ = [nSLayoutConstraint firstItem];
    return result_;
}

int C_NSLayoutConstraint_FirstAttribute(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutAttribute result_ = [nSLayoutConstraint firstAttribute];
    return result_;
}

int C_NSLayoutConstraint_Relation(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutRelation result_ = [nSLayoutConstraint relation];
    return result_;
}

void* C_NSLayoutConstraint_SecondItem(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    id result_ = [nSLayoutConstraint secondItem];
    return result_;
}

int C_NSLayoutConstraint_SecondAttribute(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutAttribute result_ = [nSLayoutConstraint secondAttribute];
    return result_;
}

double C_NSLayoutConstraint_Multiplier(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    CGFloat result_ = [nSLayoutConstraint multiplier];
    return result_;
}

double C_NSLayoutConstraint_Constant(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    CGFloat result_ = [nSLayoutConstraint constant];
    return result_;
}

void C_NSLayoutConstraint_SetConstant(void* ptr, double value) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    [nSLayoutConstraint setConstant:value];
}

void* C_NSLayoutConstraint_FirstAnchor(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutAnchor* result_ = [nSLayoutConstraint firstAnchor];
    return result_;
}

void* C_NSLayoutConstraint_SecondAnchor(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutAnchor* result_ = [nSLayoutConstraint secondAnchor];
    return result_;
}

float C_NSLayoutConstraint_Priority(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutPriority result_ = [nSLayoutConstraint priority];
    return result_;
}

void C_NSLayoutConstraint_SetPriority(void* ptr, float value) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    [nSLayoutConstraint setPriority:value];
}

void* C_NSLayoutConstraint_Identifier(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSString* result_ = [nSLayoutConstraint identifier];
    return result_;
}

void C_NSLayoutConstraint_SetIdentifier(void* ptr, void* value) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    [nSLayoutConstraint setIdentifier:(NSString*)value];
}

bool C_NSLayoutConstraint_ShouldBeArchived(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    BOOL result_ = [nSLayoutConstraint shouldBeArchived];
    return result_;
}

void C_NSLayoutConstraint_SetShouldBeArchived(void* ptr, bool value) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    [nSLayoutConstraint setShouldBeArchived:value];
}
