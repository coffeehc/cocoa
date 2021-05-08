#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

void* C_Coder_Alloc();

void* C_NSCoder_Init(void* ptr);
bool C_NSCoder_ContainsValueForKey(void* ptr, void* key);
void C_NSCoder_EncodeBool_ForKey(void* ptr, bool value, void* key);
void C_NSCoder_EncodeBycopyObject(void* ptr, void* anObject);
void C_NSCoder_EncodeByrefObject(void* ptr, void* anObject);
void C_NSCoder_EncodeConditionalObject(void* ptr, void* object);
void C_NSCoder_EncodeConditionalObject_ForKey(void* ptr, void* object, void* key);
void C_NSCoder_EncodeDataObject(void* ptr, Array data);
void C_NSCoder_EncodeDouble_ForKey(void* ptr, double value, void* key);
void C_NSCoder_EncodeInteger_ForKey(void* ptr, int value, void* key);
void C_NSCoder_EncodeInt32_ForKey(void* ptr, int32_t value, void* key);
void C_NSCoder_EncodeObject(void* ptr, void* object);
void C_NSCoder_EncodeObject_ForKey(void* ptr, void* object, void* key);
void C_NSCoder_EncodePropertyList(void* ptr, void* aPropertyList);
void C_NSCoder_EncodeRootObject(void* ptr, void* rootObject);
bool C_NSCoder_DecodeBoolForKey(void* ptr, void* key);
Array C_NSCoder_DecodeDataObject(void* ptr);
double C_NSCoder_DecodeDoubleForKey(void* ptr, void* key);
int C_NSCoder_DecodeIntegerForKey(void* ptr, void* key);
int32_t C_NSCoder_DecodeInt32ForKey(void* ptr, void* key);
void* C_NSCoder_DecodeObject(void* ptr);
void* C_NSCoder_DecodeObjectForKey(void* ptr, void* key);
void* C_NSCoder_DecodePropertyList(void* ptr);
void* C_NSCoder_DecodePropertyListForKey(void* ptr, void* key);
void C_NSCoder_FailWithError(void* ptr, void* error);
int C_NSCoder_VersionForClassName(void* ptr, void* className);
bool C_NSCoder_AllowsKeyedCoding(void* ptr);
bool C_NSCoder_RequiresSecureCoding(void* ptr);
void* C_NSCoder_Error(void* ptr);
