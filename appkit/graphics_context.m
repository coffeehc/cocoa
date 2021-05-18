#import <Appkit/Appkit.h>
#import "graphics_context.h"

void* C_GraphicsContext_Alloc() {
    return [NSGraphicsContext alloc];
}

void* C_NSGraphicsContext_Init(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    NSGraphicsContext* result_ = [nSGraphicsContext init];
    return result_;
}

void* C_NSGraphicsContext_GraphicsContextWithBitmapImageRep(void* bitmapRep) {
    NSGraphicsContext* result_ = [NSGraphicsContext graphicsContextWithBitmapImageRep:(NSBitmapImageRep*)bitmapRep];
    return result_;
}

void* C_NSGraphicsContext_GraphicsContextWithCGContext_Flipped(CGContextRef graphicsPort, bool initialFlippedState) {
    NSGraphicsContext* result_ = [NSGraphicsContext graphicsContextWithCGContext:graphicsPort flipped:initialFlippedState];
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

CGContextRef C_NSGraphicsContext_CGContext(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    CGContextRef result_ = [nSGraphicsContext CGContext];
    return result_;
}

bool C_NSGraphicsContext_IsDrawingToScreen(void* ptr) {
    NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
    BOOL result_ = [nSGraphicsContext isDrawingToScreen];
    return result_;
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
