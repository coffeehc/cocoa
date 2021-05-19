package appkit

// #include "text_block.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TextBlock interface {
	objc.Object
	SetValue_Type_ForDimension(val coregraphics.Float, _type TextBlockValueType, dimension TextBlockDimension)
	ValueForDimension(dimension TextBlockDimension) coregraphics.Float
	ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType
	SetContentWidth_Type(val coregraphics.Float, _type TextBlockValueType)
	SetWidth_Type_ForLayer_Edge(val coregraphics.Float, _type TextBlockValueType, layer TextBlockLayer, edge foundation.RectEdge)
	SetWidth_Type_ForLayer(val coregraphics.Float, _type TextBlockValueType, layer TextBlockLayer)
	WidthForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) coregraphics.Float
	WidthValueTypeForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) TextBlockValueType
	SetBorderColor_ForEdge(color Color, edge foundation.RectEdge)
	SetBorderColor(color Color)
	BorderColorForEdge(edge foundation.RectEdge) Color
	RectForLayoutAtPoint_InRect_TextContainer_CharacterRange(startingPoint foundation.Point, rect foundation.Rect, textContainer TextContainer, charRange foundation.Range) foundation.Rect
	BoundsRectForContentRect_InRect_TextContainer_CharacterRange(contentRect foundation.Rect, rect foundation.Rect, textContainer TextContainer, charRange foundation.Range) foundation.Rect
	DrawBackgroundWithFrame_InView_CharacterRange_LayoutManager(frameRect foundation.Rect, controlView View, charRange foundation.Range, layoutManager LayoutManager)
	ContentWidth() coregraphics.Float
	ContentWidthValueType() TextBlockValueType
	VerticalAlignment() TextBlockVerticalAlignment
	SetVerticalAlignment(value TextBlockVerticalAlignment)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
}

type NSTextBlock struct {
	objc.NSObject
}

