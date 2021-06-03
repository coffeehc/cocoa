package appkit

// #include "tab_view_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TabViewDelegate struct {
	TabViewDidChangeNumberOfTabViewItems func(tabView TabView)
	TabView_ShouldSelectTabViewItem      func(tabView TabView, tabViewItem TabViewItem) bool
	TabView_WillSelectTabViewItem        func(tabView TabView, tabViewItem TabViewItem)
	TabView_DidSelectTabViewItem         func(tabView TabView, tabViewItem TabViewItem)
}

func WrapTabViewDelegate(delegate *TabViewDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapTabViewDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export tabViewDelegate_TabViewDidChangeNumberOfTabViewItems
func tabViewDelegate_TabViewDidChangeNumberOfTabViewItems(id int64, tabView unsafe.Pointer) {
	delegate := resources.Get(id).(*TabViewDelegate)
	delegate.TabViewDidChangeNumberOfTabViewItems(MakeTabView(tabView))
}

//export tabViewDelegate_TabView_ShouldSelectTabViewItem
func tabViewDelegate_TabView_ShouldSelectTabViewItem(id int64, tabView unsafe.Pointer, tabViewItem unsafe.Pointer) C.bool {
	delegate := resources.Get(id).(*TabViewDelegate)
	result := delegate.TabView_ShouldSelectTabViewItem(MakeTabView(tabView), MakeTabViewItem(tabViewItem))
	return C.bool(result)
}

//export tabViewDelegate_TabView_WillSelectTabViewItem
func tabViewDelegate_TabView_WillSelectTabViewItem(id int64, tabView unsafe.Pointer, tabViewItem unsafe.Pointer) {
	delegate := resources.Get(id).(*TabViewDelegate)
	delegate.TabView_WillSelectTabViewItem(MakeTabView(tabView), MakeTabViewItem(tabViewItem))
}

//export tabViewDelegate_TabView_DidSelectTabViewItem
func tabViewDelegate_TabView_DidSelectTabViewItem(id int64, tabView unsafe.Pointer, tabViewItem unsafe.Pointer) {
	delegate := resources.Get(id).(*TabViewDelegate)
	delegate.TabView_DidSelectTabViewItem(MakeTabView(tabView), MakeTabViewItem(tabViewItem))
}

//export TabViewDelegate_RespondsTo
func TabViewDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*TabViewDelegate)
	switch selName {
	case "tabViewDidChangeNumberOfTabViewItems:":
		return delegate.TabViewDidChangeNumberOfTabViewItems != nil
	case "tabView:shouldSelectTabViewItem:":
		return delegate.TabView_ShouldSelectTabViewItem != nil
	case "tabView:willSelectTabViewItem:":
		return delegate.TabView_WillSelectTabViewItem != nil
	case "tabView:didSelectTabViewItem:":
		return delegate.TabView_DidSelectTabViewItem != nil
	default:
		return false
	}
}

//export deleteTabViewDelegate
func deleteTabViewDelegate(id int64) {
	resources.Delete(id)
}
