#import "glyph_generator.h"
#import <AppKit/NSGlyphGenerator.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_GlyphGenerator_Alloc() {
    return [NSGlyphGenerator alloc];
}

void* C_NSGlyphGenerator_AllocGlyphGenerator() {
    NSGlyphGenerator* result_ = [NSGlyphGenerator alloc];
    return result_;
}

void* C_NSGlyphGenerator_Init(void* ptr) {
    NSGlyphGenerator* nSGlyphGenerator = (NSGlyphGenerator*)ptr;
    NSGlyphGenerator* result_ = [nSGlyphGenerator init];
    return result_;
}

void* C_NSGlyphGenerator_NewGlyphGenerator() {
    NSGlyphGenerator* result_ = [NSGlyphGenerator new];
    return result_;
}

void* C_NSGlyphGenerator_Autorelease(void* ptr) {
    NSGlyphGenerator* nSGlyphGenerator = (NSGlyphGenerator*)ptr;
    NSGlyphGenerator* result_ = [nSGlyphGenerator autorelease];
    return result_;
}

void* C_NSGlyphGenerator_Retain(void* ptr) {
    NSGlyphGenerator* nSGlyphGenerator = (NSGlyphGenerator*)ptr;
    NSGlyphGenerator* result_ = [nSGlyphGenerator retain];
    return result_;
}

void* C_NSGlyphGenerator_SharedGlyphGenerator() {
    NSGlyphGenerator* result_ = [NSGlyphGenerator sharedGlyphGenerator];
    return result_;
}
