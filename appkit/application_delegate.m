#import <Appkit/Appkit.h>
#import "application_delegate.h"
#import "_cgo_export.h"

@implementation NSApplicationDelegateAdaptor


- (void)applicationWillFinishLaunching:(NSNotification*)notification {
    applicationDelegate_ApplicationWillFinishLaunching([self goID], notification);
}

- (void)applicationDidFinishLaunching:(NSNotification*)notification {
    applicationDelegate_ApplicationDidFinishLaunching([self goID], notification);
}

- (void)applicationWillBecomeActive:(NSNotification*)notification {
    applicationDelegate_ApplicationWillBecomeActive([self goID], notification);
}

- (void)applicationDidBecomeActive:(NSNotification*)notification {
    applicationDelegate_ApplicationDidBecomeActive([self goID], notification);
}

- (void)applicationWillResignActive:(NSNotification*)notification {
    applicationDelegate_ApplicationWillResignActive([self goID], notification);
}

- (void)applicationDidResignActive:(NSNotification*)notification {
    applicationDelegate_ApplicationDidResignActive([self goID], notification);
}

- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication*)sender {
    unsigned int result_ = applicationDelegate_ApplicationShouldTerminate([self goID], sender);
    return result_;
}

- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication*)sender {
    bool result_ = applicationDelegate_ApplicationShouldTerminateAfterLastWindowClosed([self goID], sender);
    return result_;
}

- (void)applicationWillTerminate:(NSNotification*)notification {
    applicationDelegate_ApplicationWillTerminate([self goID], notification);
}

- (void)applicationWillHide:(NSNotification*)notification {
    applicationDelegate_ApplicationWillHide([self goID], notification);
}

- (void)applicationDidHide:(NSNotification*)notification {
    applicationDelegate_ApplicationDidHide([self goID], notification);
}

- (void)applicationWillUnhide:(NSNotification*)notification {
    applicationDelegate_ApplicationWillUnhide([self goID], notification);
}

- (void)applicationDidUnhide:(NSNotification*)notification {
    applicationDelegate_ApplicationDidUnhide([self goID], notification);
}

- (void)applicationWillUpdate:(NSNotification*)notification {
    applicationDelegate_ApplicationWillUpdate([self goID], notification);
}

- (void)applicationDidUpdate:(NSNotification*)notification {
    applicationDelegate_ApplicationDidUpdate([self goID], notification);
}

- (BOOL)applicationShouldHandleReopen:(NSApplication*)sender hasVisibleWindows:(BOOL)flag {
    bool result_ = applicationDelegate_ApplicationShouldHandleReopen_HasVisibleWindows([self goID], sender, flag);
    return result_;
}

- (NSMenu*)applicationDockMenu:(NSApplication*)sender {
    void* result_ = applicationDelegate_ApplicationDockMenu([self goID], sender);
    return (NSMenu*)result_;
}

- (NSError*)application:(NSApplication*)application willPresentError:(NSError*)error {
    void* result_ = applicationDelegate_Application_WillPresentError([self goID], application, error);
    return (NSError*)result_;
}

- (void)applicationDidChangeScreenParameters:(NSNotification*)notification {
    applicationDelegate_ApplicationDidChangeScreenParameters([self goID], notification);
}

- (BOOL)application:(NSApplication*)application willContinueUserActivityWithType:(NSString*)userActivityType {
    bool result_ = applicationDelegate_Application_WillContinueUserActivityWithType([self goID], application, userActivityType);
    return result_;
}

- (void)application:(NSApplication*)application didFailToContinueUserActivityWithType:(NSString*)userActivityType error:(NSError*)error {
    applicationDelegate_Application_DidFailToContinueUserActivityWithType_Error([self goID], application, userActivityType, error);
}

- (void)application:(NSApplication*)application didUpdateUserActivity:(NSUserActivity*)userActivity {
    applicationDelegate_Application_DidUpdateUserActivity([self goID], application, userActivity);
}

- (void)application:(NSApplication*)application didRegisterForRemoteNotificationsWithDeviceToken:(NSData*)deviceToken {
    Array deviceTokenarray;
    deviceTokenarray.data = [deviceToken bytes];
    deviceTokenarray.len = deviceToken.length;
    applicationDelegate_Application_DidRegisterForRemoteNotificationsWithDeviceToken([self goID], application, deviceTokenarray);
}

- (void)application:(NSApplication*)application didFailToRegisterForRemoteNotificationsWithError:(NSError*)error {
    applicationDelegate_Application_DidFailToRegisterForRemoteNotificationsWithError([self goID], application, error);
}

- (void)application:(NSApplication*)application didReceiveRemoteNotification:(NSDictionary*)userInfo {
    Dictionary userInfoArray;
    NSArray * userInfoKeys = [userInfo allKeys];
    int userInfoCount = [userInfoKeys count];
    if (userInfoCount > 0) {
    	void** userInfoKeyData = malloc(userInfoCount * sizeof(void*));
    	void** userInfoValueData = malloc(userInfoCount * sizeof(void*));
    	for (int i = 0; i < userInfoCount; i++) {
    		NSString* kp = [userInfoKeys objectAtIndex:i];
    		id vp = userInfo[kp];
    		 userInfoKeyData[i] = kp;
    		 userInfoValueData[i] = vp;
    	}
    	userInfoArray.key_data = userInfoKeyData;
    	userInfoArray.value_data = userInfoValueData;
    	userInfoArray.len = userInfoCount;
    }
    applicationDelegate_Application_DidReceiveRemoteNotification([self goID], application, userInfoArray);
}

