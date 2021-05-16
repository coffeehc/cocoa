#import <Appkit/Appkit.h>
#import "layout_guide.h"

void* C_LayoutGuide_Alloc() {
    return [NSLayoutGuide alloc];
}

void* C_NSLayoutGuide_Init(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutGuide* result = [nSLayoutGuide init];
    return result;
}

Array C_NSLayoutGuide_ConstraintsAffectingLayoutForOrientation(void* ptr, int orientation) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSArray* result = [nSLayoutGuide constraintsAffectingLayoutForOrientation:orientation];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void* C_NSLayoutGuide_Identifier(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSUserInterfaceItemIdentifier result = [nSLayoutGuide identifier];
    return result;
}

void C_NSLayoutGuide_SetIdentifier(void* ptr, void* value) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    [nSLayoutGuide setIdentifier:(NSString*)value];
}

CGRect C_NSLayoutGuide_Frame(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSRect result = [nSLayoutGuide frame];
    return result;
}

void* C_NSLayoutGuide_OwningView(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSView* result = [nSLayoutGuide owningView];
    return result;
}

void C_NSLayoutGuide_SetOwningView(void* ptr, void* value) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    [nSLayoutGuide setOwningView:(NSView*)value];
}

void* C_NSLayoutGuide_BottomAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutYAxisAnchor* result = [nSLayoutGuide bottomAnchor];
    return result;
}

void* C_NSLayoutGuide_CenterXAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutXAxisAnchor* result = [nSLayoutGuide centerXAnchor];
    return result;
}

void* C_NSLayoutGuide_CenterYAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutYAxisAnchor* result = [nSLayoutGuide centerYAnchor];
    return result;
}

void* C_NSLayoutGuide_HeightAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutDimension* result = [nSLayoutGuide heightAnchor];
    return result;
}

void* C_NSLayoutGuide_LeadingAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutXAxisAnchor* result = [nSLayoutGuide leadingAnchor];
    return result;
}

void* C_NSLayoutGuide_LeftAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutXAxisAnchor* result = [nSLayoutGuide leftAnchor];
    return result;
}

void* C_NSLayoutGuide_RightAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutXAxisAnchor* result = [nSLayoutGuide rightAnchor];
    return result;
}

void* C_NSLayoutGuide_TopAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutYAxisAnchor* result = [nSLayoutGuide topAnchor];
    return result;
}

void* C_NSLayoutGuide_TrailingAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutXAxisAnchor* result = [nSLayoutGuide trailingAnchor];
    return result;
}

void* C_NSLayoutGuide_WidthAnchor(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    NSLayoutDimension* result = [nSLayoutGuide widthAnchor];
    return result;
}

bool C_NSLayoutGuide_HasAmbiguousLayout(void* ptr) {
    NSLayoutGuide* nSLayoutGuide = (NSLayoutGuide*)ptr;
    BOOL result = [nSLayoutGuide hasAmbiguousLayout];
    return result;
}
