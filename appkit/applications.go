package appkit

import "unsafe"

type ApplicationActivationPolicy int

const (
	ApplicationActivationPolicyRegular ApplicationActivationPolicy = iota
	ApplicationActivationPolicyAccessory
	ApplicationActivationPolicyProhibited
)

type ModalResponse int

const (
	ModalResponseStop       ModalResponse = -1000
	ModalResponseAbort      ModalResponse = -1001
	ModalResponseContinue   ModalResponse = -1002
	ModalResponseOK         ModalResponse = 1
	ModalResponseCancel     ModalResponse = 0
	AlertFirstButtonReturn  ModalResponse = 1000
	AlertSecondButtonReturn ModalResponse = 1001
	AlertThirdButtonReturn  ModalResponse = 1002
)

type ModalSession struct {
}

func FromNSModalSessionPointer(p unsafe.Pointer) ModalSession {
	panic("to be implemented")
}

func ToNSModalSessionPointer(m ModalSession) unsafe.Pointer {
	panic("to be implemented")
}
