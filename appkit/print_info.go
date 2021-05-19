package appkit

// #include "print_info.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PrintInfo interface {
	objc.Object
	SetUpPrintOperationDefaultValues()
	UpdateFromPMPageFormat()
	UpdateFromPMPrintSettings()
	TakeSettingsFromPDFInfo(inPDFInfo PDFInfo)
	PaperSize() foundation.Size
	SetPaperSize(value foundation.Size)
	TopMargin() coregraphics.Float
	SetTopMargin(value coregraphics.Float)
	BottomMargin() coregraphics.Float
	SetBottomMargin(value coregraphics.Float)
	LeftMargin() coregraphics.Float
	SetLeftMargin(value coregraphics.Float)
	RightMargin() coregraphics.Float
	SetRightMargin(value coregraphics.Float)
	ImageablePageBounds() foundation.Rect
	Orientation() PaperOrientation
	SetOrientation(value PaperOrientation)
	PaperName() PrinterPaperName
	SetPaperName(value PrinterPaperName)
	LocalizedPaperName() string
	HorizontalPagination() PrintingPaginationMode
	SetHorizontalPagination(value PrintingPaginationMode)
	VerticalPagination() PrintingPaginationMode
	SetVerticalPagination(value PrintingPaginationMode)
	IsHorizontallyCentered() bool
	SetHorizontallyCentered(value bool)
	IsVerticallyCentered() bool
	SetVerticallyCentered(value bool)
	Printer() Printer
	SetPrinter(value Printer)
	JobDisposition() PrintJobDispositionValue
	SetJobDisposition(value PrintJobDispositionValue)
	IsSelectionOnly() bool
	SetSelectionOnly(value bool)
	ScalingFactor() coregraphics.Float
	SetScalingFactor(value coregraphics.Float)
}

type NSPrintInfo struct {
	objc.NSObject
}

