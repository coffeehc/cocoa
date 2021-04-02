#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "image_symbol_configuration.h"

void* ImageSymbolConfiguration_ConfigurationWithScale(long scale) {
	return [NSImageSymbolConfiguration configurationWithScale:scale];
}

void* ImageSymbolConfiguration_ConfigurationWithPointSize(double pointSize, double weight) {
	return [NSImageSymbolConfiguration configurationWithPointSize:pointSize weight:weight];
}

void* ImageSymbolConfiguration_ConfigurationWithPointSizeAndScale(double pointSize, double weight, long scale) {
	return [NSImageSymbolConfiguration configurationWithPointSize:pointSize weight:weight scale:scale];
}

void* ImageSymbolConfiguration_ConfigurationWithTextStyle(const char* style) {
	return [NSImageSymbolConfiguration configurationWithTextStyle:[NSString stringWithUTF8String:style]];
}

void* ImageSymbolConfiguration_configurationWithTextStyleAndScale(const char* style, long scale) {
	return [NSImageSymbolConfiguration configurationWithTextStyle:[NSString stringWithUTF8String:style] scale:scale];
}
