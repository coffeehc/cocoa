#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "image.h"

NSSize Image_Size(void* ptr) {
	NSImage* image = (NSImage*)ptr;
	return [image size];
}

void Image_SetSize(void* ptr, NSSize size) {
	NSImage* image = (NSImage*)ptr;
	[image setSize:size];
}

bool Image_IsTemplate(void* ptr) {
	NSImage* image = (NSImage*)ptr;
	return [image isTemplate];
}

void Image_SetTemplate(void* ptr, bool template) {
	NSImage* image = (NSImage*)ptr;
	[image setTemplate:template];
}

const char* Image_AccessibilityDescription(void* ptr) {
	NSImage* image = (NSImage*)ptr;
	return [[image accessibilityDescription] UTF8String];
}

void Image_SetAccessibilityDescription(void* ptr, const char* accessibilityDescription) {
	NSImage* image = (NSImage*)ptr;
	[image setAccessibilityDescription:[NSString stringWithUTF8String:accessibilityDescription]];
}

bool Image_MatchesOnlyOnBestFittingAxis(void* ptr) {
	NSImage* image = (NSImage*)ptr;
	return [image matchesOnlyOnBestFittingAxis];
}

void Image_SetMatchesOnlyOnBestFittingAxis(void* ptr, bool matchesOnlyOnBestFittingAxis) {
	NSImage* image = (NSImage*)ptr;
	[image setMatchesOnlyOnBestFittingAxis:matchesOnlyOnBestFittingAxis];
}

Array Image_ImageTypes() {
	NSArray* ns_array = [NSImage imageTypes];
	int count = [ns_array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = (void*)[[ns_array objectAtIndex:i] UTF8String];}
	Array array;
	array.data = data;
	array.len = count;
	return array;
}

Array Image_ImageUnfilteredTypes() {
	NSArray* ns_array = [NSImage imageUnfilteredTypes];
	int count = [ns_array count];
	void** data = malloc(count * sizeof(void*));
	for (int i = 0; i < count; i++) {
		 data[i] = (void*)[[ns_array objectAtIndex:i] UTF8String];}
	Array array;
	array.data = data;
	array.len = count;
	return array;
}

void* Image_ImageNamed(const char* name) {
	return [NSImage imageNamed:[NSString stringWithUTF8String:name]];
}

const char* Image_Name(void* ptr) {
	NSImage* image = (NSImage*)ptr;
	return [[image name] UTF8String];
}

void Image_SetName(void* ptr, const char* name) {
	NSImage* image = (NSImage*)ptr;
	[image setName:[NSString stringWithUTF8String:name]];
}

void* Image_ImageWithSystemSymbolName(const char* symbolName, const char* description) {
	return [NSImage imageWithSystemSymbolName:[NSString stringWithUTF8String:symbolName] accessibilityDescription:[NSString stringWithUTF8String:description]];
}

void Image_ImageWithSymbolConfiguration(void* ptr, void* configuration) {
	NSImage* image = (NSImage*)ptr;
	[image imageWithSymbolConfiguration:(NSImageSymbolConfiguration*)configuration];
}

void* Image_InitByReferencingFile(void* ptr, const char* fileName) {
	NSImage* image = (NSImage*)ptr;
	return [image initByReferencingFile:[NSString stringWithUTF8String:fileName]];
}

void* Image_InitByReferencingURL(void* ptr, void* url) {
	NSImage* image = (NSImage*)ptr;
	return [image initByReferencingURL:(NSURL*)url];
}

void* Image_InitWithContentsOfFile(void* ptr, const char* fileName) {
	NSImage* image = (NSImage*)ptr;
	return [image initWithContentsOfFile:[NSString stringWithUTF8String:fileName]];
}

void* Image_InitWithContentsOfURL(void* ptr, void* fileName) {
	NSImage* image = (NSImage*)ptr;
	return [image initWithContentsOfURL:(NSURL*)fileName];
}

void* Image_InitWithData(void* ptr, Array data) {
	NSImage* image = (NSImage*)ptr;
	return [image initWithData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
}

void* Image_InitWithDataIgnoringOrientation(void* ptr, Array data) {
	NSImage* image = (NSImage*)ptr;
	return [image initWithDataIgnoringOrientation:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
}

void* Image_InitWithSize(void* ptr, NSSize szie) {
	NSImage* image = (NSImage*)ptr;
	return [image initWithSize:szie];
}
