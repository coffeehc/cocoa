#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>

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
