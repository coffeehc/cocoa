#import "text_checking_result.h"
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <Foundation/NSTextCheckingResult.h>

void* C_TextCheckingResult_Alloc() {
    return [NSTextCheckingResult alloc];
}

void* C_NSTextCheckingResult_AllocTextCheckingResult() {
    NSTextCheckingResult* result_ = [NSTextCheckingResult alloc];
    return result_;
}

void* C_NSTextCheckingResult_Init(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSTextCheckingResult* result_ = [nSTextCheckingResult init];
    return result_;
}

void* C_NSTextCheckingResult_NewTextCheckingResult() {
    NSTextCheckingResult* result_ = [NSTextCheckingResult new];
    return result_;
}

void* C_NSTextCheckingResult_Autorelease(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSTextCheckingResult* result_ = [nSTextCheckingResult autorelease];
    return result_;
}

void* C_NSTextCheckingResult_Retain(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSTextCheckingResult* result_ = [nSTextCheckingResult retain];
    return result_;
}

NSRange C_NSTextCheckingResult_RangeAtIndex(void* ptr, unsigned int idx) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSRange result_ = [nSTextCheckingResult rangeAtIndex:idx];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_ReplacementCheckingResultWithRange_ReplacementString(NSRange _range, void* replacementString) {
    NSTextCheckingResult* result_ = [NSTextCheckingResult replacementCheckingResultWithRange:_range replacementString:(NSString*)replacementString];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_LinkCheckingResultWithRange_URL(NSRange _range, void* url) {
    NSTextCheckingResult* result_ = [NSTextCheckingResult linkCheckingResultWithRange:_range URL:(NSURL*)url];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_AddressCheckingResultWithRange_Components(NSRange _range, Dictionary components) {
    NSMutableDictionary* objcComponents = [NSMutableDictionary dictionaryWithCapacity:components.len];
    if (components.len > 0) {
    	void** componentsKeyData = (void**)components.key_data;
    	void** componentsValueData = (void**)components.value_data;
    	for (int i = 0; i < components.len; i++) {
    		void* kp = componentsKeyData[i];
    		void* vp = componentsValueData[i];
    		[objcComponents setObject:(NSString*)vp forKey:(id<NSCopying>)(NSString*)kp];
    	}
    }
    NSTextCheckingResult* result_ = [NSTextCheckingResult addressCheckingResultWithRange:_range components:objcComponents];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_TransitInformationCheckingResultWithRange_Components(NSRange _range, Dictionary components) {
    NSMutableDictionary* objcComponents = [NSMutableDictionary dictionaryWithCapacity:components.len];
    if (components.len > 0) {
    	void** componentsKeyData = (void**)components.key_data;
    	void** componentsValueData = (void**)components.value_data;
    	for (int i = 0; i < components.len; i++) {
    		void* kp = componentsKeyData[i];
    		void* vp = componentsValueData[i];
    		[objcComponents setObject:(NSString*)vp forKey:(id<NSCopying>)(NSString*)kp];
    	}
    }
    NSTextCheckingResult* result_ = [NSTextCheckingResult transitInformationCheckingResultWithRange:_range components:objcComponents];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_PhoneNumberCheckingResultWithRange_PhoneNumber(NSRange _range, void* phoneNumber) {
    NSTextCheckingResult* result_ = [NSTextCheckingResult phoneNumberCheckingResultWithRange:_range phoneNumber:(NSString*)phoneNumber];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_DateCheckingResultWithRange_Date(NSRange _range, void* date) {
    NSTextCheckingResult* result_ = [NSTextCheckingResult dateCheckingResultWithRange:_range date:(NSDate*)date];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_DateCheckingResultWithRange_Date_TimeZone_Duration(NSRange _range, void* date, void* timeZone, double duration) {
    NSTextCheckingResult* result_ = [NSTextCheckingResult dateCheckingResultWithRange:_range date:(NSDate*)date timeZone:(NSTimeZone*)timeZone duration:duration];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_DashCheckingResultWithRange_ReplacementString(NSRange _range, void* replacementString) {
    NSTextCheckingResult* result_ = [NSTextCheckingResult dashCheckingResultWithRange:_range replacementString:(NSString*)replacementString];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_QuoteCheckingResultWithRange_ReplacementString(NSRange _range, void* replacementString) {
    NSTextCheckingResult* result_ = [NSTextCheckingResult quoteCheckingResultWithRange:_range replacementString:(NSString*)replacementString];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_SpellCheckingResultWithRange(NSRange _range) {
    NSTextCheckingResult* result_ = [NSTextCheckingResult spellCheckingResultWithRange:_range];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_CorrectionCheckingResultWithRange_ReplacementString(NSRange _range, void* replacementString) {
    NSTextCheckingResult* result_ = [NSTextCheckingResult correctionCheckingResultWithRange:_range replacementString:(NSString*)replacementString];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_OrthographyCheckingResultWithRange_Orthography(NSRange _range, void* orthography) {
    NSTextCheckingResult* result_ = [NSTextCheckingResult orthographyCheckingResultWithRange:_range orthography:(NSOrthography*)orthography];
    return result_;
}

void* C_NSTextCheckingResult_ResultByAdjustingRangesWithOffset(void* ptr, int offset) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSTextCheckingResult* result_ = [nSTextCheckingResult resultByAdjustingRangesWithOffset:offset];
    return result_;
}

NSRange C_NSTextCheckingResult_RangeWithName(void* ptr, void* name) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSRange result_ = [nSTextCheckingResult rangeWithName:(NSString*)name];
    return result_;
}

void* C_NSTextCheckingResult_TextCheckingResult_CorrectionCheckingResultWithRange_ReplacementString_AlternativeStrings(NSRange _range, void* replacementString, Array alternativeStrings) {
    NSMutableArray* objcAlternativeStrings = [NSMutableArray arrayWithCapacity:alternativeStrings.len];
    if (alternativeStrings.len > 0) {
    	void** alternativeStringsData = (void**)alternativeStrings.data;
    	for (int i = 0; i < alternativeStrings.len; i++) {
    		void* p = alternativeStringsData[i];
    		[objcAlternativeStrings addObject:(NSString*)p];
    	}
    }
    NSTextCheckingResult* result_ = [NSTextCheckingResult correctionCheckingResultWithRange:_range replacementString:(NSString*)replacementString alternativeStrings:objcAlternativeStrings];
    return result_;
}

NSRange C_NSTextCheckingResult_Range(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSRange result_ = [nSTextCheckingResult range];
    return result_;
}

unsigned int C_NSTextCheckingResult_NumberOfRanges(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSUInteger result_ = [nSTextCheckingResult numberOfRanges];
    return result_;
}

void* C_NSTextCheckingResult_ReplacementString(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSString* result_ = [nSTextCheckingResult replacementString];
    return result_;
}

void* C_NSTextCheckingResult_RegularExpression(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSRegularExpression* result_ = [nSTextCheckingResult regularExpression];
    return result_;
}

Dictionary C_NSTextCheckingResult_Components(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSDictionary* result_ = [nSTextCheckingResult components];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSTextCheckingKey kp = [result_Keys objectAtIndex:i];
    		NSString* vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void* C_NSTextCheckingResult_URL(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSURL* result_ = [nSTextCheckingResult URL];
    return result_;
}

Dictionary C_NSTextCheckingResult_AddressComponents(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSDictionary* result_ = [nSTextCheckingResult addressComponents];
    Dictionary result_Array;
    NSArray * result_Keys = [result_ allKeys];
    int result_Count = [result_Keys count];
    if (result_Count > 0) {
    	void** result_KeyData = malloc(result_Count * sizeof(void*));
    	void** result_ValueData = malloc(result_Count * sizeof(void*));
    	for (int i = 0; i < result_Count; i++) {
    		NSTextCheckingKey kp = [result_Keys objectAtIndex:i];
    		NSString* vp = result_[kp];
    		 result_KeyData[i] = kp;
    		 result_ValueData[i] = vp;
    	}
    	result_Array.key_data = result_KeyData;
    	result_Array.value_data = result_ValueData;
    	result_Array.len = result_Count;
    }
    return result_Array;
}

void* C_NSTextCheckingResult_PhoneNumber(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSString* result_ = [nSTextCheckingResult phoneNumber];
    return result_;
}

void* C_NSTextCheckingResult_Date(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSDate* result_ = [nSTextCheckingResult date];
    return result_;
}

double C_NSTextCheckingResult_Duration(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSTimeInterval result_ = [nSTextCheckingResult duration];
    return result_;
}

void* C_NSTextCheckingResult_TimeZone(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSTimeZone* result_ = [nSTextCheckingResult timeZone];
    return result_;
}

void* C_NSTextCheckingResult_Orthography(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSOrthography* result_ = [nSTextCheckingResult orthography];
    return result_;
}

Array C_NSTextCheckingResult_AlternativeStrings(void* ptr) {
    NSTextCheckingResult* nSTextCheckingResult = (NSTextCheckingResult*)ptr;
    NSArray* result_ = [nSTextCheckingResult alternativeStrings];
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
