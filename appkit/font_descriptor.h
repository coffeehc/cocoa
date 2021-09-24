#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_FontDescriptor_Alloc();

void* C_NSFontDescriptor_InitWithFontAttributes(void* ptr, Dictionary attributes);
void* C_NSFontDescriptor_FontDescriptorWithDesign(void* ptr, void* design);
void* C_NSFontDescriptor_AllocFontDescriptor();
void* C_NSFontDescriptor_Init(void* ptr);
void* C_NSFontDescriptor_NewFontDescriptor();
void* C_NSFontDescriptor_Autorelease(void* ptr);
void* C_NSFontDescriptor_Retain(void* ptr);
void* C_NSFontDescriptor_PreferredFontDescriptorForTextStyle_Options(void* style, Dictionary options);
void* C_NSFontDescriptor_FontDescriptorWithFontAttributes(Dictionary attributes);
void* C_NSFontDescriptor_FontDescriptorWithName_Matrix(void* fontName, void* matrix);
void* C_NSFontDescriptor_FontDescriptorWithName_Size(void* fontName, double size);
void* C_NSFontDescriptor_FontDescriptorByAddingAttributes(void* ptr, Dictionary attributes);
void* C_NSFontDescriptor_FontDescriptorWithFace(void* ptr, void* newFace);
void* C_NSFontDescriptor_FontDescriptorWithFamily(void* ptr, void* newFamily);
void* C_NSFontDescriptor_FontDescriptorWithMatrix(void* ptr, void* matrix);
void* C_NSFontDescriptor_FontDescriptorWithSize(void* ptr, double newPointSize);
void* C_NSFontDescriptor_FontDescriptorWithSymbolicTraits(void* ptr, uint32_t symbolicTraits);
Array C_NSFontDescriptor_MatchingFontDescriptorsWithMandatoryKeys(void* ptr, void* mandatoryKeys);
void* C_NSFontDescriptor_MatchingFontDescriptorWithMandatoryKeys(void* ptr, void* mandatoryKeys);
void* C_NSFontDescriptor_ObjectForKey(void* ptr, void* attribute);
Dictionary C_NSFontDescriptor_FontAttributes(void* ptr);
void* C_NSFontDescriptor_Matrix(void* ptr);
double C_NSFontDescriptor_PointSize(void* ptr);
void* C_NSFontDescriptor_PostscriptName(void* ptr);
uint32_t C_NSFontDescriptor_SymbolicTraits(void* ptr);
bool C_NSFontDescriptor_RequiresFontAssetRequest(void* ptr);
