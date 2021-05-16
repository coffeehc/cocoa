package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "menu.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Menu interface {
	objc.Object
	InsertItem_AtIndex(newItem MenuItem, index int)
	InsertItemWithTitle_Action_KeyEquivalent_AtIndex(_string string, selector *objc.Selector, charCode string, index int) MenuItem
	AddItem(newItem MenuItem)
	AddItemWithTitle_Action_KeyEquivalent(_string string, selector *objc.Selector, charCode string) MenuItem
	RemoveItem(item MenuItem)
	RemoveItemAtIndex(index int)
	ItemChanged(item MenuItem)
	RemoveAllItems()
	ItemWithTag(tag int) MenuItem
	ItemWithTitle(title string) MenuItem
	ItemAtIndex(index int) MenuItem
	IndexOfItem(item MenuItem) int
	IndexOfItemWithTitle(title string) int
	IndexOfItemWithTag(tag int) int
	IndexOfItemWithTarget_AndAction(target objc.Object, actionSelector *objc.Selector) int
	IndexOfItemWithRepresentedObject(object objc.Object) int
	IndexOfItemWithSubmenu(submenu Menu) int
	SetSubmenu_ForItem(menu Menu, item MenuItem)
	SubmenuAction(sender objc.Object)
	Update()
	PerformKeyEquivalent(event Event) bool
	PerformActionForItemAtIndex(index int)
	PopUpMenuPositioningItem_AtLocation_InView(item MenuItem, location foundation.Point, view View) bool
	CancelTracking()
	CancelTrackingWithoutAnimation()
	MenuBarHeight() coregraphics.Float
	NumberOfItems() int
	ItemArray() []MenuItem
	SetItemArray(value []MenuItem)
	Supermenu() Menu
	SetSupermenu(value Menu)
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	Font() Font
	SetFont(value Font)
	Title() string
	SetTitle(value string)
	MinimumWidth() coregraphics.Float
	SetMinimumWidth(value coregraphics.Float)
	Size() foundation.Size
	PropertiesToUpdate() MenuProperties
	AllowsContextMenuPlugIns() bool
	SetAllowsContextMenuPlugIns(value bool)
	ShowsStateColumn() bool
	SetShowsStateColumn(value bool)
	HighlightedItem() MenuItem
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
}

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

func AllocMenu() *NSMenu {
	return MakeMenu(C.C_Menu_Alloc())
}

func (n *NSMenu) InitWithTitle(title string) Menu {
	result := C.C_NSMenu_InitWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
	return MakeMenu(result)
}

func (n *NSMenu) InitWithCoder(coder foundation.Coder) Menu {
	result := C.C_NSMenu_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeMenu(result)
}

func (n *NSMenu) Init() Menu {
	result := C.C_NSMenu_Init(n.Ptr())
	return MakeMenu(result)
}

func MenuBarVisible() bool {
	result := C.C_NSMenu_MenuBarVisible()
	return bool(result)
}

func Menu_SetMenuBarVisible(visible bool) {
	C.C_NSMenu_Menu_SetMenuBarVisible(C.bool(visible))
}

func (n *NSMenu) InsertItem_AtIndex(newItem MenuItem, index int) {
	C.C_NSMenu_InsertItem_AtIndex(n.Ptr(), objc.ExtractPtr(newItem), C.int(index))
}

func (n *NSMenu) InsertItemWithTitle_Action_KeyEquivalent_AtIndex(_string string, selector *objc.Selector, charCode string, index int) MenuItem {
	result := C.C_NSMenu_InsertItemWithTitle_Action_KeyEquivalent_AtIndex(n.Ptr(), foundation.NewString(_string).Ptr(), objc.ExtractPtr(selector), foundation.NewString(charCode).Ptr(), C.int(index))
	return MakeMenuItem(result)
}

func (n *NSMenu) AddItem(newItem MenuItem) {
	C.C_NSMenu_AddItem(n.Ptr(), objc.ExtractPtr(newItem))
}

