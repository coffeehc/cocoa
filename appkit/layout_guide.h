#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_LayoutGuide_Alloc();

void* C_NSLayoutGuide_AllocLayoutGuide();
void* C_NSLayoutGuide_Init(void* ptr);
void* C_NSLayoutGuide_NewLayoutGuide();
void* C_NSLayoutGuide_Autorelease(void* ptr);
void* C_NSLayoutGuide_Retain(void* ptr);
Array C_NSLayoutGuide_ConstraintsAffectingLayoutForOrientation(void* ptr, int orientation);
void* C_NSLayoutGuide_Identifier(void* ptr);
void C_NSLayoutGuide_SetIdentifier(void* ptr, void* value);
CGRect C_NSLayoutGuide_Frame(void* ptr);
void* C_NSLayoutGuide_OwningView(void* ptr);
void C_NSLayoutGuide_SetOwningView(void* ptr, void* value);
void* C_NSLayoutGuide_BottomAnchor(void* ptr);
void* C_NSLayoutGuide_CenterXAnchor(void* ptr);
void* C_NSLayoutGuide_CenterYAnchor(void* ptr);
void* C_NSLayoutGuide_HeightAnchor(void* ptr);
void* C_NSLayoutGuide_LeadingAnchor(void* ptr);
void* C_NSLayoutGuide_LeftAnchor(void* ptr);
void* C_NSLayoutGuide_RightAnchor(void* ptr);
void* C_NSLayoutGuide_TopAnchor(void* ptr);
void* C_NSLayoutGuide_TrailingAnchor(void* ptr);
void* C_NSLayoutGuide_WidthAnchor(void* ptr);
bool C_NSLayoutGuide_HasAmbiguousLayout(void* ptr);
