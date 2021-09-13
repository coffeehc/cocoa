package appkits

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/helper/actions"
	"github.com/hsiafan/cocoa/objc"
)

// NewPlainButton return a new common used Button
func NewPlainButton(title string) appkit.NSButton {
	btn := appkit.NewButton().Autorelease()
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	btn.SetBezelStyle(appkit.BezelStyleRounded)
	btn.SetTitle(title)
	return btn
}

// NewCheckBox return a new common used switch Button
func NewCheckBox(title string) appkit.NSButton {
	btn := appkit.NewButton().Autorelease()
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	btn.SetButtonType(appkit.ButtonTypeSwitch)
	btn.SetTitle(title)
	return btn
}

// NewRadioButton return a new common used radio Button
func NewRadioButton(title string) appkit.NSButton {
	btn := appkit.NewButton().Autorelease()
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	btn.SetButtonType(appkit.ButtonTypeRadio)
	btn.SetTitle(title)
	return btn
}

// NewLabel create a text field, which looks like a Label
func NewLabel(title string) appkit.NSTextField {
	tf := appkit.NewTextField().Autorelease()
	tf.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tf.SetBezeled(false)
	tf.SetDrawsBackground(false)
	tf.SetEditable(false)
	tf.SetSelectable(false)
	tf.SetStringValue(title)
	return tf
}

// NewTextField return a plain TextField
func NewTextField() appkit.NSTextField {
	tf := appkit.NewTextField().Autorelease()
	tf.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tf.SetUsesSingleLineMode(true)
	cell := tf.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return tf
}

// NewSecureTextField return a plain SecureTextField
func NewSecureTextField() appkit.NSSecureTextField {
	stf := appkit.NewSecureTextField().Autorelease()
	stf.SetTranslatesAutoresizingMaskIntoConstraints(false)
	stf.SetUsesSingleLineMode(true)
	cell := stf.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return stf
}

// NewWindow create a common window with close/minimize buttons
func NewWindow(width, height float64) appkit.NSWindow {
	return appkit.AllocWindow().InitWithContentRect_StyleMask_Backing_Defer(
		foundation.MakeRect(0, 0, width, height),
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable|appkit.WindowStyleMaskResizable|appkit.WindowStyleMaskMiniaturizable,
		appkit.BackingStoreBuffered,
		false,
	).Autorelease()
}

// NewWindowWithStyle create a common window with styles
func NewWindowWithStyle(width, height float64, style appkit.WindowStyleMask) appkit.NSWindow {
	return appkit.AllocWindow().InitWithContentRect_StyleMask_Backing_Defer(
		foundation.MakeRect(0, 0, width, height),
		style,
		appkit.BackingStoreBuffered,
		false,
	).Autorelease()
}

// NewMenuItem create a new menu item, with selector
func NewMenuItem(title string, charCode string, selector objc.Selector) appkit.NSMenuItem {
	return appkit.AllocMenuItem().InitWithTitle_Action_KeyEquivalent(title, selector, charCode).Autorelease()
}

// NewMenuItemWithAction create a new menu item with action
func NewMenuItemWithAction(title string, charCode string, handler actions.ActionHandler) appkit.NSMenuItem {
	item := appkit.AllocMenuItem().InitWithTitle_Action_KeyEquivalent(title, nil, charCode).Autorelease()
	actions.Set(item, handler)
	return item
}

// NewSubMenuItem create a menu item that hold a sub menu
func NewSubMenuItem(menu appkit.Menu) appkit.NSMenuItem {
	item := appkit.AllocMenuItem().InitWithTitle_Action_KeyEquivalent("", nil, "").Autorelease()
	item.SetMenu(menu)
	return item
}

// NewView create new View
func NewView() appkit.NSView {
	v := appkit.NewView().Autorelease()
	v.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return v
}

// TextScrollView is appkit.ScrollView that contains a TextView
type TextScrollView struct {
	appkit.NSScrollView
	textView appkit.NSTextView
}

// TextView return the inner TextView
func (t *TextScrollView) TextView() appkit.NSTextView {
	return t.textView
}

// NewScrollableTextView create and return new scrollable text view.
func NewScrollableTextView() *TextScrollView {
	stv := appkit.ScrollableTextView()
	stv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tv := appkit.MakeTextView(stv.DocumentView().Ptr())
	tv.SetAllowsUndo(true)
	return &TextScrollView{
		NSScrollView: stv.(appkit.NSScrollView),
		textView:     tv,
	}
}

// NewVerticalStackView return a new vertical StackView
func NewVerticalStackView() appkit.NSStackView {
	sv := appkit.NewStackView().Autorelease()
	sv.SetOrientation(appkit.UserInterfaceLayoutOrientationVertical)
	sv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return sv
}

// NewHorizontalStackView return a new horizontal StackView
func NewHorizontalStackView() appkit.NSStackView {
	sv := appkit.NewStackView().Autorelease()
	sv.SetOrientation(appkit.UserInterfaceLayoutOrientationHorizontal)
	sv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return sv
}
