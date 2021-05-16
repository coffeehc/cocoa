package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "pop_up_button.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PopUpButton interface {
	Button
	AddItemWithTitle(title string)
	AddItemsWithTitles(itemTitles []string)
	InsertItemWithTitle_AtIndex(title string, index int)
	RemoveAllItems()
	RemoveItemWithTitle(title string)
	RemoveItemAtIndex(index int)
	SelectItem(item MenuItem)
	SelectItemAtIndex(index int)
	SelectItemWithTag(tag int) bool
	SelectItemWithTitle(title string)
	ItemAtIndex(index int) MenuItem
	ItemTitleAtIndex(index int) string
	ItemWithTitle(title string) MenuItem
	IndexOfItem(item MenuItem) int
	IndexOfItemWithTag(tag int) int
	IndexOfItemWithTitle(title string) int
	IndexOfItemWithRepresentedObject(obj objc.Object) int
	IndexOfItemWithTarget_AndAction(target objc.Object, actionSelector *objc.Selector) int
	SynchronizeTitleAndSelectedItem()
	PullsDown() bool
	SetPullsDown(value bool)
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	SelectedItem() MenuItem
	TitleOfSelectedItem() string
	IndexOfSelectedItem() int
	SelectedTag() int
	NumberOfItems() int
	ItemArray() []MenuItem
	ItemTitles() []string
	LastItem() MenuItem
	PreferredEdge() foundation.RectEdge
	SetPreferredEdge(value foundation.RectEdge)
}

type NSPopUpButton struct {
	NSButton
}

func MakePopUpButton(ptr unsafe.Pointer) *NSPopUpButton {
	if ptr == nil {
		return nil
	}
	return &NSPopUpButton{
		NSButton: *MakeButton(ptr),
	}
}

func AllocPopUpButton() *NSPopUpButton {
	return MakePopUpButton(C.C_PopUpButton_Alloc())
}

func (n *NSPopUpButton) InitWithFrame_PullsDown(buttonFrame foundation.Rect, flag bool) PopUpButton {
	result := C.C_NSPopUpButton_InitWithFrame_PullsDown(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(buttonFrame))), C.bool(flag))
	return MakePopUpButton(result)
}

func (n *NSPopUpButton) InitWithCoder(coder foundation.Coder) PopUpButton {
	result := C.C_NSPopUpButton_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakePopUpButton(result)
}

func (n *NSPopUpButton) Init() PopUpButton {
	result := C.C_NSPopUpButton_Init(n.Ptr())
	return MakePopUpButton(result)
}

func (n *NSPopUpButton) AddItemWithTitle(title string) {
	C.C_NSPopUpButton_AddItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
}

func (n *NSPopUpButton) AddItemsWithTitles(itemTitles []string) {
	cItemTitlesData := make([]unsafe.Pointer, len(itemTitles))
	for idx, v := range itemTitles {
		cItemTitlesData[idx] = foundation.NewString(v).Ptr()
	}
	cItemTitles := C.Array{data: unsafe.Pointer(&cItemTitlesData[0]), len: C.int(len(itemTitles))}
	C.C_NSPopUpButton_AddItemsWithTitles(n.Ptr(), cItemTitles)
}

func (n *NSPopUpButton) InsertItemWithTitle_AtIndex(title string, index int) {
	C.C_NSPopUpButton_InsertItemWithTitle_AtIndex(n.Ptr(), foundation.NewString(title).Ptr(), C.int(index))
}

func (n *NSPopUpButton) RemoveAllItems() {
	C.C_NSPopUpButton_RemoveAllItems(n.Ptr())
}

func (n *NSPopUpButton) RemoveItemWithTitle(title string) {
	C.C_NSPopUpButton_RemoveItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
}

func (n *NSPopUpButton) RemoveItemAtIndex(index int) {
	C.C_NSPopUpButton_RemoveItemAtIndex(n.Ptr(), C.int(index))
}

func (n *NSPopUpButton) SelectItem(item MenuItem) {
	C.C_NSPopUpButton_SelectItem(n.Ptr(), objc.ExtractPtr(item))
}

func (n *NSPopUpButton) SelectItemAtIndex(index int) {
	C.C_NSPopUpButton_SelectItemAtIndex(n.Ptr(), C.int(index))
}

func (n *NSPopUpButton) SelectItemWithTag(tag int) bool {
	result := C.C_NSPopUpButton_SelectItemWithTag(n.Ptr(), C.int(tag))
	return bool(result)
}

