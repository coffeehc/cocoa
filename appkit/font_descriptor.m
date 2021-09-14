#import <Appkit/Appkit.h>
#import "font_descriptor.h"

void* C_FontDescriptor_Alloc() {
    return [NSFontDescriptor alloc];
}

void* C_NSFontDescriptor_InitWithFontAttributes(void* ptr, Dictionary attributes) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSMutableDictionary* objcAttributes = [NSMutableDictionary dictionaryWithCapacity:attributes.len];
    if (attributes.len > 0) {
    	void** attributesKeyData = (void**)attributes.key_data;
    	void** attributesValueData = (void**)attributes.value_data;
    	for (int i = 0; i < attributes.len; i++) {
    		void* kp = attributesKeyData[i];
    		void* vp = attributesValueData[i];
    		[objcAttributes setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    NSFontDescriptor* result_ = [nSFontDescriptor initWithFontAttributes:objcAttributes];
    return result_;
}

void* C_NSFontDescriptor_FontDescriptorWithDesign(void* ptr, void* design) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor fontDescriptorWithDesign:(NSString*)design];
    return result_;
}

void* C_NSFontDescriptor_AllocFontDescriptor() {
    NSFontDescriptor* result_ = [NSFontDescriptor alloc];
    return result_;
}

void* C_NSFontDescriptor_Init(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor init];
    return result_;
}

void* C_NSFontDescriptor_NewFontDescriptor() {
    NSFontDescriptor* result_ = [NSFontDescriptor new];
    return result_;
}

void* C_NSFontDescriptor_Autorelease(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor autorelease];
    return result_;
}

void* C_NSFontDescriptor_Retain(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSFontDescriptor* result_ = [nSFontDescriptor retain];
    return result_;
}

void* C_NSFontDescriptor_PreferredFontDescriptorForTextStyle_Options(void* style, Dictionary options) {
    NSMutableDictionary* objcOptions = [NSMutableDictionary dictionaryWithCapacity:options.len];
    if (options.len > 0) {
    	void** optionsKeyData = (void**)options.key_data;
    	void** optionsValueData = (void**)options.value_data;
    	for (int i = 0; i < options.len; i++) {
    		void* kp = optionsKeyData[i];
    		void* vp = optionsValueData[i];
    		[objcOptions setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    NSFontDescriptor* result_ = [NSFontDescriptor preferredFontDescriptorForTextStyle:(NSString*)style options:objcOptions];
    return result_;
}

void* C_NSFontDescriptor_FontDescriptorWithFontAttributes(Dictionary attributes) {
    NSMutableDictionary* objcAttributes = [NSMutableDictionary dictionaryWithCapacity:attributes.len];
    if (attributes.len > 0) {
    	void** attributesKeyData = (void**)attributes.key_data;
    	void** attributesValueData = (void**)attributes.value_data;
    	for (int i = 0; i < attributes.len; i++) {
    		void* kp = attributesKeyData[i];
    		void* vp = attributesValueData[i];
    		[objcAttributes setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    NSFontDescriptor* result_ = [NSFontDescriptor fontDescriptorWithFontAttributes:objcAttributes];
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

void* C_NSFontDescriptor_FontDescriptorByAddingAttributes(void* ptr, Dictionary attributes) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSMutableDictionary* objcAttributes = [NSMutableDictionary dictionaryWithCapacity:attributes.len];
    if (attributes.len > 0) {
    	void** attributesKeyData = (void**)attributes.key_data;
    	void** attributesValueData = (void**)attributes.value_data;
    	for (int i = 0; i < attributes.len; i++) {
    		void* kp = attributesKeyData[i];
    		void* vp = attributesValueData[i];
    		[objcAttributes setObject:(NSString*)kp forKey:(id)vp];
    	}
    }
    NSFontDescriptor* result_ = [nSFontDescriptor fontDescriptorByAddingAttributes:objcAttributes];
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

Dictionary C_NSFontDescriptor_FontAttributes(void* ptr) {
    NSFontDescriptor* nSFontDescriptor = (NSFontDescriptor*)ptr;
    NSDictionary* result_ = [nSFontDescriptor fontAttributes];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSFontDescriptorAttributeName kp = [result_Keys objectAtIndex:i];
    		id vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
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
