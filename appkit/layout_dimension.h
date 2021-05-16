#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_LayoutDimension_Alloc();

void* C_NSLayoutDimension_Init(void* ptr);
void* C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier(void* ptr, void* anchor, double m);
void* C_NSLayoutDimension_ConstraintEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c);
void* C_NSLayoutDimension_ConstraintEqualToConstant(void* ptr, double c);
void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier(void* ptr, void* anchor, double m);
void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c);
void* C_NSLayoutDimension_ConstraintGreaterThanOrEqualToConstant(void* ptr, double c);
void* C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier(void* ptr, void* anchor, double m);
void* C_NSLayoutDimension_ConstraintLessThanOrEqualToAnchor_Multiplier_Constant(void* ptr, void* anchor, double m, double c);
void* C_NSLayoutDimension_ConstraintLessThanOrEqualToConstant(void* ptr, double c);
