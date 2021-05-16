#import <Appkit/Appkit.h>
#import "layout_constraint.h"

void* C_LayoutConstraint_Alloc() {
    return [NSLayoutConstraint alloc];
}

void* C_NSLayoutConstraint_Init(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutConstraint* result = [nSLayoutConstraint init];
    return result;
}

void* C_NSLayoutConstraint_LayoutConstraint_ConstraintWithItem_Attribute_RelatedBy_ToItem_Attribute_Multiplier_Constant(void* view1, int attr1, int relation, void* view2, int attr2, double multiplier, double c) {
    NSLayoutConstraint* result = [NSLayoutConstraint constraintWithItem:(id)view1 attribute:attr1 relatedBy:relation toItem:(id)view2 attribute:attr2 multiplier:multiplier constant:c];
    return result;
}

void C_NSLayoutConstraint_LayoutConstraint_ActivateConstraints(Array constraints) {
    NSMutableArray* objcConstraints = [[NSMutableArray alloc] init];
    void** constraintsData = (void**)constraints.data;
    for (int i = 0; i < constraints.len; i++) {
    	void* p = constraintsData[i];
    	[objcConstraints addObject:(NSLayoutConstraint*)(NSLayoutConstraint*)p];
    }
    [NSLayoutConstraint activateConstraints:objcConstraints];
}

void C_NSLayoutConstraint_LayoutConstraint_DeactivateConstraints(Array constraints) {
    NSMutableArray* objcConstraints = [[NSMutableArray alloc] init];
    void** constraintsData = (void**)constraints.data;
    for (int i = 0; i < constraints.len; i++) {
    	void* p = constraintsData[i];
    	[objcConstraints addObject:(NSLayoutConstraint*)(NSLayoutConstraint*)p];
    }
    [NSLayoutConstraint deactivateConstraints:objcConstraints];
}

bool C_NSLayoutConstraint_IsActive(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    BOOL result = [nSLayoutConstraint isActive];
    return result;
}

void C_NSLayoutConstraint_SetActive(void* ptr, bool value) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    [nSLayoutConstraint setActive:value];
}

void* C_NSLayoutConstraint_FirstItem(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    id result = [nSLayoutConstraint firstItem];
    return result;
}

int C_NSLayoutConstraint_FirstAttribute(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutAttribute result = [nSLayoutConstraint firstAttribute];
    return result;
}

int C_NSLayoutConstraint_Relation(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutRelation result = [nSLayoutConstraint relation];
    return result;
}

void* C_NSLayoutConstraint_SecondItem(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    id result = [nSLayoutConstraint secondItem];
    return result;
}

int C_NSLayoutConstraint_SecondAttribute(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutAttribute result = [nSLayoutConstraint secondAttribute];
    return result;
}

double C_NSLayoutConstraint_Multiplier(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    CGFloat result = [nSLayoutConstraint multiplier];
    return result;
}

double C_NSLayoutConstraint_Constant(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    CGFloat result = [nSLayoutConstraint constant];
    return result;
}

void C_NSLayoutConstraint_SetConstant(void* ptr, double value) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    [nSLayoutConstraint setConstant:value];
}

void* C_NSLayoutConstraint_FirstAnchor(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutAnchor* result = [nSLayoutConstraint firstAnchor];
    return result;
}

void* C_NSLayoutConstraint_SecondAnchor(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutAnchor* result = [nSLayoutConstraint secondAnchor];
    return result;
}

float C_NSLayoutConstraint_Priority(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSLayoutPriority result = [nSLayoutConstraint priority];
    return result;
}

void C_NSLayoutConstraint_SetPriority(void* ptr, float value) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    [nSLayoutConstraint setPriority:value];
}

void* C_NSLayoutConstraint_Identifier(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    NSString* result = [nSLayoutConstraint identifier];
    return result;
}

void C_NSLayoutConstraint_SetIdentifier(void* ptr, void* value) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    [nSLayoutConstraint setIdentifier:(NSString*)value];
}

bool C_NSLayoutConstraint_ShouldBeArchived(void* ptr) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    BOOL result = [nSLayoutConstraint shouldBeArchived];
    return result;
}

void C_NSLayoutConstraint_SetShouldBeArchived(void* ptr, bool value) {
    NSLayoutConstraint* nSLayoutConstraint = (NSLayoutConstraint*)ptr;
    [nSLayoutConstraint setShouldBeArchived:value];
}
