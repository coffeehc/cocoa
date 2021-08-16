package main

import (
	"fmt"
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/cocoa/uihelper/layouts"
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
	w := appkit.NewWindow(foundation.MakeRect(0, 0, 600, 400))
	w.SetTitle("Test Layout")

	label := appkit.NewLabel("label")
	button := appkit.NewPlainButton("button")
	button2 := appkit.NewPlainButton("button2")
	textView := appkit.ScrollableTextView()

	gridView := appkit.AllocGridView().Init()
	for i := 0; i < 3; i++ {
		var views []appkit.View
		for j := 0; j < 4; j++ {
			label := appkit.NewLabel(fmt.Sprintf("label-%v-%v", i, j))
			views = append(views, label)
		}
		gridView.AddRowWithViews(views)
	}

	stackView := appkit.StackViewWithViews([]appkit.View{label, button, button2, textView, gridView})
	stackView.SetOrientation(appkit.UserInterfaceLayoutOrientationVertical)
	stackView.SetDistribution(appkit.StackViewDistributionFillEqually)
	stackView.SetAlignment(appkit.LayoutAttributeCenterX)
	stackView.SetSpacing(10)

	layouts.AddViewWithPadding(w.ContentView(), stackView, 10, 10, 10, 10)

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
