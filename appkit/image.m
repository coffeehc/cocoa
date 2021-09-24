#import "image.h"
#import <AppKit/NSImage.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_Image_Alloc() {
    return [NSImage alloc];
}

void* C_NSImage_ImageWithSystemSymbolName_AccessibilityDescription(void* symbolName, void* description) {
    NSImage* result_ = [NSImage imageWithSystemSymbolName:(NSString*)symbolName accessibilityDescription:(NSString*)description];
    return result_;
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

void* C_NSImage_InitWithData(void* ptr, void* data) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initWithData:(NSData*)data];
    return result_;
}

void* C_NSImage_InitWithDataIgnoringOrientation(void* ptr, void* data) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initWithDataIgnoringOrientation:(NSData*)data];
    return result_;
}

void* C_NSImage_InitWithCGImage_Size(void* ptr, void* cgImage, CGSize size) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage initWithCGImage:(CGImageRef)cgImage size:size];
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

void* C_NSImage_AllocImage() {
    NSImage* result_ = [NSImage alloc];
    return result_;
}

void* C_NSImage_Init(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage init];
    return result_;
}

void* C_NSImage_NewImage() {
    NSImage* result_ = [NSImage new];
    return result_;
}

void* C_NSImage_Autorelease(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage autorelease];
    return result_;
}

void* C_NSImage_Retain(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSImage* result_ = [nSImage retain];
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
    NSMutableArray* objcImageReps = [NSMutableArray arrayWithCapacity:imageReps.len];
    if (imageReps.len > 0) {
    	void** imageRepsData = (void**)imageReps.data;
    	for (int i = 0; i < imageReps.len; i++) {
    		void* p = imageRepsData[i];
    		[objcImageReps addObject:(NSImageRep*)p];
    	}
    }
    [nSImage addRepresentations:objcImageReps];
}

void C_NSImage_RemoveRepresentation(void* ptr, void* imageRep) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage removeRepresentation:(NSImageRep*)imageRep];
}

void* C_NSImage_BestRepresentationForRect_Context_Hints(void* ptr, CGRect rect, void* referenceContext, Dictionary hints) {
    NSImage* nSImage = (NSImage*)ptr;
    NSMutableDictionary* objcHints = [NSMutableDictionary dictionaryWithCapacity:hints.len];
    if (hints.len > 0) {
    	void** hintsKeyData = (void**)hints.key_data;
    	void** hintsValueData = (void**)hints.value_data;
    	for (int i = 0; i < hints.len; i++) {
    		void* kp = hintsKeyData[i];
    		void* vp = hintsValueData[i];
    		[objcHints setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    NSImageRep* result_ = [nSImage bestRepresentationForRect:rect context:(NSGraphicsContext*)referenceContext hints:objcHints];
    return result_;
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

void C_NSImage_DrawInRect_FromRect_Operation_Fraction_RespectFlipped_Hints(void* ptr, CGRect dstSpacePortionRect, CGRect srcSpacePortionRect, unsigned int op, double requestedAlpha, bool respectContextIsFlipped, Dictionary hints) {
    NSImage* nSImage = (NSImage*)ptr;
    NSMutableDictionary* objcHints = [NSMutableDictionary dictionaryWithCapacity:hints.len];
    if (hints.len > 0) {
    	void** hintsKeyData = (void**)hints.key_data;
    	void** hintsValueData = (void**)hints.value_data;
    	for (int i = 0; i < hints.len; i++) {
    		void* kp = hintsKeyData[i];
    		void* vp = hintsValueData[i];
    		[objcHints setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    [nSImage drawInRect:dstSpacePortionRect fromRect:srcSpacePortionRect operation:op fraction:requestedAlpha respectFlipped:respectContextIsFlipped hints:objcHints];
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

void* C_NSImage_TIFFRepresentationUsingCompression_Factor(void* ptr, unsigned int comp, float factor) {
    NSImage* nSImage = (NSImage*)ptr;
    NSData* result_ = [nSImage TIFFRepresentationUsingCompression:comp factor:factor];
    return result_;
}

void C_NSImage_CancelIncrementalLoad(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    [nSImage cancelIncrementalLoad];
}

bool C_NSImage_HitTestRect_WithImageDestinationRect_Context_Hints_Flipped(void* ptr, CGRect testRectDestSpace, CGRect imageRectDestSpace, void* context, Dictionary hints, bool flipped) {
    NSImage* nSImage = (NSImage*)ptr;
    NSMutableDictionary* objcHints = [NSMutableDictionary dictionaryWithCapacity:hints.len];
    if (hints.len > 0) {
    	void** hintsKeyData = (void**)hints.key_data;
    	void** hintsValueData = (void**)hints.value_data;
    	for (int i = 0; i < hints.len; i++) {
    		void* kp = hintsKeyData[i];
    		void* vp = hintsValueData[i];
    		[objcHints setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    BOOL result_ = [nSImage hitTestRect:testRectDestSpace withImageDestinationRect:imageRectDestSpace context:(NSGraphicsContext*)context hints:objcHints flipped:flipped];
    return result_;
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
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_NSImage_ImageUnfilteredTypes() {
    NSArray* result_ = [NSImage imageUnfilteredTypes];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

Array C_NSImage_Representations(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSArray* result_ = [nSImage representations];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
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

void* C_NSImage_TIFFRepresentation(void* ptr) {
    NSImage* nSImage = (NSImage*)ptr;
    NSData* result_ = [nSImage TIFFRepresentation];
    return result_;
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

void* C_ImageSymbolConfiguration_Alloc() {
    return [NSImageSymbolConfiguration alloc];
}

void* C_NSImageSymbolConfiguration_AllocImageSymbolConfiguration() {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration alloc];
    return result_;
}

void* C_NSImageSymbolConfiguration_Init(void* ptr) {
    NSImageSymbolConfiguration* nSImageSymbolConfiguration = (NSImageSymbolConfiguration*)ptr;
    NSImageSymbolConfiguration* result_ = [nSImageSymbolConfiguration init];
    return result_;
}

void* C_NSImageSymbolConfiguration_NewImageSymbolConfiguration() {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration new];
    return result_;
}

void* C_NSImageSymbolConfiguration_Autorelease(void* ptr) {
    NSImageSymbolConfiguration* nSImageSymbolConfiguration = (NSImageSymbolConfiguration*)ptr;
    NSImageSymbolConfiguration* result_ = [nSImageSymbolConfiguration autorelease];
    return result_;
}

void* C_NSImageSymbolConfiguration_Retain(void* ptr) {
    NSImageSymbolConfiguration* nSImageSymbolConfiguration = (NSImageSymbolConfiguration*)ptr;
    NSImageSymbolConfiguration* result_ = [nSImageSymbolConfiguration retain];
    return result_;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithPointSize_Weight(double pointSize, double weight) {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration configurationWithPointSize:pointSize weight:weight];
    return result_;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithPointSize_Weight_Scale(double pointSize, double weight, int scale) {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration configurationWithPointSize:pointSize weight:weight scale:scale];
    return result_;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithTextStyle(void* style) {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration configurationWithTextStyle:(NSString*)style];
    return result_;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithTextStyle_Scale(void* style, int scale) {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration configurationWithTextStyle:(NSString*)style scale:scale];
    return result_;
}

void* C_NSImageSymbolConfiguration_ImageSymbolConfiguration_ConfigurationWithScale(int scale) {
    NSImageSymbolConfiguration* result_ = [NSImageSymbolConfiguration configurationWithScale:scale];
    return result_;
}
