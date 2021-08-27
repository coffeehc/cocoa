package appkit

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
	InsertItemWithTitle_Action_KeyEquivalent_AtIndex(_string string, selector objc.Selector, charCode string, index int) MenuItem
	AddItem(newItem MenuItem)
	AddItemWithTitle_Action_KeyEquivalent(_string string, selector objc.Selector, charCode string) MenuItem
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
	IndexOfItemWithTarget_AndAction(target objc.Object, actionSelector objc.Selector) int
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
	Delegate() objc.Object
	SetDelegate(value objc.Object)
}

type NSMenu struct {
	objc.NSObject
}

func MakeMenu(ptr unsafe.Pointer) NSMenu {
	return NSMenu{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocMenu() NSMenu {
	return MakeMenu(C.C_Menu_Alloc())
}

func (n NSMenu) InitWithTitle(title string) Menu {
	result_ := C.C_NSMenu_InitWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
	return MakeMenu(result_)
}

func (n NSMenu) InitWithCoder(coder foundation.Coder) Menu {
	result_ := C.C_NSMenu_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeMenu(result_)
}

func (n NSMenu) Init() Menu {
	result_ := C.C_NSMenu_Init(n.Ptr())
	return MakeMenu(result_)
}

func MenuBarVisible() bool {
	result_ := C.C_NSMenu_MenuBarVisible()
	return bool(result_)
}

func SetMenuBarVisible(visible bool) {
	C.C_NSMenu_SetMenuBarVisible(C.bool(visible))
}

func (n NSMenu) InsertItem_AtIndex(newItem MenuItem, index int) {
	C.C_NSMenu_InsertItem_AtIndex(n.Ptr(), objc.ExtractPtr(newItem), C.int(index))
}

func (n NSMenu) InsertItemWithTitle_Action_KeyEquivalent_AtIndex(_string string, selector objc.Selector, charCode string, index int) MenuItem {
	result_ := C.C_NSMenu_InsertItemWithTitle_Action_KeyEquivalent_AtIndex(n.Ptr(), foundation.NewString(_string).Ptr(), unsafe.Pointer(selector), foundation.NewString(charCode).Ptr(), C.int(index))
	return MakeMenuItem(result_)
}

func (n NSMenu) AddItem(newItem MenuItem) {
	C.C_NSMenu_AddItem(n.Ptr(), objc.ExtractPtr(newItem))
}

func (n NSMenu) AddItemWithTitle_Action_KeyEquivalent(_string string, selector objc.Selector, charCode string) MenuItem {
	result_ := C.C_NSMenu_AddItemWithTitle_Action_KeyEquivalent(n.Ptr(), foundation.NewString(_string).Ptr(), unsafe.Pointer(selector), foundation.NewString(charCode).Ptr())
	return MakeMenuItem(result_)
}

func (n NSMenu) RemoveItem(item MenuItem) {
	C.C_NSMenu_RemoveItem(n.Ptr(), objc.ExtractPtr(item))
}

func (n NSMenu) RemoveItemAtIndex(index int) {
	C.C_NSMenu_RemoveItemAtIndex(n.Ptr(), C.int(index))
}

func (n NSMenu) ItemChanged(item MenuItem) {
	C.C_NSMenu_ItemChanged(n.Ptr(), objc.ExtractPtr(item))
}

func (n NSMenu) RemoveAllItems() {
	C.C_NSMenu_RemoveAllItems(n.Ptr())
}

func (n NSMenu) ItemWithTag(tag int) MenuItem {
	result_ := C.C_NSMenu_ItemWithTag(n.Ptr(), C.int(tag))
	return MakeMenuItem(result_)
}

func (n NSMenu) ItemWithTitle(title string) MenuItem {
	result_ := C.C_NSMenu_ItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
	return MakeMenuItem(result_)
}

func (n NSMenu) ItemAtIndex(index int) MenuItem {
	result_ := C.C_NSMenu_ItemAtIndex(n.Ptr(), C.int(index))
	return MakeMenuItem(result_)
}

func (n NSMenu) IndexOfItem(item MenuItem) int {
	result_ := C.C_NSMenu_IndexOfItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result_)
}

func (n NSMenu) IndexOfItemWithTitle(title string) int {
	result_ := C.C_NSMenu_IndexOfItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
	return int(result_)
}

func (n NSMenu) IndexOfItemWithTag(tag int) int {
	result_ := C.C_NSMenu_IndexOfItemWithTag(n.Ptr(), C.int(tag))
	return int(result_)
}

func (n NSMenu) IndexOfItemWithTarget_AndAction(target objc.Object, actionSelector objc.Selector) int {
	result_ := C.C_NSMenu_IndexOfItemWithTarget_AndAction(n.Ptr(), objc.ExtractPtr(target), unsafe.Pointer(actionSelector))
	return int(result_)
}

func (n NSMenu) IndexOfItemWithRepresentedObject(object objc.Object) int {
	result_ := C.C_NSMenu_IndexOfItemWithRepresentedObject(n.Ptr(), objc.ExtractPtr(object))
	return int(result_)
}

func (n NSMenu) IndexOfItemWithSubmenu(submenu Menu) int {
	result_ := C.C_NSMenu_IndexOfItemWithSubmenu(n.Ptr(), objc.ExtractPtr(submenu))
	return int(result_)
}

func (n NSMenu) SetSubmenu_ForItem(menu Menu, item MenuItem) {
	C.C_NSMenu_SetSubmenu_ForItem(n.Ptr(), objc.ExtractPtr(menu), objc.ExtractPtr(item))
}

func (n NSMenu) SubmenuAction(sender objc.Object) {
	C.C_NSMenu_SubmenuAction(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSMenu) Update() {
	C.C_NSMenu_Update(n.Ptr())
}

func (n NSMenu) PerformKeyEquivalent(event Event) bool {
	result_ := C.C_NSMenu_PerformKeyEquivalent(n.Ptr(), objc.ExtractPtr(event))
	return bool(result_)
}

func (n NSMenu) PerformActionForItemAtIndex(index int) {
	C.C_NSMenu_PerformActionForItemAtIndex(n.Ptr(), C.int(index))
}

func PopUpContextMenu_WithEvent_ForView(menu Menu, event Event, view View) {
	C.C_NSMenu_PopUpContextMenu_WithEvent_ForView(objc.ExtractPtr(menu), objc.ExtractPtr(event), objc.ExtractPtr(view))
}

func PopUpContextMenu_WithEvent_ForView_WithFont(menu Menu, event Event, view View, font Font) {
	C.C_NSMenu_PopUpContextMenu_WithEvent_ForView_WithFont(objc.ExtractPtr(menu), objc.ExtractPtr(event), objc.ExtractPtr(view), objc.ExtractPtr(font))
}

func (n NSMenu) PopUpMenuPositioningItem_AtLocation_InView(item MenuItem, location foundation.Point, view View) bool {
	result_ := C.C_NSMenu_PopUpMenuPositioningItem_AtLocation_InView(n.Ptr(), objc.ExtractPtr(item), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(location))), objc.ExtractPtr(view))
	return bool(result_)
}

