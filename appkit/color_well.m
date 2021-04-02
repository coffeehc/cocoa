#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "color_well.h"

void* ColorWell_Color(void* ptr) {
	NSColorWell* colorWell = (NSColorWell*)ptr;
	return [colorWell color];
}

void ColorWell_SetColor(void* ptr, void* color) {
	NSColorWell* colorWell = (NSColorWell*)ptr;
	[colorWell setColor:(NSColor*)color];
}

bool ColorWell_IsActive(void* ptr) {
	NSColorWell* colorWell = (NSColorWell*)ptr;
	return [colorWell isActive];
}

bool ColorWell_IsBordered(void* ptr) {
	NSColorWell* colorWell = (NSColorWell*)ptr;
	return [colorWell isBordered];
}

void ColorWell_SetBordered(void* ptr, bool bordered) {
	NSColorWell* colorWell = (NSColorWell*)ptr;
	[colorWell setBordered:bordered];
}

void* ColorWell_NewColorWell(NSRect frame) {
	NSColorWell* colorWell = [NSColorWell alloc];
	return [[colorWell initWithFrame:frame] autorelease];
}

void ColorWell_TakeColorFrom(void* ptr, void* sender) {
	NSColorWell* colorWell = (NSColorWell*)ptr;
	[colorWell takeColorFrom:(NSObject*)sender];
}

void ColorWell_Activate(void* ptr, bool exclusive) {
	NSColorWell* colorWell = (NSColorWell*)ptr;
	[colorWell activate:exclusive];
}

void ColorWell_Deactivate(void* ptr) {
	NSColorWell* colorWell = (NSColorWell*)ptr;
	[colorWell deactivate];
}

void ColorWell_DrawWellInside(void* ptr, NSRect insideRect) {
	NSColorWell* colorWell = (NSColorWell*)ptr;
	[colorWell drawWellInside:insideRect];
}
