#import <Foundation/Foundation.h>
#import "attributed_string.h"

void* C_AttributedString_Alloc() {
	return [NSAttributedString alloc];
}

void* C_NSAttributedString_Init(void* ptr) {
	NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
	NSAttributedString* result = [nSAttributedString init];
	return result;
}

void* C_NSAttributedString_String(void* ptr) {
	NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
	NSString* result = [nSAttributedString string];
	return result;
}

unsigned int C_NSAttributedString_Length(void* ptr) {
	NSAttributedString* nSAttributedString = (NSAttributedString*)ptr;
	unsigned int result = [nSAttributedString length];
	return result;
}
