#import <Appkit/Appkit.h>
#import "layout_anchor.h"

void* C_LayoutAnchor_Alloc() {
    return [NSLayoutAnchor alloc];
}

void* C_NSLayoutAnchor_Init(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutAnchor* result = [nSLayoutAnchor init];
    return result;
}

void* C_NSLayoutAnchor_ConstraintEqualToAnchor(void* ptr, void* anchor) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result = [nSLayoutAnchor constraintEqualToAnchor:(NSLayoutAnchor*)anchor];
    return result;
}

void* C_NSLayoutAnchor_ConstraintEqualToAnchor_Constant(void* ptr, void* anchor, double c) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result = [nSLayoutAnchor constraintEqualToAnchor:(NSLayoutAnchor*)anchor constant:c];
    return result;
}

void* C_NSLayoutAnchor_ConstraintGreaterThanOrEqualToAnchor(void* ptr, void* anchor) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result = [nSLayoutAnchor constraintGreaterThanOrEqualToAnchor:(NSLayoutAnchor*)anchor];
    return result;
}

void* C_NSLayoutAnchor_ConstraintGreaterThanOrEqualToAnchor_Constant(void* ptr, void* anchor, double c) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result = [nSLayoutAnchor constraintGreaterThanOrEqualToAnchor:(NSLayoutAnchor*)anchor constant:c];
    return result;
}

void* C_NSLayoutAnchor_ConstraintLessThanOrEqualToAnchor(void* ptr, void* anchor) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result = [nSLayoutAnchor constraintLessThanOrEqualToAnchor:(NSLayoutAnchor*)anchor];
    return result;
}

void* C_NSLayoutAnchor_ConstraintLessThanOrEqualToAnchor_Constant(void* ptr, void* anchor, double c) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result = [nSLayoutAnchor constraintLessThanOrEqualToAnchor:(NSLayoutAnchor*)anchor constant:c];
    return result;
}

Array C_NSLayoutAnchor_ConstraintsAffectingLayout(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSArray* result = [nSLayoutAnchor constraintsAffectingLayout];
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

bool C_NSLayoutAnchor_HasAmbiguousLayout(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    BOOL result = [nSLayoutAnchor hasAmbiguousLayout];
    return result;
}

void* C_NSLayoutAnchor_Name(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSString* result = [nSLayoutAnchor name];
    return result;
}

void* C_NSLayoutAnchor_Item(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    id result = [nSLayoutAnchor item];
    return result;
}
