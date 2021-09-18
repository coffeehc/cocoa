#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_RulerView_Alloc();

void* C_NSRulerView_InitWithScrollView_Orientation(void* ptr, void* scrollView, unsigned int orientation);
void* C_NSRulerView_InitWithCoder(void* ptr, void* coder);
void* C_NSRulerView_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSRulerView_Init(void* ptr);
void* C_NSRulerView_AllocRulerView();
void* C_NSRulerView_NewRulerView();
void* C_NSRulerView_Autorelease(void* ptr);
void* C_NSRulerView_Retain(void* ptr);
void C_NSRulerView_RulerView_RegisterUnitWithName_Abbreviation_UnitToPointsConversionFactor_StepUpCycle_StepDownCycle(void* unitName, void* abbreviation, double conversionFactor, Array stepUpCycle, Array stepDownCycle);
void C_NSRulerView_AddMarker(void* ptr, void* marker);
void C_NSRulerView_RemoveMarker(void* ptr, void* marker);
bool C_NSRulerView_TrackMarker_WithMouseEvent(void* ptr, void* marker, void* event);
void C_NSRulerView_MoveRulerlineFromLocation_ToLocation(void* ptr, double oldLocation, double newLocation);
void C_NSRulerView_DrawHashMarksAndLabelsInRect(void* ptr, CGRect rect);
void C_NSRulerView_DrawMarkersInRect(void* ptr, CGRect rect);
void C_NSRulerView_InvalidateHashMarks(void* ptr);
void* C_NSRulerView_MeasurementUnits(void* ptr);
void C_NSRulerView_SetMeasurementUnits(void* ptr, void* value);
void* C_NSRulerView_ClientView(void* ptr);
void C_NSRulerView_SetClientView(void* ptr, void* value);
void* C_NSRulerView_AccessoryView(void* ptr);
void C_NSRulerView_SetAccessoryView(void* ptr, void* value);
double C_NSRulerView_OriginOffset(void* ptr);
void C_NSRulerView_SetOriginOffset(void* ptr, double value);
Array C_NSRulerView_Markers(void* ptr);
void C_NSRulerView_SetMarkers(void* ptr, Array value);
void* C_NSRulerView_ScrollView(void* ptr);
void C_NSRulerView_SetScrollView(void* ptr, void* value);
unsigned int C_NSRulerView_Orientation(void* ptr);
void C_NSRulerView_SetOrientation(void* ptr, unsigned int value);
double C_NSRulerView_ReservedThicknessForAccessoryView(void* ptr);
void C_NSRulerView_SetReservedThicknessForAccessoryView(void* ptr, double value);
double C_NSRulerView_ReservedThicknessForMarkers(void* ptr);
void C_NSRulerView_SetReservedThicknessForMarkers(void* ptr, double value);
double C_NSRulerView_RuleThickness(void* ptr);
void C_NSRulerView_SetRuleThickness(void* ptr, double value);
double C_NSRulerView_RequiredThickness(void* ptr);
double C_NSRulerView_BaselineLocation(void* ptr);
