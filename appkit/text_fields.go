package appkit

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
