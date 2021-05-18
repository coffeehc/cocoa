#import <Appkit/Appkit.h>
#import "font_descriptor.h"

void* C_FontDescriptor_Alloc() {
    return [NSFontDescriptor alloc];
}

void* C_NSFontDescriptor_FontDescriptorWithDesign(void* ptr, void* design) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor fontDescriptorWithDesign:(NSString*)design];
    return result_;
}

void* C_NSFontDescriptor_Init(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor init];
    return result_;
}

void* C_NSFontDescriptor_FontDescriptorWithName_Matrix(void* fontName, void* matrix) {
    NSFontDescriptor* result_ = [NSFontDescriptor fontDescriptorWithName:(NSString*)fontName matrix:(NSAffineTransform*)matrix];
    return result_;
}

void* C_NSFontDescriptor_FontDescriptorWithName_Size(void* fontName, double size) {
    NSFontDescriptor* result_ = [NSFontDescriptor fontDescriptorWithName:(NSString*)fontName size:size];
    return result_;
}

void* C_NSFontDescriptor_FontDescriptorWithFace(void* ptr, void* newFace) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor fontDescriptorWithFace:(NSString*)newFace];
    return result_;
}

void* C_NSFontDescriptor_FontDescriptorWithFamily(void* ptr, void* newFamily) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor fontDescriptorWithFamily:(NSString*)newFamily];
    return result_;
}

void* C_NSFontDescriptor_FontDescriptorWithMatrix(void* ptr, void* matrix) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor fontDescriptorWithMatrix:(NSAffineTransform*)matrix];
    return result_;
}

void* C_NSFontDescriptor_FontDescriptorWithSize(void* ptr, double newPointSize) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor fontDescriptorWithSize:newPointSize];
    return result_;
}

void* C_NSFontDescriptor_FontDescriptorWithSymbolicTraits(void* ptr, uint32_t symbolicTraits) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor fontDescriptorWithSymbolicTraits:symbolicTraits];
    return result_;
}

Array C_NSFontDescriptor_MatchingFontDescriptorsWithMandatoryKeys(void* ptr, void* mandatoryKeys) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSArray* result_ = [nSFontDescriptor matchingFontDescriptorsWithMandatoryKeys:(NSSet*)mandatoryKeys];
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

void* C_NSFontDescriptor_MatchingFontDescriptorWithMandatoryKeys(void* ptr, void* mandatoryKeys) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor matchingFontDescriptorWithMandatoryKeys:(NSSet*)mandatoryKeys];
    return result_;
}

void* C_NSFontDescriptor_ObjectForKey(void* ptr, void* attribute) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    id result_ = [nSFontDescriptor objectForKey:(NSString*)attribute];
    return result_;
}

void* C_NSFontDescriptor_Matrix(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSAffineTransform* result_ = [nSFontDescriptor matrix];
    return result_;
}

double C_NSFontDescriptor_PointSize(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    CGFloat result_ = [nSFontDescriptor pointSize];
    return result_;
}

void* C_NSFontDescriptor_PostscriptName(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSString* result_ = [nSFontDescriptor postscriptName];
    return result_;
}

uint32_t C_NSFontDescriptor_SymbolicTraits(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptorSymbolicTraits result_ = [nSFontDescriptor symbolicTraits];
    return result_;
}

bool C_NSFontDescriptor_RequiresFontAssetRequest(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    BOOL result_ = [nSFontDescriptor requiresFontAssetRequest];
    return result_;
}
