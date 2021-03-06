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
	"time"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.SharedApplication()
	w := appkits.NewWindow(600, 400)

	w.SetTitle("Test widgets")

	filePathField := appkit.AllocTextField().InitWithFrame(foundation.MakeRect(10, 330, 200, 20)).Autorelease()
	cell := filePathField.Cell()
	cell.SetWraps(false)
	cell.SetScrollable(true)
	filePathField.SetEditable(false)
	w.ContentView().AddSubview(filePathField)

	saveButton := appkits.NewPlainButton("Save...")
	saveButton.SetFrame(foundation.MakeRect(250, 330, 80, 20))
	actions.Set(saveButton, func(sender objc.Object) {
		savePanel := appkit.NewSavePanel().Autorelease()
		if savePanel.RunModal() == appkit.ModalResponseOK {
			filePathField.SetStringValue(savePanel.URL().Path())
		}
	})
	w.ContentView().AddSubview(saveButton)

	presentationTF := appkit.AllocTextField().InitWithFrame(foundation.MakeRect(10, 290, 100, 25)).Autorelease()
	w.ContentView().AddSubview(presentationTF)

	stepper := appkit.AllocStepper().InitWithFrame(foundation.MakeRect(130, 290, 16, 25)).Autorelease()
	stepper.SetDoubleValue(100)
	w.ContentView().AddSubview(stepper)

	colorWell := appkit.AllocColorWell().InitWithFrame(foundation.MakeRect(160, 290, 30, 25)).Autorelease()
	w.ContentView().AddSubview(colorWell)

	comboBox := appkit.AllocComboBox().InitWithFrame(foundation.MakeRect(210, 290, 100, 25)).Autorelease()
	comboBox.AddItemsWithObjectValues([]objc.Object{
		foundation.NewString("Test1"),
		foundation.NewString("Test2"),
	})
	comboBox.SelectItemAtIndex(0)
	w.ContentView().AddSubview(comboBox)

	slider := appkit.AllocSlider().InitWithFrame(foundation.MakeRect(330, 290, 100, 25)).Autorelease()
	actions.Set(slider, func(sender objc.Object) {
		presentationTF.SetDoubleValue(slider.DoubleValue())
	})
	slider.SetMaxValue(10)
	w.ContentView().AddSubview(slider)

	datePicker := appkit.AllocDatePicker().InitWithFrame(foundation.MakeRect(450, 290, 140, 25)).Autorelease()
	w.ContentView().AddSubview(datePicker)

	// buttons
	cb := appkits.NewCheckBox("check box")
	cb.SetFrame(foundation.MakeRect(10, 250, 80, 25))
	w.ContentView().AddSubview(cb)

	rb := appkits.NewRadioButton("radio button")
	rb.SetFrame(foundation.MakeRect(150, 250, 120, 25))
	w.ContentView().AddSubview(rb)

	sw := appkit.AllocSwitch().InitWithFrame(foundation.MakeRect(260, 250, 120, 25)).Autorelease()
	w.ContentView().AddSubview(sw)

	li := appkit.AllocLevelIndicator().InitWithFrame(foundation.MakeRect(370, 250, 120, 25)).Autorelease()
	li.SetCriticalValue(4)
	li.SetDoubleValue(3)
	w.ContentView().AddSubview(li)

	btn := appkits.NewPlainButton("change color")
	btn.SetFrame(foundation.MakeRect(10, 160, 120, 25))
	w.ContentView().AddSubview(btn)

	quitBtn := appkits.NewPlainButton("Quit")
	quitBtn.SetFrame(foundation.MakeRect(10, 130, 80, 25))
	actions.Set(quitBtn, func(sender objc.Object) {
		app.Terminate(nil)
	})
	w.ContentView().AddSubview(quitBtn)

	// text field
	tf := appkits.NewTextField()
	w.ContentView().AddSubview(tf)
	tf.SetFrame(foundation.MakeRect(10, 100, 150, 25))

	// label
	label := appkits.NewLabel("")
	label.SetFrame(foundation.MakeRect(170, 100, 150, 25))
	w.ContentView().AddSubview(label)
	delegates.Set(tf, &appkit.TextFieldDelegate{
		ControlTextDidChange: func(obj foundation.Notification) {
			objc.DispatchAsyncToMainQueue(func() {
				label.SetStringValue(tf.StringValue())
			})
		},
	})
	actions.Set(btn, func(sender objc.Object) {
		label.SetTextColor(appkit.RedColor())
	})

	// password
	stf := appkits.NewSecureTextField()
	stf.SetFrame(foundation.MakeRect(340, 100, 150, 25))
	delegates.Set(stf, &appkit.TextFieldDelegate{
		ControlTextDidChange: func(obj foundation.Notification) {
			objc.DispatchAsyncToMainQueue(func() {
				label.SetStringValue(tf.StringValue())
			})
		},
	})
	w.ContentView().AddSubview(stf)

	// progress indicator
	indicator := appkit.AllocProgressIndicator().InitWithFrame(foundation.MakeRect(10, 70, 350, 25)).Autorelease()
	indicator.SetMinValue(0)
	indicator.SetMaxValue(1)
	indicator.SetIndeterminate(false)
	indicator.SetDisplayedWhenStopped(false)
	w.ContentView().AddSubview(indicator)
	go func() {
		objc.DispatchAsyncToMainQueue(func() {
			indicator.StartAnimation(indicator)
		})
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			objc.DispatchAsyncToMainQueue(func() {
				indicator.SetDoubleValue(0.1 * float64(i))
			})
		}
		objc.DispatchAsyncToMainQueue(func() {
			indicator.StopAnimation(indicator)
		})
	}()

	// text view & scroll view
	sv := appkit.ScrollableTextView()
	sv.SetFrame(foundation.MakeRect(10, 200, 200, 30))
	appkit.MakeTextView(sv.DocumentView().Ptr()).SetAllowsUndo(true)
	w.ContentView().AddSubview(sv)

	sv2 := appkit.ScrollableTextView()
	sv2.SetFrame(foundation.MakeRect(250, 200, 200, 30))
	appkit.MakeTextView(sv2.DocumentView().Ptr()).SetAllowsUndo(true)
	w.ContentView().AddSubview(sv2)

	delegates.Set(w, &appkit.WindowDelegate{
		WindowDidMove: func(notification foundation.Notification) {
			frame := w.Frame()
			fmt.Println("window move to ", frame.Origin.X, frame.Origin.Y)
		},
	})
	w.Center()
	w.MakeKeyAndOrderFront(nil)

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
