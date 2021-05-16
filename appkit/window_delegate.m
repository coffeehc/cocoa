#import <Appkit/Appkit.h>
#import "window_delegate.h"
#import "_cgo_export.h"

@implementation NSWindowDelegateAdaptor


- (NSRect)window:(NSWindow*)window willPositionSheet:(NSWindow*)sheet usingRect:(NSRect)rect {
    CGRect result = Window_WillPositionSheet_UsingRect([self goID], window, sheet, rect);
    return result;
}

- (void)windowWillBeginSheet:(NSNotification*)notification {
    WindowWillBeginSheet([self goID], notification);
}

- (void)windowDidEndSheet:(NSNotification*)notification {
    WindowDidEndSheet([self goID], notification);
}

- (NSSize)windowWillResize:(NSWindow*)sender toSize:(NSSize)frameSize {
    CGSize result = WindowWillResize_ToSize([self goID], sender, frameSize);
    return result;
}

- (void)windowDidResize:(NSNotification*)notification {
    WindowDidResize([self goID], notification);
}

- (void)windowWillStartLiveResize:(NSNotification*)notification {
    WindowWillStartLiveResize([self goID], notification);
}

- (void)windowDidEndLiveResize:(NSNotification*)notification {
    WindowDidEndLiveResize([self goID], notification);
}

- (void)windowWillMiniaturize:(NSNotification*)notification {
    WindowWillMiniaturize([self goID], notification);
}

- (void)windowDidMiniaturize:(NSNotification*)notification {
    WindowDidMiniaturize([self goID], notification);
}

- (void)windowDidDeminiaturize:(NSNotification*)notification {
    WindowDidDeminiaturize([self goID], notification);
}

- (NSRect)windowWillUseStandardFrame:(NSWindow*)window defaultFrame:(NSRect)newFrame {
    CGRect result = WindowWillUseStandardFrame_DefaultFrame([self goID], window, newFrame);
    return result;
}

- (BOOL)windowShouldZoom:(NSWindow*)window toFrame:(NSRect)newFrame {
    bool result = WindowShouldZoom_ToFrame([self goID], window, newFrame);
    return result;
}

- (NSSize)window:(NSWindow*)window willUseFullScreenContentSize:(NSSize)proposedSize {
    CGSize result = Window_WillUseFullScreenContentSize([self goID], window, proposedSize);
    return result;
}

- (NSApplicationPresentationOptions)window:(NSWindow*)window willUseFullScreenPresentationOptions:(NSApplicationPresentationOptions)proposedOptions {
    unsigned int result = Window_WillUseFullScreenPresentationOptions([self goID], window, proposedOptions);
    return result;
}

- (void)windowWillEnterFullScreen:(NSNotification*)notification {
    WindowWillEnterFullScreen([self goID], notification);
}

- (void)windowDidEnterFullScreen:(NSNotification*)notification {
    WindowDidEnterFullScreen([self goID], notification);
}

- (void)windowWillExitFullScreen:(NSNotification*)notification {
    WindowWillExitFullScreen([self goID], notification);
}

- (void)windowDidExitFullScreen:(NSNotification*)notification {
    WindowDidExitFullScreen([self goID], notification);
}

- (void)window:(NSWindow*)window startCustomAnimationToEnterFullScreenWithDuration:(NSTimeInterval)duration {
    Window_StartCustomAnimationToEnterFullScreenWithDuration([self goID], window, duration);
}

- (void)window:(NSWindow*)window startCustomAnimationToEnterFullScreenOnScreen:(NSScreen*)screen withDuration:(NSTimeInterval)duration {
    Window_StartCustomAnimationToEnterFullScreenOnScreen_WithDuration([self goID], window, screen, duration);
}

- (void)windowDidFailToEnterFullScreen:(NSWindow*)window {
    WindowDidFailToEnterFullScreen([self goID], window);
}

- (void)window:(NSWindow*)window startCustomAnimationToExitFullScreenWithDuration:(NSTimeInterval)duration {
    Window_StartCustomAnimationToExitFullScreenWithDuration([self goID], window, duration);
}

