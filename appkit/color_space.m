#import <Appkit/Appkit.h>
#import "color_space.h"

void* C_ColorSpace_Alloc() {
    return [NSColorSpace alloc];
}

void* C_NSColorSpace_InitWithICCProfileData(void* ptr, Array iccData) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSColorSpace* result_ = [nSColorSpace initWithICCProfileData:[[NSData alloc] initWithBytes:(Byte *)iccData.data length:iccData.len]];
    return result_;
}

void* C_NSColorSpace_Init(void* ptr) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSColorSpace* result_ = [nSColorSpace init];
    return result_;
}

Array C_NSColorSpace_ColorSpace_AvailableColorSpacesWithModel(int model) {
    NSArray* result_ = [NSColorSpace availableColorSpacesWithModel:model];
    int result_count = [result_ count];
    void** result_Data = malloc(result_count * sizeof(void*));
    for (int i = 0; i < result_count; i++) {
    	 void* p = [result_ objectAtIndex:i];
    	 result_Data[i] = p;
    }
    Array result_Array;
    result_Array.data = result_Data;
    result_Array.len = result_count;
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

int C_NSColorSpace_ColorSpaceModel(void* ptr) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSColorSpaceModel result_ = [nSColorSpace colorSpaceModel];
    return result_;
}

Array C_NSColorSpace_ICCProfileData(void* ptr) {
    NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
    NSData* result_ = [nSColorSpace ICCProfileData];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
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
