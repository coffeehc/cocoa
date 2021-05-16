package appkit

import "github.com/hsiafan/cocoa/uihelper"

type MenuProperties uint

const (
	MenuPropertyItemTitle                    MenuProperties = 1 << 0 // the menu item's title
	MenuPropertyItemAttributedTitle          MenuProperties = 1 << 1 // the menu item's attributed title
	MenuPropertyItemKeyEquivalent            MenuProperties = 1 << 2 // the menu item's key equivalent
	MenuPropertyItemImage                    MenuProperties = 1 << 3 // the menu item's image
	MenuPropertyItemEnabled                  MenuProperties = 1 << 4 // whether the menu item is enabled or disabled
	MenuPropertyItemAccessibilityDescription MenuProperties = 1 << 5 // the menu item's accessibility description
)

func NewMenuItem(title string, charCode string, handler uihelper.ActionHandler) MenuItem {
	item := AllocMenuItem().InitWithTitle_Action_KeyEquivalent(title, nil, charCode)
	uihelper.SetAction(item, handler)
	return item
}
