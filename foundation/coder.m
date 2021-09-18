#import "coder.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSCoder.h>

void* C_Coder_Alloc() {
    return [NSCoder alloc];
}

void* C_NSCoder_AllocCoder() {
    NSCoder* result_ = [NSCoder alloc];
    return result_;
}

void* C_NSCoder_Init(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSCoder* result_ = [nSCoder init];
    return result_;
}

void* C_NSCoder_NewCoder() {
    NSCoder* result_ = [NSCoder new];
    return result_;
}

void* C_NSCoder_Autorelease(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSCoder* result_ = [nSCoder autorelease];
    return result_;
}

void* C_NSCoder_Retain(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSCoder* result_ = [nSCoder retain];
    return result_;
}

bool C_NSCoder_ContainsValueForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    BOOL result_ = [nSCoder containsValueForKey:(NSString*)key];
    return result_;
}

void C_NSCoder_EncodeBool_ForKey(void* ptr, bool value, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeBool:value forKey:(NSString*)key];
}

void C_NSCoder_EncodeBycopyObject(void* ptr, void* anObject) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeBycopyObject:(id)anObject];
}

void C_NSCoder_EncodeByrefObject(void* ptr, void* anObject) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeByrefObject:(id)anObject];
}

void C_NSCoder_EncodeConditionalObject(void* ptr, void* object) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeConditionalObject:(id)object];
}

void C_NSCoder_EncodeConditionalObject_ForKey(void* ptr, void* object, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeConditionalObject:(id)object forKey:(NSString*)key];
}

void C_NSCoder_EncodeDataObject(void* ptr, void* data) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeDataObject:(NSData*)data];
}

void C_NSCoder_EncodeDouble_ForKey(void* ptr, double value, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeDouble:value forKey:(NSString*)key];
}

void C_NSCoder_EncodeFloat_ForKey(void* ptr, float value, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeFloat:value forKey:(NSString*)key];
}

void C_NSCoder_EncodeInteger_ForKey(void* ptr, int value, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeInteger:value forKey:(NSString*)key];
}

void C_NSCoder_EncodeInt32_ForKey(void* ptr, int32_t value, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeInt32:value forKey:(NSString*)key];
}

void C_NSCoder_EncodeInt64_ForKey(void* ptr, long value, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeInt64:value forKey:(NSString*)key];
}

void C_NSCoder_EncodeObject(void* ptr, void* object) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeObject:(id)object];
}

void C_NSCoder_EncodeObject_ForKey(void* ptr, void* object, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeObject:(id)object forKey:(NSString*)key];
}

void C_NSCoder_EncodePoint(void* ptr, CGPoint point) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodePoint:point];
}

void C_NSCoder_EncodePoint_ForKey(void* ptr, CGPoint point, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodePoint:point forKey:(NSString*)key];
}

void C_NSCoder_EncodePropertyList(void* ptr, void* aPropertyList) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodePropertyList:(id)aPropertyList];
}

void C_NSCoder_EncodeRect(void* ptr, CGRect rect) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeRect:rect];
}

void C_NSCoder_EncodeRect_ForKey(void* ptr, CGRect rect, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeRect:rect forKey:(NSString*)key];
}

void C_NSCoder_EncodeRootObject(void* ptr, void* rootObject) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeRootObject:(id)rootObject];
}

void C_NSCoder_EncodeSize(void* ptr, CGSize size) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeSize:size];
}

void C_NSCoder_EncodeSize_ForKey(void* ptr, CGSize size, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder encodeSize:size forKey:(NSString*)key];
}

bool C_NSCoder_DecodeBoolForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    BOOL result_ = [nSCoder decodeBoolForKey:(NSString*)key];
    return result_;
}

void* C_NSCoder_DecodeDataObject(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSData* result_ = [nSCoder decodeDataObject];
    return result_;
}

double C_NSCoder_DecodeDoubleForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    double result_ = [nSCoder decodeDoubleForKey:(NSString*)key];
    return result_;
}

float C_NSCoder_DecodeFloatForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    float result_ = [nSCoder decodeFloatForKey:(NSString*)key];
    return result_;
}

int C_NSCoder_DecodeIntegerForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSInteger result_ = [nSCoder decodeIntegerForKey:(NSString*)key];
    return result_;
}

int32_t C_NSCoder_DecodeInt32ForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    int32_t result_ = [nSCoder decodeInt32ForKey:(NSString*)key];
    return result_;
}

long C_NSCoder_DecodeInt64ForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSInteger result_ = [nSCoder decodeInt64ForKey:(NSString*)key];
    return result_;
}

void* C_NSCoder_DecodeObject(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    id result_ = [nSCoder decodeObject];
    return result_;
}

void* C_NSCoder_DecodeObjectForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    id result_ = [nSCoder decodeObjectForKey:(NSString*)key];
    return result_;
}

CGPoint C_NSCoder_DecodePoint(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSPoint result_ = [nSCoder decodePoint];
    return result_;
}

CGPoint C_NSCoder_DecodePointForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSPoint result_ = [nSCoder decodePointForKey:(NSString*)key];
    return result_;
}

void* C_NSCoder_DecodePropertyList(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    id result_ = [nSCoder decodePropertyList];
    return result_;
}

CGRect C_NSCoder_DecodeRect(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSRect result_ = [nSCoder decodeRect];
    return result_;
}

CGRect C_NSCoder_DecodeRectForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSRect result_ = [nSCoder decodeRectForKey:(NSString*)key];
    return result_;
}

CGSize C_NSCoder_DecodeSize(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSSize result_ = [nSCoder decodeSize];
    return result_;
}

CGSize C_NSCoder_DecodeSizeForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSSize result_ = [nSCoder decodeSizeForKey:(NSString*)key];
    return result_;
}

void* C_NSCoder_DecodeObjectOfClasses_ForKey(void* ptr, void* classes, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    id result_ = [nSCoder decodeObjectOfClasses:(NSSet*)classes forKey:(NSString*)key];
    return result_;
}

void* C_NSCoder_DecodePropertyListForKey(void* ptr, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    id result_ = [nSCoder decodePropertyListForKey:(NSString*)key];
    return result_;
}

void C_NSCoder_FailWithError(void* ptr, void* error) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    [nSCoder failWithError:(NSError*)error];
}

int C_NSCoder_VersionForClassName(void* ptr, void* className) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSInteger result_ = [nSCoder versionForClassName:(NSString*)className];
    return result_;
}

Array C_NSCoder_DecodeArrayOfObjectsOfClasses_ForKey(void* ptr, void* classes, void* key) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSArray* result_ = [nSCoder decodeArrayOfObjectsOfClasses:(NSSet*)classes forKey:(NSString*)key];
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

bool C_NSCoder_AllowsKeyedCoding(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    BOOL result_ = [nSCoder allowsKeyedCoding];
    return result_;
}

int C_NSCoder_DecodingFailurePolicy(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSDecodingFailurePolicy result_ = [nSCoder decodingFailurePolicy];
    return result_;
}

bool C_NSCoder_RequiresSecureCoding(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    BOOL result_ = [nSCoder requiresSecureCoding];
    return result_;
}

void* C_NSCoder_AllowedClasses(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSSet* result_ = [nSCoder allowedClasses];
    return result_;
}

void* C_NSCoder_Error(void* ptr) {
    NSCoder* nSCoder = (NSCoder*)ptr;
    NSError* result_ = [nSCoder error];
    return result_;
}
