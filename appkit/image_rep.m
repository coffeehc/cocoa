#import <Appkit/Appkit.h>
#import "image_rep.h"

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

Array C_NSImageRep_ImageRepsWithContentsOfFile(void* filename) {
    NSArray* result_ = [NSImageRep imageRepsWithContentsOfFile:(NSString*)filename];
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

Array C_NSImageRep_ImageRepsWithPasteboard(void* pasteboard) {
    NSArray* result_ = [NSImageRep imageRepsWithPasteboard:(NSPasteboard*)pasteboard];
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

Array C_NSImageRep_ImageRepsWithContentsOfURL(void* url) {
    NSArray* result_ = [NSImageRep imageRepsWithContentsOfURL:(NSURL*)url];
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

bool C_NSImageRep_ImageRep_CanInitWithData(Array data) {
    BOOL result_ = [NSImageRep canInitWithData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
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
    NSMutableDictionary* objcHints = [[NSMutableDictionary alloc] initWithCapacity: hints.len];
    if (hints.len > 0) {
    	void** hintsKeyData = (void**)hints.key_data;
    	void** hintsValueData = (void**)hints.value_data;
    	for (int i = 0; i < hints.len; i++) {
    		void* kp = hintsKeyData[i];
    		void* vp = hintsValueData[i];
    		[objcHints setObject:(NSImageHintKey)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    BOOL result_ = [nSImageRep drawInRect:dstSpacePortionRect fromRect:srcSpacePortionRect operation:op fraction:requestedAlpha respectFlipped:respectContextIsFlipped hints:objcHints];
    return result_;
}

Array C_NSImageRep_ImageRep_ImageTypes() {
    NSArray* result_ = [NSImageRep imageTypes];
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

Array C_NSImageRep_ImageRep_ImageUnfilteredTypes() {
    NSArray* result_ = [NSImageRep imageUnfilteredTypes];
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
