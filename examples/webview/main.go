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
	app := appkit.InitSharedApplication()
	w := appkit.NewPlainWindow(foundation.MakeRect(0, 0, 600, 400))
	w.SetTitle("Test widgets")

	webView := webkit.NewWKWebView(foundation.ZeroRect)
	webView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	webView.LoadRequest(foundation.NewURLRequest(foundation.URLWithString("https://www.baidu.com")))
	//webView.LoadHTMLString("<h1>Test</h1>", "https://www.baidu.com")
	w.ContentView().AddSubview(webView)
	w.ContentView().LeadingAnchor().ConstraintEqualToAnchor2(webView.LeadingAnchor(), -10).SetActive(true)
	w.ContentView().TopAnchor().ConstraintEqualToAnchor2(webView.TopAnchor(), -10).SetActive(true)
	w.ContentView().TrailingAnchor().ConstraintEqualToAnchor2(webView.TrailingAnchor(), 10).SetActive(true)
	w.ContentView().BottomAnchor().ConstraintEqualToAnchor2(webView.BottomAnchor(), 10).SetActive(true)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	app.ApplicationDidFinishLaunching(func(notification foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
	})

	app.Run()
}

func main() {
	objc.WithAutoreleasePool(func() {
		initAndRun()
	})
}
