package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "pop_up_button.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type PopUpButton interface {
	Button
	PullsDown() bool
	SetPullsDown(pullsDown bool)
	AutoenablesItems() bool
	SetAutoenablesItems(autoenablesItems bool)
	SelectedItem() MenuItem
	TitleOfSelectedItem() string
	SelectedTag() int
	NumberOfItems() int
	ItemArray() []MenuItem
	ItemTitles() []string
	LastItem() MenuItem
	PreferredEdge() foundation.RectEdge
	SetPreferredEdge(preferredEdge foundation.RectEdge)
	AddItemWithTitle(title string)
	AddItemsWithTitles(itemTitles []string)
	InsertItemWithTitle(title string, index int)
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
	SynchronizeTitleAndSelectedItem()
}

var _ PopUpButton = (*NSPopUpButton)(nil)

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

func (p *NSPopUpButton) PullsDown() bool {
	return bool(C.PopUpButton_PullsDown(p.Ptr()))
}

func (p *NSPopUpButton) SetPullsDown(pullsDown bool) {
	C.PopUpButton_SetPullsDown(p.Ptr(), C.bool(pullsDown))
}

func (p *NSPopUpButton) AutoenablesItems() bool {
	return bool(C.PopUpButton_AutoenablesItems(p.Ptr()))
}

func (p *NSPopUpButton) SetAutoenablesItems(autoenablesItems bool) {
	C.PopUpButton_SetAutoenablesItems(p.Ptr(), C.bool(autoenablesItems))
}

func (p *NSPopUpButton) SelectedItem() MenuItem {
	return MakeMenuItem(C.PopUpButton_SelectedItem(p.Ptr()))
}

func (p *NSPopUpButton) TitleOfSelectedItem() string {
	return C.GoString(C.PopUpButton_TitleOfSelectedItem(p.Ptr()))
}

func (p *NSPopUpButton) SelectedTag() int {
	return int(C.PopUpButton_SelectedTag(p.Ptr()))
}

func (p *NSPopUpButton) Menu() Menu {
	return MakeMenu(C.PopUpButton_Menu(p.Ptr()))
}

func (p *NSPopUpButton) SetMenu(menu Menu) {
	C.PopUpButton_SetMenu(p.Ptr(), toPointer(menu))
}

func (p *NSPopUpButton) NumberOfItems() int {
	return int(C.PopUpButton_NumberOfItems(p.Ptr()))
}

func (p *NSPopUpButton) ItemArray() []MenuItem {
	var cArray C.Array = C.PopUpButton_ItemArray(p.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]MenuItem, len(result))
	for idx, r := range result {
		goArray[idx] = MakeMenuItem(r)
	}
	return goArray
}

func (p *NSPopUpButton) ItemTitles() []string {
	var cArray C.Array = C.PopUpButton_ItemTitles(p.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]string, len(result))
	for idx, r := range result {
		goArray[idx] = C.GoString((*C.char)(r))
	}
	return goArray
}

func (p *NSPopUpButton) LastItem() MenuItem {
	return MakeMenuItem(C.PopUpButton_LastItem(p.Ptr()))
}

func (p *NSPopUpButton) PreferredEdge() foundation.RectEdge {
	return foundation.RectEdge(C.PopUpButton_PreferredEdge(p.Ptr()))
}

func (p *NSPopUpButton) SetPreferredEdge(preferredEdge foundation.RectEdge) {
	C.PopUpButton_SetPreferredEdge(p.Ptr(), C.long(preferredEdge))
}

func NewPopUpButton(buttonFrame foundation.Rect, flag bool) PopUpButton {
	return MakePopUpButton(C.PopUpButton_NewPopUpButton(toNSRect(buttonFrame), C.bool(flag)))
}

func (p *NSPopUpButton) AddItemWithTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.PopUpButton_AddItemWithTitle(p.Ptr(), cTitle)
}

func (p *NSPopUpButton) AddItemsWithTitles(itemTitles []string) {
	c_itemTitlesData := make([]unsafe.Pointer, len(itemTitles))
	for idx, v := range itemTitles {
		c_itemTitlesData[idx] = unsafe.Pointer(C.CString(v))
	}
	c_itemTitles := C.Array{data: unsafe.Pointer(&c_itemTitlesData[0]), len: C.int(len(itemTitles))}
	defer func() {
		for _, p := range c_itemTitlesData {
			C.free(p)
		}
	}()
	C.PopUpButton_AddItemsWithTitles(p.Ptr(), c_itemTitles)
}

func (p *NSPopUpButton) InsertItemWithTitle(title string, index int) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.PopUpButton_InsertItemWithTitle(p.Ptr(), cTitle, C.long(index))
}

func (p *NSPopUpButton) RemoveAllItems() {
	C.PopUpButton_RemoveAllItems(p.Ptr())
}

func (p *NSPopUpButton) RemoveItemWithTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.PopUpButton_RemoveItemWithTitle(p.Ptr(), cTitle)
}

func (p *NSPopUpButton) RemoveItemAtIndex(index int) {
	C.PopUpButton_RemoveItemAtIndex(p.Ptr(), C.long(index))
}

func (p *NSPopUpButton) SelectItem(item MenuItem) {
	C.PopUpButton_SelectItem(p.Ptr(), toPointer(item))
}

func (p *NSPopUpButton) SelectItemAtIndex(index int) {
	C.PopUpButton_SelectItemAtIndex(p.Ptr(), C.long(index))
}

func (p *NSPopUpButton) SelectItemWithTag(tag int) bool {
	return bool(C.PopUpButton_SelectItemWithTag(p.Ptr(), C.long(tag)))
}

func (p *NSPopUpButton) SelectItemWithTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.PopUpButton_SelectItemWithTitle(p.Ptr(), cTitle)
}

func (p *NSPopUpButton) ItemAtIndex(index int) MenuItem {
	return MakeMenuItem(C.PopUpButton_ItemAtIndex(p.Ptr(), C.long(index)))
}

func (p *NSPopUpButton) ItemTitleAtIndex(index int) string {
	return C.GoString(C.PopUpButton_ItemTitleAtIndex(p.Ptr(), C.long(index)))
}

func (p *NSPopUpButton) ItemWithTitle(title string) MenuItem {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return MakeMenuItem(C.PopUpButton_ItemWithTitle(p.Ptr(), cTitle))
}

func (p *NSPopUpButton) IndexOfItem(item MenuItem) int {
	return int(C.PopUpButton_IndexOfItem(p.Ptr(), toPointer(item)))
}

func (p *NSPopUpButton) IndexOfItemWithTag(tag int) int {
	return int(C.PopUpButton_IndexOfItemWithTag(p.Ptr(), C.long(tag)))
}

func (p *NSPopUpButton) IndexOfItemWithTitle(title string) int {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return int(C.PopUpButton_IndexOfItemWithTitle(p.Ptr(), cTitle))
}

func (p *NSPopUpButton) SynchronizeTitleAndSelectedItem() {
	C.PopUpButton_SynchronizeTitleAndSelectedItem(p.Ptr())
}
