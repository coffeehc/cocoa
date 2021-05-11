#import <Appkit/Appkit.h>
#import "ruler_view.h"

void* C_RulerView_Alloc() {
	return [NSRulerView alloc];
}

void* C_NSRulerView_InitWithCoder(void* ptr, void* coder) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	NSRulerView* result = [nSRulerView initWithCoder:(NSCoder*)coder];
	return result;
}

void* C_NSRulerView_InitWithFrame(void* ptr, CGRect frameRect) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	NSRulerView* result = [nSRulerView initWithFrame:frameRect];
	return result;
}

void* C_NSRulerView_Init(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	NSRulerView* result = [nSRulerView init];
	return result;
}

void C_NSRulerView_AddMarker(void* ptr, void* marker) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView addMarker:(NSRulerMarker*)marker];
}

void C_NSRulerView_RemoveMarker(void* ptr, void* marker) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView removeMarker:(NSRulerMarker*)marker];
}

bool C_NSRulerView_TrackMarker_WithMouseEvent(void* ptr, void* marker, void* event) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	bool result = [nSRulerView trackMarker:(NSRulerMarker*)marker withMouseEvent:(NSEvent*)event];
	return result;
}

void C_NSRulerView_MoveRulerlineFromLocation_ToLocation(void* ptr, double oldLocation, double newLocation) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView moveRulerlineFromLocation:oldLocation toLocation:newLocation];
}

void C_NSRulerView_DrawHashMarksAndLabelsInRect(void* ptr, CGRect rect) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView drawHashMarksAndLabelsInRect:rect];
}

void C_NSRulerView_DrawMarkersInRect(void* ptr, CGRect rect) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView drawMarkersInRect:rect];
}

void C_NSRulerView_InvalidateHashMarks(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView invalidateHashMarks];
}

void* C_NSRulerView_MeasurementUnits(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	NSString* result = [nSRulerView measurementUnits];
	return result;
}

void C_NSRulerView_SetMeasurementUnits(void* ptr, void* value) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView setMeasurementUnits:(NSString*)value];
}

void* C_NSRulerView_ClientView(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	NSView* result = [nSRulerView clientView];
	return result;
}

void C_NSRulerView_SetClientView(void* ptr, void* value) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView setClientView:(NSView*)value];
}

void* C_NSRulerView_AccessoryView(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	NSView* result = [nSRulerView accessoryView];
	return result;
}

void C_NSRulerView_SetAccessoryView(void* ptr, void* value) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView setAccessoryView:(NSView*)value];
}

double C_NSRulerView_OriginOffset(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	double result = [nSRulerView originOffset];
	return result;
}

void C_NSRulerView_SetOriginOffset(void* ptr, double value) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView setOriginOffset:value];
}

void* C_NSRulerView_ScrollView(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	NSScrollView* result = [nSRulerView scrollView];
	return result;
}

void C_NSRulerView_SetScrollView(void* ptr, void* value) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView setScrollView:(NSScrollView*)value];
}

double C_NSRulerView_ReservedThicknessForAccessoryView(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	double result = [nSRulerView reservedThicknessForAccessoryView];
	return result;
}

void C_NSRulerView_SetReservedThicknessForAccessoryView(void* ptr, double value) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView setReservedThicknessForAccessoryView:value];
}

double C_NSRulerView_ReservedThicknessForMarkers(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	double result = [nSRulerView reservedThicknessForMarkers];
	return result;
}

void C_NSRulerView_SetReservedThicknessForMarkers(void* ptr, double value) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView setReservedThicknessForMarkers:value];
}

double C_NSRulerView_RuleThickness(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	double result = [nSRulerView ruleThickness];
	return result;
}

void C_NSRulerView_SetRuleThickness(void* ptr, double value) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	[nSRulerView setRuleThickness:value];
}

double C_NSRulerView_RequiredThickness(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	double result = [nSRulerView requiredThickness];
	return result;
}

double C_NSRulerView_BaselineLocation(void* ptr) {
	NSRulerView* nSRulerView = (NSRulerView*)ptr;
	double result = [nSRulerView baselineLocation];
	return result;
}
