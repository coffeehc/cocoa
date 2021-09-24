#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_ParagraphStyle_Alloc();

void* C_NSParagraphStyle_AllocParagraphStyle();
void* C_NSParagraphStyle_Init(void* ptr);
void* C_NSParagraphStyle_NewParagraphStyle();
void* C_NSParagraphStyle_Autorelease(void* ptr);
void* C_NSParagraphStyle_Retain(void* ptr);
int C_NSParagraphStyle_ParagraphStyle_DefaultWritingDirectionForLanguage(void* languageName);
void* C_NSParagraphStyle_DefaultParagraphStyle();
int C_NSParagraphStyle_Alignment(void* ptr);
double C_NSParagraphStyle_FirstLineHeadIndent(void* ptr);
double C_NSParagraphStyle_HeadIndent(void* ptr);
double C_NSParagraphStyle_TailIndent(void* ptr);
double C_NSParagraphStyle_LineHeightMultiple(void* ptr);
double C_NSParagraphStyle_MaximumLineHeight(void* ptr);
double C_NSParagraphStyle_MinimumLineHeight(void* ptr);
double C_NSParagraphStyle_LineSpacing(void* ptr);
double C_NSParagraphStyle_ParagraphSpacing(void* ptr);
double C_NSParagraphStyle_ParagraphSpacingBefore(void* ptr);
Array C_NSParagraphStyle_TabStops(void* ptr);
double C_NSParagraphStyle_DefaultTabInterval(void* ptr);
Array C_NSParagraphStyle_TextBlocks(void* ptr);
Array C_NSParagraphStyle_TextLists(void* ptr);
unsigned int C_NSParagraphStyle_LineBreakMode(void* ptr);
unsigned int C_NSParagraphStyle_LineBreakStrategy(void* ptr);
float C_NSParagraphStyle_HyphenationFactor(void* ptr);
float C_NSParagraphStyle_TighteningFactorForTruncation(void* ptr);
bool C_NSParagraphStyle_AllowsDefaultTighteningForTruncation(void* ptr);
int C_NSParagraphStyle_HeaderLevel(void* ptr);
int C_NSParagraphStyle_BaseWritingDirection(void* ptr);

void* C_MutableParagraphStyle_Alloc();

void* C_NSMutableParagraphStyle_AllocMutableParagraphStyle();
void* C_NSMutableParagraphStyle_Init(void* ptr);
void* C_NSMutableParagraphStyle_NewMutableParagraphStyle();
void* C_NSMutableParagraphStyle_Autorelease(void* ptr);
void* C_NSMutableParagraphStyle_Retain(void* ptr);
void C_NSMutableParagraphStyle_SetParagraphStyle(void* ptr, void* obj);
void C_NSMutableParagraphStyle_AddTabStop(void* ptr, void* anObject);
void C_NSMutableParagraphStyle_RemoveTabStop(void* ptr, void* anObject);
void C_NSMutableParagraphStyle_SetAlignment(void* ptr, int value);
void C_NSMutableParagraphStyle_SetFirstLineHeadIndent(void* ptr, double value);
void C_NSMutableParagraphStyle_SetHeadIndent(void* ptr, double value);
void C_NSMutableParagraphStyle_SetTailIndent(void* ptr, double value);
void C_NSMutableParagraphStyle_SetLineHeightMultiple(void* ptr, double value);
void C_NSMutableParagraphStyle_SetMaximumLineHeight(void* ptr, double value);
void C_NSMutableParagraphStyle_SetMinimumLineHeight(void* ptr, double value);
void C_NSMutableParagraphStyle_SetLineSpacing(void* ptr, double value);
void C_NSMutableParagraphStyle_SetParagraphSpacing(void* ptr, double value);
void C_NSMutableParagraphStyle_SetParagraphSpacingBefore(void* ptr, double value);
void C_NSMutableParagraphStyle_SetBaseWritingDirection(void* ptr, int value);
void C_NSMutableParagraphStyle_SetTabStops(void* ptr, Array value);
void C_NSMutableParagraphStyle_SetDefaultTabInterval(void* ptr, double value);
void C_NSMutableParagraphStyle_SetTextBlocks(void* ptr, Array value);
void C_NSMutableParagraphStyle_SetTextLists(void* ptr, Array value);
void C_NSMutableParagraphStyle_SetLineBreakMode(void* ptr, unsigned int value);
void C_NSMutableParagraphStyle_SetLineBreakStrategy(void* ptr, unsigned int value);
void C_NSMutableParagraphStyle_SetHyphenationFactor(void* ptr, float value);
void C_NSMutableParagraphStyle_SetTighteningFactorForTruncation(void* ptr, float value);
void C_NSMutableParagraphStyle_SetAllowsDefaultTighteningForTruncation(void* ptr, bool value);
void C_NSMutableParagraphStyle_SetHeaderLevel(void* ptr, int value);

void* C_TextTab_Alloc();

void* C_NSTextTab_InitWithTextAlignment_Location_Options(void* ptr, int alignment, double loc, Dictionary options);
void* C_NSTextTab_AllocTextTab();
void* C_NSTextTab_Init(void* ptr);
void* C_NSTextTab_NewTextTab();
void* C_NSTextTab_Autorelease(void* ptr);
void* C_NSTextTab_Retain(void* ptr);
void* C_NSTextTab_TextTab_ColumnTerminatorsForLocale(void* aLocale);
double C_NSTextTab_Location(void* ptr);
int C_NSTextTab_Alignment(void* ptr);
Dictionary C_NSTextTab_Options(void* ptr);
