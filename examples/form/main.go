package main

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/appkits"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.SharedApplication()
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	w := appkits.NewWindow(600, 400)
	w.SetTitle("Form")

	fv := appkits.NewFormView()
	fv.SetLabelWidth(100)
	fv.SetLabelAlignment(appkits.LabelAlignmentTrailing)
	fv.SetLabelControlSpacing(10)
	fv.AddRow("user", appkits.NewTextField())
	fv.SetLabelFont(appkit.UserFontOfSize(13))
	fv.AddRow("password", appkits.NewSecureTextField())
	cb := appkits.NewCheckBox("")
	fv.AddRow("males", cb)
	fv.AddExpandRow()

	w.ContentView().AddSubview(fv)
	fv.LeftAnchor().ConstraintEqualToAnchor_Constant(w.ContentView().LeftAnchor(), 10).SetActive(true)
	fv.TopAnchor().ConstraintEqualToAnchor_Constant(w.ContentView().TopAnchor(), 10).SetActive(true)
	fv.RightAnchor().ConstraintEqualToAnchor_Constant(w.ContentView().RightAnchor(), -10).SetActive(true)
	fv.BottomAnchor().ConstraintEqualToAnchor_Constant(w.ContentView().BottomAnchor(), -10).SetActive(true)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	appDelegate := &appkit.ApplicationDelegate{
		ApplicationDidFinishLaunching: func(notification foundation.Notification) {
			app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
			app.ActivateIgnoringOtherApps(true)
		},
		ApplicationShouldTerminateAfterLastWindowClosed: func(sender appkit.Application) bool {
			return true
		},
	}
	app.SetDelegate(appDelegate.ToObjc())

	app.Run()
}

func main() {
	objc.WithAutoreleasePool(initAndRun)

}
