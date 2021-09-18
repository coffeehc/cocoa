#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_LayoutConstraint_Alloc();

void* C_NSLayoutConstraint_LayoutConstraint_ConstraintWithItem_Attribute_RelatedBy_ToItem_Attribute_Multiplier_Constant(void* view1, int attr1, int relation, void* view2, int attr2, double multiplier, double c);
void* C_NSLayoutConstraint_AllocLayoutConstraint();
void* C_NSLayoutConstraint_Init(void* ptr);
void* C_NSLayoutConstraint_NewLayoutConstraint();
void* C_NSLayoutConstraint_Autorelease(void* ptr);
void* C_NSLayoutConstraint_Retain(void* ptr);
Array C_NSLayoutConstraint_LayoutConstraint_ConstraintsWithVisualFormat_Options_Metrics_Views(void* format, unsigned int opts, Dictionary metrics, Dictionary views);
void C_NSLayoutConstraint_LayoutConstraint_ActivateConstraints(Array constraints);
void C_NSLayoutConstraint_LayoutConstraint_DeactivateConstraints(Array constraints);
bool C_NSLayoutConstraint_IsActive(void* ptr);
void C_NSLayoutConstraint_SetActive(void* ptr, bool value);
void* C_NSLayoutConstraint_FirstItem(void* ptr);
int C_NSLayoutConstraint_FirstAttribute(void* ptr);
int C_NSLayoutConstraint_Relation(void* ptr);
void* C_NSLayoutConstraint_SecondItem(void* ptr);
int C_NSLayoutConstraint_SecondAttribute(void* ptr);
double C_NSLayoutConstraint_Multiplier(void* ptr);
double C_NSLayoutConstraint_Constant(void* ptr);
void C_NSLayoutConstraint_SetConstant(void* ptr, double value);
void* C_NSLayoutConstraint_FirstAnchor(void* ptr);
void* C_NSLayoutConstraint_SecondAnchor(void* ptr);
float C_NSLayoutConstraint_Priority(void* ptr);
void C_NSLayoutConstraint_SetPriority(void* ptr, float value);
void* C_NSLayoutConstraint_Identifier(void* ptr);
void C_NSLayoutConstraint_SetIdentifier(void* ptr, void* value);
bool C_NSLayoutConstraint_ShouldBeArchived(void* ptr);
void C_NSLayoutConstraint_SetShouldBeArchived(void* ptr, bool value);
