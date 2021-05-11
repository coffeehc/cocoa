#import <Appkit/Appkit.h>
#import "image_rep.h"

void* C_ImageRep_Alloc() {
	return [NSImageRep alloc];
}

void* C_NSImageRep_Init(void* ptr) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	NSImageRep* result = [nSImageRep init];
	return result;
}

void* C_NSImageRep_InitWithCoder(void* ptr, void* coder) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	NSImageRep* result = [nSImageRep initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSImageRep_ImageRepWithContentsOfFile(void* filename) {
	NSImageRep* result = [NSImageRep imageRepWithContentsOfFile:(NSString*)filename];
	return result;
}

void* C_NSImageRep_ImageRepWithPasteboard(void* pasteboard) {
	NSImageRep* result = [NSImageRep imageRepWithPasteboard:(NSPasteboard*)pasteboard];
	return result;
}

void* C_NSImageRep_ImageRepWithContentsOfURL(void* url) {
	NSImageRep* result = [NSImageRep imageRepWithContentsOfURL:(NSURL*)url];
	return result;
}

bool C_NSImageRep_ImageRepCanInitWithData(Array data) {
	bool result = [NSImageRep canInitWithData:[[NSData alloc] initWithBytes:(Byte *)data.data length:data.len]];
	return result;
}

bool C_NSImageRep_ImageRepCanInitWithPasteboard(void* pasteboard) {
	bool result = [NSImageRep canInitWithPasteboard:(NSPasteboard*)pasteboard];
	return result;
}

bool C_NSImageRep_Draw(void* ptr) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	bool result = [nSImageRep draw];
	return result;
}

bool C_NSImageRep_DrawAtPoint(void* ptr, CGPoint point) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	bool result = [nSImageRep drawAtPoint:point];
	return result;
}

bool C_NSImageRep_DrawInRect(void* ptr, CGRect rect) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	bool result = [nSImageRep drawInRect:rect];
	return result;
}

CGSize C_NSImageRep_Size(void* ptr) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	CGSize result = [nSImageRep size];
	return result;
}

void C_NSImageRep_SetSize(void* ptr, CGSize value) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	[nSImageRep setSize:value];
}

int C_NSImageRep_BitsPerSample(void* ptr) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	int result = [nSImageRep bitsPerSample];
	return result;
}

void C_NSImageRep_SetBitsPerSample(void* ptr, int value) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	[nSImageRep setBitsPerSample:value];
}

void* C_NSImageRep_ColorSpaceName(void* ptr) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	NSString* result = [nSImageRep colorSpaceName];
	return result;
}

void C_NSImageRep_SetColorSpaceName(void* ptr, void* value) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	[nSImageRep setColorSpaceName:(NSString*)value];
}

bool C_NSImageRep_HasAlpha(void* ptr) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	bool result = [nSImageRep hasAlpha];
	return result;
}

void C_NSImageRep_SetAlpha(void* ptr, bool value) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	[nSImageRep setAlpha:value];
}

bool C_NSImageRep_IsOpaque(void* ptr) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	bool result = [nSImageRep isOpaque];
	return result;
}

void C_NSImageRep_SetOpaque(void* ptr, bool value) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	[nSImageRep setOpaque:value];
}

int C_NSImageRep_PixelsHigh(void* ptr) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	int result = [nSImageRep pixelsHigh];
	return result;
}

void C_NSImageRep_SetPixelsHigh(void* ptr, int value) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	[nSImageRep setPixelsHigh:value];
}

int C_NSImageRep_PixelsWide(void* ptr) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	int result = [nSImageRep pixelsWide];
	return result;
}

void C_NSImageRep_SetPixelsWide(void* ptr, int value) {
	NSImageRep* nSImageRep = (NSImageRep*)ptr;
	[nSImageRep setPixelsWide:value];
}
