#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

void* C_TextCheckingResult_Alloc();

void* C_NSTextCheckingResult_AllocTextCheckingResult();
void* C_NSTextCheckingResult_Init(void* ptr);
void* C_NSTextCheckingResult_NewTextCheckingResult();
void* C_NSTextCheckingResult_Autorelease(void* ptr);
void* C_NSTextCheckingResult_Retain(void* ptr);
NSRange C_NSTextCheckingResult_RangeAtIndex(void* ptr, unsigned int idx);
void* C_NSTextCheckingResult_TextCheckingResult_ReplacementCheckingResultWithRange_ReplacementString(NSRange _range, void* replacementString);
void* C_NSTextCheckingResult_TextCheckingResult_LinkCheckingResultWithRange_URL(NSRange _range, void* url);
void* C_NSTextCheckingResult_TextCheckingResult_AddressCheckingResultWithRange_Components(NSRange _range, Dictionary components);
void* C_NSTextCheckingResult_TextCheckingResult_TransitInformationCheckingResultWithRange_Components(NSRange _range, Dictionary components);
void* C_NSTextCheckingResult_TextCheckingResult_PhoneNumberCheckingResultWithRange_PhoneNumber(NSRange _range, void* phoneNumber);
void* C_NSTextCheckingResult_TextCheckingResult_DateCheckingResultWithRange_Date(NSRange _range, void* date);
void* C_NSTextCheckingResult_TextCheckingResult_DateCheckingResultWithRange_Date_TimeZone_Duration(NSRange _range, void* date, void* timeZone, double duration);
void* C_NSTextCheckingResult_TextCheckingResult_DashCheckingResultWithRange_ReplacementString(NSRange _range, void* replacementString);
void* C_NSTextCheckingResult_TextCheckingResult_QuoteCheckingResultWithRange_ReplacementString(NSRange _range, void* replacementString);
void* C_NSTextCheckingResult_TextCheckingResult_SpellCheckingResultWithRange(NSRange _range);
void* C_NSTextCheckingResult_TextCheckingResult_CorrectionCheckingResultWithRange_ReplacementString(NSRange _range, void* replacementString);
void* C_NSTextCheckingResult_TextCheckingResult_OrthographyCheckingResultWithRange_Orthography(NSRange _range, void* orthography);
void* C_NSTextCheckingResult_ResultByAdjustingRangesWithOffset(void* ptr, int offset);
NSRange C_NSTextCheckingResult_RangeWithName(void* ptr, void* name);
void* C_NSTextCheckingResult_TextCheckingResult_CorrectionCheckingResultWithRange_ReplacementString_AlternativeStrings(NSRange _range, void* replacementString, Array alternativeStrings);
NSRange C_NSTextCheckingResult_Range(void* ptr);
unsigned int C_NSTextCheckingResult_NumberOfRanges(void* ptr);
void* C_NSTextCheckingResult_ReplacementString(void* ptr);
void* C_NSTextCheckingResult_RegularExpression(void* ptr);
Dictionary C_NSTextCheckingResult_Components(void* ptr);
void* C_NSTextCheckingResult_URL(void* ptr);
Dictionary C_NSTextCheckingResult_AddressComponents(void* ptr);
void* C_NSTextCheckingResult_PhoneNumber(void* ptr);
void* C_NSTextCheckingResult_Date(void* ptr);
double C_NSTextCheckingResult_Duration(void* ptr);
void* C_NSTextCheckingResult_TimeZone(void* ptr);
void* C_NSTextCheckingResult_Orthography(void* ptr);
Array C_NSTextCheckingResult_AlternativeStrings(void* ptr);