func (n NSMenu) CancelTracking() {
	C.C_NSMenu_CancelTracking(n.Ptr())
}

func (n NSMenu) CancelTrackingWithoutAnimation() {
	C.C_NSMenu_CancelTrackingWithoutAnimation(n.Ptr())
}

func (n NSMenu) MenuBarHeight() coregraphics.Float {
	result_ := C.C_NSMenu_MenuBarHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSMenu) NumberOfItems() int {
	result_ := C.C_NSMenu_NumberOfItems(n.Ptr())
	return int(result_)
}

func (n NSMenu) ItemArray() []MenuItem {
	result_ := C.C_NSMenu_ItemArray(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]MenuItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeMenuItem(r)
	}
	return goResult_
}

func (n NSMenu) SetItemArray(value []MenuItem) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSMenu_SetItemArray(n.Ptr(), cValue)
}

func (n NSMenu) Supermenu() Menu {
	result_ := C.C_NSMenu_Supermenu(n.Ptr())
	return MakeMenu(result_)
}

func (n NSMenu) SetSupermenu(value Menu) {
	C.C_NSMenu_SetSupermenu(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSMenu) AutoenablesItems() bool {
	result_ := C.C_NSMenu_AutoenablesItems(n.Ptr())
	return bool(result_)
}

func (n NSMenu) SetAutoenablesItems(value bool) {
	C.C_NSMenu_SetAutoenablesItems(n.Ptr(), C.bool(value))
}

func (n NSMenu) Font() Font {
	result_ := C.C_NSMenu_Font(n.Ptr())
	return MakeFont(result_)
}

func (n NSMenu) SetFont(value Font) {
	C.C_NSMenu_SetFont(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSMenu) Title() string {
	result_ := C.C_NSMenu_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSMenu) SetTitle(value string) {
	C.C_NSMenu_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSMenu) MinimumWidth() coregraphics.Float {
	result_ := C.C_NSMenu_MinimumWidth(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSMenu) SetMinimumWidth(value coregraphics.Float) {
	C.C_NSMenu_SetMinimumWidth(n.Ptr(), C.double(float64(value)))
}

func (n NSMenu) Size() foundation.Size {
	result_ := C.C_NSMenu_Size(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSMenu) PropertiesToUpdate() MenuProperties {
	result_ := C.C_NSMenu_PropertiesToUpdate(n.Ptr())
	return MenuProperties(uint(result_))
}

func (n NSMenu) AllowsContextMenuPlugIns() bool {
	result_ := C.C_NSMenu_AllowsContextMenuPlugIns(n.Ptr())
	return bool(result_)
}

func (n NSMenu) SetAllowsContextMenuPlugIns(value bool) {
	C.C_NSMenu_SetAllowsContextMenuPlugIns(n.Ptr(), C.bool(value))
}

func (n NSMenu) ShowsStateColumn() bool {
	result_ := C.C_NSMenu_ShowsStateColumn(n.Ptr())
	return bool(result_)
}

func (n NSMenu) SetShowsStateColumn(value bool) {
	C.C_NSMenu_SetShowsStateColumn(n.Ptr(), C.bool(value))
}

func (n NSMenu) HighlightedItem() MenuItem {
	result_ := C.C_NSMenu_HighlightedItem(n.Ptr())
	return MakeMenuItem(result_)
}

func (n NSMenu) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	result_ := C.C_NSMenu_UserInterfaceLayoutDirection(n.Ptr())
	return UserInterfaceLayoutDirection(int(result_))
}

func (n NSMenu) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	C.C_NSMenu_SetUserInterfaceLayoutDirection(n.Ptr(), C.int(int(value)))
}

func (n NSMenu) Delegate() objc.Object {
	result_ := C.C_NSMenu_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSMenu) SetDelegate(value objc.Object) {
	C.C_NSMenu_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}
