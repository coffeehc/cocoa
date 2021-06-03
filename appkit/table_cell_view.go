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

func AllocTableCellView() NSTableCellView {
	return MakeTableCellView(C.C_TableCellView_Alloc())
}

func (n NSTableCellView) InitWithFrame(frameRect foundation.Rect) TableCellView {
	result_ := C.C_NSTableCellView_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeTableCellView(result_)
}

func (n NSTableCellView) InitWithCoder(coder foundation.Coder) TableCellView {
	result_ := C.C_NSTableCellView_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTableCellView(result_)
}

func (n NSTableCellView) Init() TableCellView {
	result_ := C.C_NSTableCellView_Init(n.Ptr())
	return MakeTableCellView(result_)
}

func (n NSTableCellView) ObjectValue() objc.Object {
	result_ := C.C_NSTableCellView_ObjectValue(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSTableCellView) SetObjectValue(value objc.Object) {
	C.C_NSTableCellView_SetObjectValue(n.Ptr(), objc.ExtractPtr(value))
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
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]DraggingImageComponent, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeDraggingImageComponent(r)
	}
	return goResult_
}
