#import <Appkit/Appkit.h>
#import "menu.h"

void* C_Menu_Alloc() {
    return [NSMenu alloc];
}

void* C_NSMenu_InitWithTitle(void* ptr, void* title) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenu* result = [nSMenu initWithTitle:(NSString*)title];
    return result;
}

void* C_NSMenu_InitWithCoder(void* ptr, void* coder) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenu* result = [nSMenu initWithCoder:(NSCoder*)coder];
    return result;
}

void* C_NSMenu_Init(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenu* result = [nSMenu init];
    return result;
}

bool C_NSMenu_MenuBarVisible() {
    BOOL result = [NSMenu menuBarVisible];
    return result;
}

void C_NSMenu_Menu_SetMenuBarVisible(bool visible) {
    [NSMenu setMenuBarVisible:visible];
}

void C_NSMenu_InsertItem_AtIndex(void* ptr, void* newItem, int index) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu insertItem:(NSMenuItem*)newItem atIndex:index];
}

void* C_NSMenu_InsertItemWithTitle_Action_KeyEquivalent_AtIndex(void* ptr, void* _string, void* selector, void* charCode, int index) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuItem* result = [nSMenu insertItemWithTitle:(NSString*)_string action:(SEL)selector keyEquivalent:(NSString*)charCode atIndex:index];
    return result;
}

void C_NSMenu_AddItem(void* ptr, void* newItem) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu addItem:(NSMenuItem*)newItem];
}

void* C_NSMenu_AddItemWithTitle_Action_KeyEquivalent(void* ptr, void* _string, void* selector, void* charCode) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuItem* result = [nSMenu addItemWithTitle:(NSString*)_string action:(SEL)selector keyEquivalent:(NSString*)charCode];
    return result;
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
    NSMenuItem* result = [nSMenu itemWithTag:tag];
    return result;
}

void* C_NSMenu_ItemWithTitle(void* ptr, void* title) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuItem* result = [nSMenu itemWithTitle:(NSString*)title];
    return result;
}

void* C_NSMenu_ItemAtIndex(void* ptr, int index) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuItem* result = [nSMenu itemAtIndex:index];
    return result;
}

int C_NSMenu_IndexOfItem(void* ptr, void* item) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result = [nSMenu indexOfItem:(NSMenuItem*)item];
    return result;
}

int C_NSMenu_IndexOfItemWithTitle(void* ptr, void* title) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result = [nSMenu indexOfItemWithTitle:(NSString*)title];
    return result;
}

int C_NSMenu_IndexOfItemWithTag(void* ptr, int tag) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result = [nSMenu indexOfItemWithTag:tag];
    return result;
}

int C_NSMenu_IndexOfItemWithTarget_AndAction(void* ptr, void* target, void* actionSelector) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result = [nSMenu indexOfItemWithTarget:(id)target andAction:(SEL)actionSelector];
    return result;
}

int C_NSMenu_IndexOfItemWithRepresentedObject(void* ptr, void* object) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result = [nSMenu indexOfItemWithRepresentedObject:(id)object];
    return result;
}

int C_NSMenu_IndexOfItemWithSubmenu(void* ptr, void* submenu) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result = [nSMenu indexOfItemWithSubmenu:(NSMenu*)submenu];
    return result;
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
    BOOL result = [nSMenu performKeyEquivalent:(NSEvent*)event];
    return result;
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
    BOOL result = [nSMenu popUpMenuPositioningItem:(NSMenuItem*)item atLocation:location inView:(NSView*)view];
    return result;
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
    CGFloat result = [nSMenu menuBarHeight];
    return result;
}

int C_NSMenu_NumberOfItems(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSInteger result = [nSMenu numberOfItems];
    return result;
}

Array C_NSMenu_ItemArray(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSArray* result = [nSMenu itemArray];
    int resultcount = [result count];
    void** resultData = malloc(resultcount * sizeof(void*));
    for (int i = 0; i < resultcount; i++) {
    	 void* p = [result objectAtIndex:i];
    	 resultData[i] = p;
    }
    Array resultArray;
    resultArray.data = resultData;
    resultArray.len = resultcount;
    return resultArray;
}

void C_NSMenu_SetItemArray(void* ptr, Array value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMutableArray* objcValue = [[NSMutableArray alloc] init];
    void** valueData = (void**)value.data;
    for (int i = 0; i < value.len; i++) {
    	void* p = valueData[i];
    	[objcValue addObject:(NSMenuItem*)(NSMenuItem*)p];
    }
    [nSMenu setItemArray:objcValue];
}

void* C_NSMenu_Supermenu(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenu* result = [nSMenu supermenu];
    return result;
}

void C_NSMenu_SetSupermenu(void* ptr, void* value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setSupermenu:(NSMenu*)value];
}

bool C_NSMenu_AutoenablesItems(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    BOOL result = [nSMenu autoenablesItems];
    return result;
}

void C_NSMenu_SetAutoenablesItems(void* ptr, bool value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setAutoenablesItems:value];
}

void* C_NSMenu_Font(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSFont* result = [nSMenu font];
    return result;
}

void C_NSMenu_SetFont(void* ptr, void* value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setFont:(NSFont*)value];
}

void* C_NSMenu_Title(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSString* result = [nSMenu title];
    return result;
}

void C_NSMenu_SetTitle(void* ptr, void* value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setTitle:(NSString*)value];
}

double C_NSMenu_MinimumWidth(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    CGFloat result = [nSMenu minimumWidth];
    return result;
}

void C_NSMenu_SetMinimumWidth(void* ptr, double value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setMinimumWidth:value];
}

CGSize C_NSMenu_Size(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSSize result = [nSMenu size];
    return result;
}

unsigned int C_NSMenu_PropertiesToUpdate(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuProperties result = [nSMenu propertiesToUpdate];
    return result;
}

bool C_NSMenu_AllowsContextMenuPlugIns(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    BOOL result = [nSMenu allowsContextMenuPlugIns];
    return result;
}

void C_NSMenu_SetAllowsContextMenuPlugIns(void* ptr, bool value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setAllowsContextMenuPlugIns:value];
}

bool C_NSMenu_ShowsStateColumn(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    BOOL result = [nSMenu showsStateColumn];
    return result;
}

void C_NSMenu_SetShowsStateColumn(void* ptr, bool value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setShowsStateColumn:value];
}

void* C_NSMenu_HighlightedItem(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSMenuItem* result = [nSMenu highlightedItem];
    return result;
}

int C_NSMenu_UserInterfaceLayoutDirection(void* ptr) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    NSUserInterfaceLayoutDirection result = [nSMenu userInterfaceLayoutDirection];
    return result;
}

void C_NSMenu_SetUserInterfaceLayoutDirection(void* ptr, int value) {
    NSMenu* nSMenu = (NSMenu*)ptr;
    [nSMenu setUserInterfaceLayoutDirection:value];
}
