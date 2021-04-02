#import <AppKit/AppKit.h>
#import "_cgo_export.h"
#import "image_view.h"

void* ImageView_Image(void* ptr) {
	NSImageView* imageView = (NSImageView*)ptr;
	return [imageView image];
}

void ImageView_SetImage(void* ptr, void* image) {
	NSImageView* imageView = (NSImageView*)ptr;
	[imageView setImage:(NSImage*)image];
}

unsigned long ImageView_ImageFrameStyle(void* ptr) {
	NSImageView* imageView = (NSImageView*)ptr;
	return [imageView imageFrameStyle];
}

void ImageView_SetImageFrameStyle(void* ptr, unsigned long imageFrameStyle) {
	NSImageView* imageView = (NSImageView*)ptr;
	[imageView setImageFrameStyle:imageFrameStyle];
}

unsigned long ImageView_ImageAlignment(void* ptr) {
	NSImageView* imageView = (NSImageView*)ptr;
	return [imageView imageAlignment];
}

void ImageView_SetImageAlignment(void* ptr, unsigned long imageAlignment) {
	NSImageView* imageView = (NSImageView*)ptr;
	[imageView setImageAlignment:imageAlignment];
}

unsigned long ImageView_ImageScaling(void* ptr) {
	NSImageView* imageView = (NSImageView*)ptr;
	return [imageView imageScaling];
}

void ImageView_SetImageScaling(void* ptr, unsigned long imageScaling) {
	NSImageView* imageView = (NSImageView*)ptr;
	[imageView setImageScaling:imageScaling];
}

bool ImageView_Animates(void* ptr) {
	NSImageView* imageView = (NSImageView*)ptr;
	return [imageView animates];
}

void ImageView_SetAnimates(void* ptr, bool animates) {
	NSImageView* imageView = (NSImageView*)ptr;
	[imageView setAnimates:animates];
}

bool ImageView_IsEditable(void* ptr) {
	NSImageView* imageView = (NSImageView*)ptr;
	return [imageView isEditable];
}

void ImageView_SetEditable(void* ptr, bool editable) {
	NSImageView* imageView = (NSImageView*)ptr;
	[imageView setEditable:editable];
}

bool ImageView_AllowsCutCopyPaste(void* ptr) {
	NSImageView* imageView = (NSImageView*)ptr;
	return [imageView allowsCutCopyPaste];
}

void ImageView_SetAllowsCutCopyPaste(void* ptr, bool allowsCutCopyPaste) {
	NSImageView* imageView = (NSImageView*)ptr;
	[imageView setAllowsCutCopyPaste:allowsCutCopyPaste];
}

void* ImageView_ContentTintColor(void* ptr) {
	NSImageView* imageView = (NSImageView*)ptr;
	return [imageView contentTintColor];
}

void ImageView_SetContentTintColor(void* ptr, void* contentTintColor) {
	NSImageView* imageView = (NSImageView*)ptr;
	[imageView setContentTintColor:(NSColor*)contentTintColor];
}

void* ImageView_SymbolConfiguration(void* ptr) {
	NSImageView* imageView = (NSImageView*)ptr;
	return [imageView symbolConfiguration];
}

void ImageView_SetSymbolConfiguration(void* ptr, void* symbolConfiguration) {
	NSImageView* imageView = (NSImageView*)ptr;
	[imageView setSymbolConfiguration:(NSImageSymbolConfiguration*)symbolConfiguration];
}

void ImageView_ImageViewWithImage(void* image) {
	[NSImageView imageViewWithImage:(NSImage*)image];
}
