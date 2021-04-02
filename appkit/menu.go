package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "menu.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Menu interface {
	objc.Object
	MenuBarHeight() float64
	Font() Font
	SetFont(font Font)
	AutoenablesItems() bool
	SetAutoenablesItems(autoenablesItems bool)
	Title() string
	SetTitle(title string)
	MinimumWidth() float64
	SetMinimumWidth(minimumWidth float64)
	Size() foundation.Size
	PropertiesToUpdate() MenuProperties
	AllowsContextMenuPlugIns() bool
	SetAllowsContextMenuPlugIns(allowsContextMenuPlugIns bool)
	ShowsStateColumn() bool
	SetShowsStateColumn(showsStateColumn bool)
	HighlightedItem() MenuItem
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(userInterfaceLayoutDirection UserInterfaceLayoutDirection)
	NumberOfItems() int
	ItemArray() []MenuItem
	SetItemArray(itemArray []MenuItem)
	InsertItem(newItem MenuItem, index int)
	AddItem(newItem MenuItem)
	RemoveItem(newItem MenuItem)
	RemoveItemAtIndex(index int)
	RemoveAllItems()
	ItemAtIndex(index int) MenuItem
	ItemWithTitle(title string) MenuItem
	ItemWithTag(tag int) MenuItem
	IndexOfItem(item MenuItem) int
	IndexOfItemWithTitle(title string) int
	IndexOfItemWithTag(tag int) int
	IndexOfItemWithSubmenu(subMenu Menu) int
	SetSubmenu(subMenu Menu, item MenuItem)
	Update()
	CancelTracking()
	CancelTrackingWithoutAnimation()
}

var _ Menu = (*NSMenu)(nil)

type NSMenu struct {
	objc.NSObject
}

func MakeMenu(ptr unsafe.Pointer) *NSMenu {
	if ptr == nil {
		return nil
	}
	return &NSMenu{
		NSObject: *objc.MakeObject(ptr),
	}
}

func (m *NSMenu) MenuBarHeight() float64 {
	return float64(C.Menu_MenuBarHeight(m.Ptr()))
}

func (m *NSMenu) Font() Font {
	return MakeFont(C.Menu_Font(m.Ptr()))
}

func (m *NSMenu) SetFont(font Font) {
	C.Menu_SetFont(m.Ptr(), toPointer(font))
}

func (m *NSMenu) AutoenablesItems() bool {
	return bool(C.Menu_AutoenablesItems(m.Ptr()))
}

func (m *NSMenu) SetAutoenablesItems(autoenablesItems bool) {
	C.Menu_SetAutoenablesItems(m.Ptr(), C.bool(autoenablesItems))
}

func (m *NSMenu) Title() string {
	return C.GoString(C.Menu_Title(m.Ptr()))
}

func (m *NSMenu) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Menu_SetTitle(m.Ptr(), cTitle)
}

func (m *NSMenu) MinimumWidth() float64 {
	return float64(C.Menu_MinimumWidth(m.Ptr()))
}

func (m *NSMenu) SetMinimumWidth(minimumWidth float64) {
	C.Menu_SetMinimumWidth(m.Ptr(), C.double(minimumWidth))
}

func (m *NSMenu) Size() foundation.Size {
	return toSize(C.Menu_Size(m.Ptr()))
}

func (m *NSMenu) PropertiesToUpdate() MenuProperties {
	return MenuProperties(C.Menu_PropertiesToUpdate(m.Ptr()))
}

func (m *NSMenu) AllowsContextMenuPlugIns() bool {
	return bool(C.Menu_AllowsContextMenuPlugIns(m.Ptr()))
}

func (m *NSMenu) SetAllowsContextMenuPlugIns(allowsContextMenuPlugIns bool) {
	C.Menu_SetAllowsContextMenuPlugIns(m.Ptr(), C.bool(allowsContextMenuPlugIns))
}

func (m *NSMenu) ShowsStateColumn() bool {
	return bool(C.Menu_ShowsStateColumn(m.Ptr()))
}

func (m *NSMenu) SetShowsStateColumn(showsStateColumn bool) {
	C.Menu_SetShowsStateColumn(m.Ptr(), C.bool(showsStateColumn))
}

