#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_Coder_Alloc();

void* C_NSCoder_Init(void* ptr);
bool C_NSCoder_ContainsValueForKey(void* ptr, void* key);
void C_NSCoder_EncodeBool_ForKey(void* ptr, bool value, void* key);
void C_NSCoder_EncodeBycopyObject(void* ptr, void* anObject);
void C_NSCoder_EncodeByrefObject(void* ptr, void* anObject);
void C_NSCoder_EncodeConditionalObject(void* ptr, void* object);
void C_NSCoder_EncodeConditionalObject_ForKey(void* ptr, void* object, void* key);
void C_NSCoder_EncodeDataObject(void* ptr, void* data);
void C_NSCoder_EncodeDouble_ForKey(void* ptr, double value, void* key);
void C_NSCoder_EncodeFloat_ForKey(void* ptr, float value, void* key);
void C_NSCoder_EncodeInteger_ForKey(void* ptr, int value, void* key);
void C_NSCoder_EncodeInt32_ForKey(void* ptr, int32_t value, void* key);
void C_NSCoder_EncodeInt64_ForKey(void* ptr, long value, void* key);
void C_NSCoder_EncodeObject(void* ptr, void* object);
void C_NSCoder_EncodeObject_ForKey(void* ptr, void* object, void* key);
void C_NSCoder_EncodePoint(void* ptr, CGPoint point);
void C_NSCoder_EncodePoint_ForKey(void* ptr, CGPoint point, void* key);
void C_NSCoder_EncodePropertyList(void* ptr, void* aPropertyList);
void C_NSCoder_EncodeRect(void* ptr, CGRect rect);
void C_NSCoder_EncodeRect_ForKey(void* ptr, CGRect rect, void* key);
void C_NSCoder_EncodeRootObject(void* ptr, void* rootObject);
void C_NSCoder_EncodeSize(void* ptr, CGSize size);
void C_NSCoder_EncodeSize_ForKey(void* ptr, CGSize size, void* key);
bool C_NSCoder_DecodeBoolForKey(void* ptr, void* key);
void* C_NSCoder_DecodeDataObject(void* ptr);
double C_NSCoder_DecodeDoubleForKey(void* ptr, void* key);
float C_NSCoder_DecodeFloatForKey(void* ptr, void* key);
int C_NSCoder_DecodeIntegerForKey(void* ptr, void* key);
int32_t C_NSCoder_DecodeInt32ForKey(void* ptr, void* key);
long C_NSCoder_DecodeInt64ForKey(void* ptr, void* key);
void* C_NSCoder_DecodeObject(void* ptr);
void* C_NSCoder_DecodeObjectForKey(void* ptr, void* key);
CGPoint C_NSCoder_DecodePoint(void* ptr);
CGPoint C_NSCoder_DecodePointForKey(void* ptr, void* key);
void* C_NSCoder_DecodePropertyList(void* ptr);
CGRect C_NSCoder_DecodeRect(void* ptr);
CGRect C_NSCoder_DecodeRectForKey(void* ptr, void* key);
CGSize C_NSCoder_DecodeSize(void* ptr);
CGSize C_NSCoder_DecodeSizeForKey(void* ptr, void* key);
void* C_NSCoder_DecodeObjectOfClasses_ForKey(void* ptr, void* classes, void* key);
void* C_NSCoder_DecodePropertyListForKey(void* ptr, void* key);
void C_NSCoder_FailWithError(void* ptr, void* error);
int C_NSCoder_VersionForClassName(void* ptr, void* className);
Array C_NSCoder_DecodeArrayOfObjectsOfClasses_ForKey(void* ptr, void* classes, void* key);
bool C_NSCoder_AllowsKeyedCoding(void* ptr);
int C_NSCoder_DecodingFailurePolicy(void* ptr);
bool C_NSCoder_RequiresSecureCoding(void* ptr);
void* C_NSCoder_AllowedClasses(void* ptr);
void* C_NSCoder_Error(void* ptr);
