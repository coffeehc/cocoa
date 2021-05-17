#import <Appkit/Appkit.h>
#import "tab_view_delegate.h"
#import "_cgo_export.h"

@implementation NSTabViewDelegateAdaptor


- (void)tabViewDidChangeNumberOfTabViewItems:(NSTabView*)tabView {
    TabViewDelegate_TabViewDidChangeNumberOfTabViewItems([self goID], tabView);
}

- (BOOL)tabView:(NSTabView*)tabView shouldSelectTabViewItem:(NSTabViewItem*)tabViewItem {
    bool result_ = TabViewDelegate_TabView_ShouldSelectTabViewItem([self goID], tabView, tabViewItem);
    return result_;
}

- (void)tabView:(NSTabView*)tabView willSelectTabViewItem:(NSTabViewItem*)tabViewItem {
    TabViewDelegate_TabView_WillSelectTabViewItem([self goID], tabView, tabViewItem);
}

- (void)tabView:(NSTabView*)tabView didSelectTabViewItem:(NSTabViewItem*)tabViewItem {
    TabViewDelegate_TabView_DidSelectTabViewItem([self goID], tabView, tabViewItem);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return TabViewDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteTabViewDelegate([self goID]);
	[super dealloc];
}
@end

void* WrapTabViewDelegate(long goID) {
    NSTabViewDelegateAdaptor* adaptor = [[NSTabViewDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
