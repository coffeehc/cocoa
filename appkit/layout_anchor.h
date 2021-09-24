#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_LayoutAnchor_Alloc();

void* C_NSLayoutAnchor_AllocLayoutAnchor();
void* C_NSLayoutAnchor_Init(void* ptr);
void* C_NSLayoutAnchor_NewLayoutAnchor();
void* C_NSLayoutAnchor_Autorelease(void* ptr);
void* C_NSLayoutAnchor_Retain(void* ptr);
void* C_NSLayoutAnchor_ConstraintEqualToAnchor(void* ptr, void* anchor);
void* C_NSLayoutAnchor_ConstraintEqualToAnchor_Constant(void* ptr, void* anchor, double c);
void* C_NSLayoutAnchor_ConstraintGreaterThanOrEqualToAnchor(void* ptr, void* anchor);
void* C_NSLayoutAnchor_ConstraintGreaterThanOrEqualToAnchor_Constant(void* ptr, void* anchor, double c);
void* C_NSLayoutAnchor_ConstraintLessThanOrEqualToAnchor(void* ptr, void* anchor);
void* C_NSLayoutAnchor_ConstraintLessThanOrEqualToAnchor_Constant(void* ptr, void* anchor, double c);
Array C_NSLayoutAnchor_ConstraintsAffectingLayout(void* ptr);
bool C_NSLayoutAnchor_HasAmbiguousLayout(void* ptr);
void* C_NSLayoutAnchor_Name(void* ptr);
void* C_NSLayoutAnchor_Item(void* ptr);

void* C_LayoutDimension_Alloc();

void* C_NSLayoutDimension_AllocLayoutDimension();
void* C_NSLayoutDimension_Init(void* ptr);
void* C_NSLayoutDimension_NewLayoutDimension();
void* C_NSLayoutDimension_Autorelease(void* ptr);
void* C_NSLayoutDimension_Retain(void* ptr);
void* C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier(void* ptr, void* anchor, double m);
void* C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c);
void* C_NSLayoutDimension_ConstraintEqualToConstant(void* ptr, double c);
void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier(void* ptr, void* anchor, double m);
void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c);
void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToConstant(void* ptr, double c);
void* C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier(void* ptr, void* anchor, double m);
void* C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c);
void* C_NSLayoutDimension_ConstraintLessThanOrEqualToConstant(void* ptr, double c);

void* C_LayoutXAxisAnchor_Alloc();

void* C_NSLayoutXAxisAnchor_AllocLayoutXAxisAnchor();
void* C_NSLayoutXAxisAnchor_Init(void* ptr);
void* C_NSLayoutXAxisAnchor_NewLayoutXAxisAnchor();
void* C_NSLayoutXAxisAnchor_Autorelease(void* ptr);
void* C_NSLayoutXAxisAnchor_Retain(void* ptr);
void* C_NSLayoutXAxisAnchor_ConstraintEqualToSystemSpacingAfterAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutXAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutXAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingAfterAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutXAxisAnchor_AnchorWithOffsetToAnchor(void* ptr, void* otherAnchor);

void* C_LayoutYAxisAnchor_Alloc();

void* C_NSLayoutYAxisAnchor_AllocLayoutYAxisAnchor();
void* C_NSLayoutYAxisAnchor_Init(void* ptr);
void* C_NSLayoutYAxisAnchor_NewLayoutYAxisAnchor();
void* C_NSLayoutYAxisAnchor_Autorelease(void* ptr);
void* C_NSLayoutYAxisAnchor_Retain(void* ptr);
void* C_NSLayoutYAxisAnchor_ConstraintEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutYAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutYAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingBelowAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutYAxisAnchor_AnchorWithOffsetToAnchor(void* ptr, void* otherAnchor);
