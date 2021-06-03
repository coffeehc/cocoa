package appkit

// #include "print_operation.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PrintOperation interface {
	objc.Object
	RunOperation() bool
	CleanUpOperation()
	DeliverResult() bool
	CreateContext() GraphicsContext
	DestroyContext()
	IsCopyingOperation() bool
	PrintInfo() PrintInfo
	SetPrintInfo(value PrintInfo)
	View() View
	PreferredRenderingQuality() PrintRenderingQuality
	ShowsPrintPanel() bool
	SetShowsPrintPanel(value bool)
	ShowsProgressPanel() bool
	SetShowsProgressPanel(value bool)
	JobTitle() string
	SetJobTitle(value string)
	PrintPanel() PrintPanel
	SetPrintPanel(value PrintPanel)
	PDFPanel() PDFPanel
	SetPDFPanel(value PDFPanel)
	Context() GraphicsContext
	CurrentPage() int
	PageRange() foundation.Range
	PageOrder() PrintingPageOrder
	SetPageOrder(value PrintingPageOrder)
	CanSpawnSeparateThread() bool
	SetCanSpawnSeparateThread(value bool)
}

type NSPrintOperation struct {
	objc.NSObject
}

func MakePrintOperation(ptr unsafe.Pointer) NSPrintOperation {
	return NSPrintOperation{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocPrintOperation() NSPrintOperation {
	return MakePrintOperation(C.C_PrintOperation_Alloc())
}

func (n NSPrintOperation) Init() PrintOperation {
	result_ := C.C_NSPrintOperation_Init(n.Ptr())
	return MakePrintOperation(result_)
}

func PrintOperation_EPSOperationWithView_InsideRect_ToData(view View, rect foundation.Rect, data []byte) PrintOperation {
	result_ := C.C_NSPrintOperation_PrintOperation_EPSOperationWithView_InsideRect_ToData(objc.ExtractPtr(view), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return MakePrintOperation(result_)
}

func PrintOperation_EPSOperationWithView_InsideRect_ToData_PrintInfo(view View, rect foundation.Rect, data []byte, printInfo PrintInfo) PrintOperation {
	result_ := C.C_NSPrintOperation_PrintOperation_EPSOperationWithView_InsideRect_ToData_PrintInfo(objc.ExtractPtr(view), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, objc.ExtractPtr(printInfo))
	return MakePrintOperation(result_)
}

func PrintOperation_EPSOperationWithView_InsideRect_ToPath_PrintInfo(view View, rect foundation.Rect, path string, printInfo PrintInfo) PrintOperation {
	result_ := C.C_NSPrintOperation_PrintOperation_EPSOperationWithView_InsideRect_ToPath_PrintInfo(objc.ExtractPtr(view), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), foundation.NewString(path).Ptr(), objc.ExtractPtr(printInfo))
	return MakePrintOperation(result_)
}

func PrintOperation_PDFOperationWithView_InsideRect_ToData(view View, rect foundation.Rect, data []byte) PrintOperation {
	result_ := C.C_NSPrintOperation_PrintOperation_PDFOperationWithView_InsideRect_ToData(objc.ExtractPtr(view), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))})
	return MakePrintOperation(result_)
}

func PrintOperation_PDFOperationWithView_InsideRect_ToData_PrintInfo(view View, rect foundation.Rect, data []byte, printInfo PrintInfo) PrintOperation {
	result_ := C.C_NSPrintOperation_PrintOperation_PDFOperationWithView_InsideRect_ToData_PrintInfo(objc.ExtractPtr(view), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, objc.ExtractPtr(printInfo))
	return MakePrintOperation(result_)
}

func PrintOperation_PDFOperationWithView_InsideRect_ToPath_PrintInfo(view View, rect foundation.Rect, path string, printInfo PrintInfo) PrintOperation {
	result_ := C.C_NSPrintOperation_PrintOperation_PDFOperationWithView_InsideRect_ToPath_PrintInfo(objc.ExtractPtr(view), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), foundation.NewString(path).Ptr(), objc.ExtractPtr(printInfo))
	return MakePrintOperation(result_)
}

func PrintOperationWithView(view View) PrintOperation {
	result_ := C.C_NSPrintOperation_PrintOperationWithView(objc.ExtractPtr(view))
	return MakePrintOperation(result_)
}

func PrintOperationWithView_PrintInfo(view View, printInfo PrintInfo) PrintOperation {
	result_ := C.C_NSPrintOperation_PrintOperationWithView_PrintInfo(objc.ExtractPtr(view), objc.ExtractPtr(printInfo))
	return MakePrintOperation(result_)
}

func (n NSPrintOperation) RunOperation() bool {
	result_ := C.C_NSPrintOperation_RunOperation(n.Ptr())
	return bool(result_)
}

