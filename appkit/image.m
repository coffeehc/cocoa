#import <Appkit/Appkit.h>
#import "image.h"

void* C_Image_Alloc() {
    return [NSImage alloc];
}

void* C_NSImage_InitByReferencingFile(void* ptr, void* fileName) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage initByReferencingFile:(NSString*)fileName];
    return result;
}

void* C_NSImage_InitByReferencingURL(void* ptr, void* url) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage initByReferencingURL:(NSURL*)url];
    return result;
}

void* C_NSImage_InitWithContentsOfFile(void* ptr, void* fileName) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage initWithContentsOfFile:(NSString*)fileName];
    return result;
}

void* C_NSImage_InitWithContentsOfURL(void* ptr, void* url) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage initWithContentsOfURL:(NSURL*)url];
    return result;
}

void* C_NSImage_InitWithData(void* ptr, Array data) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage initWithData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
    return result;
}

void* C_NSImage_InitWithDataIgnoringOrientation(void* ptr, Array data) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage initWithDataIgnoringOrientation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
    return result;
}

void* C_NSImage_InitWithCGImage_Size(void* ptr, CGImageRef cgImage, CGSize size) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage initWithCGImage:cgImage size:size];
    return result;
}

void* C_NSImage_InitWithPasteboard(void* ptr, void* pasteboard) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage initWithPasteboard:(NSPasteboard*)pasteboard];
    return result;
}

void* C_NSImage_InitWithCoder(void* ptr, void* coder) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSImage_InitWithSize(void* ptr, CGSize size) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage initWithSize:size];
    return result;
}

void* C_NSImage_Init(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage init];
    return result;
}

void* C_NSImage_ImageNamed(void* name) {
    NSImage* result = [NSImage imageNamed:(NSString*)name];
    return result;
}

bool C_NSImage_SetName(void* ptr, void* _string) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result = [nSImage setName:(NSString*)_string];
    return result;
}

void* C_NSImage_Name(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImageName result = [nSImage name];
    return result;
}

void* C_NSImage_ImageWithSystemSymbolName_AccessibilityDescription(void* symbolName, void* description) {
    NSImage* result = [NSImage imageWithSystemSymbolName:(NSString*)symbolName accessibilityDescription:(NSString*)description];
    return result;
}

void* C_NSImage_ImageWithSymbolConfiguration(void* ptr, void* configuration) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result = [nSImage imageWithSymbolConfiguration:(NSImageSymbolConfiguration*)configuration];
    return result;
}

bool C_NSImage_Image_CanInitWithPasteboard(void* pasteboard) {
    BOOL result = [NSImage canInitWithPasteboard:(NSPasteboard*)pasteboard];
    return result;
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
    BOOL result = [nSImage drawRepresentation:(NSImageRep*)imageRep inRect:rect];
    return result;
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
    NSData* result = [nSImage TIFFRepresentationUsingCompression:comp factor:factor];
    Array resultarray;
    resultarray.data = [result bytes];
    resultarray.len = result.length;
    return resultarray;
}

void C_NSImage_CancelIncrementalLoad(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage cancelIncrementalLoad];
}

void* C_NSImage_LayerContentsForContentsScale(void* ptr, double layerContentsScale) {
    NSImage* nSImage = (NSImage*)ptr;
    id result = [nSImage layerContentsForContentsScale:layerContentsScale];
    return result;
}

double C_NSImage_RecommendedLayerContentsScale(void* ptr, double preferredContentsScale) {
    NSImage* nSImage = (NSImage*)ptr;
    CGFloat result = [nSImage recommendedLayerContentsScale:preferredContentsScale];
    return result;
}

CGSize C_NSImage_Size(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSSize result = [nSImage size];
    return result;
}

void C_NSImage_SetSize(void* ptr, CGSize value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setSize:value];
}

bool C_NSImage_IsTemplate(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result = [nSImage isTemplate];
    return result;
}

void C_NSImage_SetTemplate(void* ptr, bool value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setTemplate:value];
}

Array C_NSImage_ImageTypes() {
    NSArray* result = [NSImage imageTypes];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

Array C_NSImage_ImageUnfilteredTypes() {
    NSArray* result = [NSImage imageUnfilteredTypes];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

Array C_NSImage_Representations(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSArray* result = [nSImage representations];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

bool C_NSImage_PrefersColorMatch(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result = [nSImage prefersColorMatch];
    return result;
}

void C_NSImage_SetPrefersColorMatch(void* ptr, bool value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setPrefersColorMatch:value];
}

bool C_NSImage_UsesEPSOnResolutionMismatch(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result = [nSImage usesEPSOnResolutionMismatch];
    return result;
}

void C_NSImage_SetUsesEPSOnResolutionMismatch(void* ptr, bool value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setUsesEPSOnResolutionMismatch:value];
}

bool C_NSImage_MatchesOnMultipleResolution(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result = [nSImage matchesOnMultipleResolution];
    return result;
}

void C_NSImage_SetMatchesOnMultipleResolution(void* ptr, bool value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setMatchesOnMultipleResolution:value];
}

bool C_NSImage_IsValid(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result = [nSImage isValid];
    return result;
}

void* C_NSImage_BackgroundColor(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSColor* result = [nSImage backgroundColor];
    return result;
}

void C_NSImage_SetBackgroundColor(void* ptr, void* value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setBackgroundColor:(NSColor*)value];
}

NSEdgeInsets C_NSImage_CapInsets(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSEdgeInsets result = [nSImage capInsets];
    return result;
}

void C_NSImage_SetCapInsets(void* ptr, NSEdgeInsets value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setCapInsets:value];
}

int C_NSImage_ResizingMode(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImageResizingMode result = [nSImage resizingMode];
    return result;
}

void C_NSImage_SetResizingMode(void* ptr, int value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setResizingMode:value];
}

CGRect C_NSImage_AlignmentRect(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSRect result = [nSImage alignmentRect];
    return result;
}

void C_NSImage_SetAlignmentRect(void* ptr, CGRect value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setAlignmentRect:value];
}

unsigned int C_NSImage_CacheMode(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImageCacheMode result = [nSImage cacheMode];
    return result;
}

void C_NSImage_SetCacheMode(void* ptr, unsigned int value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setCacheMode:value];
}

Array C_NSImage_TIFFRepresentation(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSData* result = [nSImage TIFFRepresentation];
    Array resultarray;
    resultarray.data = [result bytes];
    resultarray.len = result.length;
    return resultarray;
}

void* C_NSImage_AccessibilityDescription(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSString* result = [nSImage accessibilityDescription];
    return result;
}

void C_NSImage_SetAccessibilityDescription(void* ptr, void* value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setAccessibilityDescription:(NSString*)value];
}

bool C_NSImage_MatchesOnlyOnBestFittingAxis(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    BOOL result = [nSImage matchesOnlyOnBestFittingAxis];
    return result;
}

void C_NSImage_SetMatchesOnlyOnBestFittingAxis(void* ptr, bool value) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage setMatchesOnlyOnBestFittingAxis:value];
}
