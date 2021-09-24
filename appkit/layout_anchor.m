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
