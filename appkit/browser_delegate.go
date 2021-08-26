package appkit

// #include "browser_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type BrowserDelegate struct {
	Browser_IsColumnValid                                       func(sender Browser, column int) bool
	Browser_NumberOfRowsInColumn                                func(sender Browser, column int) int
	Browser_NumberOfChildrenOfItem                              func(browser Browser, item objc.Object) int
	Browser_TitleOfColumn                                       func(sender Browser, column int) string
	Browser_ShouldTypeSelectForEvent_WithCurrentSearchString    func(browser Browser, event Event, searchString string) bool
	Browser_TypeSelectStringForRow_InColumn                     func(browser Browser, row int, column int) string
	Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString func(browser Browser, startRow int, endRow int, column int, searchString string) int
	Browser_SelectCellWithString_InColumn                       func(sender Browser, title string, column int) bool
	Browser_SelectRow_InColumn                                  func(sender Browser, row int, column int) bool
	Browser_SelectionIndexesForProposedSelection_InColumn       func(browser Browser, proposedSelectionIndexes foundation.IndexSet, column int) foundation.IndexSet
	Browser_Child_OfItem                                        func(browser Browser, index int, item objc.Object) objc.Object
	Browser_IsLeafItem                                          func(browser Browser, item objc.Object) bool
	Browser_ShouldEditItem                                      func(browser Browser, item objc.Object) bool
	Browser_ObjectValueForItem                                  func(browser Browser, item objc.Object) objc.Object
	Browser_SetObjectValue_ForItem                              func(browser Browser, object objc.Object, item objc.Object)
	RootItemForBrowser                                          func(browser Browser) objc.Object
	Browser_PreviewViewControllerForLeafItem                    func(browser Browser, item objc.Object) ViewController
	Browser_HeaderViewControllerForItem                         func(browser Browser, item objc.Object) ViewController
	Browser_CreateRowsForColumn_InMatrix                        func(sender Browser, column int, matrix Matrix)
	Browser_WillDisplayCell_AtRow_Column                        func(sender Browser, cell objc.Object, row int, column int)
	Browser_DidChangeLastColumn_ToColumn                        func(browser Browser, oldLastColumn int, column int)
	BrowserWillScroll                                           func(sender Browser)
	BrowserDidScroll                                            func(sender Browser)
	Browser_CanDragRowsWithIndexes_InColumn_WithEvent           func(browser Browser, rowIndexes foundation.IndexSet, column int, event Event) bool
	Browser_AcceptDrop_AtRow_Column_DropOperation               func(browser Browser, info objc.Object, row int, column int, dropOperation BrowserDropOperation) bool
	Browser_WriteRowsWithIndexes_InColumn_ToPasteboard          func(browser Browser, rowIndexes foundation.IndexSet, column int, pasteboard Pasteboard) bool
	Browser_ShouldSizeColumn_ForUserResize_ToWidth              func(browser Browser, columnIndex int, forUserResize bool, suggestedWidth coregraphics.Float) coregraphics.Float
	Browser_SizeToFitWidthOfColumn                              func(browser Browser, columnIndex int) coregraphics.Float
	BrowserColumnConfigurationDidChange                         func(notification foundation.Notification)
	Browser_HeightOfRow_InColumn                                func(browser Browser, row int, columnIndex int) coregraphics.Float
	Browser_ShouldShowCellExpansionForRow_Column                func(browser Browser, row int, column int) bool
}

