#import "layout_anchor.h"
#import <AppKit/NSLayoutAnchor.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_LayoutAnchor_Alloc() {
    return [NSLayoutAnchor alloc];
}

void* C_NSLayoutAnchor_AllocLayoutAnchor() {
    NSLayoutAnchor* result_ = [NSLayoutAnchor alloc];
    return result_;
}

void* C_NSLayoutAnchor_Init(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutAnchor* result_ = [nSLayoutAnchor init];
    return result_;
}

void* C_NSLayoutAnchor_NewLayoutAnchor() {
    NSLayoutAnchor* result_ = [NSLayoutAnchor new];
    return result_;
}

void* C_NSLayoutAnchor_Autorelease(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutAnchor* result_ = [nSLayoutAnchor autorelease];
    return result_;
}

void* C_NSLayoutAnchor_Retain(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutAnchor* result_ = [nSLayoutAnchor retain];
    return result_;
}

void* C_NSLayoutAnchor_ConstraintEqualToAnchor(void* ptr, void* anchor) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutAnchor constraintEqualToAnchor:(NSLayoutAnchor*)anchor];
    return result_;
}

void* C_NSLayoutAnchor_ConstraintEqualToAnchor_Constant(void* ptr, void* anchor, double c) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutAnchor constraintEqualToAnchor:(NSLayoutAnchor*)anchor constant:c];
    return result_;
}

void* C_NSLayoutAnchor_ConstraintGreaterThanOrEqualToAnchor(void* ptr, void* anchor) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutAnchor constraintGreaterThanOrEqualToAnchor:(NSLayoutAnchor*)anchor];
    return result_;
}

void* C_NSLayoutAnchor_ConstraintGreaterThanOrEqualToAnchor_Constant(void* ptr, void* anchor, double c) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutAnchor constraintGreaterThanOrEqualToAnchor:(NSLayoutAnchor*)anchor constant:c];
    return result_;
}

void* C_NSLayoutAnchor_ConstraintLessThanOrEqualToAnchor(void* ptr, void* anchor) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutAnchor constraintLessThanOrEqualToAnchor:(NSLayoutAnchor*)anchor];
    return result_;
}

void* C_NSLayoutAnchor_ConstraintLessThanOrEqualToAnchor_Constant(void* ptr, void* anchor, double c) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutAnchor constraintLessThanOrEqualToAnchor:(NSLayoutAnchor*)anchor constant:c];
    return result_;
}

Array C_NSLayoutAnchor_ConstraintsAffectingLayout(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSArray* result_ = [nSLayoutAnchor constraintsAffectingLayout];
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

bool C_NSLayoutAnchor_HasAmbiguousLayout(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    BOOL result_ = [nSLayoutAnchor hasAmbiguousLayout];
    return result_;
}

void* C_NSLayoutAnchor_Name(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    NSString* result_ = [nSLayoutAnchor name];
    return result_;
}

void* C_NSLayoutAnchor_Item(void* ptr) {
    NSLayoutAnchor* nSLayoutAnchor = (NSLayoutAnchor*)ptr;
    id result_ = [nSLayoutAnchor item];
    return result_;
}

void* C_LayoutDimension_Alloc() {
    return [NSLayoutDimension alloc];
}

void* C_NSLayoutDimension_AllocLayoutDimension() {
    NSLayoutDimension* result_ = [NSLayoutDimension alloc];
    return result_;
}

void* C_NSLayoutDimension_Init(void* ptr) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutDimension* result_ = [nSLayoutDimension init];
    return result_;
}

void* C_NSLayoutDimension_NewLayoutDimension() {
    NSLayoutDimension* result_ = [NSLayoutDimension new];
    return result_;
}

void* C_NSLayoutDimension_Autorelease(void* ptr) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutDimension* result_ = [nSLayoutDimension autorelease];
    return result_;
}

void* C_NSLayoutDimension_Retain(void* ptr) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutDimension* result_ = [nSLayoutDimension retain];
    return result_;
}

void* C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier(void* ptr, void* anchor, double m) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutDimension constraintEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m];
    return result_;
}

void* C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutDimension constraintEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m constant:c];
    return result_;
}

void* C_NSLayoutDimension_ConstraintEqualToConstant(void* ptr, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutDimension constraintEqualToConstant:c];
    return result_;
}

void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier(void* ptr, void* anchor, double m) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutDimension constraintGreaterThanOrEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m];
    return result_;
}

void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutDimension constraintGreaterThanOrEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m constant:c];
    return result_;
}

void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToConstant(void* ptr, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutDimension constraintGreaterThanOrEqualToConstant:c];
    return result_;
}

void* C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier(void* ptr, void* anchor, double m) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutDimension constraintLessThanOrEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m];
    return result_;
}

