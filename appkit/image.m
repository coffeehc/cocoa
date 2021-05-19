#import <Appkit/Appkit.h>
#import "image.h"

void* C_Image_Alloc() {
    return [NSImage alloc];
}

void* C_NSImage_InitByReferencingFile(void* ptr, void* fileName) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initByReferencingFile:(NSString*)fileName];
    return result_;
}

void* C_NSImage_InitByReferencingURL(void* ptr, void* url) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initByReferencingURL:(NSURL*)url];
    return result_;
}

void* C_NSImage_InitWithContentsOfFile(void* ptr, void* fileName) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initWithContentsOfFile:(NSString*)fileName];
    return result_;
}

void* C_NSImage_InitWithContentsOfURL(void* ptr, void* url) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initWithContentsOfURL:(NSURL*)url];
    return result_;
}

void* C_NSImage_InitWithData(void* ptr, Array data) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initWithData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
    return result_;
}

void* C_NSImage_InitWithDataIgnoringOrientation(void* ptr, Array data) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initWithDataIgnoringOrientation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
    return result_;
}

void* C_NSImage_InitWithPasteboard(void* ptr, void* pasteboard) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initWithPasteboard:(NSPasteboard*)pasteboard];
    return result_;
}

void* C_NSImage_InitWithCoder(void* ptr, void* coder) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSImage_InitWithSize(void* ptr, CGSize size) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initWithSize:size];
    return result_;
}

void* C_NSImage_Init(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage init];
    return result_;
}

void* C_NSImage_ImageNamed(void* name) {
    NSImage* result_ = [NSImage imageNamed:(NSString*)name];
    return result_;
}

bool C_NSImage_SetName(void* ptr, void* _string) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result_ = [nSImage setName:(NSString*)_string];
    return result_;
}

void* C_NSImage_Name(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImageName result_ = [nSImage name];
    return result_;
}

void* C_NSImage_ImageWithSystemSymbolName_AccessibilityDescription(void* symbolName, void* description) {
    NSImage* result_ = [NSImage imageWithSystemSymbolName:(NSString*)symbolName accessibilityDescription:(NSString*)description];
    return result_;
}

void* C_NSImage_ImageWithSymbolConfiguration(void* ptr, void* configuration) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage imageWithSymbolConfiguration:(NSImageSymbolConfiguration*)configuration];
    return result_;
}

bool C_NSImage_Image_CanInitWithPasteboard(void* pasteboard) {
    BOOL result_ = [NSImage canInitWithPasteboard:(NSPasteboard*)pasteboard];
    return result_;
}

void C_NSImage_AddRepresentation(void* ptr, void* imageRep) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage addRepresentation:(NSImageRep*)imageRep];
}

void C_NSImage_AddRepresentations(void* ptr, Array imageReps) {
    NSImage* nSImage = (NSImage*)ptr;
    NSMutableArray* objcImageReps = [[NSMutableArray alloc] init];
    void** imageRepsData = (void**)imageReps.data;
    for (int i = 0; i < imageReps.len; i++) {
    	void* p = imageRepsData[i];
    	[objcImageReps addObject:(NSImageRep*)(NSImageRep*)p];
    }
    [nSImage addRepresentations:objcImageReps];
}

void C_NSImage_RemoveRepresentation(void* ptr, void* imageRep) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage removeRepresentation:(NSImageRep*)imageRep];
}

void C_NSImage_DrawInRect(void* ptr, CGRect rect) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage drawInRect:rect];
}

void C_NSImage_DrawAtPoint_FromRect_Operation_Fraction(void* ptr, CGPoint point, CGRect fromRect, unsigned int op, double delta) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage drawAtPoint:point fromRect:fromRect operation:op fraction:delta];
}

void C_NSImage_DrawInRect_FromRect_Operation_Fraction(void* ptr, CGRect rect, CGRect fromRect, unsigned int op, double delta) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage drawInRect:rect fromRect:fromRect operation:op fraction:delta];
}

bool C_NSImage_DrawRepresentation_InRect(void* ptr, void* imageRep, CGRect rect) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result_ = [nSImage drawRepresentation:(NSImageRep*)imageRep inRect:rect];
    return result_;
}

void C_NSImage_LockFocus(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage lockFocus];
}

void C_NSImage_LockFocusFlipped(void* ptr, bool flipped) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage lockFocusFlipped:flipped];
}

void C_NSImage_UnlockFocus(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage unlockFocus];
}

void C_NSImage_Recache(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage recache];
}

