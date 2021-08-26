#import <Appkit/Appkit.h>
#import "bitmap_image_rep.h"

void* C_BitmapImageRep_Alloc() {
    return [NSBitmapImageRep alloc];
}

void* C_NSBitmapImageRep_InitWithCGImage(void* ptr, void* cgImage) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSBitmapImageRep* result_ = [nSBitmapImageRep initWithCGImage:(CGImageRef)cgImage];
    return result_;
}

void* C_NSBitmapImageRep_InitWithData(void* ptr, Array data) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSBitmapImageRep* result_ = [nSBitmapImageRep initWithData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
    return result_;
}

void* C_NSBitmapImageRep_InitForIncrementalLoad(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSBitmapImageRep* result_ = [nSBitmapImageRep initForIncrementalLoad];
    return result_;
}

void* C_NSBitmapImageRep_Init(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSBitmapImageRep* result_ = [nSBitmapImageRep init];
    return result_;
}

void* C_NSBitmapImageRep_InitWithCoder(void* ptr, void* coder) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSBitmapImageRep* result_ = [nSBitmapImageRep initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSBitmapImageRep_BitmapImageRep_ImageRepWithData(Array data) {
    NSBitmapImageRep* result_ = [NSBitmapImageRep imageRepWithData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
    return result_;
}

Array C_NSBitmapImageRep_BitmapImageRep_ImageRepsWithData(Array data) {
    NSArray* result_ = [NSBitmapImageRep imageRepsWithData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
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

void C_NSBitmapImageRep_ColorizeByMappingGray_ToColor_BlackMapping_WhiteMapping(void* ptr, double midPoint, void* midPointColor, void* shadowColor, void* lightColor) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    [nSBitmapImageRep colorizeByMappingGray:midPoint toColor:(NSColor*)midPointColor blackMapping:(NSColor*)shadowColor whiteMapping:(NSColor*)lightColor];
}

Array C_NSBitmapImageRep_BitmapImageRep_TIFFRepresentationOfImageRepsInArray(Array array) {
    NSMutableArray* objcArray = [[NSMutableArray alloc] init];
    if (array.len > 0) {
    	void** arrayData = (void**)array.data;
    	for (int i = 0; i < array.len; i++) {
    		void* p = arrayData[i];
    		[objcArray addObject:(NSImageRep*)(NSImageRep*)p];
    	}
    }
    NSData* result_ = [NSBitmapImageRep TIFFRepresentationOfImageRepsInArray:objcArray];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

Array C_NSBitmapImageRep_BitmapImageRep_TIFFRepresentationOfImageRepsInArray_UsingCompression_Factor(Array array, unsigned int comp, float factor) {
    NSMutableArray* objcArray = [[NSMutableArray alloc] init];
    if (array.len > 0) {
    	void** arrayData = (void**)array.data;
    	for (int i = 0; i < array.len; i++) {
    		void* p = arrayData[i];
    		[objcArray addObject:(NSImageRep*)(NSImageRep*)p];
    	}
    }
    NSData* result_ = [NSBitmapImageRep TIFFRepresentationOfImageRepsInArray:objcArray usingCompression:comp factor:factor];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

Array C_NSBitmapImageRep_TIFFRepresentationUsingCompression_Factor(void* ptr, unsigned int comp, float factor) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSData* result_ = [nSBitmapImageRep TIFFRepresentationUsingCompression:comp factor:factor];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

Array C_NSBitmapImageRep_BitmapImageRep_RepresentationOfImageRepsInArray_UsingType_Properties(Array imageReps, unsigned int storageType, Dictionary properties) {
    NSMutableArray* objcImageReps = [[NSMutableArray alloc] init];
    if (imageReps.len > 0) {
    	void** imageRepsData = (void**)imageReps.data;
    	for (int i = 0; i < imageReps.len; i++) {
    		void* p = imageRepsData[i];
    		[objcImageReps addObject:(NSImageRep*)(NSImageRep*)p];
    	}
    }
    NSMutableDictionary* objcProperties = [[NSMutableDictionary alloc] initWithCapacity: properties.len];
    if (properties.len > 0) {
    	void** propertiesKeyData = (void**)properties.key_data;
    	void** propertiesValueData = (void**)properties.value_data;
    	for (int i = 0; i < properties.len; i++) {
    		void* kp = propertiesKeyData[i];
    		void* vp = propertiesValueData[i];
    		[objcProperties setObject:(NSBitmapImageRepPropertyKey)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    NSData* result_ = [NSBitmapImageRep representationOfImageRepsInArray:objcImageReps usingType:storageType properties:objcProperties];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

Array C_NSBitmapImageRep_RepresentationUsingType_Properties(void* ptr, unsigned int storageType, Dictionary properties) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSMutableDictionary* objcProperties = [[NSMutableDictionary alloc] initWithCapacity: properties.len];
    if (properties.len > 0) {
    	void** propertiesKeyData = (void**)properties.key_data;
    	void** propertiesValueData = (void**)properties.value_data;
    	for (int i = 0; i < properties.len; i++) {
    		void* kp = propertiesKeyData[i];
    		void* vp = propertiesValueData[i];
    		[objcProperties setObject:(NSBitmapImageRepPropertyKey)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    NSData* result_ = [nSBitmapImageRep representationUsingType:storageType properties:objcProperties];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

void* C_NSBitmapImageRep_BitmapImageRep_LocalizedNameForTIFFCompressionType(unsigned int compression) {
    NSString* result_ = [NSBitmapImageRep localizedNameForTIFFCompressionType:compression];
    return result_;
}

bool C_NSBitmapImageRep_CanBeCompressedUsing(void* ptr, unsigned int compression) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    BOOL result_ = [nSBitmapImageRep canBeCompressedUsing:compression];
    return result_;
}

void C_NSBitmapImageRep_SetCompression_Factor(void* ptr, unsigned int compression, float factor) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    [nSBitmapImageRep setCompression:compression factor:factor];
}

void C_NSBitmapImageRep_SetProperty_WithValue(void* ptr, void* property, void* value) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    [nSBitmapImageRep setProperty:(NSString*)property withValue:(id)value];
}

void* C_NSBitmapImageRep_ValueForProperty(void* ptr, void* property) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    id result_ = [nSBitmapImageRep valueForProperty:(NSString*)property];
    return result_;
}

int C_NSBitmapImageRep_IncrementalLoadFromData_Complete(void* ptr, Array data, bool complete) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSInteger result_ = [nSBitmapImageRep incrementalLoadFromData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len] complete:complete];
    return result_;
}

void C_NSBitmapImageRep_SetColor_AtX_Y(void* ptr, void* color, int x, int y) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    [nSBitmapImageRep setColor:(NSColor*)color atX:x y:y];
}

void* C_NSBitmapImageRep_ColorAtX_Y(void* ptr, int x, int y) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSColor* result_ = [nSBitmapImageRep colorAtX:x y:y];
    return result_;
}

void* C_NSBitmapImageRep_BitmapImageRepByConvertingToColorSpace_RenderingIntent(void* ptr, void* targetSpace, int renderingIntent) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSBitmapImageRep* result_ = [nSBitmapImageRep bitmapImageRepByConvertingToColorSpace:(NSColorSpace*)targetSpace renderingIntent:renderingIntent];
    return result_;
}

void* C_NSBitmapImageRep_BitmapImageRepByRetaggingWithColorSpace(void* ptr, void* newSpace) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSBitmapImageRep* result_ = [nSBitmapImageRep bitmapImageRepByRetaggingWithColorSpace:(NSColorSpace*)newSpace];
    return result_;
}

unsigned int C_NSBitmapImageRep_BitmapFormat(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSBitmapFormat result_ = [nSBitmapImageRep bitmapFormat];
    return result_;
}

int C_NSBitmapImageRep_BitsPerPixel(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSInteger result_ = [nSBitmapImageRep bitsPerPixel];
    return result_;
}

int C_NSBitmapImageRep_BytesPerPlane(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSInteger result_ = [nSBitmapImageRep bytesPerPlane];
    return result_;
}

int C_NSBitmapImageRep_BytesPerRow(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSInteger result_ = [nSBitmapImageRep bytesPerRow];
    return result_;
}

bool C_NSBitmapImageRep_IsPlanar(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    BOOL result_ = [nSBitmapImageRep isPlanar];
    return result_;
}

int C_NSBitmapImageRep_NumberOfPlanes(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSInteger result_ = [nSBitmapImageRep numberOfPlanes];
    return result_;
}

int C_NSBitmapImageRep_SamplesPerPixel(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSInteger result_ = [nSBitmapImageRep samplesPerPixel];
    return result_;
}

Array C_NSBitmapImageRep_TIFFRepresentation(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSData* result_ = [nSBitmapImageRep TIFFRepresentation];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

void* C_NSBitmapImageRep_CGImage(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    CGImageRef result_ = [nSBitmapImageRep CGImage];
    return result_;
}

void* C_NSBitmapImageRep_ColorSpace(void* ptr) {
    NSBitmapImageRep* nSBitmapImageRep = (NSBitmapImageRep*)ptr;
    NSColorSpace* result_ = [nSBitmapImageRep colorSpace];
    return result_;
}
