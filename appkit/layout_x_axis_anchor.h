#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_LayoutXAxisAnchor_Alloc();

void* C_NSLayoutXAxisAnchor_Init(void* ptr);
void* C_NSLayoutXAxisAnchor_ConstraintEqualToSystemSpacingAfterAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutXAxisAnchor_ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutXAxisAnchor_ConstraintLessThanOrEqualToSystemSpacingAfterAnchor_Multiplier(void* ptr, void* anchor, double multiplier);
void* C_NSLayoutXAxisAnchor_AnchorWithOffsetToAnchor(void* ptr, void* otherAnchor);
