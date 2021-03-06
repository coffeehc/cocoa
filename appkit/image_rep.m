#import "image_rep.h"
#import <AppKit/NSImageRep.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_ImageRep_Alloc() {
    return [NSImageRep alloc];
}

void* C_NSImageRep_Init(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    NSImageRep* result_ = [nSImageRep init];
    return result_;
}

void* C_NSImageRep_InitWithCoder(void* ptr, void* coder) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    NSImageRep* result_ = [nSImageRep initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSImageRep_AllocImageRep() {
    NSImageRep* result_ = [NSImageRep alloc];
    return result_;
}

void* C_NSImageRep_NewImageRep() {
    NSImageRep* result_ = [NSImageRep new];
    return result_;
}

void* C_NSImageRep_Autorelease(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    NSImageRep* result_ = [nSImageRep autorelease];
    return result_;
}

void* C_NSImageRep_Retain(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    NSImageRep* result_ = [nSImageRep retain];
    return result_;
}

Array C_NSImageRep_ImageRepsWithContentsOfFile(void* filename) {
    NSArray* result_ = [NSImageRep imageRepsWithContentsOfFile:(NSString*)filename];
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

Array C_NSImageRep_ImageRepsWithPasteboard(void* pasteboard) {
    NSArray* result_ = [NSImageRep imageRepsWithPasteboard:(NSPasteboard*)pasteboard];
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

Array C_NSImageRep_ImageRepsWithContentsOfURL(void* url) {
    NSArray* result_ = [NSImageRep imageRepsWithContentsOfURL:(NSURL*)url];
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

void* C_NSImageRep_ImageRepWithContentsOfFile(void* filename) {
    NSImageRep* result_ = [NSImageRep imageRepWithContentsOfFile:(NSString*)filename];
    return result_;
}

void* C_NSImageRep_ImageRepWithPasteboard(void* pasteboard) {
    NSImageRep* result_ = [NSImageRep imageRepWithPasteboard:(NSPasteboard*)pasteboard];
    return result_;
}

void* C_NSImageRep_ImageRepWithContentsOfURL(void* url) {
    NSImageRep* result_ = [NSImageRep imageRepWithContentsOfURL:(NSURL*)url];
    return result_;
}

bool C_NSImageRep_ImageRep_CanInitWithData(void* data) {
    BOOL result_ = [NSImageRep canInitWithData:(NSData*)data];
    return result_;
}

bool C_NSImageRep_ImageRep_CanInitWithPasteboard(void* pasteboard) {
    BOOL result_ = [NSImageRep canInitWithPasteboard:(NSPasteboard*)pasteboard];
    return result_;
}

bool C_NSImageRep_Draw(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    BOOL result_ = [nSImageRep draw];
    return result_;
}

bool C_NSImageRep_DrawAtPoint(void* ptr, CGPoint point) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    BOOL result_ = [nSImageRep drawAtPoint:point];
    return result_;
}

bool C_NSImageRep_DrawInRect(void* ptr, CGRect rect) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    BOOL result_ = [nSImageRep drawInRect:rect];
    return result_;
}

bool C_NSImageRep_DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(void* ptr, CGRect dstSpacePortionRect, CGRect srcSpacePortionRect, unsigned int op, double requestedAlpha, bool respectContextIsFlipped, Dictionary hints) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    NSMutableDictionary* objcHints = [NSMutableDictionary dictionaryWithCapacity:hints.len];
    if (hints.len > 0) {
    	void** hintsKeyData = (void**)hints.key_data;
    	void** hintsValueData = (void**)hints.value_data;
    	for (int i = 0; i < hints.len; i++) {
    		void* kp = hintsKeyData[i];
    		void* vp = hintsValueData[i];
    		[objcHints setObject:(id)vp forKey:(id<NSCopying>)(NSString*)kp];
    	}
    }
    BOOL result_ = [nSImageRep drawInRect:dstSpacePortionRect fromRect:srcSpacePortionRect operation:op fraction:requestedAlpha respectFlipped:respectContextIsFlipped hints:objcHints];
    return result_;
}

Array C_NSImageRep_ImageRep_ImageTypes() {
    NSArray* result_ = [NSImageRep imageTypes];
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

Array C_NSImageRep_ImageRep_ImageUnfilteredTypes() {
    NSArray* result_ = [NSImageRep imageUnfilteredTypes];
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

CGSize C_NSImageRep_Size(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    NSSize result_ = [nSImageRep size];
    return result_;
}

void C_NSImageRep_SetSize(void* ptr, CGSize value) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    [nSImageRep setSize:value];
}

int C_NSImageRep_BitsPerSample(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    NSInteger result_ = [nSImageRep bitsPerSample];
    return result_;
}

void C_NSImageRep_SetBitsPerSample(void* ptr, int value) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    [nSImageRep setBitsPerSample:value];
}

void* C_NSImageRep_ColorSpaceName(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    NSColorSpaceName result_ = [nSImageRep colorSpaceName];
    return result_;
}

void C_NSImageRep_SetColorSpaceName(void* ptr, void* value) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    [nSImageRep setColorSpaceName:(NSString*)value];
}

bool C_NSImageRep_HasAlpha(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    BOOL result_ = [nSImageRep hasAlpha];
    return result_;
}

void C_NSImageRep_SetAlpha(void* ptr, bool value) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    [nSImageRep setAlpha:value];
}

bool C_NSImageRep_IsOpaque(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    BOOL result_ = [nSImageRep isOpaque];
    return result_;
}

void C_NSImageRep_SetOpaque(void* ptr, bool value) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    [nSImageRep setOpaque:value];
}

int C_NSImageRep_PixelsHigh(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    NSInteger result_ = [nSImageRep pixelsHigh];
    return result_;
}

void C_NSImageRep_SetPixelsHigh(void* ptr, int value) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    [nSImageRep setPixelsHigh:value];
}

int C_NSImageRep_PixelsWide(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    NSInteger result_ = [nSImageRep pixelsWide];
    return result_;
}

void C_NSImageRep_SetPixelsWide(void* ptr, int value) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    [nSImageRep setPixelsWide:value];
}

int C_NSImageRep_LayoutDirection(void* ptr) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    NSImageLayoutDirection result_ = [nSImageRep layoutDirection];
    return result_;
}

void C_NSImageRep_SetLayoutDirection(void* ptr, int value) {
    NSImageRep* nSImageRep = (NSImageRep*)ptr;
    [nSImageRep setLayoutDirection:value];
}
