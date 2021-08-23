package main

import (
	"fmt"
	"github.com/hsiafan/cocoa/uihelper"
	"github.com/hsiafan/cocoa/uihelper/layouts"
	"os"
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

	layouts.AddViewWithPadding(w.ContentView(), webView, 10, 10, 10, 20)
	cb := appkit.NewPlainButton("capture")
	uihelper.SetAction(cb, func(sender objc.Object) {
		bitMapImageRep := webView.BitmapImageRepForCachingDisplayInRect(webView.Bounds())
		data := bitMapImageRep.RepresentationUsingType_Properties(appkit.BitmapImageFileTypePNG, map[appkit.BitmapImageRepPropertyKey]objc.Object{
			"xxxxx": foundation.NewString("1"),
		})
		if err := os.WriteFile("a.png", data, os.ModePerm); err != nil {
			fmt.Println("write image to file error:", err)
		}
	})
	w.ContentView().AddSubview(cb)
	cb.LeftAnchor().ConstraintEqualToAnchor_Constant(w.ContentView().LeftAnchor(), -10)
	cb.BottomAnchor().ConstraintEqualToAnchor_Constant(w.ContentView().BottomAnchor(), -10)

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
