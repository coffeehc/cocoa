#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_GlyphGenerator_Alloc();

void* C_NSGlyphGenerator_AllocGlyphGenerator();
void* C_NSGlyphGenerator_Init(void* ptr);
void* C_NSGlyphGenerator_NewGlyphGenerator();
void* C_NSGlyphGenerator_Autorelease(void* ptr);
void* C_NSGlyphGenerator_Retain(void* ptr);
void* C_NSGlyphGenerator_SharedGlyphGenerator();
