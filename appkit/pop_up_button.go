package appkit

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
	IndexOfItemWithTarget_AndAction(target objc.Object, actionSelector objc.Selector) int
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

func MakePopUpButton(ptr unsafe.Pointer) NSPopUpButton {
	return NSPopUpButton{
		NSButton: MakeButton(ptr),
	}
}

func (n NSPopUpButton) InitWithFrame_PullsDown(buttonFrame foundation.Rect, flag bool) NSPopUpButton {
	result_ := C.C_NSPopUpButton_InitWithFrame_PullsDown(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(buttonFrame))), C.bool(flag))
	return MakePopUpButton(result_)
}

func PopUpButton_CheckboxWithTitle_Target_Action(title string, target objc.Object, action objc.Selector) NSPopUpButton {
	result_ := C.C_NSPopUpButton_PopUpButton_CheckboxWithTitle_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakePopUpButton(result_)
}

func PopUpButton_ButtonWithImage_Target_Action(image Image, target objc.Object, action objc.Selector) NSPopUpButton {
	result_ := C.C_NSPopUpButton_PopUpButton_ButtonWithImage_Target_Action(objc.ExtractPtr(image), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakePopUpButton(result_)
}

func PopUpButton_RadioButtonWithTitle_Target_Action(title string, target objc.Object, action objc.Selector) NSPopUpButton {
	result_ := C.C_NSPopUpButton_PopUpButton_RadioButtonWithTitle_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakePopUpButton(result_)
}

func PopUpButton_ButtonWithTitle_Image_Target_Action(title string, image Image, target objc.Object, action objc.Selector) NSPopUpButton {
	result_ := C.C_NSPopUpButton_PopUpButton_ButtonWithTitle_Image_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(image), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakePopUpButton(result_)
}

func PopUpButton_ButtonWithTitle_Target_Action(title string, target objc.Object, action objc.Selector) NSPopUpButton {
	result_ := C.C_NSPopUpButton_PopUpButton_ButtonWithTitle_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakePopUpButton(result_)
}

func (n NSPopUpButton) InitWithFrame(frameRect foundation.Rect) NSPopUpButton {
	result_ := C.C_NSPopUpButton_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakePopUpButton(result_)
}

func (n NSPopUpButton) InitWithCoder(coder foundation.Coder) NSPopUpButton {
	result_ := C.C_NSPopUpButton_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakePopUpButton(result_)
}

func (n NSPopUpButton) Init() NSPopUpButton {
	result_ := C.C_NSPopUpButton_Init(n.Ptr())
	return MakePopUpButton(result_)
}

func AllocPopUpButton() NSPopUpButton {
	result_ := C.C_NSPopUpButton_AllocPopUpButton()
	return MakePopUpButton(result_)
}

func NewPopUpButton() NSPopUpButton {
	result_ := C.C_NSPopUpButton_NewPopUpButton()
	return MakePopUpButton(result_)
}

func (n NSPopUpButton) Autorelease() NSPopUpButton {
	result_ := C.C_NSPopUpButton_Autorelease(n.Ptr())
	return MakePopUpButton(result_)
}

func (n NSPopUpButton) Retain() NSPopUpButton {
	result_ := C.C_NSPopUpButton_Retain(n.Ptr())
	return MakePopUpButton(result_)
}

func (n NSPopUpButton) AddItemWithTitle(title string) {
	C.C_NSPopUpButton_AddItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
}

func (n NSPopUpButton) AddItemsWithTitles(itemTitles []string) {
	var cItemTitles C.Array
	if len(itemTitles) > 0 {
		cItemTitlesData := make([]unsafe.Pointer, len(itemTitles))
		for idx, v := range itemTitles {
			cItemTitlesData[idx] = foundation.NewString(v).Ptr()
		}
		cItemTitles.data = unsafe.Pointer(&cItemTitlesData[0])
		cItemTitles.len = C.int(len(itemTitles))
	}
	C.C_NSPopUpButton_AddItemsWithTitles(n.Ptr(), cItemTitles)
}

func (n NSPopUpButton) InsertItemWithTitle_AtIndex(title string, index int) {
	C.C_NSPopUpButton_InsertItemWithTitle_AtIndex(n.Ptr(), foundation.NewString(title).Ptr(), C.int(index))
}

func (n NSPopUpButton) RemoveAllItems() {
	C.C_NSPopUpButton_RemoveAllItems(n.Ptr())
}

func (n NSPopUpButton) RemoveItemWithTitle(title string) {
	C.C_NSPopUpButton_RemoveItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
}

func (n NSPopUpButton) RemoveItemAtIndex(index int) {
	C.C_NSPopUpButton_RemoveItemAtIndex(n.Ptr(), C.int(index))
}

func (n NSPopUpButton) SelectItem(item MenuItem) {
	C.C_NSPopUpButton_SelectItem(n.Ptr(), objc.ExtractPtr(item))
}

func (n NSPopUpButton) SelectItemAtIndex(index int) {
	C.C_NSPopUpButton_SelectItemAtIndex(n.Ptr(), C.int(index))
}

func (n NSPopUpButton) SelectItemWithTag(tag int) bool {
	result_ := C.C_NSPopUpButton_SelectItemWithTag(n.Ptr(), C.int(tag))
	return bool(result_)
}

func (n NSPopUpButton) SelectItemWithTitle(title string) {
	C.C_NSPopUpButton_SelectItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
}

func (n NSPopUpButton) ItemAtIndex(index int) MenuItem {
	result_ := C.C_NSPopUpButton_ItemAtIndex(n.Ptr(), C.int(index))
	return MakeMenuItem(result_)
}

func (n NSPopUpButton) ItemTitleAtIndex(index int) string {
	result_ := C.C_NSPopUpButton_ItemTitleAtIndex(n.Ptr(), C.int(index))
	return foundation.MakeString(result_).String()
}

func (n NSPopUpButton) ItemWithTitle(title string) MenuItem {
	result_ := C.C_NSPopUpButton_ItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
	return MakeMenuItem(result_)
}

func (n NSPopUpButton) IndexOfItem(item MenuItem) int {
	result_ := C.C_NSPopUpButton_IndexOfItem(n.Ptr(), objc.ExtractPtr(item))
	return int(result_)
}

func (n NSPopUpButton) IndexOfItemWithTag(tag int) int {
	result_ := C.C_NSPopUpButton_IndexOfItemWithTag(n.Ptr(), C.int(tag))
	return int(result_)
}

func (n NSPopUpButton) IndexOfItemWithTitle(title string) int {
	result_ := C.C_NSPopUpButton_IndexOfItemWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
	return int(result_)
}

func (n NSPopUpButton) IndexOfItemWithRepresentedObject(obj objc.Object) int {
	result_ := C.C_NSPopUpButton_IndexOfItemWithRepresentedObject(n.Ptr(), objc.ExtractPtr(obj))
	return int(result_)
}

func (n NSPopUpButton) IndexOfItemWithTarget_AndAction(target objc.Object, actionSelector objc.Selector) int {
	result_ := C.C_NSPopUpButton_IndexOfItemWithTarget_AndAction(n.Ptr(), objc.ExtractPtr(target), unsafe.Pointer(actionSelector))
	return int(result_)
}

func (n NSPopUpButton) SynchronizeTitleAndSelectedItem() {
	C.C_NSPopUpButton_SynchronizeTitleAndSelectedItem(n.Ptr())
}

func (n NSPopUpButton) PullsDown() bool {
	result_ := C.C_NSPopUpButton_PullsDown(n.Ptr())
	return bool(result_)
}

func (n NSPopUpButton) SetPullsDown(value bool) {
	C.C_NSPopUpButton_SetPullsDown(n.Ptr(), C.bool(value))
}

func (n NSPopUpButton) AutoenablesItems() bool {
	result_ := C.C_NSPopUpButton_AutoenablesItems(n.Ptr())
	return bool(result_)
}

func (n NSPopUpButton) SetAutoenablesItems(value bool) {
	C.C_NSPopUpButton_SetAutoenablesItems(n.Ptr(), C.bool(value))
}

func (n NSPopUpButton) SelectedItem() MenuItem {
	result_ := C.C_NSPopUpButton_SelectedItem(n.Ptr())
	return MakeMenuItem(result_)
}

func (n NSPopUpButton) TitleOfSelectedItem() string {
	result_ := C.C_NSPopUpButton_TitleOfSelectedItem(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSPopUpButton) IndexOfSelectedItem() int {
	result_ := C.C_NSPopUpButton_IndexOfSelectedItem(n.Ptr())
	return int(result_)
}

func (n NSPopUpButton) SelectedTag() int {
	result_ := C.C_NSPopUpButton_SelectedTag(n.Ptr())
	return int(result_)
}

func (n NSPopUpButton) NumberOfItems() int {
	result_ := C.C_NSPopUpButton_NumberOfItems(n.Ptr())
	return int(result_)
}

func (n NSPopUpButton) ItemArray() []MenuItem {
	result_ := C.C_NSPopUpButton_ItemArray(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]MenuItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeMenuItem(r)
	}
	return goResult_
}

func (n NSPopUpButton) ItemTitles() []string {
	result_ := C.C_NSPopUpButton_ItemTitles(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSPopUpButton) LastItem() MenuItem {
	result_ := C.C_NSPopUpButton_LastItem(n.Ptr())
	return MakeMenuItem(result_)
}

func (n NSPopUpButton) PreferredEdge() foundation.RectEdge {
	result_ := C.C_NSPopUpButton_PreferredEdge(n.Ptr())
	return foundation.RectEdge(uint(result_))
}

func (n NSPopUpButton) SetPreferredEdge(value foundation.RectEdge) {
	C.C_NSPopUpButton_SetPreferredEdge(n.Ptr(), C.uint(uint(value)))
}
