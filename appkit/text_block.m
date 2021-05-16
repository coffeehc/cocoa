#import <Appkit/Appkit.h>
#import "text_block.h"

void* C_TextBlock_Alloc() {
    return [NSTextBlock alloc];
}

void* C_NSTextBlock_Init(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlock* result = [nSTextBlock init];
    return result;
}

void C_NSTextBlock_SetValue_Type_ForDimension(void* ptr, double val, unsigned int _type, unsigned int dimension) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock setValue:val type:_type forDimension:dimension];
}

double C_NSTextBlock_ValueForDimension(void* ptr, unsigned int dimension) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    CGFloat result = [nSTextBlock valueForDimension:dimension];
    return result;
}

unsigned int C_NSTextBlock_ValueTypeForDimension(void* ptr, unsigned int dimension) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlockValueType result = [nSTextBlock valueTypeForDimension:dimension];
    return result;
}

void C_NSTextBlock_SetContentWidth_Type(void* ptr, double val, unsigned int _type) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock setContentWidth:val type:_type];
}

void C_NSTextBlock_SetWidth_Type_ForLayer_Edge(void* ptr, double val, unsigned int _type, int layer, unsigned int edge) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock setWidth:val type:_type forLayer:layer edge:edge];
}

void C_NSTextBlock_SetWidth_Type_ForLayer(void* ptr, double val, unsigned int _type, int layer) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock setWidth:val type:_type forLayer:layer];
}

double C_NSTextBlock_WidthForLayer_Edge(void* ptr, int layer, unsigned int edge) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    CGFloat result = [nSTextBlock widthForLayer:layer edge:edge];
    return result;
}

unsigned int C_NSTextBlock_WidthValueTypeForLayer_Edge(void* ptr, int layer, unsigned int edge) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlockValueType result = [nSTextBlock widthValueTypeForLayer:layer edge:edge];
    return result;
}

void C_NSTextBlock_SetBorderColor_ForEdge(void* ptr, void* color, unsigned int edge) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock setBorderColor:(NSColor*)color forEdge:edge];
}

void C_NSTextBlock_SetBorderColor(void* ptr, void* color) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock setBorderColor:(NSColor*)color];
}

void* C_NSTextBlock_BorderColorForEdge(void* ptr, unsigned int edge) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSColor* result = [nSTextBlock borderColorForEdge:edge];
    return result;
}

CGRect C_NSTextBlock_RectForLayoutAtPoint_InRect_TextContainer_CharacterRange(void* ptr, CGPoint startingPoint, CGRect rect, void* textContainer, NSRange charRange) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSRect result = [nSTextBlock rectForLayoutAtPoint:startingPoint inRect:rect textContainer:(NSTextContainer*)textContainer characterRange:charRange];
    return result;
}

CGRect C_NSTextBlock_BoundsRectForContentRect_InRect_TextContainer_CharacterRange(void* ptr, CGRect contentRect, CGRect rect, void* textContainer, NSRange charRange) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSRect result = [nSTextBlock boundsRectForContentRect:contentRect inRect:rect textContainer:(NSTextContainer*)textContainer characterRange:charRange];
    return result;
}

void C_NSTextBlock_DrawBackgroundWithFrame_InView_CharacterRange_LayoutManager(void* ptr, CGRect frameRect, void* controlView, NSRange charRange, void* layoutManager) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock drawBackgroundWithFrame:frameRect inView:(NSView*)controlView characterRange:charRange layoutManager:(NSLayoutManager*)layoutManager];
}

double C_NSTextBlock_ContentWidth(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    CGFloat result = [nSTextBlock contentWidth];
    return result;
}

unsigned int C_NSTextBlock_ContentWidthValueType(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlockValueType result = [nSTextBlock contentWidthValueType];
    return result;
}

unsigned int C_NSTextBlock_VerticalAlignment(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlockVerticalAlignment result = [nSTextBlock verticalAlignment];
    return result;
}

void C_NSTextBlock_SetVerticalAlignment(void* ptr, unsigned int value) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock setVerticalAlignment:value];
}

void* C_NSTextBlock_BackgroundColor(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSColor* result = [nSTextBlock backgroundColor];
    return result;
}

void C_NSTextBlock_SetBackgroundColor(void* ptr, void* value) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock setBackgroundColor:(NSColor*)value];
}
