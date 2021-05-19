#import <Appkit/Appkit.h>
#import "menu_delegate.h"
#import "_cgo_export.h"

@implementation NSMenuDelegateAdaptor


- (BOOL)menu:(NSMenu*)menu updateItem:(NSMenuItem*)item atIndex:(NSInteger)index shouldCancel:(BOOL)shouldCancel {
    bool result_ = MenuDelegate_Menu_UpdateItem_AtIndex_ShouldCancel([self goID], menu, item, index, shouldCancel);
    return result_;
}

- (NSRect)confinementRectForMenu:(NSMenu*)menu onScreen:(NSScreen*)screen {
    CGRect result_ = MenuDelegate_ConfinementRectForMenu_OnScreen([self goID], menu, screen);
    return result_;
}

- (void)menu:(NSMenu*)menu willHighlightItem:(NSMenuItem*)item {
    MenuDelegate_Menu_WillHighlightItem([self goID], menu, item);
}

- (void)menuWillOpen:(NSMenu*)menu {
    MenuDelegate_MenuWillOpen([self goID], menu);
}

- (void)menuDidClose:(NSMenu*)menu {
    MenuDelegate_MenuDidClose([self goID], menu);
}

- (NSInteger)numberOfItemsInMenu:(NSMenu*)menu {
    int result_ = MenuDelegate_NumberOfItemsInMenu([self goID], menu);
    return result_;
}

- (void)menuNeedsUpdate:(NSMenu*)menu {
    MenuDelegate_MenuNeedsUpdate([self goID], menu);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return MenuDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteMenuDelegate([self goID]);
	[super dealloc];
}
@end

void* WrapMenuDelegate(long goID) {
    NSMenuDelegateAdaptor* adaptor = [[NSMenuDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
