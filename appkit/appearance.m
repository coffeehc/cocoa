#import "appearance.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSAppearance.h>

void* C_Appearance_Alloc() {
    return [NSAppearance alloc];
}

void* C_NSAppearance_InitWithAppearanceNamed_Bundle(void* ptr, void* name, void* bundle) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    NSAppearance* result_ = [nSAppearance initWithAppearanceNamed:(NSString*)name bundle:(NSBundle*)bundle];
    return result_;
}

void* C_NSAppearance_InitWithCoder(void* ptr, void* coder) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    NSAppearance* result_ = [nSAppearance initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSAppearance_AllocAppearance() {
    NSAppearance* result_ = [NSAppearance alloc];
    return result_;
}

void* C_NSAppearance_Init(void* ptr) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    NSAppearance* result_ = [nSAppearance init];
    return result_;
}

void* C_NSAppearance_NewAppearance() {
    NSAppearance* result_ = [NSAppearance new];
    return result_;
}

void* C_NSAppearance_Autorelease(void* ptr) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    NSAppearance* result_ = [nSAppearance autorelease];
    return result_;
}

void* C_NSAppearance_Retain(void* ptr) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    NSAppearance* result_ = [nSAppearance retain];
    return result_;
}

void* C_NSAppearance_AppearanceNamed(void* name) {
    NSAppearance* result_ = [NSAppearance appearanceNamed:(NSString*)name];
    return result_;
}

void* C_NSAppearance_BestMatchFromAppearancesWithNames(void* ptr, Array appearances) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    NSMutableArray* objcAppearances = [NSMutableArray arrayWithCapacity:appearances.len];
    if (appearances.len > 0) {
    	void** appearancesData = (void**)appearances.data;
    	for (int i = 0; i < appearances.len; i++) {
    		void* p = appearancesData[i];
    		[objcAppearances addObject:(NSString*)p];
    	}
    }
    NSAppearanceName result_ = [nSAppearance bestMatchFromAppearancesWithNames:objcAppearances];
    return result_;
}

void* C_NSAppearance_Name(void* ptr) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    NSAppearanceName result_ = [nSAppearance name];
    return result_;
}

void* C_NSAppearance_CurrentDrawingAppearance() {
    NSAppearance* result_ = [NSAppearance currentDrawingAppearance];
    return result_;
}

bool C_NSAppearance_AllowsVibrancy(void* ptr) {
    NSAppearance* nSAppearance = (NSAppearance*)ptr;
    BOOL result_ = [nSAppearance allowsVibrancy];
    return result_;
}
