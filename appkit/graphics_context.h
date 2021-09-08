#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_GraphicsContext_Alloc();

void* C_NSGraphicsContext_AllocGraphicsContext();
void* C_NSGraphicsContext_Init(void* ptr);
void* C_NSGraphicsContext_NewGraphicsContext();
void* C_NSGraphicsContext_Autorelease(void* ptr);
void* C_NSGraphicsContext_Retain(void* ptr);
void* C_NSGraphicsContext_GraphicsContextWithAttributes(Dictionary attributes);
void* C_NSGraphicsContext_GraphicsContextWithBitmapImageRep(void* bitmapRep);
void* C_NSGraphicsContext_GraphicsContextWithCGContext_Flipped(void* graphicsPort, bool initialFlippedState);
void* C_NSGraphicsContext_GraphicsContextWithWindow(void* window);
void C_NSGraphicsContext_GraphicsContext_RestoreGraphicsState();
void C_NSGraphicsContext_GraphicsContext_SaveGraphicsState();
bool C_NSGraphicsContext_GraphicsContext_CurrentContextDrawingToScreen();
void C_NSGraphicsContext_FlushGraphics(void* ptr);
void* C_NSGraphicsContext_GraphicsContext_CurrentContext();
void C_NSGraphicsContext_GraphicsContext_SetCurrentContext(void* value);
void* C_NSGraphicsContext_CGContext(void* ptr);
bool C_NSGraphicsContext_IsDrawingToScreen(void* ptr);
Dictionary C_NSGraphicsContext_Attributes(void* ptr);
bool C_NSGraphicsContext_IsFlipped(void* ptr);
unsigned int C_NSGraphicsContext_CompositingOperation(void* ptr);
void C_NSGraphicsContext_SetCompositingOperation(void* ptr, unsigned int value);
unsigned int C_NSGraphicsContext_ImageInterpolation(void* ptr);
void C_NSGraphicsContext_SetImageInterpolation(void* ptr, unsigned int value);
bool C_NSGraphicsContext_ShouldAntialias(void* ptr);
void C_NSGraphicsContext_SetShouldAntialias(void* ptr, bool value);
CGPoint C_NSGraphicsContext_PatternPhase(void* ptr);
void C_NSGraphicsContext_SetPatternPhase(void* ptr, CGPoint value);
int C_NSGraphicsContext_ColorRenderingIntent(void* ptr);
void C_NSGraphicsContext_SetColorRenderingIntent(void* ptr, int value);