void* C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutDimension constraintLessThanOrEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m constant:c];
    return result_;
}

void* C_NSLayoutDimension_ConstraintLessThanOrEqualToConstant(void* ptr, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutDimension constraintLessThanOrEqualToConstant:c];
    return result_;
}

void* C_LayoutXAxisAnchor_Alloc() {
    return [NSLayoutXAxisAnchor alloc];
}

void* C_NSLayoutXAxisAnchor_AllocLayoutXAxisAnchor() {
    NSLayoutXAxisAnchor* result_ = [NSLayoutXAxisAnchor alloc];
    return result_;
}

void* C_NSLayoutXAxisAnchor_Init(void* ptr) {
    NSLayoutXAxisAnchor* nSLayoutXAxisAnchor = (NSLayoutXAxisAnchor*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSLayoutXAxisAnchor init];
    return result_;
}

void* C_NSLayoutXAxisAnchor_NewLayoutXAxisAnchor() {
    NSLayoutXAxisAnchor* result_ = [NSLayoutXAxisAnchor new];
    return result_;
}

void* C_NSLayoutXAxisAnchor_Autorelease(void* ptr) {
    NSLayoutXAxisAnchor* nSLayoutXAxisAnchor = (NSLayoutXAxisAnchor*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSLayoutXAxisAnchor autorelease];
    return result_;
}

void* C_NSLayoutXAxisAnchor_Retain(void* ptr) {
    NSLayoutXAxisAnchor* nSLayoutXAxisAnchor = (NSLayoutXAxisAnchor*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSLayoutXAxisAnchor retain];
    return result_;
}

void* C_NSLayoutXAxisAnchor_ConstraintEqualToSystemSpacingAfterAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutXAxisAnchor* nSLayoutXAxisAnchor = (NSLayoutXAxisAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutXAxisAnchor constraintEqualToSystemSpacingAfterAnchor:(NSLayoutXAxisAnchor*)anchor multiplier:multiplier];
    return result_;
}

void* C_NSLayoutXAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutXAxisAnchor* nSLayoutXAxisAnchor = (NSLayoutXAxisAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutXAxisAnchor constraintGreaterThanOrEqualToSystemSpacingAfterAnchor:(NSLayoutXAxisAnchor*)anchor multiplier:multiplier];
    return result_;
}

void* C_NSLayoutXAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingAfterAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutXAxisAnchor* nSLayoutXAxisAnchor = (NSLayoutXAxisAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutXAxisAnchor constraintLessThanOrEqualToSystemSpacingAfterAnchor:(NSLayoutXAxisAnchor*)anchor multiplier:multiplier];
    return result_;
}

void* C_NSLayoutXAxisAnchor_AnchorWithOffsetToAnchor(void* ptr, void* otherAnchor) {
    NSLayoutXAxisAnchor* nSLayoutXAxisAnchor = (NSLayoutXAxisAnchor*)ptr;
    NSLayoutDimension* result_ = [nSLayoutXAxisAnchor anchorWithOffsetToAnchor:(NSLayoutXAxisAnchor*)otherAnchor];
    return result_;
}

void* C_LayoutYAxisAnchor_Alloc() {
    return [NSLayoutYAxisAnchor alloc];
}

void* C_NSLayoutYAxisAnchor_AllocLayoutYAxisAnchor() {
    NSLayoutYAxisAnchor* result_ = [NSLayoutYAxisAnchor alloc];
    return result_;
}

void* C_NSLayoutYAxisAnchor_Init(void* ptr) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSLayoutYAxisAnchor init];
    return result_;
}

void* C_NSLayoutYAxisAnchor_NewLayoutYAxisAnchor() {
    NSLayoutYAxisAnchor* result_ = [NSLayoutYAxisAnchor new];
    return result_;
}

void* C_NSLayoutYAxisAnchor_Autorelease(void* ptr) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSLayoutYAxisAnchor autorelease];
    return result_;
}

void* C_NSLayoutYAxisAnchor_Retain(void* ptr) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSLayoutYAxisAnchor retain];
    return result_;
}

void* C_NSLayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutYAxisAnchor constraintEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
    return result_;
}

void* C_NSLayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutYAxisAnchor constraintGreaterThanOrEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
    return result_;
}

void* C_NSLayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutConstraint* result_ = [nSLayoutYAxisAnchor constraintLessThanOrEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
    return result_;
}

void* C_NSLayoutYAxisAnchor_AnchorWithOffsetToAnchor(void* ptr, void* otherAnchor) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutDimension* result_ = [nSLayoutYAxisAnchor anchorWithOffsetToAnchor:(NSLayoutYAxisAnchor*)otherAnchor];
    return result_;
}
