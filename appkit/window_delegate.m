#import <Appkit/Appkit.h>
#import "window_delegate.h"
#import "_cgo_export.h"

@implementation NSWindowDelegateAdaptor


- (NSRect)window:(NSWindow*)window willPositionSheet:(NSWindow*)sheet usingRect:(NSRect)rect {
    CGRect result_ = windowDelegate_Window_WillPositionSheet_UsingRect([self goID], window, sheet, rect);
    return result_;
}

- (void)windowWillBeginSheet:(NSNotification*)notification {
    windowDelegate_WindowWillBeginSheet([self goID], notification);
}

- (void)windowDidEndSheet:(NSNotification*)notification {
    windowDelegate_WindowDidEndSheet([self goID], notification);
}

- (NSSize)windowWillResize:(NSWindow*)sender toSize:(NSSize)frameSize {
    CGSize result_ = windowDelegate_WindowWillResize_ToSize([self goID], sender, frameSize);
    return result_;
}

- (void)windowDidResize:(NSNotification*)notification {
    windowDelegate_WindowDidResize([self goID], notification);
}

- (void)windowWillStartLiveResize:(NSNotification*)notification {
    windowDelegate_WindowWillStartLiveResize([self goID], notification);
}

- (void)windowDidEndLiveResize:(NSNotification*)notification {
    windowDelegate_WindowDidEndLiveResize([self goID], notification);
}

- (void)windowWillMiniaturize:(NSNotification*)notification {
    windowDelegate_WindowWillMiniaturize([self goID], notification);
}

- (void)windowDidMiniaturize:(NSNotification*)notification {
    windowDelegate_WindowDidMiniaturize([self goID], notification);
}

- (void)windowDidDeminiaturize:(NSNotification*)notification {
    windowDelegate_WindowDidDeminiaturize([self goID], notification);
}

- (NSRect)windowWillUseStandardFrame:(NSWindow*)window defaultFrame:(NSRect)newFrame {
    CGRect result_ = windowDelegate_WindowWillUseStandardFrame_DefaultFrame([self goID], window, newFrame);
    return result_;
}

- (BOOL)windowShouldZoom:(NSWindow*)window toFrame:(NSRect)newFrame {
    bool result_ = windowDelegate_WindowShouldZoom_ToFrame([self goID], window, newFrame);
    return result_;
}

- (NSSize)window:(NSWindow*)window willUseFullScreenContentSize:(NSSize)proposedSize {
    CGSize result_ = windowDelegate_Window_WillUseFullScreenContentSize([self goID], window, proposedSize);
    return result_;
}

- (NSApplicationPresentationOptions)window:(NSWindow*)window willUseFullScreenPresentationOptions:(NSApplicationPresentationOptions)proposedOptions {
    unsigned int result_ = windowDelegate_Window_WillUseFullScreenPresentationOptions([self goID], window, proposedOptions);
    return result_;
}

- (void)windowWillEnterFullScreen:(NSNotification*)notification {
    windowDelegate_WindowWillEnterFullScreen([self goID], notification);
}

- (void)windowDidEnterFullScreen:(NSNotification*)notification {
    windowDelegate_WindowDidEnterFullScreen([self goID], notification);
}

- (void)windowWillExitFullScreen:(NSNotification*)notification {
    windowDelegate_WindowWillExitFullScreen([self goID], notification);
}

- (void)windowDidExitFullScreen:(NSNotification*)notification {
    windowDelegate_WindowDidExitFullScreen([self goID], notification);
}

- (NSArray*)customWindowsToEnterFullScreenForWindow:(NSWindow*)window {
    Array result_ = windowDelegate_CustomWindowsToEnterFullScreenForWindow([self goID], window);
    NSMutableArray* objcResult_ = [[NSMutableArray alloc] init];
    void** result_Data = (void**)result_.data;
    for (int i = 0; i < result_.len; i++) {
    	void* p = result_Data[i];
    	[objcResult_ addObject:(NSWindow*)(NSWindow*)p];
    }
    return objcResult_;
}

- (NSArray*)customWindowsToEnterFullScreenForWindow:(NSWindow*)window onScreen:(NSScreen*)screen {
    Array result_ = windowDelegate_CustomWindowsToEnterFullScreenForWindow_OnScreen([self goID], window, screen);
    NSMutableArray* objcResult_ = [[NSMutableArray alloc] init];
    void** result_Data = (void**)result_.data;
    for (int i = 0; i < result_.len; i++) {
    	void* p = result_Data[i];
    	[objcResult_ addObject:(NSWindow*)(NSWindow*)p];
    }
    return objcResult_;
}

- (void)window:(NSWindow*)window startCustomAnimationToEnterFullScreenWithDuration:(NSTimeInterval)duration {
    windowDelegate_Window_StartCustomAnimationToEnterFullScreenWithDuration([self goID], window, duration);
}

- (void)window:(NSWindow*)window startCustomAnimationToEnterFullScreenOnScreen:(NSScreen*)screen withDuration:(NSTimeInterval)duration {
    windowDelegate_Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration([self goID], window, screen, duration);
}

- (void)windowDidFailToEnterFullScreen:(NSWindow*)window {
    windowDelegate_WindowDidFailToEnterFullScreen([self goID], window);
}

