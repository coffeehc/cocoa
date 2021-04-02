#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "box.h"

NSRect Box_BorderRect(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box borderRect];
}

unsigned long Box_BoxType(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box boxType];
}

void Box_SetBoxType(void* ptr, unsigned long boxType) {
	NSBox* box = (NSBox*)ptr;
	[box setBoxType:boxType];
}

bool Box_IsTransparent(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box isTransparent];
}

void Box_SetTransparent(void* ptr, bool transparent) {
	NSBox* box = (NSBox*)ptr;
	[box setTransparent:transparent];
}

const char* Box_Title(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [[box title] UTF8String];
}

void Box_SetTitle(void* ptr, const char* title) {
	NSBox* box = (NSBox*)ptr;
	[box setTitle:[NSString stringWithUTF8String:title]];
}

void* Box_TitleFont(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box titleFont];
}

void Box_SetTitleFont(void* ptr, void* titleFont) {
	NSBox* box = (NSBox*)ptr;
	[box setTitleFont:(NSFont*)titleFont];
}

unsigned long Box_TitlePosition(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box titlePosition];
}

void Box_SetTitlePosition(void* ptr, unsigned long titlePosition) {
	NSBox* box = (NSBox*)ptr;
	[box setTitlePosition:titlePosition];
}

void* Box_TitleCell(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box titleCell];
}

NSRect Box_TitleRect(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box titleRect];
}

void* Box_BorderColor(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box borderColor];
}

void Box_SetBorderColor(void* ptr, void* borderColor) {
	NSBox* box = (NSBox*)ptr;
	[box setBorderColor:(NSColor*)borderColor];
}

double Box_BorderWidth(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box borderWidth];
}

void Box_SetBorderWidth(void* ptr, double borderWidth) {
	NSBox* box = (NSBox*)ptr;
	[box setBorderWidth:borderWidth];
}

double Box_CornerRadius(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box cornerRadius];
}

void Box_SetCornerRadius(void* ptr, double cornerRadius) {
	NSBox* box = (NSBox*)ptr;
	[box setCornerRadius:cornerRadius];
}

void* Box_FillColor(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box fillColor];
}

void Box_SetFillColor(void* ptr, void* fillColor) {
	NSBox* box = (NSBox*)ptr;
	[box setFillColor:(NSColor*)fillColor];
}

void* Box_ContentView(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box contentView];
}

void Box_SetContentView(void* ptr, void* contentView) {
	NSBox* box = (NSBox*)ptr;
	[box setContentView:(NSView*)contentView];
}

NSSize Box_ContentViewMargins(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	return [box contentViewMargins];
}

void Box_SetContentViewMargins(void* ptr, NSSize contentViewMargins) {
	NSBox* box = (NSBox*)ptr;
	[box setContentViewMargins:contentViewMargins];
}

void* Box_NewBox(NSRect frame) {
	NSBox* box = [NSBox alloc];
	return [[box initWithFrame:frame] autorelease];
}

void Box_SetFrameFromContentFrame(void* ptr, NSRect contentFrame) {
	NSBox* box = (NSBox*)ptr;
	[box setFrameFromContentFrame:contentFrame];
}

void Box_SizeToFit(void* ptr) {
	NSBox* box = (NSBox*)ptr;
	[box sizeToFit];
}
