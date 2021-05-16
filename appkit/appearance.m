#import <Appkit/Appkit.h>
#import "appearance.h"

void* C_Appearance_Alloc() {
    return [NSAppearance alloc];
}

void* C_NSAppearance_InitWithAppearanceNamed_Bundle(void* ptr, void* name, void* bundle) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    NSAppearance* result = [nSAppearance initWithAppearanceNamed:(NSString*)name bundle:(NSBundle*)bundle];
    return result;
}

void* C_NSAppearance_InitWithCoder(void* ptr, void* coder) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    NSAppearance* result = [nSAppearance initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSAppearance_Init(void* ptr) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    NSAppearance* result = [nSAppearance init];
    return result;
}

void* C_NSAppearance_AppearanceNamed(void* name) {
    NSAppearance* result = [NSAppearance appearanceNamed:(NSString*)name];
    return result;
}

void* C_NSAppearance_Name(void* ptr) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    NSAppearanceName result = [nSAppearance name];
    return result;
}

void* C_NSAppearance_CurrentDrawingAppearance() {
    NSAppearance* result = [NSAppearance currentDrawingAppearance];
    return result;
}

bool C_NSAppearance_AllowsVibrancy(void* ptr) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    BOOL result = [nSAppearance allowsVibrancy];
    return result;
}
