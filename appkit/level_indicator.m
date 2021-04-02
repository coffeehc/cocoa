#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "level_indicator.h"

double LevelIndicator_MinValue(void* ptr) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	return [levelIndicator minValue];
}

void LevelIndicator_SetMinValue(void* ptr, double minValue) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	[levelIndicator setMinValue:minValue];
}

double LevelIndicator_MaxValue(void* ptr) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	return [levelIndicator maxValue];
}

void LevelIndicator_SetMaxValue(void* ptr, double maxValue) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	[levelIndicator setMaxValue:maxValue];
}

double LevelIndicator_WarningValue(void* ptr) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	return [levelIndicator warningValue];
}

void LevelIndicator_SetWarningValue(void* ptr, double warningValue) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	[levelIndicator setWarningValue:warningValue];
}

double LevelIndicator_CriticalValue(void* ptr) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	return [levelIndicator criticalValue];
}

void LevelIndicator_SetCriticalValue(void* ptr, double criticalValue) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	[levelIndicator setCriticalValue:criticalValue];
}

bool LevelIndicator_IsEditable(void* ptr) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	return [levelIndicator isEditable];
}

void LevelIndicator_SetEditable(void* ptr, bool editable) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	[levelIndicator setEditable:editable];
}

unsigned long LevelIndicator_TickMarkPosition(void* ptr) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	return [levelIndicator tickMarkPosition];
}

void LevelIndicator_SetTickMarkPosition(void* ptr, unsigned long tickMarkPosition) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	[levelIndicator setTickMarkPosition:tickMarkPosition];
}

long LevelIndicator_NumberOfTickMarks(void* ptr) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	return [levelIndicator numberOfTickMarks];
}

void LevelIndicator_SetNumberOfTickMarks(void* ptr, long numberOfTickMarks) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	[levelIndicator setNumberOfTickMarks:numberOfTickMarks];
}

long LevelIndicator_NumberOfMajorTickMarks(void* ptr) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	return [levelIndicator numberOfMajorTickMarks];
}

void LevelIndicator_SetNumberOfMajorTickMarks(void* ptr, long numberOfMajorTickMarks) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	[levelIndicator setNumberOfMajorTickMarks:numberOfMajorTickMarks];
}

unsigned long LevelIndicator_LevelIndicatorStyle(void* ptr) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	return [levelIndicator levelIndicatorStyle];
}

void LevelIndicator_SetLevelIndicatorStyle(void* ptr, unsigned long levelIndicatorStyle) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	[levelIndicator setLevelIndicatorStyle:levelIndicatorStyle];
}

void* LevelIndicator_NewLevelIndicator(NSRect frame) {
	NSLevelIndicator* levelIndicator = [NSLevelIndicator alloc];
	return [[levelIndicator initWithFrame:frame] autorelease];
}

double LevelIndicator_TickMarkValueAtIndex(void* ptr, long index) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	return [levelIndicator tickMarkValueAtIndex:index];
}

NSRect LevelIndicator_RectOfTickMarkAtIndex(void* ptr, long index) {
	NSLevelIndicator* levelIndicator = (NSLevelIndicator*)ptr;
	return [levelIndicator rectOfTickMarkAtIndex:index];
}
