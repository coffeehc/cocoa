#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_BitmapImageRep_Alloc();

void* C_NSBitmapImageRep_BitmapImageRep_ImageRepWithData(void* data);
void* C_NSBitmapImageRep_InitWithCGImage(void* ptr, void* cgImage);
void* C_NSBitmapImageRep_InitWithData(void* ptr, void* data);
void* C_NSBitmapImageRep_InitForIncrementalLoad(void* ptr);
void* C_NSBitmapImageRep_Init(void* ptr);
void* C_NSBitmapImageRep_InitWithCoder(void* ptr, void* coder);
void* C_NSBitmapImageRep_AllocBitmapImageRep();
void* C_NSBitmapImageRep_NewBitmapImageRep();
void* C_NSBitmapImageRep_Autorelease(void* ptr);
void* C_NSBitmapImageRep_Retain(void* ptr);
Array C_NSBitmapImageRep_BitmapImageRep_ImageRepsWithData(void* data);
void C_NSBitmapImageRep_ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(void* ptr, double midPoint, void* midPointColor, void* shadowColor, void* lightColor);
void* C_NSBitmapImageRep_BitmapImageRep_TIFFRepresentationOfImageRepsInArray(Array array);
void* C_NSBitmapImageRep_BitmapImageRep_TIFFRepresentationOfImageRepsInArray_UsingCompression_Factor(Array array, unsigned int comp, float factor);
void* C_NSBitmapImageRep_TIFFRepresentationUsingCompression_Factor(void* ptr, unsigned int comp, float factor);
void* C_NSBitmapImageRep_BitmapImageRep_RepresentationOfImageRepsInArray_UsingType_Properties(Array imageReps, unsigned int storageType, Dictionary properties);
void* C_NSBitmapImageRep_RepresentationUsingType_Properties(void* ptr, unsigned int storageType, Dictionary properties);
void* C_NSBitmapImageRep_BitmapImageRep_LocalizedNameForTIFFCompressionType(unsigned int compression);
bool C_NSBitmapImageRep_CanBeCompressedUsing(void* ptr, unsigned int compression);
void C_NSBitmapImageRep_SetCompression_Factor(void* ptr, unsigned int compression, float factor);
void C_NSBitmapImageRep_SetProperty_WithValue(void* ptr, void* property, void* value);
void* C_NSBitmapImageRep_ValueForProperty(void* ptr, void* property);
int C_NSBitmapImageRep_IncrementalLoadFromData_Complete(void* ptr, void* data, bool complete);
void C_NSBitmapImageRep_SetColor_AtX_Y(void* ptr, void* color, int x, int y);
void* C_NSBitmapImageRep_ColorAtX_Y(void* ptr, int x, int y);
void* C_NSBitmapImageRep_BitmapImageRepByConvertingToColorSpace_RenderingIntent(void* ptr, void* targetSpace, int renderingIntent);
void* C_NSBitmapImageRep_BitmapImageRepByRetaggingWithColorSpace(void* ptr, void* newSpace);
unsigned int C_NSBitmapImageRep_BitmapFormat(void* ptr);
int C_NSBitmapImageRep_BitsPerPixel(void* ptr);
int C_NSBitmapImageRep_BytesPerPlane(void* ptr);
int C_NSBitmapImageRep_BytesPerRow(void* ptr);
bool C_NSBitmapImageRep_IsPlanar(void* ptr);
int C_NSBitmapImageRep_NumberOfPlanes(void* ptr);
int C_NSBitmapImageRep_SamplesPerPixel(void* ptr);
void* C_NSBitmapImageRep_TIFFRepresentation(void* ptr);
void* C_NSBitmapImageRep_CGImage(void* ptr);
void* C_NSBitmapImageRep_ColorSpace(void* ptr);
