package appkit

import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

//TODO: fina a way to support ComboBoxDataSource
type ComboBoxDataSource interface {
	objc.Object
	completedString(box ComboBox, completedString string) string
	indexOfItemWithStringValue(box ComboBox, str string) uint
	objectValueForItemAtIndex(box ComboBox, index int) objc.Object
	numberOfItemsInComboBox(box ComboBox) int
}

type ComboBoxDataSourceDelegate struct {
	objc.NSObject
}

func (c *ComboBoxDataSourceDelegate) completedString(box ComboBox, completedString string) string {
	panic("implement me")
}

func (c *ComboBoxDataSourceDelegate) indexOfItemWithStringValue(box ComboBox, str string) uint {
	panic("implement me")
}

func (c *ComboBoxDataSourceDelegate) objectValueForItemAtIndex(box ComboBox, index int) objc.Object {
	panic("implement me")
}

func (c *ComboBoxDataSourceDelegate) numberOfItemsInComboBox(box ComboBox) int {
	panic("implement me")
}

func MakeComboBoxDataSource(ptr unsafe.Pointer) *ComboBoxDataSourceDelegate {
	return nil
}