Array C_NSImage_TIFFRepresentationUsingCompression_Factor(void* ptr, unsigned int comp, float factor) {
    NSImage* nSImage = (NSImage*)ptr;
    NSData* result_ = [nSImage TIFFRepresentationUsingCompression:comp factor:factor];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

void C_NSImage_CancelIncrementalLoad(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage cancelIncrementalLoad];
}

void* C_NSImage_LayerContentsForContentsScale(void* ptr, double layerContentsScale) {
    NSImage* nSImage = (NSImage*)ptr;
    id result_ = [nSImage layerContentsForContentsScale:layerContentsScale];
    return result_;
}

double C_NSImage_RecommendedLayerContentsScale(void* ptr, double preferredContentsScale) {
    NSImage* nSImage = (NSImage*)ptr;
    CGFloat result_ = [nSImage recommendedLayerContentsScale:preferredContentsScale];
    return result_;
}

void* C_NSImage_Delegate(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    id result_ = [nSImage delegate];
    return result_;
}

void C_NSImage_SetDelegate(void* ptr, void* value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setDelegate:(id)value];
}

CGSize C_NSImage_Size(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSSize result_ = [nSImage size];
    return result_;
}

void C_NSImage_SetSize(void* ptr, CGSize value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setSize:value];
}

bool C_NSImage_IsTemplate(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result_ = [nSImage isTemplate];
    return result_;
}

void C_NSImage_SetTemplate(void* ptr, bool value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setTemplate:value];
}

Array C_NSImage_ImageTypes() {
    NSArray* result_ = [NSImage imageTypes];
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

Array C_NSImage_ImageUnfilteredTypes() {
    NSArray* result_ = [NSImage imageUnfilteredTypes];
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

Array C_NSImage_Representations(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSArray* result_ = [nSImage representations];
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

bool C_NSImage_PrefersColorMatch(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result_ = [nSImage prefersColorMatch];
    return result_;
}

void C_NSImage_SetPrefersColorMatch(void* ptr, bool value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setPrefersColorMatch:value];
}

bool C_NSImage_UsesEPSOnResolutionMismatch(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result_ = [nSImage usesEPSOnResolutionMismatch];
    return result_;
}

void C_NSImage_SetUsesEPSOnResolutionMismatch(void* ptr, bool value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setUsesEPSOnResolutionMismatch:value];
}

bool C_NSImage_MatchesOnMultipleResolution(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result_ = [nSImage matchesOnMultipleResolution];
    return result_;
}

void C_NSImage_SetMatchesOnMultipleResolution(void* ptr, bool value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setMatchesOnMultipleResolution:value];
}

bool C_NSImage_IsValid(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result_ = [nSImage isValid];
    return result_;
}

void* C_NSImage_BackgroundColor(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSColor* result_ = [nSImage backgroundColor];
    return result_;
}

void C_NSImage_SetBackgroundColor(void* ptr, void* value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setBackgroundColor:(NSColor*)value];
}

NSEdgeInsets C_NSImage_CapInsets(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSEdgeInsets result_ = [nSImage capInsets];
    return result_;
}

void C_NSImage_SetCapInsets(void* ptr, NSEdgeInsets value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setCapInsets:value];
}

int C_NSImage_ResizingMode(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImageResizingMode result_ = [nSImage resizingMode];
    return result_;
}

void C_NSImage_SetResizingMode(void* ptr, int value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setResizingMode:value];
}

CGRect C_NSImage_AlignmentRect(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSRect result_ = [nSImage alignmentRect];
    return result_;
}

void C_NSImage_SetAlignmentRect(void* ptr, CGRect value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setAlignmentRect:value];
}

unsigned int C_NSImage_CacheMode(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImageCacheMode result_ = [nSImage cacheMode];
    return result_;
}

void C_NSImage_SetCacheMode(void* ptr, unsigned int value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setCacheMode:value];
}

Array C_NSImage_TIFFRepresentation(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSData* result_ = [nSImage TIFFRepresentation];
    Array result_array;
    result_array.data = [result_ bytes];
    result_array.len = result_.length;
    return result_array;
}

void* C_NSImage_AccessibilityDescription(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSString* result_ = [nSImage accessibilityDescription];
    return result_;
}

void C_NSImage_SetAccessibilityDescription(void* ptr, void* value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setAccessibilityDescription:(NSString*)value];
}

bool C_NSImage_MatchesOnlyOnBestFittingAxis(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result_ = [nSImage matchesOnlyOnBestFittingAxis];
    return result_;
}

void C_NSImage_SetMatchesOnlyOnBestFittingAxis(void* ptr, bool value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setMatchesOnlyOnBestFittingAxis:value];
}
