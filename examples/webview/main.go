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
	w := appkit.NewWindow(600, 400)
	w.SetTitle("Test widgets")

	webView := webkit.AllocWebView().Init()
	webView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	webView.LoadRequest(foundation.AllocURLRequest().InitWithURL(foundation.URLWithString("https://www.baidu.com")))

	layouts.AddViewWithPadding(w.ContentView(), webView, 10, 10, 10, 20)
	cb := appkit.NewPlainButton("capture")

	webView.SetNavigationDelegate((&webkit.NavigationDelegate{
		WebView_DidFinishNavigation: func(webView webkit.WebView, navigation webkit.Navigation) {
			webView.EvaluateJavaScript("document.documentElement.outerHTML.toString()", func(value objc.Object, err foundation.Error) {
				content := foundation.MakeString(value.Ptr()).String()
				_ = content
			})
			webView.EvaluateJavaScript("document.body.scrollWidth", func(value objc.Object, err foundation.Error) {
				width := foundation.MakeNumber(value.Ptr()).DoubleValue()
				webView.EvaluateJavaScript("document.body.scrollHeight", func(value objc.Object, err foundation.Error) {
					height := foundation.MakeNumber(value.Ptr()).DoubleValue()
					nw := appkit.NewWindow(width, height)
					nw.SetTitle("Test widgets")

					nwv := webkit.AllocWebView().Init()
					nwv.SetTranslatesAutoresizingMaskIntoConstraints(false)
					//nwv.SetFrame(foundation.MakeRect(0, 0, width, height))
					//nwv.WidthAnchor().ConstraintEqualToConstant(width).SetActive(true)
					nwv.LoadRequest(foundation.AllocURLRequest().InitWithURL(foundation.URLWithString("https://www.baidu.com")))
					nw.SetContentView(nwv)

					uihelper.SetAction(cb, func(sender objc.Object) {
						nwv.TakeSnapshotWithConfiguration(nil, func(image appkit.Image, err foundation.Error) {
							imageRef := image.CGImageForProposedRect_Context_Hints()
							imageRepo := appkit.AllocBitmapImageRep().InitWithCGImage(imageRef)
							imageRepo.SetSize(image.Size())
							pngData := imageRepo.RepresentationUsingType_Properties(appkit.BitmapImageFileTypePNG, nil)

							if err := os.WriteFile("webview_screenshot.png", pngData, os.ModePerm); err != nil {
								fmt.Println("write image to file error:", err)
							} else {
								fmt.Println("image captured to webview_screenshot.png")
							}
							nw.Close()
						})
					})
				})
			})
		},
	}).ToObjc())

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
