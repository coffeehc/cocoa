#import <Appkit/Appkit.h>
#import "graphics_context.h"

void* C_GraphicsContext_Alloc() {
	return [NSGraphicsContext alloc];
}

void* C_NSGraphicsContext_Init(void* ptr) {
	NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
	NSGraphicsContext* result = [nSGraphicsContext init];
	return result;
}

void* C_NSGraphicsContext_GraphicsContextWithBitmapImageRep(void* bitmapRep) {
	NSGraphicsContext* result = [NSGraphicsContext graphicsContextWithBitmapImageRep:(NSBitmapImageRep*)bitmapRep];
	return result;
}

void* C_NSGraphicsContext_GraphicsContextWithCGContext_Flipped(CGContextRef graphicsPort, bool initialFlippedState) {
	NSGraphicsContext* result = [NSGraphicsContext graphicsContextWithCGContext:graphicsPort flipped:initialFlippedState];
	return result;
}

void* C_NSGraphicsContext_GraphicsContextWithWindow(void* window) {
	NSGraphicsContext* result = [NSGraphicsContext graphicsContextWithWindow:(NSWindow*)window];
	return result;
}

void C_NSGraphicsContext_GraphicsContextRestoreGraphicsState() {
	[NSGraphicsContext restoreGraphicsState];
}

void C_NSGraphicsContext_RestoreGraphicsState(void* ptr) {
	NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
	[nSGraphicsContext restoreGraphicsState];
}

void C_NSGraphicsContext_GraphicsContextSaveGraphicsState() {
	[NSGraphicsContext saveGraphicsState];
}

void C_NSGraphicsContext_SaveGraphicsState(void* ptr) {
	NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
	[nSGraphicsContext saveGraphicsState];
}

bool C_NSGraphicsContext_GraphicsContextCurrentContextDrawingToScreen() {
	bool result = [NSGraphicsContext currentContextDrawingToScreen];
	return result;
}

void C_NSGraphicsContext_FlushGraphics(void* ptr) {
	NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
	[nSGraphicsContext flushGraphics];
}

CGContextRef C_NSGraphicsContext_CGContext(void* ptr) {
	NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
	CGContextRef result = [nSGraphicsContext CGContext];
	return result;
}

bool C_NSGraphicsContext_IsDrawingToScreen(void* ptr) {
	NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
	bool result = [nSGraphicsContext isDrawingToScreen];
	return result;
}

bool C_NSGraphicsContext_IsFlipped(void* ptr) {
	NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
	bool result = [nSGraphicsContext isFlipped];
	return result;
}

bool C_NSGraphicsContext_ShouldAntialias(void* ptr) {
	NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
	bool result = [nSGraphicsContext shouldAntialias];
	return result;
}

void C_NSGraphicsContext_SetShouldAntialias(void* ptr, bool value) {
	NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
	[nSGraphicsContext setShouldAntialias:value];
}

CGPoint C_NSGraphicsContext_PatternPhase(void* ptr) {
	NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
	CGPoint result = [nSGraphicsContext patternPhase];
	return result;
}

void C_NSGraphicsContext_SetPatternPhase(void* ptr, CGPoint value) {
	NSGraphicsContext* nSGraphicsContext = (NSGraphicsContext*)ptr;
	[nSGraphicsContext setPatternPhase:value];
}
