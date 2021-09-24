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
