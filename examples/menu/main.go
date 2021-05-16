package main

import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime"
)

// menus and tray

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.SharedApplication()
	w := appkit.NewPlainWindow(foundation.MakeRect(0, 0, 600, 400))
	w.SetTitle("Test")

	// text field
	textView := appkit.ScrollableTextView()
	textView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tv := appkit.MakeTextView(textView.DocumentView().Ptr())
	tv.SetAllowsUndo(true)
	tv.SetRichText(false)
	w.ContentView().AddSubview(textView)
	w.ContentView().LeadingAnchor().ConstraintEqualToAnchor2(textView.LeadingAnchor(), -10).SetActive(true)
	w.ContentView().TopAnchor().ConstraintEqualToAnchor2(textView.TopAnchor(), -10).SetActive(true)
	w.ContentView().TrailingAnchor().ConstraintEqualToAnchor2(textView.TrailingAnchor(), 10).SetActive(true)
	w.ContentView().BottomAnchor().ConstraintEqualToAnchor2(textView.BottomAnchor(), 10).SetActive(true)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	app.SetDelegate(appkit.WrapApplicationDelegate(&appkit.ApplicationDelegate{
		ApplicationDidFinishLaunching: func(notification foundation.Notification) {
			app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
			app.ActivateIgnoringOtherApps(true)
		},
		ApplicationWillFinishLaunching: func(notification foundation.Notification) {
			// should set menu bar at ApplicationWillFinishLaunching
			setMainMenu(app)
		},
		ApplicationShouldTerminateAfterLastWindowClosed: func(sender appkit.Application) bool {
			return true
		},
	}))

	// should set system bar after window show
	setSystemBar(app)

	app.Run()
}

func setMainMenu(app appkit.Application) {
	menu := appkit.NewMenu("main")
	app.SetMainMenu(menu)

	mainMenuItem := appkit.NewMenuItemWithHandler("", "", func(sender objc.Object) {
	})
	mainMenuMenu := appkit.NewMenu("App")
	mainMenuMenu.AddItem(appkit.NewMenuItemWithHandler("Hide", "h", func(sender objc.Object) {
		app.Hide(nil)
	}))
	mainMenuMenu.AddItem(appkit.NewMenuItemWithHandler("Quit", "q", func(sender objc.Object) {
		app.Terminate(nil)
	}))
	mainMenuItem.SetSubmenu(mainMenuMenu)
	menu.AddItem(mainMenuItem)

	testMenuItem := appkit.NewMenuItemWithHandler("", "", func(sender objc.Object) {
	})
	testMenu := appkit.NewMenu("Test")
	testMenu.AddItem(appkit.NewMenuItem("Select All", foundation.SelectorFromString("selectAll:"), "a"))
	testMenu.AddItem(appkit.NewSeparatorItem())
	testMenu.AddItem(appkit.NewMenuItem("Copy", foundation.SelectorFromString("copy:"), "c"))
	testMenu.AddItem(appkit.NewMenuItem("Paste", foundation.SelectorFromString("paste:"), "v"))
	testMenu.AddItem(appkit.NewMenuItem("Cut", foundation.SelectorFromString("cut:"), "x"))
	testMenu.AddItem(appkit.NewMenuItem("Undo", foundation.SelectorFromString("undo:"), "z"))
	testMenu.AddItem(appkit.NewMenuItem("Redo", foundation.SelectorFromString("redo:"), "Z"))
	testMenuItem.SetSubmenu(testMenu)
	menu.AddItem(testMenuItem)
}

func setSystemBar(app appkit.Application) {
	bar := appkit.SystemStatusBar()
	item := bar.StatusItemWithLength(appkit.VariableStatusItemLength)
	button := item.Button()
	button.SetTitle("TestTray")

	menu := appkit.NewMenu("main")
	menu.AddItem(appkit.NewMenuItemWithHandler("Hide", "h", func(sender objc.Object) {
		app.Hide(nil)
	}))
	menu.AddItem(appkit.NewMenuItemWithHandler("Quit", "q", func(sender objc.Object) {
		app.Terminate(nil)
	}))
	item.SetMenu(menu)
}

func main() {
	objc.WithAutoreleasePool(func() {
		initAndRun()
	})
}
