#import <Appkit/Appkit.h>
#import "color_space.h"

void* C_ColorSpace_Alloc() {
	return [NSColorSpace alloc];
}

void* C_NSColorSpace_InitWithCGColorSpace(void* ptr, CGColorSpaceRef cgColorSpace) {
	NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
	NSColorSpace* result = [nSColorSpace initWithCGColorSpace:cgColorSpace];
	return result;
}

void* C_NSColorSpace_InitWithICCProfileData(void* ptr, Array iccData) {
	NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
	NSColorSpace* result = [nSColorSpace initWithICCProfileData:[[NSData alloc] initWithBytes:(Byte *)iccData.data length:iccData.len]];
	return result;
}

void* C_NSColorSpace_Init(void* ptr) {
	NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
	NSColorSpace* result = [nSColorSpace init];
	return result;
}

CGColorSpaceRef C_NSColorSpace_CGColorSpace(void* ptr) {
	NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
	CGColorSpaceRef result = [nSColorSpace CGColorSpace];
	return result;
}

Array C_NSColorSpace_ICCProfileData(void* ptr) {
	NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
	NSData* result = [nSColorSpace ICCProfileData];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

void* C_NSColorSpace_LocalizedName(void* ptr) {
	NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
	NSString* result = [nSColorSpace localizedName];
	return result;
}

int C_NSColorSpace_NumberOfColorComponents(void* ptr) {
	NSColorSpace* nSColorSpace = (NSColorSpace*)ptr;
	int result = [nSColorSpace numberOfColorComponents];
	return result;
}
