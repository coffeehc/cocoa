package main

import (
	"fmt"
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/appkits"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/helper/actions"
	"github.com/hsiafan/cocoa/helper/delegates"
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
	w.SetTitle("Test Layout")

	label := appkits.NewLabel("label")
	mdButton := appkits.NewPlainButton("modal dialog")
	dButton := appkits.NewPlainButton("dialog")
	textView := appkit.ScrollableTextView()

	actions.Set(mdButton, func(sender objc.Object) {
		d := appkits.NewDialog(400, 300)
		d.SetView(appkits.NewLabel("test modal dialog"))
		if d.RunModal() == appkit.ModalResponseOK {
			fmt.Println("ok!")
		}
	})

	actions.Set(dButton, func(sender objc.Object) {
		d := appkits.NewDialog(400, 300)
		d.SetView(appkits.NewLabel("test dialog"))
		d.Center()
		d.Show(func() {
			fmt.Println("ok!")
		})
	})

	gridView := appkit.NewGridView().Autorelease()
	for i := 0; i < 3; i++ {
		var views []appkit.View
		for j := 0; j < 4; j++ {
			label := appkits.NewLabel(fmt.Sprintf("label-%v-%v", i, j))
			views = append(views, label)
		}
		gridView.AddRowWithViews(views)
	}
	gridView.SetContentHuggingPriority_ForOrientation(appkit.LayoutPriorityDefaultHigh, appkit.LayoutConstraintOrientationHorizontal)
	gridView.ColumnAtIndex(0).SetXPlacement(appkit.GridCellPlacementTrailing)
	gridView.SetRowAlignment(appkit.GridRowAlignmentLastBaseline)

	stackView := appkit.StackViewWithViews([]appkit.View{label, mdButton, dButton, textView, gridView})
	stackView.SetOrientation(appkit.UserInterfaceLayoutOrientationVertical)
	stackView.SetDistribution(appkit.StackViewDistributionFillEqually)
	stackView.SetAlignment(appkit.LayoutAttributeCenterX)
	stackView.SetSpacing(10)

	w.SetContentView(stackView)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	delegates.Set(app, &appkit.ApplicationDelegate{
		ApplicationDidFinishLaunching: func(notification foundation.Notification) {
			app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
			app.ActivateIgnoringOtherApps(true)
		},
		ApplicationShouldTerminateAfterLastWindowClosed: func(sender appkit.Application) bool {
			return true
		},
	})

	app.Run()
}

func main() {
	objc.WithAutoreleasePool(initAndRun)

}
