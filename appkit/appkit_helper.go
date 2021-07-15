package appkit

import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type ModalSession struct {
}

func FromNSModalSessionPointer(p unsafe.Pointer) ModalSession {
	panic("to be implemented")
}

func ToNSModalSessionPointer(m ModalSession) unsafe.Pointer {
	panic("to be implemented")
}

// NewPlainButton return a new common used Button
func NewPlainButton() Button {
	btn := AllocButton().Init()
	btn.SetBezelStyle(BezelStyleRounded)
	return btn
}

// NewCheckBox return a new common used switch Button
func NewCheckBox() Button {
	btn := AllocButton().Init()
	btn.SetButtonType(ButtonTypeSwitch)
	return btn
}

// NewRadioButton return a new common used radio Button
func NewRadioButton() Button {
	btn := AllocButton().Init()
	btn.SetButtonType(ButtonTypeRadio)
	return btn
}

// NewLabel create a text field, which looks like a Label
func NewLabel() TextField {
	tf := AllocTextField().Init()
	tf.SetBezeled(false)
	tf.SetDrawsBackground(false)
	tf.SetEditable(false)
	tf.SetSelectable(false)
	return tf
}

// NewTextField return a plain TextField
func NewTextField() TextField {
	field := AllocTextField().Init()
	cell := field.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return field
}

// NewSecureTextField return a plain SecureTextField
func NewSecureTextField() SecureTextField {
	field := AllocSecureTextField().Init()
	cell := field.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return field
}

// NewWindow create a common window with close/minimize buttons
func NewWindow(rect foundation.Rect) Window {
	return AllocWindow().InitWithContentRect_StyleMask_Backing_Defer(
		rect,
		WindowStyleMaskTitled|WindowStyleMaskClosable|WindowStyleMaskResizable|WindowStyleMaskMiniaturizable,
		BackingStoreBuffered,
		false,
	)
}
