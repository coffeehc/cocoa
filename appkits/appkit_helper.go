package appkits

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/helper/actions"
	"github.com/hsiafan/cocoa/objc"
)

// NewPlainButton return a new common used Button
func NewPlainButton(title string) appkit.Button {
	btn := appkit.AllocButton().Init()
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	btn.SetBezelStyle(appkit.BezelStyleRounded)
	btn.SetTitle(title)
	return btn
}

// NewCheckBox return a new common used switch Button
func NewCheckBox(title string) appkit.Button {
	btn := appkit.AllocButton().Init()
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	btn.SetButtonType(appkit.ButtonTypeSwitch)
	btn.SetTitle(title)
	return btn
}

// NewRadioButton return a new common used radio Button
func NewRadioButton(title string) appkit.Button {
	btn := appkit.AllocButton().Init()
	btn.SetTranslatesAutoresizingMaskIntoConstraints(false)
	btn.SetButtonType(appkit.ButtonTypeRadio)
	btn.SetTitle(title)
	return btn
}

// NewLabel create a text field, which looks like a Label
func NewLabel(title string) appkit.TextField {
	tf := appkit.AllocTextField().Init()
	tf.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tf.SetBezeled(false)
	tf.SetDrawsBackground(false)
	tf.SetEditable(false)
	tf.SetSelectable(false)
	tf.SetStringValue(title)
	return tf
}

// NewTextField return a plain TextField
func NewTextField() appkit.TextField {
	field := appkit.AllocTextField().Init()
	field.SetTranslatesAutoresizingMaskIntoConstraints(false)
	field.SetUsesSingleLineMode(true)
	cell := field.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return field
}

// NewSecureTextField return a plain SecureTextField
func NewSecureTextField() appkit.SecureTextField {
	field := appkit.AllocSecureTextField().Init()
	field.SetTranslatesAutoresizingMaskIntoConstraints(false)
	field.SetUsesSingleLineMode(true)
	cell := field.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return field
}

// NewWindow create a common window with close/minimize buttons
func NewWindow(width, height float64) appkit.Window {
	return appkit.AllocWindow().InitWithContentRect_StyleMask_Backing_Defer(
		foundation.MakeRect(0, 0, width, height),
		appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable|appkit.WindowStyleMaskResizable|appkit.WindowStyleMaskMiniaturizable,
		appkit.BackingStoreBuffered,
		false,
	)
}

// NewWindowWithStyle create a common window with styles
func NewWindowWithStyle(width, height float64, style appkit.WindowStyleMask) appkit.Window {
	return appkit.AllocWindow().InitWithContentRect_StyleMask_Backing_Defer(
		foundation.MakeRect(0, 0, width, height),
		style,
		appkit.BackingStoreBuffered,
		false,
	)
}

// NewMenuItem create a new menu item, with selector
func NewMenuItem(title string, charCode string, selector objc.Selector) appkit.MenuItem {
	return appkit.AllocMenuItem().InitWithTitle_Action_KeyEquivalent(title, selector, charCode)
}

// NewMenuItemWithAction create a new menu item with action
func NewMenuItemWithAction(title string, charCode string, handler actions.ActionHandler) appkit.MenuItem {
	item := appkit.AllocMenuItem().InitWithTitle_Action_KeyEquivalent(title, nil, charCode)
	actions.Set(item, handler)
	return item
}

// NewSubMenuItem create a menu item that hold a sub menu
func NewSubMenuItem(menu appkit.Menu) appkit.MenuItem {
	item := appkit.AllocMenuItem().InitWithTitle_Action_KeyEquivalent("", nil, "")
	item.SetMenu(menu)
	return item
}

// NewView create new View
func NewView() appkit.View {
	v := appkit.AllocView().Init()
	v.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return v
}

// A TextScrollView that contains a TextView
type TextScrollView struct {
	appkit.ScrollView
	textView appkit.TextView
}

// TextView return the inner TextView
func (t TextScrollView) TextView() appkit.TextView {
	return t.textView
}

// NewScrollableTextView create and return new scrollable text view.
func NewScrollableTextView() TextScrollView {
	stv := appkit.ScrollableTextView()
	stv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tv := appkit.MakeTextView(stv.DocumentView().Ptr())
	tv.SetAllowsUndo(true)
	return TextScrollView{
		ScrollView: stv,
		textView:   tv,
	}
}

// NewVerticalStackView return a new vertical StackView
func NewVerticalStackView() appkit.StackView {
	sv := appkit.AllocStackView().Init()
	sv.SetOrientation(appkit.UserInterfaceLayoutOrientationVertical)
	sv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return sv
}

// NewHorizontalStackView return a new horizontal StackView
func NewHorizontalStackView() appkit.StackView {
	sv := appkit.AllocStackView().Init()
	sv.SetOrientation(appkit.UserInterfaceLayoutOrientationHorizontal)
	sv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	return sv
}
