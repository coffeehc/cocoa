#import "color_picking.h"
#import <AppKit/NSColorPicking.h>
#import <Foundation/NSArray.h>
#import <Foundation/NSDictionary.h>
#import <_cgo_export.h>

@interface NSColorPickingCustomAdaptor : NSObject <NSColorPickingCustom>
@property (assign) uintptr_t goID;
@end

@implementation NSColorPickingCustomAdaptor


- (void)setColor:(NSColor*)newColor {
    colorPickingCustom_SetColor([self goID], newColor);
}

- (NSColorPanelMode)currentMode {
    int result_ = colorPickingCustom_CurrentMode([self goID]);
    return result_;
}

- (BOOL)supportsMode:(NSColorPanelMode)mode {
    bool result_ = colorPickingCustom_SupportsMode([self goID], mode);
    return result_;
}

- (NSView*)provideNewView:(BOOL)initialRequest {
    void* result_ = colorPickingCustom_ProvideNewView([self goID], initialRequest);
    return (NSView*)result_;
}

- (NSObject*)initWithPickerMask:(NSUInteger)mask colorPanel:(NSColorPanel*)owningColorPanel {
    void* result_ = colorPickingCustom_InitWithPickerMask_ColorPanel([self goID], mask, owningColorPanel);
    return (NSObject*)result_;
}

- (void)setMode:(NSColorPanelMode)mode {
    colorPickingCustom_SetMode([self goID], mode);
}

- (void)insertNewButtonImage:(NSImage*)newButtonImage in:(NSButtonCell*)buttonCell {
    colorPickingCustom_InsertNewButtonImage_In([self goID], newButtonImage, buttonCell);
}

- (NSImage*)provideNewButtonImage {
    void* result_ = colorPickingCustom_ProvideNewButtonImage([self goID]);
    return (NSImage*)result_;
}

- (NSSize)minContentSize {
    CGSize result_ = colorPickingCustom_MinContentSize([self goID]);
    return result_;
}

- (NSString*)buttonToolTip {
    void* result_ = colorPickingCustom_ButtonToolTip([self goID]);
    return (NSString*)result_;
}

- (void)alphaControlAddedOrRemoved:(id)sender {
    colorPickingCustom_AlphaControlAddedOrRemoved([self goID], sender);
}

- (void)viewSizeChanged:(id)sender {
    colorPickingCustom_ViewSizeChanged([self goID], sender);
}

- (void)attachColorList:(NSColorList*)colorList {
    colorPickingCustom_AttachColorList([self goID], colorList);
}

- (void)detachColorList:(NSColorList*)colorList {
    colorPickingCustom_DetachColorList([self goID], colorList);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return ColorPickingCustom_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapColorPickingCustom(uintptr_t goID) {
    NSColorPickingCustomAdaptor* adaptor = [[NSColorPickingCustomAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}

@interface NSColorPickingDefaultAdaptor : NSObject <NSColorPickingDefault>
@property (assign) uintptr_t goID;
@end

@implementation NSColorPickingDefaultAdaptor


- (NSObject*)initWithPickerMask:(NSUInteger)mask colorPanel:(NSColorPanel*)owningColorPanel {
    void* result_ = colorPickingDefault_InitWithPickerMask_ColorPanel([self goID], mask, owningColorPanel);
    return (NSObject*)result_;
}

- (void)setMode:(NSColorPanelMode)mode {
    colorPickingDefault_SetMode([self goID], mode);
}

- (void)insertNewButtonImage:(NSImage*)newButtonImage in:(NSButtonCell*)buttonCell {
    colorPickingDefault_InsertNewButtonImage_In([self goID], newButtonImage, buttonCell);
}

- (NSImage*)provideNewButtonImage {
    void* result_ = colorPickingDefault_ProvideNewButtonImage([self goID]);
    return (NSImage*)result_;
}

- (NSSize)minContentSize {
    CGSize result_ = colorPickingDefault_MinContentSize([self goID]);
    return result_;
}

- (NSString*)buttonToolTip {
    void* result_ = colorPickingDefault_ButtonToolTip([self goID]);
    return (NSString*)result_;
}

- (void)alphaControlAddedOrRemoved:(id)sender {
    colorPickingDefault_AlphaControlAddedOrRemoved([self goID], sender);
}

- (void)viewSizeChanged:(id)sender {
    colorPickingDefault_ViewSizeChanged([self goID], sender);
}

- (void)attachColorList:(NSColorList*)colorList {
    colorPickingDefault_AttachColorList([self goID], colorList);
}

- (void)detachColorList:(NSColorList*)colorList {
    colorPickingDefault_DetachColorList([self goID], colorList);
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return ColorPickingDefault_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteAppKitHandle([self goID]);
	[super dealloc];
}
@end

void* WrapColorPickingDefault(uintptr_t goID) {
    NSColorPickingDefaultAdaptor* adaptor = [[NSColorPickingDefaultAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