- (void)application:(NSApplication*)application openURLs:(NSArray*)urls {
    Array urlsArray;
    int urlscount = [urls count];
    if (urlscount > 0) {
    	void** urlsData = malloc(urlscount * sizeof(void*));
    	for (int i = 0; i < urlscount; i++) {
    		 void* p = [urls objectAtIndex:i];
    		 urlsData[i] = p;
    	}
    	urlsArray.data = urlsData;
    	urlsArray.len = urlscount;
    }
    applicationDelegate_Application_OpenURLs([self goID], application, urlsArray);
}

- (BOOL)application:(NSApplication*)sender openFile:(NSString*)filename {
    bool result_ = applicationDelegate_Application_OpenFile([self goID], sender, filename);
    return result_;
}

- (BOOL)application:(id)sender openFileWithoutUI:(NSString*)filename {
    bool result_ = applicationDelegate_Application_OpenFileWithoutUI([self goID], sender, filename);
    return result_;
}

- (BOOL)application:(NSApplication*)sender openTempFile:(NSString*)filename {
    bool result_ = applicationDelegate_Application_OpenTempFile([self goID], sender, filename);
    return result_;
}

- (void)application:(NSApplication*)sender openFiles:(NSArray*)filenames {
    Array filenamesArray;
    int filenamescount = [filenames count];
    if (filenamescount > 0) {
    	void** filenamesData = malloc(filenamescount * sizeof(void*));
    	for (int i = 0; i < filenamescount; i++) {
    		 void* p = [filenames objectAtIndex:i];
    		 filenamesData[i] = p;
    	}
    	filenamesArray.data = filenamesData;
    	filenamesArray.len = filenamescount;
    }
    applicationDelegate_Application_OpenFiles([self goID], sender, filenamesArray);
}

- (BOOL)applicationOpenUntitledFile:(NSApplication*)sender {
    bool result_ = applicationDelegate_ApplicationOpenUntitledFile([self goID], sender);
    return result_;
}

- (BOOL)applicationShouldOpenUntitledFile:(NSApplication*)sender {
    bool result_ = applicationDelegate_ApplicationShouldOpenUntitledFile([self goID], sender);
    return result_;
}

- (BOOL)application:(NSApplication*)sender printFile:(NSString*)filename {
    bool result_ = applicationDelegate_Application_PrintFile([self goID], sender, filename);
    return result_;
}

- (NSApplicationPrintReply)application:(NSApplication*)application printFiles:(NSArray*)fileNames withSettings:(NSDictionary*)printSettings showPrintPanels:(BOOL)showPrintPanels {
    Array fileNamesArray;
    int fileNamescount = [fileNames count];
    if (fileNamescount > 0) {
    	void** fileNamesData = malloc(fileNamescount * sizeof(void*));
    	for (int i = 0; i < fileNamescount; i++) {
    		 void* p = [fileNames objectAtIndex:i];
    		 fileNamesData[i] = p;
    	}
    	fileNamesArray.data = fileNamesData;
    	fileNamesArray.len = fileNamescount;
    }
    Dictionary printSettingsArray;
    NSArray * printSettingsKeys = [printSettings allKeys];
    int printSettingsCount = [printSettingsKeys count];
    if (printSettingsCount > 0) {
    	void** printSettingsKeyData = malloc(printSettingsCount * sizeof(void*));
    	void** printSettingsValueData = malloc(printSettingsCount * sizeof(void*));
    	for (int i = 0; i < printSettingsCount; i++) {
    		NSPrintInfoAttributeKey kp = [printSettingsKeys objectAtIndex:i];
    		id vp = printSettings[kp];
    		 printSettingsKeyData[i] = kp;
    		 printSettingsValueData[i] = vp;
    	}
    	printSettingsArray.key_data = printSettingsKeyData;
    	printSettingsArray.value_data = printSettingsValueData;
    	printSettingsArray.len = printSettingsCount;
    }
    unsigned int result_ = applicationDelegate_Application_PrintFiles_WithSettings_ShowPrintPanels([self goID], application, fileNamesArray, printSettingsArray, showPrintPanels);
    return result_;
}

- (void)application:(NSApplication*)app didDecodeRestorableState:(NSCoder*)coder {
    applicationDelegate_Application_DidDecodeRestorableState([self goID], app, coder);
}

- (void)application:(NSApplication*)app willEncodeRestorableState:(NSCoder*)coder {
    applicationDelegate_Application_WillEncodeRestorableState([self goID], app, coder);
}

- (void)applicationDidChangeOcclusionState:(NSNotification*)notification {
    applicationDelegate_ApplicationDidChangeOcclusionState([self goID], notification);
}

- (BOOL)application:(NSApplication*)sender delegateHandlesKey:(NSString*)key {
    bool result_ = applicationDelegate_Application_DelegateHandlesKey([self goID], sender, key);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return ApplicationDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapApplicationDelegate(uintptr_t goID) {
    NSApplicationDelegateAdaptor* adaptor = [[NSApplicationDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
