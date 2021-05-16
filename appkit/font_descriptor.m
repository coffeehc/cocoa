#import <Appkit/Appkit.h>
#import "font_descriptor.h"

void* C_FontDescriptor_Alloc() {
    return [NSFontDescriptor alloc];
}

void* C_NSFontDescriptor_FontDescriptorWithDesign(void* ptr, void* design) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result = [nSFontDescriptor fontDescriptorWithDesign:(NSString*)design];
    return result;
}

void* C_NSFontDescriptor_Init(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result = [nSFontDescriptor init];
    return result;
}

void* C_NSFontDescriptor_FontDescriptorWithName_Matrix(void* fontName, void* matrix) {
    NSFontDescriptor* result = [NSFontDescriptor fontDescriptorWithName:(NSString*)fontName matrix:(NSAffineTransform*)matrix];
    return result;
}

void* C_NSFontDescriptor_FontDescriptorWithName_Size(void* fontName, double size) {
    NSFontDescriptor* result = [NSFontDescriptor fontDescriptorWithName:(NSString*)fontName size:size];
    return result;
}

void* C_NSFontDescriptor_FontDescriptorWithFace(void* ptr, void* newFace) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result = [nSFontDescriptor fontDescriptorWithFace:(NSString*)newFace];
    return result;
}

void* C_NSFontDescriptor_FontDescriptorWithFamily(void* ptr, void* newFamily) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result = [nSFontDescriptor fontDescriptorWithFamily:(NSString*)newFamily];
    return result;
}

void* C_NSFontDescriptor_FontDescriptorWithMatrix(void* ptr, void* matrix) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result = [nSFontDescriptor fontDescriptorWithMatrix:(NSAffineTransform*)matrix];
    return result;
}

void* C_NSFontDescriptor_FontDescriptorWithSize(void* ptr, double newPointSize) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result = [nSFontDescriptor fontDescriptorWithSize:newPointSize];
    return result;
}

void* C_NSFontDescriptor_FontDescriptorWithSymbolicTraits(void* ptr, uint32_t symbolicTraits) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result = [nSFontDescriptor fontDescriptorWithSymbolicTraits:symbolicTraits];
    return result;
}

void* C_NSFontDescriptor_ObjectForKey(void* ptr, void* attribute) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    id result = [nSFontDescriptor objectForKey:(NSString*)attribute];
    return result;
}

void* C_NSFontDescriptor_Matrix(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSAffineTransform* result = [nSFontDescriptor matrix];
    return result;
}

double C_NSFontDescriptor_PointSize(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    CGFloat result = [nSFontDescriptor pointSize];
    return result;
}

void* C_NSFontDescriptor_PostscriptName(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSString* result = [nSFontDescriptor postscriptName];
    return result;
}

uint32_t C_NSFontDescriptor_SymbolicTraits(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptorSymbolicTraits result = [nSFontDescriptor symbolicTraits];
    return result;
}

bool C_NSFontDescriptor_RequiresFontAssetRequest(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    BOOL result = [nSFontDescriptor requiresFontAssetRequest];
    return result;
}
