#import <Appkit/Appkit.h>
#import "bitmap_image_rep.h"

void* C_BitmapImageRep_Alloc() {
	return [NSBitmapImageRep alloc];
}

void* C_NSBitmapImageRep_InitWithCGImage(void* ptr, CGImageRef cgImage) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	NSBitmapImageRep* result = [nSBitmapImageRep initWithCGImage:cgImage];
	return result;
}

void* C_NSBitmapImageRep_InitWithData(void* ptr, Array data) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	NSBitmapImageRep* result = [nSBitmapImageRep initWithData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
	return result;
}

void* C_NSBitmapImageRep_InitForIncrementalLoad(void* ptr) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	NSBitmapImageRep* result = [nSBitmapImageRep initForIncrementalLoad];
	return result;
}

void* C_NSBitmapImageRep_Init(void* ptr) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	NSBitmapImageRep* result = [nSBitmapImageRep init];
	return result;
}

void* C_NSBitmapImageRep_InitWithCoder(void* ptr, void* coder) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	NSBitmapImageRep* result = [nSBitmapImageRep initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSBitmapImageRep_BitmapImageRepImageRepWithData(Array data) {
	NSBitmapImageRep* result = [NSBitmapImageRep imageRepWithData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
	return result;
}

void C_NSBitmapImageRep_ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(void* ptr, double midPoint, void* midPointColor, void* shadowColor, void* lightColor) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	[nSBitmapImageRep colorizeByMappingGray:midPoint toColor:(NSColor*)midPointColor blackMapping:(NSColor*)shadowColor whiteMapping:(NSColor*)lightColor];
}

void C_NSBitmapImageRep_SetProperty_WithValue(void* ptr, void* property, void* value) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	[nSBitmapImageRep setProperty:(NSString*)property withValue:(id)value];
}

void* C_NSBitmapImageRep_ValueForProperty(void* ptr, void* property) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	id result = [nSBitmapImageRep valueForProperty:(NSString*)property];
	return result;
}

int C_NSBitmapImageRep_IncrementalLoadFromData_Complete(void* ptr, Array data, bool complete) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	int result = [nSBitmapImageRep incrementalLoadFromData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] complete:complete];
	return result;
}

void C_NSBitmapImageRep_SetColor_AtX_Y(void* ptr, void* color, int x, int y) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	[nSBitmapImageRep setColor:(NSColor*)color atX:x y:y];
}

void* C_NSBitmapImageRep_ColorAtX_Y(void* ptr, int x, int y) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	NSColor* result = [nSBitmapImageRep colorAtX:x y:y];
	return result;
}

void* C_NSBitmapImageRep_BitmapImageRepByRetaggingWithColorSpace(void* ptr, void* newSpace) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	NSBitmapImageRep* result = [nSBitmapImageRep bitmapImageRepByRetaggingWithColorSpace:(NSColorSpace*)newSpace];
	return result;
}

int C_NSBitmapImageRep_BitsPerPixel(void* ptr) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	int result = [nSBitmapImageRep bitsPerPixel];
	return result;
}

int C_NSBitmapImageRep_BytesPerPlane(void* ptr) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	int result = [nSBitmapImageRep bytesPerPlane];
	return result;
}

int C_NSBitmapImageRep_BytesPerRow(void* ptr) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	int result = [nSBitmapImageRep bytesPerRow];
	return result;
}

bool C_NSBitmapImageRep_IsPlanar(void* ptr) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	bool result = [nSBitmapImageRep isPlanar];
	return result;
}

int C_NSBitmapImageRep_NumberOfPlanes(void* ptr) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	int result = [nSBitmapImageRep numberOfPlanes];
	return result;
}

int C_NSBitmapImageRep_SamplesPerPixel(void* ptr) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	int result = [nSBitmapImageRep samplesPerPixel];
	return result;
}

Array C_NSBitmapImageRep_TIFFRepresentation(void* ptr) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	NSData* result = [nSBitmapImageRep TIFFRepresentation];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

CGImageRef C_NSBitmapImageRep_CGImage(void* ptr) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	CGImageRef result = [nSBitmapImageRep CGImage];
	return result;
}

void* C_NSBitmapImageRep_ColorSpace(void* ptr) {
	NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
	NSColorSpace* result = [nSBitmapImageRep colorSpace];
	return result;
}
