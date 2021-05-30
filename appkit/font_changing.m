#import <Appkit/Appkit.h>
#import "font_changing.h"
#import "_cgo_export.h"

@implementation NSFontChangingAdaptor


- (void)changeFont:(NSFontManager*)sender {
    fontChanging_ChangeFont([self goID], sender);
}

- (NSFontPanelModeMask)validModesForFontPanel:(NSFontPanel*)fontPanel {
    unsigned int result_ = fontChanging_ValidModesForFontPanel([self goID], fontPanel);
    return result_;
}


- (BOOL)respondsToSelector:(SEL)aSelector {
	return FontChanging_RespondsTo([self goID], aSelector);
}

- (void)dealloc {
	deleteFontChanging([self goID]);
	[super dealloc];
}
@end

void* WrapFontChanging(long goID) {
    NSFontChangingAdaptor* adaptor = [[NSFontChangingAdaptor alloc] init];
    adaptor.goID = goID;
    return adaptor;
}
