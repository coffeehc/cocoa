#import <Foundation/Foundation.h>
#import "url.h"

void* C_Coder_Alloc() {
	return [NSCoder alloc];
}

void* C_NSCoder_Init(void* ptr) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	NSCoder* result = [nSCoder init];
	return result;
}

bool C_NSCoder_ContainsValueForKey(void* ptr, void* key) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	bool result = [nSCoder containsValueForKey:(NSString*)key];
	return result;
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

void C_NSCoder_EncodeDataObject(void* ptr, Array data) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	[nSCoder encodeDataObject:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
}

void C_NSCoder_EncodeDouble_ForKey(void* ptr, double value, void* key) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	[nSCoder encodeDouble:value forKey:(NSString*)key];
}

void C_NSCoder_EncodeInteger_ForKey(void* ptr, int value, void* key) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	[nSCoder encodeInteger:value forKey:(NSString*)key];
}

void C_NSCoder_EncodeInt32_ForKey(void* ptr, int32_t value, void* key) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	[nSCoder encodeInt32:value forKey:(NSString*)key];
}

void C_NSCoder_EncodeObject(void* ptr, void* object) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	[nSCoder encodeObject:(id)object];
}

void C_NSCoder_EncodeObject_ForKey(void* ptr, void* object, void* key) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	[nSCoder encodeObject:(id)object forKey:(NSString*)key];
}

void C_NSCoder_EncodePropertyList(void* ptr, void* aPropertyList) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	[nSCoder encodePropertyList:(id)aPropertyList];
}

void C_NSCoder_EncodeRootObject(void* ptr, void* rootObject) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	[nSCoder encodeRootObject:(id)rootObject];
}

bool C_NSCoder_DecodeBoolForKey(void* ptr, void* key) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	bool result = [nSCoder decodeBoolForKey:(NSString*)key];
	return result;
}

Array C_NSCoder_DecodeDataObject(void* ptr) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	NSData* result = [nSCoder decodeDataObject];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

double C_NSCoder_DecodeDoubleForKey(void* ptr, void* key) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	double result = [nSCoder decodeDoubleForKey:(NSString*)key];
	return result;
}

int C_NSCoder_DecodeIntegerForKey(void* ptr, void* key) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	int result = [nSCoder decodeIntegerForKey:(NSString*)key];
	return result;
}

int32_t C_NSCoder_DecodeInt32ForKey(void* ptr, void* key) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	int32_t result = [nSCoder decodeInt32ForKey:(NSString*)key];
	return result;
}

void* C_NSCoder_DecodeObject(void* ptr) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	id result = [nSCoder decodeObject];
	return result;
}

void* C_NSCoder_DecodeObjectForKey(void* ptr, void* key) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	id result = [nSCoder decodeObjectForKey:(NSString*)key];
	return result;
}

void* C_NSCoder_DecodePropertyList(void* ptr) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	id result = [nSCoder decodePropertyList];
	return result;
}

void* C_NSCoder_DecodePropertyListForKey(void* ptr, void* key) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	id result = [nSCoder decodePropertyListForKey:(NSString*)key];
	return result;
}

void C_NSCoder_FailWithError(void* ptr, void* error) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	[nSCoder failWithError:(NSError*)error];
}

int C_NSCoder_VersionForClassName(void* ptr, void* className) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	int result = [nSCoder versionForClassName:(NSString*)className];
	return result;
}

bool C_NSCoder_AllowsKeyedCoding(void* ptr) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	bool result = [nSCoder allowsKeyedCoding];
	return result;
}

bool C_NSCoder_RequiresSecureCoding(void* ptr) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	bool result = [nSCoder requiresSecureCoding];
	return result;
}

void* C_NSCoder_Error(void* ptr) {
	NSCoder* nSCoder = (NSCoder*)ptr;
	NSError* result = [nSCoder error];
	return result;
}
