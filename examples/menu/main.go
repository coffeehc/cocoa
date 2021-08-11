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
	w := appkit.NewWindow(foundation.MakeRect(0, 0, 600, 400))
	w.SetTitle("Test")

	// text field
	textView := appkit.ScrollableTextView()
	textView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	tv := appkit.MakeTextView(textView.DocumentView().Ptr())
	tv.SetAllowsUndo(true)
	tv.SetRichText(false)
	w.ContentView().AddSubview(textView)
	w.ContentView().LeadingAnchor().ConstraintEqualToAnchor_Constant(textView.LeadingAnchor(), -10).SetActive(true)
	w.ContentView().TopAnchor().ConstraintEqualToAnchor_Constant(textView.TopAnchor(), -10).SetActive(true)
	w.ContentView().TrailingAnchor().ConstraintEqualToAnchor_Constant(textView.TrailingAnchor(), 10).SetActive(true)
	w.ContentView().BottomAnchor().ConstraintEqualToAnchor_Constant(textView.BottomAnchor(), 10).SetActive(true)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	appDelegate := &appkit.ApplicationDelegate{
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
	}
	app.SetDelegate(appDelegate.ToObjc())

	// should set system bar after window show
	setSystemBar(app)

	app.Run()
}

func setMainMenu(app appkit.Application) {
	menu := appkit.AllocMenu().InitWithTitle("main")
	app.SetMainMenu(menu)

	mainMenuItem := appkit.NewMenuItem("", "", nil)
	mainMenuMenu := appkit.AllocMenu().InitWithTitle("App")
	mainMenuMenu.AddItem(appkit.NewMenuItemWithAction("Hide", "h", app.Hide))
	mainMenuMenu.AddItem(appkit.NewMenuItemWithAction("Quit", "q", app.Terminate))
	mainMenuItem.SetSubmenu(mainMenuMenu)
	menu.AddItem(mainMenuItem)

	testMenuItem := appkit.NewMenuItem("", "", nil)
	testMenu := appkit.AllocMenu().InitWithTitle("Test")
	testMenu.AddItem(appkit.NewMenuItem("Select All", "a", foundation.SelectorFromString("selectAll:")))
	testMenu.AddItem(appkit.MenuItem_SeparatorItem())
	testMenu.AddItem(appkit.NewMenuItem("Copy", "c", foundation.SelectorFromString("copy:")))
	testMenu.AddItem(appkit.NewMenuItem("Paste", "v", foundation.SelectorFromString("paste:")))
	testMenu.AddItem(appkit.NewMenuItem("Cut", "x", foundation.SelectorFromString("cut:")))
	testMenu.AddItem(appkit.NewMenuItem("Undo", "z", foundation.SelectorFromString("undo:")))
	testMenu.AddItem(appkit.NewMenuItem("Redo", "Z", foundation.SelectorFromString("redo:")))
	testMenuItem.SetSubmenu(testMenu)
	menu.AddItem(testMenuItem)
}

func setSystemBar(app appkit.Application) {
	bar := appkit.SystemStatusBar()
	item := bar.StatusItemWithLength(appkit.VariableStatusItemLength)
	button := item.Button()
	button.SetTitle("TestTray")

	menu := appkit.AllocMenu().InitWithTitle("main")
	menu.AddItem(appkit.NewMenuItemWithAction("Hide", "h", app.Hide))
	menu.AddItem(appkit.NewMenuItemWithAction("Quit", "q", app.Terminate))
	item.SetMenu(menu)
}

func main() {
	objc.WithAutoreleasePool(func() {
		initAndRun()
	})
}
