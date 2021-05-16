#import <Appkit/Appkit.h>
#import "layout_dimension.h"

void* C_LayoutDimension_Alloc() {
    return [NSLayoutDimension alloc];
}

void* C_NSLayoutDimension_Init(void* ptr) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutDimension* result = [nSLayoutDimension init];
    return result;
}

void* C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier(void* ptr, void* anchor, double m) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result = [nSLayoutDimension constraintEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m];
    return result;
}

void* C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result = [nSLayoutDimension constraintEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m constant:c];
    return result;
}

void* C_NSLayoutDimension_ConstraintEqualToConstant(void* ptr, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result = [nSLayoutDimension constraintEqualToConstant:c];
    return result;
}

void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier(void* ptr, void* anchor, double m) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result = [nSLayoutDimension constraintGreaterThanOrEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m];
    return result;
}

void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result = [nSLayoutDimension constraintGreaterThanOrEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m constant:c];
    return result;
}

void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToConstant(void* ptr, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result = [nSLayoutDimension constraintGreaterThanOrEqualToConstant:c];
    return result;
}

void* C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier(void* ptr, void* anchor, double m) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result = [nSLayoutDimension constraintLessThanOrEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m];
    return result;
}

void* C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result = [nSLayoutDimension constraintLessThanOrEqualToAnchor:(NSLayoutDimension*)anchor multiplier:m constant:c];
    return result;
}

void* C_NSLayoutDimension_ConstraintLessThanOrEqualToConstant(void* ptr, double c) {
    NSLayoutDimension* nSLayoutDimension = (NSLayoutDimension*)ptr;
    NSLayoutConstraint* result = [nSLayoutDimension constraintLessThanOrEqualToConstant:c];
    return result;
}
