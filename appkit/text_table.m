#import "text_table.h"
#import <AppKit/NSTextTable.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>

void* C_TextBlock_Alloc() {
    return [NSTextBlock alloc];
}

void* C_NSTextBlock_Init(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlock* result_ = [nSTextBlock init];
    return result_;
}

void* C_NSTextBlock_AllocTextBlock() {
    NSTextBlock* result_ = [NSTextBlock alloc];
    return result_;
}

void* C_NSTextBlock_NewTextBlock() {
    NSTextBlock* result_ = [NSTextBlock new];
    return result_;
}

void* C_NSTextBlock_Autorelease(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlock* result_ = [nSTextBlock autorelease];
    return result_;
}

void* C_NSTextBlock_Retain(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlock* result_ = [nSTextBlock retain];
    return result_;
}

void C_NSTextBlock_SetValue_Type_ForDimension(void* ptr, double val, unsigned int _type, unsigned int dimension) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock setValue:val type:_type forDimension:dimension];
}

double C_NSTextBlock_ValueForDimension(void* ptr, unsigned int dimension) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    CGFloat result_ = [nSTextBlock valueForDimension:dimension];
    return result_;
}

unsigned int C_NSTextBlock_ValueTypeForDimension(void* ptr, unsigned int dimension) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlockValueType result_ = [nSTextBlock valueTypeForDimension:dimension];
    return result_;
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
    CGFloat result_ = [nSTextBlock widthForLayer:layer edge:edge];
    return result_;
}

unsigned int C_NSTextBlock_WidthValueTypeForLayer_Edge(void* ptr, int layer, unsigned int edge) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlockValueType result_ = [nSTextBlock widthValueTypeForLayer:layer edge:edge];
    return result_;
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
    NSColor* result_ = [nSTextBlock borderColorForEdge:edge];
    return result_;
}

CGRect C_NSTextBlock_RectForLayoutAtPoint_InRect_TextContainer_CharacterRange(void* ptr, CGPoint startingPoint, CGRect rect, void* textContainer, NSRange charRange) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSRect result_ = [nSTextBlock rectForLayoutAtPoint:startingPoint inRect:rect textContainer:(NSTextContainer*)textContainer characterRange:charRange];
    return result_;
}

CGRect C_NSTextBlock_BoundsRectForContentRect_InRect_TextContainer_CharacterRange(void* ptr, CGRect contentRect, CGRect rect, void* textContainer, NSRange charRange) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSRect result_ = [nSTextBlock boundsRectForContentRect:contentRect inRect:rect textContainer:(NSTextContainer*)textContainer characterRange:charRange];
    return result_;
}

void C_NSTextBlock_DrawBackgroundWithFrame_InView_CharacterRange_LayoutManager(void* ptr, CGRect frameRect, void* controlView, NSRange charRange, void* layoutManager) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock drawBackgroundWithFrame:frameRect inView:(NSView*)controlView characterRange:charRange layoutManager:(NSLayoutManager*)layoutManager];
}

double C_NSTextBlock_ContentWidth(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    CGFloat result_ = [nSTextBlock contentWidth];
    return result_;
}

unsigned int C_NSTextBlock_ContentWidthValueType(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlockValueType result_ = [nSTextBlock contentWidthValueType];
    return result_;
}

unsigned int C_NSTextBlock_VerticalAlignment(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSTextBlockVerticalAlignment result_ = [nSTextBlock verticalAlignment];
    return result_;
}

void C_NSTextBlock_SetVerticalAlignment(void* ptr, unsigned int value) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock setVerticalAlignment:value];
}

void* C_NSTextBlock_BackgroundColor(void* ptr) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    NSColor* result_ = [nSTextBlock backgroundColor];
    return result_;
}

void C_NSTextBlock_SetBackgroundColor(void* ptr, void* value) {
    NSTextBlock* nSTextBlock = (NSTextBlock*)ptr;
    [nSTextBlock setBackgroundColor:(NSColor*)value];
}
