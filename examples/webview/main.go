package main

import (
	"runtime"

	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"github.com/hsiafan/cocoa/webkit"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.SharedApplication()
	w := appkit.NewWindow(foundation.MakeRect(0, 0, 600, 400))
	w.SetTitle("Test widgets")

	webView := webkit.AllocWebView().Init()
	webView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	webView.LoadRequest(foundation.AllocURLRequest().InitWithURL(foundation.URLWithString("https://www.baidu.com")))
	//webView.LoadHTMLString("<h1>Test</h1>", "https://www.baidu.com")
	w.ContentView().AddSubview(webView)
	w.ContentView().LeadingAnchor().ConstraintEqualToAnchor_Constant(webView.LeadingAnchor(), -10).SetActive(true)
	w.ContentView().TopAnchor().ConstraintEqualToAnchor_Constant(webView.TopAnchor(), -10).SetActive(true)
	w.ContentView().TrailingAnchor().ConstraintEqualToAnchor_Constant(webView.TrailingAnchor(), 10).SetActive(true)
	w.ContentView().BottomAnchor().ConstraintEqualToAnchor_Constant(webView.BottomAnchor(), 10).SetActive(true)

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
	objc.WithAutoreleasePool(func() {
		initAndRun()
	})
}