func (n *NSPopUpButton) SelectItemWithTitle(title string) {
	C.C_NSPopUpButton_SelectItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
}

func (n *NSPopUpButton) ItemAtIndex(index int) MenuItem {
	result := C.C_NSPopUpButton_ItemAtIndex(n.Ptr(), C.int(index))
	return MakeMenuItem(result)
}

func (n *NSPopUpButton) ItemTitleAtIndex(index int) string {
	result := C.C_NSPopUpButton_ItemTitleAtIndex(n.Ptr(), C.int(index))
	return foundation.MakeString(result).String()
}

func (n *NSPopUpButton) ItemWithTitle(title string) MenuItem {
	result := C.C_NSPopUpButton_ItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
	return MakeMenuItem(result)
}

func (n *NSPopUpButton) IndexOfItem(item MenuItem) int {
	result := C.C_NSPopUpButton_IndexOfItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result)
}

func (n *NSPopUpButton) IndexOfItemWithTag(tag int) int {
	result := C.C_NSPopUpButton_IndexOfItemWithTag(n.Ptr(), C.int(tag))
	return int(result)
}

func (n *NSPopUpButton) IndexOfItemWithTitle(title string) int {
	result := C.C_NSPopUpButton_IndexOfItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
	return int(result)
}

func (n *NSPopUpButton) IndexOfItemWithRepresentedObject(obj objc.Object) int {
	result := C.C_NSPopUpButton_IndexOfItemWithRepresentedObject(n.Ptr(), objc.ExtractPtr(obj))
	return int(result)
}

func (n *NSPopUpButton) IndexOfItemWithTarget_AndAction(target objc.Object, actionSelector *objc.Selector) int {
	result := C.C_NSPopUpButton_IndexOfItemWithTarget_AndAction(n.Ptr(), objc.ExtractPtr(target), objc.ExtractPtr(actionSelector))
	return int(result)
}

func (n *NSPopUpButton) SynchronizeTitleAndSelectedItem() {
	C.C_NSPopUpButton_SynchronizeTitleAndSelectedItem(n.Ptr())
}

func (n *NSPopUpButton) PullsDown() bool {
	result := C.C_NSPopUpButton_PullsDown(n.Ptr())
	return bool(result)
}

func (n *NSPopUpButton) SetPullsDown(value bool) {
	C.C_NSPopUpButton_SetPullsDown(n.Ptr(), C.bool(value))
}

func (n *NSPopUpButton) AutoenablesItems() bool {
	result := C.C_NSPopUpButton_AutoenablesItems(n.Ptr())
	return bool(result)
}

func (n *NSPopUpButton) SetAutoenablesItems(value bool) {
	C.C_NSPopUpButton_SetAutoenablesItems(n.Ptr(), C.bool(value))
}

func (n *NSPopUpButton) SelectedItem() MenuItem {
	result := C.C_NSPopUpButton_SelectedItem(n.Ptr())
	return MakeMenuItem(result)
}

func (n *NSPopUpButton) TitleOfSelectedItem() string {
	result := C.C_NSPopUpButton_TitleOfSelectedItem(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSPopUpButton) IndexOfSelectedItem() int {
	result := C.C_NSPopUpButton_IndexOfSelectedItem(n.Ptr())
	return int(result)
}

func (n *NSPopUpButton) SelectedTag() int {
	result := C.C_NSPopUpButton_SelectedTag(n.Ptr())
	return int(result)
}

func (n *NSPopUpButton) NumberOfItems() int {
	result := C.C_NSPopUpButton_NumberOfItems(n.Ptr())
	return int(result)
}

func (n *NSPopUpButton) ItemArray() []MenuItem {
	result := C.C_NSPopUpButton_ItemArray(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]MenuItem, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeMenuItem(r)
	}
	return goResult
}

func (n *NSPopUpButton) ItemTitles() []string {
	result := C.C_NSPopUpButton_ItemTitles(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]string, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = foundation.MakeString(r).String()
	}
	return goResult
}

func (n *NSPopUpButton) LastItem() MenuItem {
	result := C.C_NSPopUpButton_LastItem(n.Ptr())
	return MakeMenuItem(result)
}

func (n *NSPopUpButton) PreferredEdge() foundation.RectEdge {
	result := C.C_NSPopUpButton_PreferredEdge(n.Ptr())
	return foundation.RectEdge(uint(result))
}

func (n *NSPopUpButton) SetPreferredEdge(value foundation.RectEdge) {
	C.C_NSPopUpButton_SetPreferredEdge(n.Ptr(), C.uint(uint(value)))
}