- (NSArray*)customWindowsToExitFullScreenForWindow:(NSWindow*)window {
    Array result_ = windowDelegate_CustomWindowsToExitFullScreenForWindow([self goID], window);
    NSMutableArray* objcResult_ = [[NSMutableArray alloc] init];
    void** result_Data = (void**)result_.data;
    for (int i = 0; i < result_.len; i++) {
    	void* p = result_Data[i];
    	[objcResult_ addObject:(NSWindow*)(NSWindow*)p];
    }
    return objcResult_;
}

- (void)window:(NSWindow*)window startCustomAnimationToExitFullScreenWithDuration:(NSTimeInterval)duration {
    windowDelegate_Window_StartCustomAnimationToExitFullScreenWithDuration([self goID], window, duration);
}

- (void)windowDidFailToExitFullScreen:(NSWindow*)window {
    windowDelegate_WindowDidFailToExitFullScreen([self goID], window);
}

- (void)windowWillMove:(NSNotification*)notification {
    windowDelegate_WindowWillMove([self goID], notification);
}

- (void)windowDidMove:(NSNotification*)notification {
    windowDelegate_WindowDidMove([self goID], notification);
}

- (void)windowDidChangeScreen:(NSNotification*)notification {
    windowDelegate_WindowDidChangeScreen([self goID], notification);
}

- (void)windowDidChangeScreenProfile:(NSNotification*)notification {
    windowDelegate_WindowDidChangeScreenProfile([self goID], notification);
}

- (void)windowDidChangeBackingProperties:(NSNotification*)notification {
    windowDelegate_WindowDidChangeBackingProperties([self goID], notification);
}

- (BOOL)windowShouldClose:(NSWindow*)sender {
    bool result_ = windowDelegate_WindowShouldClose([self goID], sender);
    return result_;
}

- (void)windowWillClose:(NSNotification*)notification {
    windowDelegate_WindowWillClose([self goID], notification);
}

- (void)windowDidBecomeKey:(NSNotification*)notification {
    windowDelegate_WindowDidBecomeKey([self goID], notification);
}

- (void)windowDidResignKey:(NSNotification*)notification {
    windowDelegate_WindowDidResignKey([self goID], notification);
}

- (void)windowDidBecomeMain:(NSNotification*)notification {
    windowDelegate_WindowDidBecomeMain([self goID], notification);
}

- (void)windowDidResignMain:(NSNotification*)notification {
    windowDelegate_WindowDidResignMain([self goID], notification);
}

- (id)windowWillReturnFieldEditor:(NSWindow*)sender toObject:(id)client {
    void* result_ = windowDelegate_WindowWillReturnFieldEditor_ToObject([self goID], sender, client);
    return (id)result_;
}

- (void)windowDidUpdate:(NSNotification*)notification {
    windowDelegate_WindowDidUpdate([self goID], notification);
}

- (void)windowDidExpose:(NSNotification*)notification {
    windowDelegate_WindowDidExpose([self goID], notification);
}

- (void)windowDidChangeOcclusionState:(NSNotification*)notification {
    windowDelegate_WindowDidChangeOcclusionState([self goID], notification);
}

- (BOOL)window:(NSWindow*)window shouldDragDocumentWithEvent:(NSEvent*)event from:(NSPoint)dragImageLocation withPasteboard:(NSPasteboard*)pasteboard {
    bool result_ = windowDelegate_Window_ShouldDragDocumentWithEvent_From_WithPasteboard([self goID], window, event, dragImageLocation, pasteboard);
    return result_;
}

- (NSUndoManager*)windowWillReturnUndoManager:(NSWindow*)window {
    void* result_ = windowDelegate_WindowWillReturnUndoManager([self goID], window);
    return (NSUndoManager*)result_;
}

- (BOOL)window:(NSWindow*)window shouldPopUpDocumentPathMenu:(NSMenu*)menu {
    bool result_ = windowDelegate_Window_ShouldPopUpDocumentPathMenu([self goID], window, menu);
    return result_;
}

- (void)window:(NSWindow*)window willEncodeRestorableState:(NSCoder*)state {
    windowDelegate_Window_WillEncodeRestorableState([self goID], window, state);
}

- (void)window:(NSWindow*)window didDecodeRestorableState:(NSCoder*)state {
    windowDelegate_Window_DidDecodeRestorableState([self goID], window, state);
}

- (NSSize)window:(NSWindow*)window willResizeForVersionBrowserWithMaxPreferredSize:(NSSize)maxPreferredFrameSize maxAllowedSize:(NSSize)maxAllowedFrameSize {
    CGSize result_ = windowDelegate_Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize([self goID], window, maxPreferredFrameSize, maxAllowedFrameSize);
    return result_;
}

- (void)windowWillEnterVersionBrowser:(NSNotification*)notification {
    windowDelegate_WindowWillEnterVersionBrowser([self goID], notification);
}

- (void)windowDidEnterVersionBrowser:(NSNotification*)notification {
    windowDelegate_WindowDidEnterVersionBrowser([self goID], notification);
}

- (void)windowWillExitVersionBrowser:(NSNotification*)notification {
    windowDelegate_WindowWillExitVersionBrowser([self goID], notification);
}

- (void)windowDidExitVersionBrowser:(NSNotification*)notification {
    windowDelegate_WindowDidExitVersionBrowser([self goID], notification);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return WindowDelegate_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteWindowDelegate([self goID]);
	[super dealloc];
}
@end

void* WrapWindowDelegate(long goID) {
    NSWindowDelegateAdaptor* adaptor = [[NSWindowDelegateAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
