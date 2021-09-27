package appkit

// #include "table_cell_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TableCellView interface {
	View
	ObjectValue() objc.Object
	SetObjectValue(value objc.Object)
	ImageView() ImageView
	SetImageView(value ImageView)
	TextField() TextField
	SetTextField(value TextField)
	BackgroundStyle() BackgroundStyle
	SetBackgroundStyle(value BackgroundStyle)
	RowSizeStyle() TableViewRowSizeStyle
	SetRowSizeStyle(value TableViewRowSizeStyle)
	DraggingImageComponents() []DraggingImageComponent
}

type NSTableCellView struct {
	NSView
}

func MakeTableCellView(ptr unsafe.Pointer) NSTableCellView {
	return NSTableCellView{
		NSView: MakeView(ptr),
	}
}

func (n NSTableCellView) InitWithFrame(frameRect foundation.Rect) NSTableCellView {
	result_ := C.C_NSTableCellView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeTableCellView(result_)
}

func (n NSTableCellView) InitWithCoder(coder foundation.Coder) NSTableCellView {
	result_ := C.C_NSTableCellView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTableCellView(result_)
}

func (n NSTableCellView) Init() NSTableCellView {
	result_ := C.C_NSTableCellView_Init(n.Ptr())
	return MakeTableCellView(result_)
}

func AllocTableCellView() NSTableCellView {
	result_ := C.C_NSTableCellView_AllocTableCellView()
	return MakeTableCellView(result_)
}

func NewTableCellView() NSTableCellView {
	result_ := C.C_NSTableCellView_NewTableCellView()
	return MakeTableCellView(result_)
}

func (n NSTableCellView) Autorelease() NSTableCellView {
	result_ := C.C_NSTableCellView_Autorelease(n.Ptr())
	return MakeTableCellView(result_)
}

func (n NSTableCellView) Retain() NSTableCellView {
	result_ := C.C_NSTableCellView_Retain(n.Ptr())
	return MakeTableCellView(result_)
}

func (n NSTableCellView) ObjectValue() objc.Object {
	result_ := C.C_NSTableCellView_ObjectValue(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTableCellView) SetObjectValue(value objc.Object) {
	C.C_NSTableCellView_SetObjectValue(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableCellView) ImageView() ImageView {
	result_ := C.C_NSTableCellView_ImageView(n.Ptr())
	return MakeImageView(result_)
}

func (n NSTableCellView) SetImageView(value ImageView) {
	C.C_NSTableCellView_SetImageView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableCellView) TextField() TextField {
	result_ := C.C_NSTableCellView_TextField(n.Ptr())
	return MakeTextField(result_)
}

func (n NSTableCellView) SetTextField(value TextField) {
	C.C_NSTableCellView_SetTextField(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSTableCellView) BackgroundStyle() BackgroundStyle {
	result_ := C.C_NSTableCellView_BackgroundStyle(n.Ptr())
	return BackgroundStyle(int(result_))
}

func (n NSTableCellView) SetBackgroundStyle(value BackgroundStyle) {
	C.C_NSTableCellView_SetBackgroundStyle(n.Ptr(), C.int(int(value)))
}

func (n NSTableCellView) RowSizeStyle() TableViewRowSizeStyle {
	result_ := C.C_NSTableCellView_RowSizeStyle(n.Ptr())
	return TableViewRowSizeStyle(int(result_))
}

func (n NSTableCellView) SetRowSizeStyle(value TableViewRowSizeStyle) {
	C.C_NSTableCellView_SetRowSizeStyle(n.Ptr(), C.int(int(value)))
}

func (n NSTableCellView) DraggingImageComponents() []DraggingImageComponent {
	result_ := C.C_NSTableCellView_DraggingImageComponents(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]DraggingImageComponent, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeDraggingImageComponent(r)
	}
	return goResult_
}