func (delegate *BrowserDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapBrowserDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export browserDelegate_Browser_IsColumnValid
func browserDelegate_Browser_IsColumnValid(hp C.uintptr_t, sender unsafe.Pointer, column C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_IsColumnValid(MakeBrowser(sender), int(column))
	return C.bool(result)
}

//export browserDelegate_Browser_NumberOfRowsInColumn
func browserDelegate_Browser_NumberOfRowsInColumn(hp C.uintptr_t, sender unsafe.Pointer, column C.int) C.int {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_NumberOfRowsInColumn(MakeBrowser(sender), int(column))
	return C.int(result)
}

//export browserDelegate_Browser_NumberOfChildrenOfItem
func browserDelegate_Browser_NumberOfChildrenOfItem(hp C.uintptr_t, browser unsafe.Pointer, item unsafe.Pointer) C.int {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_NumberOfChildrenOfItem(MakeBrowser(browser), objc.MakeObject(item))
	return C.int(result)
}

//export browserDelegate_Browser_TitleOfColumn
func browserDelegate_Browser_TitleOfColumn(hp C.uintptr_t, sender unsafe.Pointer, column C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_TitleOfColumn(MakeBrowser(sender), int(column))
	return foundation.NewString(result).Ptr()
}

//export browserDelegate_Browser_ShouldTypeSelectForEvent_WithCurrentSearchString
func browserDelegate_Browser_ShouldTypeSelectForEvent_WithCurrentSearchString(hp C.uintptr_t, browser unsafe.Pointer, event unsafe.Pointer, searchString unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_ShouldTypeSelectForEvent_WithCurrentSearchString(MakeBrowser(browser), MakeEvent(event), foundation.MakeString(searchString).String())
	return C.bool(result)
}

//export browserDelegate_Browser_TypeSelectStringForRow_InColumn
func browserDelegate_Browser_TypeSelectStringForRow_InColumn(hp C.uintptr_t, browser unsafe.Pointer, row C.int, column C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_TypeSelectStringForRow_InColumn(MakeBrowser(browser), int(row), int(column))
	return foundation.NewString(result).Ptr()
}

//export browserDelegate_Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString
func browserDelegate_Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString(hp C.uintptr_t, browser unsafe.Pointer, startRow C.int, endRow C.int, column C.int, searchString unsafe.Pointer) C.int {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString(MakeBrowser(browser), int(startRow), int(endRow), int(column), foundation.MakeString(searchString).String())
	return C.int(result)
}

//export browserDelegate_Browser_SelectCellWithString_InColumn
func browserDelegate_Browser_SelectCellWithString_InColumn(hp C.uintptr_t, sender unsafe.Pointer, title unsafe.Pointer, column C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_SelectCellWithString_InColumn(MakeBrowser(sender), foundation.MakeString(title).String(), int(column))
	return C.bool(result)
}

//export browserDelegate_Browser_SelectRow_InColumn
func browserDelegate_Browser_SelectRow_InColumn(hp C.uintptr_t, sender unsafe.Pointer, row C.int, column C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_SelectRow_InColumn(MakeBrowser(sender), int(row), int(column))
	return C.bool(result)
}

//export browserDelegate_Browser_SelectionIndexesForProposedSelection_InColumn
func browserDelegate_Browser_SelectionIndexesForProposedSelection_InColumn(hp C.uintptr_t, browser unsafe.Pointer, proposedSelectionIndexes unsafe.Pointer, column C.int) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_SelectionIndexesForProposedSelection_InColumn(MakeBrowser(browser), foundation.MakeIndexSet(proposedSelectionIndexes), int(column))
	return objc.ExtractPtr(result)
}

