package appkits

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/helper/actions"
	"github.com/hsiafan/cocoa/objc"
)

// Dialog is appkit.Panel with optional OK|CANCEL buttons
type Dialog interface {
	appkit.Panel
	// SetView set the view for Dialog
	SetView(view appkit.View)
	// Show display dialog in non-modal mode
	Show(handle func())
	// RunModal display dialog in modal mode
	RunModal() appkit.ModalResponse
}

// GoDialog implements Dialog
type GoDialog struct {
	appkit.NSPanel
	content appkit.View
	ok      appkit.Button
	cancel  appkit.Button
}

// NewDialog create new GoDialog
func NewDialog(width, height float64) *GoDialog {
	panel := appkit.NewPanel().Autorelease()
	panel.SetFrame_Display(foundation.MakeRect(0, 0, width, height), true)

	contentView := panel.ContentView()
	view := appkit.NewView().Autorelease()
	view.SetTranslatesAutoresizingMaskIntoConstraints(false)
	ok := NewPlainButton("OK")
	cancel := NewPlainButton("Cancel")
	contentView.AddSubview(view)
	contentView.AddSubview(ok)
	contentView.AddSubview(cancel)

	ok.BottomAnchor().ConstraintEqualToAnchor_Constant(contentView.BottomAnchor(), -10).SetActive(true)
	ok.RightAnchor().ConstraintEqualToAnchor_Constant(contentView.RightAnchor(), -15).SetActive(true)
	cancel.BottomAnchor().ConstraintEqualToAnchor(ok.BottomAnchor()).SetActive(true)
	cancel.RightAnchor().ConstraintEqualToAnchor_Constant(ok.LeftAnchor(), -10).SetActive(true)
	ok.WidthAnchor().ConstraintEqualToAnchor(cancel.WidthAnchor()).SetActive(true)
	view.BottomAnchor().ConstraintEqualToAnchor_Constant(ok.TopAnchor(), -10).SetActive(true)
	view.LeftAnchor().ConstraintEqualToAnchor(contentView.LeftAnchor()).SetActive(true)
	view.TopAnchor().ConstraintEqualToAnchor(contentView.TopAnchor()).SetActive(true)
	view.RightAnchor().ConstraintEqualToAnchor(contentView.RightAnchor()).SetActive(true)

	return &GoDialog{
		NSPanel: panel,
		content: view,
		ok:      ok,
		cancel:  cancel,
	}
}

// SetView set inner content view for GoDialog
func (d *GoDialog) SetView(view appkit.View) {
	d.content.AddSubview(view)
	view.LeftAnchor().ConstraintEqualToAnchor(d.content.LeftAnchor()).SetActive(true)
	view.TopAnchor().ConstraintEqualToAnchor(d.content.TopAnchor()).SetActive(true)
	view.RightAnchor().ConstraintEqualToAnchor(d.content.RightAnchor()).SetActive(true)
	view.BottomAnchor().ConstraintEqualToAnchor(d.content.BottomAnchor()).SetActive(true)
}

func (d *GoDialog) Show(handle func()) {
	actions.Set(d.ok, func(sender objc.Object) {
		handle()
		d.Close()
	})

	actions.Set(d.cancel, func(sender objc.Object) {
		d.Close()
	})

	d.MakeKeyAndOrderFront(d.NSPanel)
}

func (d *GoDialog) RunModal() appkit.ModalResponse {
	app := appkit.SharedApplication()

	actions.Set(d.ok, func(sender objc.Object) {
		app.StopModalWithCode(appkit.ModalResponseOK)
		d.Close()
	})

	actions.Set(d.cancel, func(sender objc.Object) {
		app.StopModalWithCode(appkit.ModalResponseCancel)
		d.Close()
	})

	d.SetDelegate((&appkit.WindowDelegate{
		WindowShouldClose: func(sender appkit.Window) bool {
			app.StopModalWithCode(appkit.ModalResponseCancel)
			return true
		},
	}).ToObjc())

	return app.RunModalForWindow(d)
}
