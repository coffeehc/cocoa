#import <Appkit/Appkit.h>
#import "combo_box_data_source.h"
#import "_cgo_export.h"

@implementation NSComboBoxDataSourceAdaptor


- (NSString*)comboBox:(NSComboBox*)comboBox completedString:(NSString*)_string {
    void* result_ = ComboBoxDataSource_ComboBox_CompletedString([self goID], comboBox, _string);
    return (NSString*)result_;
}

- (NSUInteger)comboBox:(NSComboBox*)comboBox indexOfItemWithStringValue:(NSString*)_string {
    unsigned int result_ = ComboBoxDataSource_ComboBox_IndexOfItemWithStringValue([self goID], comboBox, _string);
    return result_;
}

- (id)comboBox:(NSComboBox*)comboBox objectValueForItemAtIndex:(NSInteger)index {
    void* result_ = ComboBoxDataSource_ComboBox_ObjectValueForItemAtIndex([self goID], comboBox, index);
    return (id)result_;
}

- (NSInteger)numberOfItemsInComboBox:(NSComboBox*)comboBox {
    int result_ = ComboBoxDataSource_NumberOfItemsInComboBox([self goID], comboBox);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return ComboBoxDataSource_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteComboBoxDataSource([self goID]);
	[super dealloc];
}
@end

void* WrapComboBoxDataSource(long goID) {
    NSComboBoxDataSourceAdaptor* adaptor = [[NSComboBoxDataSourceAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