func MakePrintInfo(ptr unsafe.Pointer) *NSPrintInfo {
	if ptr == nil {
		return nil
	}
	return &NSPrintInfo{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocPrintInfo() *NSPrintInfo {
	return MakePrintInfo(C.C_PrintInfo_Alloc())
}

func (n *NSPrintInfo) Init() PrintInfo {
	result_ := C.C_NSPrintInfo_Init(n.Ptr())
	return MakePrintInfo(result_)
}

func (n *NSPrintInfo) InitWithCoder(coder foundation.Coder) PrintInfo {
	result_ := C.C_NSPrintInfo_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakePrintInfo(result_)
}

func (n *NSPrintInfo) SetUpPrintOperationDefaultValues() {
	C.C_NSPrintInfo_SetUpPrintOperationDefaultValues(n.Ptr())
}

func (n *NSPrintInfo) UpdateFromPMPageFormat() {
	C.C_NSPrintInfo_UpdateFromPMPageFormat(n.Ptr())
}

func (n *NSPrintInfo) UpdateFromPMPrintSettings() {
	C.C_NSPrintInfo_UpdateFromPMPrintSettings(n.Ptr())
}

func (n *NSPrintInfo) TakeSettingsFromPDFInfo(inPDFInfo PDFInfo) {
	C.C_NSPrintInfo_TakeSettingsFromPDFInfo(n.Ptr(), objc.ExtractPtr(inPDFInfo))
}

func SharedPrintInfo() PrintInfo {
	result_ := C.C_NSPrintInfo_SharedPrintInfo()
	return MakePrintInfo(result_)
}

func SetSharedPrintInfo(value PrintInfo) {
	C.C_NSPrintInfo_SetSharedPrintInfo(objc.ExtractPtr(value))
}

func (n *NSPrintInfo) PaperSize() foundation.Size {
	result_ := C.C_NSPrintInfo_PaperSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n *NSPrintInfo) SetPaperSize(value foundation.Size) {
	C.C_NSPrintInfo_SetPaperSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n *NSPrintInfo) TopMargin() coregraphics.Float {
	result_ := C.C_NSPrintInfo_TopMargin(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSPrintInfo) SetTopMargin(value coregraphics.Float) {
	C.C_NSPrintInfo_SetTopMargin(n.Ptr(), C.double(float64(value)))
}

func (n *NSPrintInfo) BottomMargin() coregraphics.Float {
	result_ := C.C_NSPrintInfo_BottomMargin(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSPrintInfo) SetBottomMargin(value coregraphics.Float) {
	C.C_NSPrintInfo_SetBottomMargin(n.Ptr(), C.double(float64(value)))
}

func (n *NSPrintInfo) LeftMargin() coregraphics.Float {
	result_ := C.C_NSPrintInfo_LeftMargin(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSPrintInfo) SetLeftMargin(value coregraphics.Float) {
	C.C_NSPrintInfo_SetLeftMargin(n.Ptr(), C.double(float64(value)))
}

func (n *NSPrintInfo) RightMargin() coregraphics.Float {
	result_ := C.C_NSPrintInfo_RightMargin(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSPrintInfo) SetRightMargin(value coregraphics.Float) {
	C.C_NSPrintInfo_SetRightMargin(n.Ptr(), C.double(float64(value)))
}

func (n *NSPrintInfo) ImageablePageBounds() foundation.Rect {
	result_ := C.C_NSPrintInfo_ImageablePageBounds(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSPrintInfo) Orientation() PaperOrientation {
	result_ := C.C_NSPrintInfo_Orientation(n.Ptr())
	return PaperOrientation(int(result_))
}

func (n *NSPrintInfo) SetOrientation(value PaperOrientation) {
	C.C_NSPrintInfo_SetOrientation(n.Ptr(), C.int(int(value)))
}

func (n *NSPrintInfo) PaperName() PrinterPaperName {
	result_ := C.C_NSPrintInfo_PaperName(n.Ptr())
	return PrinterPaperName(foundation.MakeString(result_).String())
}

func (n *NSPrintInfo) SetPaperName(value PrinterPaperName) {
	C.C_NSPrintInfo_SetPaperName(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSPrintInfo) LocalizedPaperName() string {
	result_ := C.C_NSPrintInfo_LocalizedPaperName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSPrintInfo) HorizontalPagination() PrintingPaginationMode {
	result_ := C.C_NSPrintInfo_HorizontalPagination(n.Ptr())
	return PrintingPaginationMode(uint(result_))
}

func (n *NSPrintInfo) SetHorizontalPagination(value PrintingPaginationMode) {
	C.C_NSPrintInfo_SetHorizontalPagination(n.Ptr(), C.uint(uint(value)))
}

func (n *NSPrintInfo) VerticalPagination() PrintingPaginationMode {
	result_ := C.C_NSPrintInfo_VerticalPagination(n.Ptr())
	return PrintingPaginationMode(uint(result_))
}

func (n *NSPrintInfo) SetVerticalPagination(value PrintingPaginationMode) {
	C.C_NSPrintInfo_SetVerticalPagination(n.Ptr(), C.uint(uint(value)))
}

func (n *NSPrintInfo) IsHorizontallyCentered() bool {
	result_ := C.C_NSPrintInfo_IsHorizontallyCentered(n.Ptr())
	return bool(result_)
}

func (n *NSPrintInfo) SetHorizontallyCentered(value bool) {
	C.C_NSPrintInfo_SetHorizontallyCentered(n.Ptr(), C.bool(value))
}

func (n *NSPrintInfo) IsVerticallyCentered() bool {
	result_ := C.C_NSPrintInfo_IsVerticallyCentered(n.Ptr())
	return bool(result_)
}

func (n *NSPrintInfo) SetVerticallyCentered(value bool) {
	C.C_NSPrintInfo_SetVerticallyCentered(n.Ptr(), C.bool(value))
}

func (n *NSPrintInfo) Printer() Printer {
	result_ := C.C_NSPrintInfo_Printer(n.Ptr())
	return MakePrinter(result_)
}

func (n *NSPrintInfo) SetPrinter(value Printer) {
	C.C_NSPrintInfo_SetPrinter(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSPrintInfo) JobDisposition() PrintJobDispositionValue {
	result_ := C.C_NSPrintInfo_JobDisposition(n.Ptr())
	return PrintJobDispositionValue(foundation.MakeString(result_).String())
}

func (n *NSPrintInfo) SetJobDisposition(value PrintJobDispositionValue) {
	C.C_NSPrintInfo_SetJobDisposition(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSPrintInfo) IsSelectionOnly() bool {
	result_ := C.C_NSPrintInfo_IsSelectionOnly(n.Ptr())
	return bool(result_)
}

func (n *NSPrintInfo) SetSelectionOnly(value bool) {
	C.C_NSPrintInfo_SetSelectionOnly(n.Ptr(), C.bool(value))
}

func (n *NSPrintInfo) ScalingFactor() coregraphics.Float {
	result_ := C.C_NSPrintInfo_ScalingFactor(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSPrintInfo) SetScalingFactor(value coregraphics.Float) {
	C.C_NSPrintInfo_SetScalingFactor(n.Ptr(), C.double(float64(value)))
}

func PrintInfo_DefaultPrinter() Printer {
	result_ := C.C_NSPrintInfo_PrintInfo_DefaultPrinter()
	return MakePrinter(result_)
}