//export browserDelegate_Browser_Child_OfItem
func browserDelegate_Browser_Child_OfItem(hp C.uintptr_t, browser unsafe.Pointer, index C.int, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_Child_OfItem(MakeBrowser(browser), int(index), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export browserDelegate_Browser_IsLeafItem
func browserDelegate_Browser_IsLeafItem(hp C.uintptr_t, browser unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_IsLeafItem(MakeBrowser(browser), objc.MakeObject(item))
	return C.bool(result)
}

//export browserDelegate_Browser_ShouldEditItem
func browserDelegate_Browser_ShouldEditItem(hp C.uintptr_t, browser unsafe.Pointer, item unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_ShouldEditItem(MakeBrowser(browser), objc.MakeObject(item))
	return C.bool(result)
}

//export browserDelegate_Browser_ObjectValueForItem
func browserDelegate_Browser_ObjectValueForItem(hp C.uintptr_t, browser unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_ObjectValueForItem(MakeBrowser(browser), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export browserDelegate_Browser_SetObjectValue_ForItem
func browserDelegate_Browser_SetObjectValue_ForItem(hp C.uintptr_t, browser unsafe.Pointer, object unsafe.Pointer, item unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	delegate.Browser_SetObjectValue_ForItem(MakeBrowser(browser), objc.MakeObject(object), objc.MakeObject(item))
}

//export browserDelegate_RootItemForBrowser
func browserDelegate_RootItemForBrowser(hp C.uintptr_t, browser unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.RootItemForBrowser(MakeBrowser(browser))
	return objc.ExtractPtr(result)
}

//export browserDelegate_Browser_PreviewViewControllerForLeafItem
func browserDelegate_Browser_PreviewViewControllerForLeafItem(hp C.uintptr_t, browser unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_PreviewViewControllerForLeafItem(MakeBrowser(browser), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export browserDelegate_Browser_HeaderViewControllerForItem
func browserDelegate_Browser_HeaderViewControllerForItem(hp C.uintptr_t, browser unsafe.Pointer, item unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_HeaderViewControllerForItem(MakeBrowser(browser), objc.MakeObject(item))
	return objc.ExtractPtr(result)
}

//export browserDelegate_Browser_CreateRowsForColumn_InMatrix
func browserDelegate_Browser_CreateRowsForColumn_InMatrix(hp C.uintptr_t, sender unsafe.Pointer, column C.int, matrix unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	delegate.Browser_CreateRowsForColumn_InMatrix(MakeBrowser(sender), int(column), MakeMatrix(matrix))
}

//export browserDelegate_Browser_WillDisplayCell_AtRow_Column
func browserDelegate_Browser_WillDisplayCell_AtRow_Column(hp C.uintptr_t, sender unsafe.Pointer, cell unsafe.Pointer, row C.int, column C.int) {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	delegate.Browser_WillDisplayCell_AtRow_Column(MakeBrowser(sender), objc.MakeObject(cell), int(row), int(column))
}

//export browserDelegate_Browser_DidChangeLastColumn_ToColumn
func browserDelegate_Browser_DidChangeLastColumn_ToColumn(hp C.uintptr_t, browser unsafe.Pointer, oldLastColumn C.int, column C.int) {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	delegate.Browser_DidChangeLastColumn_ToColumn(MakeBrowser(browser), int(oldLastColumn), int(column))
}

//export browserDelegate_BrowserWillScroll
func browserDelegate_BrowserWillScroll(hp C.uintptr_t, sender unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	delegate.BrowserWillScroll(MakeBrowser(sender))
}

//export browserDelegate_BrowserDidScroll
func browserDelegate_BrowserDidScroll(hp C.uintptr_t, sender unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	delegate.BrowserDidScroll(MakeBrowser(sender))
}

//export browserDelegate_Browser_CanDragRowsWithIndexes_InColumn_WithEvent
func browserDelegate_Browser_CanDragRowsWithIndexes_InColumn_WithEvent(hp C.uintptr_t, browser unsafe.Pointer, rowIndexes unsafe.Pointer, column C.int, event unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_CanDragRowsWithIndexes_InColumn_WithEvent(MakeBrowser(browser), foundation.MakeIndexSet(rowIndexes), int(column), MakeEvent(event))
	return C.bool(result)
}

//export browserDelegate_Browser_AcceptDrop_AtRow_Column_DropOperation
func browserDelegate_Browser_AcceptDrop_AtRow_Column_DropOperation(hp C.uintptr_t, browser unsafe.Pointer, info unsafe.Pointer, row C.int, column C.int, dropOperation C.uint) C.bool {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_AcceptDrop_AtRow_Column_DropOperation(MakeBrowser(browser), objc.MakeObject(info), int(row), int(column), BrowserDropOperation(uint(dropOperation)))
	return C.bool(result)
}

//export browserDelegate_Browser_WriteRowsWithIndexes_InColumn_ToPasteboard
func browserDelegate_Browser_WriteRowsWithIndexes_InColumn_ToPasteboard(hp C.uintptr_t, browser unsafe.Pointer, rowIndexes unsafe.Pointer, column C.int, pasteboard unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_WriteRowsWithIndexes_InColumn_ToPasteboard(MakeBrowser(browser), foundation.MakeIndexSet(rowIndexes), int(column), MakePasteboard(pasteboard))
	return C.bool(result)
}

//export browserDelegate_Browser_ShouldSizeColumn_ForUserResize_ToWidth
func browserDelegate_Browser_ShouldSizeColumn_ForUserResize_ToWidth(hp C.uintptr_t, browser unsafe.Pointer, columnIndex C.int, forUserResize C.bool, suggestedWidth C.double) C.double {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_ShouldSizeColumn_ForUserResize_ToWidth(MakeBrowser(browser), int(columnIndex), bool(forUserResize), coregraphics.Float(float64(suggestedWidth)))
	return C.double(float64(result))
}

//export browserDelegate_Browser_SizeToFitWidthOfColumn
func browserDelegate_Browser_SizeToFitWidthOfColumn(hp C.uintptr_t, browser unsafe.Pointer, columnIndex C.int) C.double {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_SizeToFitWidthOfColumn(MakeBrowser(browser), int(columnIndex))
	return C.double(float64(result))
}

//export browserDelegate_BrowserColumnConfigurationDidChange
func browserDelegate_BrowserColumnConfigurationDidChange(hp C.uintptr_t, notification unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	delegate.BrowserColumnConfigurationDidChange(foundation.MakeNotification(notification))
}

//export browserDelegate_Browser_HeightOfRow_InColumn
func browserDelegate_Browser_HeightOfRow_InColumn(hp C.uintptr_t, browser unsafe.Pointer, row C.int, columnIndex C.int) C.double {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_HeightOfRow_InColumn(MakeBrowser(browser), int(row), int(columnIndex))
	return C.double(float64(result))
}

//export browserDelegate_Browser_ShouldShowCellExpansionForRow_Column
func browserDelegate_Browser_ShouldShowCellExpansionForRow_Column(hp C.uintptr_t, browser unsafe.Pointer, row C.int, column C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	result := delegate.Browser_ShouldShowCellExpansionForRow_Column(MakeBrowser(browser), int(row), int(column))
	return C.bool(result)
}

//export BrowserDelegate_RespondsTo
func BrowserDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*BrowserDelegate)
	switch selName {
	case "browser:isColumnValid:":
		return delegate.Browser_IsColumnValid != nil
	case "browser:numberOfRowsInColumn:":
		return delegate.Browser_NumberOfRowsInColumn != nil
	case "browser:numberOfChildrenOfItem:":
		return delegate.Browser_NumberOfChildrenOfItem != nil
	case "browser:titleOfColumn:":
		return delegate.Browser_TitleOfColumn != nil
	case "browser:shouldTypeSelectForEvent:withCurrentSearchString:":
		return delegate.Browser_ShouldTypeSelectForEvent_WithCurrentSearchString != nil
	case "browser:typeSelectStringForRow:inColumn:":
		return delegate.Browser_TypeSelectStringForRow_InColumn != nil
	case "browser:nextTypeSelectMatchFromRow:toRow:inColumn:forString:":
		return delegate.Browser_NextTypeSelectMatchFromRow_ToRow_InColumn_ForString != nil
	case "browser:selectCellWithString:inColumn:":
		return delegate.Browser_SelectCellWithString_InColumn != nil
	case "browser:selectRow:inColumn:":
		return delegate.Browser_SelectRow_InColumn != nil
	case "browser:selectionIndexesForProposedSelection:inColumn:":
		return delegate.Browser_SelectionIndexesForProposedSelection_InColumn != nil
	case "browser:child:ofItem:":
		return delegate.Browser_Child_OfItem != nil
	case "browser:isLeafItem:":
		return delegate.Browser_IsLeafItem != nil
	case "browser:shouldEditItem:":
		return delegate.Browser_ShouldEditItem != nil
	case "browser:objectValueForItem:":
		return delegate.Browser_ObjectValueForItem != nil
	case "browser:setObjectValue:forItem:":
		return delegate.Browser_SetObjectValue_ForItem != nil
	case "rootItemForBrowser:":
		return delegate.RootItemForBrowser != nil
	case "browser:previewViewControllerForLeafItem:":
		return delegate.Browser_PreviewViewControllerForLeafItem != nil
	case "browser:headerViewControllerForItem:":
		return delegate.Browser_HeaderViewControllerForItem != nil
	case "browser:createRowsForColumn:inMatrix:":
		return delegate.Browser_CreateRowsForColumn_InMatrix != nil
	case "browser:willDisplayCell:atRow:column:":
		return delegate.Browser_WillDisplayCell_AtRow_Column != nil
	case "browser:didChangeLastColumn:toColumn:":
		return delegate.Browser_DidChangeLastColumn_ToColumn != nil
	case "browserWillScroll:":
		return delegate.BrowserWillScroll != nil
	case "browserDidScroll:":
		return delegate.BrowserDidScroll != nil
	case "browser:canDragRowsWithIndexes:inColumn:withEvent:":
		return delegate.Browser_CanDragRowsWithIndexes_InColumn_WithEvent != nil
	case "browser:acceptDrop:atRow:column:dropOperation:":
		return delegate.Browser_AcceptDrop_AtRow_Column_DropOperation != nil
	case "browser:writeRowsWithIndexes:inColumn:toPasteboard:":
		return delegate.Browser_WriteRowsWithIndexes_InColumn_ToPasteboard != nil
	case "browser:shouldSizeColumn:forUserResize:toWidth:":
		return delegate.Browser_ShouldSizeColumn_ForUserResize_ToWidth != nil
	case "browser:sizeToFitWidthOfColumn:":
		return delegate.Browser_SizeToFitWidthOfColumn != nil
	case "browserColumnConfigurationDidChange:":
		return delegate.BrowserColumnConfigurationDidChange != nil
	case "browser:heightOfRow:inColumn:":
		return delegate.Browser_HeightOfRow_InColumn != nil
	case "browser:shouldShowCellExpansionForRow:column:":
		return delegate.Browser_ShouldShowCellExpansionForRow_Column != nil
	default:
		return false
	}
}
