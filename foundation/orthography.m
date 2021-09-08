#import <Foundation/Foundation.h>
#import "orthography.h"

void* C_Orthography_Alloc() {
    return [NSOrthography alloc];
}

void* C_NSOrthography_DefaultOrthographyForLanguage(void* language) {
    NSOrthography* result_ = [NSOrthography defaultOrthographyForLanguage:(NSString*)language];
    return result_;
}

void* C_NSOrthography_InitWithCoder(void* ptr, void* coder) {
    NSOrthography* nSOrthography = (NSOrthography*)ptr;
    NSOrthography* result_ = [nSOrthography initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSOrthography_AllocOrthography() {
    NSOrthography* result_ = [NSOrthography alloc];
    return result_;
}

void* C_NSOrthography_Init(void* ptr) {
    NSOrthography* nSOrthography = (NSOrthography*)ptr;
    NSOrthography* result_ = [nSOrthography init];
    return result_;
}

void* C_NSOrthography_NewOrthography() {
    NSOrthography* result_ = [NSOrthography new];
    return result_;
}

void* C_NSOrthography_Autorelease(void* ptr) {
    NSOrthography* nSOrthography = (NSOrthography*)ptr;
    NSOrthography* result_ = [nSOrthography autorelease];
    return result_;
}

void* C_NSOrthography_Retain(void* ptr) {
    NSOrthography* nSOrthography = (NSOrthography*)ptr;
    NSOrthography* result_ = [nSOrthography retain];
    return result_;
}

void* C_NSOrthography_DominantLanguageForScript(void* ptr, void* script) {
    NSOrthography* nSOrthography = (NSOrthography*)ptr;
    NSString* result_ = [nSOrthography dominantLanguageForScript:(NSString*)script];
    return result_;
}

Array C_NSOrthography_LanguagesForScript(void* ptr, void* script) {
    NSOrthography* nSOrthography = (NSOrthography*)ptr;
    NSArray* result_ = [nSOrthography languagesForScript:(NSString*)script];
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

void* C_NSOrthography_DominantLanguage(void* ptr) {
    NSOrthography* nSOrthography = (NSOrthography*)ptr;
    NSString* result_ = [nSOrthography dominantLanguage];
    return result_;
}

void* C_NSOrthography_DominantScript(void* ptr) {
    NSOrthography* nSOrthography = (NSOrthography*)ptr;
    NSString* result_ = [nSOrthography dominantScript];
    return result_;
}

Array C_NSOrthography_AllScripts(void* ptr) {
    NSOrthography* nSOrthography = (NSOrthography*)ptr;
    NSArray* result_ = [nSOrthography allScripts];
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

Array C_NSOrthography_AllLanguages(void* ptr) {
    NSOrthography* nSOrthography = (NSOrthography*)ptr;
    NSArray* result_ = [nSOrthography allLanguages];
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
