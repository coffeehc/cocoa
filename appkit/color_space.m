#import <Appkit/Appkit.h>
#import "color_space.h"

void* C_ColorSpace_Alloc() {
    return [NSColorSpace alloc];
}

void* C_NSColorSpace_InitWithCGColorSpace(void* ptr, void* cgColorSpace) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSColorSpace* result_ = [nSColorSpace initWithCGColorSpace:(CGColorSpaceRef)cgColorSpace];
    return result_;
}

void* C_NSColorSpace_InitWithICCProfileData(void* ptr, void* iccData) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSColorSpace* result_ = [nSColorSpace initWithICCProfileData:(NSData*)iccData];
    return result_;
}

void* C_NSColorSpace_AllocColorSpace() {
    NSColorSpace* result_ = [NSColorSpace alloc];
    return result_;
}

void* C_NSColorSpace_Init(void* ptr) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSColorSpace* result_ = [nSColorSpace init];
    return result_;
}

void* C_NSColorSpace_NewColorSpace() {
    NSColorSpace* result_ = [NSColorSpace new];
    return result_;
}

void* C_NSColorSpace_Autorelease(void* ptr) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSColorSpace* result_ = [nSColorSpace autorelease];
    return result_;
}

void* C_NSColorSpace_Retain(void* ptr) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSColorSpace* result_ = [nSColorSpace retain];
    return result_;
}

Array C_NSColorSpace_AvailableColorSpacesWithModel(int model) {
    NSArray* result_ = [NSColorSpace availableColorSpacesWithModel:model];
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

void* C_NSColorSpace_DeviceRGBColorSpace() {
    NSColorSpace* result_ = [NSColorSpace deviceRGBColorSpace];
    return result_;
}

void* C_NSColorSpace_GenericRGBColorSpace() {
    NSColorSpace* result_ = [NSColorSpace genericRGBColorSpace];
    return result_;
}

void* C_NSColorSpace_DeviceCMYKColorSpace() {
    NSColorSpace* result_ = [NSColorSpace deviceCMYKColorSpace];
    return result_;
}

void* C_NSColorSpace_GenericCMYKColorSpace() {
    NSColorSpace* result_ = [NSColorSpace genericCMYKColorSpace];
    return result_;
}

void* C_NSColorSpace_DeviceGrayColorSpace() {
    NSColorSpace* result_ = [NSColorSpace deviceGrayColorSpace];
    return result_;
}

void* C_NSColorSpace_GenericGrayColorSpace() {
    NSColorSpace* result_ = [NSColorSpace genericGrayColorSpace];
    return result_;
}

void* C_NSColorSpace_SRGBColorSpace() {
    NSColorSpace* result_ = [NSColorSpace sRGBColorSpace];
    return result_;
}

void* C_NSColorSpace_ExtendedSRGBColorSpace() {
    NSColorSpace* result_ = [NSColorSpace extendedSRGBColorSpace];
    return result_;
}

void* C_NSColorSpace_DisplayP3ColorSpace() {
    NSColorSpace* result_ = [NSColorSpace displayP3ColorSpace];
    return result_;
}

void* C_NSColorSpace_GenericGamma22GrayColorSpace() {
    NSColorSpace* result_ = [NSColorSpace genericGamma22GrayColorSpace];
    return result_;
}

void* C_NSColorSpace_ExtendedGenericGamma22GrayColorSpace() {
    NSColorSpace* result_ = [NSColorSpace extendedGenericGamma22GrayColorSpace];
    return result_;
}

void* C_NSColorSpace_AdobeRGB1998ColorSpace() {
    NSColorSpace* result_ = [NSColorSpace adobeRGB1998ColorSpace];
    return result_;
}

void* C_NSColorSpace_CGColorSpace(void* ptr) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    CGColorSpaceRef result_ = [nSColorSpace CGColorSpace];
    return result_;
}

int C_NSColorSpace_ColorSpaceModel(void* ptr) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSColorSpaceModel result_ = [nSColorSpace colorSpaceModel];
    return result_;
}

void* C_NSColorSpace_ICCProfileData(void* ptr) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSData* result_ = [nSColorSpace ICCProfileData];
    return result_;
}

void* C_NSColorSpace_LocalizedName(void* ptr) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSString* result_ = [nSColorSpace localizedName];
    return result_;
}

int C_NSColorSpace_NumberOfColorComponents(void* ptr) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSInteger result_ = [nSColorSpace numberOfColorComponents];
    return result_;
}
