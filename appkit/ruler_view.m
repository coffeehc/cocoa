#import <Appkit/Appkit.h>
#import "ruler_view.h"

void* C_RulerView_Alloc() {
    return [NSRulerView alloc];
}

void* C_NSRulerView_InitWithScrollView_Orientation(void* ptr, void* scrollView, unsigned int orientation) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    NSRulerView* result_ = [nSRulerView initWithScrollView:(NSScrollView*)scrollView orientation:orientation];
    return result_;
}

void* C_NSRulerView_InitWithCoder(void* ptr, void* coder) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    NSRulerView* result_ = [nSRulerView initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSRulerView_InitWithFrame(void* ptr, CGRect frameRect) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    NSRulerView* result_ = [nSRulerView initWithFrame:frameRect];
    return result_;
}

void* C_NSRulerView_Init(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    NSRulerView* result_ = [nSRulerView init];
    return result_;
}

void C_NSRulerView_RulerView_RegisterUnitWithName_Abbreviation_UnitToPointsConversionFactor_StepUpCycle_StepDownCycle(void* unitName, void* abbreviation, double conversionFactor, Array stepUpCycle, Array stepDownCycle) {
    NSMutableArray* objcStepUpCycle = [[NSMutableArray alloc] init];
    if (stepUpCycle.len > 0) {
    	void** stepUpCycleData = (void**)stepUpCycle.data;
    	for (int i = 0; i < stepUpCycle.len; i++) {
    		void* p = stepUpCycleData[i];
    		[objcStepUpCycle addObject:(NSNumber*)(NSNumber*)p];
    	}
    }
    NSMutableArray* objcStepDownCycle = [[NSMutableArray alloc] init];
    if (stepDownCycle.len > 0) {
    	void** stepDownCycleData = (void**)stepDownCycle.data;
    	for (int i = 0; i < stepDownCycle.len; i++) {
    		void* p = stepDownCycleData[i];
    		[objcStepDownCycle addObject:(NSNumber*)(NSNumber*)p];
    	}
    }
    [NSRulerView registerUnitWithName:(NSString*)unitName abbreviation:(NSString*)abbreviation unitToPointsConversionFactor:conversionFactor stepUpCycle:objcStepUpCycle stepDownCycle:objcStepDownCycle];
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
    BOOL result_ = [nSRulerView trackMarker:(NSRulerMarker*)marker withMouseEvent:(NSEvent*)event];
    return result_;
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
    NSRulerViewUnitName result_ = [nSRulerView measurementUnits];
    return result_;
}

void C_NSRulerView_SetMeasurementUnits(void* ptr, void* value) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    [nSRulerView setMeasurementUnits:(NSString*)value];
}

void* C_NSRulerView_ClientView(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    NSView* result_ = [nSRulerView clientView];
    return result_;
}

void C_NSRulerView_SetClientView(void* ptr, void* value) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    [nSRulerView setClientView:(NSView*)value];
}

void* C_NSRulerView_AccessoryView(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    NSView* result_ = [nSRulerView accessoryView];
    return result_;
}

void C_NSRulerView_SetAccessoryView(void* ptr, void* value) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    [nSRulerView setAccessoryView:(NSView*)value];
}

double C_NSRulerView_OriginOffset(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    CGFloat result_ = [nSRulerView originOffset];
    return result_;
}

void C_NSRulerView_SetOriginOffset(void* ptr, double value) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    [nSRulerView setOriginOffset:value];
}

Array C_NSRulerView_Markers(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    NSArray* result_ = [nSRulerView markers];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSRulerView_SetMarkers(void* ptr, Array value) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSRulerMarker*)(NSRulerMarker*)p];
    	}
    }
    [nSRulerView setMarkers:objcValue];
}

void* C_NSRulerView_ScrollView(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    NSScrollView* result_ = [nSRulerView scrollView];
    return result_;
}

void C_NSRulerView_SetScrollView(void* ptr, void* value) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    [nSRulerView setScrollView:(NSScrollView*)value];
}

unsigned int C_NSRulerView_Orientation(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    NSRulerOrientation result_ = [nSRulerView orientation];
    return result_;
}

void C_NSRulerView_SetOrientation(void* ptr, unsigned int value) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    [nSRulerView setOrientation:value];
}

double C_NSRulerView_ReservedThicknessForAccessoryView(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    CGFloat result_ = [nSRulerView reservedThicknessForAccessoryView];
    return result_;
}

void C_NSRulerView_SetReservedThicknessForAccessoryView(void* ptr, double value) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    [nSRulerView setReservedThicknessForAccessoryView:value];
}

double C_NSRulerView_ReservedThicknessForMarkers(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    CGFloat result_ = [nSRulerView reservedThicknessForMarkers];
    return result_;
}

void C_NSRulerView_SetReservedThicknessForMarkers(void* ptr, double value) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    [nSRulerView setReservedThicknessForMarkers:value];
}

double C_NSRulerView_RuleThickness(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    CGFloat result_ = [nSRulerView ruleThickness];
    return result_;
}

void C_NSRulerView_SetRuleThickness(void* ptr, double value) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    [nSRulerView setRuleThickness:value];
}

double C_NSRulerView_RequiredThickness(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    CGFloat result_ = [nSRulerView requiredThickness];
    return result_;
}

double C_NSRulerView_BaselineLocation(void* ptr) {
    NSRulerView* nSRulerView = (NSRulerView*)ptr;
    CGFloat result_ = [nSRulerView baselineLocation];
    return result_;
}
