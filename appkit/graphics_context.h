#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_GraphicsContext_Alloc();

void* C_NSGraphicsContext_Init(void* ptr);
void* C_NSGraphicsContext_GraphicsContextWithBitmapImageRep(void* bitmapRep);
void* C_NSGraphicsContext_GraphicsContextWithCGContext_Flipped(CGContextRef graphicsPort, bool initialFlippedState);
void* C_NSGraphicsContext_GraphicsContextWithWindow(void* window);
void C_NSGraphicsContext_GraphicsContextRestoreGraphicsState();
void C_NSGraphicsContext_RestoreGraphicsState(void* ptr);
void C_NSGraphicsContext_GraphicsContextSaveGraphicsState();
void C_NSGraphicsContext_SaveGraphicsState(void* ptr);
bool C_NSGraphicsContext_GraphicsContextCurrentContextDrawingToScreen();
void C_NSGraphicsContext_FlushGraphics(void* ptr);
CGContextRef C_NSGraphicsContext_CGContext(void* ptr);
bool C_NSGraphicsContext_IsDrawingToScreen(void* ptr);
bool C_NSGraphicsContext_IsFlipped(void* ptr);
bool C_NSGraphicsContext_ShouldAntialias(void* ptr);
void C_NSGraphicsContext_SetShouldAntialias(void* ptr, bool value);
CGPoint C_NSGraphicsContext_PatternPhase(void* ptr);
void C_NSGraphicsContext_SetPatternPhase(void* ptr, CGPoint value);
