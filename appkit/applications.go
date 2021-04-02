package appkit

// InitSharedApplication initializes and return the global application instance.
func InitSharedApplication() Application {
	app := SharedApplication()
	app.(*NSApplication).setDelegate()
	return app
}

/* The following activation policies control whether and how an application may be activated.  They are determined by the Info.plist. */
type ApplicationActivationPolicy int

const (
	/* The application is an ordinary app that appears in the Dock and may have a user interface.  This is the default for bundled apps, unless overridden in the Info.plist. */
	ApplicationActivationPolicyRegular ApplicationActivationPolicy = iota

	/* The application does not appear in the Dock and does not have a menu bar, but it may be activated programmatically or by clicking on one of its windows.  This corresponds to LSUIElement=1 in the Info.plist. */
	ApplicationActivationPolicyAccessory

	/* The application does not appear in the Dock and may not create windows or be activated.  This corresponds to LSBackgroundOnly=1 in the Info.plist.  This is also the default for unbundled executables that do not have Info.plists. */
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
