#import <Appkit/Appkit.h>
#import "layout_y_axis_anchor.h"

void* C_LayoutYAxisAnchor_Alloc() {
    return [NSLayoutYAxisAnchor alloc];
}

void* C_NSLayoutYAxisAnchor_Init(void* ptr) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutYAxisAnchor* result = [nSLayoutYAxisAnchor init];
    return result;
}

void* C_NSLayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutConstraint* result = [nSLayoutYAxisAnchor constraintEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
    return result;
}

void* C_NSLayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutConstraint* result = [nSLayoutYAxisAnchor constraintGreaterThanOrEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
    return result;
}

void* C_NSLayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutConstraint* result = [nSLayoutYAxisAnchor constraintLessThanOrEqualToSystemSpacingBelowAnchor:(NSLayoutYAxisAnchor*)anchor multiplier:multiplier];
    return result;
}

void* C_NSLayoutYAxisAnchor_AnchorWithOffsetToAnchor(void* ptr, void* otherAnchor) {
    NSLayoutYAxisAnchor* nSLayoutYAxisAnchor = (NSLayoutYAxisAnchor*)ptr;
    NSLayoutDimension* result = [nSLayoutYAxisAnchor anchorWithOffsetToAnchor:(NSLayoutYAxisAnchor*)otherAnchor];
    return result;
}
