#import <Appkit/Appkit.h>
#import "ruler_marker.h"

void* C_RulerMarker_Alloc() {
	return [NSRulerMarker alloc];
}

void* C_NSRulerMarker_InitWithRulerView_MarkerLocation_Image_ImageOrigin(void* ptr, void* ruler, double location, void* image, CGPoint imageOrigin) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	NSRulerMarker* result = [nSRulerMarker initWithRulerView:(NSRulerView*)ruler markerLocation:location image:(NSImage*)image imageOrigin:imageOrigin];
	return result;
}

void* C_NSRulerMarker_InitWithCoder(void* ptr, void* coder) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	NSRulerMarker* result = [nSRulerMarker initWithCoder:(NSCoder*)coder];
	return result;
}

void C_NSRulerMarker_DrawRect(void* ptr, CGRect rect) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	[nSRulerMarker drawRect:rect];
}

bool C_NSRulerMarker_TrackMouse_Adding(void* ptr, void* mouseDownEvent, bool isAdding) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	bool result = [nSRulerMarker trackMouse:(NSEvent*)mouseDownEvent adding:isAdding];
	return result;
}

void* C_NSRulerMarker_Ruler(void* ptr) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	NSRulerView* result = [nSRulerMarker ruler];
	return result;
}

void* C_NSRulerMarker_Image(void* ptr) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	NSImage* result = [nSRulerMarker image];
	return result;
}

void C_NSRulerMarker_SetImage(void* ptr, void* value) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	[nSRulerMarker setImage:(NSImage*)value];
}

CGPoint C_NSRulerMarker_ImageOrigin(void* ptr) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	CGPoint result = [nSRulerMarker imageOrigin];
	return result;
}

void C_NSRulerMarker_SetImageOrigin(void* ptr, CGPoint value) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	[nSRulerMarker setImageOrigin:value];
}

CGRect C_NSRulerMarker_ImageRectInRuler(void* ptr) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	CGRect result = [nSRulerMarker imageRectInRuler];
	return result;
}

double C_NSRulerMarker_ThicknessRequiredInRuler(void* ptr) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	double result = [nSRulerMarker thicknessRequiredInRuler];
	return result;
}

bool C_NSRulerMarker_IsMovable(void* ptr) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	bool result = [nSRulerMarker isMovable];
	return result;
}

void C_NSRulerMarker_SetMovable(void* ptr, bool value) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	[nSRulerMarker setMovable:value];
}

bool C_NSRulerMarker_IsRemovable(void* ptr) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	bool result = [nSRulerMarker isRemovable];
	return result;
}

void C_NSRulerMarker_SetRemovable(void* ptr, bool value) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	[nSRulerMarker setRemovable:value];
}

double C_NSRulerMarker_MarkerLocation(void* ptr) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	double result = [nSRulerMarker markerLocation];
	return result;
}

void C_NSRulerMarker_SetMarkerLocation(void* ptr, double value) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	[nSRulerMarker setMarkerLocation:value];
}

bool C_NSRulerMarker_IsDragging(void* ptr) {
	NSRulerMarker* nSRulerMarker = (NSRulerMarker*)ptr;
	bool result = [nSRulerMarker isDragging];
	return result;
}