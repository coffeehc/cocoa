package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/cocoa/uihelper"
	"unsafe"
)

type ModalSession unsafe.Pointer

func FromNSModalSessionPointer(p unsafe.Pointer) ModalSession {
	return ModalSession(*(*unsafe.Pointer)(p))
}

func ToNSModalSessionPointer(m ModalSession) unsafe.Pointer {
	return unsafe.Pointer(&m)
}

// NewPlainButton return a new common used Button
func NewPlainButton(title string) Button {
	btn := AllocButton().Init()
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	btn.SetBezelStyle(BezelStyleRounded)
	btn.SetTitle(title)
	return btn
}

// NewCheckBox return a new common used switch Button
func NewCheckBox(title string) Button {
	btn := AllocButton().Init()
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	btn.SetButtonType(ButtonTypeSwitch)
	btn.SetTitle(title)
	return btn
}

// NewRadioButton return a new common used radio Button
func NewRadioButton(title string) Button {
	btn := AllocButton().Init()
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	btn.SetButtonType(ButtonTypeRadio)
	btn.SetTitle(title)
	return btn
}

// NewLabel create a text field, which looks like a Label
func NewLabel(title string) TextField {
	tf := AllocTextField().Init()
	tf.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tf.SetBezeled(false)
	tf.SetDrawsBackground(false)
	tf.SetEditable(false)
	tf.SetSelectable(false)
	tf.SetStringValue(title)
	return tf
}

// NewTextField return a plain TextField
func NewTextField() TextField {
	field := AllocTextField().Init()
	field.SetTranslatesAutoresizingMaskIntoConstraints(false)
	field.SetUsesSingleLineMode(true)
	cell := field.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return field
}

// NewSecureTextField return a plain SecureTextField
func NewSecureTextField() SecureTextField {
	field := AllocSecureTextField().Init()
	field.SetTranslatesAutoresizingMaskIntoConstraints(false)
	field.SetUsesSingleLineMode(true)
	cell := field.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return field
}

// NewWindow create a common window with close/minimize buttons
func NewWindow(width, height float64) Window {
	return AllocWindow().InitWithContentRect_StyleMask_Backing_Defer(
		foundation.MakeRect(0, 0, width, height),
		WindowStyleMaskTitled|WindowStyleMaskClosable|WindowStyleMaskResizable|WindowStyleMaskMiniaturizable,
		BackingStoreBuffered,
		false,
	)
}

// NewWindowWithStyle create a common window with styles
func NewWindowWithStyle(width, height float64, style WindowStyleMask) Window {
	return AllocWindow().InitWithContentRect_StyleMask_Backing_Defer(
		foundation.MakeRect(0, 0, width, height),
		style,
		BackingStoreBuffered,
		false,
	)
}

// NewMenuItem create a new menu item, with selector
func NewMenuItem(title string, charCode string, selector objc.Selector) MenuItem {
	return AllocMenuItem().InitWithTitle_Action_KeyEquivalent(title, selector, charCode)
}

// NewMenuItemWithAction create a new menu item with action
func NewMenuItemWithAction(title string, charCode string, handler uihelper.ActionHandler) MenuItem {
	item := AllocMenuItem().InitWithTitle_Action_KeyEquivalent(title, nil, charCode)
	uihelper.SetAction(item, handler)
	return item
}

// NewSubMenuItem create a menu item that hold a sub menu
func NewSubMenuItem(menu Menu) MenuItem {
	item := AllocMenuItem().InitWithTitle_Action_KeyEquivalent("", nil, "")
	item.SetMenu(menu)
	return item
}

// A TextScrollView that contains a TextView
type TextScrollView struct {
	ScrollView
	textView TextView
}

// TextView return the inner TextView
func (t TextScrollView) TextView() TextView {
	return t.textView
}

// NewScrollableTextView create and return new scrollable text view.
func NewScrollableTextView() TextScrollView {
	stv := ScrollableTextView()
	stv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tv := MakeTextView(stv.DocumentView().Ptr())
	tv.SetAllowsUndo(true)
	return TextScrollView{
		ScrollView: stv,
		textView:   tv,
	}
}

// NewVerticalStackView return a new vertical StackView
func NewVerticalStackView() StackView {
	sv := AllocStackView().Init()
	sv.SetOrientation(UserInterfaceLayoutOrientationVertical)
	sv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return sv
}

// NewHorizontalStackView return a new horizontal StackView
func NewHorizontalStackView() StackView {
	sv := AllocStackView().Init()
	sv.SetOrientation(UserInterfaceLayoutOrientationHorizontal)
	sv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return sv
}
