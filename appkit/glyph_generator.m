#import <Appkit/Appkit.h>
#import "glyph_generator.h"

void* C_GlyphGenerator_Alloc() {
    return [NSGlyphGenerator alloc];
}

void* C_NSGlyphGenerator_Init(void* ptr) {
    NSGlyphGenerator* nSGlyphGenerator = (NSGlyphGenerator*)ptr;
    NSGlyphGenerator* result_ = [nSGlyphGenerator init];
    return result_;
}

void* C_NSGlyphGenerator_SharedGlyphGenerator() {
    NSGlyphGenerator* result_ = [NSGlyphGenerator sharedGlyphGenerator];
    return result_;
}
