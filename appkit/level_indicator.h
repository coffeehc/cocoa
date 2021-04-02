#import <Foundation/NSGeometry.h>
#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>

double LevelIndicator_MinValue(void* ptr);
void LevelIndicator_SetMinValue(void* ptr, double minValue);
double LevelIndicator_MaxValue(void* ptr);
void LevelIndicator_SetMaxValue(void* ptr, double maxValue);
double LevelIndicator_WarningValue(void* ptr);
void LevelIndicator_SetWarningValue(void* ptr, double warningValue);
double LevelIndicator_CriticalValue(void* ptr);
void LevelIndicator_SetCriticalValue(void* ptr, double criticalValue);
bool LevelIndicator_IsEditable(void* ptr);
void LevelIndicator_SetEditable(void* ptr, bool editable);
unsigned long LevelIndicator_TickMarkPosition(void* ptr);
void LevelIndicator_SetTickMarkPosition(void* ptr, unsigned long tickMarkPosition);
long LevelIndicator_NumberOfTickMarks(void* ptr);
void LevelIndicator_SetNumberOfTickMarks(void* ptr, long numberOfTickMarks);
long LevelIndicator_NumberOfMajorTickMarks(void* ptr);
void LevelIndicator_SetNumberOfMajorTickMarks(void* ptr, long numberOfMajorTickMarks);
unsigned long LevelIndicator_LevelIndicatorStyle(void* ptr);
void LevelIndicator_SetLevelIndicatorStyle(void* ptr, unsigned long levelIndicatorStyle);

void* LevelIndicator_NewLevelIndicator(NSRect frame);
double LevelIndicator_TickMarkValueAtIndex(void* ptr, long index);
NSRect LevelIndicator_RectOfTickMarkAtIndex(void* ptr, long index);
