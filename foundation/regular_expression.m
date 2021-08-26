#import <Foundation/Foundation.h>
#import "regular_expression.h"

void* C_RegularExpression_Alloc() {
    return [NSRegularExpression alloc];
}

void* C_NSRegularExpression_Init(void* ptr) {
    NSRegularExpression* nSRegularExpression = (NSRegularExpression*)ptr;
    NSRegularExpression* result_ = [nSRegularExpression init];
    return result_;
}

unsigned int C_NSRegularExpression_NumberOfMatchesInString_Options_Range(void* ptr, void* _string, unsigned int options, NSRange _range) {
    NSRegularExpression* nSRegularExpression = (NSRegularExpression*)ptr;
    NSUInteger result_ = [nSRegularExpression numberOfMatchesInString:(NSString*)_string options:options range:_range];
    return result_;
}

Array C_NSRegularExpression_MatchesInString_Options_Range(void* ptr, void* _string, unsigned int options, NSRange _range) {
    NSRegularExpression* nSRegularExpression = (NSRegularExpression*)ptr;
    NSArray* result_ = [nSRegularExpression matchesInString:(NSString*)_string options:options range:_range];
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

void* C_NSRegularExpression_FirstMatchInString_Options_Range(void* ptr, void* _string, unsigned int options, NSRange _range) {
    NSRegularExpression* nSRegularExpression = (NSRegularExpression*)ptr;
    NSTextCheckingResult* result_ = [nSRegularExpression firstMatchInString:(NSString*)_string options:options range:_range];
    return result_;
}

NSRange C_NSRegularExpression_RangeOfFirstMatchInString_Options_Range(void* ptr, void* _string, unsigned int options, NSRange _range) {
    NSRegularExpression* nSRegularExpression = (NSRegularExpression*)ptr;
    NSRange result_ = [nSRegularExpression rangeOfFirstMatchInString:(NSString*)_string options:options range:_range];
    return result_;
}

unsigned int C_NSRegularExpression_ReplaceMatchesInString_Options_Range_WithTemplate(void* ptr, void* _string, unsigned int options, NSRange _range, void* templ) {
    NSRegularExpression* nSRegularExpression = (NSRegularExpression*)ptr;
    NSUInteger result_ = [nSRegularExpression replaceMatchesInString:(NSMutableString*)_string options:options range:_range withTemplate:(NSString*)templ];
    return result_;
}

void* C_NSRegularExpression_StringByReplacingMatchesInString_Options_Range_WithTemplate(void* ptr, void* _string, unsigned int options, NSRange _range, void* templ) {
    NSRegularExpression* nSRegularExpression = (NSRegularExpression*)ptr;
    NSString* result_ = [nSRegularExpression stringByReplacingMatchesInString:(NSString*)_string options:options range:_range withTemplate:(NSString*)templ];
    return result_;
}

void* C_NSRegularExpression_RegularExpression_EscapedTemplateForString(void* _string) {
    NSString* result_ = [NSRegularExpression escapedTemplateForString:(NSString*)_string];
    return result_;
}

void* C_NSRegularExpression_RegularExpression_EscapedPatternForString(void* _string) {
    NSString* result_ = [NSRegularExpression escapedPatternForString:(NSString*)_string];
    return result_;
}

void* C_NSRegularExpression_ReplacementStringForResult_InString_Offset_Template(void* ptr, void* result, void* _string, int offset, void* templ) {
    NSRegularExpression* nSRegularExpression = (NSRegularExpression*)ptr;
    NSString* result_ = [nSRegularExpression replacementStringForResult:(NSTextCheckingResult*)result inString:(NSString*)_string offset:offset template:(NSString*)templ];
    return result_;
}

void* C_NSRegularExpression_Pattern(void* ptr) {
    NSRegularExpression* nSRegularExpression = (NSRegularExpression*)ptr;
    NSString* result_ = [nSRegularExpression pattern];
    return result_;
}

unsigned int C_NSRegularExpression_Options(void* ptr) {
    NSRegularExpression* nSRegularExpression = (NSRegularExpression*)ptr;
    NSRegularExpressionOptions result_ = [nSRegularExpression options];
    return result_;
}

unsigned int C_NSRegularExpression_NumberOfCaptureGroups(void* ptr) {
    NSRegularExpression* nSRegularExpression = (NSRegularExpression*)ptr;
    NSUInteger result_ = [nSRegularExpression numberOfCaptureGroups];
    return result_;
}
