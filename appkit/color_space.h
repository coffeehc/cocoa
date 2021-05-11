#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_ColorSpace_Alloc();

void* C_NSColorSpace_InitWithCGColorSpace(void* ptr, CGColorSpaceRef cgColorSpace);
void* C_NSColorSpace_InitWithICCProfileData(void* ptr, Array iccData);
void* C_NSColorSpace_Init(void* ptr);
CGColorSpaceRef C_NSColorSpace_CGColorSpace(void* ptr);
Array C_NSColorSpace_ICCProfileData(void* ptr);
void* C_NSColorSpace_LocalizedName(void* ptr);
int C_NSColorSpace_NumberOfColorComponents(void* ptr);
