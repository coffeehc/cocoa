#import <Appkit/Appkit.h>
#import "layout_x_axis_anchor.h"

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