func (n *NSMenu) AddItemWithTitle_Action_KeyEquivalent(_string string, selector *objc.Selector, charCode string) MenuItem {
	result := C.C_NSMenu_AddItemWithTitle_Action_KeyEquivalent(n.Ptr(), foundation.NewString(_string).Ptr(), objc.ExtractPtr(selector), foundation.NewString(charCode).Ptr())
	return MakeMenuItem(result)
}

func (n *NSMenu) RemoveItem(item MenuItem) {
	C.C_NSMenu_RemoveItem(n.Ptr(), objc.ExtractPtr(item))
}

func (n *NSMenu) RemoveItemAtIndex(index int) {
	C.C_NSMenu_RemoveItemAtIndex(n.Ptr(), C.int(index))
}

func (n *NSMenu) ItemChanged(item MenuItem) {
	C.C_NSMenu_ItemChanged(n.Ptr(), objc.ExtractPtr(item))
}

func (n *NSMenu) RemoveAllItems() {
	C.C_NSMenu_RemoveAllItems(n.Ptr())
}

func (n *NSMenu) ItemWithTag(tag int) MenuItem {
	result := C.C_NSMenu_ItemWithTag(n.Ptr(), C.int(tag))
	return MakeMenuItem(result)
}

func (n *NSMenu) ItemWithTitle(title string) MenuItem {
	result := C.C_NSMenu_ItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
	return MakeMenuItem(result)
}

func (n *NSMenu) ItemAtIndex(index int) MenuItem {
	result := C.C_NSMenu_ItemAtIndex(n.Ptr(), C.int(index))
	return MakeMenuItem(result)
}

func (n *NSMenu) IndexOfItem(item MenuItem) int {
	result := C.C_NSMenu_IndexOfItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result)
}

func (n *NSMenu) IndexOfItemWithTitle(title string) int {
	result := C.C_NSMenu_IndexOfItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
	return int(result)
}

func (n *NSMenu) IndexOfItemWithTag(tag int) int {
	result := C.C_NSMenu_IndexOfItemWithTag(n.Ptr(), C.int(tag))
	return int(result)
}

func (n *NSMenu) IndexOfItemWithTarget_AndAction(target objc.Object, actionSelector *objc.Selector) int {
	result := C.C_NSMenu_IndexOfItemWithTarget_AndAction(n.Ptr(), objc.ExtractPtr(target), objc.ExtractPtr(actionSelector))
	return int(result)
}

func (n *NSMenu) IndexOfItemWithRepresentedObject(object objc.Object) int {
	result := C.C_NSMenu_IndexOfItemWithRepresentedObject(n.Ptr(), objc.ExtractPtr(object))
	return int(result)
}

func (n *NSMenu) IndexOfItemWithSubmenu(submenu Menu) int {
	result := C.C_NSMenu_IndexOfItemWithSubmenu(n.Ptr(), objc.ExtractPtr(submenu))
	return int(result)
}

func (n *NSMenu) SetSubmenu_ForItem(menu Menu, item MenuItem) {
	C.C_NSMenu_SetSubmenu_ForItem(n.Ptr(), objc.ExtractPtr(menu), objc.ExtractPtr(item))
}

func (n *NSMenu) SubmenuAction(sender objc.Object) {
	C.C_NSMenu_SubmenuAction(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSMenu) Update() {
	C.C_NSMenu_Update(n.Ptr())
}

func (n *NSMenu) PerformKeyEquivalent(event Event) bool {
	result := C.C_NSMenu_PerformKeyEquivalent(n.Ptr(), objc.ExtractPtr(event))
	return bool(result)
}

func (n *NSMenu) PerformActionForItemAtIndex(index int) {
	C.C_NSMenu_PerformActionForItemAtIndex(n.Ptr(), C.int(index))
}

func PopUpContextMenu_WithEvent_ForView(menu Menu, event Event, view View) {
	C.C_NSMenu_PopUpContextMenu_WithEvent_ForView(objc.ExtractPtr(menu), objc.ExtractPtr(event), objc.ExtractPtr(view))
}

func PopUpContextMenu_WithEvent_ForView_WithFont(menu Menu, event Event, view View, font Font) {
	C.C_NSMenu_PopUpContextMenu_WithEvent_ForView_WithFont(objc.ExtractPtr(menu), objc.ExtractPtr(event), objc.ExtractPtr(view), objc.ExtractPtr(font))
}

func (n *NSMenu) PopUpMenuPositioningItem_AtLocation_InView(item MenuItem, location foundation.Point, view View) bool {
	result := C.C_NSMenu_PopUpMenuPositioningItem_AtLocation_InView(n.Ptr(), objc.ExtractPtr(item), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(location))), objc.ExtractPtr(view))
	return bool(result)
}

