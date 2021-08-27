package widgets

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/cocoa/uihelper"
)

// Dialog is appkit.Panel with optional OK|CANCEL buttons
type Dialog struct {
	appkit.Panel
	content appkit.View
	ok      appkit.Button
	cancel  appkit.Button
}

// NewDialog create new Dialog
func NewDialog(width, height float64) *Dialog {
	panel := appkit.AllocPanel().Init()
	panel.SetFrame_Display(foundation.MakeRect(0, 0, width, height), true)

	contentView := panel.ContentView()
	view := appkit.AllocView().Init()
	view.SetTranslatesAutoresizingMaskIntoConstraints(false)
	ok := appkit.NewPlainButton("OK")
	cancel := appkit.NewPlainButton("Cancel")
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

	return &Dialog{
		Panel:   panel,
		content: view,
		ok:      ok,
		cancel:  cancel,
	}
}

// SetView set inner content view for Dialog
func (d *Dialog) SetView(view appkit.View) {
	d.content.AddSubview(view)
	view.LeftAnchor().ConstraintEqualToAnchor(d.content.LeftAnchor()).SetActive(true)
	view.TopAnchor().ConstraintEqualToAnchor(d.content.TopAnchor()).SetActive(true)
	view.RightAnchor().ConstraintEqualToAnchor(d.content.RightAnchor()).SetActive(true)
	view.BottomAnchor().ConstraintEqualToAnchor(d.content.BottomAnchor()).SetActive(true)
}

func (d *Dialog) Show(handle func()) {
	uihelper.SetAction(d.ok, func(sender objc.Object) {
		handle()
		d.Close()
	})

	uihelper.SetAction(d.cancel, func(sender objc.Object) {
		d.Close()
	})

	d.MakeKeyAndOrderFront(d.Panel)
}

func (d *Dialog) RunModal() appkit.ModalResponse {
	app := appkit.SharedApplication()

	uihelper.SetAction(d.ok, func(sender objc.Object) {
		app.StopModalWithCode(appkit.ModalResponseOK)
		d.Close()
	})

	uihelper.SetAction(d.cancel, func(sender objc.Object) {
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
