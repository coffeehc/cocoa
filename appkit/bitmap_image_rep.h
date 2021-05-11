#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_BitmapImageRep_Alloc();

void* C_NSBitmapImageRep_InitWithCGImage(void* ptr, CGImageRef cgImage);
void* C_NSBitmapImageRep_InitWithData(void* ptr, Array data);
void* C_NSBitmapImageRep_InitForIncrementalLoad(void* ptr);
void* C_NSBitmapImageRep_Init(void* ptr);
void* C_NSBitmapImageRep_InitWithCoder(void* ptr, void* coder);
void* C_NSBitmapImageRep_BitmapImageRepImageRepWithData(Array data);
void C_NSBitmapImageRep_ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(void* ptr, double midPoint, void* midPointColor, void* shadowColor, void* lightColor);
void C_NSBitmapImageRep_SetProperty_WithValue(void* ptr, void* property, void* value);
void* C_NSBitmapImageRep_ValueForProperty(void* ptr, void* property);
int C_NSBitmapImageRep_IncrementalLoadFromData_Complete(void* ptr, Array data, bool complete);
void C_NSBitmapImageRep_SetColor_AtX_Y(void* ptr, void* color, int x, int y);
void* C_NSBitmapImageRep_ColorAtX_Y(void* ptr, int x, int y);
void* C_NSBitmapImageRep_BitmapImageRepByRetaggingWithColorSpace(void* ptr, void* newSpace);
int C_NSBitmapImageRep_BitsPerPixel(void* ptr);
int C_NSBitmapImageRep_BytesPerPlane(void* ptr);
int C_NSBitmapImageRep_BytesPerRow(void* ptr);
bool C_NSBitmapImageRep_IsPlanar(void* ptr);
int C_NSBitmapImageRep_NumberOfPlanes(void* ptr);
int C_NSBitmapImageRep_SamplesPerPixel(void* ptr);
Array C_NSBitmapImageRep_TIFFRepresentation(void* ptr);
CGImageRef C_NSBitmapImageRep_CGImage(void* ptr);
void* C_NSBitmapImageRep_ColorSpace(void* ptr);