func (n NSPrintOperation) CleanUpOperation() {
	C.C_NSPrintOperation_CleanUpOperation(n.Ptr())
}

func (n NSPrintOperation) DeliverResult() bool {
	result_ := C.C_NSPrintOperation_DeliverResult(n.Ptr())
	return bool(result_)
}

func (n NSPrintOperation) CreateContext() GraphicsContext {
	result_ := C.C_NSPrintOperation_CreateContext(n.Ptr())
	return MakeGraphicsContext(result_)
}

func (n NSPrintOperation) DestroyContext() {
	C.C_NSPrintOperation_DestroyContext(n.Ptr())
}

func PrintOperation_CurrentOperation() PrintOperation {
	result_ := C.C_NSPrintOperation_PrintOperation_CurrentOperation()
	return MakePrintOperation(result_)
}

func PrintOperation_SetCurrentOperation(value PrintOperation) {
	C.C_NSPrintOperation_PrintOperation_SetCurrentOperation(objc.ExtractPtr(value))
}

func (n NSPrintOperation) IsCopyingOperation() bool {
	result_ := C.C_NSPrintOperation_IsCopyingOperation(n.Ptr())
	return bool(result_)
}

func (n NSPrintOperation) PrintInfo() PrintInfo {
	result_ := C.C_NSPrintOperation_PrintInfo(n.Ptr())
	return MakePrintInfo(result_)
}

func (n NSPrintOperation) SetPrintInfo(value PrintInfo) {
	C.C_NSPrintOperation_SetPrintInfo(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPrintOperation) View() View {
	result_ := C.C_NSPrintOperation_View(n.Ptr())
	return MakeView(result_)
}

func (n NSPrintOperation) PreferredRenderingQuality() PrintRenderingQuality {
	result_ := C.C_NSPrintOperation_PreferredRenderingQuality(n.Ptr())
	return PrintRenderingQuality(int(result_))
}

func (n NSPrintOperation) ShowsPrintPanel() bool {
	result_ := C.C_NSPrintOperation_ShowsPrintPanel(n.Ptr())
	return bool(result_)
}

func (n NSPrintOperation) SetShowsPrintPanel(value bool) {
	C.C_NSPrintOperation_SetShowsPrintPanel(n.Ptr(), C.bool(value))
}

func (n NSPrintOperation) ShowsProgressPanel() bool {
	result_ := C.C_NSPrintOperation_ShowsProgressPanel(n.Ptr())
	return bool(result_)
}

func (n NSPrintOperation) SetShowsProgressPanel(value bool) {
	C.C_NSPrintOperation_SetShowsProgressPanel(n.Ptr(), C.bool(value))
}

func (n NSPrintOperation) JobTitle() string {
	result_ := C.C_NSPrintOperation_JobTitle(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSPrintOperation) SetJobTitle(value string) {
	C.C_NSPrintOperation_SetJobTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSPrintOperation) PrintPanel() PrintPanel {
	result_ := C.C_NSPrintOperation_PrintPanel(n.Ptr())
	return MakePrintPanel(result_)
}

func (n NSPrintOperation) SetPrintPanel(value PrintPanel) {
	C.C_NSPrintOperation_SetPrintPanel(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPrintOperation) PDFPanel() PDFPanel {
	result_ := C.C_NSPrintOperation_PDFPanel(n.Ptr())
	return MakePDFPanel(result_)
}

func (n NSPrintOperation) SetPDFPanel(value PDFPanel) {
	C.C_NSPrintOperation_SetPDFPanel(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPrintOperation) Context() GraphicsContext {
	result_ := C.C_NSPrintOperation_Context(n.Ptr())
	return MakeGraphicsContext(result_)
}

func (n NSPrintOperation) CurrentPage() int {
	result_ := C.C_NSPrintOperation_CurrentPage(n.Ptr())
	return int(result_)
}

func (n NSPrintOperation) PageRange() foundation.Range {
	result_ := C.C_NSPrintOperation_PageRange(n.Ptr())
	return foundation.FromNSRangePointer(unsafe.Pointer(&result_))
}

func (n NSPrintOperation) PageOrder() PrintingPageOrder {
	result_ := C.C_NSPrintOperation_PageOrder(n.Ptr())
	return PrintingPageOrder(int(result_))
}

func (n NSPrintOperation) SetPageOrder(value PrintingPageOrder) {
	C.C_NSPrintOperation_SetPageOrder(n.Ptr(), C.int(int(value)))
}

func (n NSPrintOperation) CanSpawnSeparateThread() bool {
	result_ := C.C_NSPrintOperation_CanSpawnSeparateThread(n.Ptr())
	return bool(result_)
}

func (n NSPrintOperation) SetCanSpawnSeparateThread(value bool) {
	C.C_NSPrintOperation_SetCanSpawnSeparateThread(n.Ptr(), C.bool(value))
}
