#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ColorSpace_Alloc();

void* C_NSColorSpace_InitWithCGColorSpace(void* ptr, void* cgColorSpace);
void* C_NSColorSpace_InitWithICCProfileData(void* ptr, Array iccData);
void* C_NSColorSpace_Init(void* ptr);
Array C_NSColorSpace_AvailableColorSpacesWithModel(int model);
void* C_NSColorSpace_DeviceRGBColorSpace();
void* C_NSColorSpace_GenericRGBColorSpace();
void* C_NSColorSpace_DeviceCMYKColorSpace();
void* C_NSColorSpace_GenericCMYKColorSpace();
void* C_NSColorSpace_DeviceGrayColorSpace();
void* C_NSColorSpace_GenericGrayColorSpace();
void* C_NSColorSpace_SRGBColorSpace();
void* C_NSColorSpace_ExtendedSRGBColorSpace();
void* C_NSColorSpace_DisplayP3ColorSpace();
void* C_NSColorSpace_GenericGamma22GrayColorSpace();
void* C_NSColorSpace_ExtendedGenericGamma22GrayColorSpace();
void* C_NSColorSpace_AdobeRGB1998ColorSpace();
void* C_NSColorSpace_CGColorSpace(void* ptr);
int C_NSColorSpace_ColorSpaceModel(void* ptr);
Array C_NSColorSpace_ICCProfileData(void* ptr);
void* C_NSColorSpace_LocalizedName(void* ptr);
int C_NSColorSpace_NumberOfColorComponents(void* ptr);
