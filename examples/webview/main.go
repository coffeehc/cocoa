package main

import (
	"fmt"
	"github.com/hsiafan/cocoa/appkits"
	"github.com/hsiafan/cocoa/helper/actions"
	"github.com/hsiafan/cocoa/helper/delegates"
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
	w := appkits.NewWindow(600, 400)
	w.SetTitle("Test widgets")

	sv := appkits.NewVerticalStackView()
	w.SetContentView(sv)

	webView := webkit.NewWebView().Autorelease()
	webView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	webView.LoadRequest(foundation.AllocURLRequest().InitWithURL(foundation.URLWithString("https://www.baidu.com")).Autorelease())
	sv.AddView_InGravity(webView, appkit.StackViewGravityTop)

	snapshotButton := appkits.NewPlainButton("capture")

	snapshotWin := appkits.NewWindow(0, 0)
	snapshotWin.SetTitle("Test widgets")

	snapshotWebView := webkit.NewWebView().Autorelease()
	snapshotWebView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	snapshotWin.SetContentView(snapshotWebView)

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
					snapshotWin.SetFrame_Display(foundation.MakeRect(0, 0, width, height), true)
					snapshotWebView.LoadRequest(foundation.AllocURLRequest().InitWithURL(foundation.URLWithString("https://www.baidu.com")))

					actions.Set(snapshotButton, func(sender objc.Object) {
						snapshotWebView.TakeSnapshotWithConfiguration(nil, func(image appkit.Image, err foundation.Error) {
							imageRef := image.CGImageForProposedRect_Context_Hints()
							imageRepo := appkit.AllocBitmapImageRep().InitWithCGImage(imageRef).Autorelease()
							imageRepo.SetSize(image.Size())
							pngData := imageRepo.RepresentationUsingType_Properties(appkit.BitmapImageFileTypePNG, nil)

							if err := os.WriteFile("webview_screenshot.png", pngData, os.ModePerm); err != nil {
								fmt.Println("write image to file error:", err)
							} else {
								fmt.Println("image captured to webview_screenshot.png")
							}
						})
					})
				})
			})
		},
	}).ToObjc())

	sv.AddView_InGravity(snapshotButton, appkit.StackViewGravityTop)

	delegates.Set(w, &appkit.WindowDelegate{
		WindowWillClose: func(notification foundation.Notification) {
			snapshotWin.Close()
		},
	})

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
	initAndRun()
}
