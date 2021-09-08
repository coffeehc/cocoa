#import <Appkit/Appkit.h>
#import "graphics_context.h"

void* C_GraphicsContext_Alloc() {
    return [NSGraphicsContext alloc];
}

void* C_NSGraphicsContext_AllocGraphicsContext() {
    NSGraphicsContext* result_ = [NSGraphicsContext alloc];
    return result_;
}

void* C_NSGraphicsContext_Init(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    NSGraphicsContext* result_ = [nSGraphicsContext init];
    return result_;
}

void* C_NSGraphicsContext_NewGraphicsContext() {
    NSGraphicsContext* result_ = [NSGraphicsContext new];
    return result_;
}

void* C_NSGraphicsContext_Autorelease(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    NSGraphicsContext* result_ = [nSGraphicsContext autorelease];
    return result_;
}

void* C_NSGraphicsContext_Retain(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    NSGraphicsContext* result_ = [nSGraphicsContext retain];
    return result_;
}

void* C_NSGraphicsContext_GraphicsContextWithAttributes(Dictionary attributes) {
    NSMutableDictionary* objcAttributes = [NSMutableDictionary dictionaryWithCapacity:attributes.len];
    if (attributes.len > 0) {
    	void** attributesKeyData = (void**)attributes.key_data;
    	void** attributesValueData = (void**)attributes.value_data;
    	for (int i = 0; i < attributes.len; i++) {
    		void* kp = attributesKeyData[i];
    		void* vp = attributesValueData[i];
    		[objcAttributes setObject:(NSGraphicsContextAttributeKey)(NSString*)kp forKey:(id)(NSString*)vp];
    	}
    }
    NSGraphicsContext* result_ = [NSGraphicsContext graphicsContextWithAttributes:objcAttributes];
    return result_;
}

void* C_NSGraphicsContext_GraphicsContextWithBitmapImageRep(void* bitmapRep) {
    NSGraphicsContext* result_ = [NSGraphicsContext graphicsContextWithBitmapImageRep:(NSBitmapImageRep*)bitmapRep];
    return result_;
}

void* C_NSGraphicsContext_GraphicsContextWithCGContext_Flipped(void* graphicsPort, bool initialFlippedState) {
    NSGraphicsContext* result_ = [NSGraphicsContext graphicsContextWithCGContext:(CGContextRef)graphicsPort flipped:initialFlippedState];
    return result_;
}

void* C_NSGraphicsContext_GraphicsContextWithWindow(void* window) {
    NSGraphicsContext* result_ = [NSGraphicsContext graphicsContextWithWindow:(NSWindow*)window];
    return result_;
}

void C_NSGraphicsContext_GraphicsContext_RestoreGraphicsState() {
    [NSGraphicsContext restoreGraphicsState];
}

void C_NSGraphicsContext_GraphicsContext_SaveGraphicsState() {
    [NSGraphicsContext saveGraphicsState];
}

bool C_NSGraphicsContext_GraphicsContext_CurrentContextDrawingToScreen() {
    BOOL result_ = [NSGraphicsContext currentContextDrawingToScreen];
    return result_;
}

void C_NSGraphicsContext_FlushGraphics(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    [nSGraphicsContext flushGraphics];
}

void* C_NSGraphicsContext_GraphicsContext_CurrentContext() {
    NSGraphicsContext* result_ = [NSGraphicsContext currentContext];
    return result_;
}

void C_NSGraphicsContext_GraphicsContext_SetCurrentContext(void* value) {
    [NSGraphicsContext setCurrentContext:(NSGraphicsContext*)value];
}

void* C_NSGraphicsContext_CGContext(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    CGContextRef result_ = [nSGraphicsContext CGContext];
    return result_;
}

bool C_NSGraphicsContext_IsDrawingToScreen(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    BOOL result_ = [nSGraphicsContext isDrawingToScreen];
    return result_;
}

Dictionary C_NSGraphicsContext_Attributes(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    NSDictionary* result_ = [nSGraphicsContext attributes];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSGraphicsContextAttributeKey kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

bool C_NSGraphicsContext_IsFlipped(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    BOOL result_ = [nSGraphicsContext isFlipped];
    return result_;
}

unsigned int C_NSGraphicsContext_CompositingOperation(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    NSCompositingOperation result_ = [nSGraphicsContext compositingOperation];
    return result_;
}

void C_NSGraphicsContext_SetCompositingOperation(void* ptr, unsigned int value) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    [nSGraphicsContext setCompositingOperation:value];
}

unsigned int C_NSGraphicsContext_ImageInterpolation(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    NSImageInterpolation result_ = [nSGraphicsContext imageInterpolation];
    return result_;
}

void C_NSGraphicsContext_SetImageInterpolation(void* ptr, unsigned int value) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    [nSGraphicsContext setImageInterpolation:value];
}

bool C_NSGraphicsContext_ShouldAntialias(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    BOOL result_ = [nSGraphicsContext shouldAntialias];
    return result_;
}

void C_NSGraphicsContext_SetShouldAntialias(void* ptr, bool value) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    [nSGraphicsContext setShouldAntialias:value];
}

CGPoint C_NSGraphicsContext_PatternPhase(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    NSPoint result_ = [nSGraphicsContext patternPhase];
    return result_;
}

void C_NSGraphicsContext_SetPatternPhase(void* ptr, CGPoint value) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    [nSGraphicsContext setPatternPhase:value];
}

int C_NSGraphicsContext_ColorRenderingIntent(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    NSColorRenderingIntent result_ = [nSGraphicsContext colorRenderingIntent];
    return result_;
}

void C_NSGraphicsContext_SetColorRenderingIntent(void* ptr, int value) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    [nSGraphicsContext setColorRenderingIntent:value];
}
