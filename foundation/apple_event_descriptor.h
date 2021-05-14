#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/Foundation.h>

void* C_AppleEventDescriptor_Alloc();

void* C_NSAppleEventDescriptor_InitListDescriptor(void* ptr);
void* C_NSAppleEventDescriptor_InitRecordDescriptor(void* ptr);
void* C_NSAppleEventDescriptor_Init(void* ptr);
void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithString(void* _string);
void* C_NSAppleEventDescriptor_AppleEventDescriptorListDescriptor();
void* C_NSAppleEventDescriptor_AppleEventDescriptorNullDescriptor();
void* C_NSAppleEventDescriptor_AppleEventDescriptorRecordDescriptor();
void* C_NSAppleEventDescriptor_DescriptorAtIndex(void* ptr, int index);
void C_NSAppleEventDescriptor_InsertDescriptor_AtIndex(void* ptr, void* descriptor, int index);
void C_NSAppleEventDescriptor_RemoveDescriptorAtIndex(void* ptr, int index);
void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithApplicationURL(void* applicationURL);
void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithBundleIdentifier(void* bundleIdentifier);
void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithDate(void* date);
void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithDouble(double doubleValue);
void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithFileURL(void* fileURL);
void* C_NSAppleEventDescriptor_AppleEventDescriptorCurrentProcessDescriptor();
Array C_NSAppleEventDescriptor_Data(void* ptr);
int C_NSAppleEventDescriptor_NumberOfItems(void* ptr);
void* C_NSAppleEventDescriptor_StringValue(void* ptr);
void* C_NSAppleEventDescriptor_DateValue(void* ptr);
double C_NSAppleEventDescriptor_DoubleValue(void* ptr);
void* C_NSAppleEventDescriptor_FileURLValue(void* ptr);
bool C_NSAppleEventDescriptor_IsRecordDescriptor(void* ptr);
