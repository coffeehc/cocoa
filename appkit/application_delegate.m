#import <Appkit/Appkit.h>
#import "application_delegate.h"
#import "_cgo_export.h"

@implementation NSApplicationDelegateAdaptor


- (void)applicationWillFinishLaunching:(NSNotification*)notification {
    ApplicationWillFinishLaunching([self goID], notification);
}

- (void)applicationDidFinishLaunching:(NSNotification*)notification {
    ApplicationDidFinishLaunching([self goID], notification);
}

- (void)applicationWillBecomeActive:(NSNotification*)notification {
    ApplicationWillBecomeActive([self goID], notification);
}

- (void)applicationDidBecomeActive:(NSNotification*)notification {
    ApplicationDidBecomeActive([self goID], notification);
}

- (void)applicationWillResignActive:(NSNotification*)notification {
    ApplicationWillResignActive([self goID], notification);
}

- (void)applicationDidResignActive:(NSNotification*)notification {
    ApplicationDidResignActive([self goID], notification);
}

- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication*)sender {
    unsigned int result = ApplicationShouldTerminate([self goID], sender);
    return result;
}

- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication*)sender {
    bool result = ApplicationShouldTerminateAfterLastWindowClosed([self goID], sender);
    return result;
}

- (void)applicationWillTerminate:(NSNotification*)notification {
    ApplicationWillTerminate([self goID], notification);
}

- (void)applicationWillHide:(NSNotification*)notification {
    ApplicationWillHide([self goID], notification);
}

- (void)applicationDidHide:(NSNotification*)notification {
    ApplicationDidHide([self goID], notification);
}

- (void)applicationWillUnhide:(NSNotification*)notification {
    ApplicationWillUnhide([self goID], notification);
}

- (void)applicationDidUnhide:(NSNotification*)notification {
    ApplicationDidUnhide([self goID], notification);
}

- (void)applicationWillUpdate:(NSNotification*)notification {
    ApplicationWillUpdate([self goID], notification);
}

- (void)applicationDidUpdate:(NSNotification*)notification {
    ApplicationDidUpdate([self goID], notification);
}

- (BOOL)applicationShouldHandleReopen:(NSApplication*)sender hasVisibleWindows:(BOOL)flag {
    bool result = ApplicationShouldHandleReopen_HasVisibleWindows([self goID], sender, flag);
    return result;
}

- (NSMenu*)applicationDockMenu:(NSApplication*)sender {
    void* result = ApplicationDockMenu([self goID], sender);
    return (NSMenu*)result;
}

- (NSError*)application:(NSApplication*)application willPresentError:(NSError*)error {
    void* result = Application_WillPresentError([self goID], application, error);
    return (NSError*)result;
}

- (void)applicationDidChangeScreenParameters:(NSNotification*)notification {
    ApplicationDidChangeScreenParameters([self goID], notification);
}

- (BOOL)application:(NSApplication*)application willContinueUserActivityWithType:(NSString*)userActivityType {
    bool result = Application_WillContinueUserActivityWithType([self goID], application, userActivityType);
    return result;
}

- (void)application:(NSApplication*)application didFailToContinueUserActivityWithType:(NSString*)userActivityType error:(NSError*)error {
    Application_DidFailToContinueUserActivityWithType_Error([self goID], application, userActivityType, error);
}

- (void)application:(NSApplication*)application didUpdateUserActivity:(NSUserActivity*)userActivity {
    Application_DidUpdateUserActivity([self goID], application, userActivity);
}

- (void)application:(NSApplication*)application didRegisterForRemoteNotificationsWithDeviceToken:(NSData*)deviceToken {
    Array deviceTokenarray;
    deviceTokenarray.data = [deviceToken bytes];
    deviceTokenarray.len = deviceToken.length;
    Application_DidRegisterForRemoteNotificationsWithDeviceToken([self goID], application, deviceTokenarray);
}

- (void)application:(NSApplication*)application didFailToRegisterForRemoteNotificationsWithError:(NSError*)error {
    Application_DidFailToRegisterForRemoteNotificationsWithError([self goID], application, error);
}

- (BOOL)application:(NSApplication*)sender openFile:(NSString*)filename {
    bool result = Application_OpenFile([self goID], sender, filename);
    return result;
}

- (BOOL)application:(id)sender openFileWithoutUI:(NSString*)filename {
    bool result = Application_OpenFileWithoutUI([self goID], sender, filename);
    return result;
}

- (BOOL)application:(NSApplication*)sender openTempFile:(NSString*)filename {
    bool result = Application_OpenTempFile([self goID], sender, filename);
    return result;
}

- (BOOL)applicationOpenUntitledFile:(NSApplication*)sender {
    bool result = ApplicationOpenUntitledFile([self goID], sender);
    return result;
}

- (BOOL)applicationShouldOpenUntitledFile:(NSApplication*)sender {
    bool result = ApplicationShouldOpenUntitledFile([self goID], sender);
    return result;
}

- (BOOL)application:(NSApplication*)sender printFile:(NSString*)filename {
    bool result = Application_PrintFile([self goID], sender, filename);
    return result;
}

- (void)application:(NSApplication*)app didDecodeRestorableState:(NSCoder*)coder {
    Application_DidDecodeRestorableState([self goID], app, coder);
}

- (void)application:(NSApplication*)app willEncodeRestorableState:(NSCoder*)coder {
    Application_WillEncodeRestorableState([self goID], app, coder);
}

- (void)applicationDidChangeOcclusionState:(NSNotification*)notification {
    ApplicationDidChangeOcclusionState([self goID], notification);
}

- (BOOL)application:(NSApplication*)sender delegateHandlesKey:(NSString*)key {
    bool result = Application_DelegateHandlesKey([self goID], sender, key);
    return result;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return ApplicationDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteApplicationDelegate([self goID]);
	[super dealloc];
}
@end

void* WrapApplicationDelegate(long goID) {
    NSApplicationDelegateAdaptor* adaptor = [[NSApplicationDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
