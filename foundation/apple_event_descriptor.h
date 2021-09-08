#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_AppleEventDescriptor_Alloc();

void* C_NSAppleEventDescriptor_InitListDescriptor(void* ptr);
void* C_NSAppleEventDescriptor_InitRecordDescriptor(void* ptr);
void* C_NSAppleEventDescriptor_AllocAppleEventDescriptor();
void* C_NSAppleEventDescriptor_Init(void* ptr);
void* C_NSAppleEventDescriptor_NewAppleEventDescriptor();
void* C_NSAppleEventDescriptor_Autorelease(void* ptr);
void* C_NSAppleEventDescriptor_Retain(void* ptr);
void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithString(void* _string);
void* C_NSAppleEventDescriptor_AppleEventDescriptor_ListDescriptor();
void* C_NSAppleEventDescriptor_AppleEventDescriptor_NullDescriptor();
void* C_NSAppleEventDescriptor_AppleEventDescriptor_RecordDescriptor();
void* C_NSAppleEventDescriptor_DescriptorAtIndex(void* ptr, int index);
void C_NSAppleEventDescriptor_InsertDescriptor_AtIndex(void* ptr, void* descriptor, int index);
void C_NSAppleEventDescriptor_RemoveDescriptorAtIndex(void* ptr, int index);
void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithApplicationURL(void* applicationURL);
void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithBundleIdentifier(void* bundleIdentifier);
void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithDate(void* date);
void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithDouble(double doubleValue);
void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithFileURL(void* fileURL);
void* C_NSAppleEventDescriptor_AppleEventDescriptor_CurrentProcessDescriptor();
void* C_NSAppleEventDescriptor_Data(void* ptr);
int C_NSAppleEventDescriptor_NumberOfItems(void* ptr);
void* C_NSAppleEventDescriptor_StringValue(void* ptr);
void* C_NSAppleEventDescriptor_DateValue(void* ptr);
double C_NSAppleEventDescriptor_DoubleValue(void* ptr);
void* C_NSAppleEventDescriptor_FileURLValue(void* ptr);
bool C_NSAppleEventDescriptor_IsRecordDescriptor(void* ptr);
