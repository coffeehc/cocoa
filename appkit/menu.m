#import "menu.h"
#import <AppKit/NSMenu.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

void* C_Menu_Alloc() {
    return [NSMenu alloc];
}

void* C_NSMenu_InitWithTitle(void* ptr, void* title) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenu* result_ = [nSMenu initWithTitle:(NSString*)title];
    return result_;
}

void* C_NSMenu_InitWithCoder(void* ptr, void* coder) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenu* result_ = [nSMenu initWithCoder:(NSCoder*)coder];
    return result_;
}

void* C_NSMenu_AllocMenu() {
    NSMenu* result_ = [NSMenu alloc];
    return result_;
}

void* C_NSMenu_Init(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenu* result_ = [nSMenu init];
    return result_;
}

void* C_NSMenu_NewMenu() {
    NSMenu* result_ = [NSMenu new];
    return result_;
}

void* C_NSMenu_Autorelease(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenu* result_ = [nSMenu autorelease];
    return result_;
}

void* C_NSMenu_Retain(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenu* result_ = [nSMenu retain];
    return result_;
}

bool C_NSMenu_MenuBarVisible() {
    BOOL result_ = [NSMenu menuBarVisible];
    return result_;
}

void C_NSMenu_SetMenuBarVisible(bool visible) {
    [NSMenu setMenuBarVisible:visible];
}

void C_NSMenu_InsertItem_AtIndex(void* ptr, void* newItem, int index) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu insertItem:(NSMenuItem*)newItem atIndex:index];
}

void* C_NSMenu_InsertItemWithTitle_Action_KeyEquivalent_AtIndex(void* ptr, void* _string, void* selector, void* charCode, int index) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuItem* result_ = [nSMenu insertItemWithTitle:(NSString*)_string action:(SEL)selector keyEquivalent:(NSString*)charCode atIndex:index];
    return result_;
}

void C_NSMenu_AddItem(void* ptr, void* newItem) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu addItem:(NSMenuItem*)newItem];
}

void* C_NSMenu_AddItemWithTitle_Action_KeyEquivalent(void* ptr, void* _string, void* selector, void* charCode) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuItem* result_ = [nSMenu addItemWithTitle:(NSString*)_string action:(SEL)selector keyEquivalent:(NSString*)charCode];
    return result_;
}

void C_NSMenu_RemoveItem(void* ptr, void* item) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu removeItem:(NSMenuItem*)item];
}

void C_NSMenu_RemoveItemAtIndex(void* ptr, int index) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu removeItemAtIndex:index];
}

void C_NSMenu_ItemChanged(void* ptr, void* item) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu itemChanged:(NSMenuItem*)item];
}

void C_NSMenu_RemoveAllItems(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu removeAllItems];
}

void* C_NSMenu_ItemWithTag(void* ptr, int tag) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuItem* result_ = [nSMenu itemWithTag:tag];
    return result_;
}

void* C_NSMenu_ItemWithTitle(void* ptr, void* title) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuItem* result_ = [nSMenu itemWithTitle:(NSString*)title];
    return result_;
}

void* C_NSMenu_ItemAtIndex(void* ptr, int index) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuItem* result_ = [nSMenu itemAtIndex:index];
    return result_;
}

int C_NSMenu_IndexOfItem(void* ptr, void* item) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result_ = [nSMenu indexOfItem:(NSMenuItem*)item];
    return result_;
}

int C_NSMenu_IndexOfItemWithTitle(void* ptr, void* title) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result_ = [nSMenu indexOfItemWithTitle:(NSString*)title];
    return result_;
}

int C_NSMenu_IndexOfItemWithTag(void* ptr, int tag) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result_ = [nSMenu indexOfItemWithTag:tag];
    return result_;
}

int C_NSMenu_IndexOfItemWithTarget_AndAction(void* ptr, void* target, void* actionSelector) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result_ = [nSMenu indexOfItemWithTarget:(id)target andAction:(SEL)actionSelector];
    return result_;
}

int C_NSMenu_IndexOfItemWithRepresentedObject(void* ptr, void* object) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result_ = [nSMenu indexOfItemWithRepresentedObject:(id)object];
    return result_;
}

int C_NSMenu_IndexOfItemWithSubmenu(void* ptr, void* submenu) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result_ = [nSMenu indexOfItemWithSubmenu:(NSMenu*)submenu];
    return result_;
}

void C_NSMenu_SetSubmenu_ForItem(void* ptr, void* menu, void* item) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setSubmenu:(NSMenu*)menu forItem:(NSMenuItem*)item];
}

void C_NSMenu_SubmenuAction(void* ptr, void* sender) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu submenuAction:(id)sender];
}

void C_NSMenu_Update(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu update];
}

bool C_NSMenu_PerformKeyEquivalent(void* ptr, void* event) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    BOOL result_ = [nSMenu performKeyEquivalent:(NSEvent*)event];
    return result_;
}