- (void)windowDidFailToExitFullScreen:(NSWindow*)window {
    WindowDidFailToExitFullScreen([self goID], window);
}

- (void)windowWillMove:(NSNotification*)notification {
    WindowWillMove([self goID], notification);
}

- (void)windowDidMove:(NSNotification*)notification {
    WindowDidMove([self goID], notification);
}

- (void)windowDidChangeScreen:(NSNotification*)notification {
    WindowDidChangeScreen([self goID], notification);
}

- (void)windowDidChangeScreenProfile:(NSNotification*)notification {
    WindowDidChangeScreenProfile([self goID], notification);
}

- (void)windowDidChangeBackingProperties:(NSNotification*)notification {
    WindowDidChangeBackingProperties([self goID], notification);
}

- (BOOL)windowShouldClose:(NSWindow*)sender {
    bool result = WindowShouldClose([self goID], sender);
    return result;
}

- (void)windowWillClose:(NSNotification*)notification {
    WindowWillClose([self goID], notification);
}

- (void)windowDidBecomeKey:(NSNotification*)notification {
    WindowDidBecomeKey([self goID], notification);
}

- (void)windowDidResignKey:(NSNotification*)notification {
    WindowDidResignKey([self goID], notification);
}

- (void)windowDidBecomeMain:(NSNotification*)notification {
    WindowDidBecomeMain([self goID], notification);
}

- (void)windowDidResignMain:(NSNotification*)notification {
    WindowDidResignMain([self goID], notification);
}

- (id)windowWillReturnFieldEditor:(NSWindow*)sender toObject:(id)client {
    void* result = WindowWillReturnFieldEditor_ToObject([self goID], sender, client);
    return (id)result;
}

- (void)windowDidUpdate:(NSNotification*)notification {
    WindowDidUpdate([self goID], notification);
}

- (void)windowDidExpose:(NSNotification*)notification {
    WindowDidExpose([self goID], notification);
}

- (void)windowDidChangeOcclusionState:(NSNotification*)notification {
    WindowDidChangeOcclusionState([self goID], notification);
}

- (BOOL)window:(NSWindow*)window shouldDragDocumentWithEvent:(NSEvent*)event from:(NSPoint)dragImageLocation withPasteboard:(NSPasteboard*)pasteboard {
    bool result = Window_ShouldDragDocumentWithEvent_From_WithPasteboard([self goID], window, event, dragImageLocation, pasteboard);
    return result;
}

- (NSUndoManager*)windowWillReturnUndoManager:(NSWindow*)window {
    void* result = WindowWillReturnUndoManager([self goID], window);
    return (NSUndoManager*)result;
}

- (BOOL)window:(NSWindow*)window shouldPopUpDocumentPathMenu:(NSMenu*)menu {
    bool result = Window_ShouldPopUpDocumentPathMenu([self goID], window, menu);
    return result;
}

- (void)window:(NSWindow*)window willEncodeRestorableState:(NSCoder*)state {
    Window_WillEncodeRestorableState([self goID], window, state);
}

- (void)window:(NSWindow*)window didDecodeRestorableState:(NSCoder*)state {
    Window_DidDecodeRestorableState([self goID], window, state);
}

- (NSSize)window:(NSWindow*)window willResizeForVersionBrowserWithMaxPreferredSize:(NSSize)maxPreferredFrameSize maxAllowedSize:(NSSize)maxAllowedFrameSize {
    CGSize result = Window_WillResizeForVersionBrowserWithMaxPreferredSize_MaxAllowedSize([self goID], window, maxPreferredFrameSize, maxAllowedFrameSize);
    return result;
}

- (void)windowWillEnterVersionBrowser:(NSNotification*)notification {
    WindowWillEnterVersionBrowser([self goID], notification);
}

- (void)windowDidEnterVersionBrowser:(NSNotification*)notification {
    WindowDidEnterVersionBrowser([self goID], notification);
}

- (void)windowWillExitVersionBrowser:(NSNotification*)notification {
    WindowWillExitVersionBrowser([self goID], notification);
}

- (void)windowDidExitVersionBrowser:(NSNotification*)notification {
    WindowDidExitVersionBrowser([self goID], notification);
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
