#import <Foundation/Foundation.h>
#import "apple_event_descriptor.h"

void* C_AppleEventDescriptor_Alloc() {
	return [NSAppleEventDescriptor alloc];
}

void* C_NSAppleEventDescriptor_InitListDescriptor(void* ptr) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	NSAppleEventDescriptor* result = [nSAppleEventDescriptor initListDescriptor];
	return result;
}

void* C_NSAppleEventDescriptor_InitRecordDescriptor(void* ptr) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	NSAppleEventDescriptor* result = [nSAppleEventDescriptor initRecordDescriptor];
	return result;
}

void* C_NSAppleEventDescriptor_Init(void* ptr) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	NSAppleEventDescriptor* result = [nSAppleEventDescriptor init];
	return result;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithString(void* _string) {
	NSAppleEventDescriptor* result = [NSAppleEventDescriptor descriptorWithString:(NSString*)_string];
	return result;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptorListDescriptor() {
	NSAppleEventDescriptor* result = [NSAppleEventDescriptor listDescriptor];
	return result;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptorNullDescriptor() {
	NSAppleEventDescriptor* result = [NSAppleEventDescriptor nullDescriptor];
	return result;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptorRecordDescriptor() {
	NSAppleEventDescriptor* result = [NSAppleEventDescriptor recordDescriptor];
	return result;
}

void* C_NSAppleEventDescriptor_DescriptorAtIndex(void* ptr, int index) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	NSAppleEventDescriptor* result = [nSAppleEventDescriptor descriptorAtIndex:index];
	return result;
}

void C_NSAppleEventDescriptor_InsertDescriptor_AtIndex(void* ptr, void* descriptor, int index) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	[nSAppleEventDescriptor insertDescriptor:(NSAppleEventDescriptor*)descriptor atIndex:index];
}

void C_NSAppleEventDescriptor_RemoveDescriptorAtIndex(void* ptr, int index) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	[nSAppleEventDescriptor removeDescriptorAtIndex:index];
}

void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithApplicationURL(void* applicationURL) {
	NSAppleEventDescriptor* result = [NSAppleEventDescriptor descriptorWithApplicationURL:(NSURL*)applicationURL];
	return result;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithBundleIdentifier(void* bundleIdentifier) {
	NSAppleEventDescriptor* result = [NSAppleEventDescriptor descriptorWithBundleIdentifier:(NSString*)bundleIdentifier];
	return result;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithDate(void* date) {
	NSAppleEventDescriptor* result = [NSAppleEventDescriptor descriptorWithDate:(NSDate*)date];
	return result;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithDouble(double doubleValue) {
	NSAppleEventDescriptor* result = [NSAppleEventDescriptor descriptorWithDouble:doubleValue];
	return result;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptorDescriptorWithFileURL(void* fileURL) {
	NSAppleEventDescriptor* result = [NSAppleEventDescriptor descriptorWithFileURL:(NSURL*)fileURL];
	return result;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptorCurrentProcessDescriptor() {
	NSAppleEventDescriptor* result = [NSAppleEventDescriptor currentProcessDescriptor];
	return result;
}

Array C_NSAppleEventDescriptor_Data(void* ptr) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	NSData* result = [nSAppleEventDescriptor data];
	Array resultarray;
	resultarray.data = [result bytes];
	resultarray.len = result.length;
	return resultarray;
}

int C_NSAppleEventDescriptor_NumberOfItems(void* ptr) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	int result = [nSAppleEventDescriptor numberOfItems];
	return result;
}

void* C_NSAppleEventDescriptor_StringValue(void* ptr) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	NSString* result = [nSAppleEventDescriptor stringValue];
	return result;
}

void* C_NSAppleEventDescriptor_DateValue(void* ptr) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	NSDate* result = [nSAppleEventDescriptor dateValue];
	return result;
}

double C_NSAppleEventDescriptor_DoubleValue(void* ptr) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	double result = [nSAppleEventDescriptor doubleValue];
	return result;
}

void* C_NSAppleEventDescriptor_FileURLValue(void* ptr) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	NSURL* result = [nSAppleEventDescriptor fileURLValue];
	return result;
}

bool C_NSAppleEventDescriptor_IsRecordDescriptor(void* ptr) {
	NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
	bool result = [nSAppleEventDescriptor isRecordDescriptor];
	return result;
}