void C_NSMenu_PerformActionForItemAtIndex(void* ptr, int index) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu performActionForItemAtIndex:index];
}

void C_NSMenu_PopUpContextMenu_WithEvent_ForView(void* menu, void* event, void* view) {
    [NSMenu popUpContextMenu:(NSMenu*)menu withEvent:(NSEvent*)event forView:(NSView*)view];
}

void C_NSMenu_PopUpContextMenu_WithEvent_ForView_WithFont(void* menu, void* event, void* view, void* font) {
    [NSMenu popUpContextMenu:(NSMenu*)menu withEvent:(NSEvent*)event forView:(NSView*)view withFont:(NSFont*)font];
}

bool C_NSMenu_PopUpMenuPositioningItem_AtLocation_InView(void* ptr, void* item, CGPoint location, void* view) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    BOOL result_ = [nSMenu popUpMenuPositioningItem:(NSMenuItem*)item atLocation:location inView:(NSView*)view];
    return result_;
}

void C_NSMenu_CancelTracking(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu cancelTracking];
}

void C_NSMenu_CancelTrackingWithoutAnimation(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu cancelTrackingWithoutAnimation];
}

double C_NSMenu_MenuBarHeight(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    CGFloat result_ = [nSMenu menuBarHeight];
    return result_;
}

int C_NSMenu_NumberOfItems(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result_ = [nSMenu numberOfItems];
    return result_;
}

Array C_NSMenu_ItemArray(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSArray* result_ = [nSMenu itemArray];
    Array result_Array;
    int result_count = [result_ count];
    if (result_count > 0) {
    	void** result_Data = malloc(result_count * sizeof(void*));
    	for (int i = 0; i < result_count; i++) {
    		 void* p = [result_ objectAtIndex:i];
    		 result_Data[i] = p;
    	}
    	result_Array.data = result_Data;
    	result_Array.len = result_count;
    }
    return result_Array;
}

void C_NSMenu_SetItemArray(void* ptr, Array value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMutableArray* objcValue = [NSMutableArray arrayWithCapacity:value.len];
    if (value.len > 0) {
    	void** valueData = (void**)value.data;
    	for (int i = 0; i < value.len; i++) {
    		void* p = valueData[i];
    		[objcValue addObject:(NSMenuItem*)p];
    	}
    }
    [nSMenu setItemArray:objcValue];
}

void* C_NSMenu_Supermenu(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenu* result_ = [nSMenu supermenu];
    return result_;
}

void C_NSMenu_SetSupermenu(void* ptr, void* value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setSupermenu:(NSMenu*)value];
}

bool C_NSMenu_AutoenablesItems(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    BOOL result_ = [nSMenu autoenablesItems];
    return result_;
}

void C_NSMenu_SetAutoenablesItems(void* ptr, bool value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setAutoenablesItems:value];
}

void* C_NSMenu_Font(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSFont* result_ = [nSMenu font];
    return result_;
}

void C_NSMenu_SetFont(void* ptr, void* value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setFont:(NSFont*)value];
}

void* C_NSMenu_Title(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSString* result_ = [nSMenu title];
    return result_;
}

void C_NSMenu_SetTitle(void* ptr, void* value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setTitle:(NSString*)value];
}

double C_NSMenu_MinimumWidth(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    CGFloat result_ = [nSMenu minimumWidth];
    return result_;
}

void C_NSMenu_SetMinimumWidth(void* ptr, double value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setMinimumWidth:value];
}

CGSize C_NSMenu_Size(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSSize result_ = [nSMenu size];
    return result_;
}

unsigned int C_NSMenu_PropertiesToUpdate(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuProperties result_ = [nSMenu propertiesToUpdate];
    return result_;
}

bool C_NSMenu_AllowsContextMenuPlugIns(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    BOOL result_ = [nSMenu allowsContextMenuPlugIns];
    return result_;
}

void C_NSMenu_SetAllowsContextMenuPlugIns(void* ptr, bool value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setAllowsContextMenuPlugIns:value];
}

bool C_NSMenu_ShowsStateColumn(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    BOOL result_ = [nSMenu showsStateColumn];
    return result_;
}

void C_NSMenu_SetShowsStateColumn(void* ptr, bool value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setShowsStateColumn:value];
}

void* C_NSMenu_HighlightedItem(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuItem* result_ = [nSMenu highlightedItem];
    return result_;
}

int C_NSMenu_UserInterfaceLayoutDirection(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSUserInterfaceLayoutDirection result_ = [nSMenu userInterfaceLayoutDirection];
    return result_;
}

void C_NSMenu_SetUserInterfaceLayoutDirection(void* ptr, int value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setUserInterfaceLayoutDirection:value];
}

void* C_NSMenu_Delegate(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    id result_ = [nSMenu delegate];
    return result_;
}

void C_NSMenu_SetDelegate(void* ptr, void* value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setDelegate:(id)value];
}

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
