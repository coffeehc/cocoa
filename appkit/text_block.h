#import <stdbool.h>
#import <stdlib.h>
#import <utils.h>
#import <Appkit/Appkit.h>

void* C_TextBlock_Alloc();

void* C_NSTextBlock_Init(void* ptr);
void* C_NSTextBlock_AllocTextBlock();
void* C_NSTextBlock_NewTextBlock();
void* C_NSTextBlock_Autorelease(void* ptr);
void* C_NSTextBlock_Retain(void* ptr);
void C_NSTextBlock_SetValue_Type_ForDimension(void* ptr, double val, unsigned int _type, unsigned int dimension);
double C_NSTextBlock_ValueForDimension(void* ptr, unsigned int dimension);
unsigned int C_NSTextBlock_ValueTypeForDimension(void* ptr, unsigned int dimension);
void C_NSTextBlock_SetContentWidth_Type(void* ptr, double val, unsigned int _type);
void C_NSTextBlock_SetWidth_Type_ForLayer_Edge(void* ptr, double val, unsigned int _type, int layer, unsigned int edge);
void C_NSTextBlock_SetWidth_Type_ForLayer(void* ptr, double val, unsigned int _type, int layer);
double C_NSTextBlock_WidthForLayer_Edge(void* ptr, int layer, unsigned int edge);
unsigned int C_NSTextBlock_WidthValueTypeForLayer_Edge(void* ptr, int layer, unsigned int edge);
void C_NSTextBlock_SetBorderColor_ForEdge(void* ptr, void* color, unsigned int edge);
void C_NSTextBlock_SetBorderColor(void* ptr, void* color);
void* C_NSTextBlock_BorderColorForEdge(void* ptr, unsigned int edge);
CGRect C_NSTextBlock_RectForLayoutAtPoint_InRect_TextContainer_CharacterRange(void* ptr, CGPoint startingPoint, CGRect rect, void* textContainer, NSRange charRange);
CGRect C_NSTextBlock_BoundsRectForContentRect_InRect_TextContainer_CharacterRange(void* ptr, CGRect contentRect, CGRect rect, void* textContainer, NSRange charRange);
void C_NSTextBlock_DrawBackgroundWithFrame_InView_CharacterRange_LayoutManager(void* ptr, CGRect frameRect, void* controlView, NSRange charRange, void* layoutManager);
double C_NSTextBlock_ContentWidth(void* ptr);
unsigned int C_NSTextBlock_ContentWidthValueType(void* ptr);
unsigned int C_NSTextBlock_VerticalAlignment(void* ptr);
void C_NSTextBlock_SetVerticalAlignment(void* ptr, unsigned int value);
void* C_NSTextBlock_BackgroundColor(void* ptr);
void C_NSTextBlock_SetBackgroundColor(void* ptr, void* value);
