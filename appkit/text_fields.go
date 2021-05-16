package appkit

import "github.com/hsiafan/cocoa/foundation"

// NewLabel create a text field, which looks like a Label
func NewLabel(frame foundation.Rect) TextField {
	tf := AllocTextField().InitWithFrame(frame)
	tf.SetBezeled(false)
	tf.SetDrawsBackground(false)
	tf.SetEditable(false)
	tf.SetSelectable(false)
	return tf
}

// NewTextField return a plain TextField
func NewTextField(frame foundation.Rect) TextField {
	field := AllocTextField().InitWithFrame(frame)
	cell := field.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return field
}

// NewSecureTextField return a plain SecureTextField
func NewSecureTextField(frame foundation.Rect) SecureTextField {
	field := AllocSecureTextField().InitWithFrame(frame)
	cell := field.Cell()
	cell.SetScrollable(true)
	cell.SetWraps(false)
	return field
}
