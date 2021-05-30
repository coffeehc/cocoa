#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_MutableParagraphStyle_Alloc();

void* C_NSMutableParagraphStyle_Init(void* ptr);
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