func (n *NSMenu) CancelTracking() {
	C.C_NSMenu_CancelTracking(n.Ptr())
}

func (n *NSMenu) CancelTrackingWithoutAnimation() {
	C.C_NSMenu_CancelTrackingWithoutAnimation(n.Ptr())
}

func (n *NSMenu) MenuBarHeight() coregraphics.Float {
	result := C.C_NSMenu_MenuBarHeight(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSMenu) NumberOfItems() int {
	result := C.C_NSMenu_NumberOfItems(n.Ptr())
	return int(result)
}

func (n *NSMenu) ItemArray() []MenuItem {
	result := C.C_NSMenu_ItemArray(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]MenuItem, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = MakeMenuItem(r)
	}
	return goResult
}

func (n *NSMenu) SetItemArray(value []MenuItem) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = objc.ExtractPtr(v)
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
	C.C_NSMenu_SetItemArray(n.Ptr(), cValue)
}

func (n *NSMenu) Supermenu() Menu {
	result := C.C_NSMenu_Supermenu(n.Ptr())
	return MakeMenu(result)
}

func (n *NSMenu) SetSupermenu(value Menu) {
	C.C_NSMenu_SetSupermenu(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenu) AutoenablesItems() bool {
	result := C.C_NSMenu_AutoenablesItems(n.Ptr())
	return bool(result)
}

func (n *NSMenu) SetAutoenablesItems(value bool) {
	C.C_NSMenu_SetAutoenablesItems(n.Ptr(), C.bool(value))
}

func (n *NSMenu) Font() Font {
	result := C.C_NSMenu_Font(n.Ptr())
	return MakeFont(result)
}

func (n *NSMenu) SetFont(value Font) {
	C.C_NSMenu_SetFont(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenu) Title() string {
	result := C.C_NSMenu_Title(n.Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSMenu) SetTitle(value string) {
	C.C_NSMenu_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSMenu) MinimumWidth() coregraphics.Float {
	result := C.C_NSMenu_MinimumWidth(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSMenu) SetMinimumWidth(value coregraphics.Float) {
	C.C_NSMenu_SetMinimumWidth(n.Ptr(), C.double(float64(value)))
}

func (n *NSMenu) Size() foundation.Size {
	result := C.C_NSMenu_Size(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result)))
}

func (n *NSMenu) PropertiesToUpdate() MenuProperties {
	result := C.C_NSMenu_PropertiesToUpdate(n.Ptr())
	return MenuProperties(uint(result))
}

func (n *NSMenu) AllowsContextMenuPlugIns() bool {
	result := C.C_NSMenu_AllowsContextMenuPlugIns(n.Ptr())
	return bool(result)
}

func (n *NSMenu) SetAllowsContextMenuPlugIns(value bool) {
	C.C_NSMenu_SetAllowsContextMenuPlugIns(n.Ptr(), C.bool(value))
}

func (n *NSMenu) ShowsStateColumn() bool {
	result := C.C_NSMenu_ShowsStateColumn(n.Ptr())
	return bool(result)
}

func (n *NSMenu) SetShowsStateColumn(value bool) {
	C.C_NSMenu_SetShowsStateColumn(n.Ptr(), C.bool(value))
}

func (n *NSMenu) HighlightedItem() MenuItem {
	result := C.C_NSMenu_HighlightedItem(n.Ptr())
	return MakeMenuItem(result)
}

func (n *NSMenu) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	result := C.C_NSMenu_UserInterfaceLayoutDirection(n.Ptr())
	return UserInterfaceLayoutDirection(int(result))
}

func (n *NSMenu) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	C.C_NSMenu_SetUserInterfaceLayoutDirection(n.Ptr(), C.int(int(value)))
}
