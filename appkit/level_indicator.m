#import "level_indicator.h"
#import <Foundation/NSDictionary.h>
#import <Foundation/NSArray.h>
#import <AppKit/NSLevelIndicator.h>

void* C_LevelIndicator_Alloc() {
    return [NSLevelIndicator alloc];
}

void* C_NSLevelIndicator_InitWithFrame(void* ptr, CGRect frameRect) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSLevelIndicator* result_ = [nSLevelIndicator initWithFrame:frameRect];
    return result_;
}

void* C_NSLevelIndicator_InitWithCoder(void* ptr, void* coder) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSLevelIndicator* result_ = [nSLevelIndicator initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSLevelIndicator_Init(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSLevelIndicator* result_ = [nSLevelIndicator init];
    return result_;
}

void* C_NSLevelIndicator_AllocLevelIndicator() {
    NSLevelIndicator* result_ = [NSLevelIndicator alloc];
    return result_;
}

void* C_NSLevelIndicator_NewLevelIndicator() {
    NSLevelIndicator* result_ = [NSLevelIndicator new];
    return result_;
}

void* C_NSLevelIndicator_Autorelease(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSLevelIndicator* result_ = [nSLevelIndicator autorelease];
    return result_;
}

void* C_NSLevelIndicator_Retain(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSLevelIndicator* result_ = [nSLevelIndicator retain];
    return result_;
}

double C_NSLevelIndicator_TickMarkValueAtIndex(void* ptr, int index) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    double result_ = [nSLevelIndicator tickMarkValueAtIndex:index];
    return result_;
}

CGRect C_NSLevelIndicator_RectOfTickMarkAtIndex(void* ptr, int index) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSRect result_ = [nSLevelIndicator rectOfTickMarkAtIndex:index];
    return result_;
}

double C_NSLevelIndicator_MinValue(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    double result_ = [nSLevelIndicator minValue];
    return result_;
}

void C_NSLevelIndicator_SetMinValue(void* ptr, double value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setMinValue:value];
}

double C_NSLevelIndicator_MaxValue(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    double result_ = [nSLevelIndicator maxValue];
    return result_;
}

void C_NSLevelIndicator_SetMaxValue(void* ptr, double value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setMaxValue:value];
}

double C_NSLevelIndicator_WarningValue(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    double result_ = [nSLevelIndicator warningValue];
    return result_;
}

void C_NSLevelIndicator_SetWarningValue(void* ptr, double value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setWarningValue:value];
}

double C_NSLevelIndicator_CriticalValue(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    double result_ = [nSLevelIndicator criticalValue];
    return result_;
}

void C_NSLevelIndicator_SetCriticalValue(void* ptr, double value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setCriticalValue:value];
}

unsigned int C_NSLevelIndicator_TickMarkPosition(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSTickMarkPosition result_ = [nSLevelIndicator tickMarkPosition];
    return result_;
}

void C_NSLevelIndicator_SetTickMarkPosition(void* ptr, unsigned int value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setTickMarkPosition:value];
}

int C_NSLevelIndicator_NumberOfTickMarks(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSInteger result_ = [nSLevelIndicator numberOfTickMarks];
    return result_;
}

void C_NSLevelIndicator_SetNumberOfTickMarks(void* ptr, int value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setNumberOfTickMarks:value];
}

int C_NSLevelIndicator_NumberOfMajorTickMarks(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSInteger result_ = [nSLevelIndicator numberOfMajorTickMarks];
    return result_;
}

void C_NSLevelIndicator_SetNumberOfMajorTickMarks(void* ptr, int value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setNumberOfMajorTickMarks:value];
}

unsigned int C_NSLevelIndicator_LevelIndicatorStyle(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSLevelIndicatorStyle result_ = [nSLevelIndicator levelIndicatorStyle];
    return result_;
}

void C_NSLevelIndicator_SetLevelIndicatorStyle(void* ptr, unsigned int value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setLevelIndicatorStyle:value];
}

void* C_NSLevelIndicator_RatingImage(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSImage* result_ = [nSLevelIndicator ratingImage];
    return result_;
}

void C_NSLevelIndicator_SetRatingImage(void* ptr, void* value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setRatingImage:(NSImage*)value];
}

bool C_NSLevelIndicator_DrawsTieredCapacityLevels(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    BOOL result_ = [nSLevelIndicator drawsTieredCapacityLevels];
    return result_;
}

void C_NSLevelIndicator_SetDrawsTieredCapacityLevels(void* ptr, bool value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setDrawsTieredCapacityLevels:value];
}

void* C_NSLevelIndicator_FillColor(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSColor* result_ = [nSLevelIndicator fillColor];
    return result_;
}

void C_NSLevelIndicator_SetFillColor(void* ptr, void* value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setFillColor:(NSColor*)value];
}

void* C_NSLevelIndicator_WarningFillColor(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSColor* result_ = [nSLevelIndicator warningFillColor];
    return result_;
}

void C_NSLevelIndicator_SetWarningFillColor(void* ptr, void* value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setWarningFillColor:(NSColor*)value];
}

void* C_NSLevelIndicator_CriticalFillColor(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSColor* result_ = [nSLevelIndicator criticalFillColor];
    return result_;
}

void C_NSLevelIndicator_SetCriticalFillColor(void* ptr, void* value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setCriticalFillColor:(NSColor*)value];
}

void* C_NSLevelIndicator_RatingPlaceholderImage(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSImage* result_ = [nSLevelIndicator ratingPlaceholderImage];
    return result_;
}

void C_NSLevelIndicator_SetRatingPlaceholderImage(void* ptr, void* value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setRatingPlaceholderImage:(NSImage*)value];
}

int C_NSLevelIndicator_PlaceholderVisibility(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    NSLevelIndicatorPlaceholderVisibility result_ = [nSLevelIndicator placeholderVisibility];
    return result_;
}

void C_NSLevelIndicator_SetPlaceholderVisibility(void* ptr, int value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setPlaceholderVisibility:value];
}

bool C_NSLevelIndicator_IsEditable(void* ptr) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    BOOL result_ = [nSLevelIndicator isEditable];
    return result_;
}

void C_NSLevelIndicator_SetEditable(void* ptr, bool value) {
    NSLevelIndicator* nSLevelIndicator = (NSLevelIndicator*)ptr;
    [nSLevelIndicator setEditable:value];
}
