#import <Foundation/NSGeometry.h>
#import <Foundation/NSRange.h>
#import <stdbool.h>
#import <stdint.h>
#import <stdlib.h>
#import <utils.h>

void* C_TextContainer_Alloc();

void* C_NSTextContainer_InitWithSize(void* ptr, CGSize size);
void* C_NSTextContainer_InitWithCoder(void* ptr, void* coder);
void* C_NSTextContainer_AllocTextContainer();
void* C_NSTextContainer_Init(void* ptr);
void* C_NSTextContainer_NewTextContainer();
void* C_NSTextContainer_Autorelease(void* ptr);
void* C_NSTextContainer_Retain(void* ptr);
void C_NSTextContainer_ReplaceLayoutManager(void* ptr, void* newLayoutManager);
void* C_NSTextContainer_LayoutManager(void* ptr);
void C_NSTextContainer_SetLayoutManager(void* ptr, void* value);
void* C_NSTextContainer_TextView(void* ptr);
void C_NSTextContainer_SetTextView(void* ptr, void* value);
CGSize C_NSTextContainer_Size(void* ptr);
void C_NSTextContainer_SetSize(void* ptr, CGSize value);
Array C_NSTextContainer_ExclusionPaths(void* ptr);
void C_NSTextContainer_SetExclusionPaths(void* ptr, Array value);
unsigned int C_NSTextContainer_LineBreakMode(void* ptr);
void C_NSTextContainer_SetLineBreakMode(void* ptr, unsigned int value);
bool C_NSTextContainer_WidthTracksTextView(void* ptr);
void C_NSTextContainer_SetWidthTracksTextView(void* ptr, bool value);
bool C_NSTextContainer_HeightTracksTextView(void* ptr);
void C_NSTextContainer_SetHeightTracksTextView(void* ptr, bool value);
unsigned int C_NSTextContainer_MaximumNumberOfLines(void* ptr);
void C_NSTextContainer_SetMaximumNumberOfLines(void* ptr, unsigned int value);
double C_NSTextContainer_LineFragmentPadding(void* ptr);
void C_NSTextContainer_SetLineFragmentPadding(void* ptr, double value);
bool C_NSTextContainer_IsSimpleRectangularTextContainer(void* ptr);
