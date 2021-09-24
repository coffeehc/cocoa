#import "menu_delegate.h"
#import <AppKit/NSMenu.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

@interface NSMenuDelegateAdaptor : NSObject <NSMenuDelegate>
@property (assign) uintptr_t goID;
@end

@implementation NSMenuDelegateAdaptor


- (BOOL)menu:(NSMenu*)menu updateItem:(NSMenuItem*)item atIndex:(NSInteger)index shouldCancel:(BOOL)shouldCancel {
    bool result_ = menuDelegate_Menu_UpdateItem_AtIndex_ShouldCancel([self goID], menu, item, index, shouldCancel);
    return result_;
}

- (NSRect)confinementRectForMenu:(NSMenu*)menu onScreen:(NSScreen*)screen {
    CGRect result_ = menuDelegate_ConfinementRectForMenu_OnScreen([self goID], menu, screen);
    return result_;
}

- (void)menu:(NSMenu*)menu willHighlightItem:(NSMenuItem*)item {
    menuDelegate_Menu_WillHighlightItem([self goID], menu, item);
}

- (void)menuWillOpen:(NSMenu*)menu {
    menuDelegate_MenuWillOpen([self goID], menu);
}

- (void)menuDidClose:(NSMenu*)menu {
    menuDelegate_MenuDidClose([self goID], menu);
}

- (NSInteger)numberOfItemsInMenu:(NSMenu*)menu {
    int result_ = menuDelegate_NumberOfItemsInMenu([self goID], menu);
    return result_;
}

- (void)menuNeedsUpdate:(NSMenu*)menu {
    menuDelegate_MenuNeedsUpdate([self goID], menu);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return MenuDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapMenuDelegate(uintptr_t goID) {
    NSMenuDelegateAdaptor* adaptor = [[NSMenuDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
