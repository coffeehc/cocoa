package main

import (
	"fmt"
	"github.com/hsiafan/cocoa/appkit"
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
	w := appkit.NewWindow(foundation.MakeRect(0, 0, 600, 400))
	w.SetTitle("Test Layout")

	label := appkit.NewLabel()
	label.SetStringValue("label")
	button := appkit.NewPlainButton()
	button.SetStringValue("button")
	button2 := appkit.NewPlainButton()
	button2.SetStringValue("button2")
	textView := appkit.ScrollableTextView()

	gridView := appkit.AllocGridView().Init()
	for i := 0; i < 3; i++ {
		var views []appkit.View
		for j := 0; j < 4; j++ {
			label := appkit.NewLabel()
			label.SetStringValue(fmt.Sprintf("label-%v-%v", i, j))
			views = append(views, label)
		}
		gridView.AddRowWithViews(views)
	}

	stackView := appkit.StackViewWithViews([]appkit.View{label, button, button2, textView, gridView})
	stackView.SetOrientation(appkit.UserInterfaceLayoutOrientationVertical)
	stackView.SetDistribution(appkit.NSStackViewDistributionFillEqually)
	stackView.SetAlignment(appkit.LayoutAttributeCenterX)
	stackView.SetSpacing(10)

	w.ContentView().AddSubview(stackView)
	w.ContentView().LeadingAnchor().ConstraintEqualToAnchor_Constant(stackView.LeadingAnchor(), -10).SetActive(true)
	w.ContentView().TopAnchor().ConstraintEqualToAnchor_Constant(stackView.TopAnchor(), -10).SetActive(true)
	w.ContentView().TrailingAnchor().ConstraintEqualToAnchor_Constant(stackView.TrailingAnchor(), 10).SetActive(true)
	w.ContentView().BottomAnchor().ConstraintEqualToAnchor_Constant(stackView.BottomAnchor(), 10).SetActive(true)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	app.SetDelegate(appkit.WrapApplicationDelegate(&appkit.ApplicationDelegate{
		ApplicationDidFinishLaunching: func(notification foundation.Notification) {
			app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
			app.ActivateIgnoringOtherApps(true)
		},
		ApplicationShouldTerminateAfterLastWindowClosed: func(sender appkit.Application) bool {
			return true
		},
	}))

	app.Run()
}

func main() {
	objc.WithAutoreleasePool(func() {
		initAndRun()
	})

}