func (m *NSMenu) HighlightedItem() MenuItem {
	return MakeMenuItem(C.Menu_HighlightedItem(m.Ptr()))
}

func (m *NSMenu) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	return UserInterfaceLayoutDirection(C.Menu_UserInterfaceLayoutDirection(m.Ptr()))
}

func (m *NSMenu) SetUserInterfaceLayoutDirection(userInterfaceLayoutDirection UserInterfaceLayoutDirection) {
	C.Menu_SetUserInterfaceLayoutDirection(m.Ptr(), C.long(userInterfaceLayoutDirection))
}

func (m *NSMenu) NumberOfItems() int {
	return int(C.Menu_NumberOfItems(m.Ptr()))
}

func (m *NSMenu) ItemArray() []MenuItem {
	var cArray C.Array = C.Menu_ItemArray(m.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]MenuItem, len(result))
	for idx, r := range result {
		goArray[idx] = MakeMenuItem(r)
	}
	return goArray
}

func (m *NSMenu) SetItemArray(itemArray []MenuItem) {
	c_itemArrayData := make([]unsafe.Pointer, len(itemArray))
	for idx, v := range itemArray {
		c_itemArrayData[idx] = v.Ptr()
	}
	c_itemArray := C.Array{data: unsafe.Pointer(&c_itemArrayData[0]), len: C.int(len(itemArray))}
	C.Menu_SetItemArray(m.Ptr(), c_itemArray)
}

func NewMenu(title string) Menu {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return MakeMenu(C.Menu_NewMenu(cTitle))
}

func MenuBarVisible() bool {
	return bool(C.Menu_MenuBarVisible())
}

func SetMenuBarVisible(visible bool) {
	C.Menu_SetMenuBarVisible(C.bool(visible))
}

func (m *NSMenu) InsertItem(newItem MenuItem, index int) {
	C.Menu_InsertItem(m.Ptr(), toPointer(newItem), C.long(index))
}

func (m *NSMenu) AddItem(newItem MenuItem) {
	C.Menu_AddItem(m.Ptr(), toPointer(newItem))
}

func (m *NSMenu) RemoveItem(newItem MenuItem) {
	C.Menu_RemoveItem(m.Ptr(), toPointer(newItem))
}

func (m *NSMenu) RemoveItemAtIndex(index int) {
	C.Menu_RemoveItemAtIndex(m.Ptr(), C.long(index))
}

func (m *NSMenu) RemoveAllItems() {
	C.Menu_RemoveAllItems(m.Ptr())
}

func (m *NSMenu) ItemAtIndex(index int) MenuItem {
	return MakeMenuItem(C.Menu_ItemAtIndex(m.Ptr(), C.long(index)))
}

func (m *NSMenu) ItemWithTitle(title string) MenuItem {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return MakeMenuItem(C.Menu_ItemWithTitle(m.Ptr(), cTitle))
}

func (m *NSMenu) ItemWithTag(tag int) MenuItem {
	return MakeMenuItem(C.Menu_ItemWithTag(m.Ptr(), C.long(tag)))
}

func (m *NSMenu) IndexOfItem(item MenuItem) int {
	return int(C.Menu_IndexOfItem(m.Ptr(), toPointer(item)))
}

func (m *NSMenu) IndexOfItemWithTitle(title string) int {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	return int(C.Menu_IndexOfItemWithTitle(m.Ptr(), cTitle))
}

func (m *NSMenu) IndexOfItemWithTag(tag int) int {
	return int(C.Menu_IndexOfItemWithTag(m.Ptr(), C.long(tag)))
}

func (m *NSMenu) IndexOfItemWithSubmenu(subMenu Menu) int {
	return int(C.Menu_IndexOfItemWithSubmenu(m.Ptr(), toPointer(subMenu)))
}

func (m *NSMenu) SetSubmenu(subMenu Menu, item MenuItem) {
	C.Menu_SetSubmenu(m.Ptr(), toPointer(subMenu), toPointer(item))
}

func (m *NSMenu) Update() {
	C.Menu_Update(m.Ptr())
}

func (m *NSMenu) CancelTracking() {
	C.Menu_CancelTracking(m.Ptr())
}

func (m *NSMenu) CancelTrackingWithoutAnimation() {
	C.Menu_CancelTrackingWithoutAnimation(m.Ptr())
}
