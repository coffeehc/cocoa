#import <Appkit/Appkit.h>
#import "layout_guide.h"

void* C_LayoutGuide_Alloc() {
    return [NSLayoutGuide alloc];
}

void* C_NSLayoutGuide_AllocLayoutGuide() {
    NSLayoutGuide* result_ = [NSLayoutGuide alloc];
    return result_;
}

void* C_NSLayoutGuide_Init(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutGuide* result_ = [nSLayoutGuide init];
    return result_;
}

void* C_NSLayoutGuide_NewLayoutGuide() {
    NSLayoutGuide* result_ = [NSLayoutGuide new];
    return result_;
}

void* C_NSLayoutGuide_Autorelease(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutGuide* result_ = [nSLayoutGuide autorelease];
    return result_;
}

void* C_NSLayoutGuide_Retain(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutGuide* result_ = [nSLayoutGuide retain];
    return result_;
}

Array C_NSLayoutGuide_ConstraintsAffectingLayoutForOrientation(void* ptr, int orientation) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSArray* result_ = [nSLayoutGuide constraintsAffectingLayoutForOrientation:orientation];
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

void* C_NSLayoutGuide_Identifier(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSUserInterfaceItemIdentifier result_ = [nSLayoutGuide identifier];
    return result_;
}

void C_NSLayoutGuide_SetIdentifier(void* ptr, void* value) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    [nSLayoutGuide setIdentifier:(NSString*)value];
}

CGRect C_NSLayoutGuide_Frame(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSRect result_ = [nSLayoutGuide frame];
    return result_;
}

void* C_NSLayoutGuide_OwningView(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSView* result_ = [nSLayoutGuide owningView];
    return result_;
}

void C_NSLayoutGuide_SetOwningView(void* ptr, void* value) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    [nSLayoutGuide setOwningView:(NSView*)value];
}

void* C_NSLayoutGuide_BottomAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSLayoutGuide bottomAnchor];
    return result_;
}

void* C_NSLayoutGuide_CenterXAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSLayoutGuide centerXAnchor];
    return result_;
}

void* C_NSLayoutGuide_CenterYAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSLayoutGuide centerYAnchor];
    return result_;
}

void* C_NSLayoutGuide_HeightAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutDimension* result_ = [nSLayoutGuide heightAnchor];
    return result_;
}

void* C_NSLayoutGuide_LeadingAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSLayoutGuide leadingAnchor];
    return result_;
}

void* C_NSLayoutGuide_LeftAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSLayoutGuide leftAnchor];
    return result_;
}

void* C_NSLayoutGuide_RightAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSLayoutGuide rightAnchor];
    return result_;
}

void* C_NSLayoutGuide_TopAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutYAxisAnchor* result_ = [nSLayoutGuide topAnchor];
    return result_;
}

void* C_NSLayoutGuide_TrailingAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutXAxisAnchor* result_ = [nSLayoutGuide trailingAnchor];
    return result_;
}

void* C_NSLayoutGuide_WidthAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutDimension* result_ = [nSLayoutGuide widthAnchor];
    return result_;
}

bool C_NSLayoutGuide_HasAmbiguousLayout(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    BOOL result_ = [nSLayoutGuide hasAmbiguousLayout];
    return result_;
}
