#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_LevelIndicator_Alloc();

void* C_NSLevelIndicator_InitWithFrame(void* ptr, CGRect frameRect);
void* C_NSLevelIndicator_InitWithCoder(void* ptr, void* coder);
void* C_NSLevelIndicator_Init(void* ptr);
void* C_NSLevelIndicator_AllocLevelIndicator();
void* C_NSLevelIndicator_NewLevelIndicator();
void* C_NSLevelIndicator_Autorelease(void* ptr);
void* C_NSLevelIndicator_Retain(void* ptr);
double C_NSLevelIndicator_TickMarkValueAtIndex(void* ptr, int index);
CGRect C_NSLevelIndicator_RectOfTickMarkAtIndex(void* ptr, int index);
double C_NSLevelIndicator_MinValue(void* ptr);
void C_NSLevelIndicator_SetMinValue(void* ptr, double value);
double C_NSLevelIndicator_MaxValue(void* ptr);
void C_NSLevelIndicator_SetMaxValue(void* ptr, double value);
double C_NSLevelIndicator_WarningValue(void* ptr);
void C_NSLevelIndicator_SetWarningValue(void* ptr, double value);
double C_NSLevelIndicator_CriticalValue(void* ptr);
void C_NSLevelIndicator_SetCriticalValue(void* ptr, double value);
unsigned int C_NSLevelIndicator_TickMarkPosition(void* ptr);
void C_NSLevelIndicator_SetTickMarkPosition(void* ptr, unsigned int value);
int C_NSLevelIndicator_NumberOfTickMarks(void* ptr);
void C_NSLevelIndicator_SetNumberOfTickMarks(void* ptr, int value);
int C_NSLevelIndicator_NumberOfMajorTickMarks(void* ptr);
void C_NSLevelIndicator_SetNumberOfMajorTickMarks(void* ptr, int value);
unsigned int C_NSLevelIndicator_LevelIndicatorStyle(void* ptr);
void C_NSLevelIndicator_SetLevelIndicatorStyle(void* ptr, unsigned int value);
void* C_NSLevelIndicator_RatingImage(void* ptr);
void C_NSLevelIndicator_SetRatingImage(void* ptr, void* value);
bool C_NSLevelIndicator_DrawsTieredCapacityLevels(void* ptr);
void C_NSLevelIndicator_SetDrawsTieredCapacityLevels(void* ptr, bool value);
void* C_NSLevelIndicator_FillColor(void* ptr);
void C_NSLevelIndicator_SetFillColor(void* ptr, void* value);
void* C_NSLevelIndicator_WarningFillColor(void* ptr);
void C_NSLevelIndicator_SetWarningFillColor(void* ptr, void* value);
void* C_NSLevelIndicator_CriticalFillColor(void* ptr);
void C_NSLevelIndicator_SetCriticalFillColor(void* ptr, void* value);
void* C_NSLevelIndicator_RatingPlaceholderImage(void* ptr);
void C_NSLevelIndicator_SetRatingPlaceholderImage(void* ptr, void* value);
int C_NSLevelIndicator_PlaceholderVisibility(void* ptr);
void C_NSLevelIndicator_SetPlaceholderVisibility(void* ptr, int value);
bool C_NSLevelIndicator_IsEditable(void* ptr);
void C_NSLevelIndicator_SetEditable(void* ptr, bool value);
