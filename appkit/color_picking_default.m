#import <Appkit/Appkit.h>
#import "color_picking_default.h"
#import "_cgo_export.h"

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
