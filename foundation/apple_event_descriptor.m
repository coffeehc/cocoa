#import "apple_event_descriptor.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSAppleEventDescriptor.h>

void* C_AppleEventDescriptor_Alloc() {
    return [NSAppleEventDescriptor alloc];
}

void* C_NSAppleEventDescriptor_InitListDescriptor(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    NSAppleEventDescriptor* result_ = [nSAppleEventDescriptor initListDescriptor];
    return result_;
}

void* C_NSAppleEventDescriptor_InitRecordDescriptor(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    NSAppleEventDescriptor* result_ = [nSAppleEventDescriptor initRecordDescriptor];
    return result_;
}

void* C_NSAppleEventDescriptor_AllocAppleEventDescriptor() {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor alloc];
    return result_;
}

void* C_NSAppleEventDescriptor_Init(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    NSAppleEventDescriptor* result_ = [nSAppleEventDescriptor init];
    return result_;
}

void* C_NSAppleEventDescriptor_NewAppleEventDescriptor() {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor new];
    return result_;
}

void* C_NSAppleEventDescriptor_Autorelease(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    NSAppleEventDescriptor* result_ = [nSAppleEventDescriptor autorelease];
    return result_;
}

void* C_NSAppleEventDescriptor_Retain(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    NSAppleEventDescriptor* result_ = [nSAppleEventDescriptor retain];
    return result_;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithString(void* _string) {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor descriptorWithString:(NSString*)_string];
    return result_;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptor_ListDescriptor() {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor listDescriptor];
    return result_;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptor_NullDescriptor() {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor nullDescriptor];
    return result_;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptor_RecordDescriptor() {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor recordDescriptor];
    return result_;
}

void* C_NSAppleEventDescriptor_DescriptorAtIndex(void* ptr, int index) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    NSAppleEventDescriptor* result_ = [nSAppleEventDescriptor descriptorAtIndex:index];
    return result_;
}

void C_NSAppleEventDescriptor_InsertDescriptor_AtIndex(void* ptr, void* descriptor, int index) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    [nSAppleEventDescriptor insertDescriptor:(NSAppleEventDescriptor*)descriptor atIndex:index];
}

void C_NSAppleEventDescriptor_RemoveDescriptorAtIndex(void* ptr, int index) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    [nSAppleEventDescriptor removeDescriptorAtIndex:index];
}

void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithApplicationURL(void* applicationURL) {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor descriptorWithApplicationURL:(NSURL*)applicationURL];
    return result_;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithBundleIdentifier(void* bundleIdentifier) {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor descriptorWithBundleIdentifier:(NSString*)bundleIdentifier];
    return result_;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithDate(void* date) {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor descriptorWithDate:(NSDate*)date];
    return result_;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithDouble(double doubleValue) {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor descriptorWithDouble:doubleValue];
    return result_;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptor_DescriptorWithFileURL(void* fileURL) {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor descriptorWithFileURL:(NSURL*)fileURL];
    return result_;
}

void* C_NSAppleEventDescriptor_AppleEventDescriptor_CurrentProcessDescriptor() {
    NSAppleEventDescriptor* result_ = [NSAppleEventDescriptor currentProcessDescriptor];
    return result_;
}

void* C_NSAppleEventDescriptor_Data(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    NSData* result_ = [nSAppleEventDescriptor data];
    return result_;
}

int C_NSAppleEventDescriptor_NumberOfItems(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    NSInteger result_ = [nSAppleEventDescriptor numberOfItems];
    return result_;
}

void* C_NSAppleEventDescriptor_StringValue(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    NSString* result_ = [nSAppleEventDescriptor stringValue];
    return result_;
}

void* C_NSAppleEventDescriptor_DateValue(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    NSDate* result_ = [nSAppleEventDescriptor dateValue];
    return result_;
}

double C_NSAppleEventDescriptor_DoubleValue(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    double result_ = [nSAppleEventDescriptor doubleValue];
    return result_;
}

void* C_NSAppleEventDescriptor_FileURLValue(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    NSURL* result_ = [nSAppleEventDescriptor fileURLValue];
    return result_;
}

bool C_NSAppleEventDescriptor_IsRecordDescriptor(void* ptr) {
    NSAppleEventDescriptor* nSAppleEventDescriptor = (NSAppleEventDescriptor*)ptr;
    BOOL result_ = [nSAppleEventDescriptor isRecordDescriptor];
    return result_;
}
