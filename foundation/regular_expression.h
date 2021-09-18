#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_RegularExpression_Alloc();

void* C_NSRegularExpression_AllocRegularExpression();
void* C_NSRegularExpression_Init(void* ptr);
void* C_NSRegularExpression_NewRegularExpression();
void* C_NSRegularExpression_Autorelease(void* ptr);
void* C_NSRegularExpression_Retain(void* ptr);
unsigned int C_NSRegularExpression_NumberOfMatchesInString_Options_Range(void* ptr, void* _string, unsigned int options, NSRange _range);
Array C_NSRegularExpression_MatchesInString_Options_Range(void* ptr, void* _string, unsigned int options, NSRange _range);
void* C_NSRegularExpression_FirstMatchInString_Options_Range(void* ptr, void* _string, unsigned int options, NSRange _range);
NSRange C_NSRegularExpression_RangeOfFirstMatchInString_Options_Range(void* ptr, void* _string, unsigned int options, NSRange _range);
unsigned int C_NSRegularExpression_ReplaceMatchesInString_Options_Range_WithTemplate(void* ptr, void* _string, unsigned int options, NSRange _range, void* templ);
void* C_NSRegularExpression_StringByReplacingMatchesInString_Options_Range_WithTemplate(void* ptr, void* _string, unsigned int options, NSRange _range, void* templ);
void* C_NSRegularExpression_RegularExpression_EscapedTemplateForString(void* _string);
void* C_NSRegularExpression_RegularExpression_EscapedPatternForString(void* _string);
void* C_NSRegularExpression_ReplacementStringForResult_InString_Offset_Template(void* ptr, void* result, void* _string, int offset, void* templ);
void* C_NSRegularExpression_Pattern(void* ptr);
unsigned int C_NSRegularExpression_Options(void* ptr);
unsigned int C_NSRegularExpression_NumberOfCaptureGroups(void* ptr);