func MakeTextBlock(ptr unsafe.Pointer) *NSTextBlock {
	if ptr == nil {
		return nil
	}
	return &NSTextBlock{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTextBlock() *NSTextBlock {
	return MakeTextBlock(C.C_TextBlock_Alloc())
}

func (n *NSTextBlock) Init() TextBlock {
	result_ := C.C_NSTextBlock_Init(n.Ptr())
	return MakeTextBlock(result_)
}

func (n *NSTextBlock) SetValue_Type_ForDimension(val coregraphics.Float, _type TextBlockValueType, dimension TextBlockDimension) {
	C.C_NSTextBlock_SetValue_Type_ForDimension(n.Ptr(), C.double(float64(val)), C.uint(uint(_type)), C.uint(uint(dimension)))
}

func (n *NSTextBlock) ValueForDimension(dimension TextBlockDimension) coregraphics.Float {
	result_ := C.C_NSTextBlock_ValueForDimension(n.Ptr(), C.uint(uint(dimension)))
	return coregraphics.Float(float64(result_))
}

func (n *NSTextBlock) ValueTypeForDimension(dimension TextBlockDimension) TextBlockValueType {
	result_ := C.C_NSTextBlock_ValueTypeForDimension(n.Ptr(), C.uint(uint(dimension)))
	return TextBlockValueType(uint(result_))
}

func (n *NSTextBlock) SetContentWidth_Type(val coregraphics.Float, _type TextBlockValueType) {
	C.C_NSTextBlock_SetContentWidth_Type(n.Ptr(), C.double(float64(val)), C.uint(uint(_type)))
}

func (n *NSTextBlock) SetWidth_Type_ForLayer_Edge(val coregraphics.Float, _type TextBlockValueType, layer TextBlockLayer, edge foundation.RectEdge) {
	C.C_NSTextBlock_SetWidth_Type_ForLayer_Edge(n.Ptr(), C.double(float64(val)), C.uint(uint(_type)), C.int(int(layer)), C.uint(uint(edge)))
}

func (n *NSTextBlock) SetWidth_Type_ForLayer(val coregraphics.Float, _type TextBlockValueType, layer TextBlockLayer) {
	C.C_NSTextBlock_SetWidth_Type_ForLayer(n.Ptr(), C.double(float64(val)), C.uint(uint(_type)), C.int(int(layer)))
}

func (n *NSTextBlock) WidthForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) coregraphics.Float {
	result_ := C.C_NSTextBlock_WidthForLayer_Edge(n.Ptr(), C.int(int(layer)), C.uint(uint(edge)))
	return coregraphics.Float(float64(result_))
}

func (n *NSTextBlock) WidthValueTypeForLayer_Edge(layer TextBlockLayer, edge foundation.RectEdge) TextBlockValueType {
	result_ := C.C_NSTextBlock_WidthValueTypeForLayer_Edge(n.Ptr(), C.int(int(layer)), C.uint(uint(edge)))
	return TextBlockValueType(uint(result_))
}

func (n *NSTextBlock) SetBorderColor_ForEdge(color Color, edge foundation.RectEdge) {
	C.C_NSTextBlock_SetBorderColor_ForEdge(n.Ptr(), objc.ExtractPtr(color), C.uint(uint(edge)))
}

func (n *NSTextBlock) SetBorderColor(color Color) {
	C.C_NSTextBlock_SetBorderColor(n.Ptr(), objc.ExtractPtr(color))
}

func (n *NSTextBlock) BorderColorForEdge(edge foundation.RectEdge) Color {
	result_ := C.C_NSTextBlock_BorderColorForEdge(n.Ptr(), C.uint(uint(edge)))
	return MakeColor(result_)
}

func (n *NSTextBlock) RectForLayoutAtPoint_InRect_TextContainer_CharacterRange(startingPoint foundation.Point, rect foundation.Rect, textContainer TextContainer, charRange foundation.Range) foundation.Rect {
	result_ := C.C_NSTextBlock_RectForLayoutAtPoint_InRect_TextContainer_CharacterRange(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(startingPoint))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(textContainer), *(*C.NSRange)(foundation.ToNSRangePointer(charRange)))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSTextBlock) BoundsRectForContentRect_InRect_TextContainer_CharacterRange(contentRect foundation.Rect, rect foundation.Rect, textContainer TextContainer, charRange foundation.Range) foundation.Rect {
	result_ := C.C_NSTextBlock_BoundsRectForContentRect_InRect_TextContainer_CharacterRange(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(textContainer), *(*C.NSRange)(foundation.ToNSRangePointer(charRange)))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSTextBlock) DrawBackgroundWithFrame_InView_CharacterRange_LayoutManager(frameRect foundation.Rect, controlView View, charRange foundation.Range, layoutManager LayoutManager) {
	C.C_NSTextBlock_DrawBackgroundWithFrame_InView_CharacterRange_LayoutManager(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))), objc.ExtractPtr(controlView), *(*C.NSRange)(foundation.ToNSRangePointer(charRange)), objc.ExtractPtr(layoutManager))
}

func (n *NSTextBlock) ContentWidth() coregraphics.Float {
	result_ := C.C_NSTextBlock_ContentWidth(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSTextBlock) ContentWidthValueType() TextBlockValueType {
	result_ := C.C_NSTextBlock_ContentWidthValueType(n.Ptr())
	return TextBlockValueType(uint(result_))
}

func (n *NSTextBlock) VerticalAlignment() TextBlockVerticalAlignment {
	result_ := C.C_NSTextBlock_VerticalAlignment(n.Ptr())
	return TextBlockVerticalAlignment(uint(result_))
}

func (n *NSTextBlock) SetVerticalAlignment(value TextBlockVerticalAlignment) {
	C.C_NSTextBlock_SetVerticalAlignment(n.Ptr(), C.uint(uint(value)))
}

func (n *NSTextBlock) BackgroundColor() Color {
	result_ := C.C_NSTextBlock_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n *NSTextBlock) SetBackgroundColor(value Color) {
	C.C_NSTextBlock_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}
